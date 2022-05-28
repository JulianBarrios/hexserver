package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	/*r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080*/
	mc := memcache.New("127.0.0.1:11211")

	//	mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})

	it, err := mc.Get("foo")
	if err == memcache.ErrNotStored {
		fmt.Print(err)
	}
	fmt.Println(string(it.Value))
}
