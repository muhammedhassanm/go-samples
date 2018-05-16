package main
import "fmt"


func main(){
var n int
fmt.Print("Enter The Number To Print The Multiplication Table")
fmt.Scan(&n)
  for i := 1; i <= 10; i++ {
    fmt.Println("| ",i,"| * |",n,"=",i*n)
    fmt.Println("--------------")
  }

}
