package client

//Databases -
type Databases struct {
	ID        int
	Databases []Database
}

//Database -
type Database struct {
	database_id   string
	database_name string
	region        string
	tls           bool
	api_enabled   bool
}
