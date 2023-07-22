package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gtihub.com/popooq/Gophkeeper/server/internal/types"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Database struct {
	db  *sql.DB
	ctx context.Context
	dba string
}

var (
	registerEntry string = "INSERT INTO keeper (username, service, entry, metadata) VALUES ($1, $2, $3, $4)"
	getEntry      string = "SELECT entry, metadata FROM keeper WHERE username = $1 AND service = $2"
	updateEntry   string = "UPDATE keeper SET entry = $1, metadata = $2 WHERE username = $3 AND service = $4"
	deleteEntry   string = "DELETE FROM keeper WHERE username = $1 AND service = $2"
	registerUser  string = "INSERT INTO users (login, password_hash) VALUES ($1, $2) returning id"
	loginUser     string = "SELECT EXISTS(SELECT login, password_hash FROM users WHERE login = $1 AND password_hash = $2)"
)

func New(context context.Context, address string) *Database {
	if address == "" {
		log.Fatalf("cannot open DB there is no DB address %s", address)
		return nil
	}
	db, err := sql.Open("pgx", address)
	if err != nil {
		log.Fatalf("cannot open DB with address %s", address)
		return nil
	}
	return &Database{
		db:  db,
		ctx: context,
		dba: address,
	}
}

func (d *Database) Migrate() {
	driver, err := postgres.WithInstance(d.db, &postgres.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://server/internal/database/migrations",
		d.dba,
		driver,
	)
	if err != nil {
		log.Fatalln(err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalln(err)
	}
}
func (d *Database) Registration(username, password string) (types.User, error) {
	user := types.User{
		Login: username,
	}

	if d.db == nil {
		err := fmt.Errorf("you haven`t opened the database connection")
		return user, err
	}

	row := d.db.QueryRowContext(context.Background(), registerUser, username, password)
	if err := row.Scan(&user.ID); err != nil {
		return types.User{}, err
	}

	return user, nil
}

func (d *Database) Login(username, password string) bool {
	if d.db == nil {
		log.Println("you haven`t opened the database connection")
		return false
	}
	var exist bool
	tx, err := d.db.BeginTx(d.ctx, nil)
	if err != nil {
		log.Printf("error while creating tx %s", err)
		return false
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(d.ctx, loginUser)
	if err != nil {
		log.Printf("error while creating stmt %s", err)
		return false
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(d.ctx, username, password)
	err = row.Scan(&exist)
	if errors.Is(err, sql.ErrNoRows) {
		log.Println(err)
		return exist
	}
	if err != nil {
		log.Println(err)
		return exist
	}
	return exist

}

func (d *Database) NewEntry(entry types.Entry) error {
	if d.db == nil {
		err := fmt.Errorf("you haven`t opened the database connection")
		return err
	}

	tx, err := d.db.BeginTx(d.ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(d.ctx, registerEntry)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(d.ctx, entry.User, entry.Service, entry.Entry, entry.Metadata)
	if err != nil {
		err = fmt.Errorf("exec: %w", err)
		return err
	}

	return tx.Commit()
}

func (d *Database) UpdateEntry(entry types.Entry) error {
	if d.db == nil {
		err := fmt.Errorf("you haven`t opened the database connection")

		return err
	}

	tx, err := d.db.BeginTx(d.ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(d.ctx, updateEntry)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(d.ctx, entry.Entry, entry.Metadata, entry.User, entry.Service)
	if err != nil {
		err = fmt.Errorf("exec: %w", err)
		return err
	}

	return tx.Commit()
}

func (d *Database) GetEntry(username, service string) (types.Entry, error) {
	var entry types.Entry

	if d.db == nil {
		err := fmt.Errorf("you haven`t opened the database connection")

		return entry, err
	}

	row := d.db.QueryRowContext(d.ctx, getEntry, username, service)

	err := row.Scan(&entry.Entry, &entry.Metadata)
	if err != nil {
		log.Printf("error in scanning %s", err)
		return entry, err
	}

	err = row.Err()
	if err != nil {
		return entry, err
	}
	return entry, nil
}

func (d *Database) DeleteEntry(username, service string) (int, error) {

	if d.db == nil {
		err := fmt.Errorf("you haven`t opened the database connection")
		return http.StatusInternalServerError, err
	}

	resp, err := d.db.ExecContext(d.ctx, deleteEntry, username, service)
	if err != nil {
		log.Printf("error in exec %s", err)
		return http.StatusInternalServerError, err
	}

	status, err := resp.RowsAffected()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if status != 1 {
		return http.StatusNotFound, err
	}

	return http.StatusOK, err
}
