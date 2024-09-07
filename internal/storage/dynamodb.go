package storage

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Client wraps the DynamoDB service client
type Client struct {
	svc *dynamodb.Client
}

// NewClient initializes a new DynamoDB client
func NewClient(region string, profile string) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile),
		config.WithRegion("ap-south-1"))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %w", err)
	}

	return &Client{
		svc: dynamodb.NewFromConfig(cfg),
	}, nil
}

func (db *Client) PutItemGeneric(ctx context.Context, tableName string, item interface{}) error {
	itemToInsert, err := attributevalue.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("unable to marshal item to insert into DynamoDB: %w", err)
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      itemToInsert,
	}
	_, err = db.svc.PutItem(ctx, input)
	if err != nil {
		return fmt.Errorf("unable to put item into DynamoDB: %w", err)
	}
	return nil
}

func (db *Client) GetItemGeneric(ctx context.Context, tableName string) ([]map[string]types.AttributeValue, error) {
	var items []map[string]types.AttributeValue
	var lastEvaluatedKeys map[string]types.AttributeValue

	for {
		input := &dynamodb.ScanInput{
			TableName: aws.String(tableName),
		}

		if lastEvaluatedKeys != nil {
			input.ExclusiveStartKey = lastEvaluatedKeys
		}

		result, err := db.svc.Scan(ctx, input)
		if err != nil {
			return nil, fmt.Errorf("unable to get item from DynamoDB: %w", err)
		}
		items = append(items, result.Items...)
		lastEvaluatedKeys = result.LastEvaluatedKey
		if result.LastEvaluatedKey == nil {
			break
		}

	}
	return items, nil
}
