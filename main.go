package main

import (
	"dl-admin-go/routers"
	"fmt"
	"os"
)

func main() {
	r := routers.InitRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	fmt.Println("server listening at port: " + port)
	_ = r.Run(":" + port)
}