package main

import (
	"testing"
)

func TestIndex(t *testing.T) {
	slice := []interface{}{
		"this is a line this",
		"this is another line",
		"another line should be here",
		"we are writing a lot of lines",
		"so this is a test",
		"123 hello there",
	}
	m := MemoryStorage{}
	m.Index(slice)
	if len(m.items) != 19 {
		t.Errorf("Expected 19 items, got %d", len(m.items))
	}

	if len(m.items["this"]) != 3 {
		t.Errorf("Expected 3 items, got %d", len(m.items["this"]))
	}

	if len(m.items["123"]) != 1 {
		t.Errorf("Expected 1 item, got %d", len(m.items["123"]))
	}

	if len(m.items["hello"]) != 1 {
		t.Errorf("Expected 1 item, got %d", len(m.items["hello"]))
	}
}

func TestSearch(t *testing.T) {
	slice := []interface{}{
		"this is a line this",
		"this is another line",
		"another line should be here",
		"we are writing a lot of lines",
		"so this is a test",
		"123 hello there",
	}
	m := MemoryStorage{}
	m.Index(slice)
	results := m.Search(slice, "this")
	if len(results) != 3 {
		t.Errorf("Expected 3 results, got %d", len(results))
	}

	expected := []SearchResult{
		{Line: "this is a line this", Count: 2},
		{Line: "this is another line", Count: 1},
		{Line: "so this is a test", Count: 1},
	}

	for i, result := range results {
		if result.Line != expected[i].Line {
			t.Errorf("Expected %v, got %v", expected[i].Line, result.Line)
		}
		if result.Count != expected[i].Count {
			t.Errorf("Expected %v, got %v", expected[i].Count, result.Count)
		}
	}

	results = m.Search(slice, "123")
	if len(results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(results))
	}

	results = m.Search(slice, "hello")
	if len(results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(results))
	}
}
