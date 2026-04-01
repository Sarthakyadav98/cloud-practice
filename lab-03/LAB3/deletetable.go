package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {
	cfg, _ := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-south-1"))

	client := dynamodb.NewFromConfig(cfg)

	_, err := client.DeleteTable(context.TODO(), &dynamodb.DeleteTableInput{
		TableName: aws.String("2023BCS0044"),
	})

	if err != nil {
		log.Fatal("Delete failed:", err)
	}

	log.Println("Table deleted successfully")
}

