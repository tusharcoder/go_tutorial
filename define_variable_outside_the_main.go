/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T11:51:28+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: variables_changing_values.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-10T12:04:26+05:30
 */



package main

import "fmt"

var y string="Hello world"
//we can not use this here
//y:="Hello world"
func main(){
  fmt.Println("Outside variable is",y)
}
