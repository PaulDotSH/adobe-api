package main

import (
	adobe_api "adobe/adobe-api"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST": //on post set the cookies and whatever
		mutex.Lock()
		w.Header().Set("Access-Control-Allow-Origin", "*")

		Package := KeyMap[r.Header.Get("key")]
		if Package == 0 {
			mutex.Unlock()
			return
		}

		var user User
		user.Email = r.Header.Get("email")
		user.RemoveAt = time.Now().Unix() + (24 * 60 * 60 * int64(Package))
		Users = append(Users, user)
		err := adobe_api.AddUser(user.Email, AuthInfo.Token)
		if err != nil {
			fmt.Println("Error while adding user", user.Email, err.Error())
			w.Write([]byte("Backend error"))
			mutex.Unlock()
			return
		}
		err = adobe_api.AddAllApps(user.Email, AuthInfo.Token)
		if err != nil {
			fmt.Println("Error while adding user apps to user", user.Email, err.Error())
			w.Write([]byte("Backend error"))
			mutex.Unlock()
			return
		}
		fmt.Println("Successfully added user", user.Email, "will remove at", user.RemoveAt)
		delete(KeyMap, r.Header.Get("key"))
		SaveKeys(&KeyMap)
		SaveUsers(&Users)
		w.Write([]byte("Ok"))
		mutex.Unlock()
	case "GET":

	}
}

var secret = "P00LAC0X"

func GenerateLicenses(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST": //on post set the cookies and whatever
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Header.Get("$ecret") != secret {
			w.Write([]byte("You naughty naughty"))
			return
		}

		dayAmount, e := strconv.Atoi(r.Header.Get("Day-Amount"))
		if dayAmount == 0 || e != nil {
			return
		}
		licensesAmount, e := strconv.Atoi(r.Header.Get("Licenses-Amount"))
		if licensesAmount == 0 || e != nil {
			return
		}

		mutex.Lock()
		for i := 0; i < licensesAmount; i++ {
			KeyMap[generateKey("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 16)] = dayAmount
		}

		fmt.Printf("Successfully generated %v keys with the duration of %v days\n", licensesAmount, dayAmount)
		SaveKeys(&KeyMap)
		w.Write([]byte("Ok"))
		mutex.Unlock()
	case "GET":

	}
}

func GetLicenses(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET": //on post set the cookies and whatever
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Header.Get("$ecret") != secret {
			w.Write([]byte("You naughty naughty"))
			return
		}

		mutex.Lock()

		_json, err := json.Marshal(KeyMap)
		if err != nil {
			w.Write([]byte("Bad shit happen"))
			mutex.Unlock()
			return
		}
		w.Write(_json)
		mutex.Unlock()
	}
}
