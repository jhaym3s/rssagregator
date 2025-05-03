package main 

import (
    "fmt"
	"os"
	"log"
)

func main()  {
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}
	fmt.Println("PORT: %v", portString)
}