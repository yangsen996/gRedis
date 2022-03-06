package main

import (
	"github.com/gin-gonic/gin"
	"goRedis/lib"
)

func main() {
	////fmt.Println(greids.NewStringOperation().Get("name").UnwrapOrDefault("no value"))
	//ret := greids.NewStringOperation().MGet("name","age","addr").Iter()
	//
	//for ret.HasNext(){
	// fmt.Println(ret.Next())
	//}
	//
	//greids.NewStringOperation().Set("as","aa",greids.WithExpire(10*time.Second))
	r := gin.Default()
	r.GET("/news/:id", func(c *gin.Context) {
		newsCache := lib.NewCache()
		defer lib.RelaseNew(newsCache)
		newid := c.Param("id")
		newsCache.DBGetter = lib.DbGetNews(newid)
		c.String(200, newsCache.GetCache("new"+newid).(string))
	})
}
