package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type connection struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DB       string `json:"database"`
	Schema   string `json:"schema"`
}

func GetConnStr(filename string) (string, error) {

	buf, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	var c connection
	err = json.Unmarshal(buf, &c)
	if err != nil {
		return "", err
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?currentSchema=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DB,
		c.Schema)

	return connStr, nil

}
