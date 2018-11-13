package main

func (s *Server) routes() {
	s.addRoute("/api/genres", s.handleGenres())
	s.addRoute("/api/client_id", s.handleClientID())
}
