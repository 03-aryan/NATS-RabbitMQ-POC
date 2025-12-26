package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS
	//nc, err := nats.Connect(nats.DefaultURL)
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println("Error connecting to NATS:", err)
		return
	}
	defer nc.Close()

	start := time.Now()

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

	msg := nats.NewMsg("test.header")
	//msg = nats.NewMsg(s)
	msg.Data = s

	// Publisher sending timestamped messages
	for i := 0; i <= 1000000; i++ {

		msg.Header.Set("1", "client 1  ")

		// Publish message
		err := nc.Publish(msg.Subject, msg.Data)
		if err != nil {
			fmt.Println("Error publishing message:", err)
		}
	}

	nc.Flush() // Ensure all messages are flushed to NATS

	dur := time.Since(start)
	fmt.Println("Time taken Pub3:", dur)
}
