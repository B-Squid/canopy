package configs

import "fmt"

type AppConfig struct {
	userName            string // имя пользователя в ServiceDesk
	userSalt            string // случайный идентификатор, отличающий двух пользователей с одинаковым именем
	authtoken           string // токен доступа к ServiceDesk
	sdServer            string // адрес сервера ServiceDesk
	sdServerPort        string // порт сервера ServiceDesk
	apiPath             string // путь в api
	rmqServer           string // адрес сервера RabbitMQ
	rmqServerPort       string // порт сервера RabbitMQ
	pgServer            string // адрес сервера PosgtreSQL
	pgServerPort        string // порт сервера PosgtreSQL
	pollingEnabled      bool
	pollingFrequencySec int
	rowCount            int
	RouteRule           string
}

type MessageFilter struct {
	filterHeader string // фильтр сообщения по полю subject
	filterBody   string // фильтр сообщения по полю shortDescription
}

type configCheck struct { // проверка значений настроек
	regExpSdServer      string
	regExpSdServerPort  string
	regExpRmqServer     string
	regExpRmqServerPort string
	regExpPgServer      string
	regExpPgServerPort  string
}

func (ac *AppConfig) Setter() {
	//read config file
	ac.userName = "Some User"
	ac.userSalt = "k2m3n4n5"
	ac.authtoken = "2354ljkh25b45ik2t4r234goj4g5n924ngkl4gn69g"
	ac.sdServer = "https://servicedesk.domain"
}

func (ac AppConfig) Getter() {
	fmt.Println("User name: ", ac.userName)
}
