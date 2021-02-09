#!/bin/bash
ARM_VERSION=6
ARCH_RPI=arm
ARCH_DESKTOP=amd64
SOURCE_FILE='channel_req.go'
PROGRAM_NAME='channel_req'

echo "Building for Linux x64..."
env GOOS=linux GOARCH=$ARCH_DESKTOP go build -o ./builds/linux_x64/$PROGRAM_NAME_LINUX $SOURCE_FILE

echo "Building for Linux ARM (RPI)..."
env GOOS=linux GOARCH=$ARCH_RPI GOARM=$ARM_VERSION go build -o ./builds/rpi/$PROGRAM_NAME_RPI $SOURCE_FILE

echo "Building for Windows x64..."
env GOOS=windows GOARCH=$ARCH_DESKTOP go build -o ./builds/win_x64/$PROGRAM_NAME_WINDOWS $SOURCE_FILE

exit 0 