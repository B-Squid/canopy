package main

import (
	l "canopy/internal/logger"
	m "canopy/internal/messages"
	r "canopy/internal/repository"
	s "canopy/internal/service"
	"time"
)

func main() {
	var ch = make(chan []m.Sorter)
	go s.Transciever(ch)
	go r.Reciever(ch)

	for {
		go l.Logman()
		time.Sleep(200 * time.Millisecond)
	}
	//close(ch)
}
