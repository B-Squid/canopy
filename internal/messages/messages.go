package messages

type SdMessage struct {
	createdBy        string // автор created_by.name
	createdTime      string // время создания created_time.value
	id               string // id. Добавить метод  возврата ссылки на заявки по id. Типа https://<server>/api/<id>
	priority         string // важность priority.name
	shortDescription string // краткое описание short_description
	subject          string // тема subject
}

type RichMessage struct {
	SdMessage
	Origin string // некий идентификатор источника сообщения
}

type DbMessage struct { // что писать в БД мы пока не знаем
}
