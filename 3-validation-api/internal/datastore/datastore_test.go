package datastore

import (
	"testing"

	"demo/3-validation-api/configs"
)

func TestAddItem(t *testing.T) {
	conf := configs.LoadConfig(".env.test")

	ds := NewDataStore(conf.StoreFilename)

	defer DropFile(ds.filename)

	cases := []Item{
		{
			Email: "a@a.ru",
			Hash:  "123",
		},
		{
			Email: "b@b.ru",
			Hash:  "456",
		},
		{
			Email: "c@c.ru",
			Hash:  "789",
		},
	}
	for _, v := range cases {
		t.Run("Check add items", func(t *testing.T) {
			ds.AddItem(v)

			item := *ds.FindByEmail(v.Email)

			if item.Email != v.Email {
				t.Errorf("Emails don't matg, got: %s, want: %s.", item.Email, v.Email)
			}
		})
	}
}
