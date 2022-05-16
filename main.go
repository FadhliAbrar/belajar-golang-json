package main

import (
	"encoding/json"
	"fmt"
)

// type Message struct {
// 	Name      string
// 	Age       int
// 	Something interface{}
// }

type Message string
type familyMember struct {
	Name    string
	Age     int
	Parents []string
	Message *Message
}

func main() {
	// m := Message{"Fadhli", 20, 100}
	// var data Message

	// b, err := json.Marshal(m)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err = json.Unmarshal(b, &data)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// switch v := data.Something.(type) {
	// case int:
	// 	fmt.Println("value:", v)
	// case float64:
	// 	fmt.Println("value f:", v)
	// case string:
	// 	fmt.Println("value:", v)
	// default:
	// 	fmt.Println("nothing here")
	// }

	var f familyMember
	b := []byte(`{"Name":"Fadhli","Age":20,"Parents":["E","B"]}`)
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
	fmt.Println(f.Message)
}
