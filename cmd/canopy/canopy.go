package main

import (
	"canopy/internal/messages"
	"fmt"
)

var sliceForId1 []messages.Sorter
var sliceForId2 []messages.Sorter
var sliceForId3 []messages.Sorter

func transciever() []messages.Sorter {
	qwe := messages.SdMessage{CreatedBy: "dfg", Id: 1}
	asd := messages.RichMessage{RouteRule: "erter", Id: 2}
	zxc := messages.DbMessage{Id: 3}
	myarr := []messages.Sorter{zxc, qwe, asd, qwe}

	return myarr
}

func reciever(s []messages.Sorter) {
	//for i := 0; i < 3; i++ {
	//fmt.Printf(strconv.Itoa(s[i].Getid()))
	//	fmt.Printf("id = %d\n", s[i].Getid())
	for i := range s {
		switch v := s[i].(type) {
		case messages.SdMessage:
			fmt.Println("Processing SdMessage, Id =", v.Id)
			sliceForId1 = append(sliceForId1, s[i])
		case messages.RichMessage:
			fmt.Println("Processing RichMessage, Id =", v.Id)
			sliceForId2 = append(sliceForId2, s[i])
		case messages.DbMessage:
			fmt.Println("Processing DbMessage, Id =", v.Id)
			sliceForId3 = append(sliceForId3, s[i])
		default:
			fmt.Println("Unknown type :(")

		}
	}
	fmt.Println(sliceForId1)
}

func main() {
	hjk := transciever()
	reciever(hjk)
}
