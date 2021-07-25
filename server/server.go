package server

import (
	"context"
	"encoding/gob"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-oauth2-server/dao"
	"go-oauth2-server/db"
	"go-oauth2-server/model"
	"go-oauth2-server/service"
	"go-oauth2-server/util/list"
	"go-oauth2-server/web"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	//设置session
	gob.Register(model.OauthUser{})
	store := cookie.NewStore([]byte("secret11111"))
	store.Options(sessions.Options{
		MaxAge: int(60 * 5), //session 5分钟有效
		Path:   "/",
	})
	router.Use(sessions.Sessions("loginSession", store))

	//设置中间件
	router.Use(panicMiddleware)
	router.Use(loginMiddleware)

	//初始化数据库
	db.InitDb()

	//初始化DAO
	userDao := &dao.OauthUserDaoImpl{}
	clientDetailDao := &dao.OauthClientDetailDaoImpl{}
	codeDao := &dao.OauthCodeDaoImpl{}
	accessRefreshTokenDao := &dao.AccessRefreshTokenDaoImpl{}

	//初始化service
	userService := service.NewOauthUserServiceImpl(userDao)
	authorizeService := service.NewOauthAuthorizeServiceImpl(clientDetailDao, codeDao, accessRefreshTokenDao)
	//注册路由
	web.NewLoginRoute(userService).RegisterRoutes(router)
	web.NewOauthAuthorizeRoutes(authorizeService).RegisterRoutes(router)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	//优雅关机
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Errorf("listen: %s", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	defer db.Close()
	logrus.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	logrus.Info("Server exiting")
}

func panicMiddleware(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(200, err)
		}
	}()
	c.Next()
}

//白名单不进行拦截
var whiteList = &[]string{"/oauth/login.html", "/oauth/executeLogin.html", "/oauth/access/token"}

func loginMiddleware(c *gin.Context) {
	session := sessions.Default(c)
	path := c.Request.URL.Path
	fmt.Printf("访问路径：%s\n", path)
	if list.Contain(whiteList, path) {
		c.Next()
		return
	}
	userSession := session.Get("user")
	if userSession == nil {
		c.Redirect(http.StatusFound, "/oauth/login.html?"+c.Request.URL.RawQuery)
		c.Abort()
	} else {
		c.Next()
	}
}
