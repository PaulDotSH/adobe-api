package main

import (
	"fmt"
	"github.com/didip/tollbooth"
	"log"
	"net/http"
	"sync"
	"time"
)

func init() {
	KeyMap = ReadKeys()
	Users = ReadUsers()
	AuthInfo.GetToken(&mutex)
}

var mutex sync.Mutex

var limiter = tollbooth.NewLimiter(0.0033, nil)

func main() {
	go func() {
		for true {
			if AuthInfo.ValidUntil < time.Now().Unix() {
				fmt.Println("Token expired... grabbing new token")
				AuthInfo.GetToken(&mutex)
			}
			mutex.Lock()
			RemoveExpiredUsers(&Users)
			mutex.Unlock()
			time.Sleep(time.Minute)
		}
	}()

	//http.HandleFunc("/AddUser", AddUser)
	http.Handle("/AddUser", tollbooth.LimitFuncHandler(limiter, AddUser))
	http.HandleFunc("/GenerateLicenses", GenerateLicenses)
	http.HandleFunc("/GetLicenses", GetLicenses)

	err := http.ListenAndServe("127.0.0.1:1337", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
