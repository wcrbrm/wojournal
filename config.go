package main

import (
	"strconv"
	"strings"

	cli "github.com/jawher/mow.cli"
)

var (
	appName = app.String(cli.StringOpt{
		Name:   "name",
		Desc:   "",
		EnvVar: "APP_NAME",
		Value:  "core_offers",
	})
	envName = app.String(cli.StringOpt{
		Name:   "env",
		Desc:   "",
		EnvVar: "APP_ENV",
		Value:  "local",
	})
	logLevel = app.String(cli.StringOpt{
		Name:   "l log-level",
		Desc:   "Available levels: error, warn, info, debug.",
		EnvVar: "APP_LOG_LEVEL",
		Value:  "info",
	})

	enableMetrics = app.Bool(cli.BoolOpt{
		Name:   "m metrics",
		Desc:   "Enable prometheus metrics",
		EnvVar: "ENABLE_PROMETHEUS",
		Value:  true,
	})

	webListenAddr = app.String(cli.StringOpt{
		Name:   "http-addr",
		Desc:   "Listen address for HTTP debugging using external tools",
		EnvVar: "APP_HTTP_ADDRESS",
		Value:  "0.0.0.0:9091",
	})

	goMaxProcs = app.String(cli.StringOpt{
		Name:   "p go-procs",
		Desc:   "The maximum number of CPUs that can be used simultaneously by Go runtime.",
		EnvVar: "AN_GOMAXPROCS",
		Value:  "128",
	})
)

func toList(s string) []string {
	return strings.Split(s, ",")
}

func toBool(s string) bool {
	switch strings.ToLower(s) {
	case "true", "1", "t", "yes":
		return true
	default:
		return false
	}
}

func toNatural(s string, defaults uint64) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		// defaults in case of incorrect or empty "" value
		return int(defaults)
	} else if i < 0 {
		// not defaults, because nobody expects +100 while specifying -100
		return 0
	}
	return int(i)
}
