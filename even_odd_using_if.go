/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T16:20:30+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: print_num_using_for_loop.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-10T16:37:38+05:30
 */

package main

import "fmt"

func main(){
  for i:=1;i<=10;i++{
    if i%2==0 {
      fmt.Printf("i : %d is even\n",i)
    }else{
      fmt.Printf("i : %d is odd\n",i)
    }
  }
}
//simillarly we can do if else if
