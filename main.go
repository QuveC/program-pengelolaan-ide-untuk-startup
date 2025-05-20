package main

import (
	"os"
	"fmt"
	"time"
	"bufio"
	"os/exec"
	"strings"
)

const maxIdea = 20
const maxRating int = 100

type Idea struct {
	IdIdea      int
	ideaProject string
	Kategori    string
	totalVote   int
	tgl         time.Time
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
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		clear()
		return
	}

	reader := bufio.NewReader(os.Stdin)

	var newIdea Idea

	newIdea.IdIdea = currentId
	currentId++

	fmt.Print("Masukkan ide: ")
	ideaInput, _ := reader.ReadString('\n')
	newIdea.ideaProject = strings.TrimSpace(ideaInput)

	fmt.Print("Masukkan kategori: ")
	kategoriInput, _ := reader.ReadString('\n')
	newIdea.Kategori = strings.TrimSpace(kategoriInput)

	fmt.Print("Masukkan Tanggal (YYYY-MM-DD): ")
	tanggalInput, _ := reader.ReadString('\n')
	tanggalStr := strings.TrimSpace(tanggalInput)

	tgl, err := time.Parse("2006-01-02", tanggalStr)
	if err != nil {
		fmt.Println("Format tanggal tidak valid!")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		clear()
		return
	}
	newIdea.tgl = tgl

	newIdea.totalVote = 0
	
	ideaList[totalAmount] = newIdea
	totalAmount++

	fmt.Println("Ide berhasil ditambahkan.")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	clear()
}

//================================ UPDATE =============================
func updateIdea() {
	clear()
	for i := 0; i < totalAmount; i++ {
		fmt.Printf("ID %d | Ide: %s | Kategori: %s | Vote: %d | Tanggal: %s\n",			
			ideaList[i].IdIdea,
			ideaList[i].ideaProject,
			ideaList[i].Kategori,
			ideaList[i].totalVote,
			ideaList[i].tgl.Format("2006-01-02"),
		)
	}
	var id int
	fmt.Println("")
	fmt.Print("Masukkan ID ide yang ingin diupdate: ")
	fmt.Scanln(&id)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < totalAmount; i++ {
		if ideaList[i].IdIdea == id {
			fmt.Print("Masukkan ide baru: ")
			ideaInput, _ := reader.ReadString('\n')
			ideaList[i].ideaProject = strings.TrimSpace(ideaInput)

			fmt.Print("Masukkan kategori baru: ")
			kategoriInput, _ := reader.ReadString('\n')
			ideaList[i].Kategori = strings.TrimSpace(kategoriInput)
			
			ideaList[i].tgl = time.Now()	

			fmt.Println("Ide berhasil diupdate.")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			clear()
			return
		}
	}
	fmt.Println("ID tidak ditemukan.")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	clear()
}

//================================ DELETE =============================
func deleteIdea() {
    clear()
	for i := 0; i < totalAmount; i++ {
		fmt.Printf("ID %d | Ide: %s | Kategori: %s | Vote: %d | Tanggal: %s\n",			
			ideaList[i].IdIdea,
			ideaList[i].ideaProject,
			ideaList[i].Kategori,
			ideaList[i].totalVote,
			ideaList[i].tgl.Format("2006-01-02"),
		)
	}
	var id int
    fmt.Println("")
    fmt.Print("Masukan ID ide yang ingin dihapus: ")
    fmt.Scan(&id)

    found := false
    for i := 0; i < totalAmount; i++ {
        if ideaList[i].IdIdea == id {
            for j := i; j < totalAmount-1; j++ {
                ideaList[j] = ideaList[j+1]
            }
            totalAmount--
            found = true
            fmt.Println("Ide berhasil dihapus.")
            break
        }
    }

    if found {
        for i := 0; i < totalAmount; i++ {
            ideaList[i].IdIdea = i + 1
        }
    } else {
		fmt.Println("ID tidak ditemukan.")
    }
}

//=============================== SHOW ===============================
func selectionSortByTgl() {

	for i := 0; i < totalAmount-1; i++ {
		minIdx := i
		for j := i + 1; j < totalAmount; j++ {
			// if ideaList[j].tgl < ideaList[minIdx].tgl {
			if ideaList[j].tgl.Before(ideaList[minIdx].tgl) {
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
	if totalAmount == 0 {
		fmt.Println("Belum ada data yang di tambahkan")
	}
		fmt.Println("╔═════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                              List Ide                               ║")
		fmt.Println("╠════╦════════════════════════╦══════════════╦═══════╦════════════════╣")
		fmt.Println("║ No ║        Ide             ║   Kategori   ║ Vote  ║    Tanggal     ║")
		fmt.Println("╠════╬════════════════════════╬══════════════╬═══════╬════════════════╣")
	for i := 0; i < totalAmount; i++ {
			fmt.Printf("║ %-2d ║ %-22s ║ %-12s ║ %-5d ║ %-14s ║\n",
			ideaList[i].IdIdea,
			ideaList[i].ideaProject,
			ideaList[i].Kategori,
			ideaList[i].totalVote,
			ideaList[i].tgl.Format("2006-01-02"),
		)
	}
        fmt.Println("╚════╩════════════════════════╩══════════════╩═══════╩════════════════╝")

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
		
	for i := 0; i < totalAmount; i++ {
    fmt.Printf("ID %d | Ide: %s | Kategori: %s | Vote: %d | Tanggal: %s\n",			
			ideaList[i].IdIdea,
			ideaList[i].ideaProject,
			ideaList[i].Kategori,
			ideaList[i].totalVote,
			ideaList[i].tgl.Format("2006-01-02"),
		)
	}

	var r rating 
	fmt.Print("Masukkan nama Anda: ")
	fmt.Scanln(&r.author)
	fmt.Print("Masukkan ID ide yang ingin divote: ")
	fmt.Scanln(&r.IdIdea)

		for i := 0; i < totalAmount; i++ {
			if ideaList[i].IdIdea == r.IdIdea {
				ideaList[i].totalVote++
			}else{
				fmt.Println("Id Tidak DI Temukan!")
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
	
		fmt.Println("╔══════════════════════════════════════════════════╗")
		fmt.Println("║                 Pilih UrutaN Ide                 ║")
		fmt.Println("╠══════════════════════════════════════════════════╣")
		fmt.Println("║ 1. Berdasarkan tanggal                           ║")
		fmt.Println("║ 2. Berdasarkan voting                            ║")
		fmt.Println("║ 3. Kembali ke menu                               ║")
		fmt.Println("╚══════════════════════════════════════════════════╝")
		fmt.Print("Pilih : ")
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

		fmt.Println("╔═════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                      IDE BERDASARKAN URUTAN                         ║")
		fmt.Println("╠════╦════════════════════════╦══════════════╦═══════╦════════════════╣")
		fmt.Println("║ No ║        Ide             ║   Kategori   ║ Vote  ║    Tanggal     ║")
		fmt.Println("╠════╬════════════════════════╬══════════════╬═══════╬════════════════╣")

		if pilih == 1 || pilih == 2 {
			for i := 0; i < totalAmount; i++ {
			fmt.Printf("║ %-2d ║ %-22s ║ %-12s ║ %-5d ║ %-14s ║\n",
					i+1,
					ideaList[i].ideaProject,
					ideaList[i].Kategori,
					ideaList[i].totalVote,
					ideaList[i].tgl.Format("2006-01-02"))
			}
		fmt.Println("╚════╩════════════════════════╩══════════════╩═══════╩════════════════╝")
			fmt.Println("\nTekan Enter untuk kembali ke pilihan...")
			fmt.Scanln()
		}
	}
}



func menu(){
	clear()
	var pilihan int 
    for {
		fmt.Println("╔══════════════════════════════════════════════════╗")
		fmt.Println("║                    MENU UTAMA                    ║")
		fmt.Println("╠══════════════════════════════════════════════════╣")
		fmt.Println("║ 1. Tambah Ide                                    ║")
		fmt.Println("║ 2. Voting                                        ║")
		fmt.Println("║ 3. Lihat Ide Terpopuler                          ║")
		fmt.Println("║ 4. Lihat Semua Ide                               ║")
		fmt.Println("║ 5. Keluar                                        ║")
		fmt.Println("╚══════════════════════════════════════════════════╝")
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