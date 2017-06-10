/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T16:37:45+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: switch.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-10T16:42:12+05:30
 */



package main
import "fmt"

func main(){
  for i := 0; i < 10; i++ {
      switch i {
      case 0:
          fmt.Printf("ZERO\n")
      case 1:
          fmt.Printf("One\n")
      case 2:
          fmt.Printf("TWO\n")
      case 3:
          fmt.Printf("THREE\n")
      case 4:
          fmt.Printf("FOUR\n")
      case 5:
          fmt.Printf("FIVE\n")
      case 6:
          fmt.Printf("SIX\n")
      case 7:
          fmt.Printf("SEVEN\n")
      case 8:
          fmt.Printf("EIGHT\n")
      case 9:
          fmt.Printf("NINE\n")
      }
  }
}
