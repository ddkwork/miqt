// +build ignore

#include <QByteArray>
#include <QColorSpace>
#include <QColorTransform>
#include <QList>
#include <QPointF>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qcolorspace.h>
#include "gen_qcolorspace.h"

#ifndef _Bool
#define _Bool bool
#endif

void _GUID_Delete(_GUID* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<_GUID*>( self );
	} else {
		delete self;
	}
}

void type_info_Delete(type_info* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<type_info*>( self );
	} else {
		delete self;
	}
}

QColorSpace* QColorSpace_new() {
	return new QColorSpace();
}

QColorSpace* QColorSpace_new2(NamedColorSpace namedColorSpace) {
	return new QColorSpace(namedColorSpace);
}

QColorSpace* QColorSpace_new3(QPointF* whitePoint, TransferFunction transferFunction) {
	return new QColorSpace(*whitePoint, transferFunction);
}

QColorSpace* QColorSpace_new4(QPointF* whitePoint, struct miqt_array /* of uint16_t */  transferFunctionTable) {
	QList<uint16_t> transferFunctionTable_QList;
	transferFunctionTable_QList.reserve(transferFunctionTable.len);
	uint16_t* transferFunctionTable_arr = static_cast<uint16_t*>(transferFunctionTable.data);
	for(size_t i = 0; i < transferFunctionTable.len; ++i) {
		transferFunctionTable_QList.push_back(static_cast<uint16_t>(transferFunctionTable_arr[i]));
	}
	return new QColorSpace(*whitePoint, transferFunctionTable_QList);
}

QColorSpace* QColorSpace_new5(Primaries primaries, TransferFunction transferFunction) {
	return new QColorSpace(primaries, transferFunction);
}

QColorSpace* QColorSpace_new6(Primaries primaries, float gamma) {
	return new QColorSpace(primaries, static_cast<float>(gamma));
}

QColorSpace* QColorSpace_new7(Primaries primaries, struct miqt_array /* of uint16_t */  transferFunctionTable) {
	QList<uint16_t> transferFunctionTable_QList;
	transferFunctionTable_QList.reserve(transferFunctionTable.len);
	uint16_t* transferFunctionTable_arr = static_cast<uint16_t*>(transferFunctionTable.data);
	for(size_t i = 0; i < transferFunctionTable.len; ++i) {
		transferFunctionTable_QList.push_back(static_cast<uint16_t>(transferFunctionTable_arr[i]));
	}
	return new QColorSpace(primaries, transferFunctionTable_QList);
}

QColorSpace* QColorSpace_new8(QPointF* whitePoint, QPointF* redPoint, QPointF* greenPoint, QPointF* bluePoint, TransferFunction transferFunction) {
	return new QColorSpace(*whitePoint, *redPoint, *greenPoint, *bluePoint, transferFunction);
}

QColorSpace* QColorSpace_new9(QPointF* whitePoint, QPointF* redPoint, QPointF* greenPoint, QPointF* bluePoint, struct miqt_array /* of uint16_t */  transferFunctionTable) {
	QList<uint16_t> transferFunctionTable_QList;
	transferFunctionTable_QList.reserve(transferFunctionTable.len);
	uint16_t* transferFunctionTable_arr = static_cast<uint16_t*>(transferFunctionTable.data);
	for(size_t i = 0; i < transferFunctionTable.len; ++i) {
		transferFunctionTable_QList.push_back(static_cast<uint16_t>(transferFunctionTable_arr[i]));
	}
	return new QColorSpace(*whitePoint, *redPoint, *greenPoint, *bluePoint, transferFunctionTable_QList);
}

QColorSpace* QColorSpace_new10(QPointF* whitePoint, QPointF* redPoint, QPointF* greenPoint, QPointF* bluePoint, struct miqt_array /* of uint16_t */  redTransferFunctionTable, struct miqt_array /* of uint16_t */  greenTransferFunctionTable, struct miqt_array /* of uint16_t */  blueTransferFunctionTable) {
	QList<uint16_t> redTransferFunctionTable_QList;
	redTransferFunctionTable_QList.reserve(redTransferFunctionTable.len);
	uint16_t* redTransferFunctionTable_arr = static_cast<uint16_t*>(redTransferFunctionTable.data);
	for(size_t i = 0; i < redTransferFunctionTable.len; ++i) {
		redTransferFunctionTable_QList.push_back(static_cast<uint16_t>(redTransferFunctionTable_arr[i]));
	}
	QList<uint16_t> greenTransferFunctionTable_QList;
	greenTransferFunctionTable_QList.reserve(greenTransferFunctionTable.len);
	uint16_t* greenTransferFunctionTable_arr = static_cast<uint16_t*>(greenTransferFunctionTable.data);
	for(size_t i = 0; i < greenTransferFunctionTable.len; ++i) {
		greenTransferFunctionTable_QList.push_back(static_cast<uint16_t>(greenTransferFunctionTable_arr[i]));
	}
	QList<uint16_t> blueTransferFunctionTable_QList;
	blueTransferFunctionTable_QList.reserve(blueTransferFunctionTable.len);
	uint16_t* blueTransferFunctionTable_arr = static_cast<uint16_t*>(blueTransferFunctionTable.data);
	for(size_t i = 0; i < blueTransferFunctionTable.len; ++i) {
		blueTransferFunctionTable_QList.push_back(static_cast<uint16_t>(blueTransferFunctionTable_arr[i]));
	}
	return new QColorSpace(*whitePoint, *redPoint, *greenPoint, *bluePoint, redTransferFunctionTable_QList, greenTransferFunctionTable_QList, blueTransferFunctionTable_QList);
}

QColorSpace* QColorSpace_new11(QColorSpace* colorSpace) {
	return new QColorSpace(*colorSpace);
}

QColorSpace* QColorSpace_new12(QPointF* whitePoint, TransferFunction transferFunction, float gamma) {
	return new QColorSpace(*whitePoint, transferFunction, static_cast<float>(gamma));
}

QColorSpace* QColorSpace_new13(Primaries primaries, TransferFunction transferFunction, float gamma) {
	return new QColorSpace(primaries, transferFunction, static_cast<float>(gamma));
}

QColorSpace* QColorSpace_new14(QPointF* whitePoint, QPointF* redPoint, QPointF* greenPoint, QPointF* bluePoint, TransferFunction transferFunction, float gamma) {
	return new QColorSpace(*whitePoint, *redPoint, *greenPoint, *bluePoint, transferFunction, static_cast<float>(gamma));
}

void QColorSpace_OperatorAssign(QColorSpace* self, QColorSpace* colorSpace) {
	self->operator=(*colorSpace);
}

void QColorSpace_Swap(QColorSpace* self, QColorSpace* colorSpace) {
	self->swap(*colorSpace);
}

Primaries QColorSpace_Primaries(const QColorSpace* self) {
	return self->primaries();
}

TransferFunction QColorSpace_TransferFunction(const QColorSpace* self) {
	return self->transferFunction();
}

float QColorSpace_Gamma(const QColorSpace* self) {
	return self->gamma();
}

struct miqt_string QColorSpace_Description(const QColorSpace* self) {
	QString _ret = self->description();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QColorSpace_SetDescription(QColorSpace* self, struct miqt_string description) {
	QString description_QString = QString::fromUtf8(description.data, description.len);
	self->setDescription(description_QString);
}

void QColorSpace_SetTransferFunction(QColorSpace* self, TransferFunction transferFunction) {
	self->setTransferFunction(transferFunction);
}

void QColorSpace_SetTransferFunctionWithTransferFunctionTable(QColorSpace* self, struct miqt_array /* of uint16_t */  transferFunctionTable) {
	QList<uint16_t> transferFunctionTable_QList;
	transferFunctionTable_QList.reserve(transferFunctionTable.len);
	uint16_t* transferFunctionTable_arr = static_cast<uint16_t*>(transferFunctionTable.data);
	for(size_t i = 0; i < transferFunctionTable.len; ++i) {
		transferFunctionTable_QList.push_back(static_cast<uint16_t>(transferFunctionTable_arr[i]));
	}
	self->setTransferFunction(transferFunctionTable_QList);
}

void QColorSpace_SetTransferFunctions(QColorSpace* self, struct miqt_array /* of uint16_t */  redTransferFunctionTable, struct miqt_array /* of uint16_t */  greenTransferFunctionTable, struct miqt_array /* of uint16_t */  blueTransferFunctionTable) {
	QList<uint16_t> redTransferFunctionTable_QList;
	redTransferFunctionTable_QList.reserve(redTransferFunctionTable.len);
	uint16_t* redTransferFunctionTable_arr = static_cast<uint16_t*>(redTransferFunctionTable.data);
	for(size_t i = 0; i < redTransferFunctionTable.len; ++i) {
		redTransferFunctionTable_QList.push_back(static_cast<uint16_t>(redTransferFunctionTable_arr[i]));
	}
	QList<uint16_t> greenTransferFunctionTable_QList;
	greenTransferFunctionTable_QList.reserve(greenTransferFunctionTable.len);
	uint16_t* greenTransferFunctionTable_arr = static_cast<uint16_t*>(greenTransferFunctionTable.data);
	for(size_t i = 0; i < greenTransferFunctionTable.len; ++i) {
		greenTransferFunctionTable_QList.push_back(static_cast<uint16_t>(greenTransferFunctionTable_arr[i]));
	}
	QList<uint16_t> blueTransferFunctionTable_QList;
	blueTransferFunctionTable_QList.reserve(blueTransferFunctionTable.len);
	uint16_t* blueTransferFunctionTable_arr = static_cast<uint16_t*>(blueTransferFunctionTable.data);
	for(size_t i = 0; i < blueTransferFunctionTable.len; ++i) {
		blueTransferFunctionTable_QList.push_back(static_cast<uint16_t>(blueTransferFunctionTable_arr[i]));
	}
	self->setTransferFunctions(redTransferFunctionTable_QList, greenTransferFunctionTable_QList, blueTransferFunctionTable_QList);
}

QColorSpace* QColorSpace_WithTransferFunction(const QColorSpace* self, TransferFunction transferFunction) {
	return new QColorSpace(self->withTransferFunction(transferFunction));
}

QColorSpace* QColorSpace_WithTransferFunctionWithTransferFunctionTable(const QColorSpace* self, struct miqt_array /* of uint16_t */  transferFunctionTable) {
	QList<uint16_t> transferFunctionTable_QList;
	transferFunctionTable_QList.reserve(transferFunctionTable.len);
	uint16_t* transferFunctionTable_arr = static_cast<uint16_t*>(transferFunctionTable.data);
	for(size_t i = 0; i < transferFunctionTable.len; ++i) {
		transferFunctionTable_QList.push_back(static_cast<uint16_t>(transferFunctionTable_arr[i]));
	}
	return new QColorSpace(self->withTransferFunction(transferFunctionTable_QList));
}

QColorSpace* QColorSpace_WithTransferFunctions(const QColorSpace* self, struct miqt_array /* of uint16_t */  redTransferFunctionTable, struct miqt_array /* of uint16_t */  greenTransferFunctionTable, struct miqt_array /* of uint16_t */  blueTransferFunctionTable) {
	QList<uint16_t> redTransferFunctionTable_QList;
	redTransferFunctionTable_QList.reserve(redTransferFunctionTable.len);
	uint16_t* redTransferFunctionTable_arr = static_cast<uint16_t*>(redTransferFunctionTable.data);
	for(size_t i = 0; i < redTransferFunctionTable.len; ++i) {
		redTransferFunctionTable_QList.push_back(static_cast<uint16_t>(redTransferFunctionTable_arr[i]));
	}
	QList<uint16_t> greenTransferFunctionTable_QList;
	greenTransferFunctionTable_QList.reserve(greenTransferFunctionTable.len);
	uint16_t* greenTransferFunctionTable_arr = static_cast<uint16_t*>(greenTransferFunctionTable.data);
	for(size_t i = 0; i < greenTransferFunctionTable.len; ++i) {
		greenTransferFunctionTable_QList.push_back(static_cast<uint16_t>(greenTransferFunctionTable_arr[i]));
	}
	QList<uint16_t> blueTransferFunctionTable_QList;
	blueTransferFunctionTable_QList.reserve(blueTransferFunctionTable.len);
	uint16_t* blueTransferFunctionTable_arr = static_cast<uint16_t*>(blueTransferFunctionTable.data);
	for(size_t i = 0; i < blueTransferFunctionTable.len; ++i) {
		blueTransferFunctionTable_QList.push_back(static_cast<uint16_t>(blueTransferFunctionTable_arr[i]));
	}
	return new QColorSpace(self->withTransferFunctions(redTransferFunctionTable_QList, greenTransferFunctionTable_QList, blueTransferFunctionTable_QList));
}

void QColorSpace_SetPrimaries(QColorSpace* self, Primaries primariesId) {
	self->setPrimaries(primariesId);
}

void QColorSpace_SetPrimaries2(QColorSpace* self, QPointF* whitePoint, QPointF* redPoint, QPointF* greenPoint, QPointF* bluePoint) {
	self->setPrimaries(*whitePoint, *redPoint, *greenPoint, *bluePoint);
}

void QColorSpace_SetWhitePoint(QColorSpace* self, QPointF* whitePoint) {
	self->setWhitePoint(*whitePoint);
}

QPointF* QColorSpace_WhitePoint(const QColorSpace* self) {
	return new QPointF(self->whitePoint());
}

TransformModel QColorSpace_TransformModel(const QColorSpace* self) {
	return self->transformModel();
}

ColorModel QColorSpace_ColorModel(const QColorSpace* self) {
	return self->colorModel();
}

void QColorSpace_Detach(QColorSpace* self) {
	self->detach();
}

bool QColorSpace_IsValid(const QColorSpace* self) {
	return self->isValid();
}

bool QColorSpace_IsValidTarget(const QColorSpace* self) {
	return self->isValidTarget();
}

QColorSpace* QColorSpace_FromIccProfile(struct miqt_string iccProfile) {
	QByteArray iccProfile_QByteArray(iccProfile.data, iccProfile.len);
	return new QColorSpace(QColorSpace::fromIccProfile(iccProfile_QByteArray));
}

struct miqt_string QColorSpace_IccProfile(const QColorSpace* self) {
	QByteArray _qb = self->iccProfile();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

QColorTransform* QColorSpace_TransformationToColorSpace(const QColorSpace* self, QColorSpace* colorspace) {
	return new QColorTransform(self->transformationToColorSpace(*colorspace));
}

void QColorSpace_SetTransferFunction2(QColorSpace* self, TransferFunction transferFunction, float gamma) {
	self->setTransferFunction(transferFunction, static_cast<float>(gamma));
}

QColorSpace* QColorSpace_WithTransferFunction2(const QColorSpace* self, TransferFunction transferFunction, float gamma) {
	return new QColorSpace(self->withTransferFunction(transferFunction, static_cast<float>(gamma)));
}

void QColorSpace_Delete(QColorSpace* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QColorSpace*>( self );
	} else {
		delete self;
	}
}

