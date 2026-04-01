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
	cfg, _ := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-south-1"))

	client := dynamodb.NewFromConfig(cfg)

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("2023BCS0044"),
		Key: map[string]types.AttributeValue{
			"studentId": &types.AttributeValueMemberS{
				Value: "S101",
			},
		},
		UpdateExpression: aws.String("SET age = :a"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":a": &types.AttributeValueMemberN{
				Value: "21",
			},
		},
		ReturnValues: types.ReturnValueUpdatedNew,
	}

	result, err := client.UpdateItem(context.TODO(), input)
	if err != nil {
		log.Fatal("Update failed:", err)
	}

	log.Println("Updated attributes:", result.Attributes)
}

