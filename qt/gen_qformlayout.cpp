// +build ignore

#include <QFormLayout>
#define WORKAROUND_INNER_CLASS_DEFINITION_QFormLayout__TakeRowResult
#include <QLayout>
#include <QLayoutItem>
#include <QMetaObject>
#include <QObject>
#include <QRect>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include <qformlayout.h>
#include "gen_qformlayout.h"

class MiqtVirtualQFormLayout : public virtual QFormLayout {
public:

	MiqtVirtualQFormLayout(QWidget* parent): QFormLayout(parent) {};
	MiqtVirtualQFormLayout(): QFormLayout() {};

	virtual ~MiqtVirtualQFormLayout() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QFormLayout::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QFormLayout_MetaObject(const_cast<MiqtVirtualQFormLayout*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QFormLayout::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QFormLayout::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QFormLayout_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QFormLayout::qt_metacast(param1);

	}

};

QFormLayout* QFormLayout_new(QWidget* parent) {
	return new MiqtVirtualQFormLayout(parent);
}

QFormLayout* QFormLayout_new2() {
	return new MiqtVirtualQFormLayout();
}

void QFormLayout_virtbase(QFormLayout* src, QLayout** outptr_QLayout) {
	*outptr_QLayout = static_cast<QLayout*>(src);
}

QMetaObject* QFormLayout_MetaObject(const QFormLayout* self) {
	return (QMetaObject*) self->metaObject();
}

void* QFormLayout_Metacast(QFormLayout* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QFormLayout_Tr(const char* s) {
	QString _ret = QFormLayout::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QFormLayout_SetFieldGrowthPolicy(QFormLayout* self, FieldGrowthPolicy policy) {
	self->setFieldGrowthPolicy(policy);
}

FieldGrowthPolicy QFormLayout_FieldGrowthPolicy(const QFormLayout* self) {
	return self->fieldGrowthPolicy();
}

void QFormLayout_SetRowWrapPolicy(QFormLayout* self, RowWrapPolicy policy) {
	self->setRowWrapPolicy(policy);
}

RowWrapPolicy QFormLayout_RowWrapPolicy(const QFormLayout* self) {
	return self->rowWrapPolicy();
}

void QFormLayout_SetLabelAlignment(QFormLayout* self, int alignment) {
	self->setLabelAlignment(static_cast<Qt::Alignment>(alignment));
}

int QFormLayout_LabelAlignment(const QFormLayout* self) {
	Qt::Alignment _ret = self->labelAlignment();
	return static_cast<int>(_ret);
}

void QFormLayout_SetFormAlignment(QFormLayout* self, int alignment) {
	self->setFormAlignment(static_cast<Qt::Alignment>(alignment));
}

int QFormLayout_FormAlignment(const QFormLayout* self) {
	Qt::Alignment _ret = self->formAlignment();
	return static_cast<int>(_ret);
}

void QFormLayout_SetHorizontalSpacing(QFormLayout* self, int spacing) {
	self->setHorizontalSpacing(static_cast<int>(spacing));
}

int QFormLayout_HorizontalSpacing(const QFormLayout* self) {
	return self->horizontalSpacing();
}

void QFormLayout_SetVerticalSpacing(QFormLayout* self, int spacing) {
	self->setVerticalSpacing(static_cast<int>(spacing));
}

int QFormLayout_VerticalSpacing(const QFormLayout* self) {
	return self->verticalSpacing();
}

int QFormLayout_Spacing(const QFormLayout* self) {
	return self->spacing();
}

void QFormLayout_SetSpacing(QFormLayout* self, int spacing) {
	self->setSpacing(static_cast<int>(spacing));
}

void QFormLayout_AddRow(QFormLayout* self, QWidget* label, QWidget* field) {
	self->addRow(label, field);
}

void QFormLayout_AddRow2(QFormLayout* self, QWidget* label, QLayout* field) {
	self->addRow(label, field);
}

void QFormLayout_AddRow3(QFormLayout* self, struct miqt_string labelText, QWidget* field) {
	QString labelText_QString = QString::fromUtf8(labelText.data, labelText.len);
	self->addRow(labelText_QString, field);
}

void QFormLayout_AddRow4(QFormLayout* self, struct miqt_string labelText, QLayout* field) {
	QString labelText_QString = QString::fromUtf8(labelText.data, labelText.len);
	self->addRow(labelText_QString, field);
}

void QFormLayout_AddRowWithWidget(QFormLayout* self, QWidget* widget) {
	self->addRow(widget);
}

void QFormLayout_AddRowWithLayout(QFormLayout* self, QLayout* layout) {
	self->addRow(layout);
}

void QFormLayout_InsertRow(QFormLayout* self, int row, QWidget* label, QWidget* field) {
	self->insertRow(static_cast<int>(row), label, field);
}

void QFormLayout_InsertRow2(QFormLayout* self, int row, QWidget* label, QLayout* field) {
	self->insertRow(static_cast<int>(row), label, field);
}

void QFormLayout_InsertRow3(QFormLayout* self, int row, struct miqt_string labelText, QWidget* field) {
	QString labelText_QString = QString::fromUtf8(labelText.data, labelText.len);
	self->insertRow(static_cast<int>(row), labelText_QString, field);
}

void QFormLayout_InsertRow4(QFormLayout* self, int row, struct miqt_string labelText, QLayout* field) {
	QString labelText_QString = QString::fromUtf8(labelText.data, labelText.len);
	self->insertRow(static_cast<int>(row), labelText_QString, field);
}

void QFormLayout_InsertRow5(QFormLayout* self, int row, QWidget* widget) {
	self->insertRow(static_cast<int>(row), widget);
}

void QFormLayout_InsertRow6(QFormLayout* self, int row, QLayout* layout) {
	self->insertRow(static_cast<int>(row), layout);
}

void QFormLayout_RemoveRow(QFormLayout* self, int row) {
	self->removeRow(static_cast<int>(row));
}

void QFormLayout_RemoveRowWithWidget(QFormLayout* self, QWidget* widget) {
	self->removeRow(widget);
}

void QFormLayout_RemoveRowWithLayout(QFormLayout* self, QLayout* layout) {
	self->removeRow(layout);
}

TakeRowResult QFormLayout_TakeRow(QFormLayout* self, int row) {
	return self->takeRow(static_cast<int>(row));
}

TakeRowResult QFormLayout_TakeRowWithWidget(QFormLayout* self, QWidget* widget) {
	return self->takeRow(widget);
}

TakeRowResult QFormLayout_TakeRowWithLayout(QFormLayout* self, QLayout* layout) {
	return self->takeRow(layout);
}

void QFormLayout_SetItem(QFormLayout* self, int row, ItemRole role, QLayoutItem* item) {
	self->setItem(static_cast<int>(row), role, item);
}

void QFormLayout_SetWidget(QFormLayout* self, int row, ItemRole role, QWidget* widget) {
	self->setWidget(static_cast<int>(row), role, widget);
}

void QFormLayout_SetLayout(QFormLayout* self, int row, ItemRole role, QLayout* layout) {
	self->setLayout(static_cast<int>(row), role, layout);
}

void QFormLayout_SetRowVisible(QFormLayout* self, int row, bool on) {
	self->setRowVisible(static_cast<int>(row), on);
}

void QFormLayout_SetRowVisible2(QFormLayout* self, QWidget* widget, bool on) {
	self->setRowVisible(widget, on);
}

void QFormLayout_SetRowVisible3(QFormLayout* self, QLayout* layout, bool on) {
	self->setRowVisible(layout, on);
}

bool QFormLayout_IsRowVisible(const QFormLayout* self, int row) {
	return self->isRowVisible(static_cast<int>(row));
}

bool QFormLayout_IsRowVisibleWithWidget(const QFormLayout* self, QWidget* widget) {
	return self->isRowVisible(widget);
}

bool QFormLayout_IsRowVisibleWithLayout(const QFormLayout* self, QLayout* layout) {
	return self->isRowVisible(layout);
}

QLayoutItem* QFormLayout_ItemAt(const QFormLayout* self, int row, ItemRole role) {
	return self->itemAt(static_cast<int>(row), role);
}

void QFormLayout_GetItemPosition(const QFormLayout* self, int index, int* rowPtr, ItemRole* rolePtr) {
	self->getItemPosition(static_cast<int>(index), static_cast<int*>(rowPtr), rolePtr);
}

void QFormLayout_GetWidgetPosition(const QFormLayout* self, QWidget* widget, int* rowPtr, ItemRole* rolePtr) {
	self->getWidgetPosition(widget, static_cast<int*>(rowPtr), rolePtr);
}

void QFormLayout_GetLayoutPosition(const QFormLayout* self, QLayout* layout, int* rowPtr, ItemRole* rolePtr) {
	self->getLayoutPosition(layout, static_cast<int*>(rowPtr), rolePtr);
}

QWidget* QFormLayout_LabelForField(const QFormLayout* self, QWidget* field) {
	return self->labelForField(field);
}

QWidget* QFormLayout_LabelForFieldWithField(const QFormLayout* self, QLayout* field) {
	return self->labelForField(field);
}

void QFormLayout_AddItem(QFormLayout* self, QLayoutItem* item) {
	self->addItem(item);
}

QLayoutItem* QFormLayout_ItemAtWithIndex(const QFormLayout* self, int index) {
	return self->itemAt(static_cast<int>(index));
}

QLayoutItem* QFormLayout_TakeAt(QFormLayout* self, int index) {
	return self->takeAt(static_cast<int>(index));
}

void QFormLayout_SetGeometry(QFormLayout* self, QRect* rect) {
	self->setGeometry(*rect);
}

QSize* QFormLayout_MinimumSize(const QFormLayout* self) {
	return new QSize(self->minimumSize());
}

QSize* QFormLayout_SizeHint(const QFormLayout* self) {
	return new QSize(self->sizeHint());
}

void QFormLayout_Invalidate(QFormLayout* self) {
	self->invalidate();
}

bool QFormLayout_HasHeightForWidth(const QFormLayout* self) {
	return self->hasHeightForWidth();
}

int QFormLayout_HeightForWidth(const QFormLayout* self, int width) {
	return self->heightForWidth(static_cast<int>(width));
}

int QFormLayout_ExpandingDirections(const QFormLayout* self) {
	Qt::Orientations _ret = self->expandingDirections();
	return static_cast<int>(_ret);
}

int QFormLayout_Count(const QFormLayout* self) {
	return self->count();
}

int QFormLayout_RowCount(const QFormLayout* self) {
	return self->rowCount();
}

struct miqt_string QFormLayout_Tr2(const char* s, const char* c) {
	QString _ret = QFormLayout::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QFormLayout_Tr3(const char* s, const char* c, int n) {
	QString _ret = QFormLayout::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QFormLayout_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQFormLayout*>( (QFormLayout*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QFormLayout_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQFormLayout*)(self) )->virtualbase_MetaObject();
}

void QFormLayout_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQFormLayout*>( (QFormLayout*)(self) )->handle__Metacast = slot;
}

void* QFormLayout_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQFormLayout*)(self) )->virtualbase_Metacast(param1);
}

void QFormLayout_Delete(QFormLayout* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQFormLayout*>( self );
	} else {
		delete self;
	}
}

QFormLayout__TakeRowResult* QFormLayout__TakeRowResult_new() {
	return new QFormLayout::TakeRowResult();
}

QFormLayout__TakeRowResult* QFormLayout__TakeRowResult_new2(const TakeRowResult* param1) {
	return new QFormLayout::TakeRowResult(*param1);
}

void QFormLayout__TakeRowResult_Delete(QFormLayout__TakeRowResult* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QFormLayout::TakeRowResult*>( self );
	} else {
		delete self;
	}
}

