// +build ignore

#include <QPoint>
#define WORKAROUND_INNER_CLASS_DEFINITION_QTest__QTouchEventWidgetSequence
#include <QWidget>
#include <qtestsupport_widgets.h>
#include "gen_qtestsupport_widgets.h"

#ifndef _Bool
#define _Bool bool
#endif

class MiqtVirtualQTestQTouchEventWidgetSequence : public virtual QTest::QTouchEventWidgetSequence {
public:

	MiqtVirtualQTestQTouchEventWidgetSequence(const QTouchEventWidgetSequence& param1): QTest::QTouchEventWidgetSequence(param1) {};

	virtual ~MiqtVirtualQTestQTouchEventWidgetSequence() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Stationary = 0;

	// Subclass to allow providing a Go implementation
	virtual QTouchEventWidgetSequence& stationary(int touchId) override {
		if (handle__Stationary == 0) {
			return QTest__QTouchEventWidgetSequence::stationary(touchId);
		}
		
		int sigval1 = touchId;

		QTouchEventWidgetSequence* callback_return_value = miqt_exec_callback_QTest__QTouchEventWidgetSequence_Stationary(this, handle__Stationary, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QTouchEventWidgetSequence* virtualbase_Stationary(int touchId) {

		return &QTest__QTouchEventWidgetSequence::stationary(static_cast<int>(touchId));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Commit = 0;

	// Subclass to allow providing a Go implementation
	virtual bool commit(bool processEvents) override {
		if (handle__Commit == 0) {
			return QTest__QTouchEventWidgetSequence::commit(processEvents);
		}
		
		bool sigval1 = processEvents;

		bool callback_return_value = miqt_exec_callback_QTest__QTouchEventWidgetSequence_Commit(this, handle__Commit, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_Commit(bool processEvents) {

		return QTest__QTouchEventWidgetSequence::commit(processEvents);

	}

};

QTest__QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_new(const QTouchEventWidgetSequence* param1) {
	return new MiqtVirtualQTestQTouchEventWidgetSequence(*param1);
}

QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Press(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt) {
	return &self->press(static_cast<int>(touchId), *pt);
}

QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Move(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt) {
	return &self->move(static_cast<int>(touchId), *pt);
}

QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Release(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt) {
	return &self->release(static_cast<int>(touchId), *pt);
}

QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Stationary(QTest__QTouchEventWidgetSequence* self, int touchId) {
	return &self->stationary(static_cast<int>(touchId));
}

bool QTest__QTouchEventWidgetSequence_Commit(QTest__QTouchEventWidgetSequence* self, bool processEvents) {
	return self->commit(processEvents);
}

QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Press3(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt, QWidget* widget) {
	return &self->press(static_cast<int>(touchId), *pt, widget);
}

QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Move3(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt, QWidget* widget) {
	return &self->move(static_cast<int>(touchId), *pt, widget);
}

QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Release3(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt, QWidget* widget) {
	return &self->release(static_cast<int>(touchId), *pt, widget);
}

void QTest__QTouchEventWidgetSequence_override_virtual_Stationary(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQTestQTouchEventWidgetSequence*>( (QTest__QTouchEventWidgetSequence*)(self) )->handle__Stationary = slot;
}

QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_virtualbase_Stationary(void* self, int touchId) {
	return ( (MiqtVirtualQTestQTouchEventWidgetSequence*)(self) )->virtualbase_Stationary(touchId);
}

void QTest__QTouchEventWidgetSequence_override_virtual_Commit(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQTestQTouchEventWidgetSequence*>( (QTest__QTouchEventWidgetSequence*)(self) )->handle__Commit = slot;
}

bool QTest__QTouchEventWidgetSequence_virtualbase_Commit(void* self, bool processEvents) {
	return ( (MiqtVirtualQTestQTouchEventWidgetSequence*)(self) )->virtualbase_Commit(processEvents);
}

void QTest__QTouchEventWidgetSequence_Delete(QTest__QTouchEventWidgetSequence* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQTestQTouchEventWidgetSequence*>( self );
	} else {
		delete self;
	}
}
