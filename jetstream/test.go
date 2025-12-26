// package main

// import (
// 	"fmt"
// 	"sync"
// 	"sync/atomic"
// 	"time"

// 	"github.com/nats-io/nats.go"
// )

// func main() {
// 	const total = 100_000
// 	data := []byte("abc")

// 	nc, err := nats.Connect(nats.DefaultURL)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer nc.Close()

// 	js, err := nc.JetStream(nats.PublishAsyncMaxPending(total))
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Stream setup (safe if already exists)
// 	_, err = js.AddStream(&nats.StreamConfig{
// 		Name:     "fast_stream",
// 		Subjects: []string{"t"},
// 		Storage:  nats.MemoryStorage,
// 		Replicas: 1,
// 	})
// 	if err != nil && err != nats.ErrStreamNameAlreadyInUse {
// 		panic(err)
// 	}

// 	var acked int32
// 	var failed int32
// 	var wg sync.WaitGroup

// 	// Track publish errors
// 	errCh := js.PublishAsyncErrChan()
// 	go func() {
// 		for err := range errCh {
// 			fmt.Println("❌ Publish error:", err)
// 			atomic.AddInt32(&failed, 1)
// 			wg.Done()
// 		}
// 	}()

// 	start := time.Now()

// 	// Publish messages
// 	for i := 0; i < total; i++ {
// 		wg.Add(1)
// 		ackFuture, err := js.PublishAsync("t", data)
// 		if err != nil {
// 			fmt.Println("❌ Failed to publish:", err)
// 			atomic.AddInt32(&failed, 1)
// 			wg.Done()
// 			continue
// 		}

// 		// Wait for each individual ack
// 		go func(af nats.PubAckFuture) {
// 			_, err := af.Ok()
// 			if err != nil {
// 				atomic.AddInt32(&failed, 1)
// 			} else {
// 				atomic.AddInt32(&acked, 1)
// 			}
// 			wg.Done()
// 		}(ackFuture)
// 	}

// 	// Wait for all publish acks
// 	wg.Wait()

// 	nc.Flush()
// 	nc.Drain()

// 	elapsed := time.Since(start)

// 	fmt.Printf("✅ Published: %d\n", total)
// 	fmt.Printf("✅ Acked:     %d\n", acked)
// 	fmt.Printf("❌ Failed:    %d\n", failed)
// 	fmt.Printf("⏱️  Time:     %v\n", elapsed)
// }
