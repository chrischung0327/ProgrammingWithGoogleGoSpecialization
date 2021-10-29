package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var addr string
	fmt.Printf("Please Enter the name: ")
	fmt.Scan(&name)
	fmt.Printf("Please Enter the address: ")
	fmt.Scan(&addr)
	map_A := map[string]string{"name": name, "address": addr}
	barr, _ := json.Marshal(map_A)

	fmt.Println(string(barr))
}
