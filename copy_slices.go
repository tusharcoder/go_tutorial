/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T17:07:50+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: slice.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-10T17:28:56+05:30
 */



package main
import "fmt"

func main(){
slice1:=[]float64{1,2,3,4,5}
slice2:=make([]float64,3)
//copy the slice, but only 3 elements upto 3 are copied as due to the size limitation
copy(slice2,slice1)
fmt.Println("slice1:",slice1)
fmt.Println("slice2:",slice2)
}
//slice can be considered the subarray of the arrays but their size can be changed, to max to the length of the array they point to
