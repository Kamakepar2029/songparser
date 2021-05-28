cu=`pwd`
rm GOARCH* && echo "success deletion" || echo "not_success"
rm -rf $1/release/packages && echo "success deletion" || echo "not success"
mkdir -p $1/release/packages/
os_all='linux windows darwin freebsd'
arch_all='386 amd64 arm arm64 mips64 mips64le mips mipsle'
for os in $os_all; do
    for arch in $arch_all; do
      set GOOS=$os
      set GOARCH=$arch
      if [ $os = "windows" ]; then
        go build -o $1"_"$os"_"$arch".exe" && echo "Success build for arch "$arch" and os "$os || echo "No problem"
        mv $1"_"$os"_"$arch".exe" $1/release/packages && echo "Move success" || echo "Move not success"
      else
        go build -o $1"_"$os"_"$arch && echo "Success build for arch "$arch" and os "$os || echo "No problem"
        mv $1"_"$os"_"$arch".exe" $1/release/packages && echo "Move success" || echo "Move not success"
      fi
    done
done
echo "Success Build"
cd $cu
