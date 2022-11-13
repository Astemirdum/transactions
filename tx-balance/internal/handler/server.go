package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/Astemirdum/transactions/tx-balance/config"
)

type Serv struct {
	serv *http.Server
}

func NewServer(handler http.Handler, addr string, cfg config.Server) *Serv {
	return &Serv{
		serv: &http.Server{
			Addr:              addr,
			Handler:           handler,
			ReadTimeout:       cfg.ReadTimeout,
			WriteTimeout:      time.Second * 5,
			ReadHeaderTimeout: time.Second,
			MaxHeaderBytes:    8 * 1024, // 8KiB
		},
	}
}

func (s *Serv) Run() error {
	return s.serv.ListenAndServe()
}

func (s *Serv) Shutdown(ctx context.Context) error {
	return s.serv.Shutdown(ctx)
}
