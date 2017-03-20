package api

import (
	"testing"
	"math/rand"
	"time"
	"strings"
)



func TestRandomBytes(t *testing.T) {
	for i:= 0; i < 13; i++ {
		bytes := RandomBytes(i)

		if len(bytes) != i {
			t.Fatalf("Length must be %d. The resulted byte array was %d length", i, len(bytes))
		}

		for _, char := range bytes {
			if !strings.Contains(letterBytes, string(char)) {
				t.Fatalf("'%v' isn't in the possibles chars", string(char))
			}
		}
	}
}

func TestNumCombinations(t *testing.T) {
	NumCombinations()
}

func BenchmarkRandomBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := rand.NewSource(time.Now().UnixNano())
		length := rand.New(s).Intn(10)
		RandomBytes(length)
	}
}