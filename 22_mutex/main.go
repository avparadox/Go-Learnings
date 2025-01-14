package main

import "fmt"

type post struct{
	views int
}


func (p *post) inc(){
	p.views += 1;
}

func main(){

	myPost := post{views: 0}

	for i:= 0; i < 100 ; i++{
		myPost.inc()
	}

	fmt.Println(myPost.views)

}