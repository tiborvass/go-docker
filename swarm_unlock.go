package docker // import "docker.io/go-docker"

import (
	"docker.io/go-docker/api/types/swarm"
	"golang.org/x/net/context"
)

// SwarmUnlock unlocks locked swarm.
func (cli *Client) SwarmUnlock(ctx context.Context, req swarm.UnlockRequest) error {
	serverResp, err := cli.post(ctx, "/swarm/unlock", nil, req, nil)
	ensureReaderClosed(serverResp)
	return err
}
