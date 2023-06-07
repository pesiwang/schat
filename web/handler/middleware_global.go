package handler

import (
	"schat/web/base"
	"schat/web/define"

	mapset "github.com/deckarep/golang-set"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

var authFreePathSet = mapset.NewThreadUnsafeSet()

func init() {
	// only add in init(), no need to be thread safe
	authFreePathSet.Add("/server/health_check")

	authFreePathSet.Add("/user/check_whether_register")
	authFreePathSet.Add("/user/login")
	authFreePathSet.Add("/user/register")
	authFreePathSet.Add("/user/get_auth_code")
}

func initMiddlewareGlobal() {
	base.GinRouter().Use(authMiddleware())
}

func authMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if authFreePathSet.Contains(path) {
			c.Next()
			return
		}

		tokenString := c.GetHeader("token")
		if tokenString == "" {
			base.GinSimpleResponse(c, define.RESULT_CODE_VERIFY_TOKEN_FAILED, "verify token failed: header token is empty")
			c.Abort()
			return
		}

		claims, err := base.ParseToken(tokenString)
		if err != nil {
			ve, ok := err.(*jwt.ValidationError)
			if ok && ve.Errors&jwt.ValidationErrorExpired != 0 {
				base.GinSimpleResponse(c, define.RESULT_CODE_TOKEN_EXPIRED, err.Error())
			} else {
				base.GinSimpleResponse(c, define.RESULT_CODE_VERIFY_TOKEN_FAILED, err.Error())
			}

			c.Abort()
			return
		}

		c.Set("uid", claims.Uid)

		c.Next()
	}

}
