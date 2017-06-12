/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-12T13:37:06+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: use_embeeded_struct_type.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-12T15:15:13+05:30
 */



package main
import "fmt"
//embedded type means a struct contains a type i.e. is also a struct i.e. relation to another struct
//type Person
type Person struct{
  name string
}
//Person function
func (p *Person) talk(mobile string){
  fmt.Printf("The talked person is %s with %s\n ",p.name,mobile)
}

//struct Mobile i.e. embeddedtype containg relation to struct Person
type (Mobile struct{
  Person //anonomous field i.e. unnamed field
  mobile_model string
}
)
func main(){
  var m Mobile
  m.talk("")
  //initialize with the name
  m1:=Mobile{Person{name:"Tushar"},
    "Oneplus 3T",
  }
  m1.talk(m1.mobile_model)
}
