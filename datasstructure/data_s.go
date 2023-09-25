package datasstructure

import "time"

type gq_user struct {
	id              string
	name            string
	age             int
	address         string
	document_type   string
	document_number int
}

type trans struct {
	id       string
	rrn      int
	amount   int
	currency string
	user_id  string
	good_id  string
	place    string
	t_time   time.Time
}

type good struct {
	id             string
	name           string
	price          int
	currency       string
	country_origin string
}
