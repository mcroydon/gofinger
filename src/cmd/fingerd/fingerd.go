package main

import (
	"bytes"
	"io"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// A map of username keys to user information values.
var users = map[string]string{
	"mcroydon": "Matt Croydon\r\nhttp://postneo.com\r\nhttp://github.com/mcroydon",
	"root":     "Root User",
}

// Flags used to modify default behavior.
var (
	address = flag.String("a", "127.0.0.1", "the IP address to listen on.")
)

// Ensure that we can listen on port 79.  Log a fatal error if we cannot.
func mustListen() net.Listener {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:79", *address))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Now accepting connections on port 79.")
	return l
}

// Begin accepting connections.
func acceptAndServe(l net.Listener) {
	for {
		cn, err := l.Accept()
		if err != nil {
			log.Println(err)
		}

		go func() {
			err := serve(cn, cn)
			defer cn.Close()
			if err != nil {
				log.Println(err)
			}
		}()
	}
}

// Serve an indivudual request inside a goroutine.
func serve(r io.Reader, w io.Writer) os.Error {
	// Buffer to read request in to.
	// TODO: Handle requests > 1024 bytes
	query := make([]byte, 1024)

	// Read from network
	_, err := r.Read(query)

	// Return an error if we encountered one.
	if err != nil {
		return err
	}

	// Serve the default report if the request is a blank one.
	if bytes.Equal(query[:2], []byte("\r\n")) {
		log.Println("Default report")
		w.Write([]byte("List of available users:\r\n"))
		for k, _ := range users {
			w.Write([]byte(k))
			w.Write([]byte("\r\n"))

		}

		// Serve either a user detail report or a not found message.
	} else {
		// TODO: Handle possible edge cases such as when we don't encounter \r\n.
		user := strings.Split(string(query), "\r\n")[0]
		userInfo, present := users[user]
		log.Printf("Requested user %s (%s)", user, present)
		if present {
			log.Printf("%s found.", user)
			w.Write([]byte("User info for "))
			w.Write([]byte(user))
			w.Write([]byte("\r\n"))
			w.Write([]byte(userInfo))
			w.Write([]byte("\r\n"))
		} else {
			log.Printf("%s not found.", user)
		}
	}

	// TODO: Handle write errors with _, err = w.Write()...

	// Return write error or nil on success
	return nil
}

// Main function that parses flags then calls mustListen and acceptAndServe.
func main() {
	flag.Parse()
	acceptAndServe(mustListen())
}
