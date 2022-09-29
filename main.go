package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// https://github.com/ariejanuarb/kabayan-backend-test

// prepare the empty struct
type Passanger struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

/*func majorityElement(arr int) int {
	sort.Ints(arr)
	return arr[len(arr)/2]
}
*/
func main() {
	// load the json file
	content, err := os.ReadFile("backend-titanic-test.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	// parse json file to []Passanger slice
	var passangers []Passanger
	err2 := json.Unmarshal(content, &passangers)
	if err2 != nil {
		fmt.Println("Error JSON Unmarshalling")
		fmt.Println(err2.Error())
	}

	// find most frequent age numbers
	for _, v := range passangers {
		fmt.Println(v.Age)
	}
	// print and sort apabethically the value based on most frequent numbers
}
