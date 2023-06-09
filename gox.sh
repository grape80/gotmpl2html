#!/bin/sh

. .build/env

app=$APP_NAME
version=$VERSION
mainDir=$MAIN_DIR
distDir=$DIST_DIR

distlist=.build/godist.list
gobuild_opts=$(cat .build/gobuild.opts | tr '\n' ' ')

for target in $(cat $distlist | grep -v '#')
do
    os=$(echo $target | cut -d '/' -f 1)
    arch=$(echo $target | cut -d '/' -f 2)

    if [ $os = "windows" ]; then
        eval CGO_ENABLED=0 GOOS=$os GOARCH=$arch go build $gobuild_opts -o $distDir/$os-$arch/$app.exe ./$mainDir && \
        zip -j $distDir/$app-$version.$os-$arch.zip $distDir/$os-$arch/$app.exe
    else
        eval CGO_ENABLED=0 GOOS=$os GOARCH=$arch go build $gobuild_opts -o $distDir/$os-$arch/$app ./$mainDir && \
        tar zcvf $distDir/$app-$version.$os-$arch.tar.gz -C $distDir/$os-$arch $app
    fi
done
