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

func TestHashID_Encode(t *testing.T) {
	t.Run("negative integers can't be encoded", func(t *testing.T) {
		tables := []struct {
			minLength uint
			salt      string
			id        hashids.ID
		}{
			{0, "salt", -1},
			{1, "one-salt", -10},
			{2, "]1=-3asc", -7},
			{3, "~!/..&%%!(#", -193013},
			{4, "|=\\//f35022fj!^@*H((&&#", -323652},
		}

		for _, table := range tables {
			hash, _ := hashids.NewHashID(table.minLength, table.salt)
			res, err := hash.Encode(table.id)

			assert.NotNil(t, err)
			assert.Nil(t, res)
		}
	})

	t.Run("successfully encodes uint64 into a slice of byte length at least the same as minimal length", func(t *testing.T) {
		tables := []struct {
			minLength uint
			salt      string
			id        hashids.ID
		}{
			{0, "salt", 1},
			{1, "one-salt", 10},
			{2, "]1=-3asc", 7},
			{3, "~!/..&%%!(#", 193013},
			{4, "|=\\//f35022fj!^@*H((&&#", 323652},
		}

		for _, table := range tables {
			hash, _ := hashids.NewHashID(table.minLength, table.salt)
			res, err := hash.Encode(table.id)

			assert.Nil(t, err)
			assert.NotNil(t, res)
			assert.NotEmpty(t, res)
			assert.True(t, int(table.minLength) <= len(res))
		}
	})
}

func TestHashID_Decode(t *testing.T) {
	t.Run("zero length byte is decoded into 0 and nil error", func(t *testing.T) {
		hash, _ := hashids.NewHashID(10, "salt-is-garam")
		id, err := hash.Decode([]byte{})

		assert.Nil(t, err)
		assert.Zero(t, id)
	})
}
