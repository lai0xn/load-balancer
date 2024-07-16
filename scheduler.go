package main

type Scheduler interface {
	IsActive()
	GetLeastTraffic(severs []*Server) *Server
	CheckServer()
}

type scheduler struct{}

func (s *scheduler) IsActive() {}

func (s *scheduler) GetLeastTraffic(servers []*Server) *Server {
	leastActive := servers[0]
	for _, server := range servers {
		if server.ActiveConnections < leastActive.ActiveConnections && server.IsAlive {
			leastActive = server
		}
	}
	return leastActive
}

func (s *scheduler) CheckServer() {}
