package src

import (
	"net/http"
	Handler "local.packages/handler"
)

type Server struct {
	voiceHandler *Handler.VoiceHandler
}

func NewServer(voiceHandler *Handler.VoiceHandler) *Server {
	server := new(Server)
	server.voiceHandler = voiceHandler

	return server
}

func (s Server) Start() {
	http.HandleFunc("/", s.voiceHandler.Index)
	http.HandleFunc("/register", s.voiceHandler.Register)
	http.ListenAndServe(":8080", nil)
}
