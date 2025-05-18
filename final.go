//Aplikasi Analisis sentimen komentar media sosial

package main
import "fmt"

const NMAX int = 100
type tabInt [NMAX]string

func menu() {
	fmt.Println("=============MENU=============")
	fmt.Println("1.Tambahkan komentar")
	fmt.Println("2.Ubah komentar")
	fmt.Println("3.Hapus komentar")
	fmt.Println("4.EXIT")
	fmt.Println("==============================")
	fmt.Println("pilih menu: ")
}

func main() {
	var a tabInt
	var n int
	var pilihan string

	for {
		menu()
		fmt.Scanln(&pilihan)
		switch pilihan {
		case "1":
			fmt.Scan(&n)
			fmt.Println("Tambakan komentar : ")
			tambahKometar(a, n)
		case "2":
			fmt.Println("Komentar yang ingin di ubah : ")
			fmt.Println("Komentar yang baru : ")
			tampilKomentar(&a, n)
		case "3":
			fmt.Println("komentar yang ingin di hapus : ")
			fmt.Println("komentar setelah di hapus : ")

		case "4":
			fmt.Println("terimakasih sudah masuk di aplikasi kami: ")
			return
		default:
			fmt.Println("pilihan anda tidak valid")
		}
	}
}

func tambahKometar(a tabInt, n int) {
	for i := 0; i < n-1; i++ {
		fmt.Scan(&a[i])
	}
}

func tampilKomentar(a *tabInt, n int) {
	for i := 1; i < n; i++ {
		fmt.Println(i, ". ", a[i])
	}
}
