package models

type Enviado struct {
	TableId    string `json:"table_id"`
	Data       string `json:"data"`
	HoraInicio string `json:"hora_inicio"`
	HoraFinal  string `json:"hora_final"`
}
