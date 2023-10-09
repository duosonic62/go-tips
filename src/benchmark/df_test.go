package benchmark

import (
	"bytes"
	"github.com/go-gota/gota/dataframe"
	"github.com/tobgu/qframe"
	"os"
	"testing"
)

func BenchmarkQFrame(b *testing.B) {
	bs, err := os.ReadFile("iris.csv")
	if err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var out bytes.Buffer
		qf := qframe.ReadCSV(bytes.NewReader(bs))
		qf = qf.Select("sepal_length", "species")
		err = qf.ToCSV(&out)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGota(b *testing.B) {
	bs, err := os.ReadFile("iris.csv")
	if err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var out bytes.Buffer
		df := dataframe.ReadCSV(bytes.NewReader(bs))
		df = df.Select([]string{"sepal_length", "species"})
		err = df.WriteCSV(&out, dataframe.WriteHeader(true))
		if err != nil {
			b.Fatal(err)
		}
	}
}
