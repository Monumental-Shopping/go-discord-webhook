package godiscordwebhook

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// To run this test individually, run the command:
// go test -v -timeout 30s -run ^TestCreateField$ go-discord-webhook
func TestCreateField(t *testing.T) {
	tests := []struct {
		Name, Value string
		Inline      bool
	}{
		{Name: "Test", Value: "Test", Inline: true},
		{Name: "", Value: "", Inline: true},
	}

	for testNum, test := range tests {
		t.Run(fmt.Sprintf("Test #%v", testNum+1), func(t *testing.T) {
			result := createField(test.Name, test.Value, test.Inline)
			if assert.NotNil(t, result, "result should not be nil") {
				assert.Equal(t, test.Name, result.Name)
				assert.Equal(t, test.Value, result.Value)
				assert.Equal(t, test.Inline, result.Inline)
			}
		})
	}
}

// To run this test individually, run the command:
// go test -v -timeout 30s -run ^TestAddField$ go-discord-webhook
func TestAddField(t *testing.T) {
	tests := []struct {
		NumFields int
	}{
		{NumFields: 1},
		{NumFields: 5},
		{NumFields: 10},
		{NumFields: 50},
		{NumFields: 100},
	}

	for testNum, test := range tests {
		t.Run(fmt.Sprintf("Test #%v", testNum+1), func(t *testing.T) {
			w := newWebhook()
			for i := 0; i < test.NumFields; i++ {
				f := Field{}
				w.AddField(f)
			}
			assert.Equal(t, test.NumFields, len(w.Embeds[0].Fields))
		})
	}
}

// To run this test individually, run the command:
// go test -v -timeout 30s -run ^TestAddFields$ go-discord-webhook
func TestAddFields(t *testing.T) {
	tests := []struct {
		NumFields int
	}{
		{NumFields: 1},
		{NumFields: 5},
		{NumFields: 10},
		{NumFields: 50},
		{NumFields: 100},
	}

	for testNum, test := range tests {
		t.Run(fmt.Sprintf("Test #%v", testNum+1), func(t *testing.T) {
			fields := make([]Field, test.NumFields)
			for i := 0; i < test.NumFields; i++ {
				f := Field{}
				fields[i] = f
			}
			w := newWebhook()
			w.AddFields(fields)
			assert.Equal(t, test.NumFields, len(w.Embeds[0].Fields))
		})
	}
}
