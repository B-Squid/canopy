package repository

import (
	m "canopy/internal/messages"
	"fmt"
)

var SliceForId1 []m.SdMessage
var SliceForId2 []m.RichMessage
var SliceForId3 []m.DbMessage

func Reciever(s []m.Sorter) {

	for i := range s {
		switch v := s[i].(type) {
		case m.SdMessage:
			fmt.Println("Processing SdMessage, Id =", v.Id)
			SliceForId1 = append(SliceForId1, s[i].(m.SdMessage))
		case m.RichMessage:
			fmt.Println("Processing RichMessage, Id =", v.Id)
			SliceForId2 = append(SliceForId2, s[i].(m.RichMessage))
		case m.DbMessage:
			fmt.Println("Processing DbMessage, Id =", v.Id)
			SliceForId3 = append(SliceForId3, s[i].(m.DbMessage))
		default:
			fmt.Println("Unknown type :(")
		}
	}
}
