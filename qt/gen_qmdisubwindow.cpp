// +build ignore

#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QEvent>
#include <QFocusEvent>
#include <QHideEvent>
#include <QKeyEvent>
#include <QMdiArea>
#include <QMdiSubWindow>
#include <QMenu>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTimerEvent>
#include <QWidget>
#include <qmdisubwindow.h>
#include "gen_qmdisubwindow.h"
class MiqtVirtualQMdiSubWindow : public virtual QMdiSubWindow {
public:
MiqtVirtualQMdiSubWindow(QWidget* parent): QMdiSubWindow(parent) {};
MiqtVirtualQMdiSubWindow(): QMdiSubWindow() {};
MiqtVirtualQMdiSubWindow(QWidget* parent, Qt::WindowFlags flags): QMdiSubWindow(parent, flags) {};

virtual ~MiqtVirtualQMdiSubWindow() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QMdiSubWindow::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QMdiSubWindow_MetaObject(const_cast<MiqtVirtualQMdiSubWindow*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QMdiSubWindow::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QMdiSubWindow::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QMdiSubWindow_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QMdiSubWindow::qt_metacast(param1);

	}
};
QMdiSubWindow* QMdiSubWindow_new(QWidget* parent) {
return new MiqtVirtualQMdiSubWindow(parent);
}
QMdiSubWindow* QMdiSubWindow_new2() {
return new MiqtVirtualQMdiSubWindow();
}
QMdiSubWindow* QMdiSubWindow_new3(QWidget* parent, int flags) {
return new MiqtVirtualQMdiSubWindow(parent, static_cast<Qt::WindowFlags>(flags));
}
void QMdiSubWindow_virtbase(QMdiSubWindow* src
, QWidget** outptr_QWidget
) {
*outptr_QWidget = static_cast<QWidget*>(src);
}
QMetaObject* QMdiSubWindow_MetaObject(const QMdiSubWindow* self) {
	return (QMetaObject*) self->metaObject();
}
void* QMdiSubWindow_Metacast(QMdiSubWindow* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QMdiSubWindow_Tr(const char* s) {
	QString _ret = QMdiSubWindow::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QSize* QMdiSubWindow_SizeHint(const QMdiSubWindow* self) {
	return new QSize(self->sizeHint());
}
QSize* QMdiSubWindow_MinimumSizeHint(const QMdiSubWindow* self) {
	return new QSize(self->minimumSizeHint());
}
void QMdiSubWindow_SetWidget(QMdiSubWindow* self, QWidget* widget) {
	self->setWidget(widget);
}
QWidget* QMdiSubWindow_Widget(const QMdiSubWindow* self) {
	return self->widget();
}
QWidget* QMdiSubWindow_MaximizedButtonsWidget(const QMdiSubWindow* self) {
	return self->maximizedButtonsWidget();
}
QWidget* QMdiSubWindow_MaximizedSystemMenuIconWidget(const QMdiSubWindow* self) {
	return self->maximizedSystemMenuIconWidget();
}
bool QMdiSubWindow_IsShaded(const QMdiSubWindow* self) {
	return self->isShaded();
}
void QMdiSubWindow_SetOption(QMdiSubWindow* self, SubWindowOption option) {
	self->setOption(option);
}
bool QMdiSubWindow_TestOption(const QMdiSubWindow* self, SubWindowOption param1) {
	return self->testOption(param1);
}
void QMdiSubWindow_SetKeyboardSingleStep(QMdiSubWindow* self, int step) {
	self->setKeyboardSingleStep(static_cast<int>(step));
}
int QMdiSubWindow_KeyboardSingleStep(const QMdiSubWindow* self) {
	return self->keyboardSingleStep();
}
void QMdiSubWindow_SetKeyboardPageStep(QMdiSubWindow* self, int step) {
	self->setKeyboardPageStep(static_cast<int>(step));
}
int QMdiSubWindow_KeyboardPageStep(const QMdiSubWindow* self) {
	return self->keyboardPageStep();
}
void QMdiSubWindow_SetSystemMenu(QMdiSubWindow* self, QMenu* systemMenu) {
	self->setSystemMenu(systemMenu);
}
QMenu* QMdiSubWindow_SystemMenu(const QMdiSubWindow* self) {
	return self->systemMenu();
}
QMdiArea* QMdiSubWindow_MdiArea(const QMdiSubWindow* self) {
	return self->mdiArea();
}
void QMdiSubWindow_WindowStateChanged(QMdiSubWindow* self, int oldState, int newState) {
	self->windowStateChanged(static_cast<Qt::WindowStates>(oldState), static_cast<Qt::WindowStates>(newState));
}
void QMdiSubWindow_connect_WindowStateChanged(QMdiSubWindow* self, intptr_t slot) {
	MiqtVirtualQMdiSubWindow::connect(self, static_cast<void (QMdiSubWindow::*)(Qt::WindowStates, Qt::WindowStates)>(&QMdiSubWindow::windowStateChanged), self, [=](Qt::WindowStates oldState, Qt::WindowStates newState) {
		Qt::WindowStates oldState_ret = oldState;
		int sigval1 = static_cast<int>(oldState_ret);
		Qt::WindowStates newState_ret = newState;
		int sigval2 = static_cast<int>(newState_ret);
		miqt_exec_callback_QMdiSubWindow_WindowStateChanged(slot, sigval1, sigval2);
	});
}
void QMdiSubWindow_AboutToActivate(QMdiSubWindow* self) {
	self->aboutToActivate();
}
void QMdiSubWindow_connect_AboutToActivate(QMdiSubWindow* self, intptr_t slot) {
	MiqtVirtualQMdiSubWindow::connect(self, static_cast<void (QMdiSubWindow::*)()>(&QMdiSubWindow::aboutToActivate), self, [=]() {
		miqt_exec_callback_QMdiSubWindow_AboutToActivate(slot);
	});
}
void QMdiSubWindow_ShowSystemMenu(QMdiSubWindow* self) {
	self->showSystemMenu();
}
void QMdiSubWindow_ShowShaded(QMdiSubWindow* self) {
	self->showShaded();
}
struct miqt_string QMdiSubWindow_Tr2(const char* s, const char* c) {
	QString _ret = QMdiSubWindow::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QMdiSubWindow_Tr3(const char* s, const char* c, int n) {
	QString _ret = QMdiSubWindow::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QMdiSubWindow_SetOption2(QMdiSubWindow* self, SubWindowOption option, bool on) {
	self->setOption(option, on);
}
void QMdiSubWindow_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQMdiSubWindow*>( (QMdiSubWindow*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QMdiSubWindow_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQMdiSubWindow*)(self) )->virtualbase_MetaObject();
}
void QMdiSubWindow_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQMdiSubWindow*>( (QMdiSubWindow*)(self) )->handle__Metacast = slot;
}
void* QMdiSubWindow_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQMdiSubWindow*)(self) )->virtualbase_Metacast(param1);
}
void QMdiSubWindow_Delete(QMdiSubWindow* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQMdiSubWindow*>( self );
	} else {
		delete self;
	}
}
