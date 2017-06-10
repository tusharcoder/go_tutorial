/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T23:35:55+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: use_closure_to_add_values.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-10T23:40:37+05:30
 */



package main
import "fmt"

func main(){
  add := func(args ...int) (total int){
    for _,val:= range args {
      total+=val
    }
    //total is returned
    return
  }
  fmt.Println(add(1,2,3,4,5,6,7,8))
}
