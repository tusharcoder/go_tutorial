/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T16:48:57+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: fizzbuzz.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-10T16:53:53+05:30
 */




 package main
 import "fmt"

 func main(){
   for i := 1; i <=100; i++ {
     if i%3==0 {
       if i%5==0 {
          fmt.Printf("FIZZBUZZ\t")
       }
       fmt.Printf("FIZZ\t")
     }
     if i%5==0 {
         fmt.Printf("BUZZ\t")
     }

   }
   fmt.Println()
 }
