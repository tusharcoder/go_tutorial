/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-12T13:06:50+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: use_new_keyword_to_get_pointer.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-12T13:12:29+05:30
 */



//in this program we use the new operator to get the pointer to the particular variable location
package main
import "fmt"
import "strconv"
func makeZero(i *int){
*i=0
}
func main(){
  i:=new(int) //pointer the ihe newly allocated memory for the int
  *i=5
  fmt.Println("i value is "+strconv.Itoa(*i)+" for now")
  makeZero(i) //since it is a pointer we dont need a & before it
  fmt.Println("i value changes to "+strconv.Itoa(*i))
}
