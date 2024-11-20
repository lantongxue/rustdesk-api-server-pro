package service

import (
	"fmt"
	"rustdesk-api-server-pro/app/form/api"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/util"
	"strconv"
	"strings"
	"time"

	"github.com/golang-module/carbon"
	"github.com/kataras/iris/v12"
	"github.com/pquerna/otp/totp"
)

type UserService struct {
	config *config.ServerConfig
}

var userService *UserService

func NewUserService() *UserService {
	// 单例模式
	if userService != nil {
		return userService
	}

	config := config.GetServerConfig()
	return &UserService{
		config: config,
	}
}

// 常规登录，使用账号和密码
func (service *UserService) Login(loginForm api.LoginForm) iris.Map {
	var user model.User
	get, err := db.DbEngine.Where("username = ?", loginForm.Username).Get(&user)
	if err != nil {
		return iris.Map{
			"error": err.Error(),
		}
	}

	if !get {
		return iris.Map{
			"error": "Username Or Password Error",
		}
	}

	if !util.PasswordVerify(loginForm.Password, user.Password) {
		return iris.Map{
			"error": "Username Or Password Error",
		}
	}

	// 如果是email_check，则发送验证邮件
	if user.LoginVerify == model.LOGIN_EMAIL_CHECK {
		if user.Email == "" {
			return iris.Map{
				"error": "No Email Address",
			}
		}

		// 发送邮件
		uuid := util.GetUUID()
		s := NewMailService()

		// 查询登录验证的邮件模板
		tpl, err := s.GetMailTemplateByType(model.MAIL_TPL_TYPE_LOGIN_VERIFY)
		if err != nil {
			return iris.Map{
				"error": err.Error(),
			}
		}

		code := util.RandomString(6)
		code = strings.ToUpper(code)

		expired := 10

		// 发送邮件
		err = s.Send(user.Id, tpl.Id, user.Email, uuid, map[string]string{
			"{$username}": user.Name,
			"{$code}":     code,
			"{$expired}":  strconv.Itoa(expired),
		})

		if err != nil {
			return iris.Map{
				"error": err.Error(),
			}
		}

		db.DbEngine.Insert(&model.VerifyCode{
			UserId:     user.Id,
			Type:       model.VC_TYPE_MAIL,
			Uuid:       uuid,
			Code:       code,
			RustdeskId: loginForm.RustdeskId,
			Expired:    carbon.Now(service.config.Db.TimeZone).AddMinutes(expired).ToStdTime(),
			Status:     model.VC_STATUS_UNUSED,
		})

		return iris.Map{
			"type":     model.LOGIN_EMAIL_CHECK,
			"tfa_type": model.LOGIN_EMAIL_CHECK,
			"secret":   uuid,
		}
	}

	// 2fa 验证
	if user.LoginVerify == model.LOGIN_TFA_CHECK {
		uuid := util.GetUUID()
		db.DbEngine.Insert(&model.VerifyCode{
			UserId:     user.Id,
			Type:       model.VC_TYPE_2FA,
			Uuid:       uuid,
			RustdeskId: loginForm.RustdeskId,
			Status:     model.VC_STATUS_UNUSED,
		})
		return iris.Map{
			"type":     model.LOGIN_EMAIL_CHECK,
			"tfa_type": model.LOGIN_TFA_CHECK,
			"secret":   uuid,
		}
	}

	// 这里可以根据表单里面的type决定是登录还是什么
	// type 取值范围：account mobile sms_code email_code tfa_code
	// account 就是直接登录

	token := service.GetLoginToken(loginForm, user.Id)
	return iris.Map{
		"access_token": token,
		"type":         model.LOGIN_ACCESS_TOKEN,
		"user": iris.Map{
			"name":     user.Name,
			"email":    user.Email,
			"note":     user.Note,
			"status":   user.Status,
			"is_admin": false,
		},
	}

}

// 获取登录TOKEN
func (service *UserService) GetLoginToken(loginForm api.LoginForm, userId int) string {
	// make other tokens expired
	_, _ = db.DbEngine.Where("user_id = ? and rustdesk_id = ? and status = 1 and is_admin = 0", userId, loginForm.RustdeskId).Cols("status").Update(&model.AuthToken{
		Status: 0,
	})

	signStr := fmt.Sprintf("%s_%s_%d_%s", loginForm.RustdeskId, loginForm.Uuid, userId, time.Now().String())
	token := util.HmacSha256(signStr, service.config.SignKey)
	expired := 90 * 24 * time.Hour // 3 months

	authToken := &model.AuthToken{
		UserId:     userId,
		RustdeskId: loginForm.RustdeskId,
		Uuid:       loginForm.Uuid,
		DeviceOs:   loginForm.DeviceInfo.OS,
		DeviceType: loginForm.DeviceInfo.Type,
		DeviceName: loginForm.DeviceInfo.Name,
		Token:      token,
		Expired:    time.Now().Add(expired),
		IsAdmin:    false,
		Status:     1,
	}

	_, _ = db.DbEngine.Insert(authToken)

	return token
}

// 使用邮箱验证码登录
func (service *UserService) LoginVerifyByEmailCode(loginForm api.LoginForm) iris.Map {
	var verifyCode model.VerifyCode
	get, err := db.DbEngine.Where("type = 1 and rustdesk_id = ? and uuid = ? and code = ? and status = 1", loginForm.RustdeskId, loginForm.Secret, loginForm.VerificationCode).Desc("id").Get(&verifyCode)
	if err != nil {
		return iris.Map{
			"error": err.Error(),
		}
	}
	if !get {
		return iris.Map{
			"error": "Verification Code Error",
		}
	}
	if verifyCode.Expired.Before(time.Now()) {
		verifyCode.Status = model.VC_STATUS_EXPIRED
		db.DbEngine.ID(verifyCode.Id).Update(&verifyCode)

		return iris.Map{
			"error": "Verification Code Error",
		}
	}
	if verifyCode.Code != loginForm.VerificationCode {
		return iris.Map{
			"error": "Verification Code Error",
		}
	}

	verifyCode.Status = model.VC_STATUS_USED
	db.DbEngine.ID(verifyCode.Id).Update(&verifyCode)

	var user model.User
	_, err = db.DbEngine.Where("id = ?", verifyCode.UserId).Get(&user)
	if err != nil {
		return iris.Map{
			"error": err.Error(),
		}
	}

	token := service.GetLoginToken(loginForm, verifyCode.UserId)

	return iris.Map{
		"access_token": token,
		"type":         model.LOGIN_ACCESS_TOKEN,
		"user": iris.Map{
			"name":     user.Name,
			"email":    user.Email,
			"note":     user.Note,
			"status":   user.Status,
			"is_admin": false,
		},
	}
}

// 使用2FA验证登录
func (service *UserService) LoginVerifyBy2FACode(loginForm api.LoginForm) iris.Map {

	var verifyCode model.VerifyCode
	get, err := db.DbEngine.Where("type = 3 and rustdesk_id = ? and uuid = ? and status = 1", loginForm.RustdeskId, loginForm.Secret).Desc("id").Get(&verifyCode)
	if err != nil {
		return iris.Map{
			"error": err.Error(),
		}
	}
	if !get {
		return iris.Map{
			"error": "Verification Code Error",
		}
	}

	var user model.User
	get, err = db.DbEngine.Where("id = ?", verifyCode.UserId).Get(&user)
	if err != nil {
		return iris.Map{
			"error": err.Error(),
		}
	}

	if !get {
		return iris.Map{
			"error": "Username Or Password Error",
		}
	}

	if !totp.Validate(loginForm.TfaCode, user.TwoFactorAuthSecret) {
		return iris.Map{
			"error": "Verification Code Error",
		}
	}

	verifyCode.Status = model.VC_STATUS_USED
	db.DbEngine.ID(verifyCode.Id).Update(&verifyCode)

	token := service.GetLoginToken(loginForm, user.Id)

	return iris.Map{
		"access_token": token,
		"type":         model.LOGIN_ACCESS_TOKEN,
		"user": iris.Map{
			"name":     user.Name,
			"email":    user.Email,
			"note":     user.Note,
			"status":   user.Status,
			"is_admin": false,
		},
	}
}
