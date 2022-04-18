package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	log.Fatal(http.ListenAndServe(":8081", nil))

	fmt.Println("Study buddy server is up!")
}
