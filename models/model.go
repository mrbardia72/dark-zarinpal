package models

type Logpay struct {
	Status  		int    `json:status`
	Authority  		string `json:authority`
	Amount 			int    `json:amount`
	Email 			string `json:email`
	Description 	string `json:description`
	Mobile 			string `json:mobile`
	Date 			string `json:date`
	Time 			string `json:time`
}

type Payment struct {
	Verified 	bool   `json:verified`
	Refid 		string `json:refid:`
	Status  	int    `json:status`
	Date 		string `json:date`
	Time 		string `json:time`
}