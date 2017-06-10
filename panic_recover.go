/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-11T00:01:38+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: panic_recover.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-11T00:04:04+05:30
 */



package main

import "fmt"
func main(){
  //defered self calling function will executed at the end of the main function,hence catch the errro when occu
  defer func(){
    //recover from the pani, catch the error
    str:=recover()
    fmt.Println(str)
  }()
  panic("PANIC")
}
