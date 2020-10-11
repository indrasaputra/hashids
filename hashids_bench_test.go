package hashids_test

import (
	"testing"

	"github.com/indrasaputra/hashids"
)

func BenchmarkID_MarshalJSON(b *testing.B) {
	id := hashids.ID(66)
	for i := 0; i < b.N; i++ {
		id.MarshalJSON()
	}
}

func BenchmarkID_UnmarshalJSON(b *testing.B) {
	var id hashids.ID
	for i := 0; i < b.N; i++ {
		id.UnmarshalJSON([]byte(`"J4r0MA20No"`))
	}
}

func BenchmarkHashID_Encode(b *testing.B) {
	b.Run("length: 5", func(b *testing.B) {
		hash, _ := hashids.NewHashID(5, "common-salt")
		for i := 0; i < b.N; i++ {
			hash.Encode(hashids.ID(i))
		}
	})

	b.Run("length: 10", func(b *testing.B) {
		hash, _ := hashids.NewHashID(10, "common-salt")
		for i := 0; i < b.N; i++ {
			hash.Encode(hashids.ID(i))
		}
	})

	b.Run("length: 15", func(b *testing.B) {
		hash, _ := hashids.NewHashID(15, "common-salt")
		for i := 0; i < b.N; i++ {
			hash.Encode(hashids.ID(i))
		}
	})

	b.Run("length: 20", func(b *testing.B) {
		hash, _ := hashids.NewHashID(20, "common-salt")
		for i := 0; i < b.N; i++ {
			hash.Encode(hashids.ID(i))
		}
	})
}
