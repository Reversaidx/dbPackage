package main

import (
	"bufio"
	"dbPackage/testdb"
	"fmt"
	"os"
	"strings"
)

func main() {
	var test testdb.Database
	err := test.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter Command put/get/delete/flush/close/stats ")
		command, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		switch command {
		case "put\n":
			fmt.Println("Enter key:value to put")
			keyvalue, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			s := strings.Split(strings.Trim(keyvalue, "\n"), ":")
			err = test.Put([]byte(s[0]), []byte(s[1]))
			if err != nil {
				fmt.Println("Put failed ")

			}

		case "get\n":
			fmt.Println("Enter key to get")
			key, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			key = strings.Trim(key, "\n")
			value, ok := test.Get([]byte(key))
			if ok != nil {
				fmt.Println("No such key")
			} else {
				fmt.Println(string(value))
			}
		case "delete\n":
			fmt.Println("Enter key to delete")
			key, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			key = strings.Trim(key, "\n")
			ok := test.Delete([]byte(key))
			if ok != nil {
				fmt.Println("No such key")
			} else {
				fmt.Println("Key was deleted ")
			}
		case "flush\n":
			test.Flush()
			fmt.Println("Done")
		case "stats\n":
			test.Stats()
			fmt.Println(test.Stats())
		case "close\n":
			test.Close()
			fmt.Println("Exit")
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
		}

	}
}
