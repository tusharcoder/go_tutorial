/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T23:56:11+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: use_defer.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-11T00:01:04+05:30
 */


//what ever the statements comes after defer
//it will be executed at the end of the function
//i.e. main function in this case
//here it deferred the second() function call executes the rest of the function i.e.
//first() and then atlast use the second()
//so use the defer to closed the open file, recover panic i.e. try catch in other languages
package main
import "fmt"
func first(){
  fmt.Println("first")
}
func second(){
  fmt.Println("second")
}
func main(){
  defer second()
  first()
}
