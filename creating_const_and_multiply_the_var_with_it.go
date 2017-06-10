/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T11:51:28+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: variables_changing_values.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-10T12:15:40+05:30
 */



package main

import "fmt"

//using the const keyword to create the constant
const multiplier = 10

func multiply(n int) int{
  //this takes the input n and returned the multiplied value with the n to the main function
  //later int is the type of the value returned by the function
  return n*multiplier
}
func main(){
fmt.Println("multiple 2 with multiplier i.e. 10 returns", multiply(2))
}
