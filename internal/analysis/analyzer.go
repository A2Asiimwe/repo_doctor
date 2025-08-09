package analysis

import "repo_doctor/internal/github"

type Analyzer struct {
	client *github.Client
}

func NewAnalyzer(client *github.Client) *Analyzer {
	return &Analyzer{client: client}
}

func (a *Analyzer) Run(repos []string) error {
	// Add repository analysis logic here
	return nil
}