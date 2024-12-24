// +build ignore

#include <QBluetoothPermission>
#include <QCalendarPermission>
#include <QCameraPermission>
#include <QContactsPermission>
#include <QLocationPermission>
#include <QMetaType>
#include <QMicrophonePermission>
#include <QPermission>
#include <qpermissions.h>
#include "gen_qpermissions.h"

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

QPermission* QPermission_new() {
	return new QPermission();
}

QPermission* QPermission_new2(QPermission* param1) {
	return new QPermission(*param1);
}

int QPermission_Status(const QPermission* self) {
	Qt::PermissionStatus _ret = self->status();
	return static_cast<int>(_ret);
}

QMetaType* QPermission_Type(const QPermission* self) {
	return new QMetaType(self->type());
}

void QPermission_Delete(QPermission* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QPermission*>( self );
	} else {
		delete self;
	}
}

QLocationPermission* QLocationPermission_new() {
	return new QLocationPermission();
}

QLocationPermission* QLocationPermission_new2(QLocationPermission* other) {
	return new QLocationPermission(*other);
}

void QLocationPermission_SetAccuracy(QLocationPermission* self, Accuracy accuracy) {
	self->setAccuracy(accuracy);
}

Accuracy QLocationPermission_Accuracy(const QLocationPermission* self) {
	return self->accuracy();
}

void QLocationPermission_SetAvailability(QLocationPermission* self, Availability availability) {
	self->setAvailability(availability);
}

Availability QLocationPermission_Availability(const QLocationPermission* self) {
	return self->availability();
}

void QLocationPermission_OperatorAssign(QLocationPermission* self, QLocationPermission* other) {
	self->operator=(*other);
}

void QLocationPermission_connect_OperatorAssign(QLocationPermission* self, intptr_t slot) {
	QLocationPermission::connect(self, static_cast<void (QLocationPermission::*)(const QLocationPermission&)>(&QLocationPermission::operator=), self, [=](const QLocationPermission& other) {
		const QLocationPermission& other_ret = other;
		// Cast returned reference into pointer
		QLocationPermission* sigval1 = const_cast<QLocationPermission*>(&other_ret);
		miqt_exec_callback_QLocationPermission_OperatorAssign(slot, sigval1);
	});
}

void QLocationPermission_Swap(QLocationPermission* self, QLocationPermission* other) {
	self->swap(*other);
}

void QLocationPermission_connect_Swap(QLocationPermission* self, intptr_t slot) {
	QLocationPermission::connect(self, static_cast<void (QLocationPermission::*)(QLocationPermission&)>(&QLocationPermission::swap), self, [=](QLocationPermission& other) {
		QLocationPermission& other_ret = other;
		// Cast returned reference into pointer
		QLocationPermission* sigval1 = &other_ret;
		miqt_exec_callback_QLocationPermission_Swap(slot, sigval1);
	});
}

void QLocationPermission_Delete(QLocationPermission* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QLocationPermission*>( self );
	} else {
		delete self;
	}
}

QCalendarPermission* QCalendarPermission_new() {
	return new QCalendarPermission();
}

QCalendarPermission* QCalendarPermission_new2(QCalendarPermission* other) {
	return new QCalendarPermission(*other);
}

void QCalendarPermission_SetAccessMode(QCalendarPermission* self, AccessMode mode) {
	self->setAccessMode(mode);
}

AccessMode QCalendarPermission_AccessMode(const QCalendarPermission* self) {
	return self->accessMode();
}

void QCalendarPermission_OperatorAssign(QCalendarPermission* self, QCalendarPermission* other) {
	self->operator=(*other);
}

void QCalendarPermission_connect_OperatorAssign(QCalendarPermission* self, intptr_t slot) {
	QCalendarPermission::connect(self, static_cast<void (QCalendarPermission::*)(const QCalendarPermission&)>(&QCalendarPermission::operator=), self, [=](const QCalendarPermission& other) {
		const QCalendarPermission& other_ret = other;
		// Cast returned reference into pointer
		QCalendarPermission* sigval1 = const_cast<QCalendarPermission*>(&other_ret);
		miqt_exec_callback_QCalendarPermission_OperatorAssign(slot, sigval1);
	});
}

void QCalendarPermission_Swap(QCalendarPermission* self, QCalendarPermission* other) {
	self->swap(*other);
}

void QCalendarPermission_connect_Swap(QCalendarPermission* self, intptr_t slot) {
	QCalendarPermission::connect(self, static_cast<void (QCalendarPermission::*)(QCalendarPermission&)>(&QCalendarPermission::swap), self, [=](QCalendarPermission& other) {
		QCalendarPermission& other_ret = other;
		// Cast returned reference into pointer
		QCalendarPermission* sigval1 = &other_ret;
		miqt_exec_callback_QCalendarPermission_Swap(slot, sigval1);
	});
}

void QCalendarPermission_Delete(QCalendarPermission* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QCalendarPermission*>( self );
	} else {
		delete self;
	}
}

QContactsPermission* QContactsPermission_new() {
	return new QContactsPermission();
}

QContactsPermission* QContactsPermission_new2(QContactsPermission* other) {
	return new QContactsPermission(*other);
}

void QContactsPermission_SetAccessMode(QContactsPermission* self, AccessMode mode) {
	self->setAccessMode(mode);
}

AccessMode QContactsPermission_AccessMode(const QContactsPermission* self) {
	return self->accessMode();
}

void QContactsPermission_OperatorAssign(QContactsPermission* self, QContactsPermission* other) {
	self->operator=(*other);
}

void QContactsPermission_connect_OperatorAssign(QContactsPermission* self, intptr_t slot) {
	QContactsPermission::connect(self, static_cast<void (QContactsPermission::*)(const QContactsPermission&)>(&QContactsPermission::operator=), self, [=](const QContactsPermission& other) {
		const QContactsPermission& other_ret = other;
		// Cast returned reference into pointer
		QContactsPermission* sigval1 = const_cast<QContactsPermission*>(&other_ret);
		miqt_exec_callback_QContactsPermission_OperatorAssign(slot, sigval1);
	});
}

void QContactsPermission_Swap(QContactsPermission* self, QContactsPermission* other) {
	self->swap(*other);
}

void QContactsPermission_connect_Swap(QContactsPermission* self, intptr_t slot) {
	QContactsPermission::connect(self, static_cast<void (QContactsPermission::*)(QContactsPermission&)>(&QContactsPermission::swap), self, [=](QContactsPermission& other) {
		QContactsPermission& other_ret = other;
		// Cast returned reference into pointer
		QContactsPermission* sigval1 = &other_ret;
		miqt_exec_callback_QContactsPermission_Swap(slot, sigval1);
	});
}

void QContactsPermission_Delete(QContactsPermission* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QContactsPermission*>( self );
	} else {
		delete self;
	}
}

QBluetoothPermission* QBluetoothPermission_new() {
	return new QBluetoothPermission();
}

QBluetoothPermission* QBluetoothPermission_new2(QBluetoothPermission* other) {
	return new QBluetoothPermission(*other);
}

void QBluetoothPermission_SetCommunicationModes(QBluetoothPermission* self, CommunicationModes modes) {
	self->setCommunicationModes(modes);
}

CommunicationModes QBluetoothPermission_CommunicationModes(const QBluetoothPermission* self) {
	return self->communicationModes();
}

void QBluetoothPermission_OperatorAssign(QBluetoothPermission* self, QBluetoothPermission* other) {
	self->operator=(*other);
}

void QBluetoothPermission_connect_OperatorAssign(QBluetoothPermission* self, intptr_t slot) {
	QBluetoothPermission::connect(self, static_cast<void (QBluetoothPermission::*)(const QBluetoothPermission&)>(&QBluetoothPermission::operator=), self, [=](const QBluetoothPermission& other) {
		const QBluetoothPermission& other_ret = other;
		// Cast returned reference into pointer
		QBluetoothPermission* sigval1 = const_cast<QBluetoothPermission*>(&other_ret);
		miqt_exec_callback_QBluetoothPermission_OperatorAssign(slot, sigval1);
	});
}

void QBluetoothPermission_Swap(QBluetoothPermission* self, QBluetoothPermission* other) {
	self->swap(*other);
}

void QBluetoothPermission_connect_Swap(QBluetoothPermission* self, intptr_t slot) {
	QBluetoothPermission::connect(self, static_cast<void (QBluetoothPermission::*)(QBluetoothPermission&)>(&QBluetoothPermission::swap), self, [=](QBluetoothPermission& other) {
		QBluetoothPermission& other_ret = other;
		// Cast returned reference into pointer
		QBluetoothPermission* sigval1 = &other_ret;
		miqt_exec_callback_QBluetoothPermission_Swap(slot, sigval1);
	});
}

void QBluetoothPermission_Delete(QBluetoothPermission* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QBluetoothPermission*>( self );
	} else {
		delete self;
	}
}

QCameraPermission* QCameraPermission_new() {
	return new QCameraPermission();
}

QCameraPermission* QCameraPermission_new2(QCameraPermission* other) {
	return new QCameraPermission(*other);
}

void QCameraPermission_OperatorAssign(QCameraPermission* self, QCameraPermission* other) {
	self->operator=(*other);
}

void QCameraPermission_connect_OperatorAssign(QCameraPermission* self, intptr_t slot) {
	QCameraPermission::connect(self, static_cast<void (QCameraPermission::*)(const QCameraPermission&)>(&QCameraPermission::operator=), self, [=](const QCameraPermission& other) {
		const QCameraPermission& other_ret = other;
		// Cast returned reference into pointer
		QCameraPermission* sigval1 = const_cast<QCameraPermission*>(&other_ret);
		miqt_exec_callback_QCameraPermission_OperatorAssign(slot, sigval1);
	});
}

void QCameraPermission_Swap(QCameraPermission* self, QCameraPermission* other) {
	self->swap(*other);
}

void QCameraPermission_connect_Swap(QCameraPermission* self, intptr_t slot) {
	QCameraPermission::connect(self, static_cast<void (QCameraPermission::*)(QCameraPermission&)>(&QCameraPermission::swap), self, [=](QCameraPermission& other) {
		QCameraPermission& other_ret = other;
		// Cast returned reference into pointer
		QCameraPermission* sigval1 = &other_ret;
		miqt_exec_callback_QCameraPermission_Swap(slot, sigval1);
	});
}

void QCameraPermission_Delete(QCameraPermission* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QCameraPermission*>( self );
	} else {
		delete self;
	}
}

QMicrophonePermission* QMicrophonePermission_new() {
	return new QMicrophonePermission();
}

QMicrophonePermission* QMicrophonePermission_new2(QMicrophonePermission* other) {
	return new QMicrophonePermission(*other);
}

void QMicrophonePermission_OperatorAssign(QMicrophonePermission* self, QMicrophonePermission* other) {
	self->operator=(*other);
}

void QMicrophonePermission_connect_OperatorAssign(QMicrophonePermission* self, intptr_t slot) {
	QMicrophonePermission::connect(self, static_cast<void (QMicrophonePermission::*)(const QMicrophonePermission&)>(&QMicrophonePermission::operator=), self, [=](const QMicrophonePermission& other) {
		const QMicrophonePermission& other_ret = other;
		// Cast returned reference into pointer
		QMicrophonePermission* sigval1 = const_cast<QMicrophonePermission*>(&other_ret);
		miqt_exec_callback_QMicrophonePermission_OperatorAssign(slot, sigval1);
	});
}

void QMicrophonePermission_Swap(QMicrophonePermission* self, QMicrophonePermission* other) {
	self->swap(*other);
}

void QMicrophonePermission_connect_Swap(QMicrophonePermission* self, intptr_t slot) {
	QMicrophonePermission::connect(self, static_cast<void (QMicrophonePermission::*)(QMicrophonePermission&)>(&QMicrophonePermission::swap), self, [=](QMicrophonePermission& other) {
		QMicrophonePermission& other_ret = other;
		// Cast returned reference into pointer
		QMicrophonePermission* sigval1 = &other_ret;
		miqt_exec_callback_QMicrophonePermission_Swap(slot, sigval1);
	});
}

void QMicrophonePermission_Delete(QMicrophonePermission* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QMicrophonePermission*>( self );
	} else {
		delete self;
	}
}

