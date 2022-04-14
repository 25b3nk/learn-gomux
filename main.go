package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting of the program")
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))
	a.Run(":8010")
	fmt.Println("End of the program")
}
