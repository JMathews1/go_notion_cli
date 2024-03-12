package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/dstotijn/go-notion"
)

func main() {
	client := notion.NewClient(os.Getenv("NOTION_TOKEN"))
	databaseID := os.Getenv("NOTION_DATABASE_ID") // Ensure this is set in your environment

	reader := bufio.NewReader(os.Stdin)

	// Prompt for and read the title
	fmt.Print("Enter the title for the new database entry: ")
	inputTitle, _ := reader.ReadString('\n')
	inputTitle = strings.TrimSpace(inputTitle)

	// Prompt for and read the content
	fmt.Print("Enter the content for the new database entry: ")
	inputContent, _ := reader.ReadString('\n')
	inputContent = strings.TrimSpace(inputContent)

	// Define the properties of the new page with the user's input for the title
	properties := notion.DatabasePageProperties{
		"Name": {
			Title: []notion.RichText{
				{
					Text: &notion.Text{
						Content: inputTitle,
					},
				},
			},
		},
	}

	// Create a ParagraphBlock with the user's input content
	paragraphBlock := notion.ParagraphBlock{
		RichText: []notion.RichText{
			{
				Text: &notion.Text{
					Content: inputContent,
				},
			},
		},
		// Optionally, specify other properties like Color or add Children blocks
	}

	// Since the `Children` field in `CreatePageParams` expects a slice of `Block` interfaces,
	// and assuming the `ParagraphBlock` satisfies this interface, you might need to handle
	// the conversion or directly use a slice of `ParagraphBlock` if the library allows.

	// Assuming `CreatePageParams` allows for directly setting `ParagraphBlock` as children:
	children := []notion.Block{&paragraphBlock}

	// Create the new page with the specified properties and content block
	newPageParams := notion.CreatePageParams{
		ParentType:             notion.ParentTypeDatabase,
		ParentID:               databaseID,
		DatabasePageProperties: &properties,
		Children:               children,
	}

	newPage, err := client.CreatePage(context.Background(), newPageParams)
	if err != nil {
		fmt.Printf("Error creating new database entry: %v\n", err)
		return
	}

	fmt.Println("New database entry created successfully:", newPage.ID)
}
