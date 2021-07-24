package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-oauth2-server/model"
	"go-oauth2-server/model/dto"
	"go-oauth2-server/service"
	"go-oauth2-server/util"
	"net/url"
	"strconv"
)

type OauthAuthorizeRoutes struct {
	authorizeService service.IoauthAuthorizeService
}

func NewOauthAuthorizeRoutes(authorizeService service.IoauthAuthorizeService) *OauthAuthorizeRoutes {
	return &OauthAuthorizeRoutes{
		authorizeService: authorizeService,
	}
}

func (route *OauthAuthorizeRoutes) RegisterRoutes(engine *gin.Engine) {
	route.authorizeHtml(engine)
	route.authorize(engine)
}

func (route *OauthAuthorizeRoutes) authorizeHtml(engine *gin.Engine) {
	engine.GET("/oauth/authorize.html", func(c *gin.Context) {
		//TODO 此处需要校验客户端申请的scope值
		queryMap := *util.ParseUrlQuery(c.Request.URL.RawQuery)
		authorizeRequest := &dto.OauthAuthorizeRequest{
			ClientId:     queryMap["client_id"],
			RedirectUri:  queryMap["redirect_uri"],
			ResponseType: queryMap["response_type"],
			Scope:        queryMap["scope"],
			State:        queryMap["state"],
		}
		err := route.authorizeService.CheckAuthorizeParams(authorizeRequest)
		if err != nil {
			c.HTML(401, "authorizeError.html", gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.HTML(200, "authorize.html", gin.H{
			"query": c.Request.URL.RawQuery,
		})
	})
}

func (route *OauthAuthorizeRoutes) authorize(engine *gin.Engine) {
	engine.POST("/oauth/authorize", func(c *gin.Context) {
		query, _ := url.QueryUnescape(c.Request.URL.RawQuery)
		queryMap := *util.ParseUrlQuery(query)
		action, _ := strconv.Atoi(c.PostForm("action"))
		if action == 1 { //用户拒绝授权
			c.HTML(401, "authorizeError.html", gin.H{
				"msg": "用户拒绝授权",
			})
			return
		}
		authorizeRequest := &dto.OauthAuthorizeRequest{
			ClientId:     queryMap["client_id"],
			RedirectUri:  queryMap["redirect_uri"],
			ResponseType: queryMap["response_type"],
			Scope:        queryMap["scope"],
			State:        queryMap["state"],
		}
		session := sessions.Default(c)
		code, err := route.authorizeService.AuthorizeCode(authorizeRequest, session.Get("user").(model.OauthUser))
		if err != nil {
			c.HTML(401, "authorizeError.html", gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.Redirect(302, authorizeRequest.RedirectUri+"?code="+code+"&state="+authorizeRequest.State)
	})
}
