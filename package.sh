cu=`pwd`
rm -rf release/packages && echo "success deletion" || echo "not success"
mkdir -p release/packages
go build 
