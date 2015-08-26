package cl

import (
	"log"
	"testing"
	"unsafe"
)

const (
	PlatformInfoBufferSize = 1024
	DeviceInfoBufferSize   = 1024
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
	data := make([]byte, PlatformInfoBufferSize)
	dataptr := unsafe.Pointer(&data[0])
	size := uint64(0)

	err := GetPlatformInfo(platform, PLATFORM_PROFILE, PlatformInfoBufferSize, dataptr, &size)
	if err != SUCCESS {
		t.Fail()
	}
	err = GetPlatformInfo(platform, PLATFORM_VERSION, PlatformInfoBufferSize, dataptr, &size)
	if err != SUCCESS {
		t.Fail()
	}
	err = GetPlatformInfo(platform, PLATFORM_NAME, PlatformInfoBufferSize, dataptr, &size)
	if err != SUCCESS {
		t.Fail()
	}
	err = GetPlatformInfo(platform, PLATFORM_VENDOR, PlatformInfoBufferSize, dataptr, &size)
	if err != SUCCESS {
		t.Fail()
	}
	err = GetPlatformInfo(platform, PLATFORM_EXTENSIONS, PlatformInfoBufferSize, dataptr, &size)
	if err != SUCCESS {
		t.Fail()
	}
}

func TestGetAndReleaseDeviceIDs(t *testing.T) {
	devices := make([]DeviceId, 100)
	actualDid := uint32(0)
	err := GetDeviceIDs(getAnyPlatform(), DEVICE_TYPE_ALL, uint32(len(devices)), &devices[0], &actualDid)
	if err != SUCCESS {
		t.Fail()
	}
	for x := 0; x < int(actualDid); x++ {
		err = ReleaseDevice(devices[x])
		if err != SUCCESS {
			log.Println(err)
			t.Fail()
		}
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
	data := make([]byte, PlatformInfoBufferSize)
	dataptr := unsafe.Pointer(&data[0])
	size := uint64(0)

	infos := [...]DeviceInfo{DEVICE_TYPE, DEVICE_VENDOR_ID, DEVICE_MAX_COMPUTE_UNITS, DEVICE_MAX_WORK_ITEM_DIMENSIONS, DEVICE_MAX_WORK_GROUP_SIZE, DEVICE_MAX_WORK_ITEM_SIZES, DEVICE_PREFERRED_VECTOR_WIDTH_CHAR, DEVICE_PREFERRED_VECTOR_WIDTH_SHORT, DEVICE_PREFERRED_VECTOR_WIDTH_INT, DEVICE_PREFERRED_VECTOR_WIDTH_LONG, DEVICE_PREFERRED_VECTOR_WIDTH_FLOAT, DEVICE_PREFERRED_VECTOR_WIDTH_DOUBLE, DEVICE_MAX_CLOCK_FREQUENCY, DEVICE_ADDRESS_BITS, DEVICE_MAX_READ_IMAGE_ARGS, DEVICE_MAX_WRITE_IMAGE_ARGS, DEVICE_MAX_MEM_ALLOC_SIZE, DEVICE_IMAGE2D_MAX_WIDTH, DEVICE_IMAGE2D_MAX_HEIGHT, DEVICE_IMAGE3D_MAX_WIDTH, DEVICE_IMAGE3D_MAX_HEIGHT, DEVICE_IMAGE3D_MAX_DEPTH, DEVICE_IMAGE_SUPPORT, DEVICE_MAX_PARAMETER_SIZE, DEVICE_MAX_SAMPLERS, DEVICE_MEM_BASE_ADDR_ALIGN, DEVICE_MIN_DATA_TYPE_ALIGN_SIZE, DEVICE_SINGLE_FP_CONFIG, DEVICE_GLOBAL_MEM_CACHE_TYPE, DEVICE_GLOBAL_MEM_CACHELINE_SIZE, DEVICE_GLOBAL_MEM_CACHE_SIZE, DEVICE_GLOBAL_MEM_SIZE, DEVICE_MAX_CONSTANT_BUFFER_SIZE, DEVICE_MAX_CONSTANT_ARGS, DEVICE_LOCAL_MEM_TYPE, DEVICE_LOCAL_MEM_SIZE, DEVICE_ERROR_CORRECTION_SUPPORT, DEVICE_PROFILING_TIMER_RESOLUTION, DEVICE_ENDIAN_LITTLE, DEVICE_AVAILABLE, DEVICE_COMPILER_AVAILABLE, DEVICE_EXECUTION_CAPABILITIES, DEVICE_QUEUE_PROPERTIES, DEVICE_NAME, DEVICE_VENDOR, DRIVER_VERSION, DEVICE_PROFILE, DEVICE_VERSION, DEVICE_EXTENSIONS, DEVICE_PLATFORM, DEVICE_DOUBLE_FP_CONFIG}

	for _, paramName := range infos {
		err := GetDeviceInfo(device, paramName, DeviceInfoBufferSize, dataptr, &size)
		if err != SUCCESS {
			log.Println(ErrToStr(err))
			t.Fail()
		}
	}
}

func TestRetainDevice(t *testing.T) {
	device := getGPUDevice()
	err := RetainDevice(device)
	if err != SUCCESS {
		log.Println(ErrToStr(err))
		t.Fail()
	}
	err = ReleaseDevice(device)
	if err != SUCCESS {
		log.Println(ErrToStr(err))
		t.Fail()
	}
}

//from now on release your devices on every test

/*
//Not sure why this isn't valid. It might just be me that doesn't know how to use it
func TestCreateSubDevices(t *testing.T) {
	device := getGPUDevice()
	defer ReleaseDevice(device)
	outdevices := make([]DeviceId, 2)
	prop := [...]DevicePartitionProperty{DEVICE_PARTITION_EQUALLY, DevicePartitionProperty(2)}
	ret := uint32(0)
	err := CreateSubDevices(device, &prop[0], uint32(len(outdevices)), &outdevices[0], &ret)
	if err != SUCCESS {
		log.Println(ErrToStr(err))
		t.Fail()
	}
}*/

func TestCreateAndReleaseContext(t *testing.T) {
	device := getGPUDevice()
	defer ReleaseDevice(device)
	var errptr ErrorCode
	context := CreateContext(nil, 1, &device, nil, nil, &errptr)
	defer ReleaseContext(context)
	if errptr != SUCCESS {
		log.Println(ErrToStr(errptr))
		t.Fail()
	}
}

func getAnyContext() (DeviceId, Context) {
	device := getGPUDevice()
	defer ReleaseDevice(device)
	var errptr ErrorCode
	context := CreateContext(nil, 1, &device, nil, nil, &errptr)
	return device, context
}

func TestRetainContext(t *testing.T) {
	device, context := getAnyContext()
	defer ReleaseDevice(device)
	err := RetainContext(context)
	if err != SUCCESS {
		log.Println(ErrToStr(err))
		t.Fail()
	}
	err = ReleaseContext(context)
	if err != SUCCESS {
		log.Println(ErrToStr(err))
		t.Fail()
	}
}

func TestGetContextInfo(t *testing.T) {
	device, context := getAnyContext()
	defer ReleaseDevice(device)
	defer ReleaseContext(context)

	contextInfos := [...]ContextInfo{CONTEXT_REFERENCE_COUNT, CONTEXT_DEVICES, CONTEXT_PROPERTIES, CONTEXT_NUM_DEVICES}

	data := make([]byte, 1024)
	size := uint64(0)

	for _, info := range contextInfos {
		err := GetContextInfo(context, info, 1024, unsafe.Pointer(&data[0]), &size)
		if err != SUCCESS {
			t.Fail()
		}
	}
}

//Test left to write
func TestCreateSubDevices(t *testing.T) {

}

func TestCreateContextFromType(t *testing.T) {

}

func TestCreateCommandQueue(t *testing.T) {

}

func TestRetainCommandQueue(t *testing.T) {

}

func TestReleaseCommandQueue(t *testing.T) {

}

func TestGetCommandQueueInfo(t *testing.T) {

}

func TestCreateBuffer(t *testing.T) {

}

func TestCreateSubBuffer(t *testing.T) {

}

func TestCreateImage(t *testing.T) {

}

func TestRetainMemObject(t *testing.T) {

}

func TestReleaseMemObject(t *testing.T) {

}

func TestGetSupportedImageFormats(t *testing.T) {

}

func TestGetMemObjectInfo(t *testing.T) {

}

func TestGetImageInfo(t *testing.T) {

}

func TestSetMemObjectDestructorCallback(t *testing.T) {

}

func TestCreateSampler(t *testing.T) {

}

func TestRetainSampler(t *testing.T) {

}

func TestReleaseSampler(t *testing.T) {

}

func TestGetSamplerInfo(t *testing.T) {

}

func TestCreateProgramWithSource(t *testing.T) {

}

func TestCreateProgramWithBinary(t *testing.T) {

}

func TestCreateProgramWithBuiltInKernels(t *testing.T) {

}

func TestRetainProgram(t *testing.T) {

}

func TestReleaseProgram(t *testing.T) {

}

func TestBuildProgram(t *testing.T) {

}

func TestCompileProgram(t *testing.T) {

}

func TestLinkProgram(t *testing.T) {

}

func TestUnloadPlatformCompiler(t *testing.T) {

}

func TestGetProgramInfo(t *testing.T) {

}

func TestGetProgramBuildInfo(t *testing.T) {

}

func TestCreateKernel(t *testing.T) {

}

func TestCreateKernelsInProgram(t *testing.T) {

}

func TestRetainKernel(t *testing.T) {

}

func TestReleaseKernel(t *testing.T) {

}

func TestSetKernelArg(t *testing.T) {

}

func TestGetKernelInfo(t *testing.T) {

}

func TestGetKernelArgInfo(t *testing.T) {

}

func TestGetKernelWorkGroupInfo(t *testing.T) {

}

func TestWaitForEvents(t *testing.T) {

}

func TestGetEventInfo(t *testing.T) {

}

func TestCreateUserEvent(t *testing.T) {

}

func TestRetainEvent(t *testing.T) {

}

func TestReleaseEvent(t *testing.T) {

}

func TestSetUserEventStatus(t *testing.T) {

}

func TestSetEventCallback(t *testing.T) {

}

func TestGetEventProfilingInfo(t *testing.T) {

}

func TestFlush(t *testing.T) {

}

func TestFinish(t *testing.T) {

}

func TestEnqueueReadBuffer(t *testing.T) {

}

func TestEnqueueReadBufferRect(t *testing.T) {

}

func TestEnqueueWriteBuffer(t *testing.T) {

}

func TestEnqueueWriteBufferRect(t *testing.T) {

}

func TestEnqueueFillBuffer(t *testing.T) {

}

func TestEnqueueCopyBuffer(t *testing.T) {

}

func TestEnqueueCopyBufferRect(t *testing.T) {

}

func TestEnqueueReadImage(t *testing.T) {

}

func TestEnqueueWriteImage(t *testing.T) {

}

func TestEnqueueFillImage(t *testing.T) {

}

func TestEnqueueCopyImage(t *testing.T) {

}

func TestEnqueueCopyImageToBuffer(t *testing.T) {

}

func TestEnqueueCopyBufferToImage(t *testing.T) {

}

func TestEnqueueMapBuffer(t *testing.T) {

}

func TestEnqueueMapImage(t *testing.T) {

}

func TestEnqueueUnmapMemObject(t *testing.T) {

}

func TestEnqueueMigrateMemObjects(t *testing.T) {

}

func TestEnqueueNDRangeKernel(t *testing.T) {

}

func TestEnqueueTask(t *testing.T) {

}

func TestEnqueueNativeKernel(t *testing.T) {

}

func TestEnqueueMarkerWithWaitList(t *testing.T) {

}

func TestEnqueueBarrierWithWaitList(t *testing.T) {

}

func TestGetExtensionFunctionAddressForPlatform(t *testing.T) {

}
