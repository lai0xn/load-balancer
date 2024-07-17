package main

type Scheduler interface {
	GetServer(severs []*Server, algo Algorithm) *Server
}

type scheduler struct {
	current int
}

func (s *scheduler) GetServer(servers []*Server, algo Algorithm) *Server {
	if algo == LeastTraffic {
		leastActive := servers[0]
		for _, server := range servers {
			if server.ActiveConnections < leastActive.ActiveConnections && server.IsAlive {
				leastActive = server
			}
		}
		return leastActive

	} else {
		index := (s.current + 1) % len(servers)
		s.current++
		return servers[index]
	}
}
