package main

import (
	"github.com/go-gl/cl/v1.2/cl"
	"log"
	"unsafe"
)

const (
	// DataSize is the size of the buffer for the string we're going to get back
	// from the CL device.
	DataSize = 1024
)

func main() {
	StatInfo()
}

// StatInfo prints a bunch of usefull information about the currently available
// CL devices
func StatInfo() {
	ids := make([]cl.PlatformID, 100)
	actual := uint32(0)
	cl.GetPlatformIDs(uint32(len(ids)), &ids[0], &actual)
	for x := 0; x < int(actual); x++ {
		data := make([]byte, DataSize)
		size := uint64(0)
		cl.GetPlatformInfo(ids[x], cl.PLATFORM_PROFILE, DataSize, unsafe.Pointer(&data[0]), &size)
		profilestring := string(data[0:size])
		cl.GetPlatformInfo(ids[x], cl.PLATFORM_VERSION, DataSize, unsafe.Pointer(&data[0]), &size)
		versionstring := string(data[0:size])

		cl.GetPlatformInfo(ids[x], cl.PLATFORM_NAME, DataSize, unsafe.Pointer(&data[0]), &size)
		namestring := string(data[0:size])
		cl.GetPlatformInfo(ids[x], cl.PLATFORM_VENDOR, DataSize, unsafe.Pointer(&data[0]), &size)
		vendorstring := string(data[0:size])
		cl.GetPlatformInfo(ids[x], cl.PLATFORM_EXTENSIONS, DataSize, unsafe.Pointer(&data[0]), &size)
		extensionsstring := string(data[0:size])
		log.Print("PLATFORM_PROFILE:\t\t", profilestring)
		log.Print("PLATFORM_VERSION:\t\t", versionstring)
		log.Print("PLATFORM_NAME:\t\t", namestring)
		log.Print("PLATFORM_VENDOR:\t\t", vendorstring)
		log.Print("PLATFORM_EXTENSIONS:\t", extensionsstring)

		devices := make([]cl.DeviceId, 100)
		actualDid := uint32(0)
		cl.GetDeviceIDs(ids[x], cl.DEVICE_TYPE_ALL, uint32(len(devices)), &devices[0], &actualDid)
		log.Println("Devices: ")
		for y := 0; y < int(actualDid); y++ {
			cl.GetDeviceInfo(devices[y], cl.DEVICE_NAME, DataSize, unsafe.Pointer(&data[0]), &size)
			deviceName := string(data[0:size])
			log.Print("\tname: "+deviceName+" @ ", &devices[y])
		}
	}
}
