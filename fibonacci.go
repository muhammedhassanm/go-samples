package main
import "fmt"

func main() {
	var n int
	var sum int
	t1 := 0
	t2 := 1
	fmt.Println("Enter The Limit")
	fmt.Scan(&n)
	fmt.Println("First ", n , "terms in the series")

	for i:= 1; i<= n; i++ {
		fmt.Println("\n ",t1)
		sum = t1 + t2
		t1=t2
		t2=sum

	}

}