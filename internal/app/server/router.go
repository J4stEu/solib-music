package server

import (
	"net/http"
	"os"
	"path/filepath"
)

// spaHandler implements the http.Handler interface, so we can use it
// to respond to HTTP requests. The path to the static directory and
// path to the index file within that static directory are used to
// serve the SPA in the given static directory.
type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

// ConfigureRouter - router configuration
func (srv *Server) ConfigureRouter() {
	srv.logger.Info("Configuring router...")

	spa := spaHandler{staticPath: "./solib_frontend/dist", indexPath: "index.html"}
	srv.router.PathPrefix("/").Handler(spa)
	//srv.router.HandleFunc("/", srv.HandleRoot())
}

//// HandleRoot - root api
//func (srv *Server) HandleRoot() http.HandlerFunc {
//	return func(writer http.ResponseWriter, request *http.Request) {
//		srv.logger.Debug(fmt.Sprintf("%s: %s:%s%s", request.Method, srv.config.Server.ServerAddr, strconv.Itoa(int(srv.config.Server.ServerPort)), request.RequestURI))
//		_, err := io.WriteString(writer, "Root")
//		if err != nil {
//			srv.logger.Warn(err)
//		}
//	}
//}
