package controllers

import (
	"dreamon/controllers/msg_struct"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var form Login
	if err := c.ShouldBind(&form); err == nil {
		if form.Mail == "Rennbon" && form.Password == "123" {
			c.JSON(http.StatusOK,
				msg_struct.NewMsg(msg_struct.Success, "OK", "123"))
		} else {
			c.JSON(http.StatusOK,
				msg_struct.NewMsg(msg_struct.Error, "傻逼滚", "有多远滚多远"))
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func RegisterHandler(c *gin.Context) {
	var form Login
	if err := c.ShouldBind(&form); err == nil {
		if form.Mail == "Rennbon" && form.Password == "123" {
			c.JSON(http.StatusOK,
				msg_struct.NewMsg(msg_struct.Success, "OK", "123"))
		} else {
			c.JSON(http.StatusOK,
				msg_struct.NewMsg(msg_struct.Error, "傻逼滚", "有多远滚多远"))
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

//登录用实体
type Login struct {
	Mail     string `form:"mail" binding:"required"`
	Password string `form:"password"  binding:"required"`
}
type register struct {
	Mail     string `form:"mail" binding:"required"`
	Password string `form:"password"  binding:"required"`
}
