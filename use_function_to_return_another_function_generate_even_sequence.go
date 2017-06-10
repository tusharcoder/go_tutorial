/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T23:35:55+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: use_closure_to_add_values.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-10T23:47:51+05:30
 */



package main
import "fmt"
// this function will behave as the generator
func generateEvenSeq() func() int{
  i:=0
  return func() (ret int){
    ret = i
    i+=2
    return
  }
}
func main(){
  even_next := generateEvenSeq()
  //generating the sequence of the even numbers
  for index:= 0;  index< 100; index++ {
    fmt.Printf("%d\t",even_next())
  }
  fmt.Printf("\n")
}
