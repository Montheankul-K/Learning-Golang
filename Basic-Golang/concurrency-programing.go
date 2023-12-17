package main

import (
	"fmt"
	"sync"
)

// Worker pool
// Jobs (channel) : รับ input
type numbers struct {
	a int
	b int
}

// Results (channel): output
// jobs กับ result จะสื่อสารกันตลอดจนทำงานเสร็จ
type sums struct {
	r int
}

func workers(jobsCh <-chan numbers, resultCh chan<- sums) {
	for job := range jobsCh {
		resultCh <- sums{r: summation(job.a, job.b)}
	}
}

func summation(a, b int) int {
	return a + b
}

func main() {
	// Concurrency programing : การทำงานแบบ asynchronous
	// มี 3 แบบ Basic (sync.WaitGroup), Channel, Worker Pool
	a := []int{1, 2, 3, 4, 5}

	// Basic
	var wg sync.WaitGroup
	for i := range a {
		wg.Add(1) // มี go routine 1 ตัวในแต่ละ loop
		// Anonymous function : function ที่ทำงานแล้วจบไป
		go func(i int) {
			// go routine จะบังคับไม่ให้มี return
			defer wg.Done()
			fmt.Printf("%v ", a[i])
		}(i) // input
		// เป็นการแยก function มาทำงาน parallel ตามจำนวน item ใน slice
	}
	wg.Wait() // Done ต้องครบตามจำนวน go routine

	// จะไม่ได้แยก go routine เป็นแต่ละ item ใน slice เหมือนด้านบน
	wg.Add(len(a)) // กำหนดว่ามี go routine กี่ตัว
	go func() {
		for i := range a {
			defer wg.Done()
			fmt.Printf("%v ", a[i])
		}
	}()
	wg.Wait()

	// Channel : name := make(chan type, buffer size)
	numberCh := make(chan int) // ถ้าไม่กำหนด buffer size จะเป็น 1 channel
	msgCh := make(chan string)

	// Number
	go func(numberCh chan<- int) { // sender
		// pass params แบบ two-way รับค่า หรือ assign ค่าให้ตัวอื่นก็ได้
		// numberCh <-chan int : assign ค่าอย่างเดียว
		// numberCh chan-> int : รับค่าอย่างเดียว
		numberCh <- 10
		// close()
		// close() : close channel ถ้าไม่ใส่ จะ close เมื่อ buffer size เต็ม
	}(numberCh)

	// Message
	go func(msgCh chan<- string) {
		msgCh <- "Hello World"
	}(msgCh)

	number := <-numberCh // reseiver
	msg := <-msgCh       // assign ค่าจาก channel msg จะรอจนกว่า channel จะ close

	fmt.Println(number)
	fmt.Println(msg)

	// Worker Pool : ใช้บ่อย มีการ limit จำนวนการ parallel
	// ประกอบด้วย jobs, result
	nums := []numbers{
		{a: 1, b: 2},
		{a: 2, b: 3},
		{a: 3, b: 4},
		{a: 4, b: 5},
		{a: 5, b: 6},
		{a: 1, b: 2},
		{a: 2, b: 3},
		{a: 3, b: 4},
		{a: 4, b: 5},
		{a: 5, b: 6},
	}
	jobsCh := make(chan numbers, len(nums))
	resultsCh := make(chan sums, len(nums))

	// Assign input to channel
	for _, n := range nums {
		jobsCh <- n
	}
	close(jobsCh)

	numWorkers := 2 // ทำงาน parallel ได้มากสุด 2 ครั้งต่อรอบ
	for w := 0; w < numWorkers; w++ {
		// Worker function (go routine): ให้ jobs ทำงานและรับ output ลง result
		go workers(jobsCh, resultsCh)
	}
	// ทำจนกว่า resultCh จะเต็ม
	results := make([]sums, 0)
	for a := 0; a < len(nums); a++ {
		temp := <-resultsCh
		results = append(results, temp)
	}
	fmt.Println(results)
}
