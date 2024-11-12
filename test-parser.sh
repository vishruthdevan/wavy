if [ $# -eq 0 ]; then
    echo "Usage: ./test-parser.sh /wavy/parser/samples/<sample-file.vy>"
    exit 1
fi

docker build --rm -t parser-test .
docker run -v "./parser/samples:/wavy/parser/samples" --rm parser-test go test -v "./parser" -run ParserOutput -args -file=$1
docker rmi lexer-test
