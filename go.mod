module github.com/halladj/dis-log

require (
	github.com/armon/go-metrics v0.0.0-20190430140413-ec5e00d3c878 // indirect
	github.com/casbin/casbin v1.9.1
	github.com/gorilla/mux v1.8.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/hashicorp/go-msgpack v0.5.5 // indirect
	github.com/hashicorp/raft v1.1.1 // indirect
	github.com/hashicorp/serf v0.8.5
	github.com/stretchr/testify v1.4.0
	github.com/travisjeffery/go-dynaport v0.0.0-20171218080632-f8768fb615d5
	github.com/tysonmote/gommap v0.0.3
	go.uber.org/zap v1.10.0
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80 // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0
)

go 1.13

replace github.com/hashicorp/raft-boltdb => github.com/travisjeffery/raft-boltdb v1.0.0
