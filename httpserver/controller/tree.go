package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Tree(c *gin.Context) {
	fmt.Println("hello word !")
    c.JSON(200, "hello word !")
}

