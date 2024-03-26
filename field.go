package godiscordwebhook

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

// Creates and returns a new field
func createField(name, value string, inline bool) *Field {
	return &Field{
		Name:   name,
		Value:  value,
		Inline: inline,
	}
}

// Sets a field to a webhook
func addField(w *webhook, f Field) {
	w.createEmbedIfNotExists()
	w.Embeds[0].Fields = append(w.Embeds[0].Fields, f)
}

// Creates and adds a field to the webhook
func (w *webhook) CreateField(name, value string, inline bool) {
	// Create Field
	f := createField(name, value, inline)

	// Add field
	addField(w, *f)
}

// Wrapper function to add a field to a webhook
func (w *webhook) AddField(f Field) {
	addField(w, f)
}

// Adds a given list of fields to the webhook
func (w *webhook) AddFields(f []Field) {
	for _, field := range f {
		addField(w, field)
	}
}
