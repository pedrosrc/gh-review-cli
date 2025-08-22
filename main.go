package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type SearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []struct {
		HTMLURL string `json:"html_url"`
		Title   string `json:"title"`
		User    struct {
			Login string `json:"login"`
		} `json:"user"`
	} `json:"items"`
}

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("⚠️ GITHUB_TOKEN not set")
	}

	url := "https://api.github.com/search/issues?q=is:pr+review-requested:@me+is:open"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("🫨Opps.. Error in Request\n")
	}

	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("🚨 Error in Request\n")
	}
	defer resp.Body.Close()

	var result SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}

	if result.TotalCount == 0 {
		fmt.Printf("🎉 You have no Pull Request's to review\n")
		return
	}

	fmt.Printf("📊 You have %d PR(s) to review:\n", result.TotalCount)
	for _, pr := range result.Items {
		fmt.Printf("[ %s - %s]\n", pr.User.Login, pr.Title)
		fmt.Printf("▶ %s\n\n", pr.HTMLURL)
	}
}
