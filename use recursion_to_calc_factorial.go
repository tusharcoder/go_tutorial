/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T23:35:55+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: use_closure_to_add_values.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-10T23:52:26+05:30
 */



package main
import "fmt"
func factorial(x int) (int){
  if x==0{
    return 1
  }
  return x*factorial(x-1)
}
func main(){
fmt.Printf("factorial of %d is %d\n",9,factorial(9))
}
