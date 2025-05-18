#!/bin/bash

repo_addr='crpi-axo2tzdpreqzdzqq.cn-hangzhou.personal.cr.aliyuncs.com/wx-easy-chat/user-rpc-dev'
tag='latest'

container_name="easy-chat-user-rpc-test"

docker stop ${container_name}

docker rm ${container_name}

docker rmi ${repo_addr}:${tag}

docker pull ${repo_addr}:${tag}

docker run -p 10000:10000 --name=${container_name} -d ${repo_addr}:${tag}