#!/usr/bin/env bash
DATE=`date '+%Y%m%d%H%M'`
os="linux"
arch="amd64"
export GOPATH="/Users/david/project_code/go_project"

echo "XYAccountServer build start"
if [ ! -d "./build/account" ]; then 
	 mkdir -p ./build/account
fi
cp -R ./XYAccoutServer/conf  ./build/account/conf
GOOS=$os GOARCH=$arch go build -a -x -o ./build/account/XYAccountServer$DATE ./XYAccoutServer/main.go
echo "XYAccountServer build end"

echo "######################"
echo "XYActivityServer build start"
if [ ! -d "./build/activity" ]; then 
	mkdir  -p ./build/activity
fi
cp -R ./XYActivityServer/conf  ./build/activity/conf
GOOS=$os GOARCH=$arch go build -a -x -o ./build/activity/XYActivityServer$DATE ./XYActivityServer/main.go
echo "XYActivityServer build end"

echo "######################"
echo "XYChatServer build start"
if [ ! -d "./build/chat" ]; then 
	mkdir -p ./build/chat
fi
cp -R ./XYChatServer/conf  ./build/chat/conf
GOOS=$os GOARCH=$arch go build -a -x -o ./build/chat/XYChatServer$DATE ./XYChatServer/main.go
echo "XYChatServer build end"

echo "######################"
echo "XYDynamicServer build start"
if [ ! -d "./build/dynamic" ]; then 
	mkdir  -p ./build/dynamic
fi
cp -R ./XYDynamicServer/conf  ./build/dynamic/conf
GOOS=$os GOARCH=$arch go build -a -x -o ./build/dynamic/XYDynamicServer$DATE ./XYDynamicServer/main.go
echo "XYDynamicServer build end"

echo "######################"
echo "XYFileServer build start"
if [ ! -d "./build/file" ]; then 
	mkdir  -p ./build/file
fi
cp -R ./XYFileServer/conf  ./build/file/conf
GOOS=$os GOARCH=$arch go build -a -x -o ./build/file/XYFileServer$DATE ./XYFileServer/main.go
echo "XYFileServer build end"

echo "######################"
echo "XYPushServer build start"
if [ ! -d "./build/push" ]; then 
	mkdir  -p ./build/push
fi
cp -R ./XYPushServer/conf  ./build/push/conf
GOOS=$os GOARCH=$arch go build -a -x -o ./build/push/XYPushServer$DATE ./XYPushServer/main.go
echo "XYPushServer build end"

echo "######################"
echo "XYRobot build start"
if [ ! -d "./build/robot" ];then 
	mkdir  -p ./build/robot
fi
cp -R ./XYRobot/conf  ./build/robot/conf
GOOS=$os GOARCH=$arch go build -a -x -o ./build/robot/XYRobot$DATE ./XYRobot/main.go
echo "XYRobot build end"

echo "######################"
echo "XYShareServer build start"
if [ ! -d "./build/share" ];then 
	mkdir  -p ./build/share
fi
cp -R ./XYShareServer/conf  ./build/share/conf
GOOS=$os GOARCH=$arch go build -a -x -o ./build/share/XYShareServer$DATE ./XYShareServer/main.go
echo "XYShareServer build end"



echo "######################"
echo "XYRPCServer build start"
if [ ! -d "./build/rpc" ];then 
	mkdir  -p ./build/rpc
fi
cp -R ./XYRPCServer/conf  ./build/rpc/conf
GOOS=$os GOARCH=$arch go build -a -x -o ./build/share/XYShareServer$DATE ./XYShareServer/main.go
echo "XYRPCServer build end"
echo "all build done"
exit 