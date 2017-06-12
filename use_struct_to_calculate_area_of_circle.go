/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-12T13:19:01+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: use_struct_to_calculate_area_of_circle.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-12T13:34:16+05:30
 */
package main
import "fmt"
import "math"
//define the struct circle
type Circle struct{
  x , y , r float64
}
//function of circle to calculate the area of the circle
//*circle understood to take the pointer of the circle
func (c *Circle) calculateArea() (area float64){
  area =  math.Pi*c.r*c.r //area = Pi*square(r)
  return
}
func main(){
  c:= Circle{x:10,y:20,r:5} //circle with x=10, y=20, r is the radius i.e. 5// or
  //we can use c:= Circle{10,20,5} as we know the order of x,y,r
  fmt.Printf("area of the circle is %f\n",c.calculateArea()) //it automatically pass the c pointer to the calculateArea function
}
