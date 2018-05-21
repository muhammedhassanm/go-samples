package main

import "fmt"
import "bufio"
import "os"

func main(){

 var mat =[5][5]string{{"A","B","C","D","E"},{"F","G","H","I","J"},{"K","L","M","N","O"},
           {"P","Q","R","S","T"},{"U","V","W","X","Y"}}
		   fmt.Println("\nThe Matrix is :")
		   for i:=0; i < 5; i++ {
			   for j:=0; j < 5; j++ {
				   fmt.Print("\t", mat[i][j])
   
			   }
			   fmt.Println()
			   
		   }
	   
  scanner := bufio.NewScanner(os.Stdin)

  fmt.Println("Current posn One:")
  scanner.Scan()
  curr1:=scanner.Text()
  fmt.Println("Current Position Two")
  scanner.Scan()
  curr2:=scanner.Text()
  fmt.Println("Current Position Three")
  scanner.Scan()
  curr3:=scanner.Text()

  fmt.Println("Block posn One :")
  scanner.Scan()
  bloc1:=scanner.Text()

  fmt.Println("Block posn Two :")
  scanner.Scan()
  bloc2:=scanner.Text()

  fmt.Println("Block posn Three :")
  scanner.Scan()
  bloc3:=scanner.Text()

  var mov1,mov2,mov3 int
  mov1=0
  mov2=0
  mov3=0
  var a1x,b1x,a1y,b1y,a2x,b2x,a2y,b2y,a3x,b3x,a3y,b3y int
  a1x=-2
  a1y=-2
  b1x=6
  b1y=6

  a2x=-2
  a2y=-2
  b2x=6
  b2y=6

  a3x=-2
  a3y=-2
  b3x=6
  b3y=6

  for p:=0;p<5;p++{
    for q:=0;q<5;q++{
      if mat[p][q]==curr1{
        a1x=p
        a1y=q
	  }
	  if(mat[p][q]==curr2){
		  a2x=p
		  a2y=q
	  }
	  if(mat[p][q]==curr3){
		  a3x=p
		  a3y=q
	  }
      if (mat[p][q]==bloc1){
        b1x=p
        b1y=q
	  }
	  if (mat[p][q]==bloc2){
		  b2x=p
		  b2y=q
	  }
	  if (mat[p][q]==bloc3){
		b3x=p
		b3y=q
	}
    }
  }

  if(a1x==-2 && a2x==-2 && a3x==-2){
    return
  }
  if(a1x>0 && a1y>0 && a1x<4 && a1y<4){
	mov1=8
	

  }else if((a1x+a1y)%4==0 && (a1x==0 || a1y==0 || a1x==4 || a1y==4)){
	mov1=3
	
  }else{
	mov1=5
	
  }
  if(((a1x-b1x)==1||(b1x-a1x)==1)&&((a1y-b1y)==0||(a1y-b1y)==1||(b1y-a1y)==1)){
	mov1=mov1-1
	
  }else if(((a1y-b1y)==1||(b1y-a1y)==1)&&((a1x-b1x)==0||(a1x-b1x)==1||(b1x-a1x)==1)){
	mov1=mov1-1
  }

  //
  if(a2x>0 && a2y>0 && a2x<4 && a2y<4){
	mov2=8
	

  }else if((a2x+a2y)%4==0 && (a2x==0 || a2y==0 || a2x==4 || a2y==4)){
	mov2=3
	
  }else{
	mov2=5
	
  }
  if(((a2x-b2x)==1||(b2x-a2x)==1)&&((a2y-b2y)==0||(a2y-b2y)==1||(b2y-a2y)==1)){
	mov2=mov2-1
	
  }else if(((a2y-b2y)==1||(b2y-a2y)==1)&&((a2x-b2x)==0||(a2x-b2x)==1||(b2x-a2x)==1)){
	mov2=mov2-1
  }
  //
  if(a3x>0 && a3y>0 && a3x<4 && a3y<4){
	mov3=8
	

  }else if((a3x+a3y)%4==0 && (a3x==0 || a3y==0 || a3x==4 || a3y==4)){
	mov3=3

	
  }else{
	mov3=5
	
  }
  if(((a3x-b3x)==1||(b3x-a3x)==1)&&((a3y-b3y)==0||(a3y-b3y)==1||(b3y-a3y)==1)){
	mov3=mov3-1
	
  }else if(((a3y-b3y)==1||(b3y-a3y)==1)&&((a3x-b3x)==0||(a3x-b3x)==1||(b3x-a3x)==1)){
	mov3=mov3-1
  }
  fmt.Println("First Pos Movement",mov1)
  fmt.Println("Second Pos Movement",mov2)
  fmt.Println("Third Pos Movement",mov3)

}