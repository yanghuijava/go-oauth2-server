package web

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-oauth2-server/common"
	"go-oauth2-server/service"
)

type UserInfoRoute struct {
	userService service.IoauthUserService
}

func NewUserInfoRoute(userService service.IoauthUserService) *UserInfoRoute {
	return &UserInfoRoute{userService: userService}
}

func (route *UserInfoRoute) RegisterRoutes(engine *gin.Engine) {
	route.userInfo(engine)
}

func (route *UserInfoRoute) userInfo(engine *gin.Engine) {
	engine.GET("/user/userinfo", func(c *gin.Context) {
		token := c.Query("token")
		logrus.Infof("获取用户信息入参：%s", token)
		if token == "" {
			c.JSON(200, common.Failure(common.TOKEN_EMPTY))
			return
		}
		resp, err := route.userService.UserInfo(token)
		if err != nil {
			c.JSON(200, common.Failure(err.Err()))
		} else {
			c.JSON(200, resp)
		}
	})
}
