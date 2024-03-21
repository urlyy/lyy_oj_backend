package main

import (
	"backend/route"
)

func main() {
	r := route.GetRouter()
	r.Run(":8090")
}
