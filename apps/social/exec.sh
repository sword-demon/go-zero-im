goctl rpc protoc ./apps/social/rpc/social.proto --go_out=./apps/social/rpc/ --go-grpc_out=./apps/social/rpc/ --zrpc_out=./apps/social/rpc/

# 根据指定的social.sql文件生成模型
goctl model mysql ddl -src="./deploy/sql/social.sql" -dir="./apps/social/socialmodels/" -c