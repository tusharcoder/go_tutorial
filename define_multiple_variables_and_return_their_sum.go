/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T11:51:28+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: variables_changing_values.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-10T12:30:05+05:30
 */



package main

import "fmt"

func sum(n ...int) int{
  sum:= 0
  //we need value only, or we write key, value for both the key and value,
  // range will iterate over the values in n
  for _,value:= range n{
    sum+=value
  }
  return sum
}
func main(){

  var (
    x=10
    y=20
    //multiple assign
    c,d,e,f = 100,200,300,400
    //multiple assignment not working here this way, so next line will give error
    //a=b=100
  )

  fmt.Println("sum of the variables are",sum(x,y,c,d,e,f))
}
