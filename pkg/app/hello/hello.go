package hello

import (
	"context"
	"fmt"
)

// HelloService struct
type HelloService struct {
	ctx context.Context
}

// NewHelloService creates a new App application struct
func NewHelloService() *HelloService {
	return &HelloService{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (hs *HelloService) Start(ctx context.Context) {
	hs.ctx = ctx
}

// Greet returns a greeting for the given name
func (hs *HelloService) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
