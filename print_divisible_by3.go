/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T16:42:56+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: print_divisible_by3.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-10T16:48:48+05:30
 */



package main
import "fmt"

func main(){
  for i := 1; i <=100; i++ {
    if i%3==0 {
      fmt.Printf("%d\t",i)
    }
  }
}
