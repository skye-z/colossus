#!/usr/bin/env bash

mkdir ./out
cp LICENSE ./out/LICENSE

build_mac(){
    wails build -clean -platform darwin/$1
    mv ./build/bin/Colossus.app ./out/Colossus.app
    cd out
    tar -zcf colossus-darwin-$1.tar.gz Colossus.app LICENSE
    rm -rf ./Colossus
    cd ../
}

build_win(){
    wails build -clean -platform windows/$1
    mv ./build/bin/Colossus.exe ./out/colossus.exe
    cd out
    zip -q colossus-windows-$1.zip colossus.exe LICENSE
    rm -rf ./colossus.exe
    cd ../
}

echo "[1] MacOS from amd64"
build_mac amd64
echo "[2] MacOS from arm64"
build_mac  arm64
echo "[3] Windows from amd64"
build_win amd64
echo "[4] Windows from arm64"
build_win arm64