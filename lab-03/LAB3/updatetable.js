import { DynamoDBClient } from "@aws-sdk/client-dynamodb";
import {
  DynamoDBDocumentClient,
  UpdateCommand
} from "@aws-sdk/lib-dynamodb";

const client = new DynamoDBClient({ region: "ap-south-1" });
const docClient = DynamoDBDocumentClient.from(client);

async function updateItem() {
  const params = {
    TableName: "2023BCS0044",
    Key: {
      studentId: "S101"
    },
    UpdateExpression: "SET age = :a",
    ExpressionAttributeValues: {
      ":a": 21
    },
    ReturnValues: "UPDATED_NEW"
  };

  try {
    const result = await docClient.send(new UpdateCommand(params));
    console.log("Updated item:", result.Attributes);
  } catch (err) {
    console.error("Update failed:", err);
  }
}

updateItem();

