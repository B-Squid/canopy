package logger

import (
	r "canopy/internal/repository"
	"fmt"
	"time"
)

func Logman() {
	var cntSd int //cntRich, cntDb
	for {
		a := len(r.SliceForSd)
		if a != cntSd {
			//cntSd = r.SliceForSd
			fmt.Println("Count SliceForSD has changed: ", a)
			for i := cntSd; i < a; i++ {
				fmt.Println(r.SliceForSd[i].Id, r.SliceForSd[i].CreatedBy)
			}
			cntSd = a
		}

		time.Sleep(200 * time.Millisecond)
	}
}
