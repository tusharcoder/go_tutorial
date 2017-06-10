/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T17:07:50+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: slice.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-10T17:18:54+05:30
 */



package main
import "fmt"

func main(){
  //declare the slice
  //var x []float64
  //if we give the length then it will be an array not the slice eg: var x[5]float64
  //use the builtin make function to make a type, lets make a slice with make function
  //slice always be smaller than the array, to which it points as it is the part of the array
  //second argument is the size of the slice and third is the size of the array to which it points
  // var x:=make([]float64, 5, 10)
  //suppose there is a array of range from 10 to 100
  y:=[10]float64{1,2,3,4,5,6,7,8,9,10,}
  x:=y[0:5] //0 is the min and 5 is the max it will print values in index 0,5 but exclude high index i.e. value at index 5 i.e. 6
  for _,val:=range x{
    fmt.Printf("%f\t",val)
  }

}
