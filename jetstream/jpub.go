package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Drain()

	js, _ := nc.JetStream() //nats.PublishAsyncMaxPending(10000))
	//js.DeleteStream("TEST_STREAMM")

	subjects := make([]string, 1000000)
	for i := 0; i < 1000000; i++ {
		subjects[i] = fmt.Sprintf("t.%d", i)
	}

	// //STREAMS ALREADY CTREATED
	js.AddStream(&nats.StreamConfig{
		Name:     "TESTTTT",
		Subjects: subjects,
		MaxMsgs:  100000,
		Storage:  nats.MemoryStorage,
		Replicas: 1,
		MaxAge:   3600 * time.Second,
		//Retention: nats.WorkQueuePolicy,
	})

	s := []byte(`
		"nfInstanceId": "a52f47c0upf",
		"nfType": "UDM",
		"nfStatus": "REGISTERED",
		"heartBeatTimer": 20,
		"sNssais": [
			{
				"sst": 1,
				"sd": "0000d1"
			},
			{
				"sst": 1,
				"sd": "0000d3"
			},
			{
				"sst": 1,
				"sd": "0000d8"
			},
			{
				"sst": 1,
				"sd": "0000d5"
			}
		],
		"perPlmnSnssaiList": [
			{
				"plmnId": {
					"mcc": "311",
					"mnc": "480"
				},
				"sNssaiList": [
					{
						"sst": 1,
						"sd": "0000d1"
					},
					{
						"sst": 1,
						"sd": "0000d3"
					},
					{
						"sst": 1,
						"sd": "0000d8"
					},
					{
						"sst": 1,
						"sd": "0000d5"
					}
				]
			}
		],
		"fqdn": "http://10.30.249.1:9002",
		"load": 0,
		"upfInfo": {
			"sNssaiUpfInfoList": [
				{
					"sNssai": {
						"sst": 1,
						"sd": "0000d1"
					},
					"dnnUpfInfoList": [
						{
							"dnn": "vzwinternet",
							"dnaiList": [
								"dnai1",
								"dnai2"
							]
						},
						{
							"dnn": "vzwims",
							"dnaiList": [
								"dnai1",
								"dnai2"
							]
						}
					]
				},
				{
					"sNssai": {
						"sst": 1,
						"sd": "0000d5"
					},
					"dnnUpfInfoList": [
						{
							"dnn": "vzwinternet",
							"dnaiList": [
								"dnai1",
								"dnai2"
							]
						},
						{
							"dnn": "vzwims",
							"dnaiList": [
								"dnai1",
								"dnai2"
							]
						}
					]
				},
				{
					"sNssai": {
						"sst": 1,
						"sd": "0000d3"
					},
					"dnnUpfInfoList": [
						{
							"dnn": "vzwinternet",
							"dnaiList": [
								"dnai1",
								"dnai2"
							]
						},
						{
							"dnn": "vzwims",
							"dnaiList": [
								"dnai1",
								"dnai2"
							]
						}
					]
				},
				{
					"sNssai": {
						"sst": 1,
						"sd": "0000d8"
					},
					"dnnUpfInfoList": [
						{
							"dnn": "vzwinternet",
							"dnaiList": [
								"dnai1",
								"dnai2"
							]
						}
					]
				}
			],
			"interfaceUpfInfoList": [
				{
					"interfaceType": "N3",
					"ipv4EndpointAddresses": [
						"10.130.253.100"
					]
				},
				{
					"interfaceType": "N6",
					"ipv4EndpointAddresses": [
						"10.130.254.101"
					]
				}
			],
			"iwkEpsInd": false,
			"pduSessionTypes": [
				"IPV4"
			],
			"taiList": [
				{
					"plmnId": {
						"mcc": "311",
						"mnc": "480"
					},
					"tac": "000001"
				}
			]
		},
		"nfServicePersistence": false,
		"nfProfileChangesInd": false
	`)
	subs := "t"
	start := time.Now()

	msg := &nats.Msg{
		Subject: subs,
		Data:    s,
		Header: nats.Header{
			"Head": []string{"Hey from header"},
		},
	}

	for i := 0; i < 100; i++ {
		js.PublishAsync(subs, msg.Data)

		js.PublishAsyncComplete()

		// if i%20000 == 0 {
		// 	<-js.PublishAsyncComplete()

		// }
	//	time.Sleep(2 * time.Second)

	}
	js.PublishAsyncComplete()
	// select {
	// case <-js.PublishAsyncComplete():
	// 	fmt.Println("published 1 messages")
	// case <-time.After(time.Second):
	// 	fmt.Println("publish took too long")
	// }
	nc.Flush()

	// nc.Subscribe("tracking.processed", func(msg *nats.Msg) {
	// 	fmt.Println("Tracking:", string(msg.Data))
	// })
	fmt.Println("Jpub1 time taken  :", time.Since(start))

	// select {}
}
