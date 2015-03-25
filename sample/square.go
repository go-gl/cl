package main

import (
	"github.com/hydroflame/gopencl/v1.2/cl"
	"log"
	"math"
	"math/rand"
	"unsafe"
)

const (
	DATA_SIZE = 1024
)

var KernelSource string = `
__kernel void square(
   __global float* input,
   __global float* output,
   const unsigned int count)
{
   int i = get_global_id(0);
   if(i < count)
	   output[i] = input[i] * input[i];
}` + "\x00"

func main() {

	data := make([]float32, DATA_SIZE)
	for x := 0; x < len(data); x++ {
		data[x] = rand.Float32()*99 + 1
	}

	//Get Device
	var device cl.DeviceId
	err := cl.GetDeviceIDs(nil, cl.DEVICE_TYPE_GPU, 1, &device, nil)
	if err != cl.SUCCESS {
		log.Fatal("Failed to create device group")
	}
	var errptr *int32

	//Create Computer Context
	context := cl.CreateContext(nil, 1, &device, nil, nil, errptr)
	if errptr != nil && cl.ErrorCode(*errptr) != cl.SUCCESS {
		log.Fatal("couldnt create context")
	}

	//Create Command Queue
	cq := cl.CreateCommandQueue(*context, device, 0, errptr)
	if errptr != nil && cl.ErrorCode(*errptr) != cl.SUCCESS {
		log.Fatal("couldnt create command queue")
	}

	//Create program
	srcptr := cl.Str(KernelSource)
	program := cl.CreateProgramWithSource(*context, 1, &srcptr, nil, errptr)
	if errptr != nil && cl.ErrorCode(*errptr) != cl.SUCCESS {
		log.Fatal("couldnt create program")
	}

	err = cl.BuildProgram(program, 1, &device, nil, nil, nil)
	if err != cl.SUCCESS {
		var length uint64
		buffer := make([]byte, DATA_SIZE)

		log.Println("Error: Failed to build program executable!")
		cl.GetProgramBuildInfo(program, device, cl.PROGRAM_BUILD_LOG, uint64(len(buffer)), unsafe.Pointer(&buffer[0]), &length)
		log.Fatal(string(buffer[0:length]))
	}

	//Get Kernel
	kernel := cl.CreateKernel(program, cl.Str("square"+"\x00"), errptr)
	if errptr != nil && cl.ErrorCode(*errptr) != cl.SUCCESS {
		log.Fatal("couldnt create compute kernel")
	}

	//Create buffers
	input := cl.CreateBuffer(*context, cl.MEM_READ_ONLY, 4*DATA_SIZE, nil, errptr)
	if errptr != nil && cl.ErrorCode(*errptr) != cl.SUCCESS {
		log.Fatal("couldnt create input buffer")
	}

	output := cl.CreateBuffer(*context, cl.MEM_WRITE_ONLY, 4*DATA_SIZE, nil, errptr)
	if errptr != nil && cl.ErrorCode(*errptr) != cl.SUCCESS {
		log.Fatal("couldnt create output buffer")
	}

	//Write data
	err = cl.EnqueueWriteBuffer(cq, input, cl.TRUE, 0, 4*DATA_SIZE, unsafe.Pointer(&data[0]), 0, nil, nil)
	if err != cl.SUCCESS {
		log.Fatal("Failed to write to source array")
	}

	//Set kernel args
	count := uint32(DATA_SIZE)
	err = cl.SetKernelArg(kernel, 0, 8, unsafe.Pointer(&input))
	if err != cl.SUCCESS {
		log.Fatal("Failed to write kernel arg 0")
	}
	err = cl.SetKernelArg(kernel, 1, 8, unsafe.Pointer(&output))
	if err != cl.SUCCESS {
		log.Fatal("Failed to write kernel arg 1")
	}
	err = cl.SetKernelArg(kernel, 2, 4, unsafe.Pointer(&count))
	if err != cl.SUCCESS {
		log.Fatal("Failed to write kernel arg 2")
	}

	local := uint64(0)
	err = cl.GetKernelWorkGroupInfo(kernel, device, cl.KERNEL_WORK_GROUP_SIZE, 8, unsafe.Pointer(&local), nil)
	if err != cl.SUCCESS {
		log.Fatal("Failed to get kernel work group info")
	}

	global := local
	err = cl.EnqueueNDRangeKernel(cq, kernel, 1, nil, &global, &local, 0, nil, nil)
	if err != cl.SUCCESS {
		log.Fatal("Failed to execute kernel!")
	}

	cl.Finish(cq)

	results := make([]float32, DATA_SIZE)
	err = cl.EnqueueReadBuffer(cq, output, cl.TRUE, 0, 4*1024, unsafe.Pointer(&results[0]), 0, nil, nil)
	if err != cl.SUCCESS {
		log.Fatal("Failed to read buffer!")
	}

	success := 0
	notzero := 0
	for i, x := range data {
		if math.Abs(float64(x*x-results[i])) < 0.5 {
			success++
		}
		if results[i] > 0 {
			notzero++
		}
		log.Printf("I/O: %f\t%f", x, results[i])
	}

	log.Printf("%d/%d success", success, DATA_SIZE)
	log.Printf("values not zero: %d", notzero)

}
