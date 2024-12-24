// +build ignore

#include <QLayout>
#include <QLayoutItem>
#include <QMetaObject>
#include <QObject>
#include <QRect>
#include <QSize>
#include <QStackedLayout>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include <qstackedlayout.h>
#include "gen_qstackedlayout.h"

class MiqtVirtualQStackedLayout : public virtual QStackedLayout {
public:

	MiqtVirtualQStackedLayout(QWidget* parent): QStackedLayout(parent) {};
	MiqtVirtualQStackedLayout(): QStackedLayout() {};
	MiqtVirtualQStackedLayout(QLayout* parentLayout): QStackedLayout(parentLayout) {};

	virtual ~MiqtVirtualQStackedLayout() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QStackedLayout::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QStackedLayout_MetaObject(const_cast<MiqtVirtualQStackedLayout*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QStackedLayout::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QStackedLayout::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QStackedLayout_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QStackedLayout::qt_metacast(param1);

	}

};

QStackedLayout* QStackedLayout_new(QWidget* parent) {
	return new MiqtVirtualQStackedLayout(parent);
}

QStackedLayout* QStackedLayout_new2() {
	return new MiqtVirtualQStackedLayout();
}

QStackedLayout* QStackedLayout_new3(QLayout* parentLayout) {
	return new MiqtVirtualQStackedLayout(parentLayout);
}

void QStackedLayout_virtbase(QStackedLayout* src, QLayout** outptr_QLayout) {
	*outptr_QLayout = static_cast<QLayout*>(src);
}

QMetaObject* QStackedLayout_MetaObject(const QStackedLayout* self) {
	return (QMetaObject*) self->metaObject();
}

void* QStackedLayout_Metacast(QStackedLayout* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QStackedLayout_Tr(const char* s) {
	QString _ret = QStackedLayout::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QStackedLayout_AddWidget(QStackedLayout* self, QWidget* w) {
	return self->addWidget(w);
}

int QStackedLayout_InsertWidget(QStackedLayout* self, int index, QWidget* w) {
	return self->insertWidget(static_cast<int>(index), w);
}

QWidget* QStackedLayout_CurrentWidget(const QStackedLayout* self) {
	return self->currentWidget();
}

int QStackedLayout_CurrentIndex(const QStackedLayout* self) {
	return self->currentIndex();
}

QWidget* QStackedLayout_Widget(const QStackedLayout* self, int param1) {
	return self->widget(static_cast<int>(param1));
}

int QStackedLayout_Count(const QStackedLayout* self) {
	return self->count();
}

StackingMode QStackedLayout_StackingMode(const QStackedLayout* self) {
	return self->stackingMode();
}

void QStackedLayout_SetStackingMode(QStackedLayout* self, StackingMode stackingMode) {
	self->setStackingMode(stackingMode);
}

void QStackedLayout_AddItem(QStackedLayout* self, QLayoutItem* item) {
	self->addItem(item);
}

QSize* QStackedLayout_SizeHint(const QStackedLayout* self) {
	return new QSize(self->sizeHint());
}

QSize* QStackedLayout_MinimumSize(const QStackedLayout* self) {
	return new QSize(self->minimumSize());
}

QLayoutItem* QStackedLayout_ItemAt(const QStackedLayout* self, int param1) {
	return self->itemAt(static_cast<int>(param1));
}

QLayoutItem* QStackedLayout_TakeAt(QStackedLayout* self, int param1) {
	return self->takeAt(static_cast<int>(param1));
}

void QStackedLayout_SetGeometry(QStackedLayout* self, QRect* rect) {
	self->setGeometry(*rect);
}

bool QStackedLayout_HasHeightForWidth(const QStackedLayout* self) {
	return self->hasHeightForWidth();
}

int QStackedLayout_HeightForWidth(const QStackedLayout* self, int width) {
	return self->heightForWidth(static_cast<int>(width));
}

void QStackedLayout_WidgetRemoved(QStackedLayout* self, int index) {
	self->widgetRemoved(static_cast<int>(index));
}

void QStackedLayout_connect_WidgetRemoved(QStackedLayout* self, intptr_t slot) {
	MiqtVirtualQStackedLayout::connect(self, static_cast<void (QStackedLayout::*)(int)>(&QStackedLayout::widgetRemoved), self, [=](int index) {
		int sigval1 = index;
		miqt_exec_callback_QStackedLayout_WidgetRemoved(slot, sigval1);
	});
}

void QStackedLayout_CurrentChanged(QStackedLayout* self, int index) {
	self->currentChanged(static_cast<int>(index));
}

void QStackedLayout_connect_CurrentChanged(QStackedLayout* self, intptr_t slot) {
	MiqtVirtualQStackedLayout::connect(self, static_cast<void (QStackedLayout::*)(int)>(&QStackedLayout::currentChanged), self, [=](int index) {
		int sigval1 = index;
		miqt_exec_callback_QStackedLayout_CurrentChanged(slot, sigval1);
	});
}

void QStackedLayout_WidgetAdded(QStackedLayout* self, int index) {
	self->widgetAdded(static_cast<int>(index));
}

void QStackedLayout_connect_WidgetAdded(QStackedLayout* self, intptr_t slot) {
	MiqtVirtualQStackedLayout::connect(self, static_cast<void (QStackedLayout::*)(int)>(&QStackedLayout::widgetAdded), self, [=](int index) {
		int sigval1 = index;
		miqt_exec_callback_QStackedLayout_WidgetAdded(slot, sigval1);
	});
}

void QStackedLayout_SetCurrentIndex(QStackedLayout* self, int index) {
	self->setCurrentIndex(static_cast<int>(index));
}

void QStackedLayout_SetCurrentWidget(QStackedLayout* self, QWidget* w) {
	self->setCurrentWidget(w);
}

struct miqt_string QStackedLayout_Tr2(const char* s, const char* c) {
	QString _ret = QStackedLayout::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QStackedLayout_Tr3(const char* s, const char* c, int n) {
	QString _ret = QStackedLayout::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QStackedLayout_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQStackedLayout*>( (QStackedLayout*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QStackedLayout_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQStackedLayout*)(self) )->virtualbase_MetaObject();
}

void QStackedLayout_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQStackedLayout*>( (QStackedLayout*)(self) )->handle__Metacast = slot;
}

void* QStackedLayout_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQStackedLayout*)(self) )->virtualbase_Metacast(param1);
}

void QStackedLayout_Delete(QStackedLayout* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQStackedLayout*>( self );
	} else {
		delete self;
	}
}

