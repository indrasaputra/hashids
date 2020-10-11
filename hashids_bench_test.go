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
