package store

import (
	"database/sql"
	"testing"
	"time"
)

func TestStore_CreateUser(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		userName     string
		creationTime time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				db: tt.fields.db,
			}
			got, err := s.CreateUser(tt.args.userName, tt.args.creationTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Store.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
