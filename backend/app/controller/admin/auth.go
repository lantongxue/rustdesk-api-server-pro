package admin

import (
	"rustdesk-api-server-pro/app/form/admin"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/helper/captcha"
	"rustdesk-api-server-pro/util"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type AuthController struct {
	basicController
	Cfg *config.ServerConfig
}

func (c *AuthController) PostAuthLogin() mvc.Result {
	var loginForm admin.LoginForm
	err := c.Ctx.ReadJSON(&loginForm)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	if !captcha.VerifyCode(loginForm.CaptchaId, loginForm.Code) {
		return c.Error(nil, "CaptchaError")
	}

	var user model.User
	get, err := c.Db.Where("username = ? and is_admin = 1", loginForm.Username).Get(&user)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	if !get {
		return c.Error(nil, "UserNotExists")
	}

	if !util.PasswordVerify(loginForm.Password, user.Password) {
		return c.Error(nil, "UsernameOrPasswordError")
	}

	// make other tokens expired
	_, _ = c.Db.Where("user_id = ? and status = 1 and is_admin = 1", user.Id).Cols("status").Update(&model.AuthToken{
		Status: 0,
	})

	signStr := strconv.Itoa(user.Id) + user.Username + time.Now().String()
	token := util.HmacSha256(signStr, c.Cfg.SignKey)
	expired := 2 * time.Hour // 2 hours

	authToken := &model.AuthToken{
		UserId:  user.Id,
		Token:   token,
		Expired: time.Now().Add(expired),
		IsAdmin: true,
		Status:  1,
	}

	_, err = c.Db.Insert(authToken)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	return c.Success(iris.Map{
		"token": token,
	}, "ok")
}

func (c *AuthController) GetAuthCaptcha() mvc.Result {
	id, img := captcha.CreateCaptcha()
	return c.Success(iris.Map{
		"id":  id,
		"img": img,
	}, "ok")
}
