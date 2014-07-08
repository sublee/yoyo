package main

import "fmt"
import "net/http"
import "net/url"
import "github.com/droundy/goopt"
import "github.com/gin-gonic/gin"

var host = goopt.String(
	[]string{"-h", "--host"}, "0.0.0.0",
	"web server host address to bind.")

var port = goopt.Int(
	[]string{"-p", "--port"}, 4040,
	"web server listening port to bind.")

var yoAPIToken = goopt.String(
	[]string{"-y", "--yo"}, "",
	"the Yo API token.")

func main() {
	// parse argv
	goopt.Summary = "Runs YoYo web server."
	goopt.Parse(nil)
	// validate
	if *yoAPIToken == "" {
		panic("Yo API token required.")
	}
	// web server
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		query := c.Req.URL.Query()
		username := query["username"][0]
		form := url.Values{"api_token": {*yoAPIToken}, "username": {username}}
		http.PostForm("http://api.justyo.co/yo", form)
	})
	endpoint := fmt.Sprintf("%s:%d", *host, *port)
	r.Run(endpoint)
}
