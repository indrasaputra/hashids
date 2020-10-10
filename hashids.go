package hashids

import (
	gohashids "github.com/speps/go-hashids"
)

// ID represents a unique identifier.
// It means to replace the old int64 as unique ID.
// Using this type allows the int64 to be obfuscated
// into a random string using the Hashids algorithm.
// Read more about hashids in https://hashids.org/.
type ID int64

// Hash defines the contract to encode and decode the ID.
type Hash interface {
	// Encode encodes the ID into a slice of byte.
	// The slice of byte generated is the result of Hashids algorithm.
	Encode(ID) ([]byte, error)
	// Decode decodes the slice of byte into an ID.
	Decode([]byte) (ID, error)
}

// HashID can be used to encode and decode hashids.
// It implements the Hash interface.
type HashID struct {
	hasher *gohashids.HashID
}

// NewHashID creates an instance of HashID.
// It needs two parameters. The minimum length is used to define
// the mininum length of generated string.
// The salt is used to add the uniqueness of the generated hash.
func NewHashID(minLength uint, salt string) (*HashID, error) {
	data := gohashids.NewData()
	data.Salt = salt
	data.MinLength = int(minLength)
	hasher, err := gohashids.NewWithData(data)
	if err != nil {
		return nil, err
	}

	return &HashID{
		hasher: hasher,
	}, nil
}
