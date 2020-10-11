package hashids_test

import (
	"encoding/json"
	"testing"

	"github.com/indrasaputra/hashids"
	"github.com/stretchr/testify/assert"
)

func TestID_MarshalJSON(t *testing.T) {
	t.Run("zero id be marshalled as 'null'", func(t *testing.T) {
		id := hashids.ID(0)
		res, err := id.MarshalJSON()

		assert.Nil(t, err)
		assert.Equal(t, "null", string(res))
	})

	t.Run("negative number can't be marshalled", func(t *testing.T) {
		ids := []hashids.ID{
			hashids.ID(-1),
			hashids.ID(-43),
			hashids.ID(-66),
		}

		for _, id := range ids {
			res, err := id.MarshalJSON()

			assert.NotNil(t, err)
			assert.Nil(t, res)
		}
	})

	t.Run("successfully marshal positive number", func(t *testing.T) {
		tables := []struct {
			hash string
			id   hashids.ID
		}{
			{`"oWx0b8DZ1a"`, 1},
			{`"EO19oA6vGx"`, 43},
			{`"J4r0MA20No"`, 66},
		}

		for _, table := range tables {
			res, err := table.id.MarshalJSON()

			assert.Nil(t, err)
			assert.NotEmpty(t, res)
			assert.Equal(t, table.hash, string(res))
		}
	})
}

func TestID_UnmarshalJSON(t *testing.T) {
	t.Run("'null' is marshalled to zero ID", func(t *testing.T) {
		id := hashids.ID(10)
		id.UnmarshalJSON([]byte(`null`))

		assert.Equal(t, hashids.ID(0), id)
	})

	t.Run("invalid hash can't be marshalled", func(t *testing.T) {
		tables := []struct {
			hash string
			id   hashids.ID
		}{
			{"oWx0b8DZ1a", 1},
			{"EO19oA6vGx", 43},
			{"J4r0MA20No", 66},
		}

		for _, table := range tables {
			var id hashids.ID
			err := id.UnmarshalJSON([]byte(table.hash))

			assert.NotNil(t, err)
			assert.NotEqual(t, table.id, id)
		}
	})

	t.Run("successfully unmarshal valid hashes", func(t *testing.T) {
		tables := []struct {
			hash string
			id   hashids.ID
		}{
			{`"oWx0b8DZ1a"`, 1},
			{`"EO19oA6vGx"`, 43},
			{`"J4r0MA20No"`, 66},
		}

		for _, table := range tables {
			var id hashids.ID
			err := id.UnmarshalJSON([]byte(table.hash))

			assert.Nil(t, err)
			assert.Equal(t, table.id, id)
		}
	})
}

func TestID_MarshalAndUnmarshal(t *testing.T) {
	t.Run("ID gets back to original ID when unmarshal after marshal", func(t *testing.T) {
		ids := []hashids.ID{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 200, 1000, 10000, 100000, 1000000, 10000000, 100000000}
		for _, id := range ids {
			res, err := json.Marshal(id)
			assert.Nil(t, err)

			var tmp hashids.ID
			err = json.Unmarshal(res, &tmp)
			assert.Nil(t, err)
			assert.Equal(t, id, tmp)
		}
	})

	t.Run("ID on struct gets back to original ID when unmarshal after marshal", func(t *testing.T) {
		type Product struct {
			ID   hashids.ID `json:"id"`
			Name string     `json:"name"`
		}

		ids := []hashids.ID{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 200, 1000, 10000, 100000, 1000000, 10000000, 100000000}
		for _, id := range ids {
			prod := Product{id, "product's name"}
			res, err := json.Marshal(prod)
			assert.Nil(t, err)

			var tmp Product
			err = json.Unmarshal(res, &tmp)
			assert.Nil(t, err)
			assert.Equal(t, id, tmp.ID)
			assert.Equal(t, prod, tmp)
		}
	})
}

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

	t.Run("not hashids byte can't be decoded", func(t *testing.T) {
		hash, _ := hashids.NewHashID(10, "salt-is-garam")
		id, err := hash.Decode([]byte(`garamissalt`))

		assert.NotNil(t, err)
		assert.Zero(t, id)
	})

	t.Run("can't decode hash that contains more than one ID", func(t *testing.T) {
		inputs := []string{"lqs9SN", "YkpsL"}
		for _, inp := range inputs {
			hash, _ := hashids.NewHashID(5, "salt-is-garam")
			id, err := hash.Decode([]byte(inp))

			assert.NotNil(t, err)
			assert.Zero(t, id)
		}
	})

	t.Run("successfully decodes the hashid", func(t *testing.T) {
		inputs := []struct {
			hash string
			id   hashids.ID
		}{
			{"kYeBo", 1},
			{"bQ6Mo", 2},
			{"AYk2O", 3},
		}

		for _, inp := range inputs {
			hash, _ := hashids.NewHashID(5, "salt-is-garam")
			id, err := hash.Decode([]byte(inp.hash))

			assert.Nil(t, err)
			assert.Equal(t, inp.id, id)
		}
	})
}

func TestEncodeID(t *testing.T) {
	t.Run("negative ID can't be encoded", func(t *testing.T) {
		ids := []hashids.ID{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10}
		for _, id := range ids {
			res, err := hashids.EncodeID(id)

			assert.NotNil(t, err)
			assert.Empty(t, res)
			assert.Nil(t, res)
		}
	})

	t.Run("successfully encodes the ID", func(t *testing.T) {
		ids := []hashids.ID{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for _, id := range ids {
			res, err := hashids.EncodeID(id)

			assert.Nil(t, err)
			assert.NotEmpty(t, res)
		}
	})
}

func TestDecodeHash(t *testing.T) {
	t.Run("can't decode invalid hash", func(t *testing.T) {
		inputs := []struct {
			hash string
			id   hashids.ID
		}{
			{"oWx0DZ1a", 1},
			{"EO19ovGx", 43},
			{"J4MA20No", 66},
		}

		for _, inp := range inputs {
			id, err := hashids.DecodeHash([]byte(inp.hash))

			assert.NotNil(t, err)
			assert.NotEqual(t, inp.id, id)
		}
	})

	t.Run("successfully decodes the hashid", func(t *testing.T) {
		inputs := []struct {
			hash string
			id   hashids.ID
		}{
			{"oWx0b8DZ1a", 1},
			{"EO19oA6vGx", 43},
			{"J4r0MA20No", 66},
		}

		for _, inp := range inputs {
			id, err := hashids.DecodeHash([]byte(inp.hash))

			assert.Nil(t, err)
			assert.Equal(t, inp.id, id)
		}
	})
}

func TestSetHasher(t *testing.T) {
	t.Run("different hasher produces different hash even the minimum length is same for the same ID", func(t *testing.T) {
		hasher1, _ := hashids.NewHashID(5, "new-salt")
		hashids.SetHasher(hasher1)
		id := hashids.ID(1)
		res1, _ := id.MarshalJSON()

		hasher2, _ := hashids.NewHashID(5, "new-salt-again")
		hashids.SetHasher(hasher2)
		id = hashids.ID(1)
		res2, _ := id.MarshalJSON()

		assert.NotEqual(t, res1, res2)
	})
}
