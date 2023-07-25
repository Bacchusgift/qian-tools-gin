package main

import (
	"github.com/Bacchusgift/qian-tools-gin/sample/web/router"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	router.Register(engine)
	if err := engine.Run(); err != nil {
		return
	}
}
