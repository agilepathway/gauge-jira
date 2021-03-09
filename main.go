package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/agilepathway/gauge-jira/gauge_messages"
	"github.com/agilepathway/gauge-jira/internal/env"
	"github.com/agilepathway/gauge-jira/internal/jira"
	"github.com/agilepathway/gauge-jira/util"
	"google.golang.org/grpc"
)

const (
	gaugeSpecsDir = "GAUGE_SPEC_DIRS"
	fileSeparator = "||"
)

var projectRoot = util.GetProjectRoot() //nolint:gochecknoglobals

type handler struct {
	server *grpc.Server
}

func (h *handler) GenerateDocs(c context.Context, m *gauge_messages.SpecDetails) (*gauge_messages.Empty, error) {
	var specFilenames []string //nolint:prealloc
	for _, arg := range strings.Split(os.Getenv(gaugeSpecsDir), fileSeparator) {
		specFilenames = append(specFilenames, util.GetFiles(arg)...)
	}

	jira.PublishSpecs(specFilenames)

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
	checkRequiredConfigVars()

	err := os.Chdir(projectRoot)
	util.Fatal("failed to change directory to project root.", err)

	address, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	util.Fatal("failed to start server.", err)

	l, err := net.ListenTCP("tcp", address)
	util.Fatal("TCP listening failed.", err)

	server := grpc.NewServer(grpc.MaxRecvMsgSize(1024 * 1024 * 10)) //nolint:gomnd
	h := &handler{server: server}
	gauge_messages.RegisterDocumenterServer(server, h)
	fmt.Printf("Listening on port:%d /n", l.Addr().(*net.TCPAddr).Port)
	server.Serve(l) //nolint:errcheck,gosec
}

func checkRequiredConfigVars() {
	env.GetRequired("JIRA_BASE_URL")
	env.GetRequired("JIRA_USERNAME")
	env.GetRequired("JIRA_TOKEN")
}
