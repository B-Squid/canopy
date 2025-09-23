package service

import (
	m "canopy/internal/messages"
	"fmt"
	"math/rand"
)

func Transciever(ch chan []m.Sorter) {
	fmt.Print("Start transciever")
	for {
		sdStr := m.SdMessage{CreatedBy: "a", Id: randInt()}
		rmStr := m.RichMessage{RouteRule: "x", Id: randInt()}
		dmStr := m.DbMessage{Id: randInt()}
		myarr := []m.Sorter{sdStr, rmStr, dmStr}

		ch <- myarr
		//time.Sleep(5 * time.Second)
		//close(ch)
	}
}

func randInt() int {
	return 1 + rand.Intn(1000-1)
}
