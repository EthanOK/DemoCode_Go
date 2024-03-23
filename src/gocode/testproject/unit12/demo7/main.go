package main

func main() {
	// 管道的遍历(必须要先关闭管道，再遍历)

	var chanA chan int

	chanA = make(chan int, 20)

	for i := 0; i < 20; i++ {
		chanA <- i
	}

	close(chanA)

	for v := range chanA {
		println(v)
	}

}
