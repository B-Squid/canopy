package service

import (
	m "canopy/internal/messages"
)

func Transciever() []m.Sorter {
	qwe := m.SdMessage{CreatedBy: "dfg", Id: 1}
	asd := m.RichMessage{RouteRule: "erter", Id: 2}
	zxc := m.DbMessage{Id: 3}
	myarr := []m.Sorter{zxc, qwe, asd, qwe}

	return myarr
}
