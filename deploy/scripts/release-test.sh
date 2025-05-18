# 总脚本

need_start_server_shell=(
  # user prc 启动脚本
  user-rpc-test.sh
)

for i in ${need_start_server_shell[*]}; do
  chmod +x $i
  ./$i
done

docker ps
docker exec -it chat-etcd etcdctl get --prefix ""