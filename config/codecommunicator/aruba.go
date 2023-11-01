package codecommunicator

import (
	"context"
	"github.com/RincewindsHat/thola/internal/device"
)

type arubaCommunicator struct {
	codeCommunicator
}

func (c *arubaCommunicator) GetDiskComponentStorages(ctx context.Context) ([]device.DiskComponentStorage, error) {
	communicator := linuxCommunicator{c.codeCommunicator}
	return communicator.GetDiskComponentStorages(ctx)
}
