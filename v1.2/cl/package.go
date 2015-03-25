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
extern void memObjectDestroyCallback(cl_mem,void*);
extern void programObjectBuildCompleteCallback(cl_program, void*);
extern void programObjectCompileCompleteCallback(cl_program, void*);
*/
import "C"
import (
	"unsafe"
)

func init() {
	mochHolder = make(map[*memObjectCallbackHolder]struct{})
	pobchHolder = make(map[*programObjectBuildCompleteHolder]struct{})
	poccHolder = make(map[*programObjectCompileCompleteHolder]struct{})
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
type BuildStatus uint32
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

const ( /* Error Codes */
	SUCCESS                                   ErrorCode = C.CL_SUCCESS
	DEVICE_NOT_FOUND                                    = C.CL_DEVICE_NOT_FOUND
	DEVICE_NOT_AVAILABLE                                = C.CL_DEVICE_NOT_AVAILABLE
	COMPILER_NOT_AVAILABLE                              = C.CL_COMPILER_NOT_AVAILABLE
	MEM_OBJECT_ALLOCATION_FAILURE                       = C.CL_MEM_OBJECT_ALLOCATION_FAILURE
	OUT_OF_RESOURCES                                    = C.CL_OUT_OF_RESOURCES
	OUT_OF_HOST_MEMORY                                  = C.CL_OUT_OF_HOST_MEMORY
	PROFILING_INFO_NOT_AVAILABLE                        = C.CL_PROFILING_INFO_NOT_AVAILABLE
	MEM_COPY_OVERLAP                                    = C.CL_MEM_COPY_OVERLAP
	IMAGE_FORMAT_MISMATCH                               = C.CL_IMAGE_FORMAT_MISMATCH
	IMAGE_FORMAT_NOT_SUPPORTED                          = C.CL_IMAGE_FORMAT_NOT_SUPPORTED
	BUILD_PROGRAM_FAILURE                               = C.CL_BUILD_PROGRAM_FAILURE
	MAP_FAILURE                                         = C.CL_MAP_FAILURE
	MISALIGNED_SUB_BUFFER_OFFSET                        = C.CL_MISALIGNED_SUB_BUFFER_OFFSET
	EXEC_STATUS_ERROR_FOR_EVENTS_IN_WAIT_LIST           = C.CL_EXEC_STATUS_ERROR_FOR_EVENTS_IN_WAIT_LIST
	COMPILE_PROGRAM_FAILURE                             = C.CL_COMPILE_PROGRAM_FAILURE
	LINKER_NOT_AVAILABLE                                = C.CL_LINKER_NOT_AVAILABLE
	LINK_PROGRAM_FAILURE                                = C.CL_LINK_PROGRAM_FAILURE
	DEVICE_PARTITION_FAILED                             = C.CL_DEVICE_PARTITION_FAILED
	KERNEL_ARG_INFO_NOT_AVAILABLE                       = C.CL_KERNEL_ARG_INFO_NOT_AVAILABLE
	INVALID_VALUE                                       = C.CL_INVALID_VALUE
	INVALID_DEVICE_TYPE                                 = C.CL_INVALID_DEVICE_TYPE
	INVALID_PLATFORM                                    = C.CL_INVALID_PLATFORM
	INVALID_DEVICE                                      = C.CL_INVALID_DEVICE
	INVALID_CONTEXT                                     = C.CL_INVALID_CONTEXT
	INVALID_QUEUE_PROPERTIES                            = C.CL_INVALID_QUEUE_PROPERTIES
	INVALID_COMMAND_QUEUE                               = C.CL_INVALID_COMMAND_QUEUE
	INVALID_HOST_PTR                                    = C.CL_INVALID_HOST_PTR
	INVALID_MEM_OBJECT                                  = C.CL_INVALID_MEM_OBJECT
	INVALID_IMAGE_FORMAT_DESCRIPTOR                     = C.CL_INVALID_IMAGE_FORMAT_DESCRIPTOR
	INVALID_IMAGE_SIZE                                  = C.CL_INVALID_IMAGE_SIZE
	INVALID_SAMPLER                                     = C.CL_INVALID_SAMPLER
	INVALID_BINARY                                      = C.CL_INVALID_BINARY
	INVALID_BUILD_OPTIONS                               = C.CL_INVALID_BUILD_OPTIONS
	INVALID_PROGRAM                                     = C.CL_INVALID_PROGRAM
	INVALID_PROGRAM_EXECUTABLE                          = C.CL_INVALID_PROGRAM_EXECUTABLE
	INVALID_KERNEL_NAME                                 = C.CL_INVALID_KERNEL_NAME
	INVALID_KERNEL_DEFINITION                           = C.CL_INVALID_KERNEL_DEFINITION
	INVALID_KERNEL                                      = C.CL_INVALID_KERNEL
	INVALID_ARG_INDEX                                   = C.CL_INVALID_ARG_INDEX
	INVALID_ARG_VALUE                                   = C.CL_INVALID_ARG_VALUE
	INVALID_ARG_SIZE                                    = C.CL_INVALID_ARG_SIZE
	INVALID_KERNEL_ARGS                                 = C.CL_INVALID_KERNEL_ARGS
	INVALID_WORK_DIMENSION                              = C.CL_INVALID_WORK_DIMENSION
	INVALID_WORK_GROUP_SIZE                             = C.CL_INVALID_WORK_GROUP_SIZE
	INVALID_WORK_ITEM_SIZE                              = C.CL_INVALID_WORK_ITEM_SIZE
	INVALID_GLOBAL_OFFSET                               = C.CL_INVALID_GLOBAL_OFFSET
	INVALID_EVENT_WAIT_LIST                             = C.CL_INVALID_EVENT_WAIT_LIST
	INVALID_EVENT                                       = C.CL_INVALID_EVENT
	INVALID_OPERATION                                   = C.CL_INVALID_OPERATION
	INVALID_GL_OBJECT                                   = C.CL_INVALID_GL_OBJECT
	INVALID_BUFFER_SIZE                                 = C.CL_INVALID_BUFFER_SIZE
	INVALID_MIP_LEVEL                                   = C.CL_INVALID_MIP_LEVEL
	INVALID_GLOBAL_WORK_SIZE                            = C.CL_INVALID_GLOBAL_WORK_SIZE
	INVALID_PROPERTY                                    = C.CL_INVALID_PROPERTY
	INVALID_IMAGE_DESCRIPTOR                            = C.CL_INVALID_IMAGE_DESCRIPTOR
	INVALID_COMPILER_OPTIONS                            = C.CL_INVALID_COMPILER_OPTIONS
	INVALID_LINKER_OPTIONS                              = C.CL_INVALID_LINKER_OPTIONS
	INVALID_DEVICE_PARTITION_COUNT                      = C.CL_INVALID_DEVICE_PARTITION_COUNT
)

const ( /* OpenCL Version */
	VERSION_1_0 = C.CL_VERSION_1_0
	VERSION_1_1 = C.CL_VERSION_1_1
	VERSION_1_2 = C.CL_VERSION_1_2
)

const (
	FALSE        Bool = C.CL_FALSE
	TRUE              = C.CL_TRUE
	BLOCKING          = C.CL_BLOCKING
	NON_BLOCKING      = C.CL_NON_BLOCKING
)

const (
	PLATFORM_PROFILE    PlatformInfo = C.CL_PLATFORM_PROFILE
	PLATFORM_VERSION                 = C.CL_PLATFORM_VERSION
	PLATFORM_NAME                    = C.CL_PLATFORM_NAME
	PLATFORM_VENDOR                  = C.CL_PLATFORM_VENDOR
	PLATFORM_EXTENSIONS              = C.CL_PLATFORM_EXTENSIONS
)

const (
	DEVICE_TYPE_DEFAULT     DeviceType = C.CL_DEVICE_TYPE_DEFAULT
	DEVICE_TYPE_CPU                    = C.CL_DEVICE_TYPE_CPU
	DEVICE_TYPE_GPU                    = C.CL_DEVICE_TYPE_GPU
	DEVICE_TYPE_ACCELERATOR            = C.CL_DEVICE_TYPE_ACCELERATOR
	DEVICE_TYPE_CUSTOM                 = C.CL_DEVICE_TYPE_CUSTOM
	DEVICE_TYPE_ALL                    = C.CL_DEVICE_TYPE_ALL
)

const (
	DEVICE_TYPE                          DeviceInfo = C.CL_DEVICE_TYPE
	DEVICE_VENDOR_ID                                = C.CL_DEVICE_VENDOR_ID
	DEVICE_MAX_COMPUTE_UNITS                        = C.CL_DEVICE_MAX_COMPUTE_UNITS
	DEVICE_MAX_WORK_ITEM_DIMENSIONS                 = C.CL_DEVICE_MAX_WORK_ITEM_DIMENSIONS
	DEVICE_MAX_WORK_GROUP_SIZE                      = C.CL_DEVICE_MAX_WORK_GROUP_SIZE
	DEVICE_MAX_WORK_ITEM_SIZES                      = C.CL_DEVICE_MAX_WORK_ITEM_SIZES
	DEVICE_PREFERRED_VECTOR_WIDTH_CHAR              = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_CHAR
	DEVICE_PREFERRED_VECTOR_WIDTH_SHORT             = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_SHORT
	DEVICE_PREFERRED_VECTOR_WIDTH_INT               = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_INT
	DEVICE_PREFERRED_VECTOR_WIDTH_LONG              = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_LONG
	DEVICE_PREFERRED_VECTOR_WIDTH_FLOAT             = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_FLOAT
	DEVICE_PREFERRED_VECTOR_WIDTH_DOUBLE            = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_DOUBLE
	DEVICE_MAX_CLOCK_FREQUENCY                      = C.CL_DEVICE_MAX_CLOCK_FREQUENCY
	DEVICE_ADDRESS_BITS                             = C.CL_DEVICE_ADDRESS_BITS
	DEVICE_MAX_READ_IMAGE_ARGS                      = C.CL_DEVICE_MAX_READ_IMAGE_ARGS
	DEVICE_MAX_WRITE_IMAGE_ARGS                     = C.CL_DEVICE_MAX_WRITE_IMAGE_ARGS
	DEVICE_MAX_MEM_ALLOC_SIZE                       = C.CL_DEVICE_MAX_MEM_ALLOC_SIZE
	DEVICE_IMAGE2D_MAX_WIDTH                        = C.CL_DEVICE_IMAGE2D_MAX_WIDTH
	DEVICE_IMAGE2D_MAX_HEIGHT                       = C.CL_DEVICE_IMAGE2D_MAX_HEIGHT
	DEVICE_IMAGE3D_MAX_WIDTH                        = C.CL_DEVICE_IMAGE3D_MAX_WIDTH
	DEVICE_IMAGE3D_MAX_HEIGHT                       = C.CL_DEVICE_IMAGE3D_MAX_HEIGHT
	DEVICE_IMAGE3D_MAX_DEPTH                        = C.CL_DEVICE_IMAGE3D_MAX_DEPTH
	DEVICE_IMAGE_SUPPORT                            = C.CL_DEVICE_IMAGE_SUPPORT
	DEVICE_MAX_PARAMETER_SIZE                       = C.CL_DEVICE_MAX_PARAMETER_SIZE
	DEVICE_MAX_SAMPLERS                             = C.CL_DEVICE_MAX_SAMPLERS
	DEVICE_MEM_BASE_ADDR_ALIGN                      = C.CL_DEVICE_MEM_BASE_ADDR_ALIGN
	DEVICE_MIN_DATA_TYPE_ALIGN_SIZE                 = C.CL_DEVICE_MIN_DATA_TYPE_ALIGN_SIZE
	DEVICE_SINGLE_FP_CONFIG                         = C.CL_DEVICE_SINGLE_FP_CONFIG
	DEVICE_GLOBAL_MEM_CACHE_TYPE                    = C.CL_DEVICE_GLOBAL_MEM_CACHE_TYPE
	DEVICE_GLOBAL_MEM_CACHELINE_SIZE                = C.CL_DEVICE_GLOBAL_MEM_CACHELINE_SIZE
	DEVICE_GLOBAL_MEM_CACHE_SIZE                    = C.CL_DEVICE_GLOBAL_MEM_CACHE_SIZE
	DEVICE_GLOBAL_MEM_SIZE                          = C.CL_DEVICE_GLOBAL_MEM_SIZE
	DEVICE_MAX_CONSTANT_BUFFER_SIZE                 = C.CL_DEVICE_MAX_CONSTANT_BUFFER_SIZE
	DEVICE_MAX_CONSTANT_ARGS                        = C.CL_DEVICE_MAX_CONSTANT_ARGS
	DEVICE_LOCAL_MEM_TYPE                           = C.CL_DEVICE_LOCAL_MEM_TYPE
	DEVICE_LOCAL_MEM_SIZE                           = C.CL_DEVICE_LOCAL_MEM_SIZE
	DEVICE_ERROR_CORRECTION_SUPPORT                 = C.CL_DEVICE_ERROR_CORRECTION_SUPPORT
	DEVICE_PROFILING_TIMER_RESOLUTION               = C.CL_DEVICE_PROFILING_TIMER_RESOLUTION
	DEVICE_ENDIAN_LITTLE                            = C.CL_DEVICE_ENDIAN_LITTLE
	DEVICE_AVAILABLE                                = C.CL_DEVICE_AVAILABLE
	DEVICE_COMPILER_AVAILABLE                       = C.CL_DEVICE_COMPILER_AVAILABLE
	DEVICE_EXECUTION_CAPABILITIES                   = C.CL_DEVICE_EXECUTION_CAPABILITIES
	DEVICE_QUEUE_PROPERTIES                         = C.CL_DEVICE_QUEUE_PROPERTIES
	DEVICE_NAME                                     = C.CL_DEVICE_NAME
	DEVICE_VENDOR                                   = C.CL_DEVICE_VENDOR
	DRIVER_VERSION                                  = C.CL_DRIVER_VERSION
	DEVICE_PROFILE                                  = C.CL_DEVICE_PROFILE
	DEVICE_VERSION                                  = C.CL_DEVICE_VERSION
	DEVICE_EXTENSIONS                               = C.CL_DEVICE_EXTENSIONS
	DEVICE_PLATFORM                                 = C.CL_DEVICE_PLATFORM
	DEVICE_DOUBLE_FP_CONFIG                         = C.CL_DEVICE_DOUBLE_FP_CONFIG
)

const ( /* 0x1033 reserved for CL_DEVICE_HALF_FP_CONFIG */
	DEVICE_PREFERRED_VECTOR_WIDTH_HALF  = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_HALF
	DEVICE_HOST_UNIFIED_MEMORY          = C.CL_DEVICE_HOST_UNIFIED_MEMORY
	DEVICE_NATIVE_VECTOR_WIDTH_CHAR     = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_CHAR
	DEVICE_NATIVE_VECTOR_WIDTH_SHORT    = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_SHORT
	DEVICE_NATIVE_VECTOR_WIDTH_INT      = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_INT
	DEVICE_NATIVE_VECTOR_WIDTH_LONG     = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_LONG
	DEVICE_NATIVE_VECTOR_WIDTH_FLOAT    = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_FLOAT
	DEVICE_NATIVE_VECTOR_WIDTH_DOUBLE   = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_DOUBLE
	DEVICE_NATIVE_VECTOR_WIDTH_HALF     = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_HALF
	DEVICE_OPENCL_C_VERSION             = C.CL_DEVICE_OPENCL_C_VERSION
	DEVICE_LINKER_AVAILABLE             = C.CL_DEVICE_LINKER_AVAILABLE
	DEVICE_BUILT_IN_KERNELS             = C.CL_DEVICE_BUILT_IN_KERNELS
	DEVICE_IMAGE_MAX_BUFFER_SIZE        = C.CL_DEVICE_IMAGE_MAX_BUFFER_SIZE
	DEVICE_IMAGE_MAX_ARRAY_SIZE         = C.CL_DEVICE_IMAGE_MAX_ARRAY_SIZE
	DEVICE_PARENT_DEVICE                = C.CL_DEVICE_PARENT_DEVICE
	DEVICE_PARTITION_MAX_SUB_DEVICES    = C.CL_DEVICE_PARTITION_MAX_SUB_DEVICES
	DEVICE_PARTITION_PROPERTIES         = C.CL_DEVICE_PARTITION_PROPERTIES
	DEVICE_PARTITION_AFFINITY_DOMAIN    = C.CL_DEVICE_PARTITION_AFFINITY_DOMAIN
	DEVICE_PARTITION_TYPE               = C.CL_DEVICE_PARTITION_TYPE
	DEVICE_REFERENCE_COUNT              = C.CL_DEVICE_REFERENCE_COUNT
	DEVICE_PREFERRED_INTEROP_USER_SYNC  = C.CL_DEVICE_PREFERRED_INTEROP_USER_SYNC
	DEVICE_PRINTF_BUFFER_SIZE           = C.CL_DEVICE_PRINTF_BUFFER_SIZE
	DEVICE_IMAGE_PITCH_ALIGNMENT        = C.CL_DEVICE_IMAGE_PITCH_ALIGNMENT
	DEVICE_IMAGE_BASE_ADDRESS_ALIGNMENT = C.CL_DEVICE_IMAGE_BASE_ADDRESS_ALIGNMENT
)

const (
	FP_DENORM                        DeviceFpConfig = C.CL_FP_DENORM
	FP_INF_NAN                                      = C.CL_FP_INF_NAN
	FP_ROUND_TO_NEAREST                             = C.CL_FP_ROUND_TO_NEAREST
	FP_ROUND_TO_ZERO                                = C.CL_FP_ROUND_TO_ZERO
	FP_ROUND_TO_INF                                 = C.CL_FP_ROUND_TO_INF
	FP_FMA                                          = C.CL_FP_FMA
	FP_SOFT_FLOAT                                   = C.CL_FP_SOFT_FLOAT
	FP_CORRECTLY_ROUNDED_DIVIDE_SQRT                = C.CL_FP_CORRECTLY_ROUNDED_DIVIDE_SQRT
)

const (
	NONE             DeviceMemCacheType = C.CL_NONE
	READ_ONLY_CACHE                     = C.CL_READ_ONLY_CACHE
	READ_WRITE_CACHE                    = C.CL_READ_WRITE_CACHE
)

const (
	LOCAL  DeviceLocalMemType = C.CL_LOCAL
	GLOBAL                    = C.CL_GLOBAL
)

const (
	EXEC_KERNEL        DeviceExecCapabilities = C.CL_EXEC_KERNEL
	EXEC_NATIVE_KERNEL                        = C.CL_EXEC_NATIVE_KERNEL
)

const (
	QUEUE_OUT_OF_ORDER_EXEC_MODE_ENABLE CommandQueueProperties = C.CL_QUEUE_OUT_OF_ORDER_EXEC_MODE_ENABLE
	QUEUE_PROFILING_ENABLE                                     = C.CL_QUEUE_PROFILING_ENABLE
)

const (
	CONTEXT_REFERENCE_COUNT ContextInfo = C.CL_CONTEXT_REFERENCE_COUNT
	CONTEXT_DEVICES                     = C.CL_CONTEXT_DEVICES
	CONTEXT_PROPERTIES                  = C.CL_CONTEXT_PROPERTIES
	CONTEXT_NUM_DEVICES                 = C.CL_CONTEXT_NUM_DEVICES
)

const (
	CONTEXT_PLATFORM          ContextProperties = C.CL_CONTEXT_PLATFORM
	CONTEXT_INTEROP_USER_SYNC                   = C.CL_CONTEXT_INTEROP_USER_SYNC
)

const (
	DEVICE_PARTITION_EQUALLY            DevicePartitionProperty = C.CL_DEVICE_PARTITION_EQUALLY
	DEVICE_PARTITION_BY_COUNTS                                  = C.CL_DEVICE_PARTITION_BY_COUNTS
	DEVICE_PARTITION_BY_COUNTS_LIST_END                         = C.CL_DEVICE_PARTITION_BY_COUNTS_LIST_END
	DEVICE_PARTITION_BY_AFFINITY_DOMAIN                         = C.CL_DEVICE_PARTITION_BY_AFFINITY_DOMAIN
)

const (
	DEVICE_AFFINITY_DOMAIN_NUMA               DeviceAffinityDomain = C.CL_DEVICE_AFFINITY_DOMAIN_NUMA
	DEVICE_AFFINITY_DOMAIN_L4_CACHE                                = C.CL_DEVICE_AFFINITY_DOMAIN_L4_CACHE
	DEVICE_AFFINITY_DOMAIN_L3_CACHE                                = C.CL_DEVICE_AFFINITY_DOMAIN_L3_CACHE
	DEVICE_AFFINITY_DOMAIN_L2_CACHE                                = C.CL_DEVICE_AFFINITY_DOMAIN_L2_CACHE
	DEVICE_AFFINITY_DOMAIN_L1_CACHE                                = C.CL_DEVICE_AFFINITY_DOMAIN_L1_CACHE
	DEVICE_AFFINITY_DOMAIN_NEXT_PARTITIONABLE                      = C.CL_DEVICE_AFFINITY_DOMAIN_NEXT_PARTITIONABLE
)

const (
	QUEUE_CONTEXT         CommandQueueInfo = C.CL_QUEUE_CONTEXT
	QUEUE_DEVICE                           = C.CL_QUEUE_DEVICE
	QUEUE_REFERENCE_COUNT                  = C.CL_QUEUE_REFERENCE_COUNT
	QUEUE_PROPERTIES                       = C.CL_QUEUE_PROPERTIES
)

const (
	MEM_READ_WRITE     MemFlags = C.CL_MEM_READ_WRITE
	MEM_WRITE_ONLY              = C.CL_MEM_WRITE_ONLY
	MEM_READ_ONLY               = C.CL_MEM_READ_ONLY
	MEM_USE_HOST_PTR            = C.CL_MEM_USE_HOST_PTR
	MEM_ALLOC_HOST_PTR          = C.CL_MEM_ALLOC_HOST_PTR
	MEM_COPY_HOST_PTR           = C.CL_MEM_COPY_HOST_PTR
	// reserved                                         (1 << 6)
	MEM_HOST_WRITE_ONLY = C.CL_MEM_HOST_WRITE_ONLY
	MEM_HOST_READ_ONLY  = C.CL_MEM_HOST_READ_ONLY
	MEM_HOST_NO_ACCESS  = C.CL_MEM_HOST_NO_ACCESS
)

const (
	MIGRATE_MEM_OBJECT_HOST              MemMigrationFlags = C.CL_MIGRATE_MEM_OBJECT_HOST
	MIGRATE_MEM_OBJECT_CONTENT_UNDEFINED                   = C.CL_MIGRATE_MEM_OBJECT_CONTENT_UNDEFINED
)

const (
	R             ChannelOrder = C.CL_R
	A                          = C.CL_A
	RG                         = C.CL_RG
	RA                         = C.CL_RA
	RGB                        = C.CL_RGB
	RGBA                       = C.CL_RGBA
	BGRA                       = C.CL_BGRA
	ARGB                       = C.CL_ARGB
	INTENSITY                  = C.CL_INTENSITY
	LUMINANCE                  = C.CL_LUMINANCE
	Rx                         = C.CL_Rx
	RGx                        = C.CL_RGx
	RGBx                       = C.CL_RGBx
	DEPTH                      = C.CL_DEPTH
	DEPTH_STENCIL              = C.CL_DEPTH_STENCIL
)

const (
	SNORM_INT8       ChannelType = C.CL_SNORM_INT8
	SNORM_INT16                  = C.CL_SNORM_INT16
	UNORM_INT8                   = C.CL_UNORM_INT8
	UNORM_INT16                  = C.CL_UNORM_INT16
	UNORM_SHORT_565              = C.CL_UNORM_SHORT_565
	UNORM_SHORT_555              = C.CL_UNORM_SHORT_555
	UNORM_INT_101010             = C.CL_UNORM_INT_101010
	SIGNED_INT8                  = C.CL_SIGNED_INT8
	SIGNED_INT16                 = C.CL_SIGNED_INT16
	SIGNED_INT32                 = C.CL_SIGNED_INT32
	UNSIGNED_INT8                = C.CL_UNSIGNED_INT8
	UNSIGNED_INT16               = C.CL_UNSIGNED_INT16
	UNSIGNED_INT32               = C.CL_UNSIGNED_INT32
	HALF_FLOAT                   = C.CL_HALF_FLOAT
	FLOAT                        = C.CL_FLOAT
	UNORM_INT24                  = C.CL_UNORM_INT24
)

const (
	MEM_OBJECT_BUFFER         MemObjectType = C.CL_MEM_OBJECT_BUFFER
	MEM_OBJECT_IMAGE2D                      = C.CL_MEM_OBJECT_IMAGE2D
	MEM_OBJECT_IMAGE3D                      = C.CL_MEM_OBJECT_IMAGE3D
	MEM_OBJECT_IMAGE2D_ARRAY                = C.CL_MEM_OBJECT_IMAGE2D_ARRAY
	MEM_OBJECT_IMAGE1D                      = C.CL_MEM_OBJECT_IMAGE1D
	MEM_OBJECT_IMAGE1D_ARRAY                = C.CL_MEM_OBJECT_IMAGE1D_ARRAY
	MEM_OBJECT_IMAGE1D_BUFFER               = C.CL_MEM_OBJECT_IMAGE1D_BUFFER
)

const (
	MEM_TYPE                 MemInfo = C.CL_MEM_TYPE
	MEM_FLAGS                        = C.CL_MEM_FLAGS
	MEM_SIZE                         = C.CL_MEM_SIZE
	MEM_HOST_PTR                     = C.CL_MEM_HOST_PTR
	MEM_MAP_COUNT                    = C.CL_MEM_MAP_COUNT
	MEM_REFERENCE_COUNT              = C.CL_MEM_REFERENCE_COUNT
	MEM_CONTEXT                      = C.CL_MEM_CONTEXT
	MEM_ASSOCIATED_MEMOBJECT         = C.CL_MEM_ASSOCIATED_MEMOBJECT
	MEM_OFFSET                       = C.CL_MEM_OFFSET
)

const (
	IMAGE_FORMAT         ImageInfo = C.CL_IMAGE_FORMAT
	IMAGE_ELEMENT_SIZE             = C.CL_IMAGE_ELEMENT_SIZE
	IMAGE_ROW_PITCH                = C.CL_IMAGE_ROW_PITCH
	IMAGE_SLICE_PITCH              = C.CL_IMAGE_SLICE_PITCH
	IMAGE_WIDTH                    = C.CL_IMAGE_WIDTH
	IMAGE_HEIGHT                   = C.CL_IMAGE_HEIGHT
	IMAGE_DEPTH                    = C.CL_IMAGE_DEPTH
	IMAGE_ARRAY_SIZE               = C.CL_IMAGE_ARRAY_SIZE
	IMAGE_BUFFER                   = C.CL_IMAGE_BUFFER
	IMAGE_NUM_MIP_LEVELS           = C.CL_IMAGE_NUM_MIP_LEVELS
	IMAGE_NUM_SAMPLES              = C.CL_IMAGE_NUM_SAMPLES
)

const (
	ADDRESS_NONE            AddressingMode = C.CL_ADDRESS_NONE
	ADDRESS_CLAMP_TO_EDGE                  = C.CL_ADDRESS_CLAMP_TO_EDGE
	ADDRESS_CLAMP                          = C.CL_ADDRESS_CLAMP
	ADDRESS_REPEAT                         = C.CL_ADDRESS_REPEAT
	ADDRESS_MIRRORED_REPEAT                = C.CL_ADDRESS_MIRRORED_REPEAT
)

const (
	FILTER_NEAREST FilterMode = C.CL_FILTER_NEAREST
	FILTER_LINEAR             = C.CL_FILTER_LINEAR
)

const (
	SAMPLER_REFERENCE_COUNT   SamplerInfo = C.CL_SAMPLER_REFERENCE_COUNT
	SAMPLER_CONTEXT                       = C.CL_SAMPLER_CONTEXT
	SAMPLER_NORMALIZED_COORDS             = C.CL_SAMPLER_NORMALIZED_COORDS
	SAMPLER_ADDRESSING_MODE               = C.CL_SAMPLER_ADDRESSING_MODE
	SAMPLER_FILTER_MODE                   = C.CL_SAMPLER_FILTER_MODE
)

const (
	MAP_READ                    MapFlags = C.CL_MAP_READ
	MAP_WRITE                            = C.CL_MAP_WRITE
	MAP_WRITE_INVALIDATE_REGION          = C.CL_MAP_WRITE_INVALIDATE_REGION
)

const (
	PROGRAM_REFERENCE_COUNT ProgramInfo = C.CL_PROGRAM_REFERENCE_COUNT
	PROGRAM_CONTEXT                     = C.CL_PROGRAM_CONTEXT
	PROGRAM_NUM_DEVICES                 = C.CL_PROGRAM_NUM_DEVICES
	PROGRAM_DEVICES                     = C.CL_PROGRAM_DEVICES
	PROGRAM_SOURCE                      = C.CL_PROGRAM_SOURCE
	PROGRAM_BINARY_SIZES                = C.CL_PROGRAM_BINARY_SIZES
	PROGRAM_BINARIES                    = C.CL_PROGRAM_BINARIES
	PROGRAM_NUM_KERNELS                 = C.CL_PROGRAM_NUM_KERNELS
	PROGRAM_KERNEL_NAMES                = C.CL_PROGRAM_KERNEL_NAMES
)

const (
	PROGRAM_BUILD_STATUS  ProgramBuildInfo = C.CL_PROGRAM_BUILD_STATUS
	PROGRAM_BUILD_OPTIONS                  = C.CL_PROGRAM_BUILD_OPTIONS
	PROGRAM_BUILD_LOG                      = C.CL_PROGRAM_BUILD_LOG
	PROGRAM_BINARY_TYPE                    = C.CL_PROGRAM_BINARY_TYPE
)

const (
	PROGRAM_BINARY_TYPE_NONE            ProgramBinaryType = C.CL_PROGRAM_BINARY_TYPE_NONE
	PROGRAM_BINARY_TYPE_COMPILED_OBJECT                   = C.CL_PROGRAM_BINARY_TYPE_COMPILED_OBJECT
	PROGRAM_BINARY_TYPE_LIBRARY                           = C.CL_PROGRAM_BINARY_TYPE_LIBRARY
	PROGRAM_BINARY_TYPE_EXECUTABLE                        = C.CL_PROGRAM_BINARY_TYPE_EXECUTABLE
)

const (
	BUILD_SUCCESS     BuildStatus = C.CL_BUILD_SUCCESS
	BUILD_NONE                    = C.CL_BUILD_NONE
	BUILD_ERROR                   = C.CL_BUILD_ERROR
	BUILD_IN_PROGRESS             = C.CL_BUILD_IN_PROGRESS
)

const (
	KERNEL_FUNCTION_NAME   KernelInfo = C.CL_KERNEL_FUNCTION_NAME
	KERNEL_NUM_ARGS                   = C.CL_KERNEL_NUM_ARGS
	KERNEL_REFERENCE_COUNT            = C.CL_KERNEL_REFERENCE_COUNT
	KERNEL_CONTEXT                    = C.CL_KERNEL_CONTEXT
	KERNEL_PROGRAM                    = C.CL_KERNEL_PROGRAM
	KERNEL_ATTRIBUTES                 = C.CL_KERNEL_ATTRIBUTES
)

const (
	KERNEL_ARG_ADDRESS_QUALIFIER KernelArgInfo = C.CL_KERNEL_ARG_ADDRESS_QUALIFIER
	KERNEL_ARG_ACCESS_QUALIFIER                = C.CL_KERNEL_ARG_ACCESS_QUALIFIER
	KERNEL_ARG_TYPE_NAME                       = C.CL_KERNEL_ARG_TYPE_NAME
	KERNEL_ARG_TYPE_QUALIFIER                  = C.CL_KERNEL_ARG_TYPE_QUALIFIER
	KERNEL_ARG_NAME                            = C.CL_KERNEL_ARG_NAME
)

const (
	KERNEL_ARG_ADDRESS_GLOBAL   KernelArgAddressQualifier = C.CL_KERNEL_ARG_ADDRESS_GLOBAL
	KERNEL_ARG_ADDRESS_LOCAL                              = C.CL_KERNEL_ARG_ADDRESS_LOCAL
	KERNEL_ARG_ADDRESS_CONSTANT                           = C.CL_KERNEL_ARG_ADDRESS_CONSTANT
	KERNEL_ARG_ADDRESS_PRIVATE                            = C.CL_KERNEL_ARG_ADDRESS_PRIVATE
)

const (
	KERNEL_ARG_ACCESS_READ_ONLY  KernelArgAccessQualifier = C.CL_KERNEL_ARG_ACCESS_READ_ONLY
	KERNEL_ARG_ACCESS_WRITE_ONLY                          = C.CL_KERNEL_ARG_ACCESS_WRITE_ONLY
	KERNEL_ARG_ACCESS_READ_WRITE                          = C.CL_KERNEL_ARG_ACCESS_READ_WRITE
	KERNEL_ARG_ACCESS_NONE                                = C.CL_KERNEL_ARG_ACCESS_NONE
)

const (
	KERNEL_ARG_TYPE_NONE     KernelArgTypeQualifer = C.CL_KERNEL_ARG_TYPE_NONE
	KERNEL_ARG_TYPE_CONST                          = C.CL_KERNEL_ARG_TYPE_CONST
	KERNEL_ARG_TYPE_RESTRICT                       = C.CL_KERNEL_ARG_TYPE_RESTRICT
	KERNEL_ARG_TYPE_VOLATILE                       = C.CL_KERNEL_ARG_TYPE_VOLATILE
)

const (
	KERNEL_WORK_GROUP_SIZE                    KernelWorkGroupInfo = C.CL_KERNEL_WORK_GROUP_SIZE
	KERNEL_COMPILE_WORK_GROUP_SIZE                                = C.CL_KERNEL_COMPILE_WORK_GROUP_SIZE
	KERNEL_LOCAL_MEM_SIZE                                         = C.CL_KERNEL_LOCAL_MEM_SIZE
	KERNEL_PREFERRED_WORK_GROUP_SIZE_MULTIPLE                     = C.CL_KERNEL_PREFERRED_WORK_GROUP_SIZE_MULTIPLE
	KERNEL_PRIVATE_MEM_SIZE                                       = C.CL_KERNEL_PRIVATE_MEM_SIZE
	KERNEL_GLOBAL_WORK_SIZE                                       = C.CL_KERNEL_GLOBAL_WORK_SIZE
)

const (
	EVENT_COMMAND_QUEUE            EventInfo = C.CL_EVENT_COMMAND_QUEUE
	EVENT_COMMAND_TYPE                       = C.CL_EVENT_COMMAND_TYPE
	EVENT_REFERENCE_COUNT                    = C.CL_EVENT_REFERENCE_COUNT
	EVENT_COMMAND_EXECUTION_STATUS           = C.CL_EVENT_COMMAND_EXECUTION_STATUS
	EVENT_CONTEXT                            = C.CL_EVENT_CONTEXT
)

const (
	COMMAND_NDRANGE_KERNEL       CommandType = C.CL_COMMAND_NDRANGE_KERNEL
	COMMAND_TASK                             = C.CL_COMMAND_TASK
	COMMAND_NATIVE_KERNEL                    = C.CL_COMMAND_NATIVE_KERNEL
	COMMAND_READ_BUFFER                      = C.CL_COMMAND_READ_BUFFER
	COMMAND_WRITE_BUFFER                     = C.CL_COMMAND_WRITE_BUFFER
	COMMAND_COPY_BUFFER                      = C.CL_COMMAND_COPY_BUFFER
	COMMAND_READ_IMAGE                       = C.CL_COMMAND_READ_IMAGE
	COMMAND_WRITE_IMAGE                      = C.CL_COMMAND_WRITE_IMAGE
	COMMAND_COPY_IMAGE                       = C.CL_COMMAND_COPY_IMAGE
	COMMAND_COPY_IMAGE_TO_BUFFER             = C.CL_COMMAND_COPY_IMAGE_TO_BUFFER
	COMMAND_COPY_BUFFER_TO_IMAGE             = C.CL_COMMAND_COPY_BUFFER_TO_IMAGE
	COMMAND_MAP_BUFFER                       = C.CL_COMMAND_MAP_BUFFER
	COMMAND_MAP_IMAGE                        = C.CL_COMMAND_MAP_IMAGE
	COMMAND_UNMAP_MEM_OBJECT                 = C.CL_COMMAND_UNMAP_MEM_OBJECT
	COMMAND_MARKER                           = C.CL_COMMAND_MARKER
	COMMAND_ACQUIRE_GL_OBJECTS               = C.CL_COMMAND_ACQUIRE_GL_OBJECTS
	COMMAND_RELEASE_GL_OBJECTS               = C.CL_COMMAND_RELEASE_GL_OBJECTS
	COMMAND_READ_BUFFER_RECT                 = C.CL_COMMAND_READ_BUFFER_RECT
	COMMAND_WRITE_BUFFER_RECT                = C.CL_COMMAND_WRITE_BUFFER_RECT
	COMMAND_COPY_BUFFER_RECT                 = C.CL_COMMAND_COPY_BUFFER_RECT
	COMMAND_USER                             = C.CL_COMMAND_USER
	COMMAND_BARRIER                          = C.CL_COMMAND_BARRIER
	COMMAND_MIGRATE_MEM_OBJECTS              = C.CL_COMMAND_MIGRATE_MEM_OBJECTS
	COMMAND_FILL_BUFFER                      = C.CL_COMMAND_FILL_BUFFER
	COMMAND_FILL_IMAGE                       = C.CL_COMMAND_FILL_IMAGE
)

const ( /* command execution status */
	COMPLETE  = C.CL_COMPLETE
	RUNNING   = C.CL_RUNNING
	SUBMITTED = C.CL_SUBMITTED
	QUEUED    = C.CL_QUEUED
)

const (
	BUFFER_CREATE_TYPE_REGION BufferCreateType = C.CL_BUFFER_CREATE_TYPE_REGION
)

const (
	PROFILING_COMMAND_QUEUED ProfilingInfo = C.CL_PROFILING_COMMAND_QUEUED
	PROFILING_COMMAND_SUBMIT               = C.CL_PROFILING_COMMAND_SUBMIT
	PROFILING_COMMAND_START                = C.CL_PROFILING_COMMAND_START
	PROFILING_COMMAND_END                  = C.CL_PROFILING_COMMAND_END
)

/*=================================================================================================================
=======================================================Platform Api================================================
=================================================================================================================*/

type PlatformID C.cl_platform_id

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clGetPlatformIDs.html
func GetPlatformIDs(numentries uint32, ids *PlatformID, numplatform *uint32) ErrorCode {
	return ErrorCode(C.clGetPlatformIDs(C.cl_uint(numentries), (*C.cl_platform_id)(unsafe.Pointer(ids)), (*C.cl_uint)(numplatform)))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clGetPlatformInfo.html
func GetPlatformInfo(pid PlatformID, paramName PlatformInfo, paramValueSize uint64, data unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetPlatformInfo(pid, C.cl_platform_info(paramName), C.size_t(paramValueSize), data, (*C.size_t)(paramValueSizeRet)))
}

/*=================================================================================================================
========================================================Device Api=================================================
=================================================================================================================*/

type DeviceId C.cl_device_id

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clGetDeviceIDs.html
func GetDeviceIDs(pid PlatformID, deviceType DeviceType, numentries uint32, devices *DeviceId, numdevices *uint32) ErrorCode {
	return ErrorCode(C.clGetDeviceIDs(pid, C.cl_device_type(deviceType), C.cl_uint(numentries), (*C.cl_device_id)(unsafe.Pointer(devices)), (*C.cl_uint)(numdevices)))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clGetDeviceInfo.html
func GetDeviceInfo(did DeviceId, paramName DeviceInfo, paramValueSize uint64, data unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetDeviceInfo(did, C.cl_device_info(paramName), C.size_t(paramValueSize), data, (*C.size_t)(paramValueSizeRet)))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clCreateSubDevices.html
func CreateSubDevices(did DeviceId, properties *DevicePartitionProperty, numDevices uint32, devices *DeviceId, numDevicesRet *uint32) ErrorCode {
	return ErrorCode(C.clCreateSubDevices(did, (*C.cl_device_partition_property)(unsafe.Pointer(properties)), C.cl_uint(numDevices), (*C.cl_device_id)(unsafe.Pointer(devices)), (*C.cl_uint)(numDevicesRet)))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clRetainDevice.html
func RetainDevice(did DeviceId) ErrorCode {
	return ErrorCode(C.clRetainDevice(did))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clReleaseDevice.html
func ReleaseDevice(did DeviceId) ErrorCode {
	return ErrorCode(C.clReleaseDevice(did))
}

/*=================================================================================================================
========================================================Context Api================================================
=================================================================================================================*/

type Context struct {
	clContext   C.cl_context
	errCallback func(string, unsafe.Pointer, uint64, unsafe.Pointer)
	userdata    interface{}
}

//export contextErrorCallback
func contextErrorCallback(errinfo *C.char, privateinfo unsafe.Pointer, cb C.size_t, userData unsafe.Pointer) {
	ctx := (*Context)(userData)
	_ = ctx
	//do something
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clCreateContext.html
func CreateContext(properties *ContextProperties, numDevices uint32, devices *DeviceId, errcb func(string, unsafe.Pointer, uint64, unsafe.Pointer), userdata interface{}, errcode *int32) *Context {
	ctx := Context{nil, errcb, userdata}
	ctx.clContext = C.clCreateContext((*C.cl_context_properties)(unsafe.Pointer(properties)),
		C.cl_uint(numDevices),
		(*C.cl_device_id)(unsafe.Pointer(devices)),
		(*[0]byte)(C.contextErrorCallback),
		unsafe.Pointer(&ctx),
		(*C.cl_int)(unsafe.Pointer(errcode)))
	return &ctx
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clCreateContextFromType.html
func CreateContextFromType(properties *ContextProperties, deviceType DeviceType, errcb func(string, unsafe.Pointer, uint64, unsafe.Pointer), userdata interface{}, errcode *int32) *Context {
	ctx := Context{nil, errcb, userdata}
	var f *[0]byte
	var u unsafe.Pointer
	if errcb != nil {
		f = (*[0]byte)(C.contextErrorCallback)
		u = unsafe.Pointer(&ctx)
	}
	ctx.clContext = C.clCreateContextFromType((*C.cl_context_properties)(unsafe.Pointer(properties)),
		C.cl_device_type(deviceType), f, u,
		(*C.cl_int)(unsafe.Pointer(errcode)))
	return &ctx
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clRetainContext.html
func RetainContext(context *Context) ErrorCode {
	return ErrorCode(C.clRetainContext(context.clContext))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clReleaseContext.html
func ReleaseContext(context *Context) ErrorCode {
	return ErrorCode(C.clReleaseContext(context.clContext))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clGetContextInfo.html
func GetContextInfo(context *Context, paramName ContextInfo, paramValueSize uint64, data unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetContextInfo(context.clContext, C.cl_context_info(paramName), C.size_t(paramValueSize), data, (*C.size_t)(paramValueSizeRet)))
}

/*=================================================================================================================
=====================================================Command queue Api=============================================
=================================================================================================================*/

type CommandQueue C.cl_command_queue

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clCreateCommandQueue.html
func CreateCommandQueue(context Context, did DeviceId, properties CommandQueueProperties, errcode *int32) CommandQueue {
	return CommandQueue(C.clCreateCommandQueue(context.clContext, did, C.cl_command_queue_properties(properties), (*C.cl_int)(errcode)))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clRetainCommandQueue.html
func RetainCommandQueue(cq CommandQueue) {
	C.clRetainCommandQueue(cq)
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clReleaseCommandQueue.html
func ReleaseCommandQueue(cq CommandQueue) {
	C.clReleaseCommandQueue(cq)
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clGetCommandQueueInfo.html
func GetCommandQueueInfo(cq CommandQueue, paramName CommandQueueInfo, paramValueSize uint64, data unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetCommandQueueInfo(cq, C.cl_command_queue_info(paramName), C.size_t(paramValueSize), data, (*C.size_t)(paramValueSizeRet)))
}

/*=================================================================================================================
=====================================================Memory Object Api=============================================
=================================================================================================================*/

type ImageFormat struct {
	imageChannelOrder    ChannelOrder
	imageChannelDataType ChannelType
}

//see (https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/climf.html
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

//see (https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/climde.html
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

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clCreateBuffer.html
func CreateBuffer(context Context, flags, paramValueSize MemFlags, hostPtr unsafe.Pointer, errcode *int32) Mem {
	return Mem(C.clCreateBuffer(context.clContext, C.cl_mem_flags(flags), C.size_t(paramValueSize), hostPtr, (*C.cl_int)(unsafe.Pointer(errcode))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clCreateSubBuffer.html
func CreateSubBuffer(mem Mem, flags MemFlags, bufferCreateType BufferCreateType, bufferCreateInfo unsafe.Pointer, errcode *int32) Mem {
	return Mem(C.clCreateSubBuffer(mem, C.cl_mem_flags(flags), C.cl_buffer_create_type(bufferCreateType), bufferCreateInfo, (*C.cl_int)(unsafe.Pointer(errcode))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clCreateImage.html
func CreateImage(context Context, flags MemFlags, imageFormat ImageFormat, imageDesc ImageDesc, hostPtr unsafe.Pointer, errcode *int32) Mem {
	return Mem(C.clCreateImage(context.clContext, C.cl_mem_flags(flags), imageFormat.toC(), imageDesc.toC(), hostPtr, (*C.cl_int)(unsafe.Pointer(errcode))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clRetainMemObject.html
func RetainMemObject(mem Mem) ErrorCode {
	return ErrorCode(C.clRetainMemObject(mem))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clReleaseMemObject.html
func ReleaseMemObject(mem Mem) ErrorCode {
	return ErrorCode(C.clReleaseMemObject(mem))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clGetSupportedImageFormats.html
func GetSupportedImageFormats(context Context, flags MemFlags, memObjectType MemObjectType, numEntries uint32, imageformats *ImageFormat, numImageFormat *uint32) ErrorCode {
	return ErrorCode(C.clGetSupportedImageFormats(context.clContext,
		C.cl_mem_flags(flags),
		C.cl_mem_object_type(memObjectType),
		C.cl_uint(numEntries),
		(*C.cl_image_format)(unsafe.Pointer(imageformats)),
		(*C.cl_uint)(numImageFormat)))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clGetMemObjectInfo.html
func GetMemObjectInfo(mem Mem, paramName MemInfo, paramValueSize uint64, paramValue unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetMemObjectInfo(mem, C.cl_mem_info(paramName), C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clGetImageInfo.html
func GetImageInfo(mem Mem, paramName ImageInfo, paramValueSize uint64, paramValue unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetImageInfo(mem, C.cl_image_info(paramName), C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}

type memObjectCallbackHolder struct {
	cbfunc   func(Mem, interface{})
	userdata interface{}
}

var mochHolder map[*memObjectCallbackHolder]struct{}

//export memObjectDestroyCallback
func memObjectDestroyCallback(mem C.cl_mem, userData unsafe.Pointer) {
	moch := (*memObjectCallbackHolder)(userData)
	moch.cbfunc(Mem(mem), moch.userdata)
	delete(mochHolder, moch)
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clSetMemObjectDestructorCallback.html
func SetMemObjectDestructorCallback(mem Mem, destroyCb func(Mem, interface{}), userData interface{}) ErrorCode {
	cbh := memObjectCallbackHolder{destroyCb, userData}
	mochHolder[&cbh] = struct{}{}
	return ErrorCode(C.clSetMemObjectDestructorCallback(mem, (*[0]byte)(C.memObjectDestroyCallback), unsafe.Pointer(&cbh)))
}

/*=================================================================================================================
========================================================Sampler Api================================================
=================================================================================================================*/

type Sampler C.cl_sampler

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clCreateSampler.html
func CreateSampler(context Context, normalizedCoords Bool, addressingMode AddressingMode, filterMode FilterMode, errcode *int32) Sampler {
	return Sampler(C.clCreateSampler(context.clContext, C.cl_bool(normalizedCoords), C.cl_addressing_mode(addressingMode), C.cl_filter_mode(filterMode), (*C.cl_int)(unsafe.Pointer(errcode))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clRetainSampler.html
func RetainSampler(sampler Sampler) ErrorCode {
	return ErrorCode(C.clRetainSampler(sampler))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clReleaseSampler.html
func ReleaseSampler(sampler Sampler) ErrorCode {
	return ErrorCode(C.clReleaseSampler(sampler))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clGetSamplerInfo.html
func GetSamplerInfo(sampler Sampler, paramName SamplerInfo, paramValueSize uint64, paramValue unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetSamplerInfo(sampler, C.cl_sampler_info(paramName), C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}

/*=================================================================================================================
=====================================================Program Object Api============================================
=================================================================================================================*/

type Program C.cl_program

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clCreateProgramWithSource.html
func CreateProgramWithSource(context Context, count uint32, src **uint8, lengths *uint64, errcode *int32) Program {
	return Program(C.clCreateProgramWithSource(context.clContext, C.cl_uint(count), (**C.char)(unsafe.Pointer(src)), (*C.size_t)(unsafe.Pointer(lengths)), (*C.cl_int)(unsafe.Pointer(errcode))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clCreateProgramWithBinary.html
func CreateProgramWithBinary(context Context, numDevices uint32, devices *DeviceId, lengths *uint64, binaries **uint8, binaryStatus *int32, errcode *int32) Program {
	return Program(C.clCreateProgramWithBinary(context.clContext,
		C.cl_uint(numDevices),
		(*C.cl_device_id)(unsafe.Pointer(devices)),
		(*C.size_t)(unsafe.Pointer(lengths)),
		(**C.uchar)(unsafe.Pointer(binaries)),
		(*C.cl_int)(unsafe.Pointer(binaryStatus)),
		(*C.cl_int)(unsafe.Pointer(errcode))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clCreateProgramWithBuiltInKernels.html
func CreateProgramWithBuiltInKernels(context Context, numDevices uint32, devices *DeviceId, kernelNames *uint8, errcode *int32) Program {
	return Program(C.clCreateProgramWithBuiltInKernels(context.clContext, C.cl_uint(numDevices), (*C.cl_device_id)(unsafe.Pointer(devices)), (*C.char)(unsafe.Pointer(kernelNames)), (*C.cl_int)(unsafe.Pointer(errcode))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clRetainProgram.html
func RetainProgram(prog Program) ErrorCode {
	return ErrorCode(C.clRetainProgram(prog))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clReleaseProgram.html
func ReleaseProgram(prog Program) ErrorCode {
	return ErrorCode(C.clReleaseProgram(prog))
}

type programObjectBuildCompleteHolder struct {
	cbfunc   func(Program, interface{})
	userdata interface{}
}

var pobchHolder map[*programObjectBuildCompleteHolder]struct{}

//export programObjectBuildCompleteCallback
func programObjectBuildCompleteCallback(prog C.cl_program, userdata unsafe.Pointer) {
	pobch := (*programObjectBuildCompleteHolder)(userdata)
	pobch.cbfunc(Program(prog), pobch.userdata)
	delete(pobchHolder, pobch)
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clBuildProgram.html
func BuildProgram(prog Program, numDevices uint32, devices *DeviceId, options *uint8, buildcomplete func(Program, interface{}), userdata interface{}) ErrorCode {
	var f *[0]byte
	var u unsafe.Pointer
	if buildcomplete != nil {
		pobch := programObjectBuildCompleteHolder{buildcomplete, userdata}
		pobchHolder[&pobch] = struct{}{}
		f = (*[0]byte)(C.programObjectBuildCompleteCallback)
		u = unsafe.Pointer(&pobch)
	}
	return ErrorCode(C.clBuildProgram(prog, C.cl_uint(numDevices), (*C.cl_device_id)(unsafe.Pointer(devices)), (*C.char)(unsafe.Pointer(options)), f, u))
}

type programObjectCompileCompleteHolder struct {
	cbfunc   func(Program, interface{})
	userData interface{}
}

var poccHolder map[*programObjectCompileCompleteHolder]struct{}

//export programObjectCompileCompleteCallback
func programObjectCompileCompleteCallback(prog C.cl_program, userData unsafe.Pointer) {
	pocc := (*programObjectCompileCompleteHolder)(userData)
	pocc.cbfunc(Program(prog), pocc.userData)
	delete(poccHolder, pocc)
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clclCompileProgram.html
func CompileProgram(prog Program, numDevices uint32, devices *DeviceId, options *uint8, numInputHeaders uint32, inputHeaders *Program, headerIncludeNames **uint8, notify func(Program, interface{}), userData interface{}) ErrorCode {
	var f *[0]byte
	var u unsafe.Pointer
	if notify != nil {
		pocc := programObjectCompileCompleteHolder{notify, userData}
		poccHolder[&pocc] = struct{}{}
		f = (*[0]byte)(C.programObjectCompileCompleteCallback)
		u = unsafe.Pointer(&pocc)
	}
	return ErrorCode(C.clCompileProgram(prog, C.cl_uint(numDevices), (*C.cl_device_id)(unsafe.Pointer(devices)), (*C.char)(unsafe.Pointer(options)), C.cl_uint(numInputHeaders), (*C.cl_program)(unsafe.Pointer(inputHeaders)), (**C.char)(unsafe.Pointer(headerIncludeNames)),
		f,
		u))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clLinkProgram.html
func LinkProgram(context Context, numDevices uint32, devices *DeviceId, options *uint8, numInputPrograms uint32, inputPrograms *Program, notify func(Program, interface{}), userData interface{}, errcode *int32) Program {
	var f *[0]byte
	var u unsafe.Pointer
	if notify != nil {
		pocc := programObjectCompileCompleteHolder{notify, userData}
		poccHolder[&pocc] = struct{}{}
		f = (*[0]byte)(C.programObjectCompileCompleteCallback)
		u = unsafe.Pointer(&pocc)
	}
	return Program(C.clLinkProgram(context.clContext, C.cl_uint(numDevices), (*C.cl_device_id)(unsafe.Pointer(devices)), (*C.char)(unsafe.Pointer(options)), C.cl_uint(numInputPrograms), (*C.cl_program)(unsafe.Pointer(inputPrograms)), f, u, (*C.cl_int)(unsafe.Pointer(errcode))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clUnloadPlatformCompiler.html
func UnloadPlatformCompiler(pid PlatformID) ErrorCode {
	return ErrorCode(C.clUnloadPlatformCompiler(pid))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clGetProgramInfo.html
func GetProgramInfo(prog Program, paramName ProgramInfo, paramValueSize uint64, paramValue unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetProgramInfo(prog, C.cl_program_info(paramName), C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clGetProgramBuildInfo.html
func GetProgramBuildInfo(prog Program, device DeviceId, paramName ProgramBuildInfo, paramValueSize uint64, paramValue unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetProgramBuildInfo(prog, device, C.cl_program_build_info(paramName), C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}

/*=================================================================================================================
=====================================================Kernel Object Api=============================================
=================================================================================================================*/

type Kernel C.cl_kernel

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clCreateKernel.html
func CreateKernel(prog Program, kernelName *uint8, errcode *int32) Kernel {
	return Kernel(C.clCreateKernel(prog, (*C.char)(unsafe.Pointer(kernelName)), (*C.cl_int)(unsafe.Pointer(errcode))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clCreateKernelsInProgram.html
func CreateKernelsInProgram(prog Program, numKernels uint32, kernels *Kernel, numKernelsRet *uint32) ErrorCode {
	return ErrorCode(C.clCreateKernelsInProgram(prog, C.cl_uint(numKernels), (*C.cl_kernel)(unsafe.Pointer(kernels)), (*C.cl_uint)(unsafe.Pointer(numKernelsRet))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clRetainKernel.html
func RetainKernel(ker Kernel) ErrorCode {
	return ErrorCode(C.clRetainKernel(ker))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clReleaseKernel.html
func ReleaseKernel(ker Kernel) ErrorCode {
	return ErrorCode(C.clReleaseKernel(ker))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clSetKernelArg.html
func SetKernelArg(ker Kernel, argIndex uint32, argSize uint64, argValue unsafe.Pointer) ErrorCode {
	return ErrorCode(C.clSetKernelArg(ker, C.cl_uint(argIndex), C.size_t(argSize), argValue))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clGetKernelInfo.html
func GetKernelInfo(ker Kernel, paramName KernelInfo, paramValueSize uint64, paramValue unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetKernelInfo(ker, C.cl_kernel_info(paramName), C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clGetKernelArgInfo.html
func GetKernelArgInfo(ker Kernel, argIndex uint32, kernelArgInfo KernelArgInfo, paramValueSize uint64, paramValue unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetKernelArgInfo(ker, C.cl_uint(argIndex), C.cl_kernel_arg_info(kernelArgInfo), C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clGetKernelWorkGroupInfo.html
func GetKernelWorkGroupInfo(ker Kernel, did DeviceId, paramName KernelWorkGroupInfo, paramValueSize uint64, paramValue unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetKernelWorkGroupInfo(ker, did, C.cl_kernel_work_group_info(paramName), C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}

/*=================================================================================================================
======================================================Event Object Api=============================================
=================================================================================================================*/

type Event struct {
	clEvent C.cl_event
	cbFuncs []func()
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clWaitForEvents.html
func WaitForEvents(numEvents uint32, events *Event) ErrorCode {
	return ErrorCode(C.clWaitForEvents(C.cl_uint(numEvents), (*C.cl_event)(unsafe.Pointer(events))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clGetEventInfo.html
func GetEventInfo(e Event, paramName uint32, paramValueSize uint64, paramValue unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetEventInfo(e.clEvent, C.cl_event_info(paramName), C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clCreateUserEvent.html
func CreateUserEvent(context Context, errcode *int32) Event {
	return Event{C.clCreateUserEvent(context.clContext, (*C.cl_int)(unsafe.Pointer(errcode))), make([]func(), 0)}
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clRetainEvent.html
func RetainEvent(e Event) ErrorCode {
	return ErrorCode(C.clRetainEvent(e.clEvent))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clReleaseEvent.html
func ReleaseEvent(e Event) ErrorCode {
	return ErrorCode(C.clReleaseEvent(e.clEvent))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clSetUserEventStatus.html
func SetUserEventStatus(e Event, execStatus int32) ErrorCode {
	return ErrorCode(C.clSetUserEventStatus(e.clEvent, C.cl_int(execStatus)))
}

/*
type eventCbHolder struct {
	cbfunc   func()
	userData interface{}
}

var ecbhHolder map[*eventCbHolder]struct{}

//exp
//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/cleventCallbackCallbackort.html  eventCallbackCallback
func eventCallbackCallback() {

}

//I say hey, that doesn't wooooorkkkk!
//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clSetEventCallback.html
func SetEventCallback(e Event, commandExecCallbackType int32, notify func(Event, int32, interface{}), userData interface{}) ErrorCode {
	return ErrorCode(C.clSetEventCallback(e, C.cl_int(commandExecCallbackType), f))
}*/

var code = `
extern CL_API_ENTRY cl_int CL_API_CALL
clSetEventCallback( cl_event    /* event */,
                    cl_int      /* command_exec_callback_type */,
                    void (CL_CALLBACK * /* pfn_notify */)(cl_event, cl_int, void *),
                    void *      /* user_data */) CL_API_SUFFIX__VERSION_1_1;

`

/*=================================================================================================================
======================================================Profiling Api=============================================
=================================================================================================================*/

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clGetEventProfilingInfo.html
func GetEventProfilingInfo(e Event, paramName ProfilingInfo, paramValueSize uint64, paramValue unsafe.Pointer, paramValueSizeRet *uint64) ErrorCode {
	return ErrorCode(C.clGetEventProfilingInfo(e.clEvent, C.cl_profiling_info(paramName), C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}

/*=================================================================================================================
====================================================Flush and Finish Api===========================================
=================================================================================================================*/

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clFlush.html
func Flush(cq CommandQueue) ErrorCode {
	return ErrorCode(C.clFlush(cq))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clFinish.html
func Finish(cq CommandQueue) ErrorCode {
	return ErrorCode(C.clFinish(cq))
}

/*=================================================================================================================
=========================================================Enqueue Api===============================================
=================================================================================================================*/

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clEnqueueReadBuffer.html
func EnqueueReadBuffer(cq CommandQueue, buffer Mem, blocking_read Bool, offset uint64, size uint64, ptr unsafe.Pointer, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueReadBuffer(cq, buffer, C.cl_bool(blocking_read), C.size_t(offset), C.size_t(size), unsafe.Pointer(ptr), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clEnqueueReadBufferRect.html
func EnqueueReadBufferRect(cq CommandQueue, buffer Mem, blocking_read Bool, buffer_offset *uint64, host_offset *uint64, region *uint64, buffer_row_pitch uint64, buffer_slice_pitch uint64, host_row_pitch uint64, host_slice_pitch uint64, ptr unsafe.Pointer, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueReadBufferRect(cq, buffer, C.cl_bool(blocking_read), (*C.size_t)(unsafe.Pointer(buffer_offset)), (*C.size_t)(unsafe.Pointer(host_offset)), (*C.size_t)(unsafe.Pointer(region)), C.size_t(buffer_row_pitch), C.size_t(buffer_slice_pitch), C.size_t(host_row_pitch), C.size_t(host_slice_pitch), unsafe.Pointer(ptr), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clEnqueueWriteBuffer.html
func EnqueueWriteBuffer(cq CommandQueue, buffer Mem, blocking_write Bool, offset uint64, size uint64, ptr unsafe.Pointer, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueWriteBuffer(cq, buffer, C.cl_bool(blocking_write), C.size_t(offset), C.size_t(size), unsafe.Pointer(ptr), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clEnqueueWriteBufferRect.html
func EnqueueWriteBufferRect(cq CommandQueue, buffer Mem, blocking_write Bool, buffer_offset *uint64, host_offset *uint64, region *uint64, buffer_row_pitch uint64, buffer_slice_pitch uint64, host_row_pitch uint64, host_slice_pitch uint64, ptr unsafe.Pointer, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueWriteBufferRect(cq, buffer, C.cl_bool(blocking_write), (*C.size_t)(unsafe.Pointer(buffer_offset)), (*C.size_t)(unsafe.Pointer(host_offset)), (*C.size_t)(unsafe.Pointer(region)), C.size_t(buffer_row_pitch), C.size_t(buffer_slice_pitch), C.size_t(host_row_pitch), C.size_t(host_slice_pitch), unsafe.Pointer(ptr), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clEnqueueFillBuffer.html
func EnqueueFillBuffer(cq CommandQueue, buffer Mem, pattern unsafe.Pointer, pattern_size uint64, offset uint64, size uint64, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueFillBuffer(cq, buffer, unsafe.Pointer(pattern), C.size_t(pattern_size), C.size_t(offset), C.size_t(size), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clEnqueueCopyBuffer.html
func EnqueueCopyBuffer(cq CommandQueue, src_buffer Mem, dst_buffer Mem, src_offset uint64, dst_offset uint64, size uint64, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueCopyBuffer(cq, src_buffer, dst_buffer, C.size_t(src_offset), C.size_t(dst_offset), C.size_t(size), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clEnqueueCopyBufferRect.html
func EnqueueCopyBufferRect(cq CommandQueue, src_buffer Mem, dst_buffer Mem, src_origin *uint64, dst_origin *uint64, region *uint64, src_row_pitch uint64, src_slice_pitch uint64, dst_row_pitch uint64, dst_slice_pitch uint64, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueCopyBufferRect(cq, src_buffer, dst_buffer, (*C.size_t)(unsafe.Pointer(src_origin)), (*C.size_t)(unsafe.Pointer(dst_origin)), (*C.size_t)(unsafe.Pointer(region)), C.size_t(src_row_pitch), C.size_t(src_slice_pitch), C.size_t(dst_row_pitch), C.size_t(dst_slice_pitch), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clEnqueueReadImage.html
func EnqueueReadImage(cq CommandQueue, image Mem, blocking_read Bool, origin3 *uint64, region3 *uint64, row_pitch uint64, slice_pitch uint64, ptr unsafe.Pointer, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueReadImage(cq, image, C.cl_bool(blocking_read), (*C.size_t)(unsafe.Pointer(origin3)), (*C.size_t)(unsafe.Pointer(region3)), C.size_t(row_pitch), C.size_t(slice_pitch), unsafe.Pointer(ptr), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clEnqueueWriteImage.html
func EnqueueWriteImage(cq CommandQueue, image Mem, blocking_write Bool, origin3 *uint64, region3 *uint64, input_row_pitch uint64, input_slice_pitch uint64, ptr unsafe.Pointer, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueWriteImage(cq, image, C.cl_bool(blocking_write), (*C.size_t)(unsafe.Pointer(origin3)), (*C.size_t)(unsafe.Pointer(region3)), C.size_t(input_row_pitch), C.size_t(input_slice_pitch), unsafe.Pointer(ptr), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clEnqueueFillImage.html
func EnqueueFillImage(cq CommandQueue, image Mem, fill_color unsafe.Pointer, origin3 *uint64, region3 *uint64, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueFillImage(cq, image, unsafe.Pointer(fill_color), (*C.size_t)(unsafe.Pointer(origin3)), (*C.size_t)(unsafe.Pointer(region3)), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clEnqueueCopyImage.html
func EnqueueCopyImage(cq CommandQueue, src_image Mem, dst_image Mem, src_origin3 *uint64, dst_origin3 *uint64, region3 *uint64, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueCopyImage(cq, src_image, dst_image, (*C.size_t)(unsafe.Pointer(src_origin3)), (*C.size_t)(unsafe.Pointer(dst_origin3)), (*C.size_t)(unsafe.Pointer(region3)), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clEnqueueCopyImageToBuffer.html
func EnqueueCopyImageToBuffer(cq CommandQueue, src_image Mem, dst_buffer Mem, src_origin3 *uint64, region3 *uint64, dst_offset uint64, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueCopyImageToBuffer(cq, src_image, dst_buffer, (*C.size_t)(unsafe.Pointer(src_origin3)), (*C.size_t)(unsafe.Pointer(region3)), C.size_t(dst_offset), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clEnqueueCopyBufferToImage.html
func EnqueueCopyBufferToImage(cq CommandQueue, src_buffer Mem, dst_image Mem, src_offset uint64, dst_origin3 *uint64, region3 *uint64, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueCopyBufferToImage(cq, src_buffer, dst_image, C.size_t(src_offset), (*C.size_t)(unsafe.Pointer(dst_origin3)), (*C.size_t)(unsafe.Pointer(region3)), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clEnqueueMapBuffer.html
func EnqueueMapBuffer(cq CommandQueue, buffer Mem, blocking_map Bool, map_flags uint64, offset uint64, size uint64, num_events_in_wait_list uint32, event_wait_list *Event, event *Event, errcode_ret *int32) unsafe.Pointer {
	return unsafe.Pointer(C.clEnqueueMapBuffer(cq, buffer, C.cl_bool(blocking_map), C.cl_map_flags(map_flags), C.size_t(offset), C.size_t(size), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event)), (*C.cl_int)(unsafe.Pointer(errcode_ret))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clEnqueueMapImage.html
func EnqueueMapImage(cq CommandQueue, image Mem, blocking_map Bool, map_flags MapFlags, origin3 *uint64, region3 *uint64, image_row_pitch *uint64, image_slice_pitch *uint64, num_events_in_wait_list uint32, event_wait_list *Event, event *Event, errcode_ret *int32) unsafe.Pointer {
	return unsafe.Pointer(C.clEnqueueMapImage(cq, image, C.cl_bool(blocking_map), C.cl_map_flags(map_flags), (*C.size_t)(unsafe.Pointer(origin3)), (*C.size_t)(unsafe.Pointer(region3)), (*C.size_t)(unsafe.Pointer(image_row_pitch)), (*C.size_t)(unsafe.Pointer(image_slice_pitch)), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event)), (*C.cl_int)(unsafe.Pointer(errcode_ret))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clEnqueueUnmapMemObject.html
func EnqueueUnmapMemObject(cq CommandQueue, memobj Mem, mapped_ptr unsafe.Pointer, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueUnmapMemObject(cq, memobj, unsafe.Pointer(mapped_ptr), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clEnqueueMigrateMemObjects.html
func EnqueueMigrateMemObjects(cq CommandQueue, num_mem_objects uint32, mem_objects *Mem, flags MemMigrationFlags, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueMigrateMemObjects(cq, C.cl_uint(num_mem_objects), (*C.cl_mem)(unsafe.Pointer(mem_objects)), C.cl_mem_migration_flags(flags), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clEnqueueNDRangeKernel.html
func EnqueueNDRangeKernel(cq CommandQueue, kernel Kernel, work_dim uint32, global_work_offset *uint64, global_work_size *uint64, local_work_size *uint64, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueNDRangeKernel(cq, kernel, C.cl_uint(work_dim), (*C.size_t)(unsafe.Pointer(global_work_offset)), (*C.size_t)(unsafe.Pointer(global_work_size)), (*C.size_t)(unsafe.Pointer(local_work_size)), C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clEnqueueTask.html
func EnqueueTask(cq CommandQueue, kernel Kernel, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueTask(cq, kernel, C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

/*
//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clEnqueueNativeKernel.html
func EnqueueNativeKernel(cq CommandQueue ,userfunc func(),args unsafe.Pointer,cb_args uint64, num_mem_objects uint32 ,mem_list *Mem , args_mem_loc unsafe.Pointer ,num_events_in_wait_list uint32 ,event_wait_list *Event , event *Event) int32{
	return ErrorCode(C.clEnqueueNativeKernel(cq,userfunc,unsafe.Pointer(args),C.size_t(cb_args),C.cl_uint(num_mem_objects),(*C.cl_mem)(unsafe.Pointer(mem_list)),args_mem_loc,C.cl_uint(num_events_in_wait_list),(*C.cl_event)(unsafe.Pointer(event_wait_list)),(*C.cl_event)(unsafe.Pointer(event))))
}
*/

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clEnqueueMarkerWithWaitList.html
func EnqueueMarkerWithWaitList(cq CommandQueue, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueMarkerWithWaitList(cq, C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clEnqueueBarrierWithWaitList.html
func EnqueueBarrierWithWaitList(cq CommandQueue, num_events_in_wait_list uint32, event_wait_list *Event, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueBarrierWithWaitList(cq, C.cl_uint(num_events_in_wait_list), (*C.cl_event)(unsafe.Pointer(event_wait_list)), (*C.cl_event)(unsafe.Pointer(event))))
}

/*=================================================================================================================
===================================================Extension functions Api=========================================
=================================================================================================================*/

//see https://www.khronos.org/registry/cl/sdk/1.2/docs/man/xhtml/clGetExtensionFunctionAddressForPlatform.html
func GetExtensionFunctionAddressForPlatform(platform PlatformID, func_name *int8) unsafe.Pointer {
	return unsafe.Pointer(C.clGetExtensionFunctionAddressForPlatform(platform, (*C.char)(func_name)))
}

/*
//Deprecated OpenCL 1.1 APIs
//see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateImage2D.html
func CreateImage2D(context Context, flags uint32, image_format *ImageFormat, image_width uint64, image_height uint64, image_row_pitch uint64, host_ptr unsafe.Pointer, errcode_ret *int32) Mem {
	return Mem(C.clCreateImage2D(context.clContext, C.cl_mem_flags(flags), (*C.cl_image_format)(image_format.toC()), C.size_t(image_width), C.size_t(image_height), C.size_t(image_row_pitch), host_ptr, (*C.cl_int)(errcode_ret)))
}

//Deprecated OpenCL 1.1 APIs
//see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateImage3D.html
func CreateImage3D(context Context, flags uint32, image_format *ImageFormat, image_width uint64, image_height uint64, image_depth uint64, image_row_pitch uint64, image_slice_pitch uint64, host_ptr unsafe.Pointer, errcode_ret *int32) Mem {
	return Mem(C.clCreateImage3D(context.clContext, C.cl_mem_flags(flags), (*C.cl_image_format)(image_format.toC()), C.size_t(image_width), C.size_t(image_height), C.size_t(image_depth), C.size_t(image_row_pitch), C.size_t(image_slice_pitch), host_ptr, (*C.cl_int)(errcode_ret)))
}

//Deprecated OpenCL 1.1 APIs
//see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueMarker.html
func EnqueueMarker(cq CommandQueue, event *Event) ErrorCode {
	return ErrorCode(C.clEnqueueMarker(cq, (*C.cl_event)(unsafe.Pointer(event))))
}

//Deprecated OpenCL 1.1 APIs
//see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueWaitForEvents.html
func EnqueueWaitForEvents(cq CommandQueue, num_events uint32, event_list *Event) ErrorCode {
	return ErrorCode(C.clEnqueueWaitForEvents(cq, C.cl_uint(num_events), (*C.cl_event)(unsafe.Pointer(event_list))))
}

//Deprecated OpenCL 1.1 APIs
//see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueBarrier.html
func EnqueueBarrier(cq CommandQueue) ErrorCode {
	return ErrorCode(C.clEnqueueBarrier(cq))
}

//Deprecated OpenCL 1.1 APIs
// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clUnloadCompiler.html
func UnloadCompiler() ErrorCode {
	return ErrorCode(C.clUnloadCompiler())
}

//Deprecated OpenCL 1.1 APIs
// see https://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/.html
func GetExtensionFunctionAddress(func_name *int8) unsafe.Pointer {
	return unsafe.Pointer(C.clGetExtensionFunctionAddress((*C.char)(func_name)))
}
*/
