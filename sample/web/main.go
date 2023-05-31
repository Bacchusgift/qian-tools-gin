package main

import (
	"github.com/gin-gonic/gin"
	"qian-tools-gin/sample/web/router"
)

func main() {
	engine := gin.Default()
	router.Register(engine)
	if err := engine.Run(); err != nil {
		return
	}
}
