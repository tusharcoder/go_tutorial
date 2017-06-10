/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T18:48:05+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: find_smallest_num_in_list.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-10T23:00:37+05:30
 */



package main
import "fmt"

func main(){
  //define a list i.e. slice
  x:=[]int{
    48,96,86,68,
  57,82,63,70,
  37,34,83,27,
  19,97, 9,17,
  }
  
  small:=x[0]
  for i := 0; i < len(x); i++ {
    if small>x[i]{
      small=x[i]
    }
  }
  fmt.Printf("Smallest number in the given list: %d\n",small)
}
