package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/dstotijn/go-notion"
)

func main() {

	ctx := context.Background()
	apiKey := "<NOTION_API_SECRET>"
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	client := notion.NewClient(apiKey, notion.WithHTTPClient(httpClient))

	var parentPageID string
	flag.StringVar(&parentPageID, "parentPageId", "", "Parent page ID.")
	flag.Parse()

	dbProperties := make(notion.DatabaseProperties)
	dbProperties["Name"] = notion.DatabaseProperty{}
	if entry, ok := dbProperties["Name"]; ok {
		entry.Type = notion.DBPropTypeTitle
		entry.Title = &notion.EmptyMetadata{}
		dbProperties["Name"] = entry
	}
	dbProperties["Completed"] = notion.DatabaseProperty{}
	if entry, ok := dbProperties["Completed"]; ok {
		entry.Type = notion.DBPropTypeCheckbox
		entry.Checkbox = &notion.EmptyMetadata{}
		dbProperties["Completed"] = entry
	}
	dbProperties["Completed"] = notion.DatabaseProperty{}
	if entry, ok := dbProperties["Completed"]; ok {
		entry.Type = notion.DBPropTypeCheckbox
		entry.Checkbox = &notion.EmptyMetadata{}
		dbProperties["Completed"] = entry
	}
	dbProperties["Start Date"] = notion.DatabaseProperty{}
	if entry, ok := dbProperties["Start Date"]; ok {
		entry.Type = notion.DBPropTypeDate
		entry.Date = &notion.EmptyMetadata{}
		dbProperties["Start Date"] = entry
	}
	dbProperties["End Date"] = notion.DatabaseProperty{}
	if entry, ok := dbProperties["End Date"]; ok {
		entry.Type = notion.DBPropTypeDate
		entry.Date = &notion.EmptyMetadata{}
		dbProperties["End Date"] = entry
	}
	dbProperties["Primary"] = notion.DatabaseProperty{}
	if entry, ok := dbProperties["Primary"]; ok {
		entry.Type = notion.DBPropTypePeople
		entry.People = &notion.EmptyMetadata{}
		dbProperties["Primary"] = entry
	}

	params := notion.CreateDatabaseParams{
		ParentPageID: parentPageID,
		Title: []notion.RichText{
			{
				Text: &notion.Text{
					Content: "Db-1",
				},
			},
		},
		Properties: dbProperties,
	}

	_, err := client.CreateDatabase(ctx, params)
	if err != nil {
		fmt.Print(err)
	}
}

/*

go run create-database.go -parentPageId "<PARENT-ID>"

*/
