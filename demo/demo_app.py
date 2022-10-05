# foreign function library to call shared library functions
import ctypes

# load the shared library into c types
library = ctypes.cdll.LoadLibrary('./build/library.so')

# call the shared library function
hello_world = library.helloWorld
hello_world()