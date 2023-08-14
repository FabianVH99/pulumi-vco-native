package testdata

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func refreshJWT() string {
	url := fmt.Sprintf("https://%s/%s", os.Getenv("LAB_IAM_URL"), "/v1/oauth/jwt/refresh")

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(fmt.Errorf("failed to refresh JWT: %s", err))
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("LAB_TOKEN")))
	response, err := client.Do(req)
	if err != nil {
		panic(fmt.Errorf("failed to refresh JWT: %s", err))
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		panic(fmt.Errorf("failed to refresh JWT: %s", response.Status))
	}
	token, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return string(token)
}

func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}
