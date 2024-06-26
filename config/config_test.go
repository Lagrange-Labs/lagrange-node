package config

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

func TestConfigLoad(t *testing.T) {
	// test with nil context
	_, err := Load(nil)
	require.NoError(t, err)

	// test with empty config file path
	ctx := cli.NewContext(nil, nil, nil)
	_, err = Load(ctx)
	require.NoError(t, err)

	// test with invalid config file path
	flagSet := flag.NewFlagSet("test", flag.ContinueOnError)
	flagSet.String(FlagCfg, "invalid", "")
	ctx = cli.NewContext(nil, flagSet, nil)
	require.NoError(t, ctx.Set(FlagCfg, "invalid"))
	_, err = Load(ctx)
	require.Error(t, err)
}

func TestEnvrionmentConfig(t *testing.T) {
	// set environment variables
	require.NoError(t, os.Setenv("LAGRANGE_NODE_CLIENT_GRPCURLS", "127.0.0.1:9090,127.0.0.1:9091"))
	cfg, err := Load(nil)
	require.NoError(t, err)
	require.NotNil(t, cfg)
	require.Len(t, cfg.Client.GrpcURLs, 2)
}
