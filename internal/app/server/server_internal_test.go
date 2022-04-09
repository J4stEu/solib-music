package server

//// TestApp_HandleRoot - test server_errors root endpoint
//func TestApp_HandleRoot(t *testing.T) {
//	app := New(config.DefaultConfiguration(), logrus.New())
//	rec := httptest.NewRecorder()
//	req, _ := http.NewRequest(http.MethodGet, "/", nil)
//	app.HandleRoot().ServeHTTP(rec, req)
//	assert.Equal(t, rec.Body.String(), "Root")
//}
