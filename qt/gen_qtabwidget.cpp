#include <QIcon>
#include <QMetaObject>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTabBar>
#include <QTabWidget>
#include <QWidget>
#include "qtabwidget.h"

#include "gen_qtabwidget.h"

extern "C" {
    extern void miqt_exec_callback(void* cb, int argc, void* argv);
}

QTabWidget* QTabWidget_new() {
	return new QTabWidget();
}

QTabWidget* QTabWidget_new2(QWidget* parent) {
	return new QTabWidget(parent);
}

QMetaObject* QTabWidget_MetaObject(const QTabWidget* self) {
	return (QMetaObject*) self->metaObject();
}

void QTabWidget_Tr(const char* s, char** _out, int* _out_Strlen) {
	QString ret = QTabWidget::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QTabWidget_TrUtf8(const char* s, char** _out, int* _out_Strlen) {
	QString ret = QTabWidget::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

int QTabWidget_AddTab(QTabWidget* self, QWidget* widget, const char* param2, size_t param2_Strlen) {
	QString param2_QString = QString::fromUtf8(param2, param2_Strlen);
	return self->addTab(widget, param2_QString);
}

int QTabWidget_AddTab2(QTabWidget* self, QWidget* widget, QIcon* icon, const char* label, size_t label_Strlen) {
	QString label_QString = QString::fromUtf8(label, label_Strlen);
	return self->addTab(widget, *icon, label_QString);
}

int QTabWidget_InsertTab(QTabWidget* self, int index, QWidget* widget, const char* param3, size_t param3_Strlen) {
	QString param3_QString = QString::fromUtf8(param3, param3_Strlen);
	return self->insertTab(static_cast<int>(index), widget, param3_QString);
}

int QTabWidget_InsertTab2(QTabWidget* self, int index, QWidget* widget, QIcon* icon, const char* label, size_t label_Strlen) {
	QString label_QString = QString::fromUtf8(label, label_Strlen);
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

void QTabWidget_TabText(const QTabWidget* self, int index, char** _out, int* _out_Strlen) {
	QString ret = self->tabText(static_cast<int>(index));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QTabWidget_SetTabText(QTabWidget* self, int index, const char* text, size_t text_Strlen) {
	QString text_QString = QString::fromUtf8(text, text_Strlen);
	self->setTabText(static_cast<int>(index), text_QString);
}

QIcon* QTabWidget_TabIcon(const QTabWidget* self, int index) {
	QIcon ret = self->tabIcon(static_cast<int>(index));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QIcon*>(new QIcon(ret));
}

void QTabWidget_SetTabIcon(QTabWidget* self, int index, QIcon* icon) {
	self->setTabIcon(static_cast<int>(index), *icon);
}

void QTabWidget_SetTabToolTip(QTabWidget* self, int index, const char* tip, size_t tip_Strlen) {
	QString tip_QString = QString::fromUtf8(tip, tip_Strlen);
	self->setTabToolTip(static_cast<int>(index), tip_QString);
}

void QTabWidget_TabToolTip(const QTabWidget* self, int index, char** _out, int* _out_Strlen) {
	QString ret = self->tabToolTip(static_cast<int>(index));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QTabWidget_SetTabWhatsThis(QTabWidget* self, int index, const char* text, size_t text_Strlen) {
	QString text_QString = QString::fromUtf8(text, text_Strlen);
	self->setTabWhatsThis(static_cast<int>(index), text_QString);
}

void QTabWidget_TabWhatsThis(const QTabWidget* self, int index, char** _out, int* _out_Strlen) {
	QString ret = self->tabWhatsThis(static_cast<int>(index));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
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

uintptr_t QTabWidget_TabPosition(const QTabWidget* self) {
	QTabWidget::TabPosition ret = self->tabPosition();
	return static_cast<uintptr_t>(ret);
}

void QTabWidget_SetTabPosition(QTabWidget* self, uintptr_t position) {
	self->setTabPosition(static_cast<QTabWidget::TabPosition>(position));
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

uintptr_t QTabWidget_TabShape(const QTabWidget* self) {
	QTabWidget::TabShape ret = self->tabShape();
	return static_cast<uintptr_t>(ret);
}

void QTabWidget_SetTabShape(QTabWidget* self, uintptr_t s) {
	self->setTabShape(static_cast<QTabWidget::TabShape>(s));
}

QSize* QTabWidget_SizeHint(const QTabWidget* self) {
	QSize ret = self->sizeHint();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QSize*>(new QSize(ret));
}

QSize* QTabWidget_MinimumSizeHint(const QTabWidget* self) {
	QSize ret = self->minimumSizeHint();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QSize*>(new QSize(ret));
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

uintptr_t QTabWidget_ElideMode(const QTabWidget* self) {
	Qt::TextElideMode ret = self->elideMode();
	return static_cast<uintptr_t>(ret);
}

void QTabWidget_SetElideMode(QTabWidget* self, uintptr_t mode) {
	self->setElideMode(static_cast<Qt::TextElideMode>(mode));
}

QSize* QTabWidget_IconSize(const QTabWidget* self) {
	QSize ret = self->iconSize();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QSize*>(new QSize(ret));
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

void QTabWidget_connect_CurrentChanged(QTabWidget* self, void* slot) {
	QTabWidget::connect(self, static_cast<void (QTabWidget::*)(int)>(&QTabWidget::currentChanged), self, [=](int index) {
		miqt_exec_callback(slot, 0, nullptr);
	});
}

void QTabWidget_TabCloseRequested(QTabWidget* self, int index) {
	self->tabCloseRequested(static_cast<int>(index));
}

void QTabWidget_connect_TabCloseRequested(QTabWidget* self, void* slot) {
	QTabWidget::connect(self, static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabCloseRequested), self, [=](int index) {
		miqt_exec_callback(slot, 0, nullptr);
	});
}

void QTabWidget_TabBarClicked(QTabWidget* self, int index) {
	self->tabBarClicked(static_cast<int>(index));
}

void QTabWidget_connect_TabBarClicked(QTabWidget* self, void* slot) {
	QTabWidget::connect(self, static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabBarClicked), self, [=](int index) {
		miqt_exec_callback(slot, 0, nullptr);
	});
}

void QTabWidget_TabBarDoubleClicked(QTabWidget* self, int index) {
	self->tabBarDoubleClicked(static_cast<int>(index));
}

void QTabWidget_connect_TabBarDoubleClicked(QTabWidget* self, void* slot) {
	QTabWidget::connect(self, static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabBarDoubleClicked), self, [=](int index) {
		miqt_exec_callback(slot, 0, nullptr);
	});
}

void QTabWidget_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen) {
	QString ret = QTabWidget::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QTabWidget_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen) {
	QString ret = QTabWidget::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QTabWidget_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen) {
	QString ret = QTabWidget::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QTabWidget_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen) {
	QString ret = QTabWidget::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QTabWidget_SetCornerWidget2(QTabWidget* self, QWidget* w, uintptr_t corner) {
	self->setCornerWidget(w, static_cast<Qt::Corner>(corner));
}

QWidget* QTabWidget_CornerWidget1(const QTabWidget* self, uintptr_t corner) {
	return self->cornerWidget(static_cast<Qt::Corner>(corner));
}

void QTabWidget_Delete(QTabWidget* self) {
	delete self;
}

