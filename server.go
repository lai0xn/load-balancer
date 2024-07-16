package main

import (
	"net/url"
)

type Server struct {
	URL               string
	PORT              string
	ActiveConnections int
	IsAlive           bool
}

func BuildServer(rawURL string) (*Server, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	return &Server{
		URL:               parsed.Hostname(),
		PORT:              parsed.Port(),
		ActiveConnections: 0,
		IsAlive:           true,
	}, nil
}
