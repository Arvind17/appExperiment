package main

import (
	"fmt"
	cache2 "main.go/cache"

	//"github.com/appExperiment/cache"
)

//Initialize all the prerequisites like loggers and cache connections
func init() {
	cache2.Init()
}

func main() {
	// call Redis PING command to test connectivity
	err := cache2.Ping()
	if err != nil {
		fmt.Println(err)
	}

	// set demonstrates the redis SET command using a simple
	// string key:value pair
	err = cache2.Set()
	if err != nil {
		fmt.Println(err)
	}

	// set demonstrates the redis GET command
	err = cache2.Get()
	if err != nil {
		fmt.Println(err)
	}
}
