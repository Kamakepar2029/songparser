cu=`pwd`
rm -rf $1/release/packages && echo "success deletion" || echo "not success"
mkdir -p $1/release/packages/
os_all='linux windows darwin freebsd'
arch_all='386 amd64 arm arm64 mips64 mips64le mips mipsle'
cd $1/release/packages
for os in $os_all; do
    for arch in $arch_all; do
      set GOOS=$os
      set GOARCH=$arch
      if [ $os = "windows" ]; then
        go build -o $1"_"$os"_"$arch".exe" && echo "Success build for arch "$arch" and os "$os || echo "No problem"
      else
        go build -o $1"_"$os"_"$arch && echo "Success build for arch "$arch" and os "$os || echo "No problem"
      fi
      
      
echo "Success Build"
cd $cu
