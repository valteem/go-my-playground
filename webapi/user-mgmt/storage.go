package main

import "fmt"

var (
	credentials map[string]string
)

func initStorage() {
	credentials = make(map[string]string, 0)
}

func addUserCredentials(login string, password string) error {
	if _, ok := credentials[login]; ok {
		return fmt.Errorf("user %s already exists", login)
	}
	credentials[login] = password
	return nil
}

func findUserCredentials(login string) (string, error) {
	var storedPassword string
	if _, ok := credentials[login]; !ok {
		return storedPassword, fmt.Errorf("user %s not found", login)
	}
	return credentials[login], nil
}
