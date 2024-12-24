// +build ignore

#include <QAbstractFileIconProvider>
#include <QFileInfo>
#include <QIcon>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qabstractfileiconprovider.h>
#include "gen_qabstractfileiconprovider.h"

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

class MiqtVirtualQAbstractFileIconProvider : public virtual QAbstractFileIconProvider {
public:

	MiqtVirtualQAbstractFileIconProvider(): QAbstractFileIconProvider() {};

	virtual ~MiqtVirtualQAbstractFileIconProvider() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Icon = 0;

	// Subclass to allow providing a Go implementation
	virtual QIcon icon(IconType param1) const override {
		if (handle__Icon == 0) {
			return QAbstractFileIconProvider::icon(param1);
		}
		
		IconType sigval1 = param1;

		QIcon* callback_return_value = miqt_exec_callback_QAbstractFileIconProvider_Icon(const_cast<MiqtVirtualQAbstractFileIconProvider*>(this), handle__Icon, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QIcon* virtualbase_Icon(IconType param1) const {

		return new QIcon(QAbstractFileIconProvider::icon(param1));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__IconWithQFileInfo = 0;

	// Subclass to allow providing a Go implementation
	virtual QIcon icon(const QFileInfo& param1) const override {
		if (handle__IconWithQFileInfo == 0) {
			return QAbstractFileIconProvider::icon(param1);
		}
		
		const QFileInfo& param1_ret = param1;
		// Cast returned reference into pointer
		QFileInfo* sigval1 = const_cast<QFileInfo*>(&param1_ret);

		QIcon* callback_return_value = miqt_exec_callback_QAbstractFileIconProvider_IconWithQFileInfo(const_cast<MiqtVirtualQAbstractFileIconProvider*>(this), handle__IconWithQFileInfo, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QIcon* virtualbase_IconWithQFileInfo(QFileInfo* param1) const {

		return new QIcon(QAbstractFileIconProvider::icon(*param1));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Type = 0;

	// Subclass to allow providing a Go implementation
	virtual QString type(const QFileInfo& param1) const override {
		if (handle__Type == 0) {
			return QAbstractFileIconProvider::type(param1);
		}
		
		const QFileInfo& param1_ret = param1;
		// Cast returned reference into pointer
		QFileInfo* sigval1 = const_cast<QFileInfo*>(&param1_ret);

		struct miqt_string callback_return_value = miqt_exec_callback_QAbstractFileIconProvider_Type(const_cast<MiqtVirtualQAbstractFileIconProvider*>(this), handle__Type, sigval1);
		QString callback_return_value_QString = QString::fromUtf8(callback_return_value.data, callback_return_value.len);

		return callback_return_value_QString;
	}

	// Wrapper to allow calling protected method
	struct miqt_string virtualbase_Type(QFileInfo* param1) const {

		QString _ret = QAbstractFileIconProvider::type(*param1);
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _b = _ret.toUtf8();
		struct miqt_string _ms;
		_ms.len = _b.length();
		_ms.data = static_cast<char*>(malloc(_ms.len));
		memcpy(_ms.data, _b.data(), _ms.len);
		return _ms;

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetOptions = 0;

	// Subclass to allow providing a Go implementation
	virtual void setOptions(Options options) override {
		if (handle__SetOptions == 0) {
			QAbstractFileIconProvider::setOptions(options);
			return;
		}
		
		Options sigval1 = options;

		miqt_exec_callback_QAbstractFileIconProvider_SetOptions(this, handle__SetOptions, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetOptions(Options options) {

		QAbstractFileIconProvider::setOptions(options);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Options = 0;

	// Subclass to allow providing a Go implementation
	virtual Options options() const override {
		if (handle__Options == 0) {
			return QAbstractFileIconProvider::options();
		}
		

		Options callback_return_value = miqt_exec_callback_QAbstractFileIconProvider_Options(const_cast<MiqtVirtualQAbstractFileIconProvider*>(this), handle__Options);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	Options virtualbase_Options() const {

		return QAbstractFileIconProvider::options();

	}

};

QAbstractFileIconProvider* QAbstractFileIconProvider_new() {
	return new MiqtVirtualQAbstractFileIconProvider();
}

QIcon* QAbstractFileIconProvider_Icon(const QAbstractFileIconProvider* self, IconType param1) {
	return new QIcon(self->icon(param1));
}

QIcon* QAbstractFileIconProvider_IconWithQFileInfo(const QAbstractFileIconProvider* self, QFileInfo* param1) {
	return new QIcon(self->icon(*param1));
}

struct miqt_string QAbstractFileIconProvider_Type(const QAbstractFileIconProvider* self, QFileInfo* param1) {
	QString _ret = self->type(*param1);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QAbstractFileIconProvider_SetOptions(QAbstractFileIconProvider* self, Options options) {
	self->setOptions(options);
}

Options QAbstractFileIconProvider_Options(const QAbstractFileIconProvider* self) {
	return self->options();
}

void QAbstractFileIconProvider_override_virtual_Icon(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAbstractFileIconProvider*>( (QAbstractFileIconProvider*)(self) )->handle__Icon = slot;
}

QIcon* QAbstractFileIconProvider_virtualbase_Icon(const void* self, IconType param1) {
	return ( (const MiqtVirtualQAbstractFileIconProvider*)(self) )->virtualbase_Icon(param1);
}

void QAbstractFileIconProvider_override_virtual_IconWithQFileInfo(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAbstractFileIconProvider*>( (QAbstractFileIconProvider*)(self) )->handle__IconWithQFileInfo = slot;
}

QIcon* QAbstractFileIconProvider_virtualbase_IconWithQFileInfo(const void* self, QFileInfo* param1) {
	return ( (const MiqtVirtualQAbstractFileIconProvider*)(self) )->virtualbase_IconWithQFileInfo(param1);
}

void QAbstractFileIconProvider_override_virtual_Type(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAbstractFileIconProvider*>( (QAbstractFileIconProvider*)(self) )->handle__Type = slot;
}

struct miqt_string QAbstractFileIconProvider_virtualbase_Type(const void* self, QFileInfo* param1) {
	return ( (const MiqtVirtualQAbstractFileIconProvider*)(self) )->virtualbase_Type(param1);
}

void QAbstractFileIconProvider_override_virtual_SetOptions(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAbstractFileIconProvider*>( (QAbstractFileIconProvider*)(self) )->handle__SetOptions = slot;
}

void QAbstractFileIconProvider_virtualbase_SetOptions(void* self, Options options) {
	( (MiqtVirtualQAbstractFileIconProvider*)(self) )->virtualbase_SetOptions(options);
}

void QAbstractFileIconProvider_override_virtual_Options(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAbstractFileIconProvider*>( (QAbstractFileIconProvider*)(self) )->handle__Options = slot;
}

Options QAbstractFileIconProvider_virtualbase_Options(const void* self) {
	return ( (const MiqtVirtualQAbstractFileIconProvider*)(self) )->virtualbase_Options();
}

void QAbstractFileIconProvider_Delete(QAbstractFileIconProvider* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQAbstractFileIconProvider*>( self );
	} else {
		delete self;
	}
}

