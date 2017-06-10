/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T11:51:28+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: variables_changing_values.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-10T12:00:40+05:30
 */



package main

import "fmt"

func main(){
x:="Hello"
y:="Hello"
fmt.Println("x is:",x)
fmt.Println("y is:",y)
//adding the integer this time
//cannot change the type of the already defined variable, rather assign new value of the same type\
//so later statement will result in error
//y:=5
//so we use the new variable
z:=5
fmt.Println("z is:",z)


}
