package filters

type MessageFilter struct {
	filterHeader string // фильтр сообщения по полю subject
	filterBody   string // фильтр сообщения по полю shortDescription
}
