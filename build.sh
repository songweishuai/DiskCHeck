#########################################################################
# File Name: build.sh
# Author: 宋伟帅
# mail: songweishuai@thudner.com.cn
# Created Time: 2018年11月01日 星期四 13时16分02秒
#########################################################################
#!/bin/bash

echo "build begin"
env GOOS=linux GOARCH=arm go build main.go
echo "build end"


