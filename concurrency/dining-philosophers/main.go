package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. 일정 시간 생각을 한다.
// 2. 왼쪽 포크가 사용 가능해질 때까지 대기한다. 만약 사용 가능하다면 집어든다.
// 3. 오른쪽 포크가 사용 가능해질 때까지 대기한다. 만약 사용 가능하다면 집어든다.
// 4. 양쪽의 포크를 잡으면 일정 시간만큼 식사를 한다.
// 5. 오른쪽 포크를 내려놓는다.
// 6. 왼쪽 포크를 내려놓는다.
// 7. 다시 1번으로 돌아간다.

//간단하게 생각해, 만약 모든 철학자들이 동시에 자신의 왼쪽 포크를 잡는다면,
// 모든 철학자들이 자기 오른쪽의 포크가 사용 가능해질 때까지 기다려야 한다.
//그런데 모든 철학자들이 그러고 있다. 이 상태에서는 모든 철학자가 영원히 3번 상태에 머물러있어 아무것도 진행할 수가 없게 되는데, 이것이 교착(Deadlock)상태이다.

const hunger = 3

// variables - philosopher
var philosophers = []string{
	"kks",
	"kky",
	"kka",
	"kkb",
	"kkc",
}
var wg sync.WaitGroup
var sleepTime = 1 * time.Second
var eatTime = 3 * time.Second
var thinkTime = 1 * time.Second
var order = make(chan string, len(philosophers))

func diningProblem(philosopher string, leftFork, rightFork *sync.Mutex) {
	defer wg.Done()
	//print a message
	fmt.Println(philosopher, "is seated.")
	time.Sleep(sleepTime)

	// lock both forks
	for i := hunger; i > 0; i-- {
		fmt.Println(philosopher, "is hungry.")
		time.Sleep(sleepTime)

		leftFork.Lock()
		fmt.Printf("\t%s picked up the fork to his left\n", philosopher)

		rightFork.Lock()
		fmt.Printf("\t%s picked up the fork to his right\n", philosopher)

		// print a message
		fmt.Println(philosopher, "has both forks, and is eating.")
		time.Sleep(eatTime)

		//give the philosopher some time to think
		fmt.Println(philosopher, "is thinking")
		time.Sleep(thinkTime)

		// unlock the mutexes
		rightFork.Unlock()
		fmt.Printf("\t%s put down the fork on this right.\n", philosopher)
		leftFork.Unlock()
		fmt.Printf("\t%s put down the fork on this left.\n", philosopher)
		time.Sleep(sleepTime)
	}

	fmt.Println(philosopher, "is satisfied")
	time.Sleep(sleepTime)
	fmt.Println(philosopher, "left table")
	order <- philosopher
}

func main() {

	// print intro
	fmt.Println("The Dining Philosophers Problem")
	fmt.Println("-------------------------------")
	wg.Add(len(philosophers))
	forkLeft := &sync.Mutex{}

	// spawn one goroutine for each philosopher
	for i := 0; i < len(philosophers); i++ {
		// create a mutex for the right fork
		forkRight := &sync.Mutex{}
		// call a goroutine
		go diningProblem(philosophers[i], forkLeft, forkRight)
		forkLeft = forkRight
	}
	wg.Wait()
	fmt.Println("The table is empty")
	for i := 0; i < len(philosophers); i++ {
		fmt.Println(<-order)
	}
}
