package main

import (
        "fmt"
        "github.com/kramer0815/rss-reader/service"
)

var appName = "accountservice"

func main() {
        fmt.Printf("Starting %v\n", appName)
        service.StartWebServer("6767")
}

