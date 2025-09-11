package service

import (
	m "canopy/internal/messages"
)

func Transciever() []m.Sorter {
	sdStr1 := m.SdMessage{CreatedBy: "a", Id: 1}
	sdStr2 := m.SdMessage{CreatedBy: "b", Id: 2}
	rmStr1 := m.RichMessage{RouteRule: "x", Id: 3}
	rmStr2 := m.RichMessage{RouteRule: "y", Id: 4}
	dmStr1 := m.DbMessage{Id: 5}
	dmStr2 := m.DbMessage{Id: 6}
	myarr := []m.Sorter{sdStr1, sdStr2, rmStr1, rmStr2, dmStr1, dmStr2}

	return myarr
}
