package cl

import (
	"log"
	"testing"
	"unsafe"
)

const (
	PLATFORM_INFO_BUFFER_SIZE = 1024
	DEVICE_INFO_BUFFER_SIZE   = 1024
)

func TestGetPlatformIDs(t *testing.T) {
	ids := make([]PlatformID, 100)
	actual := uint32(0)
	err := GetPlatformIDs(uint32(len(ids)), &ids[0], &actual)
	if err != SUCCESS {
		t.Fail()
	}
}

func getAnyPlatform() PlatformID {
	ids := make([]PlatformID, 1)
	actual := uint32(0)
	err := GetPlatformIDs(uint32(len(ids)), &ids[0], &actual)
	if err != SUCCESS {
		return nil
	}
	return ids[0]
}

func TestGetPlatformInfo(t *testing.T) {
	platform := getAnyPlatform()
	data := make([]byte, PLATFORM_INFO_BUFFER_SIZE)
	dataptr := unsafe.Pointer(&data[0])
	size := uint64(0)

	err := GetPlatformInfo(platform, PLATFORM_PROFILE, PLATFORM_INFO_BUFFER_SIZE, dataptr, &size)
	if err != SUCCESS {
		t.Fail()
	}
	err = GetPlatformInfo(platform, PLATFORM_VERSION, PLATFORM_INFO_BUFFER_SIZE, dataptr, &size)
	if err != SUCCESS {
		t.Fail()
	}
	err = GetPlatformInfo(platform, PLATFORM_NAME, PLATFORM_INFO_BUFFER_SIZE, dataptr, &size)
	if err != SUCCESS {
		t.Fail()
	}
	err = GetPlatformInfo(platform, PLATFORM_VENDOR, PLATFORM_INFO_BUFFER_SIZE, dataptr, &size)
	if err != SUCCESS {
		t.Fail()
	}
	err = GetPlatformInfo(platform, PLATFORM_EXTENSIONS, PLATFORM_INFO_BUFFER_SIZE, dataptr, &size)
	if err != SUCCESS {
		t.Fail()
	}
}

func TestGetDeviceIDs(t *testing.T) {
	devices := make([]DeviceId, 100)
	actualDid := uint32(0)
	err := GetDeviceIDs(getAnyPlatform(), DEVICE_TYPE_ALL, uint32(len(devices)), &devices[0], &actualDid)
	if err != SUCCESS {
		t.Fail()
	}
}

func getGPUDevice() DeviceId {
	devices := make([]DeviceId, 1)
	actualDid := uint32(0)
	err := GetDeviceIDs(getAnyPlatform(), DEVICE_TYPE_GPU, uint32(len(devices)), &devices[0], &actualDid)
	if err != SUCCESS {
		return nil
	}
	return devices[0]
}

func TestGetDeviceInfo(t *testing.T) {
	device := getGPUDevice()
	data := make([]byte, PLATFORM_INFO_BUFFER_SIZE)
	dataptr := unsafe.Pointer(&data[0])
	size := uint64(0)

	infos := [...]DeviceInfo{DEVICE_TYPE,
		DEVICE_VENDOR_ID,
		DEVICE_MAX_COMPUTE_UNITS,
		DEVICE_MAX_WORK_ITEM_DIMENSIONS,
		DEVICE_MAX_WORK_GROUP_SIZE,
		DEVICE_MAX_WORK_ITEM_SIZES,
		DEVICE_PREFERRED_VECTOR_WIDTH_CHAR,
		DEVICE_PREFERRED_VECTOR_WIDTH_SHORT,
		DEVICE_PREFERRED_VECTOR_WIDTH_INT,
		DEVICE_PREFERRED_VECTOR_WIDTH_LONG,
		DEVICE_PREFERRED_VECTOR_WIDTH_FLOAT,
		DEVICE_PREFERRED_VECTOR_WIDTH_DOUBLE,
		DEVICE_MAX_CLOCK_FREQUENCY,
		DEVICE_ADDRESS_BITS,
		DEVICE_MAX_READ_IMAGE_ARGS,
		DEVICE_MAX_WRITE_IMAGE_ARGS,
		DEVICE_MAX_MEM_ALLOC_SIZE,
		DEVICE_IMAGE2D_MAX_WIDTH,
		DEVICE_IMAGE2D_MAX_HEIGHT,
		DEVICE_IMAGE3D_MAX_WIDTH,
		DEVICE_IMAGE3D_MAX_HEIGHT,
		DEVICE_IMAGE3D_MAX_DEPTH,
		DEVICE_IMAGE_SUPPORT,
		DEVICE_MAX_PARAMETER_SIZE,
		DEVICE_MAX_SAMPLERS,
		DEVICE_MEM_BASE_ADDR_ALIGN,
		DEVICE_MIN_DATA_TYPE_ALIGN_SIZE,
		DEVICE_SINGLE_FP_CONFIG,
		DEVICE_GLOBAL_MEM_CACHE_TYPE,
		DEVICE_GLOBAL_MEM_CACHELINE_SIZE,
		DEVICE_GLOBAL_MEM_CACHE_SIZE,
		DEVICE_GLOBAL_MEM_SIZE,
		DEVICE_MAX_CONSTANT_BUFFER_SIZE,
		DEVICE_MAX_CONSTANT_ARGS,
		DEVICE_LOCAL_MEM_TYPE,
		DEVICE_LOCAL_MEM_SIZE,
		DEVICE_ERROR_CORRECTION_SUPPORT,
		DEVICE_PROFILING_TIMER_RESOLUTION,
		DEVICE_ENDIAN_LITTLE,
		DEVICE_AVAILABLE,
		DEVICE_COMPILER_AVAILABLE,
		DEVICE_EXECUTION_CAPABILITIES,
		DEVICE_QUEUE_PROPERTIES,
		DEVICE_NAME,
		DEVICE_VENDOR,
		DRIVER_VERSION,
		DEVICE_PROFILE,
		DEVICE_VERSION,
		DEVICE_EXTENSIONS,
		DEVICE_PLATFORM,
		DEVICE_DOUBLE_FP_CONFIG}

	for _, paramName := range infos {
		err := GetDeviceInfo(device, paramName, DEVICE_INFO_BUFFER_SIZE, dataptr, &size)
		if err != SUCCESS {
			log.Println(ErrToStr(err))
			t.Fail()
		}
	}
}
