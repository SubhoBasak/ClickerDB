// This is the entrypoint of the database. The TCP server defined
// here, will help the user to interact with the database. By using
// multiple go routines this server can handle multiple clients at
// a same time.
package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net"
	"os"
)

var logger *log.Logger

func main() {
	// command line arguments
	host := flag.String("host", "127.0.0.1", "Enter the host address.\nEx - host 127.0.0.1")
	port := flag.String("port", "3778", "Enter the port number.\nEx - port 3778")
	logfile := flag.String("logfile", "", "Enter the log file name. If not provided then it will log into stdout.\nEx - db.log")

	flag.Parse()

	// logger setup
	if *logfile == "" {
		logger = log.New(os.Stdout, "GoDB:", log.LstdFlags|log.Lshortfile)
	} else {
		logf, err := os.OpenFile(*logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer logf.Close()
		logger = log.New(logf, "GoDB:", log.LstdFlags|log.Lshortfile)
	}

	// TCP server setup
	listener, err := net.Listen("tcp", *host+":"+*port)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	logger.Println("Server is running")

	for {
		conn, err := listener.Accept()
		if err == nil {
			// concurrently handle multiple connection over multiple goroutines
			go connHandler(conn)
		} else {
			logger.Println(err)
		}
	}
}

func connHandler(conn net.Conn) {
	defer conn.Close()
	logger.Println("Received new connection") // TODO: remove this log statement

	for {
		logger.Println("Waiting for new query") // TODO: remove this log statement
		buff, err := bufio.NewReader(conn).ReadString(';')
		if err != nil {
			if err == io.EOF {
				break
			}
			logger.Println(err)
			conn.Write([]byte("BAD QUERY"))
		}
		conn.Write([]byte(Execute(&buff)))
	}

	logger.Println("Closed one connection") // TODO: remove this log statement
}
