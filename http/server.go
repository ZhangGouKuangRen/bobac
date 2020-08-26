package http

import (
	"idis/cache"
	"net/http"
)

type Server struct {
	cache.Cache
}

func New(c cache.Cache)*Server  {
	return &Server{c}
}


func (s *Server)Listen()  {
    http.Handle("/cache/", s.cacheHandler())
    http.Handle("/status", s.statusHandler())
    http.ListenAndServe(":6380", nil)
}

func (s *Server)cacheHandler()http.Handler  {
	return &cacheHandler{s}
}

func (s *Server)statusHandler()http.Handler  {
	return &statusHandler{s}
}