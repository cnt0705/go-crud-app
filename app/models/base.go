package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"todo-module/config"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)

	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`create table if not exists %s(
		id integer primary key autoincrement,
		uuid string not null unique,
		name string,
		email string,
		password string,
		created_at datetime)`, tableNameUser)
	Db.Exec(cmdU)

	cmdT := fmt.Sprintf(`create table if not exists %s(
		id integer primary key autoincrement,
		content text,
		user_id integer,
		created_at datetime)`, tableNameTodo)
	Db.Exec(cmdT)

	cmdS := fmt.Sprintf(`create table if not exists %s(
		id integer primary key autoincrement,
		uuid string not null unique,
		email string,
		user_id integer,
		created_at datetime)`, tableNameSession)
	Db.Exec(cmdS)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
