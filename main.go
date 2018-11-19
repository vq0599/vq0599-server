package main

import (
	"vq0599/routers"
)


func main() {
	router := routers.InitRouter()
	router.Run(":8180")
}
