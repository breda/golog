package main

import (
	"fmt"
	"github.com/breda/golog/internal/server"
)

func main() {
	server := server.NewHTTPServer(":8000")
	fmt.Println(server.ListenAndServe())
}