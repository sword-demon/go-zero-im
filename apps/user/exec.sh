# 第一次执行脚本
goctl rpc protoc ./apps/user/rpc/user.proto --go_out=./apps/user/rpc/ --go-grpc_out=./apps/user/rpc/ --zrpc_out=./apps/user/rpc/

# 根据指定的user.sql文件生成模型
goctl model mysql ddl -src="./deploy/sql/user.sql" -dir="./apps/user/models/" -c

# 生成user服务
goctl api go -api ./apps/user/api/user.api -dir ./apps/user/api -style=gozero