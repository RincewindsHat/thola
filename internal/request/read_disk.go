package request

import "github.com/RincewindsHat/thola/internal/device"

// ReadDiskRequest
//
// ReadDiskRequest is the request struct for the read disk request.
//
// swagger:model
type ReadDiskRequest struct {
	ReadRequest
}

// ReadDiskResponse
//
// ReadDiskResponse is the response struct for the read disk response.
//
// swagger:model
type ReadDiskResponse struct {
	Disk device.DiskComponent `yaml:"disk" json:"disk" xml:"disk"`
	ReadResponse
}
