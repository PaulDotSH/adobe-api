package adobe_api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func RemoveUser(email, token string) error {
	url := "https://usermanagement.adobe.io/v2/usermanagement/action/9DC2338D627015580A495EB9@AdobeOrg"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`[
  {
    "user" : "%v",
    "do" : [
      {
        "removeFromOrg" : {}
      }
    ]
  }
]`, email))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("x-api-key", "bec443e1f1194ed78eca083310a94f9e")
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))

	return nil
}

func AddAllApps(email, token string) error {
	url := "https://usermanagement.adobe.io/v2/usermanagement/action/9DC2338D627015580A495EB9@AdobeOrg"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`[
  {
    "user" : "%v",
    "do" : [
      {
        "add" : {
          "group" : [
            "Default All Apps plan - 100 GB configuration"
          ]
        }
      }
    ]
  }
]`, email))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("x-api-key", "bec443e1f1194ed78eca083310a94f9e")
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))

	return nil
}

func AddUser(email, token string) error {
	url := "https://usermanagement.adobe.io/v2/usermanagement/action/9DC2338D627015580A495EB9@AdobeOrg"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`[
  {
    "user" : "%v",
    "do" : [
      {
        "addAdobeID" : {
          "email" : "%v"
        }
      }
    ]
  }
]`, email, email))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("x-api-key", "bec443e1f1194ed78eca083310a94f9e")
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))

	return nil
}
