package main

import (
    "context"
    "fmt"
    "os"

    "github.com/dstotijn/go-notion"
)

func main() {
    client := notion.NewClient(os.Getenv("NOTION_TOKEN"))
    databaseID := os.Getenv("NOTION_DATABASE_ID") 

    // Example: Update the database title
    updateParams := notion.UpdateDatabaseParams{
        Title: []notion.RichText{
            {
                Text: &notion.Text{
                    Content: " hello there",
                },
            },
        },
    }

    updatedDB, err := client.UpdateDatabase(context.Background(), databaseID, updateParams)
    if err != nil {
        fmt.Printf("Error updating database: %v\n", err)
        return
    }

    fmt.Println("Database updated successfully:", updatedDB.ID)
}
