package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
)

type Algorithm int

const (
	LeastTraffic Algorithm = 0
	RoundRobin   Algorithm = 1
)

type Balancer interface {
	Listen()
	HandleConn(conn net.Conn)
	AddServer(*Server)
}

func NewBalancer(port string) Balancer {
	return &balancer{
		PORT:      port,
		Scheduler: &scheduler{},
	}
}

type balancer struct {
	PORT      string
	Conns     []*net.Conn
	Servers   []*Server
	Scheduler Scheduler
}

type Config struct {
	Algorithm
}

func (b *balancer) Listen() {
	l, err := net.Listen("tcp", b.PORT)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	log.Println("tcp server started listening")
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("New Connection", conn.RemoteAddr())
		go b.HandleConn(conn)
	}
}

func (b *balancer) HandleConn(conn net.Conn) {
	defer conn.Close()
	var buff bytes.Buffer

	req, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		conn.Close()
		return
	}
	server := b.Scheduler.GetLeastTraffic(b.Servers)
	url := fmt.Sprintf("http://%s:%s%s", server.URL, server.PORT, req.URL.Path)

	request, err := http.NewRequest(req.Method, url, req.Body)
	if err != nil {
		log.Println(err)
		return
	}
	server.ActiveConnections++
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}

	resp.Write(&buff)
	conn.Write(buff.Bytes())

	server.ActiveConnections--
}

func (b *balancer) AddServer(s *Server) {
	b.Servers = append(b.Servers, s)
}
