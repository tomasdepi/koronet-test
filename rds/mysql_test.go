package rds

import "testing"

func TestDBConnection(t *testing.T) {
	InitDB()
	defer CloseDB()

	if DB == nil {
		t.Fatal("DB is nil")
	}

	err := DB.Ping()
	if err != nil {
		t.Fatalf("Ping failed: %v", err)
	}
}
