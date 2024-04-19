package main

import (
	"backend/route"
	"backend/util"
	"strconv"
)

func main() {
	r := route.GetRouter()
	r.Run(":" + strconv.Itoa(util.GetProjectConfig().Server.Port))
}
