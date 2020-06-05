#!/usr/bin/env sh
# pull
if [ "$1" = "" ];then
	printf "Please input commit message\nExample: $0 commitmessage\n"
	exit 1
fi
echo -e '开始执行命令\n'

echo "执行git pull"
git pull
[ $? -ne 0 ] && exit 2

# 保存所有的修改
echo "执行命令：git add -A"
git add -A

# 把修改的文件提交
echo "执行命令：commit -m  $1"
git commit -m "$1"
[ $? -ne 0 ] && exit 3

echo "执行命令: git push"
git push