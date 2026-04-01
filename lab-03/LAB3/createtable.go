package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-south-1"))
	if err != nil {
		log.Fatal(err)
	}

	client := dynamodb.NewFromConfig(cfg)

	input := &dynamodb.CreateTableInput{
		TableName: aws.String("2023BCS0044"),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("studentId"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("studentId"),
				KeyType:       types.KeyTypeHash,
			},
		},
		BillingMode: types.BillingModePayPerRequest,
	}

	_, err = client.CreateTable(context.TODO(), input)
	if err != nil {
		log.Fatal("Create table failed:", err)
	}

	log.Println("Table created successfully")
}

