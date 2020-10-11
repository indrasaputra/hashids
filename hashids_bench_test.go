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
