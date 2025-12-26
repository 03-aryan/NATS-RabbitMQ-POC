package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	// Connect to RabbitMQ
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()

	// Create a channel
	ch, _ := conn.Channel()
	defer ch.Close()

	// Declare a queue
	q, _ := ch.QueueDeclare(
		"taskqueue", // Name of the queue
		true,        // Durable (survive server restarts)
		false,       // Delete when unused
		false,       // Exclusive
		false,       // No-wait
		nil,         // Arguments
	)

	body := []byte(`
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

	for i := 0; i <= 1000000; i++ {
		ch.Publish(
			"",     // Exchange (default)
			q.Name, // Routing key (queue name)
			false,  // Mandatory
			false,  // Immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        []byte(body),
			})
	}

	// Sleep a bit to ensure all messages are sent
	//time.Sleep(5 * time.Second)

	// Print the time taken
	fmt.Println("Time taken Pub2:", time.Since(start))
}
