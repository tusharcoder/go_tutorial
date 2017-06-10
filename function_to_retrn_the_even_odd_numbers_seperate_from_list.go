/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T22:44:26+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: function_to_retrn_the_even_odd_numbers_seperate_from_list.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-10T23:21:59+05:30
 */



package main
import "fmt"
//function to take the array and return the seperate arrays of the odd and even arrays and also if there is any error
func sep_even_odd(x []int) ([]int,[]int){
  var even []int
  var odd []int
  for _,value := range x{
    if value % 2==0{
      even = append(even,value)
    }else{
      odd = append(odd,value)
    }
  }
  return even, odd
}

//same function but take aas many as ints at the function call but not []int
func sep_even_odd_variadic(x ...int) ([]int,[]int){
  var even []int
  var odd []int
  for _,value := range x{
    if value % 2==0{
      even = append(even,value)
    }else{
      odd = append(odd,value)
    }
  }
  return even, odd
}

func main(){
  x:=make([]int,100)
  for index,_:=range x{
    x[index]=index+1
  }
  //brackets used for scoping the even odd
  {  even, odd:=sep_even_odd(x)
    fmt.Printf("even numbers\n")
    for _,val := range even{
      fmt.Printf("%d\t",val)
    }
    fmt.Printf("\nodd numbers\n")
    for _,val := range odd{
      fmt.Printf("%d\t",val)
    }
  }
  fmt.Printf("\n")
  for i := 0; i < 70; i++ {
    fmt.Printf("==")
  }
  fmt.Printf("\nUsing the Variadic function for anyu number of ints passed\n")
  //if even odd was not scoped then this later line of code gives error, of no new variable at the left of :=
  even,odd:=sep_even_odd_variadic(1,2,3,4,5,6,7,8,9,0)
  //printing even odd again
  fmt.Printf("even numbers\n")
  for _,val := range even{
    fmt.Printf("%d\t",val)
  }
  fmt.Printf("\nodd numbers\n")
  for _,val := range odd{
    fmt.Printf("%d\t",val)
  }
  fmt.Printf("\n\n")
}
