package main

import (
	"bufio"
	"fmt"
	"net"
)

type User struct {
	Username string
	IP       string
}

func handleConnection(conn net.Conn) {

	fmt.Println("Please enter Username:")
	userScanner := bufio.NewScanner(conn)

	ok := userScanner.Scan()

	if !ok {
		fmt.Println("Something went wrong here")
	}

	Username := userScanner.Text()
	remoteAddr := User{
		Username: Username,
		IP:       conn.RemoteAddr().String(),
	}

	fmt.Println("Client " + remoteAddr.Username + " has connected.")

	scanner := bufio.NewScanner(conn)

	for {
		ok := scanner.Scan()

		if !ok {
			break
		}

		data := scanner.Text()
		if data == "Accept-Language: en-US,en;q=0.9,lt;q=0.8,ar;q=0.7" {

			println("writing")
			_, err := conn.Write([]byte("HTTP/1.1 404 Not Found\nddd\n"))
			if err != nil {
				println(err.Error())
			}

		} else {
			println(remoteAddr.Username + ": " + data)
		}

		//handleMessage(scanner.Text(), conn)
	}

	fmt.Println("Client " + remoteAddr.Username + " disconnected.")
}

func main() {

	address := "192.168.1.7:90"
	listener, _ := net.Listen("tcp", address)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Some connection error: %s\n", err)
		}
		go handleConnection(conn)
	}
}
