package repository

import (
	"database/sql"
	"iroko/config"
)

// SGDB - Define o tipo de banco de dados.
type SGDB string

//POSTGRES - Banco PostgreSQL.
const POSTGRES SGDB = "POSTGRES"

//DB - DB conection.
var DB *sql.DB

func conectionUrl() string {
}
