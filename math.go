package math

// Add menambahkan dua angka integer
func Add(a, b int) int {
	return a + b
}

// Subtract mengurangi angka a dengan b
// BUG: Fungsi ini sengaja salah, malah melakukan penjumlahan (+) bukan pengurangan (-)
func Subtract(a, b int) int {
	return a + b
}
