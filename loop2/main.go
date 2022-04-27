package main

import "fmt"

func main(){
	var n,max,sum int
    i:=1
	for i<10 {
		fmt.Scan(&n)
		if n==0{
		    break
		}
	    if n>max{
			max=n
			sum=0
		}
		if n==max{
			sum+=1
		}

	}
	fmt.Println(sum)

}