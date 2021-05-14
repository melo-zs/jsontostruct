package main

import (
	"encoding/json"
	"fmt"
)

//用户
type User struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Age 	 int    `json:"age"`
	Birthday string `json:"birthday"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

//结构体转JSON
func structToJSON() string {

	user := User{
		UserName: "itbsl",
		NickName: "jack",
		Age:   	  18,
		Birthday: "2001-08-15",
		Sex:      "itbsl@gmail.com",
		Phone:    "176XXXX6688",
	}

	data, err := json.Marshal(user)

	if err != nil {
		fmt.Println("json.marshal failed, err:", err)
		return ""
	}

	return string(data)
}

//json转结构体
func JSONToStruct(str string) {

	var user User

	err := json.Unmarshal([]byte(str), &user)

	if err != nil {

		panic(err)
	}

	fmt.Println(user)

}

func main() {

	JSONToStruct(structToJSON())
}
