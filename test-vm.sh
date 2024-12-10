if [ $# -eq 0 ]; then
    echo "Usage: ./test-vm.sh /wavy/vm/samples/<sample-file.vy>"
    exit 1
fi

docker build --rm -t vm-test .
docker run -v "./vm/samples:/wavy/vm/samples" --rm vm-test go test -v "./vm" -run TestVMOutput -args -file=$1
docker rmi vm-test
