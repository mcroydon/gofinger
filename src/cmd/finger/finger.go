package main

import (
	"fmt"
	"net"
	"io/ioutil"
)

func main() {
	// Create connection
	conn, _ := net.Dial("tcp", "localhost"+":79")

	// Close connection when we're done
	defer conn.Close()

	// Request default report (blank line <CR><LF>
	conn.Write([]byte("\r\n"))

	// Read response in to buffer
	mess, _ := ioutil.ReadAll(conn)

	// Print response
	fmt.Printf("%s", mess)

}
