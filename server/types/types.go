package types


type Email struct {
	Name 	string `json:"name"`
	Subject string `json:"subject"`
	From 	string `json:"from"` 
	To 		string `json:"to"`
	Text 	string `json:"text"`
}