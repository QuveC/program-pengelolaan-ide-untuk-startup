package main

import (
	"fmt"
	"os"
	"os/exec"
)

const maxIdea = 20
const maxRating int = 100


type Idea struct {
	IdIdea      int
	ideaProject string
	Kategori    string
	totalVote   int
	tgl         int
}

type rating struct {
	idRating int
	author   string
	IdIdea   int
}

type ideas [maxIdea]Idea
type ratingList [maxRating]rating
var totalRating int 
var currentRatingId int
var ratingMenu ratingList

var ideaList ideas
var totalAmount int
var currentId int = 1

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}


//================================ CREATE =============================
func createIdea() {
	clear()
	if totalAmount >= maxIdea {
		fmt.Println("Kapasitas Ide Penuh")
		return
	}

	var newIdea Idea
	fmt.Print("Masukkan ide: ")
	fmt.Scanln(&newIdea.ideaProject)
	fmt.Print("Masukkan kategori: ")
	fmt.Scanln(&newIdea.Kategori)
	fmt.Print("Masukan Tanggal: ")
	fmt.Scanln(&newIdea.tgl)

	newIdea.totalVote = 0
	newIdea.IdIdea = currentId
	currentId++

	ideaList[totalAmount] = newIdea
	totalAmount++

	fmt.Println("Ide berhasil ditambahkan.")
	
}

//================================ UPDATE =============================
func updateIdea() {
    clear()
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
    clear()
	var id int 
	fmt.Print("Masukan ID ide yang ingin dihapus: ")
	fmt.Scan(&id)

	for i:= 0 ;i<= totalAmount ; i++{
		if ideaList[i].IdIdea == id {
			for j := i; j < totalAmount - 1;j++{
				ideaList[j]= ideaList[j+1]
			}
			totalAmount--
			fmt.Println("Ide berhasil dihapus.")
		}

	}
   	fmt.Println("ID tidak ditemukan.")
}

//=============================== SHOW ===============================
func selectionSortByTgl() {

	for i := 0; i < totalAmount-1; i++ {
		minIdx := i
		for j := i + 1; j < totalAmount; j++ {
			if ideaList[j].tgl < ideaList[minIdx].tgl {
				minIdx = j
			}
		}
		if minIdx != i {
			ideaList[i], ideaList[minIdx] = ideaList[minIdx], ideaList[i]
		}
	}
}

func showIdea() {
	clear()
	fmt.Println("\nDaftar Ide :")
	for i := 0; i < totalAmount; i++ {
    fmt.Printf("ID %d | Ide: %s | Kategori: %s | Vote: %d | Tanggal: %d\n",			
	        ideaList[i].IdIdea,
			ideaList[i].ideaProject,
			ideaList[i].Kategori,
			ideaList[i].totalVote,
			ideaList[i].tgl,
		)
	}

	var pilih int
	fmt.Println("\n=== Pilihan ===")
	fmt.Println("1. Create")
	fmt.Println("2. Update")
	fmt.Println("3. Delete")
	fmt.Println("4. back")
	fmt.Print("Pilih: ")
	fmt.Scanln(&pilih)

	switch pilih {
	case 1:
		createIdea()
	case 2:
		updateIdea()
	case 3:
		deleteIdea()
	case 4:
		menu()	
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

//============================= DUMMY VOTE =============================
func addRating() {
	clear()
	if totalAmount == 0 {
		fmt.Println("Belum ada ide yang bisa divote.")
		return
	}

	var r rating 
	fmt.Print("Masukkan nama Anda: ")
	fmt.Scanln(&r.author)
	fmt.Print("Masukkan ID ide yang ingin divote: ")
	fmt.Scanln(&r.IdIdea)

	for i := 0; i < totalAmount; i++ {
		if ideaList[i].IdIdea == r.IdIdea {
			ideaList[i].totalVote++
		}
	}

	if totalRating >= maxRating {
		fmt.Println("Kapasitas rating penuh.")
		return
	}

	r.idRating = currentRatingId
	currentRatingId++
	ratingMenu[totalRating] = r
	totalRating++

	fmt.Println("Vote berhasil ditambahkan.")


}

func insertionShort(){
	for i := 1; i < totalAmount; i++ {
		temp := ideaList[i]
		j := i - 1
		for j >= 0 && ideaList[j].totalVote < temp.totalVote {
			ideaList[j+1] = ideaList[j]
			j--
		}
		ideaList[j+1] = temp
	}
}

func PopularIdea() {
	clear()

	if totalAmount == 0 {
		fmt.Println("Belum ada ide yang bisa ditampilkan.")
		fmt.Println("Tekan Enter untuk kembali ke menu...")
		fmt.Scanln()
		return
	}

	var pilih int
	for {
		fmt.Println("\n=== Pilih Urutan Ide ===")
		fmt.Println("1. Berdasarkan tanggal")
		fmt.Println("2. Berdasarkan voting")
		fmt.Println("3. Kembali ke menu")
		fmt.Print("Pilih: ")
		fmt.Scanln(&pilih)

		if pilih == 3 {
			return
		}

		if pilih == 1 {
			selectionSortByTgl()
			fmt.Println("\n=== Ide berdasarkan Tanggal ===")
		} else if pilih == 2 {
			insertionShort()
			fmt.Println("\n=== Ide berdasarkan Voting ===")
		} else if pilih == 3{
			fmt.Println("Pilihan tidak valid.")
			fmt.Println("Tekan Enter untuk lanjut...")
			fmt.Scanln()
		}else {
			fmt.Println("Pilihan Tidak Ada")
		}

		if pilih == 1 || pilih == 2 {
			fmt.Println("=== Daftar Ide ===")
			for i := 0; i < totalAmount; i++ {
				fmt.Printf("ID %d | Ide: %s | Kategori: %s | Vote: %d | Tanggal: %d\n",
					ideaList[i].IdIdea,
					ideaList[i].ideaProject,
					ideaList[i].Kategori,
					ideaList[i].totalVote,
					ideaList[i].tgl)
			}

			fmt.Println("\nTekan Enter untuk kembali ke pilihan...")
			fmt.Scanln()
		}
	}
}



func menu(){
	clear()
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
			PopularIdea()
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

//============================= MAIN =============================
func main() {
     menu()
	
}
