package hashids

// ID represents a unique identifier.
// It means to replace the old int64 as unique ID.
// Using this type allows the int64 to be obfuscated
// into a random string using the Hashids algorithm.
// Read more about hashids in https://hashids.org/.
type ID int64
