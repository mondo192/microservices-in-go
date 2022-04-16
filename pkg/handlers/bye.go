package handlers

import (
	"log"
	"net/http"
)

type Bye struct {
	logger *log.Logger
}

func NewBye(logger *log.Logger) *Bye {
	return &Bye{logger}
}

func (b *Bye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Byeee!"))
}
