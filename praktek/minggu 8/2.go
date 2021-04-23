package main 
import "fmt"
func main(){
  var angka int
  var c1,c2,c3 int
  angka = 0
  for angka != -1{
    fmt.Scan(&angka)
    hitungSuara(angka,&c1,&c2,&c3)
  }
  fmt.Println(c1,c2,c3)
}

func hitungSuara(suara int,c1,c2,c3 *int){
 if suara == 1{
      *c1 += 1
    }else if suara == 2{
      *c2 += 1
    }else if suara == 3{
      *c3 += 1
    }
}