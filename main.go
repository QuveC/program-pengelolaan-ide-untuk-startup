package main

import (
	"os"	
	"fmt"
	"time"
	"os/exec"
	"strings"
)

const maxIdea = 20
const maxRating int = 100

func dummyData(ideaList *ideas, ratingMenu *ratingList, totalAmount *int, currentId *int, totalRating *int, currentRatingId *int) {
	ideaList[0] = Idea{IdIdea: 1, ideaProject: "Aplikasi Donasi Online", Kategori: "Sosial", totalVote: 5, tgl: time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC)}
	ideaList[1] = Idea{IdIdea: 2, ideaProject: "Sistem Absensi QR", Kategori: "Teknologi", totalVote: 8, tgl: time.Date(2015, 3, 22, 0, 0, 0, 0, time.UTC)}
	ideaList[2] = Idea{IdIdea: 3, ideaProject: "Aplikasi Belajar Bahasa", Kategori: "Pendidikan", totalVote: 3, tgl: time.Date(2009, 5, 5, 0, 0, 0, 0, time.UTC)}
	ideaList[3] = Idea{IdIdea: 4, ideaProject: "Marketplace UMKM", Kategori: "Ekonomi", totalVote: 10, tgl: time.Date(2005, 2, 15, 0, 0, 0, 0, time.UTC)}
	*totalAmount = 4
	*currentId = 5

	ratingMenu[0] = rating{idRating: 1, author: "Budi", IdIdea: 2}
	ratingMenu[1] = rating{idRating: 2, author: "Sari", IdIdea: 4}
	ratingMenu[2] = rating{idRating: 3, author: "Ani", IdIdea: 1}
	ratingMenu[3] = rating{idRating: 4, author: "Dewi", IdIdea: 4}
	ratingMenu[4] = rating{idRating: 5, author: "Rudi", IdIdea: 2}
	ratingMenu[5] = rating{idRating: 6, author: "Andi", IdIdea: 4}
	ratingMenu[6] = rating{idRating: 7, author: "Putri", IdIdea: 2}
	ratingMenu[7] = rating{idRating: 8, author: "Agus", IdIdea: 4}
	*totalRating = 8
	*currentRatingId = 9
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

var ideaList ideas
var ratingMenu ratingList

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

//============================= Read Data =============================
func readData(ideaList ideas, totalAmount int){
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
			ideaList[i].tgl.Format("2006-01-02"),
		)
	}
	fmt.Println("╚════╩══════════════════════════════════════════════════════╩══════════════════════════╩═══════╩════════════════════╝")
}

func readPopularIdea(judul string) {
	totalWidth := 115
	judulLen := len(judul)
	leftPadding := (totalWidth - judulLen) / 2
	rightPadding := totalWidth - judulLen - leftPadding
	fmt.Println("╔═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Printf(" ║%*s%s%*s║\n", leftPadding, "", judul, rightPadding, "")
	fmt.Println("╠════╦══════════════════════════════════════════════════════╦══════════════════════════╦═══════╦════════════════════╣")
	fmt.Println("║ No ║                         Ide                          ║         Kategori         ║ Vote  ║     Tanggal        ║")
	fmt.Println("╠════╬══════════════════════════════════════════════════════╬══════════════════════════╬═══════╬════════════════════╣")
}

//================================ CREATE =============================
func createIdea(ideaList *ideas, totalAmount *int, currentId *int) {
	clear()
	if *totalAmount >= maxIdea {
		fmt.Println("Kapasitas Ide Penuh")
		fmt.Scanln()
		clear()
		return
	}
	var newIdea Idea

	newIdea.IdIdea = *currentId
	(*currentId)++

	fmt.Print("Masukkan Ide: ")
	fmt.Scanln(&newIdea.ideaProject)
	
	fmt.Print("Masukkan Kategori: ")
	fmt.Scanln(&newIdea.Kategori)

	newIdea.tgl = time.Now()
	newIdea.totalVote = 0

	ideaList[*totalAmount] = newIdea
	*totalAmount++

	fmt.Println("Ide Berhasil Ditambahkan")
	fmt.Scanln()
	clear()
}

//================================ UPDATE =============================
func updateIdea(ideaList *ideas, totalAmount int) {
	clear()
	readData(*ideaList, totalAmount)

	var id int
	fmt.Println("")
	fmt.Print("Masukkan ID Ide Yang Ingin Diupdate: ")
	fmt.Scanln(&id)

	for i := 0; i < totalAmount; i++ {
		if ideaList[i].IdIdea == id {
			fmt.Print("Masukkan Ide Baru: ")
			fmt.Scanln(&ideaList[i].ideaProject)
			
			fmt.Print("Masukkan Kategori Baru: ")
			fmt.Scanln(&ideaList[i].Kategori)
			
			ideaList[i].tgl = time.Now()	

			fmt.Println("Ide Berhasil Diupdate.")
			fmt.Scanln()
			clear()
			return
		}
	}
	fmt.Println("ID Tidak Ditemukan.")
	fmt.Scanln()
	clear()
}

//================================ DELETE =============================
func deleteIdea(ideaList *ideas, totalAmount *int) {
	clear()
	readData(*ideaList, *totalAmount)

	var id int
	fmt.Println("")
	fmt.Print("Masukan ID Ide Yang Ingin Dihapus: ")
	fmt.Scan(&id)

	foundIndex := -1
	for i := 0; i < *totalAmount; i++ {
		if ideaList[i].IdIdea == id {
			foundIndex = i
		}
	}

	if foundIndex != -1 {
		for j := foundIndex; j < *totalAmount-1; j++ {
			ideaList[j] = ideaList[j+1]
		}
		*totalAmount--

		for i := 0; i < *totalAmount; i++ {
			ideaList[i].IdIdea = i + 1
		}
		fmt.Println("Ide Berhasil Dihapus")
	} else {
		fmt.Println("ID Tidak Ditemukan.")
	}

	fmt.Scanln()
}


//============================= ADD RATING =============================
func addRating(ideaList *ideas, ratingMenu *ratingList, totalRating *int, currentRatingId *int, totalAmount int) {
	clear()

	if totalAmount == 0 {
		fmt.Println("Belum Ada Ide Yang Bisa Divote.")
		fmt.Scanln()
		return
	}

	readData(*ideaList, totalAmount)
	fmt.Print("Kembali: cancel")

	var r rating
	var inputAuthor string
	var inputId int

	fmt.Print("\nMasukkan Nama Anda: ")
	fmt.Scanln(&inputAuthor)
	if inputAuthor == "cancel" {
		return
	}
	r.author = inputAuthor

	fmt.Print("Masukkan ID Ide Yang Ingin Divote: ")
	fmt.Scanln(&inputId)
	if inputId == 0 {
		return
	}
	r.IdIdea = inputId

	if *totalRating >= maxRating {
		fmt.Println("Kapasitas Rating Penuh.")
		fmt.Scanln()
		return
	}

	found := false

	for i := 0; i < totalAmount; i++ {
		if ideaList[i].IdIdea == r.IdIdea {
			ideaList[i].totalVote++
			found = true
		}
	}

	if found {
		r.idRating = *currentRatingId
		(*currentRatingId)++
		ratingMenu[*totalRating] = r
		(*totalRating)++
		fmt.Println("Vote Berhasil Ditambahkan.")
	} else {
		fmt.Println("ID Tidak Ditemukan.")
	}

	fmt.Scanln()
}


//=============================== CRUD LIST IDEA ===============================
func menuCrudListIdea(ideaList *ideas, totalAmount *int, currentId *int) {
	clear()
	fmt.Println("\nDaftar Ide :")
	if *totalAmount == 0 {
		fmt.Println("Belum Ada Data Yang DiTambahkan")
	} else {
		readData(*ideaList, *totalAmount)
	}

	var pilih int
	fmt.Println("\n=== Pilihan ===")
	fmt.Println("1. Create")
	fmt.Println("2. Update")
	fmt.Println("3. Delete")
	fmt.Println("4. Kembali")
	fmt.Print("Pilih: ")
	fmt.Scanln(&pilih)

	switch pilih {
	case 1:
		createIdea(ideaList, totalAmount, currentId)
	case 2:
		updateIdea(ideaList, *totalAmount)
	case 3:
		deleteIdea(ideaList, totalAmount)
	case 4:
		return
	default:
		fmt.Println("Pilihan Tidak Valid.")
		fmt.Scanln()
	}
}

//====== Sorting Tanggal Dengan Selection Sort ======
func selectionSortByTglMenaik(ideaList *ideas, totalAmount int) {
	for pass := 0; pass < totalAmount-1; pass++ {
		idx := pass
		for i := pass + 1; i < totalAmount; i++ {
			if ideaList[i].tgl.Before(ideaList[idx].tgl) {
				idx = i
			}
		}
		ideaList[pass], ideaList[idx] = ideaList[idx], ideaList[pass]
	}

	for i := 0; i < totalAmount; i++ {
		ideaList[i].IdIdea = i + 1
	}
}

func selectionSortByTglMenurun(ideaList *ideas, totalAmount int) {
	for pass := 0; pass < totalAmount-1; pass++ {
		idx := pass
		for i := pass + 1; i < totalAmount; i++ {
			if ideaList[i].tgl.After(ideaList[idx].tgl) {
				idx = i
			}
		}
		ideaList[pass], ideaList[idx] = ideaList[idx], ideaList[pass]
	}

	for i := 0; i < totalAmount; i++ {
		ideaList[i].IdIdea = i + 1
	}
}

//====== Sorting Vote Dengan Insertion Sort ======
func insertionSortByVoteMenurun(ideaList *ideas, totalAmount int) {
	for i := 1; i < totalAmount; i++ {
		key := ideaList[i]
		j := i - 1

		for j >= 0 && ideaList[j].totalVote < key.totalVote {
			ideaList[j+1] = ideaList[j]
			j--
		}
		ideaList[j+1] = key
	}

	for i := 0; i < totalAmount; i++ {
		ideaList[i].IdIdea = i + 1
	}
}

func insertionSortByVoteMenaik(ideaList *ideas, totalAmount int) {
	for i := 1; i < totalAmount; i++ {
		key := ideaList[i]
		j := i - 1

		for j >= 0 && ideaList[j].totalVote > key.totalVote {
			ideaList[j+1] = ideaList[j]
			j--
		}
		ideaList[j+1] = key
	}

	for i := 0; i < totalAmount; i++ {
		ideaList[i].IdIdea = i + 1
	}
}

func PopularIdea(ideaList *ideas, totalAmount int) {
	clear()

	if totalAmount == 0 {
		fmt.Println("Belum Ada ide Yang Bisa Ditampilkan.")
		fmt.Scanln()
		return
	}

	var pilihMenu, pilihUrutan int
	for {
		clear()
		fmt.Println("╔══════════════════════════════════════════════════╗")
		fmt.Println("║                 Pilih Urutan Ide                 ║")
		fmt.Println("╠══════════════════════════════════════════════════╣")
		fmt.Println("║ 1. Berdasarkan Tanggal                           ║")
		fmt.Println("║ 2. Berdasarkan Voting                            ║")
		fmt.Println("║ 3. Kembali Ke Menu                               ║")
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
				selectionSortByTglMenaik(ideaList, totalAmount)
				readPopularIdea("IDE BERDASARKAN TANGGAL MENAIK")
			} else if pilihUrutan == 2 {
				clear()
				selectionSortByTglMenurun(ideaList, totalAmount)
				readPopularIdea("IDE BERDASARKAN TANGGAL MENURUN")
			} else if pilihUrutan == 3 {
				return
			} else {
				fmt.Println("Pilihan Tidak Ada")
				fmt.Scanln()
				return
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
				insertionSortByVoteMenaik(ideaList, totalAmount)
				readPopularIdea("IDE BERDASARKAN VOTE MENAIK")
			} else if pilihUrutan == 2 {
				clear()
				insertionSortByVoteMenurun(ideaList, totalAmount)
				readPopularIdea("IDE BERDASARKAN VOTE MENURUN")
			} else if pilihUrutan == 3 {
				return
			} else {
				fmt.Println("Pilihan Tidak Ada")
				fmt.Scanln()
				return
			}

		} else if pilihMenu == 3 {
			return
		} else {
			fmt.Println("Pilihan Tidak Ada")
			fmt.Scanln()
			return
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
			fmt.Scanln()
		}
	}
}

//============================= SEARCH =============================
func searchIdea(ideaList *ideas, totalAmount int) {
	clear()
	if totalAmount == 0 {
		fmt.Println("Belum Ada Ide Yang Tersedia")
		fmt.Scanln()
		return
	}
	var keyword string

	fmt.Print("Masukkan Kata Kunci Ide Yang Dicari: ")
	fmt.Scanln(&keyword)

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
		fmt.Println("Ide Tidak Ditemukan.")
	}
	fmt.Scanln()
}

//======================== Popular by Periode ========================
func popularIdeaByPeriod(ideaList *ideas, totalAmount int) {
	clear()
	var startStr, endStr string
	layout := "2006-01-02"

	fmt.Print("Masukkan Tanggal Awal (yyyy-mm-dd): ")
	fmt.Scanln(&startStr)
	fmt.Print("Masukkan Tanggal Akhir (yyyy-mm-dd): ")
	fmt.Scanln(&endStr)

	start, err1 := time.Parse(layout, startStr)
	end, err2 := time.Parse(layout, endStr)

	if err1 != nil || err2 != nil {
		fmt.Println("Format Tanggal Salah")
		fmt.Scanln()
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
		fmt.Println("Tidak Ada Ide Dalam Periode Ini")
		fmt.Scanln()
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

	fmt.Println("\nIde Paling Populer Dalam Periode Tersebut:")
	fmt.Println("╔═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                     List Ide                                                                      ║")
	fmt.Println("╠════╦══════════════════════════════════════════════════════╦══════════════════════════╦═══════╦════════════════════╣")
	fmt.Println("║ No ║                         Ide                          ║         Kategori         ║ Vote  ║     Tanggal        ║")
	fmt.Println("╠════╬══════════════════════════════════════════════════════╬══════════════════════════╬═══════╬════════════════════╣")

	for i := 0; i < len(filtered); i++ {
		idea := filtered[i]
		fmt.Printf("║ %-2d ║ %-52s ║ %-24s ║ %-5d ║ %-18s ║\n",
			idea.IdIdea,
			idea.ideaProject,
			idea.Kategori,
			idea.totalVote,
			idea.tgl.Format("2006-01-02"),
		)
	}

	fmt.Println("╚════╩══════════════════════════════════════════════════════╩══════════════════════════╩═══════╩════════════════════╝")

	fmt.Println()
	fmt.Scanln()
}

func menu(ideaList *ideas, ratingMenu *ratingList, totalAmount *int, currentId *int, totalRating *int, currentRatingId *int) {
	var pilihan int
	for {
		clear()
		fmt.Println("╔══════════════════════════════════════════════════╗")
		fmt.Println("║                    MENU UTAMA                    ║")
		fmt.Println("╠══════════════════════════════════════════════════╣")
		fmt.Println("║ 1. Tambah Ide                                    ║")
		fmt.Println("║ 2. Voting                                        ║")
		fmt.Println("║ 3. Filter Ide                                    ║")
		fmt.Println("║ 4. Lihat Semua Ide                               ║")
		fmt.Println("║ 5. Search Ide                                    ║")
		fmt.Println("║ 6. Ide Popular Berdasarkan Periode               ║")
		fmt.Println("║ 7. Keluar                                        ║")
		fmt.Println("╚══════════════════════════════════════════════════╝")
		fmt.Print("Pilih Menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			createIdea(ideaList, totalAmount, currentId)
		case 2:
			addRating(ideaList, ratingMenu, totalRating, currentRatingId, *totalAmount)
		case 3:
			PopularIdea(ideaList, *totalAmount)
		case 4:
			menuCrudListIdea(ideaList, totalAmount, currentId)
		case 5:
			searchIdea(ideaList, *totalAmount)
		case 6: 	
			popularIdeaByPeriod(ideaList, *totalAmount)
		case 7:
			fmt.Println("Terima Kasih")
			return
		default:
			fmt.Println("Pilihan Tidak Valid")
			fmt.Scanln()
		}
	}
}

//============================= MAIN =============================
func main() {
	var totalAmount, currentId, totalRating, currentRatingId int
	dummyData(&ideaList, &ratingMenu, &totalAmount, &currentId, &totalRating, &currentRatingId)
	menu(&ideaList, &ratingMenu, &totalAmount, &currentId, &totalRating, &currentRatingId)
}