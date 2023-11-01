package request

import (
	"context"
	"github.com/RincewindsHat/thola/internal/network"
)

// CheckSNMPRequest
//
// CheckSNMPRequest is the request struct for the check snmp request.
//
// swagger:model
type CheckSNMPRequest struct {
	CheckDeviceRequest
}

func (r *CheckSNMPRequest) setupConnection(_ context.Context) (*network.RequestDeviceConnection, error) {
	// in this specific case, we do not setup any connection here, because this check should check if an snmp
	// connection is possible. it will be checked in the process() function.
	return &network.RequestDeviceConnection{}, nil
}

// CheckSNMPResponse
//
// CheckSNMPResponse is a response struct for the check snmp request.
//
// swagger:model
type CheckSNMPResponse struct {
	CheckResponse
	SuccessfulSnmpCredentials *network.SNMPCredentials `yaml:"successful_snmp_credentials" json:"successful_snmp_credentials" xml:"successful_snmp_credentials"`
}
