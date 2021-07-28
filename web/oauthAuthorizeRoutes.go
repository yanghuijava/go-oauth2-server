package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-oauth2-server/common"
	"go-oauth2-server/model"
	"go-oauth2-server/model/dto"
	"go-oauth2-server/service"
	"go-oauth2-server/util"
	"net/http"
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
	route.accessToken(engine)
}

func (route *OauthAuthorizeRoutes) authorizeHtml(engine *gin.Engine) {
	engine.GET("/oauth/authorize.html", func(c *gin.Context) {
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

//用户授权获取授权码
func (route *OauthAuthorizeRoutes) authorize(engine *gin.Engine) {
	engine.POST("/oauth/authorize", func(c *gin.Context) {
		//通过post提交表单后，url上传递的值被urlencode了，无法通过c.Query("xxx")获取，暂时自己解析
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
		user := session.Get("user").(model.OauthUser)
		if authorizeRequest.ResponseType == common.RESP_TYPE_CODE {
			code, err := route.authorizeService.AuthorizeCode(authorizeRequest, &user)
			if err != nil {
				c.HTML(401, "authorizeError.html", gin.H{
					"msg": err.Error(),
				})
				return
			}
			c.Redirect(302, authorizeRequest.RedirectUri+"?code="+code+"&state="+authorizeRequest.State)
		} else {
			resp, err := route.authorizeService.AuthorizeToken(authorizeRequest, &user)
			if err != nil {
				c.HTML(401, "authorizeError.html", gin.H{
					"msg": err.Error(),
				})
				return
			}
			c.Redirect(302, authorizeRequest.RedirectUri+"?accessToken="+resp.AccessToken+"&ExpiresIn="+strconv.Itoa(resp.ExpiresIn)+"&Openid="+resp.Openid+"&Scope="+resp.Scope+"&state="+authorizeRequest.State)
		}
	})
}

//获取accessToken
func (route *OauthAuthorizeRoutes) accessToken(engine *gin.Engine) {
	engine.POST("/oauth/access/token", func(c *gin.Context) {
		var request dto.AccessTokenReuqest
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(http.StatusOK, common.Failure(common.PARAMS_ERROR))
			return
		}
		request.BasicAuth = c.GetHeader("Authorization") //密码模式需要的参数
		request.OauthUser = nil                          //用户信息不允许直接从当前接口传入
		logrus.Infof("获取accessToken入参：%v", request)
		if request.GrantType == "" {
			c.JSON(http.StatusOK, common.Failure(common.PARAMS_ERROR))
			return
		}
		response, e := route.authorizeService.AccessToken(&request)
		if e != nil {
			c.JSON(http.StatusOK, common.Failure(e.Err()))
			return
		}
		c.JSON(http.StatusOK, response)
	})
}
