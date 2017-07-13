package main

import (
	"net"
	"log"
	"bufio"
	"strings"
	"fmt"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	request(conn)

	response(conn)

}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			method := strings.Fields(ln)[0] //Method
			uri := strings.Fields(ln)[1] //request uri

			fmt.Println("The request METHOD ", method)
			fmt.Println("The request URI ", uri)
		}

		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func response(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}