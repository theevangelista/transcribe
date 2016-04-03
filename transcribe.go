package main

import (
	"io/ioutil"
	"strconv"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/joaoevangelista/transcribe/consul"
	"github.com/joaoevangelista/transcribe/parser"
)

func init() {
	consul.Register()
}

func main() {
	r := gin.Default()
	r.PUT("/markdown", func(c *gin.Context) {
		body := c.Request.Body
		defer body.Close()
		b, err := ioutil.ReadAll(body)
		if err != nil {
			c.AbortWithError(500, err)
		}
		resp := parser.Markdown(b)
		c.Header("Content-Length", strconv.Itoa(utf8.RuneCountInString(resp)))
		c.String(200, "%s", resp)
	})

	r.Run()
}
