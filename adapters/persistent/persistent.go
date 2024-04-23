package persistent

import "context"

type Persistent interface {
    Connect(connectionString string) error
    Close() error
    Insert(ctx context.Context, tableName string, data interface{}) error
    SelectID(ctx context.Context, tableName string, id string) (interface{}, error)
    Delete(ctx context.Context, tableName string, id string) error
}


