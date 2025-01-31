package main

import (
	"fmt"
	"sync"
)
type post struct{
	views int
	mu sync.Mutex
}


func (p *post) inc(wg *sync.WaitGroup){
	defer func ()  {	
	p.views += 1;
	wg.Done()	
	}()

	p.mu.Lock()
	p.mu.Unlock()
}

func main(){
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i:= 0; i < 100 ; i++{
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()
	fmt.Println(myPost.views)

}