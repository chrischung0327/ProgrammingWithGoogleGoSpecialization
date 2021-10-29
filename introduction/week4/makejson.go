// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// func main() {
// 	var name string
// 	var addr string
// 	fmt.Printf("Please Enter the name: ")
// 	fmt.Scan(&name)
// 	fmt.Printf("Please Enter the address: ")
// 	fmt.Scan(&addr)
// 	map_A := map[string]string{"name": name, "address": addr}
// 	barr, _ := json.Marshal(map_A)

// 	fmt.Println(string(barr))
// }
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func nameAddressMap(name string, address string) map[string]string {
	myNameAddressMap := map[string]string{
		"Name":    name,
		"Address": address,
	}
	return myNameAddressMap
}

func main() {

	fmt.Println("Enter your name please: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nameInput := scanner.Text()

	fmt.Println("Enter your address please: ")
	scanner.Scan()
	addressInput := scanner.Text()

	marshalledUserInput, err := json.Marshal(nameAddressMap(nameInput, addressInput))
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(marshalledUserInput))
}
