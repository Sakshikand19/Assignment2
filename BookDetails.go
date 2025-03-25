package main
import "fmt"
type Books struct{
	title string
	author string
	subject string
	book_id int
}
func main(){
	var Book1 Books
	var Book2 Books
	
	Book1.title="Go programming"
	Book1.author="Mhaesh kumar"
	Book1.subject="Go Programming Tutorial"
	Book1.book_id=324343

	Book2.title="Telecom billing"
	Book2.author="zara ali"
	Book2.subject="telecom billing tutorial"
	Book2.book_id=232334

	fmt.Printf("Book 1 title: &s\n",Book1.title)
	fmt.Printf("Book 1 author: &s\n",Book1.author)
	fmt.Printf("Book 1 subject: &s\n",Book1.subject)
	fmt.Printf("Book 1 Bookid: &s\n",Book1.book_id)

	fmt.Printf("Book 2 title: &s\n",Book1.title)
	fmt.Printf("Book 2 title: &s\n",Book1.author)
	fmt.Printf("Book 2 title: &s\n",Book1.subject)
	fmt.Printf("Book 2 title: &s\n",Book1.book_id)

}