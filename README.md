# gopencl
Robust cross platform opencl bindings for go

#Examples
see samples

#Compile optimisation
for better performance (less overhead) compile with 4 `-l` flag:  
`go install -gcflags="-l -l -l -l" github.com/hydroflame/gopencl/v1.2/cl`
