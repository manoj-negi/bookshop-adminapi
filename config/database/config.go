package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	DB *pgxpool.Pool
}

// func NewPostgres(ctx context.Context, connString string)
func NewPostgres(ctx context.Context, connString string) (*Postgres, error) {
	db, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}
	return &Postgres{DB: db}, nil
}

func (p *Postgres) Ping(ctx context.Context) error {
	return p.DB.Ping(ctx)
}

func (p *Postgres) Close() error {
	defer p.DB.Close()
	return nil
}


// package db

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	_ "github.com/lib/pq"
// 	"github.com/jackc/pgx/v5/pgxpool"
// )

// type Postgres struct {
// 	DB *pgxpool.Pool
// }

// func NewPostgres(ctx context.Context, connString string) (*Postgres, error) {
// 	db, err := connectToDB(ctx, connString)
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to create db connection pool: %w", err)
// 	}

// 	return &Postgres{DB: db}, nil
// }

// func connectToDB(ctx context.Context, connString string) (*pgxpool.Pool, error) {
// 	db, err := pgxpool.New(ctx, connString)
// 	if err != nil {
// 		// Check for auth error
// 		var pgErr *pgconn.PgError
// 		if errors.As(err, &pgErr) {
// 			if pgErr.Code == "28P01" {
// 				return nil, fmt.Errorf("invalid username/password")
// 			}
// 		}
// 		return nil, err
// 	}

// 	return db, nil
// }

// func (p *Postgres) Ping(ctx context.Context) error {
// 	return p.DB.Ping(ctx)
// }

// func (p *Postgres) Close() error {
// 	defer p.DB.Close()
// 	return nil
// }
