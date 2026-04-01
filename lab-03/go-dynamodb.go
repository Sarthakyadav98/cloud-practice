package main

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {

    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String("ap-south-1"),
    }))

    svc := dynamodb.New(sess)

    // CREATE TABLE
    _, err := svc.CreateTable(&dynamodb.CreateTableInput{
        TableName: aws.String("Students"),
        AttributeDefinitions: []*dynamodb.AttributeDefinition{
            {
                AttributeName: aws.String("id"),
                AttributeType: aws.String("S"),
            },
        },
        KeySchema: []*dynamodb.KeySchemaElement{
            {
                AttributeName: aws.String("id"),
                KeyType:      aws.String("HASH"),
            },
        },
        ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
            ReadCapacityUnits:  aws.Int64(5),
            WriteCapacityUnits: aws.Int64(5),
        },
    })

    if err != nil {
        fmt.Println("Error creating table:", err)
    } else {
        fmt.Println("Table Created")
    }

    // INSERT ITEM
    _, err = svc.PutItem(&dynamodb.PutItemInput{
        TableName: aws.String("Students"),
        Item: map[string]*dynamodb.AttributeValue{
            "id": {
                S: aws.String("1"),
            },
            "name": {
                S: aws.String("Sarthak"),
            },
        },
    })

    if err != nil {
        fmt.Println("Error inserting item:", err)
    } else {
        fmt.Println("Item Inserted")
    }

    // DELETE TABLE
    _, err = svc.DeleteTable(&dynamodb.DeleteTableInput{
        TableName: aws.String("Students"),
    })

    if err != nil {
        fmt.Println("Error deleting table:", err)
    } else {
        fmt.Println("Table Deleted")
    }
}