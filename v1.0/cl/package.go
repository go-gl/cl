package cl

/*
#cgo darwin  LDFLAGS: -framework OpenCL
#cgo linux   LDFLAGS: -lCL
#cgo windows LDFLAGS: -lopencl32
#ifdef __APPLE__
#include "OpenCL/opencl.h"
#else
#include "CL/opencl.h"
#endif
extern void contextErrorCallback(char *, void *, size_t, void *);
extern void programCallback(cl_program, void*);
extern void programObjectCompileCompleteCallback(cl_program, void*);
*/
import "C"
import (
	"log"
	"reflect"
	"strings"
	"unsafe"
)

func init() {
	mochHolder = make(map[*memObjectCallbackHolder]struct{})
	pochHolder = make(map[*programCallbackHolder]struct{})
	ecbHolder = make(map[*eventCbHolder]struct{})
}

type ErrorCode int32
type Bool uint32
type PlatformInfo uint32
type DeviceType uint32
type DeviceInfo uint32
type DeviceFpConfig uint32
type DeviceMemCacheType uint32
type DeviceLocalMemType uint32
type DeviceExecCapabilities uint32
type CommandQueueProperties uint32
type ContextInfo uint32
type ContextProperties uint32
type DevicePartitionProperty uint32
type DeviceAffinityDomain uint32
type CommandQueueInfo uint32
type MemFlags uint32
type MemMigrationFlags uint32
type ChannelOrder uint32
type ChannelType uint32
type MemObjectType uint32
type MemInfo uint32
type ImageInfo uint32
type AddressingMode uint32
type FilterMode uint32
type SamplerInfo uint32
type MapFlags uint32
type ProgramInfo uint32
type ProgramBuildInfo uint32
type ProgramBinaryType uint32
type BuildStatus int32
type KernelInfo uint32
type KernelArgInfo uint32
type KernelArgAddressQualifier uint32
type KernelArgAccessQualifier uint32
type KernelArgTypeQualifer uint32
type KernelWorkGroupInfo uint32
type EventInfo uint32
type CommandType uint32
type BufferCreateType uint32
type ProfilingInfo uint32
type CommandExecutionStatus int32

const (
	SUCCESS                       ErrorCode = C.CL_SUCCESS
	DEVICE_NOT_FOUND              ErrorCode = C.CL_DEVICE_NOT_FOUND
	DEVICE_NOT_AVAILABLE          ErrorCode = C.CL_DEVICE_NOT_AVAILABLE
	COMPILER_NOT_AVAILABLE        ErrorCode = C.CL_COMPILER_NOT_AVAILABLE
	MEM_OBJECT_ALLOCATION_FAILURE ErrorCode = C.CL_MEM_OBJECT_ALLOCATION_FAILURE
	OUT_OF_RESOURCES              ErrorCode = C.CL_OUT_OF_RESOURCES
	OUT_OF_HOST_MEMORY            ErrorCode = C.CL_OUT_OF_HOST_MEMORY
	PROFILING_INFO_NOT_AVAILABLE  ErrorCode = C.CL_PROFILING_INFO_NOT_AVAILABLE
	MEM_COPY_OVERLAP              ErrorCode = C.CL_MEM_COPY_OVERLAP
	IMAGE_FORMAT_MISMATCH         ErrorCode = C.CL_IMAGE_FORMAT_MISMATCH
	IMAGE_FORMAT_NOT_SUPPORTED    ErrorCode = C.CL_IMAGE_FORMAT_NOT_SUPPORTED
	BUILD_PROGRAM_FAILURE         ErrorCode = C.CL_BUILD_PROGRAM_FAILURE
	MAP_FAILURE                   ErrorCode = C.CL_MAP_FAILURE

	INVALID_VALUE                   ErrorCode = C.CL_INVALID_VALUE
	INVALID_DEVICE_TYPE             ErrorCode = C.CL_INVALID_DEVICE_TYPE
	INVALID_PLATFORM                ErrorCode = C.CL_INVALID_PLATFORM
	INVALID_DEVICE                  ErrorCode = C.CL_INVALID_DEVICE
	INVALID_CONTEXT                 ErrorCode = C.CL_INVALID_CONTEXT
	INVALID_QUEUE_PROPERTIES        ErrorCode = C.CL_INVALID_QUEUE_PROPERTIES
	INVALID_COMMAND_QUEUE           ErrorCode = C.CL_INVALID_COMMAND_QUEUE
	INVALID_HOST_PTR                ErrorCode = C.CL_INVALID_HOST_PTR
	INVALID_MEM_OBJECT              ErrorCode = C.CL_INVALID_MEM_OBJECT
	INVALID_IMAGE_FORMAT_DESCRIPTOR ErrorCode = C.CL_INVALID_IMAGE_FORMAT_DESCRIPTOR
	INVALID_IMAGE_SIZE              ErrorCode = C.CL_INVALID_IMAGE_SIZE
	INVALID_SAMPLER                 ErrorCode = C.CL_INVALID_SAMPLER
	INVALID_BINARY                  ErrorCode = C.CL_INVALID_BINARY
	INVALID_BUILD_OPTIONS           ErrorCode = C.CL_INVALID_BUILD_OPTIONS
	INVALID_PROGRAM                 ErrorCode = C.CL_INVALID_PROGRAM
	INVALID_PROGRAM_EXECUTABLE      ErrorCode = C.CL_INVALID_PROGRAM_EXECUTABLE
	INVALID_KERNEL_NAME             ErrorCode = C.CL_INVALID_KERNEL_NAME
	INVALID_KERNEL_DEFINITION       ErrorCode = C.CL_INVALID_KERNEL_DEFINITION
	INVALID_KERNEL                  ErrorCode = C.CL_INVALID_KERNEL
	INVALID_ARG_INDEX               ErrorCode = C.CL_INVALID_ARG_INDEX
	INVALID_ARG_VALUE               ErrorCode = C.CL_INVALID_ARG_VALUE
	INVALID_ARG_SIZE                ErrorCode = C.CL_INVALID_ARG_SIZE
	INVALID_KERNEL_ARGS             ErrorCode = C.CL_INVALID_KERNEL_ARGS
	INVALID_WORK_DIMENSION          ErrorCode = C.CL_INVALID_WORK_DIMENSION
	INVALID_WORK_GROUP_SIZE         ErrorCode = C.CL_INVALID_WORK_GROUP_SIZE
	INVALID_WORK_ITEM_SIZE          ErrorCode = C.CL_INVALID_WORK_ITEM_SIZE
	INVALID_GLOBAL_OFFSET           ErrorCode = C.CL_INVALID_GLOBAL_OFFSET
	INVALID_EVENT_WAIT_LIST         ErrorCode = C.CL_INVALID_EVENT_WAIT_LIST
	INVALID_EVENT                   ErrorCode = C.CL_INVALID_EVENT
	INVALID_OPERATION               ErrorCode = C.CL_INVALID_OPERATION
	INVALID_GL_OBJECT               ErrorCode = C.CL_INVALID_GL_OBJECT
	INVALID_BUFFER_SIZE             ErrorCode = C.CL_INVALID_BUFFER_SIZE
	INVALID_MIP_LEVEL               ErrorCode = C.CL_INVALID_MIP_LEVEL
	INVALID_GLOBAL_WORK_SIZE        ErrorCode = C.CL_INVALID_GLOBAL_WORK_SIZE
)

const (
	VERSION_1_0 = C.CL_VERSION_1_0
)

const (
	FALSE Bool = C.CL_FALSE
	TRUE  Bool = C.CL_TRUE
)

const (
	PLATFORM_PROFILE    PlatformInfo = C.CL_PLATFORM_PROFILE
	PLATFORM_VERSION    PlatformInfo = C.CL_PLATFORM_VERSION
	PLATFORM_NAME       PlatformInfo = C.CL_PLATFORM_NAME
	PLATFORM_VENDOR     PlatformInfo = C.CL_PLATFORM_VENDOR
	PLATFORM_EXTENSIONS PlatformInfo = C.CL_PLATFORM_EXTENSIONS
)

const (
	DEVICE_TYPE_DEFAULT     DeviceType = C.CL_DEVICE_TYPE_DEFAULT
	DEVICE_TYPE_CPU         DeviceType = C.CL_DEVICE_TYPE_CPU
	DEVICE_TYPE_GPU         DeviceType = C.CL_DEVICE_TYPE_GPU
	DEVICE_TYPE_ACCELERATOR DeviceType = C.CL_DEVICE_TYPE_ACCELERATOR
	DEVICE_TYPE_ALL         DeviceType = C.CL_DEVICE_TYPE_ALL
)

const (
	DEVICE_TYPE                          DeviceInfo = C.CL_DEVICE_TYPE
	DEVICE_VENDOR_ID                     DeviceInfo = C.CL_DEVICE_VENDOR_ID
	DEVICE_MAX_COMPUTE_UNITS             DeviceInfo = C.CL_DEVICE_MAX_COMPUTE_UNITS
	DEVICE_MAX_WORK_ITEM_DIMENSIONS      DeviceInfo = C.CL_DEVICE_MAX_WORK_ITEM_DIMENSIONS
	DEVICE_MAX_WORK_GROUP_SIZE           DeviceInfo = C.CL_DEVICE_MAX_WORK_GROUP_SIZE
	DEVICE_MAX_WORK_ITEM_SIZES           DeviceInfo = C.CL_DEVICE_MAX_WORK_ITEM_SIZES
	DEVICE_PREFERRED_VECTOR_WIDTH_CHAR   DeviceInfo = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_CHAR
	DEVICE_PREFERRED_VECTOR_WIDTH_SHORT  DeviceInfo = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_SHORT
	DEVICE_PREFERRED_VECTOR_WIDTH_INT    DeviceInfo = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_INT
	DEVICE_PREFERRED_VECTOR_WIDTH_LONG   DeviceInfo = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_LONG
	DEVICE_PREFERRED_VECTOR_WIDTH_FLOAT  DeviceInfo = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_FLOAT
	DEVICE_PREFERRED_VECTOR_WIDTH_DOUBLE DeviceInfo = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_DOUBLE
	DEVICE_MAX_CLOCK_FREQUENCY           DeviceInfo = C.CL_DEVICE_MAX_CLOCK_FREQUENCY
	DEVICE_ADDRESS_BITS                  DeviceInfo = C.CL_DEVICE_ADDRESS_BITS
	DEVICE_MAX_READ_IMAGE_ARGS           DeviceInfo = C.CL_DEVICE_MAX_READ_IMAGE_ARGS
	DEVICE_MAX_WRITE_IMAGE_ARGS          DeviceInfo = C.CL_DEVICE_MAX_WRITE_IMAGE_ARGS
	DEVICE_MAX_MEM_ALLOC_SIZE            DeviceInfo = C.CL_DEVICE_MAX_MEM_ALLOC_SIZE
	DEVICE_IMAGE2D_MAX_WIDTH             DeviceInfo = C.CL_DEVICE_IMAGE2D_MAX_WIDTH
	DEVICE_IMAGE2D_MAX_HEIGHT            DeviceInfo = C.CL_DEVICE_IMAGE2D_MAX_HEIGHT
	DEVICE_IMAGE3D_MAX_WIDTH             DeviceInfo = C.CL_DEVICE_IMAGE3D_MAX_WIDTH
	DEVICE_IMAGE3D_MAX_HEIGHT            DeviceInfo = C.CL_DEVICE_IMAGE3D_MAX_HEIGHT
	DEVICE_IMAGE3D_MAX_DEPTH             DeviceInfo = C.CL_DEVICE_IMAGE3D_MAX_DEPTH
	DEVICE_IMAGE_SUPPORT                 DeviceInfo = C.CL_DEVICE_IMAGE_SUPPORT
	DEVICE_MAX_PARAMETER_SIZE            DeviceInfo = C.CL_DEVICE_MAX_PARAMETER_SIZE
	DEVICE_MAX_SAMPLERS                  DeviceInfo = C.CL_DEVICE_MAX_SAMPLERS
	DEVICE_MEM_BASE_ADDR_ALIGN           DeviceInfo = C.CL_DEVICE_MEM_BASE_ADDR_ALIGN
	DEVICE_MIN_DATA_TYPE_ALIGN_SIZE      DeviceInfo = C.CL_DEVICE_MIN_DATA_TYPE_ALIGN_SIZE
	DEVICE_SINGLE_FP_CONFIG              DeviceInfo = C.CL_DEVICE_SINGLE_FP_CONFIG
	DEVICE_GLOBAL_MEM_CACHE_TYPE         DeviceInfo = C.CL_DEVICE_GLOBAL_MEM_CACHE_TYPE
	DEVICE_GLOBAL_MEM_CACHELINE_SIZE     DeviceInfo = C.CL_DEVICE_GLOBAL_MEM_CACHELINE_SIZE
	DEVICE_GLOBAL_MEM_CACHE_SIZE         DeviceInfo = C.CL_DEVICE_GLOBAL_MEM_CACHE_SIZE
	DEVICE_GLOBAL_MEM_SIZE               DeviceInfo = C.CL_DEVICE_GLOBAL_MEM_SIZE
	DEVICE_MAX_CONSTANT_BUFFER_SIZE      DeviceInfo = C.CL_DEVICE_MAX_CONSTANT_BUFFER_SIZE
	DEVICE_MAX_CONSTANT_ARGS             DeviceInfo = C.CL_DEVICE_MAX_CONSTANT_ARGS
	DEVICE_LOCAL_MEM_TYPE                DeviceInfo = C.CL_DEVICE_LOCAL_MEM_TYPE
	DEVICE_LOCAL_MEM_SIZE                DeviceInfo = C.CL_DEVICE_LOCAL_MEM_SIZE
	DEVICE_ERROR_CORRECTION_SUPPORT      DeviceInfo = C.CL_DEVICE_ERROR_CORRECTION_SUPPORT
	DEVICE_PROFILING_TIMER_RESOLUTION    DeviceInfo = C.CL_DEVICE_PROFILING_TIMER_RESOLUTION
	DEVICE_ENDIAN_LITTLE                 DeviceInfo = C.CL_DEVICE_ENDIAN_LITTLE
	DEVICE_AVAILABLE                     DeviceInfo = C.CL_DEVICE_AVAILABLE
	DEVICE_COMPILER_AVAILABLE            DeviceInfo = C.CL_DEVICE_COMPILER_AVAILABLE
	DEVICE_EXECUTION_CAPABILITIES        DeviceInfo = C.CL_DEVICE_EXECUTION_CAPABILITIES
	DEVICE_QUEUE_PROPERTIES              DeviceInfo = C.CL_DEVICE_QUEUE_PROPERTIES
	DEVICE_NAME                          DeviceInfo = C.CL_DEVICE_NAME
	DEVICE_VENDOR                        DeviceInfo = C.CL_DEVICE_VENDOR
	DRIVER_VERSION                       DeviceInfo = C.CL_DRIVER_VERSION
	DEVICE_PROFILE                       DeviceInfo = C.CL_DEVICE_PROFILE
	DEVICE_VERSION                       DeviceInfo = C.CL_DEVICE_VERSION
	DEVICE_EXTENSIONS                    DeviceInfo = C.CL_DEVICE_EXTENSIONS
	DEVICE_PLATFORM                      DeviceInfo = C.CL_DEVICE_PLATFORM
)

const (
	FP_DENORM           DeviceFpConfig = C.CL_FP_DENORM
	FP_INF_NAN          DeviceFpConfig = C.CL_FP_INF_NAN
	FP_ROUND_TO_NEAREST DeviceFpConfig = C.CL_FP_ROUND_TO_NEAREST
	FP_ROUND_TO_ZERO    DeviceFpConfig = C.CL_FP_ROUND_TO_ZERO
	FP_ROUND_TO_INF     DeviceFpConfig = C.CL_FP_ROUND_TO_INF
	FP_FMA              DeviceFpConfig = C.CL_FP_FMA
	FP_SOFT_FLOAT       DeviceFpConfig = C.CL_FP_SOFT_FLOAT
)

const (
	NONE             DeviceMemCacheType = C.CL_NONE
	READ_ONLY_CACHE  DeviceMemCacheType = C.CL_READ_ONLY_CACHE
	READ_WRITE_CACHE DeviceMemCacheType = C.CL_READ_WRITE_CACHE
)

const (
	LOCAL  DeviceLocalMemType = C.CL_LOCAL
	GLOBAL DeviceLocalMemType = C.CL_GLOBAL
)

const (
	EXEC_KERNEL        DeviceExecCapabilities = C.CL_EXEC_KERNEL
	EXEC_NATIVE_KERNEL DeviceExecCapabilities = C.CL_EXEC_NATIVE_KERNEL
)

const (
	QUEUE_OUT_OF_ORDER_EXEC_MODE_ENABLE CommandQueueProperties = C.CL_QUEUE_OUT_OF_ORDER_EXEC_MODE_ENABLE
	QUEUE_PROFILING_ENABLE              CommandQueueProperties = C.CL_QUEUE_PROFILING_ENABLE
)

const (
	CONTEXT_REFERENCE_COUNT ContextInfo = C.CL_CONTEXT_REFERENCE_COUNT
	CONTEXT_DEVICES         ContextInfo = C.CL_CONTEXT_DEVICES
	CONTEXT_PROPERTIES      ContextInfo = C.CL_CONTEXT_PROPERTIES
)

const (
	CONTEXT_PLATFORM ContextProperties = C.CL_CONTEXT_PLATFORM
)

const (
	QUEUE_CONTEXT         CommandQueueInfo = C.CL_QUEUE_CONTEXT
	QUEUE_DEVICE          CommandQueueInfo = C.CL_QUEUE_DEVICE
	QUEUE_REFERENCE_COUNT CommandQueueInfo = C.CL_QUEUE_REFERENCE_COUNT
	QUEUE_PROPERTIES      CommandQueueInfo = C.CL_QUEUE_PROPERTIES
)

const (
	MEM_READ_WRITE     MemFlags = C.CL_MEM_READ_WRITE
	MEM_WRITE_ONLY     MemFlags = C.CL_MEM_WRITE_ONLY
	MEM_READ_ONLY      MemFlags = C.CL_MEM_READ_ONLY
	MEM_USE_HOST_PTR   MemFlags = C.CL_MEM_USE_HOST_PTR
	MEM_ALLOC_HOST_PTR MemFlags = C.CL_MEM_ALLOC_HOST_PTR
	MEM_COPY_HOST_PTR  MemFlags = C.CL_MEM_COPY_HOST_PTR
)

const (
	R         ChannelOrder = C.CL_R
	A         ChannelOrder = C.CL_A
	RG        ChannelOrder = C.CL_RG
	RA        ChannelOrder = C.CL_RA
	RGB       ChannelOrder = C.CL_RGB
	RGBA      ChannelOrder = C.CL_RGBA
	BGRA      ChannelOrder = C.CL_BGRA
	ARGB      ChannelOrder = C.CL_ARGB
	INTENSITY ChannelOrder = C.CL_INTENSITY
	LUMINANCE ChannelOrder = C.CL_LUMINANCE
)

const (
	SNORM_INT8       ChannelType = C.CL_SNORM_INT8
	SNORM_INT16      ChannelType = C.CL_SNORM_INT16
	UNORM_INT8       ChannelType = C.CL_UNORM_INT8
	UNORM_INT16      ChannelType = C.CL_UNORM_INT16
	UNORM_SHORT_565  ChannelType = C.CL_UNORM_SHORT_565
	UNORM_SHORT_555  ChannelType = C.CL_UNORM_SHORT_555
	UNORM_INT_101010 ChannelType = C.CL_UNORM_INT_101010
	SIGNED_INT8      ChannelType = C.CL_SIGNED_INT8
	SIGNED_INT16     ChannelType = C.CL_SIGNED_INT16
	SIGNED_INT32     ChannelType = C.CL_SIGNED_INT32
	UNSIGNED_INT8    ChannelType = C.CL_UNSIGNED_INT8
	UNSIGNED_INT16   ChannelType = C.CL_UNSIGNED_INT16
	UNSIGNED_INT32   ChannelType = C.CL_UNSIGNED_INT32
	HALF_FLOAT       ChannelType = C.CL_HALF_FLOAT
	FLOAT            ChannelType = C.CL_FLOAT
)

const (
	MEM_OBJECT_BUFFER  MemObjectType = C.CL_MEM_OBJECT_BUFFER
	MEM_OBJECT_IMAGE2D MemObjectType = C.CL_MEM_OBJECT_IMAGE2D
	MEM_OBJECT_IMAGE3D MemObjectType = C.CL_MEM_OBJECT_IMAGE3D
)

const (
	MEM_TYPE            MemInfo = C.CL_MEM_TYPE
	MEM_FLAGS           MemInfo = C.CL_MEM_FLAGS
	MEM_SIZE            MemInfo = C.CL_MEM_SIZE
	MEM_HOST_PTR        MemInfo = C.CL_MEM_HOST_PTR
	MEM_MAP_COUNT       MemInfo = C.CL_MEM_MAP_COUNT
	MEM_REFERENCE_COUNT MemInfo = C.CL_MEM_REFERENCE_COUNT
	MEM_CONTEXT         MemInfo = C.CL_MEM_CONTEXT
)

const (
	IMAGE_FORMAT       ImageInfo = C.CL_IMAGE_FORMAT
	IMAGE_ELEMENT_SIZE ImageInfo = C.CL_IMAGE_ELEMENT_SIZE
	IMAGE_ROW_PITCH    ImageInfo = C.CL_IMAGE_ROW_PITCH
	IMAGE_SLICE_PITCH  ImageInfo = C.CL_IMAGE_SLICE_PITCH
	IMAGE_WIDTH        ImageInfo = C.CL_IMAGE_WIDTH
	IMAGE_HEIGHT       ImageInfo = C.CL_IMAGE_HEIGHT
	IMAGE_DEPTH        ImageInfo = C.CL_IMAGE_DEPTH
)

const (
	ADDRESS_NONE          AddressingMode = C.CL_ADDRESS_NONE
	ADDRESS_CLAMP_TO_EDGE AddressingMode = C.CL_ADDRESS_CLAMP_TO_EDGE
	ADDRESS_CLAMP         AddressingMode = C.CL_ADDRESS_CLAMP
	ADDRESS_REPEAT        AddressingMode = C.CL_ADDRESS_REPEAT
)

const (
	FILTER_NEAREST FilterMode = C.CL_FILTER_NEAREST
	FILTER_LINEAR  FilterMode = C.CL_FILTER_LINEAR
)

const (
	SAMPLER_REFERENCE_COUNT   SamplerInfo = C.CL_SAMPLER_REFERENCE_COUNT
	SAMPLER_CONTEXT           SamplerInfo = C.CL_SAMPLER_CONTEXT
	SAMPLER_NORMALIZED_COORDS SamplerInfo = C.CL_SAMPLER_NORMALIZED_COORDS
	SAMPLER_ADDRESSING_MODE   SamplerInfo = C.CL_SAMPLER_ADDRESSING_MODE
	SAMPLER_FILTER_MODE       SamplerInfo = C.CL_SAMPLER_FILTER_MODE
)

const (
	MAP_READ  MapFlags = C.CL_MAP_READ
	MAP_WRITE MapFlags = C.CL_MAP_WRITE
)

const (
	PROGRAM_REFERENCE_COUNT ProgramInfo = C.CL_PROGRAM_REFERENCE_COUNT
	PROGRAM_CONTEXT         ProgramInfo = C.CL_PROGRAM_CONTEXT
	PROGRAM_NUM_DEVICES     ProgramInfo = C.CL_PROGRAM_NUM_DEVICES
	PROGRAM_DEVICES         ProgramInfo = C.CL_PROGRAM_DEVICES
	PROGRAM_SOURCE          ProgramInfo = C.CL_PROGRAM_SOURCE
	PROGRAM_BINARY_SIZES    ProgramInfo = C.CL_PROGRAM_BINARY_SIZES
	PROGRAM_BINARIES        ProgramInfo = C.CL_PROGRAM_BINARIES
)

const (
	PROGRAM_BUILD_STATUS  ProgramBuildInfo = C.CL_PROGRAM_BUILD_STATUS
	PROGRAM_BUILD_OPTIONS ProgramBuildInfo = C.CL_PROGRAM_BUILD_OPTIONS
	PROGRAM_BUILD_LOG     ProgramBuildInfo = C.CL_PROGRAM_BUILD_LOG
)

const (
	BUILD_SUCCESS     BuildStatus = C.CL_BUILD_SUCCESS
	BUILD_NONE        BuildStatus = C.CL_BUILD_NONE
	BUILD_ERROR       BuildStatus = C.CL_BUILD_ERROR
	BUILD_IN_PROGRESS BuildStatus = C.CL_BUILD_IN_PROGRESS
)

const (
	KERNEL_FUNCTION_NAME   KernelInfo = C.CL_KERNEL_FUNCTION_NAME
	KERNEL_NUM_ARGS        KernelInfo = C.CL_KERNEL_NUM_ARGS
	KERNEL_REFERENCE_COUNT KernelInfo = C.CL_KERNEL_REFERENCE_COUNT
	KERNEL_CONTEXT         KernelInfo = C.CL_KERNEL_CONTEXT
	KERNEL_PROGRAM         KernelInfo = C.CL_KERNEL_PROGRAM
)

const (
	KERNEL_WORK_GROUP_SIZE         KernelWorkGroupInfo = C.CL_KERNEL_WORK_GROUP_SIZE
	KERNEL_COMPILE_WORK_GROUP_SIZE KernelWorkGroupInfo = C.CL_KERNEL_COMPILE_WORK_GROUP_SIZE
	KERNEL_LOCAL_MEM_SIZE          KernelWorkGroupInfo = C.CL_KERNEL_LOCAL_MEM_SIZE
)

const (
	EVENT_COMMAND_QUEUE            EventInfo = C.CL_EVENT_COMMAND_QUEUE
	EVENT_COMMAND_TYPE             EventInfo = C.CL_EVENT_COMMAND_TYPE
	EVENT_REFERENCE_COUNT          EventInfo = C.CL_EVENT_REFERENCE_COUNT
	EVENT_COMMAND_EXECUTION_STATUS EventInfo = C.CL_EVENT_COMMAND_EXECUTION_STATUS
)

const (
	COMMAND_NDRANGE_KERNEL       CommandType = C.CL_COMMAND_NDRANGE_KERNEL
	COMMAND_TASK                 CommandType = C.CL_COMMAND_TASK
	COMMAND_NATIVE_KERNEL        CommandType = C.CL_COMMAND_NATIVE_KERNEL
	COMMAND_READ_BUFFER          CommandType = C.CL_COMMAND_READ_BUFFER
	COMMAND_WRITE_BUFFER         CommandType = C.CL_COMMAND_WRITE_BUFFER
	COMMAND_COPY_BUFFER          CommandType = C.CL_COMMAND_COPY_BUFFER
	COMMAND_READ_IMAGE           CommandType = C.CL_COMMAND_READ_IMAGE
	COMMAND_WRITE_IMAGE          CommandType = C.CL_COMMAND_WRITE_IMAGE
	COMMAND_COPY_IMAGE           CommandType = C.CL_COMMAND_COPY_IMAGE
	COMMAND_COPY_IMAGE_TO_BUFFER CommandType = C.CL_COMMAND_COPY_IMAGE_TO_BUFFER
	COMMAND_COPY_BUFFER_TO_IMAGE CommandType = C.CL_COMMAND_COPY_BUFFER_TO_IMAGE
	COMMAND_MAP_BUFFER           CommandType = C.CL_COMMAND_MAP_BUFFER
	COMMAND_MAP_IMAGE            CommandType = C.CL_COMMAND_MAP_IMAGE
	COMMAND_UNMAP_MEM_OBJECT     CommandType = C.CL_COMMAND_UNMAP_MEM_OBJECT
	COMMAND_MARKER               CommandType = C.CL_COMMAND_MARKER
	COMMAND_ACQUIRE_GL_OBJECTS   CommandType = C.CL_COMMAND_ACQUIRE_GL_OBJECTS
	COMMAND_RELEASE_GL_OBJECTS   CommandType = C.CL_COMMAND_RELEASE_GL_OBJECTS
)

const (
	COMPLETE  CommandExecutionStatus = C.CL_COMPLETE
	RUNNING   CommandExecutionStatus = C.CL_RUNNING
	SUBMITTED CommandExecutionStatus = C.CL_SUBMITTED
	QUEUED    CommandExecutionStatus = C.CL_QUEUED
)

const (
	PROFILING_COMMAND_QUEUED ProfilingInfo = C.CL_PROFILING_COMMAND_QUEUED
	PROFILING_COMMAND_SUBMIT ProfilingInfo = C.CL_PROFILING_COMMAND_SUBMIT
	PROFILING_COMMAND_START  ProfilingInfo = C.CL_PROFILING_COMMAND_START
	PROFILING_COMMAND_END    ProfilingInfo = C.CL_PROFILING_COMMAND_END
)

// -----------------------------------------------------------------------------
// Platform Api

type PlatformID C.cl_platform_id

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetPlatformIDs.html
func GetPlatformIDs(numentries uint32, ids *PlatformID, numplatform *uint32) ErrorCode {
	return ErrorCode(C.clGetPlatformIDs(C.cl_uint(numentries), (*C.cl_platform_id)(unsafe.Pointer(ids)), (*C.cl_uint)(numplatform)))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetPlatformInfo.html
func GetPlatformInfo(pid PlatformID, paramName PlatformInfo, paramValueSize uint64, data unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetPlatformInfo(pid, C.cl_platform_info(paramName), C.size_t(paramValueSize), data, (*C.size_t)(paramValueSizeRet)))
}

// -----------------------------------------------------------------------------
// Device Api

type DeviceId C.cl_device_id

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetDeviceIDs.html
func GetDeviceIDs(pid PlatformID, deviceType DeviceType, numentries uint32, devices *DeviceId, numdevices *uint32) ErrorCode {
	return ErrorCode(C.clGetDeviceIDs(pid, C.cl_device_type(deviceType), C.cl_uint(numentries), (*C.cl_device_id)(unsafe.Pointer(devices)), (*C.cl_uint)(numdevices)))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetDeviceInfo.html
func GetDeviceInfo(did DeviceId, paramName DeviceInfo, paramValueSize uint64, data unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetDeviceInfo(did, C.cl_device_info(paramName), C.size_t(paramValueSize), data, (*C.size_t)(paramValueSizeRet)))
}

// -----------------------------------------------------------------------------
// Context Api

type gocontext struct {
	clContext   C.cl_context
	errCallback func(string, unsafe.Pointer, uint64, interface{})
	userdata    interface{}
}

type Context *gocontext

//export contextErrorCallback
func contextErrorCallback(errinfo *C.char, privateinfo unsafe.Pointer, cb C.size_t, userData unsafe.Pointer) {
	ctx := (*gocontext)(userData)
	ctx.errCallback(C.GoString(errinfo), privateinfo, uint64(cb), ctx.userdata)
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateContext.html
func CreateContext(properties *ContextProperties, numDevices uint32, devices *DeviceId, notify func(string, unsafe.Pointer, uint64, interface{}), userdata interface{}, errcode *ErrorCode) Context {
	ctx := gocontext{nil, notify, userdata}
	var f *[0]byte
	var u unsafe.Pointer
	if notify != nil {
		f = (*[0]byte)(C.contextErrorCallback)
		u = unsafe.Pointer(&ctx)
	}
	ctx.clContext = C.clCreateContext((*C.cl_context_properties)(unsafe.Pointer(properties)), C.cl_uint(numDevices), (*C.cl_device_id)(unsafe.Pointer(devices)), f, u, (*C.cl_int)(unsafe.Pointer(errcode)))
	return Context(&ctx)
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateContextFromType.html
func CreateContextFromType(properties *ContextProperties, deviceType DeviceType, notify func(string, unsafe.Pointer, uint64, interface{}), userdata interface{}, errcode *ErrorCode) Context {
	ctx := gocontext{nil, notify, userdata}
	var f *[0]byte
	var u unsafe.Pointer
	if notify != nil {
		f = (*[0]byte)(C.contextErrorCallback)
		u = unsafe.Pointer(&ctx)
	}
	ctx.clContext = C.clCreateContextFromType((*C.cl_context_properties)(unsafe.Pointer(properties)), C.cl_device_type(deviceType), f, u, (*C.cl_int)(unsafe.Pointer(errcode)))
	return Context(&ctx)
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clRetainContext.html
func RetainContext(context Context) ErrorCode {
	return ErrorCode(C.clRetainContext(context.clContext))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clReleaseContext.html
func ReleaseContext(context Context) ErrorCode {
	return ErrorCode(C.clReleaseContext(context.clContext))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetContextInfo.html
func GetContextInfo(context Context, paramName ContextInfo, paramValueSize uint64, data unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetContextInfo(context.clContext, C.cl_context_info(paramName), C.size_t(paramValueSize), data, (*C.size_t)(paramValueSizeRet)))
}

// -----------------------------------------------------------------------------
// Command queue Api

type CommandQueue C.cl_command_queue

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateCommandQueue.html
func CreateCommandQueue(context Context, did DeviceId, properties CommandQueueProperties, errcode *ErrorCode) CommandQueue {
	return CommandQueue(C.clCreateCommandQueue(context.clContext, did, C.cl_command_queue_properties(properties), (*C.cl_int)(errcode)))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clRetainCommandQueue.html
func RetainCommandQueue(cq CommandQueue) {
	C.clRetainCommandQueue(cq)
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clReleaseCommandQueue.html
func ReleaseCommandQueue(cq CommandQueue) {
	C.clReleaseCommandQueue(cq)
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetCommandQueueInfo.html
func GetCommandQueueInfo(cq CommandQueue, paramName CommandQueueInfo, paramValueSize uint64, data unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetCommandQueueInfo(cq, C.cl_command_queue_info(paramName), C.size_t(paramValueSize), data, (*C.size_t)(paramValueSizeRet)))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clSetCommandQueueProperty.html
func SetCommandQueueProperty(cq CommandQueue, properties CommandQueueProperties, enable BOOL, oldproperties *CommandQueueProperties) int {
	return ErrorCode(C.clSetCommandQueueProperty(cq, C.cl_command_queue_properties(properties), C.cl_bool(enable), (*C.cl_command_queue_properties)(oldproperties)))
}

// -----------------------------------------------------------------------------
// Memory Object Api

type ImageFormat struct {
	imageChannelOrder    ChannelOrder
	imageChannelDataType ChannelType
}

func (imf ImageFormat) toC() *C.cl_image_format {
	return &C.cl_image_format{
		image_channel_order:     C.cl_channel_order(imf.imageChannelOrder),
		image_channel_data_type: C.cl_channel_type(imf.imageChannelDataType)}
}

type ImageDesc struct {
	imageType                                                                           MemObjectType
	numMipLevels, numSamples                                                            uint32
	imageWidth, imageHeight, imageDepth, imageArraySize, imageRowPitch, imageSlicePitch uint64
	buffer                                                                              Mem
}

// see (https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/climde.html
func (imde ImageDesc) toC() *C.cl_image_desc {
	return &C.cl_image_desc{
		image_type:        C.cl_mem_object_type(imde.imageType),
		image_width:       C.size_t(imde.imageWidth),
		image_height:      C.size_t(imde.imageHeight),
		image_depth:       C.size_t(imde.imageDepth),
		image_array_size:  C.size_t(imde.imageArraySize),
		image_row_pitch:   C.size_t(imde.imageRowPitch),
		image_slice_pitch: C.size_t(imde.imageSlicePitch),
		num_mip_levels:    C.cl_uint(imde.numMipLevels),
		num_samples:       C.cl_uint(imde.numSamples),
		buffer:            C.cl_mem(imde.buffer)}
}

type Mem C.cl_mem

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateBuffer.html
func CreateBuffer(context Context, flags, paramValueSize MemFlags, hostPtr unsafe.Pointer, errcode *ErrorCode) Mem {
	return Mem(C.clCreateBuffer(context.clContext, C.cl_mem_flags(flags), C.size_t(paramValueSize), hostPtr, (*C.cl_int)(unsafe.Pointer(errcode))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateImage2D.html
func CreateImage2D(context Context, flags MemFlags, imageFormat ImageFormat, imageWidth, imageHeight, imageRowPitch uint64, hostptr unsafe.Pointer, errcode *ErrorCode) Mem {
	return Mem(C.clCreateImage2D(context.clContext, C.cl_mem_flags(flags), imageFormat.toC(), C.size_t(imageWidth), C.size_t(imageHeight), C.size_t(imageRowPitch), hostptr, (*C.cl_int)(unsafe.Pointer(errcode))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateImage3D.html
func CreateImage3D(context Context, flags MemFlags, imageFormat ImageFormat, imageWidth, imageHeight, imageDepth, imageRowPitch, imageSlicePitch uint64, hostptr unsafe.Pointer, errcode *ErrorCode) Mem {
	return Mem(C.clCreateImage3D(context.clContext, C.cl_mem_flags(flags), imageFormat.toC(), C.size_t(imageWidth), C.size_t(imageHeight), C.size_t(imageDepth), C.size_t(imageRowPitch), C.size_t(imageSlicePitch), hostptr, (*C.cl_int)(unsafe.Pointer(errcode))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clRetainMemObject.html
func RetainMemObject(mem Mem) ErrorCode {
	return ErrorCode(C.clRetainMemObject(mem))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clReleaseMemObject.html
func ReleaseMemObject(mem Mem) ErrorCode {
	return ErrorCode(C.clReleaseMemObject(mem))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetSupportedImageFormats.html
func GetSupportedImageFormats(context Context, flags MemFlags, memObjectType MemObjectType, numEntries uint32, imageformats *ImageFormat, numImageFormat *uint32) ErrorCode {
	return ErrorCode(C.clGetSupportedImageFormats(context.clContext,
		C.cl_mem_flags(flags),
		C.cl_mem_object_type(memObjectType),
		C.cl_uint(numEntries),
		(*C.cl_image_format)(unsafe.Pointer(imageformats)),
		(*C.cl_uint)(numImageFormat)))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetMemObjectInfo.html
func GetMemObjectInfo(mem Mem, paramName MemInfo, paramValueSize uint64, paramValue unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetMemObjectInfo(mem, C.cl_mem_info(paramName), C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetImageInfo.html
func GetImageInfo(mem Mem, paramName ImageInfo, paramValueSize uint64, paramValue unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetImageInfo(mem, C.cl_image_info(paramName), C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}

// -----------------------------------------------------------------------------
// Sampler Api

type Sampler C.cl_sampler

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateSampler.html
func CreateSampler(context Context, normalizedCoords Bool, addressingMode AddressingMode, filterMode FilterMode, errcode *ErrorCode) Sampler {
	return Sampler(C.clCreateSampler(context.clContext, C.cl_bool(normalizedCoords), C.cl_addressing_mode(addressingMode), C.cl_filter_mode(filterMode), (*C.cl_int)(unsafe.Pointer(errcode))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clRetainSampler.html
func RetainSampler(sampler Sampler) ErrorCode {
	return ErrorCode(C.clRetainSampler(sampler))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clReleaseSampler.html
func ReleaseSampler(sampler Sampler) ErrorCode {
	return ErrorCode(C.clReleaseSampler(sampler))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetSamplerInfo.html
func GetSamplerInfo(sampler Sampler, paramName SamplerInfo, paramValueSize uint64, paramValue unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetSamplerInfo(sampler, C.cl_sampler_info(paramName), C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}

// -----------------------------------------------------------------------------
// Program Object Api

type Program C.cl_program

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateProgramWithSource.html
func CreateProgramWithSource(context Context, count uint32, src **uint8, lengths *uint64, errcode *ErrorCode) Program {
	return Program(C.clCreateProgramWithSource(context.clContext, C.cl_uint(count), (**C.char)(unsafe.Pointer(src)), (*C.size_t)(unsafe.Pointer(lengths)), (*C.cl_int)(unsafe.Pointer(errcode))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateProgramWithBinary.html
func CreateProgramWithBinary(context Context, numDevices uint32, devices *DeviceId, lengths *uint64, binaries **uint8, binaryStatus *int32, errcode *ErrorCode) Program {
	return Program(C.clCreateProgramWithBinary(context.clContext,
		C.cl_uint(numDevices),
		(*C.cl_device_id)(unsafe.Pointer(devices)),
		(*C.size_t)(unsafe.Pointer(lengths)),
		(**C.uchar)(unsafe.Pointer(binaries)),
		(*C.cl_int)(unsafe.Pointer(binaryStatus)),
		(*C.cl_int)(unsafe.Pointer(errcode))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clRetainProgram.html
func RetainProgram(prog Program) ErrorCode {
	return ErrorCode(C.clRetainProgram(prog))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clReleaseProgram.html
func ReleaseProgram(prog Program) ErrorCode {
	return ErrorCode(C.clReleaseProgram(prog))
}

type programCallbackHolder struct {
	cbfunc   func(Program, interface{})
	userdata interface{}
}

var pochHolder map[*programCallbackHolder]struct{}

//export programCallback
func programCallback(prog C.cl_program, userdata unsafe.Pointer) {
	poch := (*programCallbackHolder)(userdata)
	poch.cbfunc(Program(prog), poch.userdata)
	delete(pochHolder, poch)
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clBuildProgram.html
func BuildProgram(prog Program, numDevices uint32, devices *DeviceId, options *uint8, notify func(Program, interface{}), userdata interface{}) ErrorCode {
	var f *[0]byte
	var u unsafe.Pointer
	if notify != nil {
		pobch := programCallbackHolder{notify, userdata}
		pochHolder[&pobch] = struct{}{}
		f = (*[0]byte)(C.programCallback)
		u = unsafe.Pointer(&pobch)
	}
	return ErrorCode(C.clBuildProgram(prog, C.cl_uint(numDevices), (*C.cl_device_id)(unsafe.Pointer(devices)), (*C.char)(unsafe.Pointer(options)), f, u))
}

func UnloadCompiler() ErrorCode {
	return ErrorCode(C.clUnloadCompiler())
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetProgramInfo.html
func GetProgramInfo(prog Program, paramName ProgramInfo, paramValueSize uint64, paramValue unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetProgramInfo(prog, C.cl_program_info(paramName), C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetProgramBuildInfo.html
func GetProgramBuildInfo(prog Program, device DeviceId, paramName ProgramBuildInfo, paramValueSize uint64, paramValue unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetProgramBuildInfo(prog, device, C.cl_program_build_info(paramName), C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}

// -----------------------------------------------------------------------------
// Kernel Object Api

type Kernel C.cl_kernel

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateKernel.html
func CreateKernel(prog Program, kernelName *uint8, errcode *ErrorCode) Kernel {
	return Kernel(C.clCreateKernel(prog, (*C.char)(unsafe.Pointer(kernelName)), (*C.cl_int)(unsafe.Pointer(errcode))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateKernelsInProgram.html
func CreateKernelsInProgram(prog Program, numKernels uint32, kernels *Kernel, numKernelsRet *uint32) ErrorCode {
	return ErrorCode(C.clCreateKernelsInProgram(prog, C.cl_uint(numKernels), (*C.cl_kernel)(unsafe.Pointer(kernels)), (*C.cl_uint)(unsafe.Pointer(numKernelsRet))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clRetainKernel.html
func RetainKernel(ker Kernel) ErrorCode {
	return ErrorCode(C.clRetainKernel(ker))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clReleaseKernel.html
func ReleaseKernel(ker Kernel) ErrorCode {
	return ErrorCode(C.clReleaseKernel(ker))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clSetKernelArg.html
func SetKernelArg(ker Kernel, argIndex uint32, argSize uint64, argValue unsafe.Pointer) ErrorCode {
	return ErrorCode(C.clSetKernelArg(ker, C.cl_uint(argIndex), C.size_t(argSize), argValue))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetKernelInfo.html
func GetKernelInfo(ker Kernel, paramName KernelInfo, paramValueSize uint64, paramValue unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetKernelInfo(ker, C.cl_kernel_info(paramName), C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetKernelWorkGroupInfo.html
func GetKernelWorkGroupInfo(ker Kernel, did DeviceId, paramName KernelWorkGroupInfo, paramValueSize uint64, paramValue unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetKernelWorkGroupInfo(ker, did, C.cl_kernel_work_group_info(paramName), C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}

// -----------------------------------------------------------------------------
// Event Object Api

type Event struct {
	clEvent C.cl_event
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clWaitForEvents.html
func WaitForEvents(numEvents uint32, events *Event) ErrorCode {
	return ErrorCode(C.clWaitForEvents(C.cl_uint(numEvents), (*C.cl_event)(unsafe.Pointer(events))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetEventInfo.html
func GetEventInfo(e Event, paramName uint32, paramValueSize uint64, paramValue unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetEventInfo(e.clEvent, C.cl_event_info(paramName), C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clRetainEvent.html
func RetainEvent(e Event) ErrorCode {
	return ErrorCode(C.clRetainEvent(e.clEvent))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clReleaseEvent.html
func ReleaseEvent(e Event) ErrorCode {
	return ErrorCode(C.clReleaseEvent(e.clEvent))
}

// -----------------------------------------------------------------------------
// Profiling Api

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetEventProfilingInfo.html
func GetEventProfilingInfo(e Event, paramName ProfilingInfo, paramValueSize uint64, paramValue unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetEventProfilingInfo(e.clEvent, C.cl_profiling_info(paramName), C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}

// -----------------------------------------------------------------------------
// Flush and Finish Api

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clFlush.html
func Flush(cq CommandQueue) ErrorCode {
	return ErrorCode(C.clFlush(cq))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clFinish.html
func Finish(cq CommandQueue) ErrorCode {
	return ErrorCode(C.clFinish(cq))
}

// -----------------------------------------------------------------------------
// Enqueue Api

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueReadBuffer.html
func EnqueueReadBuffer(cq CommandQueue, buffer Mem, blocking_read Bool, offset uint64, size uint64, ptr unsafe.Pointer, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueReadBuffer(cq, buffer, C.cl_bool(blocking_read), C.size_t(offset), C.size_t(size), unsafe.Pointer(ptr), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueWriteBuffer.html
func EnqueueWriteBuffer(cq CommandQueue, buffer Mem, blocking_write Bool, offset uint64, size uint64, ptr unsafe.Pointer, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueWriteBuffer(cq, buffer, C.cl_bool(blocking_write), C.size_t(offset), C.size_t(size), unsafe.Pointer(ptr), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueCopyBuffer.html
func EnqueueCopyBuffer(cq CommandQueue, src_buffer Mem, dst_buffer Mem, src_offset uint64, dst_offset uint64, size uint64, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueCopyBuffer(cq, src_buffer, dst_buffer, C.size_t(src_offset), C.size_t(dst_offset), C.size_t(size), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueReadImage.html
func EnqueueReadImage(cq CommandQueue, image Mem, blocking_read Bool, origin3 *uint64, region3 *uint64, row_pitch uint64, slice_pitch uint64, ptr unsafe.Pointer, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueReadImage(cq, image, C.cl_bool(blocking_read), (*C.size_t)(unsafe.Pointer(origin3)), (*C.size_t)(unsafe.Pointer(region3)), C.size_t(row_pitch), C.size_t(slice_pitch), unsafe.Pointer(ptr), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueWriteImage.html
func EnqueueWriteImage(cq CommandQueue, image Mem, blocking_write Bool, origin3 *uint64, region3 *uint64, input_row_pitch uint64, input_slice_pitch uint64, ptr unsafe.Pointer, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueWriteImage(cq, image, C.cl_bool(blocking_write), (*C.size_t)(unsafe.Pointer(origin3)), (*C.size_t)(unsafe.Pointer(region3)), C.size_t(input_row_pitch), C.size_t(input_slice_pitch), unsafe.Pointer(ptr), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueCopyImage.html
func EnqueueCopyImage(cq CommandQueue, src_image Mem, dst_image Mem, src_origin3 *uint64, dst_origin3 *uint64, region3 *uint64, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueCopyImage(cq, src_image, dst_image, (*C.size_t)(unsafe.Pointer(src_origin3)), (*C.size_t)(unsafe.Pointer(dst_origin3)), (*C.size_t)(unsafe.Pointer(region3)), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueCopyImageToBuffer.html
func EnqueueCopyImageToBuffer(cq CommandQueue, src_image Mem, dst_buffer Mem, src_origin3 *uint64, region3 *uint64, dst_offset uint64, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueCopyImageToBuffer(cq, src_image, dst_buffer, (*C.size_t)(unsafe.Pointer(src_origin3)), (*C.size_t)(unsafe.Pointer(region3)), C.size_t(dst_offset), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueCopyBufferToImage.html
func EnqueueCopyBufferToImage(cq CommandQueue, src_buffer Mem, dst_image Mem, src_offset uint64, dst_origin3 *uint64, region3 *uint64, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueCopyBufferToImage(cq, src_buffer, dst_image, C.size_t(src_offset), (*C.size_t)(unsafe.Pointer(dst_origin3)), (*C.size_t)(unsafe.Pointer(region3)), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueMapBuffer.html
func EnqueueMapBuffer(cq CommandQueue, buffer Mem, blocking_map Bool, map_flags uint64, offset uint64, size uint64, num_events_in_wait_list uint32, event_wait_list *Event, event *Event, errcode_ret *int32) unsafe.Pointer {
	return unsafe.Pointer(C.clEnqueueMapBuffer(cq, buffer, C.cl_bool(blocking_map), C.cl_map_flags(map_flags), C.size_t(offset), C.size_t(size), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event)), (*C.cl_int)(unsafe.Pointer(errcode_ret))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueMapImage.html
func EnqueueMapImage(cq CommandQueue, image Mem, blocking_map Bool, map_flags MapFlags, origin3 *uint64, region3 *uint64, image_row_pitch *uint64, image_slice_pitch *uint64, num_events_in_wait_list uint32, event_wait_list *Event, event *Event, errcode_ret *int32) unsafe.Pointer {
	return unsafe.Pointer(C.clEnqueueMapImage(cq, image, C.cl_bool(blocking_map), C.cl_map_flags(map_flags), (*C.size_t)(unsafe.Pointer(origin3)), (*C.size_t)(unsafe.Pointer(region3)), (*C.size_t)(unsafe.Pointer(image_row_pitch)), (*C.size_t)(unsafe.Pointer(image_slice_pitch)), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event)), (*C.cl_int)(unsafe.Pointer(errcode_ret))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueUnmapMemObject.html
func EnqueueUnmapMemObject(cq CommandQueue, memobj Mem, mapped_ptr unsafe.Pointer, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueUnmapMemObject(cq, memobj, unsafe.Pointer(mapped_ptr), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueNDRangeKernel.html
func EnqueueNDRangeKernel(cq CommandQueue, kernel Kernel, work_dim uint32, global_work_offset *uint64, global_work_size *uint64, local_work_size *uint64, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueNDRangeKernel(cq, kernel, C.cl_uint(work_dim), (*C.size_t)(unsafe.Pointer(global_work_offset)), (*C.size_t)(unsafe.Pointer(global_work_size)), (*C.size_t)(unsafe.Pointer(local_work_size)), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueTask.html
func EnqueueTask(cq CommandQueue, kernel Kernel, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueTask(cq, kernel, C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

/*
// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueNativeKernel.html
func EnqueueNativeKernel(cq CommandQueue, userfunc func(unsafe.Pointer), args unsafe.Pointer, cb_args uint64, num_mem_objects uint32, mem_list *Mem, args_mem_loc unsafe.Pointer, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) int32 {
	C.clEnqueueNativeKernel(cp)
}*/

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueMarker.html
func EnqueueMarker(cq CommandQueue, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueMarker(cq, (*C.cl_event)(unsafe.Pointer(event))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueWaitForEvents.html
func EnqueueWaitForEvents(cq CommandQueue, num_events uint32, event_list *Event) ErrorCode {
	return ErrorCode(C.clEnqueueWaitForEvents(cq, C.cl_uint(num_events), (*C.cl_event)(unsafe.Pointer(event_list))))
}

// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueBarrier.html
func EnqueueBarrier(cq CommandQueue) ErrorCode {
	return ErrorCode(C.clEnqueueBarrier(cq))
}

// -----------------------------------------------------------------------------
// Extension functions Api

func GetExtensionFunctionAddress(funcName string) unsafe.Pointer {
	return unsafe.Pointer(C.clGetExtensionFunctionAddress((*C.char)(unsafe.Pointer(str(funcName)))))
}

func str(str string) *uint8 {
	if !strings.HasSuffix(str, "\x00") {
		log.Fatal("str argument missing null terminator", str)
	}
	header := (*reflect.StringHeader)(unsafe.Pointer(&str))
	return (*uint8)(unsafe.Pointer(header.Data))
}
