package main

import (
	"fmt"
	"swy-novel-server/routers"
)

func main() {

	r := routers.SetupRouters()
	err := r.Run("127.0.0.1:8888")
	if err != nil {
		fmt.Printf("err : %s", err)
		return
	}
}
