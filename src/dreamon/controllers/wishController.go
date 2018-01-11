package controllers

import (
	"dreamon/business"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var wishService business.WishService

func AddWishHandler(c *gin.Context) {
	var form wish
	if err := c.ShouldBind(&form); err == nil {
		wish := wishService.AddWish(form.Title, form.Wish)
		fmt.Print(wish.Wish)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

type wish struct {
	Title string `form:"title" binding:"required"`
	Wish  string `form:"wish" binding:"required"`
}
