//package main
package id1

import(
	"fmt"
)


//func main(){
func Solution(){

	fmt.Println("Solution to the problem id1")

	const limit=1000

	var sum int

	sum=0

	for i:=1; i < limit; i++ {
		if( (i%3==0) || (i%5 == 0)){
			sum+=i
		}
	}
	
	fmt.Println("Total sum of multiples of 3 or 5 is",sum)
}
