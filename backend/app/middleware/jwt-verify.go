package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"rustdesk-api-server-pro/app/model"
)

func JWTVerify(signKey string) iris.Handler {
	verifier := jwt.NewVerifier(jwt.HS256, []byte(signKey))
	verifier.WithDefaultBlocklist()
	verifyMiddleware := verifier.Verify(func() interface{} {
		return new(model.User)
	})
	return verifyMiddleware
}
