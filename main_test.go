package main

import "testing"

func TestCreateTable(t *testing.T) {
	err := InitTable("bookmarktest")

	if err != nil {
		t.Error("error create table")
	}
}
