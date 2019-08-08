package main

import (
	"flag"
	"fmt"
	"net"
	"syscall"
)

const maxRead = 25


func main() {
	flag.Parse()

	if flag.NArg() != 2 {
		panic("usage: host port")
	}
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	listener := initServer(hostAndPort)
	for {
		conn, err := listener.Accept()
		checkError(err, "Accept: ")
		go connectionHandler(conn)
	}
}


func initServer(hostAndPort string) *net.TCPListener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError(err, "Resolving address: port failed: '" + hostAndPort + "'")
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError(err, "ListenTCP: ")
	println("Listening to: ", listener.Addr().String())
	return listener
}


func connectionHandler(conn net.Conn) {
	connForm := conn.RemoteAddr().String()
	println("Connect from: " + connForm)
	sayHello(conn)
	for {
		var ibuf []byte = make([]byte, maxRead+1)
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0
		switch err {

		case nil:
			handleMsg(length, err, ibuf)

		case syscall.EAGAIN:
			continue

		default:
			goto DISCONNECT
		}

	DISCONNECT:
		err1 := conn.Close()
		println("Closed connection: ", connForm)
		checkError(err1, "Close: ")
	}
}


func sayHello(to net.Conn) {
	obuf := []byte{'L', 'e', 't', '\'', ',', 'G', 'O', '!', '\n'}
	wrote, err := to.Write(obuf)
	checkError(err, "Write " + string(wrote) + " bytes.")
}


func handleMsg(length int, err error, msg []byte) {
	if length > 0 {
		print("<", length, ":")
		for i := 0; ; i++ {
			if msg[i] == 0 {
				break
			}
			fmt.Printf("%c", msg[i])
		}
		print(">")
	}
}


func checkError(err error, info string) {
	if err != nil {
		panic("ERROR: " + info + " " + err.Error())
	}
}

