package ordem

import (
	"database/sql"
	"fmt"
)

// Função auxiliar que retorna verdadeiro ou falso para o select enviado
func rowExists(query string, db *sql.DB, args ...interface{}) bool {
	var exists bool
	var aux int
	query = fmt.Sprintf("select case when exists (%s) then 1 else 0 end from dual", query)
	err := db.QueryRow(query, args...).Scan(&aux)
	if err != nil && err != sql.ErrNoRows {
		fmt.Printf("ERRO checando se existe '%s' %v", args, err)
	}
	if aux == 1 {
		exists = true
	} else {
		exists = false
	}
	return exists
}