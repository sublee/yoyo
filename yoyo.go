package main

import (
	"fmt"
	"github.com/droundy/goopt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

var host = goopt.String(
	[]string{"-h", "--host"}, "0.0.0.0", "web server host address to bind.")
var port = goopt.Int(
	[]string{"-p", "--port"}, 4040, "web server listening port to bind.")
var yoAPIToken = goopt.String(
	[]string{"-y", "--yo"}, "", "the Yo API token.")

func Parse() {
	goopt.Summary = "Runs YoYo web server."
	goopt.Parse(nil)
	if *yoAPIToken == "" {
		panic("Yo API token required.")
	}
}

func WebServer() *gin.Engine {
	www := gin.Default()
	www.GET("/", func(c *gin.Context) {
		query := c.Req.URL.Query()
		username := query["username"][0]
		form := url.Values{"api_token": {*yoAPIToken}, "username": {username}}
		http.PostForm("http://api.justyo.co/yo", form)
	})
	return www
}

func main() {
	Parse()
	endpoint := fmt.Sprintf("%s:%d", *host, *port)
	WebServer().Run(endpoint)
}
