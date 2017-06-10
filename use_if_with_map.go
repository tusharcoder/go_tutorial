/**
 * @Author: Tushar Agarwal(tusharcoder) <tushar>
 * @Date:   2017-06-10T18:25:55+05:30
 * @Email:  tamyworld@gmail.com
 * @Filename: use_if_with_map.go
 * @Last modified by:   tushar
 * @Last modified time: 2017-06-10T23:32:32+05:30
 */



package main
import "fmt"

func main(){
  //declare and initialize map
  names:=map[string]map[string]string{
      "a" :map[string]string{
        "name":"Rohit",
        "course":"BTECH",
    },
      "b":map[string]string{
        "name":"Abdul",
        "course":"MTECH",
    },
      "c":map[string]string{
        "name":"Rehman",
        "course":"MCA",
    },
      "d":map[string]string{
        "name":"Aniket",
        "course":"MSc",
    },
      "e":map[string]string{
        "name":"AZhar",
        "course":"HHonours",
    },
  }
  //check if name in names then print name and course else print not in names, check the value of the ok in if statement
  for _,value :=range [2]string{"e","f"}{
    if data,ok:=names[value];ok{
      fmt.Printf("name %s, course %s\n",data["name"],data["course"])
    }else{
      fmt.Printf("key %s not in names\n",value)
    }
  }
}
