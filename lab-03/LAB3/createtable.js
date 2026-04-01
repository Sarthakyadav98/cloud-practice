import {
  DynamoDBClient,
  CreateTableCommand
} from "@aws-sdk/client-dynamodb";

const client = new DynamoDBClient({ region: "ap-south-1" });

async function createTable() {
  const params = {
    TableName: "2023BCS0044",
    AttributeDefinitions: [
      { AttributeName: "studentId", AttributeType: "S" }
    ],
    KeySchema: [
      { AttributeName: "studentId", KeyType: "HASH" }
    ],
    BillingMode: "PAY_PER_REQUEST" // no capacity planning needed
  };

  try {
    const data = await client.send(new CreateTableCommand(params));
    console.log("Table created:", data.TableDescription.TableName);
  } catch (err) {
    console.error("Error creating table:", err);
  }
}

createTable();

