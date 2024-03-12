package gostart

import (
	"context"
	"net/http"
	"time"
)

// Эта структура, которую будем использовать для запуска http сервера. По сути, это небольшая абстракция над структурой сервера из пакета http
type Server struct {
	httServer *http.Server
}

// Структура имеет два метода: 1) Run для запуска сервера и 2) Shutdown для завершения работы
func (s *Server) Run(port string, handler http.Handler) error {
	s.httServer = &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s.httServer.ListenAndServe()
}

// метод для выхода из приложения
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httServer.Shutdown(ctx)

}
