//Aplikasi Analisis sentimen komentar media sosial

package main

import (
	"fmt"
)

const NMAX int = 100

type comment struct {
	kata1, kata2, kata3 string
	id                  int
}
type tabInt [NMAX]comment

func menu() {
	fmt.Println("|==============================|")
	fmt.Println("|             MENU             |")
	fmt.Println("|==============================|")
	fmt.Println("|1.Tambahkan komentar          |")
	fmt.Println("|2.Tampilkan komentar          |")
	fmt.Println("|3.Ubah komentar               |")
	fmt.Println("|4.Hapus komentar              |")
	fmt.Println("|5.analisis sentiment komentar |")
	fmt.Println("|6.EXIT                        |")
	fmt.Println("|==============================|")
	fmt.Print("pilih menu: ")

}

func main() {
	var a tabInt
	var n int
	var pilihan string
	var positif, negatif int

	for {
		menu()
		fmt.Scanln(&pilihan)
		switch pilihan {
		case "1":
			tambahKomentar(&a, &n)
		case "2":
			tampilKomentar(a, n)
		case "3":
			fmt.Println("List komentar : ")
			ubahKomentar(&a, n)
			tampilKomentar(a, n)
		case "4":
			fmt.Println("Komentar yang ingin dihapus : ")
			hapusKomentar(&a, &n)
			tampilKomentar(a, n)
		case "5":
			fmt.Println("daftar berapa komentar positif dan negatif")
			positif, negatif = analisisSentiment(a, n)
			fmt.Printf("Hasil Analisis:\nPositif : %d\nNegatif : %d\n", positif, negatif)
		case "6":
			fmt.Println("Terima kasih sudah menggunakan aplikasi kami.")
			return
		default:
			fmt.Println("Pilihan anda tidak valid.")
		}
	}
}

func tambahKomentar(a *tabInt, n *int) {
	var kata1, kata2, kata3 string
	var id int
	fmt.Print("Masukan id : ")
	fmt.Scanln(&id)
	for checkk(*a, *n, id) == false {
		fmt.Println("ID sudah ada, buat ID baru")
		fmt.Print("Masukan ID baru : ")
		fmt.Scanln(&id)
	}
	fmt.Print("Tambahkan komentar (3 kata) : ")
	fmt.Scanln(&kata1, &kata2, &kata3)
	a[*n].id = id
	a[*n].kata1 = kata1
	a[*n].kata2 = kata2
	a[*n].kata3 = kata3
	*n++
	fmt.Println("Komentar ditambahkan.")
}

func checkk(a tabInt, n int, id int) bool {
	for i := 0; i < n; i++ {
		if id == a[i].id {
			return false
		}
	}
	return true
}

func tampilKomentar(a tabInt, n int) {
	fmt.Printf("%-6s %-30s\n", "ID", "Komentar")
	for i := 0; i < n; i++ {
		fmt.Printf("%-6d %s %s %s\n", a[i].id, a[i].kata1, a[i].kata2, a[i].kata3)
	}
}

func ubahKomentar(a *tabInt, n int) {
	tampilKomentar(*a, n)
	var idlama, idbaru, idx int
	var kata1, kata2, kata3 string

	fmt.Print("Masukan ID komentar yang ingin di ubah : ")
	fmt.Scanln(&idlama)

	idx = -1
	for i := 0; i < n; i++ {
		if a[i].id == idlama {
			idx = i
		}
	}

	if idx == -1 {
		fmt.Println("ID tidak ditemukan")
		return
	}

	fmt.Print("Masukan ID baru : ")
	fmt.Scanln(&idbaru)
	fmt.Print("Tambahkan komentar baru (3 kata) : ")
	fmt.Scanln(&kata1, &kata2, &kata3)

	a[idx].id = idbaru
	a[idx].kata1 = kata1
	a[idx].kata2 = kata2
	a[idx].kata3 = kata3

	fmt.Println("Komentar telah diubah")
}

func hapusKomentar(a *tabInt, n *int) {
	tampilKomentar(*a, *n)
	var idx, id, i int
	fmt.Print("Masukkan ID komentar yang ingin dihapus: ")
	fmt.Scanln(&id)
	idx = -1
	for i = 0; i < *n; i++ {
		if a[i].id == id {
			idx = i
		}
	}
	if idx == -1 {
		fmt.Println("ID tidak ditemukan")
		return
	}

	for i = idx; i < *n-1; i++ {
		a[i] = a[i+1]
	}
	*n--
	fmt.Println("Komentar telah dihapus.")
}

func analisisSentiment(a tabInt, n int) (int, int) {
	kataPositif := []string{"baik", "bagus", "hebat", "ramah", "keren", "asik", "mantap", "sopan"}
	kataNegatif := []string{"jelek", "buruk", "jahat", "payah", "lemah", "cemen", "sampah", "benci"}

	var positif, negatif int
	var kata1, kata2, kata3 string
	for i := 0; i < n; i++ {
		kata1 = a[i].kata1
		kata2 = a[i].kata2
		kata3 = a[i].kata3
		if check(kataPositif, kata1) {
			positif++
		} else if check(kataNegatif, kata1) {
			negatif++
		}
		if check(kataPositif, kata2) {
			positif++
		} else if check(kataNegatif, kata2) {
			negatif++
		}
		if check(kataPositif, kata3) {
			positif++
		} else if check(kataNegatif, kata3) {
			negatif++
		}
	}

	return positif, negatif
}

func check(pnt []string, kata string) bool {
	for i := 0; i < 8; i++ {
		if pnt[i] == kata {
			return true
		}
	}
	return false
}

