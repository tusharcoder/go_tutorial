/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T17:07:50+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: slice.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-10T17:24:03+05:30
 */



package main
import "fmt"

func main(){

  slice1:=[]int{1,2,3,4,5}
  //appending the elements 6,7 in slice 1 and assign it to slice 2
  slice2:=append(slice1,6,7)
  fmt.Println(slice2)
}
