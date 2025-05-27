//Aplikasi Analisis sentimen komentar media sosial

package main

import (
	"fmt"
)

const NMAX int = 100

type tabInt [NMAX]string

var kataPositif = [5]string{"baik", "bagus", "hebat", "ramah", "keren"}
var kataNegatif = [5]string{"jelek", "buruk", "jahat", "payah", "lemah"}

func menu() {
	fmt.Println("=============MENU=============")
	fmt.Println("1.Tambahkan komentar")
	fmt.Println("2.Tampilkan komentar")
	fmt.Println("3.Ubah komentar")
	fmt.Println("4.Hapus komentar")
	fmt.Println("5.analisis sentiment komentar")
	fmt.Println("6.EXIT")
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
			fmt.Println("Tambahkan komentar : ")
			tambahKometar(&a, &n)
		case "2":
			tampilKomentar(a, n)
		case "3":
			fmt.Println("Komentar yang ingin diubah : ")
			ubahKomentar(&a, n)
		case "4":
			fmt.Println("Komentar yang ingin dihapus : ")
			hapusKomentar(&a, &n)
		case "5":
			fmt.Println("daftar berapa komentar positif dan negatif")

		case "6":
			fmt.Println("Terima kasih sudah menggunakan aplikasi kami.")
			return
		default:
			fmt.Println("Pilihan anda tidak valid.")
		}
	}
}

func tambahKometar(a *tabInt, n *int) {
	var comment string
	fmt.Scanln(&comment)
	a[*n] = comment
	*n++
	fmt.Println("Komentar ditambahkan.")
}

func tampilKomentar(a tabInt, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s\n", i+1, a[i])
	}
}

func ubahKomentar(a *tabInt, n int) {
	tampilKomentar(*a, n)
	var index int
	fmt.Print("Masukkan nomor berapa yang ingin diubah: ")
	fmt.Scanln(&index)
	var newComment string
	fmt.Print("Masukkan komentar baru: ")
	fmt.Scanln(&newComment)
	a[index-1] = newComment
	fmt.Println("Komentar berhasil diubah.")
}

func hapusKomentar(a *tabInt, n *int) {
	tampilKomentar(*a, *n)
	var Hapus int
	fmt.Print("Masukkan nomor komentar yang ingin dihapus: ")
	fmt.Scanln(&Hapus)
	if Hapus < 1 || Hapus > *n {
		fmt.Println("Nomor komentar tidak valid.")
		return
	}
	idx := Hapus - 1
	for i := idx; i < *n-1; i++ {
		a[i] = a[i+1]
	}
	a[*n-1] = ""
	*n--
	fmt.Println("Komentar berhasil dihapus.")
}
