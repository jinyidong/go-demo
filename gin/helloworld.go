package gin

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func App()  {
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	router:=gin.Default()
	
	router.GET("/", func(context *gin.Context) {
		//context.String(http.StatusOK,"hello world")
		context.JSON(http.StatusOK,gin.H{"say":"hello world"})
		//context.XML(200,gin.H{"say":"hello world"})
	})

	router.Run(":8080")
}

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

//生成证书
//go run C:\go\src\crypto\tls\generate_cert.go --host localhost
func HTTP2()  {
	logger := log.New(os.Stderr, "", 0)
	logger.Println("[WARNING] DON'T USE THE EMBED CERTS FROM THIS EXAMPLE IN PRODUCTION ENVIRONMENT, GENERATE YOUR OWN!")

	r := gin.Default()
	r.SetHTMLTemplate(html)

	r.GET("/welcome", func(c *gin.Context) {
		c.HTML(http.StatusOK, "https", gin.H{
			"status": "success",
		})
	})

	// Listen and Server in https://127.0.0.1:8080
	r.RunTLS(":8080", "./cert.pem", "./key.pem")
}
