/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T16:58:03+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: array_example_and_computre_average.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-10T17:07:43+05:30
 */



package main
import "fmt"

func main(){
  //declare array and total variable
  //we cnnot change the length of the array, so for that we have to use the slice
  numbers, total:= [5]float64{
    89,
    91,
    92,
    93,
    95,
  }, float64(0) // we type convert the 0 to float64 or it will give the error in computing the total of numbers
  //computing the average
  //as we dont need the key so we use _ as go dont allow any unused variable
  for _,value:=range numbers{
    total+=value
  }
  //type conversion to float64 of the length or it will give the error
  average := total/float64(len(numbers))
  fmt.Printf("Average is %f\n",average)
}
