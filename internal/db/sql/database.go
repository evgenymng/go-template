package database

import (
	"context"
	"fmt"

	"app/internal/config"
	"app/internal/log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var C *Connection

type Connection struct {
	db *sqlx.DB
}

// Connects to the database using credentials provided in the config.
func Connect(cfg config.Config) (c *Connection, err error) {
	var db *sqlx.DB
	if db, err = sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
			cfg.Database.Username,
			cfg.Database.Password,
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.Name,
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
			log.L().Tag(log.TagSqlQuery).Error(err),
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
) (bool, error) {
	var count []int
	if err := c.db.SelectContext(ctx, &count, selectBookById, id); err != nil {
		log.S.Error(
			"Database query has failed",
			log.L().Tag(log.TagSqlQuery).Error(err),
		)
		return false, err
	}
	return count[0] != 0, nil
}

// Closes database connection.
func (c *Connection) Close() {
	c.db.Close()
}
