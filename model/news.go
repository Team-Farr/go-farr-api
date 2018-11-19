package model

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
