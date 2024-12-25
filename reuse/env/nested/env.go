package nested

import (
	"log"
	"os"
)

func init() {

	vars := []struct {
		name  string
		value string
	}{
		{"ENV_NESTED1", "ENV_NESTED1_VALUE"},
		{"ENV_NESTED2", "ENV_NESTED2_VALUE"},
	}

	for _, v := range vars {
		err := os.Setenv(v.name, v.value)
		if err != nil {
			log.Fatalf("failed to set env variable %q: %v", v.name, err)
		}
	}

}
