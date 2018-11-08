proj_name=medium
cd ../..
export GOPATH=$(pwd)
cd src/$proj_name/
export GOAPP=$proj_name
export GOENV=local
code . && bee run