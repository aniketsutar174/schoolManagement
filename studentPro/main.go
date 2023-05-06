package main

import (
	"fmt"
	"net/http"

	"github.com/aniketsutar174/schoolManagement/routes"
)

func main() {
	fmt.Println("mongoAPI")
	r := routes.Router()
	fmt.Println("Server is getting started...")
	fmt.Println("Listening at port 4000")
	http.ListenAndServe(":4000", r)

}
