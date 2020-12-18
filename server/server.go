package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

const MAX_USERS = 100
const PORT = 6666

var gClients [MAX_USERS]Client

type Client struct {
	name        string
	ip          string
	sendMessage func()
	disconnect  func()
}

func clientConns(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	i := 0
	go func() {
		for {
			client, err := listener.Accept()
			if client == nil {
				fmt.Printf("couldn't accept: " + err.Error())
				continue
			}
			i++
			fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()
	return ch
}

func handleConn(client net.Conn) {
	b := bufio.NewReader(client)
	for {
		line, err := b.ReadBytes('\n')
		if err != nil { // EOF, or worse
			break
		}
		client.Write(line)
	}
}

func listen() {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	defer listener.Close()
	if err != nil {
		panic("couldn't start listening: " + err.Error())
	}
	conns := clientConns(listener)
	for {
		go handleConn(<-conns)
	}
}

func ranom(client Client) {
	client.sendMessage()
}

func main() {
	var clients []Client
	sendMessage := func() {
		fmt.Println("HELLO")
	}
	c := Client{"hello", "hello", sendMessage, func() {}}
	clients = append(clients, c)
	ranom(c)
	fmt.Println(c)
	listen()
}
