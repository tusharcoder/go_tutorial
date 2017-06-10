/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T22:44:26+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: function_to_retrn_the_even_odd_numbers_seperate_from_list.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-10T23:31:38+05:30
 */



package main
import "fmt"
func add(args ...int) (total int){
for _,val:= range args{
  total+=val
  }
  //it is understood that total is returned as declared in the above return type
  return
}
func main(){
  xs:=[]int{1,2,3,4,5,6,7,8,9}
  fmt.Println(add(xs...))
}
