package usecase

import (
	"context"
	"fmt"
)

type HelloUseCase struct {
	ctx context.Context
}

func NewHelloUseCase() *HelloUseCase {
	return &HelloUseCase{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (hs *HelloUseCase) Start(ctx context.Context) {
	hs.ctx = ctx
}

// Greet returns a greeting for the given name
func (hs *HelloUseCase) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
