package sql

import (
	"context"
	"fmt"

	"app/internal/config"
	"app/internal/log"
	"app/internal/model"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var C *Connection

type Connection struct {
	db *sqlx.DB
}

// Connects to the database using credentials provided in the config.
func Connect() (c *Connection, err error) {
	var db *sqlx.DB
	if db, err = sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
			config.C.Database.Username,
			config.C.Database.Password,
			config.C.Database.Host,
			config.C.Database.Port,
			config.C.Database.Name,
		),
	); err != nil {
		return
	}

	c = &Connection{
		db: db,
	}
	return
}

// Initializes database schema, if required tables do not exist.
func (c *Connection) InitSchema(ctx context.Context) error {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, schema); err != nil {
		log.S.Error(
			"Database query has failed, performing rollback",
			log.L().Error(err),
		)
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (c *Connection) GetBookById(
	ctx context.Context,
	id uuid.UUID,
) (model.Book, error) {
	var books []model.Book
	if err := c.db.SelectContext(ctx, &books, selectBookById, id); err != nil {
		log.S.Error("Database query has failed", log.L().Error(err))
		return model.Book{}, err
	}
	return books[0], nil
}

// Closes database connection.
func (c *Connection) Close() {
	c.db.Close()
}
