package middleware

import (
	"fmt"
	"gin-casbin/global"
	"gin-casbin/utils/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
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
		//检查是否合法
		claims, tCode := GetClaims(checkToken[1])
		if tCode == response.ERROR {
			response.ResultWithCode(response.TOKEN_ERROR, context)
			return
		}
		//检查是否过期
		if time.Now().Unix() > claims.ExpiresAt {
			response.ResultWithCode(response.TOKEN_TIMEOUT, context)
			return
		}
		context.Set("uid", claims.Uid)
		context.Next()
	}
}

type MyClaims struct {
	Uid         int    `json:"uid"`
	AuthorityId string `json:"authorityId"`
	jwt.StandardClaims
}

// CreateToken SetToken 生成token
func CreateToken(uid int, AuthorityId string) (string, int) {
	var JwyKey = []byte(global.GVA_Set.JwtKey)
	// 过期时间
	expireTime := time.Now().Add(100 * time.Minute)
	SetClaim := MyClaims{
		Uid:         uid,
		AuthorityId: AuthorityId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "jimyag",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaim)
	token, err := reqClaim.SignedString(JwyKey)
	if err != nil {
		return "", response.ERROR
	}
	return token, response.SUCCESS

}

//// CheckToken 验证token
//func CheckToken(token string) (*MyClaims, int) {
//	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) { return JwyKey, nil })
//	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
//		return key, response.SUCCESS
//	} else {
//		return nil, response.ERROR
//	}
//}

func GetClaims(token string) (*MyClaims, int) {
	var JwyKey = []byte(global.GVA_Set.JwtKey)
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) { return JwyKey, nil })
	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key, response.SUCCESS
	} else {
		return nil, response.ERROR
	}
}
