package main

import (
	r "canopy/internal/repository"
	s "canopy/internal/service"
	"time"
)

func main() {
	for {
		time.Sleep(5 * time.Second)
		hjk := s.Transciever()
		r.Reciever(hjk)
	}
}
