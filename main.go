package main

import (
	"encoding/json"
	"yaoplugin/cmd/hello"
	"yaoplugin/utils"

	"github.com/yaoapp/kun/grpc"
)

type GprcPlugin struct{ grpc.Plugin }

func main() {
	defer utils.CloseLog()
	plugin := &GprcPlugin{}
	grpc.Serve(plugin)
}

func (plugin *GprcPlugin) Exec(name string, args ...interface{}) (*grpc.Response, error) {
	var v interface{}
	switch name {
	case "hello":
		v = hello.Echo(args...)
	default:
		v = map[string]interface{}{"name": name, "args": args}
	}

	bytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return &grpc.Response{Bytes: bytes, Type: "map"}, nil
}

// func main() {
// 	plugin := &GprcPlugin{}
// 	plugin.Exec("hello")
// }
