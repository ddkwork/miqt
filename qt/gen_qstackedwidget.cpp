// +build ignore

#include <QEvent>
#include <QFrame>
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QStackedWidget>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include <qstackedwidget.h>
#include "gen_qstackedwidget.h"
class MiqtVirtualQStackedWidget : public virtual QStackedWidget {
public:
MiqtVirtualQStackedWidget(QWidget* parent): QStackedWidget(parent) {};
MiqtVirtualQStackedWidget(): QStackedWidget() {};

virtual ~MiqtVirtualQStackedWidget() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QStackedWidget::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QStackedWidget_MetaObject(const_cast<MiqtVirtualQStackedWidget*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QStackedWidget::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QStackedWidget::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QStackedWidget_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QStackedWidget::qt_metacast(param1);

	}
};
QStackedWidget* QStackedWidget_new(QWidget* parent) {
return new MiqtVirtualQStackedWidget(parent);
}
QStackedWidget* QStackedWidget_new2() {
return new MiqtVirtualQStackedWidget();
}
void QStackedWidget_virtbase(QStackedWidget* src
, QFrame** outptr_QFrame
) {
*outptr_QFrame = static_cast<QFrame*>(src);
}
QMetaObject* QStackedWidget_MetaObject(const QStackedWidget* self) {
	return (QMetaObject*) self->metaObject();
}
void* QStackedWidget_Metacast(QStackedWidget* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QStackedWidget_Tr(const char* s) {
	QString _ret = QStackedWidget::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
int QStackedWidget_AddWidget(QStackedWidget* self, QWidget* w) {
	return self->addWidget(w);
}
int QStackedWidget_InsertWidget(QStackedWidget* self, int index, QWidget* w) {
	return self->insertWidget(static_cast<int>(index), w);
}
void QStackedWidget_RemoveWidget(QStackedWidget* self, QWidget* w) {
	self->removeWidget(w);
}
QWidget* QStackedWidget_CurrentWidget(const QStackedWidget* self) {
	return self->currentWidget();
}
int QStackedWidget_CurrentIndex(const QStackedWidget* self) {
	return self->currentIndex();
}
int QStackedWidget_IndexOf(const QStackedWidget* self, QWidget* param1) {
	return self->indexOf(param1);
}
QWidget* QStackedWidget_Widget(const QStackedWidget* self, int param1) {
	return self->widget(static_cast<int>(param1));
}
int QStackedWidget_Count(const QStackedWidget* self) {
	return self->count();
}
void QStackedWidget_SetCurrentIndex(QStackedWidget* self, int index) {
	self->setCurrentIndex(static_cast<int>(index));
}
void QStackedWidget_SetCurrentWidget(QStackedWidget* self, QWidget* w) {
	self->setCurrentWidget(w);
}
void QStackedWidget_CurrentChanged(QStackedWidget* self, int param1) {
	self->currentChanged(static_cast<int>(param1));
}
void QStackedWidget_connect_CurrentChanged(QStackedWidget* self, intptr_t slot) {
	MiqtVirtualQStackedWidget::connect(self, static_cast<void (QStackedWidget::*)(int)>(&QStackedWidget::currentChanged), self, [=](int param1) {
		int sigval1 = param1;
		miqt_exec_callback_QStackedWidget_CurrentChanged(slot, sigval1);
	});
}
void QStackedWidget_WidgetRemoved(QStackedWidget* self, int index) {
	self->widgetRemoved(static_cast<int>(index));
}
void QStackedWidget_connect_WidgetRemoved(QStackedWidget* self, intptr_t slot) {
	MiqtVirtualQStackedWidget::connect(self, static_cast<void (QStackedWidget::*)(int)>(&QStackedWidget::widgetRemoved), self, [=](int index) {
		int sigval1 = index;
		miqt_exec_callback_QStackedWidget_WidgetRemoved(slot, sigval1);
	});
}
void QStackedWidget_WidgetAdded(QStackedWidget* self, int index) {
	self->widgetAdded(static_cast<int>(index));
}
void QStackedWidget_connect_WidgetAdded(QStackedWidget* self, intptr_t slot) {
	MiqtVirtualQStackedWidget::connect(self, static_cast<void (QStackedWidget::*)(int)>(&QStackedWidget::widgetAdded), self, [=](int index) {
		int sigval1 = index;
		miqt_exec_callback_QStackedWidget_WidgetAdded(slot, sigval1);
	});
}
struct miqt_string QStackedWidget_Tr2(const char* s, const char* c) {
	QString _ret = QStackedWidget::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QStackedWidget_Tr3(const char* s, const char* c, int n) {
	QString _ret = QStackedWidget::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QStackedWidget_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQStackedWidget*>( (QStackedWidget*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QStackedWidget_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQStackedWidget*)(self) )->virtualbase_MetaObject();
}
void QStackedWidget_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQStackedWidget*>( (QStackedWidget*)(self) )->handle__Metacast = slot;
}
void* QStackedWidget_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQStackedWidget*)(self) )->virtualbase_Metacast(param1);
}
void QStackedWidget_Delete(QStackedWidget* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQStackedWidget*>( self );
	} else {
		delete self;
	}
}
