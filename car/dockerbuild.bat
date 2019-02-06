SET GOOS=linux
SET CGO_ENABLED=0

go build -a -installsuffix cgo -o main .

docker build . -t %1