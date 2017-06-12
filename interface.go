/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-12T15:33:51+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: interface.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-12T16:11:02+05:30
 */


//for refering to interfaces go to https://www.golang-book.com/books/intro/9
package main
import "fmt"
import "math"

//interface
type shape interface{
  area() float64
}
//since the struct circle and rectangle have function area they are implementing the interface shape
type Circle struct{
  x,y,r float64
}

func (c *Circle) area() (area float64){
  area = math.Pi*c.r*c.r
  return
}

//type rectangle
type Rectangle struct{
  length float64
  width float64
}

func (r *Rectangle) area() (area float64){
  area = r.length*r.width
  return
}

type MultiShape struct{
  shapes []shape // can contain multiple shape implemented types
}

func (m *MultiShape) totalArea() (area float64){
 for _,shape:= range m.shapes{
   area+=shape.area()  //every shape call its own function area
 }
 return //return the total area
}

func main(){
  c:=Circle{10,20,5}
  r:=Rectangle{20,40}
  multishape:=MultiShape{[]shape{&c,&r}}
  fmt.Printf("totalarea is %f\n",multishape.totalArea())
}
