VERSION=$(git describe --tags)
THIS_GOOS=$(go env GOOS)
THIS_GOARCH=$(go env GOARCH)
RELEASE_DIR="release"
ARTIFACT_DIR="artifact"
RELEASE_TARGET=("linux/arm" "linux/arm64" "linux/amd64" "windows/amd64" "darwin/amd64" "darwin/arm64")
for t in ${RELEASE_TARGET[@]}
do
    DIST=$RELEASE_DIR"/"$t"/mcpick_"$VERSION
    tmp=(${t//// })
    GOOS=${tmp[0]}
    GOARCH=${tmp[1]}
    SUFFIX=""
    if [ $GOOS = "windows" ]
        then
            SUFFIX=".exe"
    fi
    DIST_BIN=$DIST"/mcpick"$SUFFIX
    go build -o $DIST_BIN
    cp LICENSE README.md $DIST
    mkdir -p $ARTIFACT_DIR 2>/dev/null
    if [ $GOOS = "linux" ]
        then
            tar cfvz $ARTIFACT_DIR"/mcpick_"$GOOS"_"$GOARCH".tar.gz" $DIST_BIN
    else
        zip -r $ARTIFACT_DIR"/mcpick_"$GOOS"_"$GOARCH".zip" $DIST_BIN
    fi
done
GOOS=THIS_GOOS
GOARCH=THIS_GOARCH
rm -rf $RELEASE_DIR