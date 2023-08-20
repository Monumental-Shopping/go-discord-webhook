package godiscordwebhook

type Author struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IconURL string `json:"icon_url"`
}

// Creates and returns a new author
func createAuthor(name, url, iconUrl string) *Author {
	return &Author{
		Name:    name,
		URL:     url,
		IconURL: iconUrl,
	}
}

// Sets a given author to a given webhook
func setAuthor(w *webhook, a Author) {
	w.Embeds[0].Author = a
}

// Creates and sets an author to the webhook
func (w *webhook) CreateAuthor(name, url, iconUrl string) {
	// Create Author
	a := createAuthor(name, url, iconUrl)

	// Set Author
	setAuthor(w, *a)
}

// Wrapper function to set an author to the webhook
func (w *webhook) SetAuthor(a Author) {
	setAuthor(w, a)
}
