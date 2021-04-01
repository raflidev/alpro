package main
import "fmt"

func main(){
  var a,b,permutasi int
  permutasi = 1
  fmt.Scanln(&a,&b)
  for true {
    permutasi = permutasi * (Faa(a)/Faa(a-b))
    fmt.Scanln(&a,&b)
    if a == 0 && b == 0 {
      break
    }
  }
  // fmt.Println(Faa(a)/Faa(a-b))
  fmt.Println(permutasi)
}
func Faa(x int) int{
  var y int
  y = 1
  for i:=1; i <= x; i++{
    y = y * i
  }
  return y
}