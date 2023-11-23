package models

type Response struct {
	CurrentPage  int     `json:"current_page"`
	Data         []Breed `json:"data"`
	FirstPageURL string  `json:"first_page_url"`
	From         int     `json:"from"`
	LastPage     int     `json:"last_page"`
	LastPageURL  string  `json:"last_page_url"`
	Links        []Links `json:"links"`
	NextPageURL  string  `json:"next_page_url"`
	Path         string  `json:"path"`
	PerPage      int     `json:"per_page"`
	PrevPageURL  string  `json:"prev_page_url"`
	To           int     `json:"to"`
	Total        int     `json:"total"`
}
type Breed struct {
	Breed   string `json:"breed"`
	Country string `json:"country"`
	Origin  string `json:"origin"`
	Coat    string `json:"coat"`
	Pattern string `json:"pattern"`
}
type Links struct {
	URL    string `json:"url"`
	Label  string `json:"label"`
	Active bool   `json:"active"`
}
