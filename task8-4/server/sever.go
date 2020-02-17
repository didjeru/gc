package server

import (
	"fmt"
	"net"
	"sync"
)

//Server - Interface for server
type Server interface {
	Start() error
	Stop() error
}

//Config - config for server
type Config struct {
	Port string
}

type message struct {
	msg        string
	remoteAddr string
}

type server struct {
	cfg      *Config
	wg       sync.WaitGroup
	mu       sync.Mutex
	messages chan message
	clients  map[net.Conn]struct{}
	done     chan struct{}
}

// New - make server
func New(cfg *Config) Server {
	return &server{
		clients:  make(map[net.Conn]struct{}),
		cfg:      cfg,
		done:     make(chan struct{}),
		messages: make(chan message, 16),
	}
}

func (s *server) handleConn(conn net.Conn) error {
	fmt.Printf("Client with ip %q joined!\n", conn.RemoteAddr().String())
	data := make([]byte, 64)
	for {
		_, err := conn.Read(data)
		if err != nil {
			fmt.Printf("Client with ip %q left\n", conn.RemoteAddr().String())
			s.kickClient(conn)
			break
		}
		fmt.Printf("New message from ip %q: %s\n", conn.RemoteAddr().String(), string(data))
	}
	return nil
}

func (s *server) addClient(cn net.Conn) {
	s.mu.Lock()
	s.clients[cn] = struct{}{}
	s.mu.Unlock()
}

func (s *server) kickClient(cn net.Conn) {
	s.mu.Lock()
	delete(s.clients, cn)
	s.mu.Unlock()
}

func (s *server) getAllClients() map[net.Conn]struct{} {
	res := make(map[net.Conn]struct{})
	s.mu.Lock()
	for k := range s.clients {
		res[k] = struct{}{}
	}
	s.mu.Unlock()
	return res
}

func (s *server) startListener(srv net.Listener) {
	s.wg.Add(1)
	defer s.wg.Done()
	for {
		conn, err := srv.Accept()
		if err != nil {
			fmt.Println(err)
		}
		s.addClient(conn)
		go s.handleConn(conn)
	}
}

func (s *server) startShutdowner(srv net.Listener) {
	s.wg.Add(1)
	select {
	case <-s.done:
		err := srv.Close()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (s *server) runBroadcaster() {
	for {
		select {
		case msg := <-s.messages:
			clients := s.getAllClients()
			for k := range clients {
				if k.RemoteAddr().String() == msg.remoteAddr {
					continue
				}
				_, err := k.Write([]byte(fmt.Sprintf("%q: %s", msg.remoteAddr, msg.msg)))
				if err != nil {
					fmt.Printf("Failed to write to %q: %s", k.RemoteAddr().String(), err)
					s.kickClient(k)
				}
			}
		case <-s.done:
			return
		}
	}
}

func (s *server) Start() error {
	srv, err := net.Listen("tcp", s.cfg.Port)
	if err != nil {
		return err
	}
	go s.startListener(srv)
	go s.startShutdowner(srv)
	go s.runBroadcaster()
	return nil
}

func (s *server) Stop() error {
	close(s.done)
	s.wg.Wait()
	return nil
}
