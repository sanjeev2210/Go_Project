package main

import(
"fmt"
)

func add(n int) (x,y int){
sum:=0
for i:=0; i<n; i++ {
sum=sum+i
}
x=sum-1
y=sum+1
return
}

func main(){
fmt.Println(add(10))
}
