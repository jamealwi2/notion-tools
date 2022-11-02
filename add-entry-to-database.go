package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/dstotijn/go-notion"
)

func main() {
	ctx := context.Background()
	apiKey := "<NOTION_API_KEY>"
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	client := notion.NewClient(apiKey, notion.WithHTTPClient(httpClient))

	var parentDatabaseId string
	flag.StringVar(&parentDatabaseId, "parentDatabaseId", "", "Parent Database ID")
	flag.Parse()

	dbPageProperties := notion.DatabasePageProperties{}
	dbPageProperties["Name"] = notion.DatabasePageProperty{}
	if entry, ok := dbPageProperties["Name"]; ok {
		entry.Title = []notion.RichText{
			{
				Text: &notion.Text{
					Content: "Entry 1",
				},
			},
		}
		dbPageProperties["Name"] = entry
	}
	dbPageProperties["Completed"] = notion.DatabasePageProperty{}
	if entry, ok := dbPageProperties["Completed"]; ok {
		entry.Checkbox = newTrue()
		dbPageProperties["Completed"] = entry
	}
	dbPageProperties["Start Date"] = notion.DatabasePageProperty{}
	if entry, ok := dbPageProperties["Start Date"]; ok {
		entry.Date = &notion.Date{
			Start: notion.DateTime{
				Time: time.Now().AddDate(0, 0, -7),
			},
			/*End: &notion.DateTime{
				Time: time.Now(),
			},*/
		}
		dbPageProperties["Start Date"] = entry
	}

	dbPageProperties["End Date"] = notion.DatabasePageProperty{}
	if entry, ok := dbPageProperties["End Date"]; ok {
		entry.Date = &notion.Date{
			Start: notion.DateTime{
				Time: time.Now(),
			},
		}
		dbPageProperties["End Date"] = entry
	}

	/*paginationQuery := &notion.PaginationQuery{
		PageSize: 1000,
	}
	res, _ := client.ListUsers(ctx, paginationQuery)
	var workingUser notion.User
	for _, workingUser = range res.Results {
		if workingUser.Name == "AJ" {
			break
		}
	}
	fmt.Print(workingUser)
	cUser := notion.BaseUser{
		ID: workingUser.ID,
	}

	d := notion.User{
		BaseUser: cUser,
	}

	dbPageProperties["Primary"] = notion.DatabasePageProperty{}
	if entry, ok := dbPageProperties["Primary"]; ok {
		entry.People = append(entry.People, d)
		dbPageProperties["Primary"] = entry
	}*/

	params := notion.CreatePageParams{
		ParentType:             notion.ParentTypeDatabase,
		ParentID:               parentDatabaseId,
		DatabasePageProperties: &dbPageProperties,
		Children: []notion.Block{
			notion.Heading1Block{
				RichText: []notion.RichText{
					{
						Text: &notion.Text{
							Content: "Heading 1",
						},
					},
				},
			},
		},
	}

	_, err := client.CreatePage(ctx, params)
	if err != nil {
		log.Fatalf("Failed to create page: %v", err)
	}
}

func newTrue() *bool {
	b := true
	return &b
}

/*

go run add-entry-to-database.go -parentDatabaseId "Parent-Database-ID"

Out: notion-db2.png

*/
