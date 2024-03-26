package godiscordwebhook

type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}

// Creates and returns a footer
func createFooter(text, iconUrl string) *Footer {
	return &Footer{
		Text:    text,
		IconURL: iconUrl,
	}
}

// Sets a given footer to a given webhook
func setFooter(w *webhook, f Footer) {
	w.createEmbedIfNotExists()
	w.Embeds[0].Footer = f
}

// Creates and sets a footer to the webhook
func (w *webhook) CreateFooter(text, iconUrl string) {
	// Create footer
	f := createFooter(text, iconUrl)

	// Set footer
	setFooter(w, *f)
}

// Wrapper function to set a given footer to the webhook
func (w *webhook) SetFooter(f Footer) {
	setFooter(w, f)
}
