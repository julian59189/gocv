package contrib

import (
	"testing"

	"gocv.io/x/gocv"
)

func TestDNNSuperres(t *testing.T) {
	// read image
	img := gocv.IMRead("../images/face.jpg", 1)
	if img.Empty() {
		t.Error("TestTracker " + name + " input img failed to load")
	}
	defer img.Close()

	// init super res
	sr := DnnSuperResImpl()
	// if sr == nil {
	// 	t.Error("DnnSuperResImpl should not be nil")
	// }

	//Read the desired model
	path := "FSRCNN_x2.pb"
	sr.readModel(path)

	//Set the desired model and scale to get correct pre- and post-processing
	sr.setModel("fsrcnn", 2)

	//Upscale
	imgNew := gocv.NewMat()
	defer imgNew.close()
	sr.upsample(img, imgNew)
}
