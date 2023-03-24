package test

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func BenchmarkMarshal(b *testing.B) {
	v := struct {
		M map[string]string
		F float64
		I int
		S string
		B []byte
	}{
		M: map[string]string{
			"banana":   "cycle",
			"elephant": "hats",
		},
		F: 3428943729.234345345345345,
		I: 357393453,
		S: "sdhfksjhfks",
		B: make([]byte, 1024),
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(&v)
		if err != nil {
			b.Fatal(err)
		}

	}
}
