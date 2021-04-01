package main
import "fmt"

func main(){
  var a,b,kombinasi,hasil int
  kombinasi = 1
  fmt.Scanln(&a,&b)
  for true {
    kombinasi = kombinasi * (Faaa(a)/Faaa(a-b)/Faaa(b))
    fmt.Scanln(&a,&b)
    if a == 0 && b == 0 {
      hasil = hasil + kombinasi
      kombinasi = 1
      fmt.Println("Kombinasi sementara:",hasil)
    }else if a < 0 && b < 0 {
      break
    }
  }
  if kombinasi > 1 {
    fmt.Println("Kombinasi:",kombinasi)
  }else if hasil > 0 {
    fmt.Println("Hasil multiple kombinasi:",hasil)
  }
}
func Faaa(x int) int{
  var y int
  y = 1
  for i:=1; i <= x; i++{
    y = y * i
  }
  return y
}