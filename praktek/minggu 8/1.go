package main 
import "fmt"
func main(){
  var a,b,c int
  fmt.Scanln(&a,&b,&c)
  if b == c {
    fmt.Println("Kita tidak pernah bertemu")
  }else if a % b == 0 && a % c == 0 {
    fmt.Println("kita akan bertemu",c,"hari lagi")
  }else if a % b == 0 || a % c == 0 {
    fmt.Println("kita akan bertemu hari ini")
  }else{
    j:= a % c
    fmt.Println("Kita akan bertemu",j,"hari lagi")
  }
}