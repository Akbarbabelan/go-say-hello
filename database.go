package database

var connection string

func init(){
	connection = "MSQL" 
}

func GetDatabase() string {
	retun connection
}