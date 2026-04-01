const AWS = require('aws-sdk');

AWS.config.update({ region: 'ap-south-1' });

const dynamodb = new AWS.DynamoDB();
const docClient = new AWS.DynamoDB.DocumentClient();

// CREATE TABLE
const params = {
    TableName: "Students",
    KeySchema: [{ AttributeName: "id", KeyType: "HASH" }],
    AttributeDefinitions: [{ AttributeName: "id", AttributeType: "S" }],
    ProvisionedThroughput: { ReadCapacityUnits: 5, WriteCapacityUnits: 5 }
};

dynamodb.createTable(params, (err, data) => {
    if (err) console.log("Error:", err);
    else console.log("Table Created");

    // INSERT / MODIFY
    const item = {
        TableName: "Students",
        Item: {
            id: "1",
            name: "Sarthak",
            status: "Success"
        }
    };

    docClient.put(item, (err, data) => {
        if (err) console.log(err);
        else console.log("Item Inserted");

        // DELETE TABLE
        dynamodb.deleteTable({ TableName: "Students" }, (err, data) => {
            if (err) console.log(err);
            else console.log("Table Deleted");
        });
    });
});