// +build ignore

#include <QAbstractItemModel>
#include <QAbstractItemView>
#include <QAbstractScrollArea>
#include <QColumnView>
#include <QFrame>
#include <QItemSelection>
#include <QItemSelectionModel>
#include <QList>
#include <QMetaObject>
#include <QModelIndex>
#include <QObject>
#include <QPaintDevice>
#include <QPoint>
#include <QRect>
#include <QRegion>
#include <QResizeEvent>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include <qcolumnview.h>
#include "gen_qcolumnview.h"
class MiqtVirtualQColumnView : public virtual QColumnView {
public:
MiqtVirtualQColumnView(QWidget* parent): QColumnView(parent) {};
MiqtVirtualQColumnView(): QColumnView() {};

virtual ~MiqtVirtualQColumnView() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QColumnView::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QColumnView_MetaObject(const_cast<MiqtVirtualQColumnView*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QColumnView::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QColumnView::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QColumnView_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QColumnView::qt_metacast(param1);

	}
};
QColumnView* QColumnView_new(QWidget* parent) {
return new MiqtVirtualQColumnView(parent);
}
QColumnView* QColumnView_new2() {
return new MiqtVirtualQColumnView();
}
void QColumnView_virtbase(QColumnView* src
, QAbstractItemView** outptr_QAbstractItemView
) {
*outptr_QAbstractItemView = static_cast<QAbstractItemView*>(src);
}
QMetaObject* QColumnView_MetaObject(const QColumnView* self) {
	return (QMetaObject*) self->metaObject();
}
void* QColumnView_Metacast(QColumnView* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QColumnView_Tr(const char* s) {
	QString _ret = QColumnView::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QColumnView_UpdatePreviewWidget(QColumnView* self, QModelIndex* index) {
	self->updatePreviewWidget(*index);
}
void QColumnView_connect_UpdatePreviewWidget(QColumnView* self, intptr_t slot) {
	MiqtVirtualQColumnView::connect(self, static_cast<void (QColumnView::*)(const QModelIndex&)>(&QColumnView::updatePreviewWidget), self, [=](const QModelIndex& index) {
		const QModelIndex& index_ret = index;
		// Cast returned reference into pointer
		QModelIndex* sigval1 = const_cast<QModelIndex*>(&index_ret);
		miqt_exec_callback_QColumnView_UpdatePreviewWidget(slot, sigval1);
	});
}
QModelIndex* QColumnView_IndexAt(const QColumnView* self, QPoint* point) {
	return new QModelIndex(self->indexAt(*point));
}
void QColumnView_ScrollTo(QColumnView* self, QModelIndex* index, ScrollHint hint) {
	self->scrollTo(*index, hint);
}
QSize* QColumnView_SizeHint(const QColumnView* self) {
	return new QSize(self->sizeHint());
}
QRect* QColumnView_VisualRect(const QColumnView* self, QModelIndex* index) {
	return new QRect(self->visualRect(*index));
}
void QColumnView_SetModel(QColumnView* self, QAbstractItemModel* model) {
	self->setModel(model);
}
void QColumnView_SetSelectionModel(QColumnView* self, QItemSelectionModel* selectionModel) {
	self->setSelectionModel(selectionModel);
}
void QColumnView_SetRootIndex(QColumnView* self, QModelIndex* index) {
	self->setRootIndex(*index);
}
void QColumnView_SelectAll(QColumnView* self) {
	self->selectAll();
}
void QColumnView_SetResizeGripsVisible(QColumnView* self, bool visible) {
	self->setResizeGripsVisible(visible);
}
bool QColumnView_ResizeGripsVisible(const QColumnView* self) {
	return self->resizeGripsVisible();
}
QWidget* QColumnView_PreviewWidget(const QColumnView* self) {
	return self->previewWidget();
}
void QColumnView_SetPreviewWidget(QColumnView* self, QWidget* widget) {
	self->setPreviewWidget(widget);
}
void QColumnView_SetColumnWidths(QColumnView* self, struct miqt_array /* of int */  list) {
	QList<int> list_QList;
	list_QList.reserve(list.len);
	int* list_arr = static_cast<int*>(list.data);
	for(size_t i = 0; i < list.len; ++i) {
		list_QList.push_back(static_cast<int>(list_arr[i]));
	}
	self->setColumnWidths(list_QList);
}
struct miqt_array /* of int */  QColumnView_ColumnWidths(const QColumnView* self) {
	QList<int> _ret = self->columnWidths();
	// Convert QList<> from C++ memory to manually-managed C memory
	int* _arr = static_cast<int*>(malloc(sizeof(int) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
struct miqt_string QColumnView_Tr2(const char* s, const char* c) {
	QString _ret = QColumnView::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QColumnView_Tr3(const char* s, const char* c, int n) {
	QString _ret = QColumnView::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QColumnView_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQColumnView*>( (QColumnView*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QColumnView_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQColumnView*)(self) )->virtualbase_MetaObject();
}
void QColumnView_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQColumnView*>( (QColumnView*)(self) )->handle__Metacast = slot;
}
void* QColumnView_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQColumnView*)(self) )->virtualbase_Metacast(param1);
}
void QColumnView_Delete(QColumnView* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQColumnView*>( self );
	} else {
		delete self;
	}
}
