package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("Listening on port :6379")

	l, err := net.Listen("tcp", ":6379") //starting tcp listener
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := l.Accept() //recieving requests
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		resp := NewResp(conn)
		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		_ = value

		writer := NewWriter(conn)
		writer.Write(Value{typ: "string", str: "OK"})
	}
}
