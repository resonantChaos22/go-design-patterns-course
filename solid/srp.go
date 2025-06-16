package main

import (
	"fmt"
	"os"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(idx int) {
	entryCount--
}

// Anti-PAtterns
func (j *Journal) Save(filename string) {
	_ = os.WriteFile(filename, []byte(j.String()), 0644)
}
