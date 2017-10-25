package application

import (
	"fmt"
	"testing"
)

func Test_findNotes(t *testing.T) {
	// note := findNote("59e0c5aa785548d795cb5c56", "59eaf4ac83293d1799b1a25d")
	note := findNotes("59e0c5aa785548d795cb5c56")
	fmt.Println(note)
}
