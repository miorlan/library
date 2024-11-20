package models

// swagger:model SongRequest
type SongRequest struct {
	// Название Group
	// example: Острые Перцы
	Group string `json:"band"`
	Song  string `json:"song"`
}

// swagger:model Song
type Song struct {
	// ID Group
	// example: 1
	ID int `json:"id"`
	// Название Group
	// example: Острые Перцы
	Group string `json:"band"`
	// Название Song
	// example: Хочу в сковородку
	Song string `json:"song"`
	// Дата релиза
	// example: 20.10.1999
	ReleaseDate string `json:"releaseDate"`
	// Text песни
	// example: Нарежьте меня и отпустите на масло
	Text string `json:"text"`
	// Ссылка на песню
	// example: www.example.com
	Link string `json:"link"`
}

// swagger:model SongDetails
type SongDetails struct {
	// Дата релиза
	// example: 20.10.1999
	ReleaseDate string `json:"releaseDate"`
	// Text песни
	// example: Нарежьте меня и отпустите на масло
	Text string `json:"text"`
	// Ссылка на песню
	// example: www.example.com
	Link string `json:"link"`
}
