package handler

import (
	"context"
	"net/http"
	"time"
)

type Serv struct {
	serv *http.Server
}

func NewServer(handler http.Handler, opts ...Option) *Serv {
	s := &http.Server{
		Addr:              "0.0.0.0",
		Handler:           handler,
		ReadTimeout:       time.Second * 5,
		WriteTimeout:      time.Second * 5,
		ReadHeaderTimeout: time.Second,
		MaxHeaderBytes:    8 * 1024,
	}

	for _, opt := range opts {
		opt(s)
	}
	return &Serv{
		serv: s,
	}
}

func (s *Serv) Run() error {
	return s.serv.ListenAndServe()
}

func (s *Serv) Shutdown(ctx context.Context) error {
	return s.serv.Shutdown(ctx)
}

type Option func(srv *http.Server)

func WithTimeOutOption(t time.Duration) Option {
	return func(srv *http.Server) {
		if t != 0 {
			srv.ReadTimeout = t
		}
	}
}

func WithAddrOption(addr string) Option {
	return func(srv *http.Server) {
		if addr != "" {
			srv.Addr = addr
		}
	}
}
