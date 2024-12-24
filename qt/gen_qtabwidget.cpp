// +build ignore

#include <QEvent>
#include <QIcon>
#include <QKeyEvent>
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyleOptionTabWidgetFrame>
#include <QTabBar>
#include <QTabWidget>
#include <QWidget>
#include <qtabwidget.h>
#include "gen_qtabwidget.h"
class MiqtVirtualQTabWidget : public virtual QTabWidget {
public:
MiqtVirtualQTabWidget(QWidget* parent): QTabWidget(parent) {};
MiqtVirtualQTabWidget(): QTabWidget() {};

virtual ~MiqtVirtualQTabWidget() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QTabWidget::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QTabWidget_MetaObject(const_cast<MiqtVirtualQTabWidget*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QTabWidget::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QTabWidget::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QTabWidget_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QTabWidget::qt_metacast(param1);

	}
};
QTabWidget* QTabWidget_new(QWidget* parent) {
return new MiqtVirtualQTabWidget(parent);
}
QTabWidget* QTabWidget_new2() {
return new MiqtVirtualQTabWidget();
}
void QTabWidget_virtbase(QTabWidget* src
, QWidget** outptr_QWidget
) {
*outptr_QWidget = static_cast<QWidget*>(src);
}
QMetaObject* QTabWidget_MetaObject(const QTabWidget* self) {
	return (QMetaObject*) self->metaObject();
}
void* QTabWidget_Metacast(QTabWidget* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QTabWidget_Tr(const char* s) {
	QString _ret = QTabWidget::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
int QTabWidget_AddTab(QTabWidget* self, QWidget* widget, struct miqt_string param2) {
	QString param2_QString = QString::fromUtf8(param2.data, param2.len);
	return self->addTab(widget, param2_QString);
}
int QTabWidget_AddTab2(QTabWidget* self, QWidget* widget, QIcon* icon, struct miqt_string label) {
	QString label_QString = QString::fromUtf8(label.data, label.len);
	return self->addTab(widget, *icon, label_QString);
}
int QTabWidget_InsertTab(QTabWidget* self, int index, QWidget* widget, struct miqt_string param3) {
	QString param3_QString = QString::fromUtf8(param3.data, param3.len);
	return self->insertTab(static_cast<int>(index), widget, param3_QString);
}
int QTabWidget_InsertTab2(QTabWidget* self, int index, QWidget* widget, QIcon* icon, struct miqt_string label) {
	QString label_QString = QString::fromUtf8(label.data, label.len);
	return self->insertTab(static_cast<int>(index), widget, *icon, label_QString);
}
void QTabWidget_RemoveTab(QTabWidget* self, int index) {
	self->removeTab(static_cast<int>(index));
}
bool QTabWidget_IsTabEnabled(const QTabWidget* self, int index) {
	return self->isTabEnabled(static_cast<int>(index));
}
void QTabWidget_SetTabEnabled(QTabWidget* self, int index, bool enabled) {
	self->setTabEnabled(static_cast<int>(index), enabled);
}
bool QTabWidget_IsTabVisible(const QTabWidget* self, int index) {
	return self->isTabVisible(static_cast<int>(index));
}
void QTabWidget_SetTabVisible(QTabWidget* self, int index, bool visible) {
	self->setTabVisible(static_cast<int>(index), visible);
}
struct miqt_string QTabWidget_TabText(const QTabWidget* self, int index) {
	QString _ret = self->tabText(static_cast<int>(index));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QTabWidget_SetTabText(QTabWidget* self, int index, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->setTabText(static_cast<int>(index), text_QString);
}
QIcon* QTabWidget_TabIcon(const QTabWidget* self, int index) {
	return new QIcon(self->tabIcon(static_cast<int>(index)));
}
void QTabWidget_SetTabIcon(QTabWidget* self, int index, QIcon* icon) {
	self->setTabIcon(static_cast<int>(index), *icon);
}
void QTabWidget_SetTabToolTip(QTabWidget* self, int index, struct miqt_string tip) {
	QString tip_QString = QString::fromUtf8(tip.data, tip.len);
	self->setTabToolTip(static_cast<int>(index), tip_QString);
}
struct miqt_string QTabWidget_TabToolTip(const QTabWidget* self, int index) {
	QString _ret = self->tabToolTip(static_cast<int>(index));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QTabWidget_SetTabWhatsThis(QTabWidget* self, int index, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->setTabWhatsThis(static_cast<int>(index), text_QString);
}
struct miqt_string QTabWidget_TabWhatsThis(const QTabWidget* self, int index) {
	QString _ret = self->tabWhatsThis(static_cast<int>(index));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
int QTabWidget_CurrentIndex(const QTabWidget* self) {
	return self->currentIndex();
}
QWidget* QTabWidget_CurrentWidget(const QTabWidget* self) {
	return self->currentWidget();
}
QWidget* QTabWidget_Widget(const QTabWidget* self, int index) {
	return self->widget(static_cast<int>(index));
}
int QTabWidget_IndexOf(const QTabWidget* self, QWidget* widget) {
	return self->indexOf(widget);
}
int QTabWidget_Count(const QTabWidget* self) {
	return self->count();
}
TabPosition QTabWidget_TabPosition(const QTabWidget* self) {
	return self->tabPosition();
}
void QTabWidget_SetTabPosition(QTabWidget* self, TabPosition position) {
	self->setTabPosition(position);
}
bool QTabWidget_TabsClosable(const QTabWidget* self) {
	return self->tabsClosable();
}
void QTabWidget_SetTabsClosable(QTabWidget* self, bool closeable) {
	self->setTabsClosable(closeable);
}
bool QTabWidget_IsMovable(const QTabWidget* self) {
	return self->isMovable();
}
void QTabWidget_SetMovable(QTabWidget* self, bool movable) {
	self->setMovable(movable);
}
TabShape QTabWidget_TabShape(const QTabWidget* self) {
	return self->tabShape();
}
void QTabWidget_SetTabShape(QTabWidget* self, TabShape s) {
	self->setTabShape(s);
}
QSize* QTabWidget_SizeHint(const QTabWidget* self) {
	return new QSize(self->sizeHint());
}
QSize* QTabWidget_MinimumSizeHint(const QTabWidget* self) {
	return new QSize(self->minimumSizeHint());
}
int QTabWidget_HeightForWidth(const QTabWidget* self, int width) {
	return self->heightForWidth(static_cast<int>(width));
}
bool QTabWidget_HasHeightForWidth(const QTabWidget* self) {
	return self->hasHeightForWidth();
}
void QTabWidget_SetCornerWidget(QTabWidget* self, QWidget* w) {
	self->setCornerWidget(w);
}
QWidget* QTabWidget_CornerWidget(const QTabWidget* self) {
	return self->cornerWidget();
}
int QTabWidget_ElideMode(const QTabWidget* self) {
	Qt::TextElideMode _ret = self->elideMode();
	return static_cast<int>(_ret);
}
void QTabWidget_SetElideMode(QTabWidget* self, int mode) {
	self->setElideMode(static_cast<Qt::TextElideMode>(mode));
}
QSize* QTabWidget_IconSize(const QTabWidget* self) {
	return new QSize(self->iconSize());
}
void QTabWidget_SetIconSize(QTabWidget* self, QSize* size) {
	self->setIconSize(*size);
}
bool QTabWidget_UsesScrollButtons(const QTabWidget* self) {
	return self->usesScrollButtons();
}
void QTabWidget_SetUsesScrollButtons(QTabWidget* self, bool useButtons) {
	self->setUsesScrollButtons(useButtons);
}
bool QTabWidget_DocumentMode(const QTabWidget* self) {
	return self->documentMode();
}
void QTabWidget_SetDocumentMode(QTabWidget* self, bool set) {
	self->setDocumentMode(set);
}
bool QTabWidget_TabBarAutoHide(const QTabWidget* self) {
	return self->tabBarAutoHide();
}
void QTabWidget_SetTabBarAutoHide(QTabWidget* self, bool enabled) {
	self->setTabBarAutoHide(enabled);
}
void QTabWidget_Clear(QTabWidget* self) {
	self->clear();
}
QTabBar* QTabWidget_TabBar(const QTabWidget* self) {
	return self->tabBar();
}
void QTabWidget_SetCurrentIndex(QTabWidget* self, int index) {
	self->setCurrentIndex(static_cast<int>(index));
}
void QTabWidget_SetCurrentWidget(QTabWidget* self, QWidget* widget) {
	self->setCurrentWidget(widget);
}
void QTabWidget_CurrentChanged(QTabWidget* self, int index) {
	self->currentChanged(static_cast<int>(index));
}
void QTabWidget_connect_CurrentChanged(QTabWidget* self, intptr_t slot) {
	MiqtVirtualQTabWidget::connect(self, static_cast<void (QTabWidget::*)(int)>(&QTabWidget::currentChanged), self, [=](int index) {
		int sigval1 = index;
		miqt_exec_callback_QTabWidget_CurrentChanged(slot, sigval1);
	});
}
void QTabWidget_TabCloseRequested(QTabWidget* self, int index) {
	self->tabCloseRequested(static_cast<int>(index));
}
void QTabWidget_connect_TabCloseRequested(QTabWidget* self, intptr_t slot) {
	MiqtVirtualQTabWidget::connect(self, static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabCloseRequested), self, [=](int index) {
		int sigval1 = index;
		miqt_exec_callback_QTabWidget_TabCloseRequested(slot, sigval1);
	});
}
void QTabWidget_TabBarClicked(QTabWidget* self, int index) {
	self->tabBarClicked(static_cast<int>(index));
}
void QTabWidget_connect_TabBarClicked(QTabWidget* self, intptr_t slot) {
	MiqtVirtualQTabWidget::connect(self, static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabBarClicked), self, [=](int index) {
		int sigval1 = index;
		miqt_exec_callback_QTabWidget_TabBarClicked(slot, sigval1);
	});
}
void QTabWidget_TabBarDoubleClicked(QTabWidget* self, int index) {
	self->tabBarDoubleClicked(static_cast<int>(index));
}
void QTabWidget_connect_TabBarDoubleClicked(QTabWidget* self, intptr_t slot) {
	MiqtVirtualQTabWidget::connect(self, static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabBarDoubleClicked), self, [=](int index) {
		int sigval1 = index;
		miqt_exec_callback_QTabWidget_TabBarDoubleClicked(slot, sigval1);
	});
}
struct miqt_string QTabWidget_Tr2(const char* s, const char* c) {
	QString _ret = QTabWidget::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QTabWidget_Tr3(const char* s, const char* c, int n) {
	QString _ret = QTabWidget::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QTabWidget_SetCornerWidget2(QTabWidget* self, QWidget* w, int corner) {
	self->setCornerWidget(w, static_cast<Qt::Corner>(corner));
}
QWidget* QTabWidget_CornerWidget1(const QTabWidget* self, int corner) {
	return self->cornerWidget(static_cast<Qt::Corner>(corner));
}
void QTabWidget_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQTabWidget*>( (QTabWidget*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QTabWidget_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQTabWidget*)(self) )->virtualbase_MetaObject();
}
void QTabWidget_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQTabWidget*>( (QTabWidget*)(self) )->handle__Metacast = slot;
}
void* QTabWidget_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQTabWidget*)(self) )->virtualbase_Metacast(param1);
}
void QTabWidget_Delete(QTabWidget* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQTabWidget*>( self );
	} else {
		delete self;
	}
}
