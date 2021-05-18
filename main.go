package main

import (
	"brianwchen/ginessential/common"

	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()
	//defer db.

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}
