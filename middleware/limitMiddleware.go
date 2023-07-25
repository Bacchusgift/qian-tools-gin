package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type RedisConfig struct {
	Ip       *string
	Port     *string
	Password *string
}

var ch = make(chan int)

var i *int

// LimitMiddleware 限流中间件
// 单位是每多少时间单位多少次请求

var pool string

func LimitMiddleware(perSecond int) func(c *gin.Context) {
	return func(c *gin.Context) {
		if pool == "ok" {
			fmt.Println("不再call")
		} else {
			go openRedisPool(ch) // 启用goroutine从通道接收值
			ch <- 1
		}
		c.JSON(200, "中间件执行")
		c.Abort()
	}
}

func openRedisPool(ch chan int) {
	if i == nil {
		ret := <-ch
		pool = "ok"
		fmt.Println("接收成功", ret)
		i = &ret
	} else {
		fmt.Println("已经接受过一次了")
		<-ch
	}
}
