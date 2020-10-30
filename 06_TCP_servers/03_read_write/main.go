package main

import (
	"log"
	"bufio"
	"fmt"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	// end the server after 10s.
	// err := conn.SetDeadline(time.Now().add(10 * time.Second))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "i heard you say: %s\n", ln)
	}

	defer conn.Close()

	fmt.Println("Code got here.")
}