#ifndef _OPENCV3_DNN_SUPERRES_H_
#define _OPENCV3_DNN_SUPERRES_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/dnn_superres.hpp>

extern "C"
{
#endif

#include "../core.h"

    DnnSuperResImpl();
    DnnSuperResImpl(const String &algo, int scale);
    void readModel(const String &weights, const String &definition);
    void setModel(const String &algo, int scale);
    void upsample(InputArray img, OutputArray result);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_DNN_SUPERRES_H_