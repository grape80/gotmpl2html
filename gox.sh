#!/bin/sh

. .build/.env

app=$APP_NAME
version=$VERSION
mainDir=$MAIN_DIR
distDir=$DIST_DIR
distlist=.build/godist.list # for local

gobuild_opts=$(cat .build/gobuild.opts | tr '\n' ' ')

if [ $GH_OS_ARCH != "" ]; then
    # for github actions
    distlist=$GH_OS_ARCH
fi

for target in $(cat $distlist | grep -v '#')
do
    os=$(echo $target | cut -d '/' -f 1)
    arch=$(echo $target | cut -d '/' -f 2)

    if [ $os = "windows" ]; then
        CGO_ENABLED=0 GOOS=$os GOARCH=$arch eval go build $gobuild_opts -o $distDir/$os-$arch/$app.exe ./$mainDir && \
        zip -j $distDir/$app-$version.$os-$arch.zip $distDir/$os-$arch/$app.exe
    else
        CGO_ENABLED=0 GOOS=$os GOARCH=$arch eval go build $gobuild_opts -o $distDir/$os-$arch/$app ./$mainDir && \
        tar zcvf $distDir/$app-$version.$os-$arch.tar.gz -C $distDir/$os-$arch $app
    fi
done
