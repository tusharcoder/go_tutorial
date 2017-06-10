/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T12:30:22+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: take_input_from_user_and_doubles_it.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-10T18:13:33+05:30
 */



package main
import "fmt"
func main(){
  fmt.Println("input please")
  //float64 type float number
  var input float64
  //%f to take float type input to the address of input i.e. &input
  fmt.Scanf("%f",&input)
  //to format any value use Printf in plate of Println
  fmt.Printf("double is %f\n",input*2)
}
