#!/bin/bash

echo "Starting pubid, pubid2 in parallel..."

go run pub.go &
go run pub2.go &
go run pub3.go &

wait
echo "All publishers finished!"

# to run
# chmod +x run.sh
#./run.sh