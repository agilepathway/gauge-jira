package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/agilepathway/gauge-jira/export"
	"github.com/agilepathway/gauge-jira/gauge_messages"
	"github.com/agilepathway/gauge-jira/util"
	"google.golang.org/grpc"
)

const (
	gaugeSpecsDir = "GAUGE_SPEC_DIRS"
	fileSeparator = "||"
)

var projectRoot = util.GetProjectRoot()

type handler struct {
	server *grpc.Server
}

func (h *handler) GenerateDocs(c context.Context, m *gauge_messages.SpecDetails) (*gauge_messages.Empty, error) {
	var files []string
	for _, arg := range strings.Split(os.Getenv(gaugeSpecsDir), fileSeparator) {
		files = append(files, util.GetFiles(arg)...)
	}
	for _, file := range files {
		export.Spec(file)
	}
	fmt.Println("Successfully exported specs to Jira")
	return &gauge_messages.Empty{}, nil
}

func (h *handler) Kill(c context.Context, m *gauge_messages.KillProcessRequest) (*gauge_messages.Empty, error) {
	defer h.stopServer()
	return &gauge_messages.Empty{}, nil
}

func (h *handler) stopServer() {
	h.server.Stop()
}

func main() {
	os.Chdir(projectRoot)
	address, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		util.Fatal("failed to start server.", err)
	}
	l, err := net.ListenTCP("tcp", address)
	if err != nil {
		util.Fatal("failed to start server.", err)
	}
	server := grpc.NewServer(grpc.MaxRecvMsgSize(1024 * 1024 * 10))
	h := &handler{server: server}
	gauge_messages.RegisterDocumenterServer(server, h)
	fmt.Println(fmt.Sprintf("Listening on port:%d", l.Addr().(*net.TCPAddr).Port))
	server.Serve(l)
}
