/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T18:13:38+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: map.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-10T18:23:08+05:30
 */



//map in go is like the dictionary in python to store the key value pair, or as associative arrays is php

package main
import "fmt"

func main(){
  //let x is a map of strings to ints, here strings are the keys and ints are the values
  //below statement will create an error, as uninitialized map cannot be assigned the value
  //var x map[string]int
  //so to initialize the map use make function to create the map, and
  //it does not take the size as second argument as in case of the arrays
  var x= make(map[string]int)
  //strings can use "" as well as `` in go
  x[`key`] = 10
  fmt.Println(x)
  //delete the key value in the map
  delete(x,`key`)
  fmt.Println(`after deleting`,x)
}
