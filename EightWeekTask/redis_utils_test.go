package main

import (
	"testing"
)

func TestBatchInsert(t *testing.T) {
	tests := []struct {
		name string
		size int
		db   int
	}{
		{
			name: "10W",
			size: 1,
			db:   0,
		},
		{
			name: "20W",
			size: 2,
			db:   1,
		},
		{
			name: "50W",
			size: 5,
			db:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BatchInsert(tt.size, tt.db)
		})
	}
}
