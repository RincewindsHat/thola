package codecommunicator

import (
	"context"
	"github.com/RincewindsHat/thola/internal/device"
)

type vmwareESXiCommunicator struct {
	codeCommunicator
}

func (c *vmwareESXiCommunicator) GetDiskComponentStorages(ctx context.Context) ([]device.DiskComponentStorage, error) {
	communicator := linuxCommunicator{c.codeCommunicator}
	return communicator.GetDiskComponentStorages(ctx)
}
