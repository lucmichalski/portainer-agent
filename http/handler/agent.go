package handler

import (
	"log"
	"net/http"
	"os"

	"bitbucket.org/portainer/agent"
	"github.com/gorilla/mux"
)

type AgentHandler struct {
	*mux.Router
	logger         *log.Logger
	clusterService agent.ClusterService
}

func NewAgentHandler(cs agent.ClusterService) *AgentHandler {
	h := &AgentHandler{
		Router:         mux.NewRouter(),
		logger:         log.New(os.Stderr, "", log.LstdFlags),
		clusterService: cs,
	}

	h.Handle("/agents",
		http.HandlerFunc(h.handleGetAgents)).Methods(http.MethodGet)

	return h
}

func (handler *AgentHandler) handleGetAgents(w http.ResponseWriter, r *http.Request) {
	handler.clusterService.Members()
}