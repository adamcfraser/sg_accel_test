package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/couchbase/sync_gateway/base"
	"github.com/couchbase/sync_gateway/rest"
)

// Simple Sync Gateway launcher tool.
func main() {

	signalchannel := make(chan os.Signal, 1)
	signal.Notify(signalchannel, syscall.SIGHUP)

	go func() {
		for _ = range signalchannel {
			base.Logf("SIGHUP: Reloading Config....\n")
			rest.ReloadConf()
		}
	}()

	rest.ServerMain(rest.SyncGatewayRunModeNormal)
}
