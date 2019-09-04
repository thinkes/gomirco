package main

import (
	"fmt"
	"gomirco/service"
)

var appName = "mircoTest"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6767")
}
