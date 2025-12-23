package notes_test

import (
	"context"
	"testing"

	"github.com/danthemo/pz8-mongo/internal/db"
	"github.com/danthemo/pz8-mongo/internal/notes"
)

func TestCreateAndGet(t *testing.T) {
	ctx := context.Background()

	uri := "mongodb://root:secret@localhost:27017/?authSource=admin"
	dbName := "pz8-test"

	deps, err := db.ConnectMongo(ctx, uri, dbName)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		deps.Client.Disconnect(context.Background())
	})

	repo, err := notes.NewRepo(deps.Database)
	if err != nil {
		t.Fatal(err)
	}

	created, err := repo.Create(ctx, "Test Title", "Test Content")
	if err != nil {
		t.Fatal(err)
	}

	got, err := repo.ByID(ctx, created.ID.Hex())
	if err != nil {
		t.Fatal(err)
	}

	if got.Title != "Test Title" {
		t.Fatalf("want 'Test Title', got '%s'", got.Title)
	}

	if got.Content != "Test Content" {
		t.Fatalf("want 'Test Content', got '%s'", got.Content)
	}
}
