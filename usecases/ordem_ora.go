package ordem

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/vanigabriel/OrdemServico-Tasy/entity"
	_ "gopkg.in/goracle.v2"
)

// RepoOracle define struct
type RepoOracle struct {
	db *sql.DB
}

func NewRepository() *RepoOracle {
	var r RepoOracle
	return &r
}

func (r *RepoOracle) initDB() error {
	var dberr error

	dbuser := os.Getenv("dbuser")
	dbpassword := os.Getenv("dbpassword")
	dbname := os.Getenv("dbname")
	dbhost := os.Getenv("dbhost")
	dbport := os.Getenv("dbport")

	dbURI := fmt.Sprintf("%s/%s@%s:%s/%s", dbuser, dbpassword, dbhost, dbport, dbname)

	db, dberr := sql.Open("goracle", dbURI)

	r.db = db

	return dberr
}

func (r *RepoOracle) closeDB() {
	r.db.Close()
}

// InsertOS insere OS
func (r *RepoOracle) InsertOS(os *entity.OrdemServico) error {
	log.Println("Iniciando InsertOS")

	err := r.initDB()
	if err != nil {
		log.Println(err)
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	ctx := context.TODO()

	sqlP := `begin 
			HBS_OS_INSERE(:col1, :col2, :col3,:col4, :col5); 
			end;`
	var ErroOut string
	var NrOrdem string
	_, err = tx.ExecContext(ctx, sqlP, os.NrCPF, os.Descricao, os.Contato,
		sql.Out{Dest: &ErroOut}, sql.Out{Dest: &NrOrdem})

	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	if ErroOut != "" {
		tx.Rollback()
		log.Println(ErroOut)
		return errors.New(ErroOut)
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
