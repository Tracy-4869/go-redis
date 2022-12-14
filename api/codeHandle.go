package api

import (
	"github.com/gin-gonic/gin"
	"go-redis/redis"
	"go-redis/utils"
	"log"
	"net/http"
)

// post
// input: phone
// 生成验证码并存入redis
func CodeHandle(ctx *gin.Context) {
	phone := ctx.Query("phone")
	res := utils.InitResBody()
	// 获取验证码
	code := utils.GetRandomCode()
	log.Println("验证码：", code)
	// 保存到redis
	err := redis.SavePhoneCode(phone, code)
	if err != nil {
		res.Message = "redis错误"
		log.Println(err)
		ctx.JSON(http.StatusOK, res)
	}
	res.Message = "获取验证码成功"
	res.Data = code
	ctx.JSON(http.StatusOK, res)
}
