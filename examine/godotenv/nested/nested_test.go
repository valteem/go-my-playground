package nested

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestNestedLoad(t *testing.T) {

	// Doesn't work - env file is in parent folder
	//	err := godotenv.Load()
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatalf("failed to load env variables: %v", err)
	}

}
