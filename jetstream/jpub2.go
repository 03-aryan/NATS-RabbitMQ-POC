package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Drain()

	js, _ := nc.JetStream(nats.PublishAsyncMaxPending(100)) //nats.PublishAsyncMaxPending(10000))
	//js.DeleteStream("TEST_STREAMM")
	// js.AddStream(&nats.StreamConfig{
	// 	Name:     "TEST",
	// 	Subjects: []string{"t"},
	// 	MaxMsgs:  10000,
	// 	Storage:  nats.MemoryStorage,
	// 	Replicas: 1,
	// })

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

	start := time.Now()
	msg := &nats.Msg{
		Subject: "t",
		Data:    s,
		Header: nats.Header{
			"Head": []string{"Hey from header2"},
		},
	}
	for i := 0; i < 5; i++ {

		js.PublishAsync("t", msg.Data)

		nc.Subscribe("tracking.processed", func(msg *nats.Msg) {
			fmt.Println("Tracking:", string(msg.Data))
		})
		js.PublishAsyncComplete()

	}

	js.PublishAsyncComplete()
	// select {
	// case <-js.PublishAsyncComplete():
	// 	//fmt.Println("published 1 messages")
	// case <-time.After(time.Second):
	// 	fmt.Println("publish took too long")
	// }
	nc.Flush()

	fmt.Println("Jpub2 time taken  :", time.Since(start))
	select {}
}

/*
Publish vs PublishAsync
In NATS JetStream, PublishAsync is designed to be non-blocking, allowing the application to continue executing without waiting for the acknowledgment from the server. However, PublishAsync can potentially block if the number of outstanding asynchronous publishes exceeds the buffer limit, as seen in the issue discussion on GitHub.

The primary drawback of PublishAsync is that it does not wait for an acknowledgment from the server, which means that if a message fails to be published, the application may not immediately know about it. This can lead to situations where the application believes a message was published successfully when it actually wasn't, and the application might need to retry the publication.

Another drawback is that PublishAsync does not provide immediate feedback on the success or failure of the message publication, which can make it harder to handle errors in real-time.

To mitigate these issues, applications can choose to publish again or abandon the message if PublishAsync is blocked, as suggested in the GitHub issue.

In summary, while PublishAsync is generally faster and more non-blocking than Publish, it requires careful handling of potential failures and may block if the buffer limit is exceeded.

*/
