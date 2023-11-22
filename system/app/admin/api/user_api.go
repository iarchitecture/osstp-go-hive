package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

// @Tags UserApi
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body systemReq.Register true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {object} response.Response{data=systemRes.SysUserResponse,msg=string} "用户注册账号,返回包括用户信息"
// @Router /auth/register [post]

func (user UserApi) Register(c *gin.Context) {

}

func (user UserApi) Login(c *gin.Context) {

	fmt.Printf("login api")

}
