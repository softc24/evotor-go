package evotor_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/softc24/evotor-go"
)

func TestPublisher_GetEvents(t *testing.T) {
	token := os.Getenv("PUBLISHER_TOKEN")
	if token == "" {
		t.Skip("PUBLISHER_TOKEN is not set")
	}

	client, _ := evotor.NewPublisher(evotor.PublisherConfig{
		BaseURL: evotor.DefaultURL,
		Token:   token,
	})

	// Test case 1: Single valid event type
	t.Run("Single valid event type", func(t *testing.T) {
		items, isErr := client.GetEvents(
			context.Background(),
			"e0f6dbdf-3150-40f9-869f-a4efdc20242d",
			[]evotor.EventType{evotor.EventTypeDocument},
			evotor.WithSince(time.Now().UnixMilli()-60*1000),
			evotor.WithLimit(10),
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
			[]evotor.EventType{evotor.EventTypeDocument, evotor.EventTypeProduct},
			evotor.WithSince(time.Now().UnixMilli()-60*1000),
			evotor.WithLimit(10),
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
			[]evotor.EventType{},
			evotor.WithSince(time.Now().UnixMilli()-60*1000),
			evotor.WithLimit(10),
		)

		for v := range items {
			t.Log(v)
		}

		if err := isErr(); err == nil {
			t.Error("Expected error for empty event type slice, got nil")
		}
	})

	// Test case 4: Invalid event type
	t.Run("Invalid event type", func(t *testing.T) {
		_, isErr := client.GetEvents(
			context.Background(),
			"e0f6dbdf-3150-40f9-869f-a4efdc20242d",
			[]evotor.EventType{"invalid_type"},
			evotor.WithSince(time.Now().UnixMilli()-60*1000),
			evotor.WithLimit(10),
		)

		if err := isErr(); err == nil {
			t.Error("Expected error for invalid event type, got nil")
		}
	})
}
