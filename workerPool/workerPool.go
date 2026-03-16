package main 

func main(){
	jobs := make(chan int, 15)
	results := make(chan int, 15)
}

func worker(id int, jobs chan<- int, results chan<- int) {
	for j := range jobs
}