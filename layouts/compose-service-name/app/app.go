package app

import (
	"fmt"
	"os"
)

var (
	ConnStr string
	vars    = map[string]string{
		"POSTGRES_DB":       "",
		"POSTGRES_USER":     "",
		"POSTGRES_PASSWORD": "",
	}
)

func init() {

	for k := range vars {
		vars[k] = os.Getenv(k)
	}

	ConnStr = fmt.Sprintf("host = db_service port = 5432 dbname = %s user = %s password = %s",
		vars["POSTGRES_DB"],
		vars["POSTGRES_USER"],
		vars["POSTGRES_PASSWORD"],
	)

}
