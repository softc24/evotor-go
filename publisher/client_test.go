package publisher_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/softc24/evotor-go/publisher"
)

func TestPublisher_GetEvents(t *testing.T) {
	token := os.Getenv("PUBLISHER_TOKEN")
	if token == "" {
		t.Skip("PUBLISHER_TOKEN is not set")
	}

	client, _ := publisher.NewClient(publisher.ClientConfig{
		Client:  nil,
		BaseURL: publisher.DefaultURL,
		Token:   token,
	})

	// Test case 1: Single valid event type
	t.Run("Single valid event type", func(t *testing.T) {
		items, isErr := client.GetEvents(
			context.Background(),
			"e0f6dbdf-3150-40f9-869f-a4efdc20242d",
			[]publisher.EventType{publisher.EventTypeDocument},
			publisher.WithSince(time.Now().UnixMilli()-60*1000),
			publisher.WithLimit(10),
		)

		for v := range items {
			t.Log(v)
		}

		if err := isErr(); err != nil {
			t.Error(err)
		}
	})

	// Test case 2: Multiple valid event types
	t.Run("Multiple valid event types", func(t *testing.T) {
		items, isErr := client.GetEvents(
			context.Background(),
			"e0f6dbdf-3150-40f9-869f-a4efdc20242d",
			[]publisher.EventType{publisher.EventTypeDocument, publisher.EventTypeProduct},
			publisher.WithSince(time.Now().UnixMilli()-60*1000),
			publisher.WithLimit(10),
		)

		for v := range items {
			t.Log(v)
		}

		if err := isErr(); err != nil {
			t.Error(err)
		}
	})

	// Test case 3: Empty event type slice
	t.Run("Empty event type slice", func(t *testing.T) {
		items, isErr := client.GetEvents(
			context.Background(),
			"e0f6dbdf-3150-40f9-869f-a4efdc20242d",
			[]publisher.EventType{},
			publisher.WithSince(time.Now().UnixMilli()-60*1000),
			publisher.WithLimit(10),
		)

		for v := range items {
			t.Log(v)
		}

		if err := isErr(); err == nil {
			t.Error("Expected error for empty event type slice, got nil")
		}
	})
}
