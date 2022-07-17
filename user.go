package main

import (
	adobe_api "adobeBoss/adobe-api"
	"fmt"
	"time"
)

type User struct {
	Email    string `json:"email"`
	RemoveAt int64  `json:"remove_at"` // Valid until, packages are 30, 60 90 months
}

func RemoveExpiredUsers(users *[]User) {
	changed := false
	for i, user := range *users {
		if user.RemoveAt > time.Now().Unix() {
			fmt.Printf("Removing expired user %v\n", user.Email)
			*users = append((*users)[:i], (*users)[i+1:]...)
			adobe_api.RemoveUser(user.Email, AuthInfo.Token)
			changed = true
		}
	}
	if changed {
		SaveUsers(users)
	}
}
