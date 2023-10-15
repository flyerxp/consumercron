package service

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/flyerxp/lib/logger"
	"github.com/flyerxp/lib/middleware/pulsarL"
	"strconv"
)

func Test(cnt int) {
	ctx := logger.GetContext(context.Background(), "cmr")
	//return
	for i := 0; i <= 100; i++ {
		pulsarL.Producer(ctx, &pulsarL.OutMessage{
			Topic:      10101001,
			Content:    map[string]string{"a": "b", "10101001": "==============" + strconv.Itoa(i) + "=================="},
			Properties: map[string]string{"prop": "prop"},
			Delay:      0,
		})
	}
	topics := pulsarL.GetStringTopics([]int{10101001})
	topics = append(topics, "test")
	c := pulsarL.NewConsumer(ctx, topics, "testConsume", pulsarL.WithDlq(pulsar.DLQPolicy{
		MaxDeliveries:    50,
		DeadLetterTopic:  "dead_test",
		RetryLetterTopic: "retry_test",
	}), pulsarL.WithMaxConsumerCnt(cnt))

	//count := 0
	c.Consumer(func(ctx context.Context, message *pulsar.ConsumerMessage, message2 *pulsarL.PulsarMessage) bool {
		//c.Stop()
		fmt.Println(message2.String())
		fmt.Println(message.Properties())
		fmt.Println(message2.Topic)

		/*if count == 9 {
			pulsarL.RetryAfter(message, time.Second*10, map[string]string{"aaa": "abdddddddddddddddddddddddddddddddddddddddddcd"})
			return true
		}*/
		return true
	})
}
