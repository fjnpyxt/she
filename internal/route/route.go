package route

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	C "github.com/zu1k/she/constant"
)

var (
	serverAddr = ""
)

// Start 启动 web api
func Start(addr string, secret string) {
	if serverAddr != "" {
		return
	}
	serverAddr = addr
	fmt.Println("API listening at: %s", addr)

	r := gin.Default()
	r.GET("/", hello)
	r.GET("/version", version)
	r.GET("/search", search)
	r.GET("/source", getSource)
	r.StaticFS("/ui", http.Dir("./dist"))

	err := r.Run(serverAddr)
	if err != nil {
		fmt.Println("API start Error: %s", err.Error())
		return
	}
}

func hello(c *gin.Context) {
	c.Redirect(http.StatusFound, "/ui/")
}

func version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": C.Version,
	})
}
