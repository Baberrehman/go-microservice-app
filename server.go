package main

import (
	"fmt"
	"mvcapp/middleware"
	"net/http"
)

func main() {
	router := middleware.LoadRoutes()

	err := http.ListenAndServe(":8811", router)
	fmt.Println(err)
}
