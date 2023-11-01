package cmd

import (
	"github.com/RincewindsHat/thola/internal/request"
	"github.com/spf13/cobra"
)

func init() {
	addDeviceFlags(readHardwareHealth)
	readCMD.AddCommand(readHardwareHealth)
}

var readHardwareHealth = &cobra.Command{
	Use:   "hardware-health",
	Short: "Read out the hardware health of a device",
	Long:  "Read out the hardware health of a device.",
	Run: func(cmd *cobra.Command, args []string) {
		request := request.ReadHardwareHealthRequest{
			ReadRequest: getReadRequest(args[0]),
		}
		handleRequest(&request)
	},
}
