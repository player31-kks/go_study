// 옵져버 패턴과 거의 유사하다.

// subject -> notify()를 보내면
// Observer(Listeners)는 subject가 notify를 발생시킨다.

//옵져버 패턴은 동기식이여서 문제가 생김 따라서 발행-구독 패턴 (pub/sub)패턴

// pub/sub
// publisher 가 channel에게 데이터를 발생시키면
// channel에 등록하고 있는 message broker인 topic들이 데이터를 받아
// 각 subscriber에게 전달하게 값들을 전달하게 된다.

// chan (chan <- string) // 일방향 채널 이 채널은 넣기만 가능한 채널이다.

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	publisher := NewPublisher(ctx)
	subscriber1 := NewSubscriber("교석", ctx)
	subscriber2 := NewSubscriber("kks", ctx)

	go publisher.Update()

	subscriber1.Subscribe(publisher)
	subscriber2.Subscribe(publisher)

	go subscriber1.Update()
	go subscriber2.Update()

	go func() {
		tick := time.Tick(time.Second * 2)
		for {
			select {
			case <-tick:
				publisher.Publish("Hello world")
			case <-ctx.Done():
				fmt.Println("main 함수 끝")
				return
			}
		}
	}()
	fmt.Scanln()
	cancel()
}
