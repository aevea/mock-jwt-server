package api

import "net/http"

func (s *Server) jwkEndpoint(w http.ResponseWriter, r *http.Request) {
	token, err := s.PrivateKeyFunc()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(token))
}
