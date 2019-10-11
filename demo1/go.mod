module demo1

go 1.12

require (
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro v1.11.1
	github.com/micro/go-plugins v1.3.0
	go-micro_demo v0.0.0
)

replace go-micro_demo v0.0.0 => ./
