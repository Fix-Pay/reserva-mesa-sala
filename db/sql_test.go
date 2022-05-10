package db

import "testing"

func TestCreateDB(t *testing.T) {
	_, err := CreateDB()
	if err != nil {
		t.Error(err)
	}

}
