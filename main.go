package main

import (
	"fmt"
	"strings"
)

type MemoryStorage struct {
	items map[string]map[int]int
}

type SearchResult struct {
	Line  interface{}
	Count int
}

func (m *MemoryStorage) Index(slice []interface{}) {
	m.items = make(map[string]map[int]int)
	for i, item := range slice {
		switch item := item.(type) {
		case string:
			words := strings.Fields(item)
			for _, word := range words {
				if _, ok := m.items[word]; !ok {
					m.items[word] = make(map[int]int)
				}
				m.items[word][i]++
			}
		default:
			if _, ok := m.items[fmt.Sprintf("%v", item)]; !ok {
				m.items[fmt.Sprintf("%v", item)] = make(map[int]int)
			}
			m.items[fmt.Sprintf("%v", item)][i]++
		}
	}
}

func (m *MemoryStorage) Search(slice []interface{}, key interface{}) []SearchResult {
	var results []SearchResult
	switch key := key.(type) {
	case string:
		for k, v := range m.items[key] {
			results = append(results, SearchResult{Line: slice[k], Count: v})
		}
	default:
		for k, v := range m.items[fmt.Sprintf("%v", key)] {
			results = append(results, SearchResult{Line: slice[k], Count: v})
		}
	}
	return results
}

func main() {
	slice := []interface{}{
		"this is a line this",
		"this is another line",
		"another line should be here",
		"we are writing a lot of lines",
		"so this is a test",
		"123 hello there",
		123,
		456,
	}
	storage := &MemoryStorage{}
	storage.Index(slice)

	result := storage.Search(slice, "this")
	fmt.Println("Searching for: 'this'")
	for _, r := range result {
		fmt.Printf("%v: %v", r.Line, r.Count)
		fmt.Println()
	}
	ind := storage.Search(slice, "123")
	fmt.Println("Searching for: '123'")
	for _, r := range ind {
		fmt.Printf("%v: %v", r.Line, r.Count)
		fmt.Println()
	}
}
