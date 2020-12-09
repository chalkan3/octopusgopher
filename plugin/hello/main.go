package main

import "fmt"

func Hello(v ...interface{}) (interface{}, error) {
	fmt.Println("hello from plugin")
	return "", nil
}

func Now(v ...interface{}) (interface{}, error) {
	for _, t := range v {
		fmt.Println(t)
	}
	return "teste", nil
}
