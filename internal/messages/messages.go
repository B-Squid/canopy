package messages

type Sorter interface {
	Getid() int
}

type SdMessage struct {
	CreatedBy string
	Id        int
}

type RichMessage struct {
	Id        int
	RouteRule string
}

type DbMessage struct {
	Id int
}

func (i SdMessage) Getid() int {
	return i.Id
}

func (i RichMessage) Getid() int {
	return i.Id
}

func (i DbMessage) Getid() int {
	return i.Id
}
