package persistent

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	dbutils "github.com/viniciusith/tinify/utils/db"
)

type PostgresAdapter struct {
	client *pgxpool.Pool
}

func NewPostgresAdapter() *PostgresAdapter {
	return &PostgresAdapter{}
}

func (p *PostgresAdapter) Connect(connectionString string) error {
	conn, err := pgxpool.New(context.Background(), connectionString)
	p.client = conn

	return err
}

func (p *PostgresAdapter) Close() error {
	if p.client != nil {
		p.client.Close()
		return nil
	}
	return fmt.Errorf("Tried to close postgres pool before it was opened")
}

// The data struct MUST be tagged using the format db:"keyname"
//
// name string `db:"name"`
func (p *PostgresAdapter) Insert(ctx context.Context, tableName string, data interface{}) error {
	mappedData := dbutils.StructToMap(data)

	columns := make([]string, 0, len(mappedData))
	placeholders := make([]string, 0, len(mappedData))
	var args []interface{}

	i := 1
	for key, value := range mappedData {
		columns = append(columns, key)
		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
		args = append(args, value)
		i++
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))

	_, err := p.client.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresAdapter) SelectID(ctx context.Context, tableName string, id string) (interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tableName)

	row := p.client.QueryRow(ctx, query, id)

	var result interface{}
	if err := row.Scan(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *PostgresAdapter) Delete(ctx context.Context, tableName string, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableName)
	_, err := p.client.Exec(ctx, query, id)

	return err
}
