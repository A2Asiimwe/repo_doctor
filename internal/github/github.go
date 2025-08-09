package github

type Client struct {
	Token string
}

func NewClient(token string) *Client {
	return &Client{Token: token}
}

// Add methods for interacting with GitHub API here
