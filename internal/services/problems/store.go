package problems

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sangharshseth/gomine-backend/internal/storage"
)

type ProblemStore interface {
	PutItemIntoDynamoDB(context.Context, string, interface{}) error
	GetItemFromTable(context.Context, string) ([]map[string]types.AttributeValue, error)
}
type Store struct {
	db *storage.Client
}

func NewStore(client *storage.Client) *Store {
	return &Store{db: client}
}

func (s *Store) PutItemIntoDynamoDB(ctx context.Context, tableName string, item interface{}) error {
	return s.db.PutItemGeneric(ctx, tableName, item)
}

func (s *Store) GetItemFromTable(ctx context.Context, tableName string) ([]map[string]types.AttributeValue, error) {
	return s.db.GetItemGeneric(ctx, tableName)
}
