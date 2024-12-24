// +build ignore

#include <QByteArray>
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QSvgRenderer>
#include <QSvgWidget>
#include <QWidget>
#include <qsvgwidget.h>
#include "gen_qsvgwidget.h"

class MiqtVirtualQSvgWidget : public virtual QSvgWidget {
public:

	MiqtVirtualQSvgWidget(QWidget* parent): QSvgWidget(parent) {};
	MiqtVirtualQSvgWidget(): QSvgWidget() {};
	MiqtVirtualQSvgWidget(const QString& file): QSvgWidget(file) {};
	MiqtVirtualQSvgWidget(const QString& file, QWidget* parent): QSvgWidget(file, parent) {};

	virtual ~MiqtVirtualQSvgWidget() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QSvgWidget::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QSvgWidget_MetaObject(const_cast<MiqtVirtualQSvgWidget*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QSvgWidget::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QSvgWidget::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QSvgWidget_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QSvgWidget::qt_metacast(param1);

	}

};

QSvgWidget* QSvgWidget_new(QWidget* parent) {
	return new MiqtVirtualQSvgWidget(parent);
}

QSvgWidget* QSvgWidget_new2() {
	return new MiqtVirtualQSvgWidget();
}

QSvgWidget* QSvgWidget_new3(struct miqt_string file) {
	QString file_QString = QString::fromUtf8(file.data, file.len);
	return new MiqtVirtualQSvgWidget(file_QString);
}

QSvgWidget* QSvgWidget_new4(struct miqt_string file, QWidget* parent) {
	QString file_QString = QString::fromUtf8(file.data, file.len);
	return new MiqtVirtualQSvgWidget(file_QString, parent);
}

void QSvgWidget_virtbase(QSvgWidget* src, QWidget** outptr_QWidget) {
	*outptr_QWidget = static_cast<QWidget*>(src);
}

QMetaObject* QSvgWidget_MetaObject(const QSvgWidget* self) {
	return (QMetaObject*) self->metaObject();
}

void* QSvgWidget_Metacast(QSvgWidget* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QSvgWidget_Tr(const char* s) {
	QString _ret = QSvgWidget::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QSvgRenderer* QSvgWidget_Renderer(const QSvgWidget* self) {
	return self->renderer();
}

QSize* QSvgWidget_SizeHint(const QSvgWidget* self) {
	return new QSize(self->sizeHint());
}

int QSvgWidget_Options(const QSvgWidget* self) {
	QtSvg::Options _ret = self->options();
	return static_cast<int>(_ret);
}

void QSvgWidget_SetOptions(QSvgWidget* self, int options) {
	self->setOptions(static_cast<QtSvg::Options>(options));
}

void QSvgWidget_Load(QSvgWidget* self, struct miqt_string file) {
	QString file_QString = QString::fromUtf8(file.data, file.len);
	self->load(file_QString);
}

void QSvgWidget_LoadWithContents(QSvgWidget* self, struct miqt_string contents) {
	QByteArray contents_QByteArray(contents.data, contents.len);
	self->load(contents_QByteArray);
}

struct miqt_string QSvgWidget_Tr2(const char* s, const char* c) {
	QString _ret = QSvgWidget::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QSvgWidget_Tr3(const char* s, const char* c, int n) {
	QString _ret = QSvgWidget::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QSvgWidget_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSvgWidget*>( (QSvgWidget*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QSvgWidget_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQSvgWidget*)(self) )->virtualbase_MetaObject();
}

void QSvgWidget_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSvgWidget*>( (QSvgWidget*)(self) )->handle__Metacast = slot;
}

void* QSvgWidget_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQSvgWidget*)(self) )->virtualbase_Metacast(param1);
}

void QSvgWidget_Delete(QSvgWidget* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQSvgWidget*>( self );
	} else {
		delete self;
	}
}

