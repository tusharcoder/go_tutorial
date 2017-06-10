/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T12:04:50+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: outside_var_shared_with_other_function_and_used_as_global.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-10T12:11:25+05:30
 */



package main
import "fmt"
var x int = 5

func incrementX(n int){
  //here n is the argument
  x+=n
}

func main(){
incrementX(10)
//this will print the changed value of x
fmt.Println("Incremented value is",x)
}
