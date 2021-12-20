package middleware

import (
	"fmt"
	"gin-casbin/global"
	"gin-casbin/utils/response"
	"github.com/gin-gonic/gin"
	"strings"
)

func CasBin() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenHerder := context.Request.Header.Get("Authorization")
		fmt.Println(tokenHerder)
		//验证是否有token
		if tokenHerder == "" {
			response.ResultWithCode(response.NO_TOKEN, context)
			return
		}
		//验证格式
		checkToken := strings.SplitN(tokenHerder, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			response.ResultWithCode(response.TOEKEN_FORMAT_ERROR, context)
			return
		}
		claims, errCode := GetClaims(checkToken[1])
		if errCode != response.SUCCESS {
			response.ResultWithCode(response.TOKEN_ERROR, context)
			context.Abort()
			return
		}
		sub := claims.AuthorityId
		// casbin 的object
		obj := context.Request.URL.Path
		// casbin 的action
		act := context.Request.Method
		success, _ := global.GVA_CASBIN.Enforce(sub, obj, act)
		if success {
			context.Next()
		} else {
			response.ResultWithCode(response.PERMISSION_DENIED, context)
			context.Abort()
			return
		}

	}

}
