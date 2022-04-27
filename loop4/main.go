// package main
// import "fmt"
// func main(){
// 	var a int
	
// 	for fmt.Scan(&a);a<100;fmt.Scan(&a){
// 		if a<10{
// 			continue
// 		}else{
// 			fmt.Println(a)
// 		} 

// 	}
// }


// Bank

// package main
// import "fmt"
// func main(){
// 	var a,b,c float32
// 	fmt.Scan(&a,&b,&c)
// 	sum:=0
// 	for a<=c{
// 		a=a+a*b/100
// 		sum+=1
// 	}
// 	fmt.Println(sum)
// }

package main
import "fmt"
func main(){
	var  workArray [16]int
	var b int
	
	for idx:= range workArray {
		fmt.Scan(&b)
		workArray[idx]=b

	}
	x:=workArray[workArray[10]]
	workArray[workArray[10]]=workArray[workArray[11]]
	workArray[workArray[11]]=x
	y:=workArray[workArray[12]]
	workArray[workArray[12]]=workArray[workArray[13]]
	workArray[workArray[13]]=y
	z:=workArray[workArray[14]]
	workArray[workArray[14]]=workArray[workArray[15]]
	workArray[workArray[15]]=z
	// fmt.Println(workArray)
	for i:=0;i<10;i++{
		fmt.Print(workArray[i]," ")
		if i==9{
			fmt.Println("OK")
		}
	}
}