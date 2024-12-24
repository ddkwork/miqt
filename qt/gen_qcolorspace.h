#pragma once
#ifndef MIQT_QT_GEN_QCOLORSPACE_H
#define MIQT_QT_GEN_QCOLORSPACE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QColorSpace;
class QColorTransform;
class QPointF;
class _GUID;
class type_info;
#else
typedef struct QColorSpace QColorSpace;
typedef struct QColorTransform QColorTransform;
typedef struct QPointF QPointF;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QColorSpace* QColorSpace_new();
extern __declspec(dllexport) QColorSpace* QColorSpace_new2(NamedColorSpace namedColorSpace);
extern __declspec(dllexport) QColorSpace* QColorSpace_new3(QPointF* whitePoint, TransferFunction transferFunction);
extern __declspec(dllexport) QColorSpace* QColorSpace_new4(QPointF* whitePoint, struct miqt_array /* of uint16_t */  transferFunctionTable);
extern __declspec(dllexport) QColorSpace* QColorSpace_new5(Primaries primaries, TransferFunction transferFunction);
extern __declspec(dllexport) QColorSpace* QColorSpace_new6(Primaries primaries, float gamma);
extern __declspec(dllexport) QColorSpace* QColorSpace_new7(Primaries primaries, struct miqt_array /* of uint16_t */  transferFunctionTable);
extern __declspec(dllexport) QColorSpace* QColorSpace_new8(QPointF* whitePoint, QPointF* redPoint, QPointF* greenPoint, QPointF* bluePoint, TransferFunction transferFunction);
extern __declspec(dllexport) QColorSpace* QColorSpace_new9(QPointF* whitePoint, QPointF* redPoint, QPointF* greenPoint, QPointF* bluePoint, struct miqt_array /* of uint16_t */  transferFunctionTable);
extern __declspec(dllexport) QColorSpace* QColorSpace_new10(QPointF* whitePoint, QPointF* redPoint, QPointF* greenPoint, QPointF* bluePoint, struct miqt_array /* of uint16_t */  redTransferFunctionTable, struct miqt_array /* of uint16_t */  greenTransferFunctionTable, struct miqt_array /* of uint16_t */  blueTransferFunctionTable);
extern __declspec(dllexport) QColorSpace* QColorSpace_new11(QColorSpace* colorSpace);
extern __declspec(dllexport) QColorSpace* QColorSpace_new12(QPointF* whitePoint, TransferFunction transferFunction, float gamma);
extern __declspec(dllexport) QColorSpace* QColorSpace_new13(Primaries primaries, TransferFunction transferFunction, float gamma);
extern __declspec(dllexport) QColorSpace* QColorSpace_new14(QPointF* whitePoint, QPointF* redPoint, QPointF* greenPoint, QPointF* bluePoint, TransferFunction transferFunction, float gamma);
extern __declspec(dllexport) void QColorSpace_OperatorAssign(QColorSpace* self, QColorSpace* colorSpace);
extern __declspec(dllexport) void QColorSpace_Swap(QColorSpace* self, QColorSpace* colorSpace);
extern __declspec(dllexport) Primaries QColorSpace_Primaries(const QColorSpace* self);
extern __declspec(dllexport) TransferFunction QColorSpace_TransferFunction(const QColorSpace* self);
extern __declspec(dllexport) float QColorSpace_Gamma(const QColorSpace* self);
extern __declspec(dllexport) struct miqt_string QColorSpace_Description(const QColorSpace* self);
extern __declspec(dllexport) void QColorSpace_SetDescription(QColorSpace* self, struct miqt_string description);
extern __declspec(dllexport) void QColorSpace_SetTransferFunction(QColorSpace* self, TransferFunction transferFunction);
extern __declspec(dllexport) void QColorSpace_SetTransferFunctionWithTransferFunctionTable(QColorSpace* self, struct miqt_array /* of uint16_t */  transferFunctionTable);
extern __declspec(dllexport) void QColorSpace_SetTransferFunctions(QColorSpace* self, struct miqt_array /* of uint16_t */  redTransferFunctionTable, struct miqt_array /* of uint16_t */  greenTransferFunctionTable, struct miqt_array /* of uint16_t */  blueTransferFunctionTable);
extern __declspec(dllexport) QColorSpace* QColorSpace_WithTransferFunction(const QColorSpace* self, TransferFunction transferFunction);
extern __declspec(dllexport) QColorSpace* QColorSpace_WithTransferFunctionWithTransferFunctionTable(const QColorSpace* self, struct miqt_array /* of uint16_t */  transferFunctionTable);
extern __declspec(dllexport) QColorSpace* QColorSpace_WithTransferFunctions(const QColorSpace* self, struct miqt_array /* of uint16_t */  redTransferFunctionTable, struct miqt_array /* of uint16_t */  greenTransferFunctionTable, struct miqt_array /* of uint16_t */  blueTransferFunctionTable);
extern __declspec(dllexport) void QColorSpace_SetPrimaries(QColorSpace* self, Primaries primariesId);
extern __declspec(dllexport) void QColorSpace_SetPrimaries2(QColorSpace* self, QPointF* whitePoint, QPointF* redPoint, QPointF* greenPoint, QPointF* bluePoint);
extern __declspec(dllexport) void QColorSpace_SetWhitePoint(QColorSpace* self, QPointF* whitePoint);
extern __declspec(dllexport) QPointF* QColorSpace_WhitePoint(const QColorSpace* self);
extern __declspec(dllexport) TransformModel QColorSpace_TransformModel(const QColorSpace* self);
extern __declspec(dllexport) ColorModel QColorSpace_ColorModel(const QColorSpace* self);
extern __declspec(dllexport) void QColorSpace_Detach(QColorSpace* self);
extern __declspec(dllexport) bool QColorSpace_IsValid(const QColorSpace* self);
extern __declspec(dllexport) bool QColorSpace_IsValidTarget(const QColorSpace* self);
extern __declspec(dllexport) QColorSpace* QColorSpace_FromIccProfile(struct miqt_string iccProfile);
extern __declspec(dllexport) struct miqt_string QColorSpace_IccProfile(const QColorSpace* self);
extern __declspec(dllexport) QColorTransform* QColorSpace_TransformationToColorSpace(const QColorSpace* self, QColorSpace* colorspace);
extern __declspec(dllexport) void QColorSpace_SetTransferFunction2(QColorSpace* self, TransferFunction transferFunction, float gamma);
extern __declspec(dllexport) QColorSpace* QColorSpace_WithTransferFunction2(const QColorSpace* self, TransferFunction transferFunction, float gamma);
extern __declspec(dllexport) void QColorSpace_Delete(QColorSpace* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
