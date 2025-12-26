# NATS-RabbitMQ-POC
Proof of Concept for NATS with/without JStream and RabbitMQ. This was used on Ubuntu os. 

## To run install Docker images of NATS and RabbitMQ or setup them locally.

## Verifying Installation
NATS (with JetStream)
1. Verify NATS CLI installation
nats --version

Expected output:
nats version x.y.z

2. Verify NATS server installation
nats-server --version

3. Check if NATS server is running
nats ping

Expected output:

PONG

4. JetStream availability

JetStream is built into modern NATS servers but must be enabled at runtime.

Check JetStream status:
nats server info

Look for:

JetStream: enabled

Enable JetStream (if not enabled):
nats-server -js

Verify JetStream is working:
nats jetstream info

RabbitMQ
1. Verify RabbitMQ installation
rabbitmqctl status

If RabbitMQ is installed and running, this command will return detailed node status information.

2. Check RabbitMQ version
rabbitmqctl version

3. Verify RabbitMQ service is running (Windows)
sc query RabbitMQ

## RUN 
To run them simultaneously, use the provided shell scripts.
