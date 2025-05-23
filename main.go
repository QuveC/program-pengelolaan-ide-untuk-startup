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

func dummyData() {
	ideaList[0] = Idea{IdIdea: 1, ideaProject: "Aplikasi Donasi Online", Kategori: "Sosial", totalVote: 5, tgl: time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC)}
	ideaList[1] = Idea{IdIdea: 2, ideaProject: "Sistem Absensi QR", Kategori: "Teknologi", totalVote: 8, tgl: time.Date(2025, 3, 22, 0, 0, 0, 0, time.UTC)}
	ideaList[2] = Idea{IdIdea: 3, ideaProject: "Aplikasi Belajar Bahasa", Kategori: "Pendidikan", totalVote: 3, tgl: time.Date(2025, 5, 5, 0, 0, 0, 0, time.UTC)}
	ideaList[3] = Idea{IdIdea: 4, ideaProject: "Marketplace UMKM", Kategori: "Ekonomi", totalVote: 10, tgl: time.Date(2025, 2, 15, 0, 0, 0, 0, time.UTC)}
	totalAmount = 4
	currentId = 5

	ratingMenu[0] = rating{idRating: 1, author: "Budi", IdIdea: 2}
	ratingMenu[1] = rating{idRating: 2, author: "Sari", IdIdea: 4}
	ratingMenu[2] = rating{idRating: 3, author: "Ani", IdIdea: 1}
	ratingMenu[3] = rating{idRating: 4, author: "Dewi", IdIdea: 4}
	ratingMenu[4] = rating{idRating: 5, author: "Rudi", IdIdea: 2}
	ratingMenu[5] = rating{idRating: 6, author: "Andi", IdIdea: 4}
	ratingMenu[6] = rating{idRating: 7, author: "Putri", IdIdea: 2}
	ratingMenu[7] = rating{idRating: 8, author: "Agus", IdIdea: 4}
	totalRating = 8
	currentRatingId = 9
}


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
		fmt.Scanln()
		clear()
		return
	}
	
	var newIdea Idea
	reader := bufio.NewReader(os.Stdin)

	newIdea.IdIdea = currentId
	currentId++

	fmt.Print("Masukkan ide: ")
	ideaInput, _ := reader.ReadString('\n')
	newIdea.ideaProject = strings.TrimSpace(ideaInput)

	fmt.Print("Masukkan kategori: ")
	kategoriInput, _ := reader.ReadString('\n')
	newIdea.Kategori = strings.TrimSpace(kategoriInput)

	newIdea.tgl = time.Now()
	newIdea.totalVote = 0
	ideaList[totalAmount] = newIdea
	totalAmount++

	fmt.Println("Ide berhasil ditambahkan.")
	fmt.Scanln()
	clear()
}

//================================ UPDATE =============================
func updateIdea() {
	clear()

    fmt.Println("╔═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
    fmt.Println("║                                     List Ide                                                                      ║")
    fmt.Println("╠════╦══════════════════════════════════════════════════════╦══════════════════════════╦═══════╦════════════════════╣")
    fmt.Println("║ No ║                         Ide                          ║         Kategori         ║ Vote  ║     Tanggal        ║")
    fmt.Println("╠════╬══════════════════════════════════════════════════════╬══════════════════════════╬═══════╬════════════════════╣")

	for i := 0; i < totalAmount; i++ {
			fmt.Printf("║ %-2d ║ %-52s ║ %-24s ║ %-5d ║ %-18s ║\n",
			ideaList[i].IdIdea,
			ideaList[i].ideaProject,
			ideaList[i].Kategori,
			ideaList[i].totalVote,
			ideaList[i].tgl.Format("2005-11-23"),
		)
	}
        fmt.Println("╚════╩══════════════════════════════════════════════════════╩══════════════════════════╩═══════╩════════════════════╝")

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
			fmt.Scanln()
			clear()
			return
		}
	}
	fmt.Println("ID tidak ditemukan.")
	fmt.Scanln()
	clear()
}

//================================ DELETE =============================
func deleteIdea() {
    clear()
	

		fmt.Println("╔═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
        fmt.Println("║                                         List Ide                                                                  ║")
        fmt.Println("╠════╦══════════════════════════════════════════════════════╦══════════════════════════╦═══════╦════════════════════╣")
        fmt.Println("║ No ║                         Ide                          ║         Kategori         ║ Vote  ║     Tanggal        ║")
        fmt.Println("╠════╬══════════════════════════════════════════════════════╬══════════════════════════╬═══════╬════════════════════╣")

	for i := 0; i < totalAmount; i++ {
			fmt.Printf("║ %-2d ║ %-52s ║ %-24s ║ %-5d ║ %-18s ║\n",
			ideaList[i].IdIdea,
			ideaList[i].ideaProject,
			ideaList[i].Kategori,
			ideaList[i].totalVote,
			ideaList[i].tgl.Format("2006-01-02"),
		)
	}
        fmt.Println("╚════╩══════════════════════════════════════════════════════╩══════════════════════════╩═══════╩════════════════════╝")

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


//============================= DUMMY VOTE =============================
func addRating() {
	clear()
	if totalAmount == 0 {
		fmt.Println("Belum ada ide yang bisa divote.")
		return
	}
		
		fmt.Println("╔═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
        fmt.Println("║                                                 List Ide                                                          ║")
        fmt.Println("╠════╦══════════════════════════════════════════════════════╦══════════════════════════╦═══════╦════════════════════╣")
        fmt.Println("║ No ║                         Ide                          ║         Kategori         ║ Vote  ║     Tanggal        ║")
        fmt.Println("╠════╬══════════════════════════════════════════════════════╬══════════════════════════╬═══════╬════════════════════╣")
	for i := 0; i < totalAmount; i++ {
			fmt.Printf("║ %-2d ║ %-52s ║ %-24s ║ %-5d ║ %-18s ║\n",
			ideaList[i].IdIdea,
			ideaList[i].ideaProject,
			ideaList[i].Kategori,
			ideaList[i].totalVote,
			ideaList[i].tgl.Format("2006-01-02"),
		)
	}
        fmt.Println("╚════╩══════════════════════════════════════════════════════╩══════════════════════════╩═══════╩════════════════════╝")

var r rating
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nMasukkan nama Anda: ")
	authorInput, _ := reader.ReadString('\n')
	r.author = strings.TrimSpace(authorInput)

	fmt.Print("Masukkan ID ide yang ingin divote: ")
	fmt.Scanln(&r.IdIdea)

	found := false
	for i := 0; i < totalAmount; i++ {
		if ideaList[i].IdIdea == r.IdIdea {
			ideaList[i].totalVote++
			found = true
			break
		}
	}

	if !found {
		fmt.Println("ID tidak ditemukan.")
		fmt.Scanln()
		return
	}

	if totalRating >= maxRating {
		fmt.Println("Kapasitas rating penuh.")
		fmt.Scanln()
		return
	}

	r.idRating = currentRatingId
	currentRatingId++
	ratingMenu[totalRating] = r
	totalRating++

	fmt.Println("Vote berhasil ditambahkan.")
	fmt.Scanln()
}

//=============================== SHOW ===============================

func showIdea() {
	clear()
	fmt.Println("\nDaftar Ide :")
	if totalAmount == 0 {
		fmt.Println("Belum ada data yang di tambahkan")
	}
		fmt.Println("╔═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
        fmt.Println("║                                                    List Ide                                                       ║")
        fmt.Println("╠════╦══════════════════════════════════════════════════════╦══════════════════════════╦═══════╦════════════════════╣")
        fmt.Println("║ No ║                         Ide                          ║         Kategori         ║ Vote  ║     Tanggal        ║")
        fmt.Println("╠════╬══════════════════════════════════════════════════════╬══════════════════════════╬═══════╬════════════════════╣")
	for i := 0; i < totalAmount; i++ {
		fmt.Printf("║ %-2d ║ %-52s ║ %-24s ║ %-5d ║ %-18s ║\n",
			ideaList[i].IdIdea,
			ideaList[i].ideaProject,
			ideaList[i].Kategori,
			ideaList[i].totalVote,
			ideaList[i].tgl.Format("2006-01-02"),
		)
	}
        fmt.Println("╚════╩══════════════════════════════════════════════════════╩══════════════════════════╩═══════╩════════════════════╝")

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


//====== Sorting Tanggal Dengan Selection Sort ======

func selectionSortByTglMenaik() {
	for pass := 0; pass < totalAmount-1; pass++ {
		idx := pass
		for i := pass + 1; i < totalAmount; i++ {
			if ideaList[i].tgl.Before(ideaList[idx].tgl) {
				idx = i
			}
		}
		ideaList[pass], ideaList[idx] = ideaList[idx], ideaList[pass]
	}

	for i := range ideaList {
		ideaList[i].IdIdea = i + 1
	}
}

func selectionSortByTglMenurun() {
	for pass := 0; pass < totalAmount-1; pass++ {
		idx := pass
		for i := pass + 1; i < totalAmount; i++ {
			if ideaList[i].tgl.After(ideaList[idx].tgl) {
				idx = i
			}
		}
		ideaList[pass], ideaList[idx] = ideaList[idx], ideaList[pass]
	}

	for i := range ideaList {
		ideaList[i].IdIdea = i + 1
	}
}


//====== Sorting Vote Dengan Insertion Sort ======

func insertionSortByVoteMenurun() {
	for i := 1; i < totalAmount; i++ {
		key := ideaList[i]
		j := i - 1

		for j >= 0 && ideaList[j].totalVote < key.totalVote {
			ideaList[j+1] = ideaList[j]
			j--
		}
		ideaList[j+1] = key
	}

	for i := range ideaList {
		ideaList[i].IdIdea = i + 1
	}
}

func insertionSortByVoteMenaik() {
	for i := 1; i < totalAmount; i++ {
		key := ideaList[i]
		j := i - 1

		for j >= 0 && ideaList[j].totalVote > key.totalVote {
			ideaList[j+1] = ideaList[j]
			j--
		}
		ideaList[j+1] = key
	}

	for i := range ideaList {
		ideaList[i].IdIdea = i + 1
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

	var pilihMenu, pilihUrutan int
	for {
		clear()
		fmt.Println("╔══════════════════════════════════════════════════╗")
		fmt.Println("║                 Pilih Urutan Ide                 ║")
		fmt.Println("╠══════════════════════════════════════════════════╣")
		fmt.Println("║ 1. Berdasarkan tanggal                           ║")
		fmt.Println("║ 2. Berdasarkan voting                            ║")
		fmt.Println("║ 3. Kembali ke menu                               ║")
		fmt.Println("╚══════════════════════════════════════════════════╝")
		fmt.Print("Pilih : ")
		fmt.Scanln(&pilihMenu)

		if pilihMenu == 1 {
			clear()
			fmt.Println("╔══════════════════════════════════════════════════╗")
			fmt.Println("║       Pilih Urutan Ide Berdasarkan Tanggal       ║")
			fmt.Println("╠══════════════════════════════════════════════════╣")
			fmt.Println("║ 1. Berdasarkan Tanggal Menaik                    ║")
			fmt.Println("║ 2. Berdasarkan Tanggal Menurun                   ║")
			fmt.Println("║ 3. Kembali                                       ║")
			fmt.Println("╚══════════════════════════════════════════════════╝")
			fmt.Scanln(&pilihUrutan)

			if pilihUrutan == 1 {
				clear()
				selectionSortByTglMenaik()
				fmt.Println("╔═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
                fmt.Println("║                                  IDE BERDASARKAN TANGGAL                                                          ║")
                fmt.Println("╠════╦══════════════════════════════════════════════════════╦══════════════════════════╦═══════╦════════════════════╣")
                fmt.Println("║ No ║                         Ide                          ║         Kategori         ║ Vote  ║     Tanggal        ║")
                fmt.Println("╠════╬══════════════════════════════════════════════════════╬══════════════════════════╬═══════╬════════════════════╣")


			}else if pilihUrutan == 2 { 
				clear()
				selectionSortByTglMenurun()
				fmt.Println("╔═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
				fmt.Println("║                                  IDE BERDASARKAN TANGGAL                                                          ║")
				fmt.Println("╠════╦══════════════════════════════════════════════════════╦══════════════════════════╦═══════╦════════════════════╣")
				fmt.Println("║ No ║                         Ide                          ║         Kategori         ║ Vote  ║     Tanggal        ║")
				fmt.Println("╠════╬══════════════════════════════════════════════════════╬══════════════════════════╬═══════╬════════════════════╣")


			}else if pilihUrutan == 3 {
				PopularIdea()
				fmt.Println("╔═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
				fmt.Println("║                                  IDE BERDASARKAN TANGGAL                                                          ║")
				fmt.Println("╠════╦══════════════════════════════════════════════════════╦══════════════════════════╦═══════╦════════════════════╣")
				fmt.Println("║ No ║                         Ide                          ║         Kategori         ║ Vote  ║     Tanggal        ║")
				fmt.Println("╠════╬══════════════════════════════════════════════════════╬══════════════════════════╬═══════╬════════════════════╣")
				return
			}else{
				fmt.Println("Pilihan Tidak Ada")
			}

		} else if pilihMenu == 2 {
			clear()
			fmt.Println("╔══════════════════════════════════════════════════╗")
			fmt.Println("║       Pilih Urutan Ide Berdasarkan Vote          ║")
			fmt.Println("╠══════════════════════════════════════════════════╣")
			fmt.Println("║ 1. Berdasarkan Vote Menaik                       ║")
			fmt.Println("║ 2. Berdasarkan Vote Menurun                      ║")
			fmt.Println("║ 3. Kembali                                       ║")
			fmt.Println("╚══════════════════════════════════════════════════╝")
			fmt.Scanln(&pilihUrutan)

			if pilihUrutan == 1 {
				clear()
				insertionSortByVoteMenaik()
				fmt.Println("╔═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
				fmt.Println("║                                     IDE BERDASARKAN VOTING                                                        ║")
				fmt.Println("╠════╦══════════════════════════════════════════════════════╦══════════════════════════╦═══════╦════════════════════╣")
				fmt.Println("║ No ║                         Ide                          ║         Kategori         ║ Vote  ║     Tanggal        ║")
				fmt.Println("╠════╬══════════════════════════════════════════════════════╬══════════════════════════╬═══════╬════════════════════╣")


			}else if pilihUrutan == 2 { 
				clear()
				insertionSortByVoteMenurun()
				fmt.Println("╔═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
				fmt.Println("║                                  IDE BERDASARKAN VOTING                                                           ║")
				fmt.Println("╠════╦══════════════════════════════════════════════════════╦══════════════════════════╦═══════╦════════════════════╣")
				fmt.Println("║ No ║                         Ide                          ║         Kategori         ║ Vote  ║     Tanggal        ║")
				fmt.Println("╠════╬══════════════════════════════════════════════════════╬══════════════════════════╬═══════╬════════════════════╣")

			}else if pilihUrutan == 3 {
				PopularIdea()
				return
			}else{
				fmt.Println("Pilihan Tidak Ada")
			}

		} else if pilihMenu == 3{
			return
		}else {
			fmt.Println("Pilihan Tidak Ada")
		}



		if pilihMenu == 1 || pilihMenu == 2 {
			for i := 0; i < totalAmount; i++ {
			fmt.Printf("║ %-2d ║ %-52s ║ %-24s ║ %-5d ║ %-18s ║\n",
					i+1,
					ideaList[i].ideaProject,
					ideaList[i].Kategori,
					ideaList[i].totalVote,
					ideaList[i].tgl.Format("2006-01-02"))
			}
			fmt.Println("╚════╩══════════════════════════════════════════════════════╩══════════════════════════╩═══════╩════════════════════╝")
			fmt.Println("\nTekan Enter untuk kembali ke pilihan...")
			fmt.Scanln()
		}
	}
}


//============================= SEARCH =============================
func searchIdea() {
	clear()
	if totalAmount == 0 {
		fmt.Println("Belum ada ide yang tersedia.")
		fmt.Scanln()
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan kata kunci ide yang dicari: ")
	keyword, _ := reader.ReadString('\n')
	keyword = strings.TrimSpace(keyword)

	found := false
	fmt.Println("\nHasil Pencarian Ide:")
	fmt.Println("╔════╦══════════════════════════════════════════════════════╦══════════════════════════╦═══════╦════════════════════╗")
	fmt.Println("║ ID ║                         Ide                          ║         Kategori         ║ Vote  ║     Tanggal        ║")
	fmt.Println("╠════╬══════════════════════════════════════════════════════╬══════════════════════════╬═══════╬════════════════════╣")
	for i := 0; i < totalAmount; i++ {
		if strings.Contains(strings.ToLower(ideaList[i].ideaProject), strings.ToLower(keyword)) {
			fmt.Printf("║ %-2d ║ %-52s ║ %-24s ║ %-5d ║ %-18s ║\n",
				ideaList[i].IdIdea,
				ideaList[i].ideaProject,
				ideaList[i].Kategori,
				ideaList[i].totalVote,
				ideaList[i].tgl.Format("2006-01-02"),
			)
			found = true
		}
	}
	fmt.Println("╚════╩══════════════════════════════════════════════════════╩══════════════════════════╩═══════╩════════════════════╝")

	if !found {
		fmt.Println("Ide tidak ditemukan.")
	}
	fmt.Scanln()
}


//======================== Popular by Periode ========================
func popularIdeaByPeriod() {
	var startStr, endStr string
	layout := "2006-01-02"

	fmt.Print("Masukkan tanggal awal (yyyy-mm-dd): ")
	fmt.Scanln(&startStr)
	fmt.Print("Masukkan tanggal akhir (yyyy-mm-dd): ")
	fmt.Scanln(&endStr)

	start, err1 := time.Parse(layout, startStr)
	end, err2 := time.Parse(layout, endStr)

	if err1 != nil || err2 != nil {
		fmt.Println("Format tanggal salah. Gunakan format yyyy-mm-dd.")
		return
	}

	// Filter ide dalam periode inklusif
	var filtered []Idea
	for i := 0; i < totalAmount; i++ {
		if (ideaList[i].tgl.Equal(start) || ideaList[i].tgl.After(start)) &&
			(ideaList[i].tgl.Equal(end) || ideaList[i].tgl.Before(end)) {
			filtered = append(filtered, ideaList[i])
		}
	}

	if len(filtered) == 0 {
		fmt.Println("Tidak ada ide dalam periode ini.")
		return
	}

	// Urutkan berdasarkan totalVote (descending)
	for i := 0; i < len(filtered)-1; i++ {
		for j := i + 1; j < len(filtered); j++ {
			if filtered[j].totalVote > filtered[i].totalVote {
				filtered[i], filtered[j] = filtered[j], filtered[i]
			}
		}
	}

	fmt.Println("\nIde paling populer dalam periode tersebut:")
	for _, idea := range filtered {
		fmt.Printf("- %s (Vote: %d, Tanggal: %s)\n", idea.ideaProject, idea.totalVote, idea.tgl.Format("2006-01-02"))
	}
	fmt.Println()

	fmt.Print("Tekan ENTER untuk kembali ke menu...")
	fmt.Scanln()
}




func menu(){
	clear()
	var pilihan int 
    for {
		clear()
		fmt.Println("╔══════════════════════════════════════════════════╗")
		fmt.Println("║                    MENU UTAMA                    ║")
		fmt.Println("╠══════════════════════════════════════════════════╣")
		fmt.Println("║ 1. Tambah Ide                                    ║")
		fmt.Println("║ 2. Voting                                        ║")
		fmt.Println("║ 3. Lihat Ide Terpopuler                          ║")
		fmt.Println("║ 4. Lihat Semua Ide                               ║")
		fmt.Println("║ 5. Search Ide                                    ║")
		fmt.Println("║ 6. Ide Popular Berdasarkan Periode               ║")
		fmt.Println("║ 7. Keluar                                        ║")
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
			searchIdea()
		case 6:
			popularIdeaByPeriod()
		case 7:
			fmt.Println("Terima kasih.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

//============================= MAIN =============================
func main() {
	dummyData()
    menu()
}