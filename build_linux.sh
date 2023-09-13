#!/usr/bin/env bash

mkdir ./out
cp LICENSE ./out/LICENSE

build(){
    wails build -clean -platform $1/$2
    mv ./build/bin/Colossus ./out/Colossus
    cd out
    tar -zcf colossus-$1-$2.tar.gz Colossus LICENSE
    rm -rf ./Colossus
    cd ../
}

echo "[1] Linux from amd64"
build linux amd64
echo "[2] Linux from arm64"
build linux arm64