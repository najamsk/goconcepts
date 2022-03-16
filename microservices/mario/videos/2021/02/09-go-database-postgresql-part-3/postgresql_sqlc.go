package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
)

//go:generate sqlc generate

type PostgreSQLC struct {
	db *sql.DB
}

func NewPostgreSQLC() (*PostgreSQLC, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return &PostgreSQLC{
		db: db,
	}, nil
}

func (p *PostgreSQLC) Close() {
	p.db.Close()
}

func (p *PostgreSQLC) FindByNConst(nconst string) (Name, error) {
	fmt.Println("sqlc helper func to get row from db")
	fmt.Println(nconst)
	row, err := New(p.db).SelectName(context.Background(), sql.NullString{String: nconst, Valid: true})

	if err != nil {
		return Name{}, err
	}

	fmt.Println("row = ", row)
	return Name{
		NConst:    row.Nconst.String,
		Name:      row.PrimaryName.String,
		BirthYear: row.BirthYear.String,
		DeathYear: row.DeathYear.String,
	}, nil
}
