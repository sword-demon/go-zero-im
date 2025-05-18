user-rpc-dev:
	@make -f deploy/mk/user-rpc.mk release-test  # 执行对应目录下的makefile文件的 release-test 包含的所有命令

release-test: user-rpc-dev

install-server:
	cd  ./deploy/scripts && chmod +x release-test.sh && ./release-test.sh