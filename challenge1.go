package main

import "fmt"

func main() {
	var (
		i                    = 21
		j            bool    = true
		base16               = 15
		unicodeRusia         = 'Я'
		k            float64 = 123.456
	)
	// menampilkan nilai i
	fmt.Printf("%v\n", i)

	// menampilkan tipe data dari variabel i
	fmt.Printf("%T\n", i)

	// menampilkan tanda %
	fmt.Println("%")

	// menampilkan nilai boolean j : true
	fmt.Println(j)

	// menampilkan nilai base 2 dari i : 10101
	fmt.Printf("%b\n", i)

	// menampilkan nilai base 10 dari i : 21
	fmt.Println(i)

	// menampilkan nilai base 8 dari i : 25
	fmt.Printf("%o\n", i)

	// menampilkan nilai base 16 : f
	fmt.Printf("%x\n", base16)

	// menampilkan nilai base 16 : F
	fmt.Printf("%X\n", base16)

	// menampilkan unicode karakter Я : U+042F
	fmt.Printf("U+%04X\n", unicodeRusia)

	// menampilkan nilai variabel k float64 = 123.456
	fmt.Println(k)

	// menampilkan float : 123.456000
	fmt.Printf("%.6f\n", k)

	// menampilkan float scientific : 1.234560E+02
	fmt.Printf("%.6E\n", k)
}
