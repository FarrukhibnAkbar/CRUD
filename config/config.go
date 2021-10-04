package config

import "fmt"

var(
	db_host = "localhost"
	db_user = "admin"
	db_password = "12345"
	db_name = "crud"
	db_port = 5432
)

var DB_CONFIG = fmt.Sprintf(
	"host = %s user = %s password = %s dbname = %s port = %d", 
	db_host, db_user, db_password, db_name, db_port, 
)
