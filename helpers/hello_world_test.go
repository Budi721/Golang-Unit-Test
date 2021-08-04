package helpers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Sebelum unit tes")
	m.Run()
	fmt.Println("Setelah unit tes")
}

func TestTable(t *testing.T) {
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name: "Budi",
			request: HelloWorld("Budi"),
			expected: "Hello, Budi",
		},
		{
			name: "Rahmawan",
			request: HelloWorld("Rahmawan"),
			expected: "Hello, Rahmawan",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.request
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		assert.Equal(t, "Hello, Eko", result, "Harus hasilnya Hello, Eko")
	})
	t.Run("Budi", func(t *testing.T) {
		result := HelloWorld("Budi")
		assert.Equal(t, "Hello, Budi", result, "Harus hasilnya Hello, Budi")
	})
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Tes tidak dijalankan di windows")
	}

	result := HelloWorld("Budi")
	assert.Equal(t, "Hello, Budi", result, "Hasil harus Hello, Budi")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eko")
	require.Equal(t, "Hello, Budi", result, "Hasil harus Hello, Budi")
	fmt.Println("Test selesai")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Budi")
	assert.Equal(t, "Hello, Budi", result, "Hasil harus Hello, Budi")
	fmt.Println("Test selesai")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello, Budi" {
		/*
		t.Error("Pesan errornya")
		t.Fatal("Harusnya Hello, Budi")
		 */
		t.Fail()
	}
	fmt.Println("Test selesai")
}