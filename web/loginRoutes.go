package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-oauth2-server/model"
	"go-oauth2-server/service"
	"net/http"
	"net/url"
)

type UserLoginRoute struct {
	userService service.IoauthUserService
}

func NewLoginRoute(userService service.IoauthUserService) *UserLoginRoute {
	return &UserLoginRoute{userService: userService}
}

func (route *UserLoginRoute) RegisterRoutes(engine *gin.Engine) {
	route.loginHtml(engine)
	route.executeLoginHtml(engine)
	route.logoutHtml(engine)
}

func (route *UserLoginRoute) logoutHtml(engine *gin.Engine) {
	engine.GET("/oauth/logout.html", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.Redirect(http.StatusFound, "/oauth/login.html")
	})
}

func (route *UserLoginRoute) loginHtml(engine *gin.Engine) {
	engine.GET("/oauth/login.html", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{
			"query": c.Request.URL.RawQuery,
		})
	})
}

func (route *UserLoginRoute) executeLoginHtml(engine *gin.Engine) {
	engine.POST("/oauth/executeLogin.html", func(c *gin.Context) {
		name := c.PostForm("name")
		password := c.PostForm("password")
		logrus.Infof("登录入参name=%s,password=%s", name, password)
		if name == "" {
			c.HTML(200, "login.html", gin.H{
				"msg": "name 不能为空",
			})
			return
		}
		if password == "" {
			c.HTML(200, "login.html", gin.H{
				"msg": "password 不能为空",
			})
			return
		}
		user := &model.OauthUser{Name: name, Password: password}
		userFind, err := route.userService.Login(user)
		if err != nil {
			c.HTML(200, "login.html", gin.H{
				"msg": err.Error(),
			})
			return
		}
		session := sessions.Default(c)
		session.Set("user", userFind)
		session.Save()

		//重定向到授权页面
		query, _ := url.QueryUnescape(c.Request.URL.RawQuery)
		c.Redirect(http.StatusFound, "/oauth/authorize.html?"+query)
	})
}
