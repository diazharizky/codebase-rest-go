package server

import (
	"context"
	"net"
	"net/http"
	"sync"
)

type server struct {
	svr *http.Server
	wg  sync.WaitGroup
}

func NewServer(handler http.Handler) *server {
	return &server{
		svr: &http.Server{
			Addr:    net.JoinHostPort("", ""),
			Handler: handler,
		},
	}
}

func (s *server) Start() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	s.wg.Add(1)

	// log.Println("{\"label\":\"server-http\",\"level\":\"info\",\"msg\":\"server master started at pid " + strconv.Itoa(os.Getpid()) + "\",\"service\":\"" + Config.GetString("SERVER_NAME") + "\",\"time\":" + fmt.Sprint(time.Now().Format(time.RFC3339Nano)) + "\"}")
	go func() {
		// log.Println("{\"label\":\"server-http\",\"level\":\"info\",\"msg\":\"server worker started at pid " + strconv.Itoa(os.Getpid()) + " listening on " + net.JoinHostPort(ServerCfg.IP, ServerCfg.Port) + "\",\"service\":\"" + Config.GetString("SERVER_NAME") + "\",\"time\":" + fmt.Sprint(time.Now().Format(time.RFC3339Nano)) + "\"}")
		s.svr.ListenAndServe()

		s.wg.Done()
	}()
}
