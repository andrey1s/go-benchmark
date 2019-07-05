package concatenation

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

func getString(limit int) []string {
	data := make([]string, 100)
	for i := 0; i < limit; i++ {
		data[i] = "some data string"
	}

	return data
}

func getInt(limit int) []int {
	data := make([]int, 100)
	for i := 0; i < limit; i++ {
		data[i] = rand.Int()
	}

	return data
}
func getFloat(limit int) []float64 {
	data := make([]float64, 100)
	for i := 0; i < limit; i++ {
		data[i] = rand.Float64()
	}

	return data
}

func BenchmarkSfmt(b *testing.B) {
	data := getString(100)
	sfotmat := strings.Repeat("%s ", 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sfmt(sfotmat, data...)
	}
}
func BenchmarkIfmt(b *testing.B) {
	data := getInt(100)
	sfotmat := strings.Repeat("%d ", 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ifmt(sfotmat, data...)
	}
}
func BenchmarkFfmt(b *testing.B) {
	data := getFloat(100)
	sfotmat := strings.Repeat("%g ", 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ffmt(sfotmat, data...)
	}
}

func BenchmarkBuff(b *testing.B) {
	data := getString(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buff(data...)
	}
}

func BenchmarkJoin(b *testing.B) {
	data := getString(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		join(data...)
	}
}

func BenchmarkJoinInt(b *testing.B) {
	in := getInt(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := make([]string, 100)
		for i, r := range in {
			data[i] = strconv.Itoa(r)
		}
		join(data...)
	}
}

func BenchmarkJoinFloat(b *testing.B) {
	in := getFloat(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := make([]string, 100)
		for i, r := range in {
			data[i] = strconv.FormatFloat(r, 'g', -1, 64)
		}
		join(data...)
	}
}

func BenchmarkFor100(b *testing.B) {
	data := getString(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for100(data...)
	}
}

func BenchmarkConc100(b *testing.B) {
	data := getString(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cnt100(data...)
	}
}

func BenchmarkConc100Int(b *testing.B) {
	data := getInt(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cnt100Int(data...)
	}
}

func BenchmarkConc100Float(b *testing.B) {
	data := getFloat(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cnt100Float(data...)
	}
}
