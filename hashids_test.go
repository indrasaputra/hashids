package hashids_test

import (
	"testing"

	"github.com/indrasaputra/hashids"
	"github.com/stretchr/testify/assert"
)

func TestNewHashID(t *testing.T) {
	t.Run("successfully create an instance of HashID", func(t *testing.T) {
		tables := []struct {
			minLength uint
			salt      string
		}{
			{0, "salt"},
			{1, "one-salt"},
			{2, "]1=-3asc"},
			{3, "~!/..&%%!(#"},
			{4, "|=\\//f35022fj!^@*H((&&#"},
		}

		for _, table := range tables {
			hash, err := hashids.NewHashID(table.minLength, table.salt)

			assert.Nil(t, err)
			assert.NotNil(t, hash)
		}
	})
}
