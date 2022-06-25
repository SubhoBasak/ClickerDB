// This is the entrypoint of the database. The TCP server defined
// here, will help the user to interact with the database. By using
// multiple go routines this server can handle multiple clients at
// a same time.
package main

import (
	"flag"
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
			go connHandler(&conn)
		} else {
			logger.Println(err)
		}
	}
}

func connHandler(conn *net.Conn) {
	defer (*conn).Close()
	for {
		result := Execute(conn)
		if result[0] == 'X' {
			// if any client don't follow the protocol connection
			// will be closed without any error message
			return
		}
		(*conn).Write([]byte(result))
	}
}
