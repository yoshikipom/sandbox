# How to test it
run container
docker run -itd --name gor -p 8000:8000 -p 8001:8001 --entrypoint '' goreplay:1.3.3 /bin/bash

run server 1
$ docker exec  gor /gor file-server :8000

run server 2
$  docker exec  gor /gor file-server :8001 

run gorepaly
$ docker exec gor /gor --input-raw :8000 --output-http="http://localhost:8001" 
