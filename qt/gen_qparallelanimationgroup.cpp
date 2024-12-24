// +build ignore

#include <QAbstractAnimation>
#include <QAnimationGroup>
#include <QEvent>
#include <QMetaObject>
#include <QObject>
#include <QParallelAnimationGroup>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qparallelanimationgroup.h>
#include "gen_qparallelanimationgroup.h"

class MiqtVirtualQParallelAnimationGroup : public virtual QParallelAnimationGroup {
public:

	MiqtVirtualQParallelAnimationGroup(): QParallelAnimationGroup() {};
	MiqtVirtualQParallelAnimationGroup(QObject* parent): QParallelAnimationGroup(parent) {};

	virtual ~MiqtVirtualQParallelAnimationGroup() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QParallelAnimationGroup::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QParallelAnimationGroup_MetaObject(const_cast<MiqtVirtualQParallelAnimationGroup*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QParallelAnimationGroup::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QParallelAnimationGroup::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QParallelAnimationGroup_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QParallelAnimationGroup::qt_metacast(param1);

	}

};

QParallelAnimationGroup* QParallelAnimationGroup_new() {
	return new MiqtVirtualQParallelAnimationGroup();
}

QParallelAnimationGroup* QParallelAnimationGroup_new2(QObject* parent) {
	return new MiqtVirtualQParallelAnimationGroup(parent);
}

void QParallelAnimationGroup_virtbase(QParallelAnimationGroup* src, QAnimationGroup** outptr_QAnimationGroup) {
	*outptr_QAnimationGroup = static_cast<QAnimationGroup*>(src);
}

QMetaObject* QParallelAnimationGroup_MetaObject(const QParallelAnimationGroup* self) {
	return (QMetaObject*) self->metaObject();
}

void* QParallelAnimationGroup_Metacast(QParallelAnimationGroup* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QParallelAnimationGroup_Tr(const char* s) {
	QString _ret = QParallelAnimationGroup::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QParallelAnimationGroup_Duration(const QParallelAnimationGroup* self) {
	return self->duration();
}

struct miqt_string QParallelAnimationGroup_Tr2(const char* s, const char* c) {
	QString _ret = QParallelAnimationGroup::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QParallelAnimationGroup_Tr3(const char* s, const char* c, int n) {
	QString _ret = QParallelAnimationGroup::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QParallelAnimationGroup_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQParallelAnimationGroup*>( (QParallelAnimationGroup*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QParallelAnimationGroup_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQParallelAnimationGroup*)(self) )->virtualbase_MetaObject();
}

void QParallelAnimationGroup_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQParallelAnimationGroup*>( (QParallelAnimationGroup*)(self) )->handle__Metacast = slot;
}

void* QParallelAnimationGroup_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQParallelAnimationGroup*)(self) )->virtualbase_Metacast(param1);
}

void QParallelAnimationGroup_Delete(QParallelAnimationGroup* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQParallelAnimationGroup*>( self );
	} else {
		delete self;
	}
}

