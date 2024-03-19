package webapp

// WebServer struct
type WebServer interface {
	RegisterRoutes()
	ListenAndServe(int) error
}
