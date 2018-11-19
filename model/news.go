package model

import "net/http"

//News struct representing the news from the various scrapers
type News struct {
	ID         string   `json:"id,omitempty"`
	Title      string   `json:"title"`
	Published  int      `json:"published"`
	Categories []string `json:"categories"`
	Tags       []string `json:"tags"`
	Summary    string   `json:"summary"`
	Body       string   `json:"body"`
}

func (n News) Validate(r *http.Request) error {
	return nil
}
