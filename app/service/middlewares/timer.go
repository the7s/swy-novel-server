package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Timer(c *gin.Context) {

	fmt.Println("request starting ...", time.Now())
	c.Next()
	fmt.Println("request end ", time.Now())
}
