package main
import(
	"fmt"
	
)

type Student struct {
	Name string
	age int

}
func main(){
	stu := Student{
		Name:"lisi",
		age:12,
	}
	fmt.Println("buoda",stu)
}
