package api

import (
	"server/global"
	"server/request"
	"server/response"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type BaseApi struct {
}

var store = base64Captcha.DefaultMemStore

func (baseApi *BaseApi) Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(
		global.Config.Captcha.Height,
		global.Config.Captcha.Width,
		global.Config.Captcha.Length,
		global.Config.Captcha.MaxSkew,
		global.Config.Captcha.DotCount,
	)

	captcha := base64Captcha.NewCaptcha(driver, store)

	id, b64s, _, err := captcha.Generate()

	if err != nil {
		global.Log.Error("Failed to generate captcha", zap.Error(err))
		response.FailWithMessage("Failed to generate captcha", c)
		return
	}

	response.OkWithData(response.Captcha{
		CaptchaID: id,
		PicPath:   b64s,
	}, c)
}

func (baseApi *BaseApi) SendEmailVerificationCode(c *gin.Context) {
	var req request.SendEmailVerificationCode
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if store.Verify(req.CaptchaID, req.Captcha, true) {
		err := baseService.SendEmailVerificationCode(c, req.Email)
		if err != nil {
			global.Log.Error("Failed to send email verification code", zap.Error(err))
			response.FailWithMessage("Failed to send email verification code", c)
			return
		}
		response.OkWithMessage("Successfully sent email verification code", c)
		return
	}
	response.FailWithMessage("Invalid captcha", c)
}

func (baseApi *BaseApi) QQLoginURL(c *gin.Context) {
	url := global.Config.QQ.QQLoginURL()
	response.OkWithData(url, c)
}
