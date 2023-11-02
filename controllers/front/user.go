package front

import (
	"osstp-go-hive/controllers"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) GetUserInfo(ctx *gin.Context) {
	controllers.SuccessResponse(ctx, 0, "success", "user info", 1)
}

// UserRegisterHandler @Tags USER
// @Summary 用户注册
// @Produce json
// @Accept json
// @Param data body service.UserService true "用户名, 密码"
// @Success 200 {object} serializer.ResponseUser "{"status":200,"data":{},"msg":"ok"}"
// @Failure 500  {object} serializer.ResponseUser "{"status":500,"data":{},"Msg":{},"Error":"error"}"
// @Router /user/register [post]
func (u UserController) UserRegisterHandle() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// var req types.UserServiceReq
		// if err := ctx.ShouldBind(&req); err == nil {
		// 	l := serv
		// }

		controllers.ErrorResponse(ctx, 500, "注册失败")

	}

}
