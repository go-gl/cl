package cl

import (
	"log"
	"reflect"
	"strings"
	"unsafe"
)

var errormap = map[ErrorCode]string{
	SUCCESS:                                   "SUCCESS",
	DEVICE_NOT_FOUND:                          "DEVICE_NOT_FOUND",
	DEVICE_NOT_AVAILABLE:                      "DEVICE_NOT_AVAILABLE",
	COMPILER_NOT_AVAILABLE:                    "COMPILER_NOT_AVAILABLE",
	MEM_OBJECT_ALLOCATION_FAILURE:             "MEM_OBJECT_ALLOCATION_FAILURE",
	OUT_OF_RESOURCES:                          "OUT_OF_RESOURCES",
	OUT_OF_HOST_MEMORY:                        "OUT_OF_HOST_MEMORY",
	PROFILING_INFO_NOT_AVAILABLE:              "PROFILING_INFO_NOT_AVAILABLE",
	MEM_COPY_OVERLAP:                          "MEM_COPY_OVERLAP",
	IMAGE_FORMAT_MISMATCH:                     "IMAGE_FORMAT_MISMATCH",
	IMAGE_FORMAT_NOT_SUPPORTED:                "IMAGE_FORMAT_NOT_SUPPORTED",
	BUILD_PROGRAM_FAILURE:                     "BUILD_PROGRAM_FAILURE",
	MAP_FAILURE:                               "MAP_FAILURE",
	MISALIGNED_SUB_BUFFER_OFFSET:              "MISALIGNED_SUB_BUFFER_OFFSET",
	EXEC_STATUS_ERROR_FOR_EVENTS_IN_WAIT_LIST: "EXEC_STATUS_ERROR_FOR_EVENTS_IN_WAIT_LIST",
	COMPILE_PROGRAM_FAILURE:                   "COMPILE_PROGRAM_FAILURE",
	LINKER_NOT_AVAILABLE:                      "LINKER_NOT_AVAILABLE",
	LINK_PROGRAM_FAILURE:                      "LINK_PROGRAM_FAILURE",
	DEVICE_PARTITION_FAILED:                   "DEVICE_PARTITION_FAILED",
	KERNEL_ARG_INFO_NOT_AVAILABLE:             "KERNEL_ARG_INFO_NOT_AVAILABLE",
	INVALID_VALUE:                             "INVALID_VALUE",
	INVALID_DEVICE_TYPE:                       "INVALID_DEVICE_TYPE",
	INVALID_PLATFORM:                          "INVALID_PLATFORM",
	INVALID_DEVICE:                            "INVALID_DEVICE",
	INVALID_CONTEXT:                           "INVALID_CONTEXT",
	INVALID_QUEUE_PROPERTIES:                  "INVALID_QUEUE_PROPERTIES",
	INVALID_COMMAND_QUEUE:                     "INVALID_COMMAND_QUEUE",
	INVALID_HOST_PTR:                          "INVALID_HOST_PTR",
	INVALID_MEM_OBJECT:                        "INVALID_MEM_OBJECT",
	INVALID_IMAGE_FORMAT_DESCRIPTOR:           "INVALID_IMAGE_FORMAT_DESCRIPTOR",
	INVALID_IMAGE_SIZE:                        "INVALID_IMAGE_SIZE",
	INVALID_SAMPLER:                           "INVALID_SAMPLER",
	INVALID_BINARY:                            "INVALID_BINARY",
	INVALID_BUILD_OPTIONS:                     "INVALID_BUILD_OPTIONS",
	INVALID_PROGRAM:                           "INVALID_PROGRAM",
	INVALID_PROGRAM_EXECUTABLE:                "INVALID_PROGRAM_EXECUTABLE",
	INVALID_KERNEL_NAME:                       "INVALID_KERNEL_NAME",
	INVALID_KERNEL_DEFINITION:                 "INVALID_KERNEL_DEFINITION",
	INVALID_KERNEL:                            "INVALID_KERNEL",
	INVALID_ARG_INDEX:                         "INVALID_ARG_INDEX",
	INVALID_ARG_VALUE:                         "INVALID_ARG_VALUE",
	INVALID_ARG_SIZE:                          "INVALID_ARG_SIZE",
	INVALID_KERNEL_ARGS:                       "INVALID_KERNEL_ARGS",
	INVALID_WORK_DIMENSION:                    "INVALID_WORK_DIMENSION",
	INVALID_WORK_GROUP_SIZE:                   "INVALID_WORK_GROUP_SIZE",
	INVALID_WORK_ITEM_SIZE:                    "INVALID_WORK_ITEM_SIZE",
	INVALID_GLOBAL_OFFSET:                     "INVALID_GLOBAL_OFFSET",
	INVALID_EVENT_WAIT_LIST:                   "INVALID_EVENT_WAIT_LIST",
	INVALID_EVENT:                             "INVALID_EVENT",
	INVALID_OPERATION:                         "INVALID_OPERATION",
	INVALID_GL_OBJECT:                         "INVALID_GL_OBJECT",
	INVALID_BUFFER_SIZE:                       "INVALID_BUFFER_SIZE",
	INVALID_MIP_LEVEL:                         "INVALID_MIP_LEVEL",
	INVALID_GLOBAL_WORK_SIZE:                  "INVALID_GLOBAL_WORK_SIZE",
	INVALID_PROPERTY:                          "INVALID_PROPERTY",
	INVALID_IMAGE_DESCRIPTOR:                  "INVALID_IMAGE_DESCRIPTOR",
	INVALID_COMPILER_OPTIONS:                  "INVALID_COMPILER_OPTIONS",
	INVALID_LINKER_OPTIONS:                    "INVALID_LINKER_OPTIONS",
	INVALID_DEVICE_PARTITION_COUNT:            "INVALID_DEVICE_PARTITION_COUNT",
}

func Str(str string) *uint8 {
	if !strings.HasSuffix(str, "\x00") {
		log.Fatal("str argument missing null terminator", str)
	}
	header := (*reflect.StringHeader)(unsafe.Pointer(&str))
	return (*uint8)(unsafe.Pointer(header.Data))
}

func ErrToStr(e ErrorCode) string {
	return errormap[e]
}
