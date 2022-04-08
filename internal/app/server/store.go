package server

import "github.com/J4stEu/solib/internal/app/store"

// ConfigureStore - database configuration
func (srv *Server) ConfigureStore() error {
	st := store.New()
	if err := st.Open(srv.config.DataBase); err != nil {
		return err
	}
	srv.store = st

	return nil
}
