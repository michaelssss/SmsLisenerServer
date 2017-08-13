package hello2

import "time"

type Message struct {
	ReciviedTime    time.Time
	ReciviedContent string
	From            string
}
