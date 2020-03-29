module github.com/islishude/gomodtest

go 1.14

require (
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/zebozhuang/gomodtest/helloworld v0.0.0-00010101000000-000000000000
	github.com/zebozhuang/gomodtest/route_guide/routeguide v0.0.0-00010101000000-000000000000
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/grpc v1.28.0
	rsc.io/quote/v3 v3.1.0
	rsc.io/sampler v1.99.99 // indirect
)

replace github.com/zebozhuang/gomodtest/helloworld => ./helloworld

replace github.com/zebozhuang/gomodtest/route_guide/routeguide => ./route_guide/routeguide
