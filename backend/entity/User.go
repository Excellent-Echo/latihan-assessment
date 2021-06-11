package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type User struct {
	name		string	`json : "name"`
	address		string	`json : "address"`
	date_birth	string	`json : "date"`
	email		string	`json : "email"`
	password	string	`json : "password"`
	created_at	time.Time	`json : "created_at"`
	updated_at	time.Time	`json : "updated_at"`
}

func main() {
	jsonFileUser, err := os.Open("user.json")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	byteValue, err := ioutil.ReadAll(jsonFileUser)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var dataUser User