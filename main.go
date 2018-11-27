package main

import (
  "vq0599/routers"
  "vq0599/conf"
)


func main() {
  router := routers.InitRouter()
  router.Run(":" + conf.SERVER_PORT)
}
