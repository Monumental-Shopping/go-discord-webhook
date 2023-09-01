# go-discord-webhook
Source code to easily create and send webhooks

# How to use
## Methods
## Responses
- When the `.Send()` method is called, it returns an`*http.Response`. If there are no errors in the webhook, the status code will be `204` and the body will be empty.
- If there are any errors with the webhook (whether a rate limit or invalid data), all of the info will be in the response body. Handling the error is as simple as the following:
```go
package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	discordwebhook "github.com/Monumental-Shopping/go-discord-webhook"
)

func main() {
	// Create a new webhook
	w := discordwebhook.NewWebhook()
	resp, err := w.Send(WEBHOOK_URL)
	if err != nil {
		// TODO handle ur errors here
	}
	handleResponse(resp)

}

func handleResponse(resp *http.Response) {
	// Make sure to close the body after reading
	resp.Close = true
	defer resp.Body.Close()

	// Parse response
	body, _ := io.ReadAll(resp.Body)
	response := &struct {
		Message    string  `json:"message"`
		RetryAfter float64 `json:"retry_after"`
	}{}
	json.Unmarshal(body, response)
	if response.Message == "You are being rate limited." || response.Message == "The resource is being rate limited." {
		time.Sleep(time.Duration(response.RetryAfter) * time.Second)
		// TODO retry to send your webhook here
		w.Send(WEBHOOK_URL)
	}
}

```
- **Do not forget the close the response. Not doing so can cause leaks.**


# Examples
## Basic
```go
package main

import (
	discordwebhook "github.com/Monumental-Shopping/go-discord-webhook"
)

func main() {
	// Create a new webhook
	w := discordwebhook.NewWebhook()

	// Add basic data to the webhook
	w.SetUsername("Webhook Name")
	w.SetTitle("Webhook Title")
	w.SetUrl("https://google.com/")
	w.SetDescription("Text message. You can use Markdown here. *Italic* **bold** __underline__ ~~strikeout~~ [hyperlink](https://google.com) `code`")

	// Add a message that is OUTSIDE of the embed. Up to 2000 characters.
	w.SetContent("Text message. Up to 2000 characters.")

	// Set the color
	w.SetColor(99833)

	// Add fields
	w.CreateField("Text", "More text", true)
	w.CreateField("__Even more text__", "**Did you notice the markdown in the ||title||?**", true)
	w.CreateField("Use `\"inline\": true` parameter, if you want to display fields in the same line.", "okay?", false)

	// Set some images
	w.SetThumbnail("https://www.techsmith.com/blog/wp-content/uploads/2021/02/video-thumbnails-hero-1.png")
	w.SetImage("https://static.wikia.nocookie.net/peanuts/images/1/1e/Joecooltrans.png/revision/latest?cb=20200228013110")

	// Set the timestamp
	w.SetCurrentTimestamp()

	// Send the webhook to the channel
	w.Send(WEBHOOK_URL)
}

```
![sorry if you are seeing this :/](/examples/basic.png)


## Everything
```go
package main

import (
	discordwebhook "github.com/Monumental-Shopping/go-discord-webhook"
)

func main() {
	// Create a new webhook
	w := discordwebhook.NewWebhook()

	// Add basic data to the webhook
	w.SetUsername("Webhook Name")
	w.SetTitle("Webhook Title")
	w.SetUrl("https://google.com/")
	w.SetDescription("Text message. You can use Markdown here. *Italic* **bold** __underline__ ~~strikeout~~ [hyperlink](https://google.com) `code`")

	// Add a message that is OUTSIDE of the embed. Up to 2000 characters.
	w.SetContent("Text message. Up to 2000 characters.")

	// Set the color
	w.SetColor(696969)

	// Add fields
	w.CreateField("Text", "More text", true)
	w.CreateField("__Even more text__", "**Did you notice the markdown in the ||title||?**", true)
	w.AddField(discordwebhook.Field{
		Name:  "Use `\"inline\": true` parameter, if you want to display fields in the same line.",
		Value: "okay?",
	})

	// Add some more fields
	fields := []discordwebhook.Field{
		{
			Name:   "Some other info",
			Value:  "**please** do not leave the value blank",
			Inline: true,
		},
		{
			Name:   "One more",
			Value:  "because why not",
			Inline: true,
		},
	}
	w.AddFields(fields)

	// Add an author to the webhook
	w.SetAuthor(discordwebhook.Author{
		Name:    "Birdie♫",
		URL:     "https://www.reddit.com/r/cats/",
		IconURL: "https://i.imgur.com/R66g1Pe.jpg",
	})

	// Add a footer
	w.SetFooter(discordwebhook.Footer{
		Text:    "Woah! So cool! :smirk:",
		IconURL: "https://i.imgur.com/fKL31aD.jpg",
	})

	// Set some images
	w.SetThumbnail("https://www.techsmith.com/blog/wp-content/uploads/2021/02/video-thumbnails-hero-1.png")
	w.SetImage("https://static.wikia.nocookie.net/peanuts/images/1/1e/Joecooltrans.png/revision/latest?cb=20200228013110")

	// Set the timestamp
	w.SetCurrentTimestamp()

	// Send the webhook to the channel
	w.Send(WEBHOOK_URL)
}

```
![Picture here](/examples/everything.png)

## Everything - Alternative
```go
package main

import (
	"time"

	discordwebhook "github.com/Monumental-Shopping/go-discord-webhook"
)

func main() {
	// Create a new webhook
	w := discordwebhook.NewWebhook()

	// Add basic data to the webhook
	w.SetUsername("Webhook Name")
	w.SetTitle("Webhook Title")
	w.SetUrl("https://google.com/")
	w.SetDescription("Text message. You can use Markdown here. *Italic* **bold** __underline__ ~~strikeout~~ [hyperlink](https://google.com) `code`")

	// Set the color of the webhook
	w.SetColor(15258703)

	// Add a message that is OUTSIDE of the embed. Up to 2000 characters.
	w.SetContent("Text message. Up to 2000 characters.")

	// Add fields
	w.CreateField("Text", "More text", true)
	w.CreateField("__Even more text__", "**Did you notice the markdown in the ||title||?**", true)
	w.AddField(discordwebhook.Field{
		Name:  "Use `\"inline\": true` parameter, if you want to display fields in the same line.",
		Value: "okay?",
	})

	// Add some more fields
	fields := []discordwebhook.Field{
		{
			Name:   "Some other info",
			Value:  "**please** do not leave the value blank",
			Inline: true,
		},
		{
			Name:   "One more",
			Value:  "because why not",
			Inline: true,
		},
	}
	w.AddFields(fields)

	// Add an author to the webhook
	w.CreateAuthor("Birdie♫", "https://www.reddit.com/r/cats/", "https://i.imgur.com/R66g1Pe.jpg")

	// Add a footer
	w.CreateFooter("Woah! So cool! :smirk:", "https://i.imgur.com/fKL31aD.jpg")

	// Set some images
	w.SetThumbnail("https://www.techsmith.com/blog/wp-content/uploads/2021/02/video-thumbnails-hero-1.png")
	w.SetImage("https://static.wikia.nocookie.net/peanuts/images/1/1e/Joecooltrans.png/revision/latest?cb=20200228013110")

	// Set the timestamp
	w.SetTimestamp(time.Unix(21600, 0))

	// Send the webhook to the channel
	w.Send(WEBHOOK_URL)
}

```
![sorry if you are seeing this :/](/examples/everything-alternative.png)