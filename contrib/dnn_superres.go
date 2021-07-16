package contrib

/*
#include <stdlib.h>
#include "dnn_superres.h"
*/
import "C"
import (
	"unsafe"
)

// Net allows you to create and manipulate comprehensive artificial neural networks.
//
// For further details, please see:
//
//
type DNNNet struct {
	// C.Net
	p unsafe.Pointer
}

// // Empty constructor allows setting model later
// //
// //
// //
// func DnnSuperResImpl() DNNNet {
// 	return DNNNet{p: unsafe.Pointer(C.DnnSuperResImpl())}
// }

// Constructor which handles setting model and scale
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d11/classcv_1_1dnn__superres_1_1DnnSuperResImpl.html#a4f6e7d88778f73bcfdf43c0d9bff2c1a
//
func DnnSuperResImpl(algo string, scale int) DNNNet {
	cAlgo := C.CString(algo)
	defer C.free(unsafe.Pointer(cAlgo), C.int(scale))

	return DNNNet{p: unsafe.Pointer(C.DnnSuperResImpl(cAlgo, scale))}
}

// Read model file
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d11/classcv_1_1dnn__superres_1_1DnnSuperResImpl.html#af56741a70ee1346efcc69789f99200d8
//
func (net *DNNNet) ReadModel(path string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	ReadModel((C.DNNNet)(net.p), cPath)
}

// // Read model file
// //
// // For further details, please see:
// // https://docs.opencv.org/master/d8/d11/classcv_1_1dnn__superres_1_1DnnSuperResImpl.html#a1d062da8e770781eb7323a44a25b6fa3
// //
// func ReadModel(weights, definition string) {
// 	cWeights := C.CString(weights)
// 	cDefinition := C.CString(definition)
// 	defer C.free(unsafe.Pointer(cWeights), unsafe.Pointer(cDefinition))

// 	ReadModel(cWeights, cDefinition)
// }

// Set the model
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d11/classcv_1_1dnn__superres_1_1DnnSuperResImpl.html#ab4d5e45240e7dbc436f077d34bff8854
//
func (net *DNNNet) SetModel(algo string, scale int) {
	cAlgo := C.CString(algo)
	defer C.free(unsafe.Pointer(cAlgo), C.int(scale))

	C.SetModel((C.DNNNet)(net.p), cAlgo)
}

// Upsample the image using the selected model and method
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d11/classcv_1_1dnn__superres_1_1DnnSuperResImpl.html#a3d8bf9e39c75939ab8d2de8396201a89
//
func (net *NeDNNNett) Upsample(src Mat, dst *Mat) {
	C.upsample((C.DNNNet)(net.p), src.p, dst.p)
}
