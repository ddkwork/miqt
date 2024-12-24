// +build ignore

#include <QColor>
#include <QEvent>
#include <QHideEvent>
#include <QIcon>
#include <QKeyEvent>
#include <QMetaObject>
#include <QMouseEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QPoint>
#include <QRect>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyleOptionTab>
#include <QTabBar>
#include <QTimerEvent>
#include <QVariant>
#include <QWheelEvent>
#include <QWidget>
#include <qtabbar.h>
#include "gen_qtabbar.h"
class MiqtVirtualQTabBar : public virtual QTabBar {
public:
MiqtVirtualQTabBar(QWidget* parent): QTabBar(parent) {};
MiqtVirtualQTabBar(): QTabBar() {};

virtual ~MiqtVirtualQTabBar() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QTabBar::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QTabBar_MetaObject(const_cast<MiqtVirtualQTabBar*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QTabBar::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QTabBar::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QTabBar_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QTabBar::qt_metacast(param1);

	}
};
QTabBar* QTabBar_new(QWidget* parent) {
return new MiqtVirtualQTabBar(parent);
}
QTabBar* QTabBar_new2() {
return new MiqtVirtualQTabBar();
}
void QTabBar_virtbase(QTabBar* src
, QWidget** outptr_QWidget
) {
*outptr_QWidget = static_cast<QWidget*>(src);
}
QMetaObject* QTabBar_MetaObject(const QTabBar* self) {
	return (QMetaObject*) self->metaObject();
}
void* QTabBar_Metacast(QTabBar* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QTabBar_Tr(const char* s) {
	QString _ret = QTabBar::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
Shape QTabBar_Shape(const QTabBar* self) {
	return self->shape();
}
void QTabBar_SetShape(QTabBar* self, Shape shape) {
	self->setShape(shape);
}
int QTabBar_AddTab(QTabBar* self, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return self->addTab(text_QString);
}
int QTabBar_AddTab2(QTabBar* self, QIcon* icon, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return self->addTab(*icon, text_QString);
}
int QTabBar_InsertTab(QTabBar* self, int index, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return self->insertTab(static_cast<int>(index), text_QString);
}
int QTabBar_InsertTab2(QTabBar* self, int index, QIcon* icon, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return self->insertTab(static_cast<int>(index), *icon, text_QString);
}
void QTabBar_RemoveTab(QTabBar* self, int index) {
	self->removeTab(static_cast<int>(index));
}
void QTabBar_MoveTab(QTabBar* self, int from, int to) {
	self->moveTab(static_cast<int>(from), static_cast<int>(to));
}
bool QTabBar_IsTabEnabled(const QTabBar* self, int index) {
	return self->isTabEnabled(static_cast<int>(index));
}
void QTabBar_SetTabEnabled(QTabBar* self, int index, bool enabled) {
	self->setTabEnabled(static_cast<int>(index), enabled);
}
bool QTabBar_IsTabVisible(const QTabBar* self, int index) {
	return self->isTabVisible(static_cast<int>(index));
}
void QTabBar_SetTabVisible(QTabBar* self, int index, bool visible) {
	self->setTabVisible(static_cast<int>(index), visible);
}
struct miqt_string QTabBar_TabText(const QTabBar* self, int index) {
	QString _ret = self->tabText(static_cast<int>(index));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QTabBar_SetTabText(QTabBar* self, int index, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->setTabText(static_cast<int>(index), text_QString);
}
QColor* QTabBar_TabTextColor(const QTabBar* self, int index) {
	return new QColor(self->tabTextColor(static_cast<int>(index)));
}
void QTabBar_SetTabTextColor(QTabBar* self, int index, QColor* color) {
	self->setTabTextColor(static_cast<int>(index), *color);
}
QIcon* QTabBar_TabIcon(const QTabBar* self, int index) {
	return new QIcon(self->tabIcon(static_cast<int>(index)));
}
void QTabBar_SetTabIcon(QTabBar* self, int index, QIcon* icon) {
	self->setTabIcon(static_cast<int>(index), *icon);
}
int QTabBar_ElideMode(const QTabBar* self) {
	Qt::TextElideMode _ret = self->elideMode();
	return static_cast<int>(_ret);
}
void QTabBar_SetElideMode(QTabBar* self, int mode) {
	self->setElideMode(static_cast<Qt::TextElideMode>(mode));
}
void QTabBar_SetTabToolTip(QTabBar* self, int index, struct miqt_string tip) {
	QString tip_QString = QString::fromUtf8(tip.data, tip.len);
	self->setTabToolTip(static_cast<int>(index), tip_QString);
}
struct miqt_string QTabBar_TabToolTip(const QTabBar* self, int index) {
	QString _ret = self->tabToolTip(static_cast<int>(index));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QTabBar_SetTabWhatsThis(QTabBar* self, int index, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->setTabWhatsThis(static_cast<int>(index), text_QString);
}
struct miqt_string QTabBar_TabWhatsThis(const QTabBar* self, int index) {
	QString _ret = self->tabWhatsThis(static_cast<int>(index));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QTabBar_SetTabData(QTabBar* self, int index, QVariant* data) {
	self->setTabData(static_cast<int>(index), *data);
}
QVariant* QTabBar_TabData(const QTabBar* self, int index) {
	return new QVariant(self->tabData(static_cast<int>(index)));
}
QRect* QTabBar_TabRect(const QTabBar* self, int index) {
	return new QRect(self->tabRect(static_cast<int>(index)));
}
int QTabBar_TabAt(const QTabBar* self, QPoint* pos) {
	return self->tabAt(*pos);
}
int QTabBar_CurrentIndex(const QTabBar* self) {
	return self->currentIndex();
}
int QTabBar_Count(const QTabBar* self) {
	return self->count();
}
QSize* QTabBar_SizeHint(const QTabBar* self) {
	return new QSize(self->sizeHint());
}
QSize* QTabBar_MinimumSizeHint(const QTabBar* self) {
	return new QSize(self->minimumSizeHint());
}
void QTabBar_SetDrawBase(QTabBar* self, bool drawTheBase) {
	self->setDrawBase(drawTheBase);
}
bool QTabBar_DrawBase(const QTabBar* self) {
	return self->drawBase();
}
QSize* QTabBar_IconSize(const QTabBar* self) {
	return new QSize(self->iconSize());
}
void QTabBar_SetIconSize(QTabBar* self, QSize* size) {
	self->setIconSize(*size);
}
bool QTabBar_UsesScrollButtons(const QTabBar* self) {
	return self->usesScrollButtons();
}
void QTabBar_SetUsesScrollButtons(QTabBar* self, bool useButtons) {
	self->setUsesScrollButtons(useButtons);
}
bool QTabBar_TabsClosable(const QTabBar* self) {
	return self->tabsClosable();
}
void QTabBar_SetTabsClosable(QTabBar* self, bool closable) {
	self->setTabsClosable(closable);
}
void QTabBar_SetTabButton(QTabBar* self, int index, ButtonPosition position, QWidget* widget) {
	self->setTabButton(static_cast<int>(index), position, widget);
}
QWidget* QTabBar_TabButton(const QTabBar* self, int index, ButtonPosition position) {
	return self->tabButton(static_cast<int>(index), position);
}
SelectionBehavior QTabBar_SelectionBehaviorOnRemove(const QTabBar* self) {
	return self->selectionBehaviorOnRemove();
}
void QTabBar_SetSelectionBehaviorOnRemove(QTabBar* self, SelectionBehavior behavior) {
	self->setSelectionBehaviorOnRemove(behavior);
}
bool QTabBar_Expanding(const QTabBar* self) {
	return self->expanding();
}
void QTabBar_SetExpanding(QTabBar* self, bool enabled) {
	self->setExpanding(enabled);
}
bool QTabBar_IsMovable(const QTabBar* self) {
	return self->isMovable();
}
void QTabBar_SetMovable(QTabBar* self, bool movable) {
	self->setMovable(movable);
}
bool QTabBar_DocumentMode(const QTabBar* self) {
	return self->documentMode();
}
void QTabBar_SetDocumentMode(QTabBar* self, bool set) {
	self->setDocumentMode(set);
}
bool QTabBar_AutoHide(const QTabBar* self) {
	return self->autoHide();
}
void QTabBar_SetAutoHide(QTabBar* self, bool hide) {
	self->setAutoHide(hide);
}
bool QTabBar_ChangeCurrentOnDrag(const QTabBar* self) {
	return self->changeCurrentOnDrag();
}
void QTabBar_SetChangeCurrentOnDrag(QTabBar* self, bool change) {
	self->setChangeCurrentOnDrag(change);
}
void QTabBar_SetCurrentIndex(QTabBar* self, int index) {
	self->setCurrentIndex(static_cast<int>(index));
}
void QTabBar_CurrentChanged(QTabBar* self, int index) {
	self->currentChanged(static_cast<int>(index));
}
void QTabBar_connect_CurrentChanged(QTabBar* self, intptr_t slot) {
	MiqtVirtualQTabBar::connect(self, static_cast<void (QTabBar::*)(int)>(&QTabBar::currentChanged), self, [=](int index) {
		int sigval1 = index;
		miqt_exec_callback_QTabBar_CurrentChanged(slot, sigval1);
	});
}
void QTabBar_TabCloseRequested(QTabBar* self, int index) {
	self->tabCloseRequested(static_cast<int>(index));
}
void QTabBar_connect_TabCloseRequested(QTabBar* self, intptr_t slot) {
	MiqtVirtualQTabBar::connect(self, static_cast<void (QTabBar::*)(int)>(&QTabBar::tabCloseRequested), self, [=](int index) {
		int sigval1 = index;
		miqt_exec_callback_QTabBar_TabCloseRequested(slot, sigval1);
	});
}
void QTabBar_TabMoved(QTabBar* self, int from, int to) {
	self->tabMoved(static_cast<int>(from), static_cast<int>(to));
}
void QTabBar_connect_TabMoved(QTabBar* self, intptr_t slot) {
	MiqtVirtualQTabBar::connect(self, static_cast<void (QTabBar::*)(int, int)>(&QTabBar::tabMoved), self, [=](int from, int to) {
		int sigval1 = from;
		int sigval2 = to;
		miqt_exec_callback_QTabBar_TabMoved(slot, sigval1, sigval2);
	});
}
void QTabBar_TabBarClicked(QTabBar* self, int index) {
	self->tabBarClicked(static_cast<int>(index));
}
void QTabBar_connect_TabBarClicked(QTabBar* self, intptr_t slot) {
	MiqtVirtualQTabBar::connect(self, static_cast<void (QTabBar::*)(int)>(&QTabBar::tabBarClicked), self, [=](int index) {
		int sigval1 = index;
		miqt_exec_callback_QTabBar_TabBarClicked(slot, sigval1);
	});
}
void QTabBar_TabBarDoubleClicked(QTabBar* self, int index) {
	self->tabBarDoubleClicked(static_cast<int>(index));
}
void QTabBar_connect_TabBarDoubleClicked(QTabBar* self, intptr_t slot) {
	MiqtVirtualQTabBar::connect(self, static_cast<void (QTabBar::*)(int)>(&QTabBar::tabBarDoubleClicked), self, [=](int index) {
		int sigval1 = index;
		miqt_exec_callback_QTabBar_TabBarDoubleClicked(slot, sigval1);
	});
}
struct miqt_string QTabBar_Tr2(const char* s, const char* c) {
	QString _ret = QTabBar::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QTabBar_Tr3(const char* s, const char* c, int n) {
	QString _ret = QTabBar::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QTabBar_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQTabBar*>( (QTabBar*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QTabBar_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQTabBar*)(self) )->virtualbase_MetaObject();
}
void QTabBar_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQTabBar*>( (QTabBar*)(self) )->handle__Metacast = slot;
}
void* QTabBar_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQTabBar*)(self) )->virtualbase_Metacast(param1);
}
void QTabBar_Delete(QTabBar* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQTabBar*>( self );
	} else {
		delete self;
	}
}
