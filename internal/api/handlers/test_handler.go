package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.ezrpro.in/godemo/pkg"
)

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

// 测试服务正常运行
func (p *TestHandler) Ping(c *gin.Context) {
	pkg.Success(c, "ok")
}

func (p *TestHandler) Gracefully(c *gin.Context) {
	wait, _ := strconv.Atoi(c.Query("wait"))

	seconds := make([]int, wait)
	for i := 0; i < wait; i++ {
		seconds[i] = i
	}

	for _, sec := range seconds {
		fmt.Printf("sleep %d seconds\n", sec)
		time.Sleep(1 * time.Second)
	}

	msg := fmt.Sprintf("cost %d seconds,gracefully shutdown......\n", len(seconds))
	fmt.Println(msg)
	pkg.Success(c, msg)
}
