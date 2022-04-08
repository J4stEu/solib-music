package server

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// ConfigureRouter - router configuration
func (srv *Server) ConfigureRouter() {
	srv.router.HandleFunc("/", srv.HandleRoot())
}

// HandleRoot - root api
func (srv *Server) HandleRoot() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		srv.logger.Debug(fmt.Sprintf("%s: %s:%s%s", request.Method, srv.config.Server.ServerAddr, strconv.Itoa(int(srv.config.Server.ServerPort)), request.RequestURI))
		_, err := io.WriteString(writer, "Root")
		if err != nil {
			srv.logger.Warn(err)
		}
	}
}
