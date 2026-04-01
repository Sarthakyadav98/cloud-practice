import {
  DynamoDBClient,
  DeleteTableCommand
} from "@aws-sdk/client-dynamodb";

const client = new DynamoDBClient({ region: "ap-south-1" });

async function deleteTable() {
  try {
    await client.send(
      new DeleteTableCommand({ TableName: "2023BCS0044" })
    );
    console.log("Table deleted successfully");
  } catch (err) {
    console.error("Error deleting table:", err);
  }
}

deleteTable();

