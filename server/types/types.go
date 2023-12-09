package types


type Email struct {
	Id 		string `json:"id"`
	Name 	string `json:"name"`
	Subject string `json:"subject"`
	To 		string `json:"to"`
	Text 	string `json:"text"`
	Date 	string `json:"date"`
}