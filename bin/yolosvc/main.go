package main

import (
	"context"
	_ "embed"
	"fmt"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/gookit/slog"

	"github.com/ChengWu-NJ/yolosvc/pkg/config"
	"github.com/ChengWu-NJ/yolosvc/pkg/grpcsvc"
)

const (
	APPNAME = `yolosvc`
)

//go:generate sh -c "printf %s $(go list -m -versions -json $(go list -m)@$(git branch --show-current) | jq -r .Version) > commit.txt"
//go:embed commit.txt
var CommitVersion string

func main() {
	var err error

	flags := flag.NewFlagSet(APPNAME, flag.ExitOnError)

	showVer := flags.Bool("version", false, "show version and exit")

	logDebugInfo := flags.Bool("debug", false, "log debug info")

	err = flags.Parse(os.Args)
	if err != nil {
		slog.Error(err)
		return
	}

	if *showVer {
		fmt.Printf("%s %s\n", APPNAME, CommitVersion)
		return
	}

	slog.SetLogLevel(slog.InfoLevel)

	if *logDebugInfo {
		slog.SetLogLevel(slog.DebugLevel)
	}

	slog.Infof("begin to creat a yolosvc...")
	address := fmt.Sprintf(`:%d`, config.GlobalConfig.PortOfGrpcSvc)
	err = grpcsvc.Run(context.Background(), `tcp`, address)

	slog.Infof("yolosvc exit with errStatus:[%+v]", err)
}
