package main

import (
	r "canopy/internal/repository"
	s "canopy/internal/service"
)

func main() {
	hjk := s.Transciever()
	r.Reciever(hjk)
}
