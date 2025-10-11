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
	//var GlobalMutex sync.Mutex

	go l.Logman()

	for i := 1; i < 5; i++ {
		go s.Transciever(ch, i)
		go r.Reciever(ch, i)
	}

	for {
		time.Sleep(200 * time.Millisecond)
	}

}
