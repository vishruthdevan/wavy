if [ $# -eq 0 ]; then
    echo "Usage: ./test-lexer.sh /wavy/<path-to-sample-file.vy>"
    exit 1
fi

docker build --rm -t lexer-test .
docker run -v "./samples:/wavy/samples" --rm lexer-test go test -v /wavy/lexer/... -args -file=$1
docker rmi lexer-test
