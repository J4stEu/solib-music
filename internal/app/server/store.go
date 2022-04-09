package server

import "github.com/J4stEu/solib/internal/app/store"

// ConfigureStore - database configuration
func (srv *Server) ConfigureStore() error {
	srv.logger.Info("Configuring store...")

	st := store.New()
	if err := st.Open(srv.config.DataBase); err != nil {
		return err
	}
	srv.store = st

	if srv.config.DataBase.DataBaseInit {
		if err := st.InitStore(srv.config.DataBase); err != nil {
			return err
		}
	}

	return nil
}
