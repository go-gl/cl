# CL
Robust cross platform opencl bindings for go

#Examples
see sample/

#Compile optimisation
for better performance (less overhead) compile with 4 `-l` flag:  
`go install -gcflags="-l -l -l -l" github.com/hydroflame/gopencl/v1.2/cl`  
This will inline most functions in the package.
