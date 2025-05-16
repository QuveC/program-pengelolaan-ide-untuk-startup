package main

import (
	"fmt"
)

const maxIdea = 20

type Idea struct {
	IdIdea      int
	ideaProject string
	Kategori    string
	totalVote   int
}

type rating struct {
	idRating int
	author   string
	IdIdea   int
}

type ideas [maxIdea]Idea

var ideaList ideas
var totalAmount int
var currentId int = 1

//================================ CREATE =============================
func createIdea() {
	if totalAmount >= maxIdea {
		fmt.Println("Kapasitas Ide Penuh")
		return
	}

	var newIdea Idea
	fmt.Print("Masukkan ide: ")
	fmt.Scanln(&newIdea.ideaProject)
	fmt.Print("Masukkan kategori: ")
	fmt.Scanln(&newIdea.Kategori)

	newIdea.totalVote = 0
	newIdea.IdIdea = currentId
	currentId++

	ideaList[totalAmount] = newIdea
	totalAmount++

	fmt.Println("Ide berhasil ditambahkan.")
}

//================================ UPDATE =============================
func updateIdea() {

	var id int
	fmt.Print("Masukkan ID ide yang ingin diupdate: ")
	fmt.Scanln(&id)

	for i := 0; i < totalAmount; i++ {
		if ideaList[i].IdIdea == id {
			fmt.Print("Masukkan ide baru: ")
			fmt.Scanln(&ideaList[i].ideaProject)
			fmt.Print("Masukkan kategori baru: ")
			fmt.Scanln(&ideaList[i].Kategori)
			fmt.Println("Ide berhasil diupdate.")
			return
		}
	}
	fmt.Println("ID tidak ditemukan.")
}

//================================ DELETE =============================
func deleteIdea() {
   

}

//=============================== SHOW ===============================
func showIdea() {
	fmt.Println("\nDaftar Ide:")
	for i := 0; i < totalAmount; i++ {
		fmt.Printf("ID %d | Ide: %s | Kategori: %s | Vote: %d\n",
			ideaList[i].IdIdea,
			ideaList[i].ideaProject,
			ideaList[i].Kategori,
			ideaList[i].totalVote)
	}

	var pilih int
	fmt.Println("\n=== Pilihan ===")
	fmt.Println("1. Create")
	fmt.Println("2. Update")
	fmt.Println("3. Delete")
	fmt.Print("Pilih: ")
	fmt.Scanln(&pilih)

	switch pilih {
	case 1:
		createIdea()
	case 2:
		updateIdea()
	case 3:
		deleteIdea()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

//============================= DUMMY VOTE =============================
func addRating() {
	// fmt.Println("Fitur voting belum diimplementasi.")
}

func popularIdea() {
	// fmt.Println("Fitur popular idea belum diimplementasi.")
}

//============================= MAIN =============================
func main() {
	var pilihan int

	for {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Tambah Ide")
		fmt.Println("2. Voting")
		fmt.Println("3. Lihat Ide Terpopuler")
		fmt.Println("4. Lihat Semua Ide")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			createIdea()
		case 2:
			addRating()
		case 3:
			popularIdea()
		case 4:
			showIdea()
		case 5:
			fmt.Println("Terima kasih.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
