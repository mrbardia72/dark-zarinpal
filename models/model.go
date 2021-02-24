package models

type Logpay struct {
	Status  		int    `json:status`
	Authority  		string    `json:authority`
	Amount 			int64 `json:amount`
	Email 			string `json:email`
	Description 	string `json:description`
	Mobile 			int64 `json:mobile`
	Date 			string `json:date`
	Time 			string `json:time`
}

type Payment struct {
	Verified 	bool 	`json:verified`
	Refid 		int64 `	json:refid:`
	Status  	int    `json:status`
	Date 		string `json:date`
	Time 		string `json:time`
}