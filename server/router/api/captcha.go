package api

/*
Captcha 验证码 API
*/
import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"time"
)

var result = base64Captcha.NewMemoryStore(20240, 3*time.Minute)
var id string

func stringConfig() *base64Captcha.DriverString {
	stringType := &base64Captcha.DriverString{
		Height:          100,
		Width:           50,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine,
		Length:          5,
		Source:          "123456789qwertyuiopasdfghjklzxcvb",
		BgColor: &color.RGBA{
			R: 40,
			G: 30,
			B: 89,
			A: 29,
		},
		Fonts: nil,
	}
	return stringType
}

func createCode(ctx *gin.Context) {
	driver := stringConfig()
	// 生成验证码
	c := base64Captcha.NewCaptcha(driver, result)
	ID, b64s, _, _ := c.Generate()
	id = ID
	item, _ := c.Driver.DrawCaptcha(b64s)
	ctx.Set("captcha", item)
	ctx.Set("captcha_id", id)
	item.WriteTo(ctx.Writer)
}

func verifyCaptcha(id, answer string) bool {
	return result.Verify(id, answer, true)
}

func verify(c *gin.Context) {
	ids := c.DefaultQuery("captcha_id", "")
	if ids == "" {
		ids = id
	}
	answer := c.DefaultQuery("answer", "")
	if !verifyCaptcha(ids, answer) {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "Error",
		},
		)
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "Verified",
		},
		)
	}
}
