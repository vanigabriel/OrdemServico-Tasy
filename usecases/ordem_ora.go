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
func (r *RepoOracle) InsertOS(os *entity.OrdemServico) (string, error) {
	log.Println("Iniciando InsertOS")

	err := r.initDB()
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer r.closeDB()

	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return "", err
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
		return "", err
	}

	if ErroOut != "" {
		tx.Rollback()
		log.Println(ErroOut)
		return "", errors.New(ErroOut)
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return "", err
	}

	return NrOrdem, nil
}

// InsertAnexos recebe o número da OS, o nome do arquivo e o arquivo
func (r *RepoOracle) InsertAnexos(ordem string, filename string, file []byte) error {
	// Primeira coisa a ser feito é recuperar o caminho onde será salvo o arquivo
	err := r.initDB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer r.closeDB()

	var filepath string

	sqlS := `select a.vl_parametro
			from FUNCAO_PARAMETRO a
			where a.nr_sequencia = 8 and a.cd_funcao = 299`

	row := r.db.QueryRow(sqlS)
	err = row.Scan(&filepath)
	if err != nil {
		return err
	}

	filepath = filepath + "OS_" + ordem

	//Create path if not exists
	os.MkdirAll(filepath, os.ModePerm)

	//Create File
	f, err := os.OpenFile(filepath+"/"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.Write(file)
	if err != nil {
		log.Fatal(err)
	}

	// Insere referência no banco de dados
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	sqlP := `insert into MAN_ORDEM_SERV_ARQ (nr_sequencia, nr_seq_ordem, dt_atualizacao, nm_usuario, ds_arquivo, ie_anexar_email)
				select
					man_ordem_serv_arq_seq.nextval,
					:col1,
					sysdate,
					'Integrado Site',
					:col2,
					'S'
				from dual`

	_, err = tx.Exec(sqlP, ordem, filepath+"/"+filename)

	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
