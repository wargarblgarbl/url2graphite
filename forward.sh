#!/bin/sh
PORT=2003
SERVER=192.168.1.138
while read x; do
        stat=$(echo $x | awk 'NR==1,/\//{sub(/\//, "")}1' | sed 's/\(.*\)\//\1 /' | sed 's/\//\./g' )
        echo "$stat `date +%s`" | nc -q0 ${SERVER} ${PORT}
done



