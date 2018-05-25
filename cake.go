package main

import (
	"fmt"
)
func main(){
	var n int
	var a [10]int
	
	fmt.Println("Number Of Candles On The Cake\n")
	fmt.Scan(&n)
	fmt.Println("Enter The Height Of Each Candles")
	for i:=0; i<n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println()

	count := 0
	max := a[0]
	for i := 0; i < n; i++ {
        if (max < a[i]){
			max = a[i]
			
		}
		if(a[i]==max){
			count++
		}
	}
fmt.Println("\nNomber Of Maximum Candles To Blow",count)
   

}
