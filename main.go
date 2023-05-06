package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/yaoapp/kun/grpc"
)

type GprcPlugin struct{ grpc.Plugin }

func main() {
	var output io.Writer = os.Stdout
	plugin := &GprcPlugin{}
	plugin.SetLogger(output, grpc.Trace)
	grpc.Serve(plugin)
}

func (plugin *GprcPlugin) Exec(name string, args ...interface{}) (*grpc.Response, error) {
	var v interface{}
	switch name {
	case "echo":
		v = plugin.Echo()
	default:
		v = map[string]interface{}{"name": name, "args": args}
	}
	bytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return &grpc.Response{Bytes: bytes, Type: "map"}, nil
}

func (plugin *GprcPlugin) Echo() map[string]interface{} {
	return map[string]interface{}{"data": "hello world"}
}
