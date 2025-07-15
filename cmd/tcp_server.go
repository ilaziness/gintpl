package main

import (
	"log/slog"

	"github.com/ilaziness/gokit/config"
	"github.com/ilaziness/gokit/server/tcp"
	"github.com/spf13/cobra"
)

var tcpCmd = &cobra.Command{
	Use:   "tcp",
	Short: "tcp server",
	Run: func(cmd *cobra.Command, args []string) {
		tcpSrv := tcp.NewDefaultTCP(&config.TCPServer{
			Address:  "127.0.0.1:8080",
			CertFile: "./cmd/server.crt",
			KeyFile:  "./cmd/server.key",
			Debug:    true,
		})
		tcpSrv.AddHandler(tcp.OpCode(0), echo)
		tcpSrv.Start()
	},
}

func echo(ctx *tcp.Context) {
	if err := ctx.Write(ctx.Payload); err != nil {
		slog.Error(err.Error())
	}
}
