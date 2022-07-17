package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadKeys() Keys {
	file, err := os.ReadFile("keys.json")

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var keys Keys

	if err = json.Unmarshal(file, &keys); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return keys
}

func ReadUsers() []User {
	file, err := os.ReadFile("users.json")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var users []User

	if err = json.Unmarshal(file, &users); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return users
}

func SaveKeys(keys *Keys) {
	f, err := os.Create("keys.json")
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	_json, err := json.Marshal(*keys)
	if err != nil {
		panic(err.Error())
	}

	_, err = f.Write(_json)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully saved keys!")
}

func SaveUsers(users *[]User) {
	f, err := os.Create("users.json")
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	_json, err := json.Marshal(*users)
	if err != nil {
		panic(err.Error())
	}

	_, err = f.Write(_json)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully saved db!")
}
