package main

//ok
import (
	"encoding/json"
	"fmt"
)

//Given an object or array obj, return a compact object.
//
//A compact object is the same as the original object, except with keys containing falsy values removed.
//	This operation applies to the object and any nested objects. Arrays are considered objects where the indices are keys. A value is considered falsy when Boolean(value) returns false.
//
//You may assume the obj is the output of JSON.parse. In other words, it is valid JSON.

func main() {

	p := &Person{}
	err := json.Unmarshal([]byte(obj), p)
	if err != nil {
		fmt.Println("Ошибка при чтении данных:", err)
		return
	}
	fmt.Println(fmt.Sprintf("%+v", *p), err)
}

var obj = `{
	"userId": "r",
	"userInfo": {
		"name": "Alex Smith",
		"email": "alex.smith@example.com",
		"isActive": true,
		"age": 29
		}
	}`

type Person struct {
	UserId int      `json:"userId"`
	UsInf  UserInfo `json:"userInfo"`
}
type UserInfo struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive bool   `json:"isActive"`
}
