package tests

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api/apistruct"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/require"
	"github.com/textileio/powergate/lotus"
	"github.com/textileio/powergate/util"
)

// LaunchDevnetDocker launches the devnet docker image.
func LaunchDevnetDocker(t *testing.T, numMiners int, ipfsMaddr string, mountVolumes bool) *dockertest.Resource {
	pool, err := dockertest.NewPool("")
	require.NoError(t, err)
	envs := []string{
		devnetEnv("NUMMINERS", strconv.Itoa(numMiners)),
		devnetEnv("SPEED", "500"),
		devnetEnv("IPFSADDR", ipfsMaddr),
		devnetEnv("BIGSECTORS", false),
	}
	var mounts []string
	if mountVolumes {
		mounts = append(mounts, "/tmp/powergate:/tmp/powergate")
	}

	repository := "textile/lotus-devnet"
	tag := "butterfly-v7.10.0-3"
	lotusDevnet, err := pool.RunWithOptions(&dockertest.RunOptions{Repository: repository, Tag: tag, Env: envs, Mounts: mounts})
	require.NoError(t, err)
	err = lotusDevnet.Expire(180)
	require.NoError(t, err)
	time.Sleep(time.Second * time.Duration(2+numMiners))
	t.Cleanup(func() {
		err = pool.Purge(lotusDevnet)
		require.NoError(t, err)
	})
	debug := false
	if debug {
		go func() {
			opts := docker.LogsOptions{
				Context: context.Background(),

				Stderr:      true,
				Stdout:      true,
				Follow:      true,
				Timestamps:  true,
				RawTerminal: true,

				Container: lotusDevnet.Container.ID,

				OutputStream: os.Stdout,
			}

			err = pool.Client.Logs(opts)
			require.NoError(t, err)
		}()
	}
	return lotusDevnet
}

// CreateLocalDevnetWithIPFS creates a local devnet connected to an IPFS node.
func CreateLocalDevnetWithIPFS(t *testing.T, numMiners int, ipfsMaddr string, mountVolumes bool) (*apistruct.FullNodeStruct, address.Address, []address.Address) {
	lotusDevnet := LaunchDevnetDocker(t, numMiners, ipfsMaddr, mountVolumes)
	c, cls, err := lotus.New(util.MustParseAddr("/ip4/127.0.0.1/tcp/"+lotusDevnet.GetPort("7777/tcp")), "", 1)
	require.NoError(t, err)
	t.Cleanup(func() { cls() })
	ctx := context.Background()
	addr, err := c.WalletDefaultAddress(ctx)
	if err != nil {
		t.Fatal(err)
	}

	miners, err := c.StateListMiners(ctx, types.EmptyTSK)
	if err != nil {
		t.Fatal(err)
	}

	return c, addr, miners
}

// CreateLocalDevnet returns an API client that targets a local devnet with numMiners number
// of miners. Refer to http://github.com/textileio/local-devnet for more information.
func CreateLocalDevnet(t *testing.T, numMiners int) (*apistruct.FullNodeStruct, address.Address, []address.Address) {
	return CreateLocalDevnetWithIPFS(t, numMiners, "", true)
}

func devnetEnv(name string, value interface{}) string {
	return fmt.Sprintf("TEXLOTUSDEVNET_%s=%s", name, value)
}
