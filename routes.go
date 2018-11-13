package main

func (s *server) routes() {
	s.addRoute("/api/genres", s.handleGenres())
	s.addRoute("/api/client_id", s.handleClientID())
}
