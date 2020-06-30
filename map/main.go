package main

import "fmt"

func main() {
	mp := make(map[string]int, 2)
	mp["name"] = 12
	mp["name"] = 122
	mp["value1"] = 123
	mp["value2"] = 123
	fmt.Println(mp["name"])
}
