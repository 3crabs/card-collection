package models

//Card карта - единица коллекционирования
type Card struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Picture string `json:"picture"`
}

//AppInfo информация о пиложении
type AppInfo struct {
	AppName    string `json:"app_name"`
	AppVersion string `json:"app_version"`
}