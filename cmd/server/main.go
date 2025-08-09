package main

import (
	"fmt"
	"log"
	"os"

	"repo_doctor/internal/config"
	"repo_doctor/internal/github"
	"repo_doctor/internal/analysis"
)

func main() {
	fmt.Println("Starting repo_doctor...")

	cfg, err := config.LoadConfig("configs/default.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	client := github.NewClient(cfg.GitHubToken)
	analyzer := analysis.NewAnalyzer(client)
	err = analyzer.Run(cfg.Repos)
	if err != nil {
		log.Fatalf("analysis failed: %v", err)
	}
}