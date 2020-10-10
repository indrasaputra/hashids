package hashids

// ID represents a unique identifier.
// It means to replace the old int64 as unique ID.
// Using this type allows the int64 to be obfuscated
// into a random string using the Hashids algorithm.
// Read more about hashids in https://hashids.org/.
type ID int64

// Hash defines the contract to encode and decode the ID.
type Hash interface {
	// Encode encodes the ID into a string.
	// The string generated is the result of Hashids algorithm.
	Encode(ID) (string, error)
	// Decode decodes the string into an ID.
	Decode(string) (ID, error)
}

// HashID can be used to encode and decode hashids.
// It implements the Hash interface.
type HashID struct{}
