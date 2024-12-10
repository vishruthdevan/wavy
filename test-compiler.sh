if [ $# -eq 0 ]; then
    echo "Usage: ./test-compiler.sh /wavy/compiler/samples/<sample-file.vy>"
    exit 1
fi

docker build --rm -t compiler-test .
docker run -v "./compiler/samples:/wavy/compiler/samples" --rm compiler-test go test -v "./compiler" -run CompilerOutput -args -file=$1
docker rmi compiler-test
