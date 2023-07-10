package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gtihub.com/popooq/Gophkeeper/server/types"

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
