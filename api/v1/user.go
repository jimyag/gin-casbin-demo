package v1

import (
	"fmt"
	"gin-casbin/global"
	"gin-casbin/middleware"
	"gin-casbin/model"
	"gin-casbin/utils/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserApi struct {
}

func (userApi *UserApi) Register(context *gin.Context) {
	var user model.User
	_ = context.ShouldBindJSON(&user)
	err := model.CreateUser(&user)
	if err != response.SUCCESS {
		response.Result(response.ERROR, nil, "创建用户失败", context)
		context.Abort()
		return
	}
	user, _ = model.GetUserByUserName(user.UserName)
	data := make(map[string]interface{})
	data["token"], _ = middleware.CreateToken(int(user.ID), user.AuthorityId)
	data["user"] = user
	if user.AuthorityId == "2" {
		e, _ := global.GVA_CASBIN.AddPolicy(strconv.Itoa(int(user.ID)), fmt.Sprintf("/user/%d", user.ID), "GET")
		if !e {
			fmt.Println("添加权限失败")
		}
	}
	response.OkWithData(data, context)
}

func (userApi *UserApi) GetUserInfo(context *gin.Context) {
	uid, _ := strconv.Atoi(context.Param("uid"))
	if uid < 1 {
		response.ResultWithCode(response.ERROR, context)
		context.Abort()
		return
	}
	user, err := model.GetUserByUid(uint(uid))
	if err != response.SUCCESS {
		response.ResultWithCode(response.USER_NOT_EXIST, context)
		context.Abort()
		return
	}
	response.OkWithData(user, context)

}
