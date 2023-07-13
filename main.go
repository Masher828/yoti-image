package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"yoti/store"
)

func main() {

	store.LoadStore()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go ProcessInput(signalCh)
	<-signalCh
	err := store.StoreInFile()
	if err != nil {
		log.Println("Error while saving the file ", err.Error())
	}

}

func ProcessInput(signalCh chan os.Signal) {
	fmt.Println(`Welcome to the Key-Value Store! Commands:
- SET key value - Set a key-value pair
- GET key - Retreieve the value for a given key
- DELETE key - Delete a key-value pair
- EXIT - Termincate the program`)

	inputReader := bufio.NewReader(os.Stdin)
	for true {
		var input string
		var err error

		input, _ = inputReader.ReadString('\n')

		input = strings.TrimSuffix(input, "\n")
		line := strings.Split(input, " ")

		if len(line) == 0 {
			fmt.Println("Invalid Input. Please try again")
			continue
		}

		switch line[0] {
		case "SET":
			if len(line) < 3 {
				fmt.Println("Invalid Input. Please try again")
				continue
			}
			err = store.Add(line[1], line[2])
			if err != nil {
				log.Println("Error while adding the value to store ", err.Error())
				continue
			}
			fmt.Printf("Key %s set successfully\n", line[1])
			break
		case "GET":
			if len(line) < 2 {
				fmt.Println("Invalid Input. Please try again")
				continue
			}
			value, err := store.Get(line[1])
			if err != nil {
				log.Println("Error while fetching the value from store ", err.Error())
				continue
			}
			fmt.Printf("Value for key %s : %s\n", line[1], value)

			break
		case "DELETE":
			if len(line) < 2 {
				fmt.Println("Invalid Input. Please try again")
				continue
			}
			err = store.Delete(line[1])
			if err != nil {
				log.Println("Error while deleting the value from store ", err.Error())
				continue
			}
			fmt.Printf("Key %s delete successfully \n", line[1])
			break
		case "EXIT":
			fmt.Println("Goodbye!")
			break

		}
	}
	signalCh <- os.Kill
}
