package api

import "net/http"

func (s *Server) tokenEndpoint(w http.ResponseWriter, r *http.Request) {
	token, err := s.TokenGeneratorFunc()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(token))
}
