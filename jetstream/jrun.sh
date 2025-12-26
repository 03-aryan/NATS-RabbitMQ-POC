#!/bin/bash

echo "Starting jpub1, jpub2 in parallel..."

go run jpub.go &
go run jpub2.go &
#go run rpub3.go &

wait
echo "All publishers finished!"

# to run
# chmod +x jrun.sh
#./jrun.sh