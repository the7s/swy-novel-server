package main

import (
	"fmt"
	"swy-novel-server/routers"
)

func main() {

	r := routers.SetupRouters()
	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("err : %s", err)
		return
	}
}
