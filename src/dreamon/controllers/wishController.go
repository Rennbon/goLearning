package controllers

import (
	"dreamon/business"
	"dreamon/controllers/msg_struct"
	"net/http"

	"github.com/gin-gonic/gin"
)

var wishService business.WishService

func AddWishHandler(c *gin.Context) {
	var form wish
	if err := c.ShouldBind(&form); err == nil {
		wish := wishService.AddWish(form.Title, form.Wish)
		if wish != nil {
			c.JSON(http.StatusOK, msg_struct.NewMsg(msg_struct.Success, "OK", wish))
		} else {
			c.JSON(http.StatusOK,
				msg_struct.NewMsg(msg_struct.Error, "傻逼滚", "有多远滚多远"))
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
func GetWishHander(c *gin.Context) {
	var form getWish
	if err := c.ShouldBind(&form); err == nil {
		wish := wishService.GetWish(form.Id)
		if wish != nil {
			c.JSON(http.StatusOK, msg_struct.NewMsg(msg_struct.Success, "OK", wish))
		} else {
			c.JSON(http.StatusOK,
				msg_struct.NewMsg(msg_struct.Error, "傻逼滚", "有多远滚多远"))
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

type wish struct {
	Title string `form:"title" binding:"required"`
	Wish  string `form:"wish" binding:"required"`
}
type getWish struct {
	Id string `form:"id" binding:"required"`
}
