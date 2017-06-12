/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-12T12:49:39+05:30
 * @Email:  tamyworld@gmail.com
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-12T13:05:25+05:30
 */

//go lang supports pointers
//this is the function  take the argument by using the pointer and changes the value of the original variable to 0
package main
import "fmt"
import "strconv"

func makeZero(i *int) {
*i=0
}

func main(){
  // var i int
  i:=5
  fmt.Println("i now is "+strconv.Itoa(i)) //use Itoa to convert the integer to string
  makeZero(&i) //address the i passsed to the function makeZero
  //print the i value giva zero before it was 5
  fmt.Printf("i changes to %d\n",i)
}
