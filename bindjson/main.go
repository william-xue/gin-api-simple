package main

// ShouldBind学习，验证和绑定
import (
	"time"

	"github.com/gin-gonic/gin"
)

// 多个用,隔开，并列用|
type Boy struct {
	Name     string    `form:"name" binding:"required"`
	Age      int       `form:"age" binding:"required,min=0,max=120"`
	Address  string    `form:"address"`
	Brithday time.Time `form:"brithday" time_format:"2006-01-02"`
}

func handleBoy(c *gin.Context) {
	var boy Boy
	if err := c.ShouldBind(&boy); err == nil {
		//c.String(200, "%v", boy)
		c.JSON(200, gin.H{
			"name":     boy.Name,
			"age":      boy.Age,
			"address":  boy.Address,
			"brithday": boy.Brithday,
		})
	} else {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		//c.String(200, "%v", err)
	}
}
func main() {
	r := gin.Default()
	r.POST("/boy", handleBoy)
	r.Run(":2017")
}
