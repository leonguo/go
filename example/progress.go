package main

func main() {
	// 处理多个任务
	//tasks := getTask()
	//for _, task := range tasks {
	//	process(task)
	//}

	ch := make(chan Task, 3)

	for i := 0; i < numWorkers; i++ {
		go worker(ch)
	}

	tasks := getTasks()

	for _ , task := range  tasks {
		ch <- task
	}
}

func worker(ch chan Task){
	for {
		task := <-ch
		process(task)
	}
}