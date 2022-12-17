package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgreSQL struct {
	pool *pgxpool.Pool
}

func NewPostgreSQL() (*PostgreSQL, error) {
	envDB := os.Getenv("DATABASE_URL")
	log.Printf("envDB : %v \n", envDB)
	sDB := "postgres://user:password@localhost:5432/dbname?sslmode=disable"
	if envDB == "" {
		envDB = sDB
	}
	pool, err := pgxpool.Connect(context.Background(), envDB)
	if err != nil {
		return nil, err
	}

	return &PostgreSQL{
		pool: pool,
	}, nil
}

func (p *PostgreSQL) Close() {
	p.pool.Close()
}

func (p *PostgreSQL) FindByNConst(nconst string) (Name, error) {
	query := `SELECT nconst, primary_name, birth_year, death_year FROM "names" WHERE nconst = $1`

	var res Name

	if err := p.pool.QueryRow(context.Background(), query, nconst).
		Scan(&res.NConst, &res.Name, &res.BirthYear, &res.DeathYear); err != nil {
		return Name{}, err
	}

	return res, nil
}
