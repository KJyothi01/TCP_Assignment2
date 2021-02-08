package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Please send message >> >> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")
		//message, _ := bufio.NewReader(c).ReadString('\n')
		//fmt.Print("->: " + message)
                if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
		if strings.TrimSpace(string(text)) == "ALL" {
			fmt.Println("Getting all previous data.....")

		}

		if strings.TrimSpace(string(text)) == "NEW" {
			fmt.Println("Enter user name....")
			user, _ := reader.ReadString('\n')
			fmt.Fprintf(c, user+"\n")

		}

		if strings.TrimSpace(string(text)) == "BRAKE" {
			fmt.Println("Do you want to stop communication?...")
			user, _ := reader.ReadString('\n')
			fmt.Fprintf(c, user+"\n")

		}

		if strings.TrimSpace(string(text)) == "FIRST" {
			fmt.Println("Whoes message do you want to see first?")
			user, _ := reader.ReadString('\n')
			fmt.Fprintf(c, user+"\n")

		}

		if strings.TrimSpace(string(text)) == "PRINT" {
			fmt.Println("Getting all previous data.....")

		}

	}
}