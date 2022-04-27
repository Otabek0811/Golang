package main
import "fmt"
func main(){
	var a,b,i int
	fmt.Scan(&a)
	sum:=0
	for i=0;i<a;i++{
		fmt.Scan(&b)
		if b<100 && b>9 && b%8==0{
			sum+=b
		}

	}
	fmt.Println(sum)
}