package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func App()  {
	router:=gin.Default()
	
	router.GET("/", func(context *gin.Context) {
		//context.String(http.StatusOK,"hello world")
		context.JSON(http.StatusOK,gin.H{"say":"hello world"})
		//context.XML(200,gin.H{"say":"hello world"})
	})

	router.Run(":8080")
}
