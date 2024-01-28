package main

import "fmt"

func incrementAllByx(arr *[]int,x int){
for i := 0; i < len(*arr); i++ {
	(*arr)[i] = (*arr)[i] + x
	fmt.Println(&(*arr)[i])
}
}

func main(){
	arr := []int{1,2,3,4}
	incrementAllByx(&arr, 1)
	fmt.Println(arr)

}