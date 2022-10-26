package main
import (
	"log"
	"os"
)

func main() {

	log.Println(os.Args[0])
	log.Println(os.Args[1:])

	for i, s := range os.Args[0:] {
		log.Println(i, s)
	}

}