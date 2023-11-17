package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type LoginApi struct{}

func (api LoginApi) Login(c *gin.Context) {

	fmt.Printf("login api")

	
}
