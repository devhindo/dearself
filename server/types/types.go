package types


type Email struct {
	Id 		int `json:"id"`
	Name 	string `json:"name"`
	CreatedAt string `json:"created_at"`
	Subject string `json:"subject"`
	To 		string `json:"to"`
	Text 	string `json:"text"`
	Date 	string `json:"date"`
}