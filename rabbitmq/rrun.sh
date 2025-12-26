#!/bin/bash

echo "Starting pubid, pubid2 in parallel..."

go run rabbitpub.go &
go run rpub2.go &
go run rpub3.go &

wait
echo "All publishers finished!"

# to run
# chmod +x run.sh
#./run.sh