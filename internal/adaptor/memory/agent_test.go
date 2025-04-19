package memory

import (
	"context"
	"testing"
)

func TestAgentStore(t *testing.T) {
	store := NewAgentStore()

	if ok, err := store.Has(context.Background(), "1"); ok || err != nil {
		t.Errorf("got %v %v, expected %v %v", ok, err, false, nil)
	}

	if err := store.Put(context.Background(), "1"); err != nil {
		t.Errorf("got %v, expected nil", err)
	}

	if ok, err := store.Has(context.Background(), "1"); !ok || err != nil {
		t.Errorf("got %v %v, expected %v %v", ok, err, true, nil)
	}

	if err := store.Del(context.Background(), "1"); err != nil {
		t.Errorf("got %v, expected nil", err)
	}

	if ok, err := store.Has(context.Background(), "1"); ok || err != nil {
		t.Errorf("got %v %v, expected %v %v", ok, err, false, nil)
	}
}
