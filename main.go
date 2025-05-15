package main
import "fmt"

const maxIdea = 100

type Idea struct {
	ID     int
	Title  string
	Upvote int
}

var ideas [maxIdea]Idea

func createIdea(){
  
   
}

func updateIdea(){

}

func deleteIdea(){

}

func upVote(){

}

func popularIdea(){
     
}

func main() {
	var pilihan int

	for {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Tambah Ide")
		fmt.Println("2. Voting")
		fmt.Println("3. List Popular Idea")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			createIdea()
		} else if pilihan == 2 {
      popularIdea()
		} else if pilihan == 3 {
			popularIdea()
		}else if pilihan == 4{
      		fmt.Println("Terima kasih.")
			break
    }else {
      fmt.Println("Pilihan tidak ada")
    }
  }
}