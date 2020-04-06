package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var user User
	file, err := os.Create("user-input.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	fmt.Println("What is your name?")
	fmt.Scan(&user.Name)
	fmt.Println("What is your age?")
	fmt.Scan(&user.Age)

	record := []string{user.Name, strconv.Itoa(user.Age)}
	error := writer.Write(record)
	checkError("User information could not be written \n", error)
	fmt.Printf("Your name is %s and you are age %v. \n", user.Name, user.Age)
	fmt.Printf("The user's information has been written to user-input.csv. \n")

}

type User struct {
	Name string
	Age  int
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
