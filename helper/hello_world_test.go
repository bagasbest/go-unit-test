package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// penamaan file diakhiri _test.go, e.g. hello_world_test.go
// untuk membuat unit test diawali dengan Test
// e.g TestHelloWorld

// command : go test -v
// command : go test -v ./...
// command : go test -v -run=TestSkip
// command : go test -v -run=TestSubTest/Bagas
// command : go test -v -run=/Bagas

// menambahkan assertion dengan Testify
// command : go get github.com/stretchr/testify
// bedanya assertion dengan require adalah require memanggil FailNow() sedangkan assertion memanggio Fail()

// command benchmark : go test -v -bench=.
// command benchmark : go test -v -run=NotMathUnitTest -bench=.
// command benchmark : go test -v -run=NotMathUnitTest -bench=BencmarkTest
// command benchmark : go test -v -run=NotMathUnitTest -bench=../...
// command benchmark : go test -v -run=TestTidakAda -bench=BencmarkSub/Pangestu

func BenchmarkTable(b *testing.B) {
	benchmark := []struct {
		name    string
		request string
	}{
		{
			name:    "Bagas",
			request: "Bagas",
		},
		{
			name:    "Pangestu",
			request: "Pangestu",
		},
		{
			name:    "BagasPangestu",
			request: "BagasPangestu",
		},
		{
			name:    "UjangSuwarman",
			request: "UjangSuwarman",
		},
	}

	for _, benchmark := range benchmark {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Bagas", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Bagas")
		}
	})

	b.Run("Pangestu", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Pangestu")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Bagas")
	}
}

func BenchmarkHelloWorldPangestu(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Pangestu")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Bagas",
			request:  "Bagas",
			expected: "Hello Bagas",
		},
		{
			name:     "Pangestu",
			request:  "Pangestu",
			expected: "Hello Pangestu",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Bagas", func(t *testing.T) {
		result := HelloWorld("Bagas")
		require.Equal(t, "Hello Bagas", result, "Result must be 'Hello Bagas'")
	})

	t.Run("Pangestu", func(t *testing.T) {
		result := HelloWorld("Pangestu")
		require.Equal(t, "Hello Pangestu", result, "Result must be 'Hello Pangestu'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("Before Unit Test")

	m.Run()

	// after
	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can run only in MAC OS")
	}

	result := HelloWorld("Bagas")
	require.Equal(t, "Hello Bagas", result, "Result must be 'Hello Bagas'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Bagas")
	require.Equal(t, "Hello Bagas", result, "Result must be 'Hello Bagas'")
	fmt.Println("TestHelloWorldRequire Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Bagas")
	assert.Equal(t, "Hello Bagas", result, "Result must be 'Hello Bagas'")
	fmt.Println("TestHelloWorldAssert Done")
}

func TestHelloWorldBagas(t *testing.T) {
	result := HelloWorld("Bagas")

	if result != "Hello Bagas" {
		// unit test error dan kode dibawahnya tetap di eksekusi
		// t.Fail()
		// untuk t.Error() mengeluarkan log dan kode dibawahnya tetap di eksekusi
		// sedangkan t.Fatal() mengeluarkan log dan kode dibawahnya tidak di eksekusi
		t.Error("Result mus be 'Hello Bagas'")
	}

	fmt.Println("TestHelloWorldBagas Done")

}

func TestHelloWorldPangestu(t *testing.T) {
	result := HelloWorld("Pangestu")

	if result != "Hello Pangestu" {
		// unit test error dan kode di bawah nya tidak di eksekusi
		//t.FailNow()
		t.Fatal("Hello must be 'Hello Pangestu'")
	}

	fmt.Println("TestHelloWorldPangestu Done")
}
