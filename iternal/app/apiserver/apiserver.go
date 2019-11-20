package apiserver

// APIServer ...
type APIServer struct {
	config *Config
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

// Start ...
func (a *APIServer) Start() error {
	return nil
}
