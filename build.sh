#!/bin/bash
###############################################################################
#编译脚本的原理是将编译结果放到output目录中，这个样例模版提供一个产生
#一个最基本golang运行程序包的编译脚本，对于特殊的需求请酌情考虑
#
#1、该脚本支持参数化，参数将传入build_package函数（内容为最终执行的编译命令）
#   ，用$1,$2....表示，第1,2...个参数
#2、部署需要启动程序，所以需要提供control文件放在当前目录中，用于启动和
#   监控程序状态

###############用户修改部分################
readonly PACKAGE_DIR_NAME="blog_service"
readonly OUTPUT=$(pwd)/output

#建立最终发布的目录
function build_dir
{
    mkdir -p ${OUTPUT}/bin || return 1
    mkdir -p ${OUTPUT}/conf || return 1
    mkdir -p ${OUTPUT}/logs || return 1
}

#清理编译构建目录操作
function clean_before_build
{
    rm -rf ${OUTPUT}
}

function build_package() 
{
  export GO111MODULE="auto" 
  go mod tidy
  go generate
  #修改部分
  go build -ldflags  "-X Practice/go-programming-tour-book/blog-service/cmd.BuildTime=`date +%Y-%m-%d,%H:%M:%S` -X Practice/go-programming-tour-book/blog-service/cmd.BuildVersion=1.0.0 -X Practice/go-programming-tour-book/blog-service/cmd.GitCommitID=`git rev-parse HEAD`"  -o ${OUTPUT}/bin/blog-service .
}

function pre_copy()
{
  cp -r ./configs/* ${OUTPUT}/conf || return 1
  cp -r ./bin/* ${OUTPUT}/bin || return 1
  chmod +x ${OUTPUT}/bin/* ||return 1
}

function main() {
  cd $(dirname $0)
  clean_before_build || exit 1
  build_dir || exit 1
  pre_copy || exit 1
  build_package || exit 1
}

main $@
