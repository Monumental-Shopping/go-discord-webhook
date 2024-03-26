package godiscordwebhook

import "time"

type Embed struct {
	Fields      []Field   `json:"fields"`
	Thumbnail   UrlSource `json:"thumbnail"`
	Image       UrlSource `json:"image"`
	Footer      Footer    `json:"footer"`
	Author      Author    `json:"author"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Timestamp   string    `json:"timestamp"`
	Description string    `json:"description"`
	Color       int       `json:"color"`
}

type UrlSource struct {
	URL string `json:"url"`
}

// createEmbedIfNotExists creates and sets a new embed slice if one does not exist already
func (w *webhook) createEmbedIfNotExists() {
	// Create a new embed if it doesn't exist
	if len(w.Embeds) == 0 {
		w.Embeds = []Embed{{Fields: make([]Field, 0)}}
	}

}

// Creates and returns a new UrlSource
func newUrlSource(s string) *UrlSource {
	return &UrlSource{s}
}

// Converts a given time to RFC-3339 and sets it to a given webhook
func setTimestamp(w *webhook, t time.Time) {
	// Convert to RFC-3339 and set
	w.createEmbedIfNotExists()
	w.Embeds[0].Timestamp = t.Format(time.RFC3339)
}

// Set the title of the embed
func (w *webhook) SetTitle(s string) {
	w.createEmbedIfNotExists()
	w.Embeds[0].Title = s
}

// Sets the URL of the embed
func (w *webhook) SetUrl(s string) {
	w.createEmbedIfNotExists()
	w.Embeds[0].URL = s
}

// Sets the description of the embed
func (w *webhook) SetDescription(s string) {
	w.createEmbedIfNotExists()
	w.Embeds[0].Description = s
}

// Sets the color of the embed
func (w *webhook) SetColor(i int) {
	w.createEmbedIfNotExists()
	w.Embeds[0].Color = i
}

// Sets the image of the embed
func (w *webhook) SetImage(url string) {
	// Create a new  UrlSource
	u := newUrlSource(url)

	// Set
	w.createEmbedIfNotExists()
	w.Embeds[0].Image = *u
}

// Sets the thumbnail of the embed
func (w *webhook) SetThumbnail(url string) {
	u := newUrlSource(url)

	// Set
	w.createEmbedIfNotExists()
	w.Embeds[0].Thumbnail = *u
}

// Sets a given time to embed
func (w *webhook) SetTimestamp(t time.Time) {
	setTimestamp(w, t)
}

// Adds the current time to the timestamp in the footer
func (w *webhook) SetCurrentTimestamp() {
	// Take the current time in RFC3339
	setTimestamp(w, time.Now())
}
