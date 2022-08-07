package main

import (
	"github.com/gin-gonic/gin"
)

type Color struct {
	Name string  `json:"name"`
	Rgb  [3]byte `json:"rgb"`
	Hex  string  `json:"hex"`
}

var colors = []Color{
	{"red", [3]byte{255, 0, 0}, "FF000"},
	{"green", [3]byte{0, 255, 0}, "FF000"},
	{"blue", [3]byte{0, 0, 255}, "FF000"},
}

func main() {
	r := gin.Default()
	r.GET("/c/:color", GetColor)
	r.Run()
}

func GetColor(c *gin.Context) {
	var query string = c.Param("color")
	for i := 0; i < len(colors); i++ {
		if colors[i].Name == query {
			c.JSON(200, colors[i])
			return
		}
	}
	c.JSON(200, gin.H{})
}
