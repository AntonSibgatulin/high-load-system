package main

import "fmt"

func main(){
	for range 100 {
		count := 0
		done := make(chan struct{})
		govno := make(chan struct{}, 1)
		govno <- struct{}{}
		go func (){
			for range 100000 {
				count++
			}
			done <- struct{}{}
			<- govno
		}()
		go func(){
			for range 100000 {
				count++
			}
			<- govno
			done <- struct{}{}
		}()
		<- done
		<- done
		fmt.Println(count)
		if count >= 200000 {
		}
	}
}