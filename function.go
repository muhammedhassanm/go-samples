package main
import "fmt"
var z int
func add(x int,y int)int {

  z=x+y
  return z
}
func sub(x int,y int)int {

  z=x-y
  return z

}

func main(){
  add(2,3)
  fmt.Println("Sum=\t",z)
  sub(5,2)
  fmt.Println("Subtraction=\t",z)

}
