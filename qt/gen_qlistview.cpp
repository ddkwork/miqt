// +build ignore

#include <QAbstractItemView>
#include <QAbstractScrollArea>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QFrame>
#include <QItemSelection>
#include <QList>
#include <QListView>
#include <QMetaObject>
#include <QModelIndex>
#include <QMouseEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QPoint>
#include <QRect>
#include <QRegion>
#include <QResizeEvent>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyleOptionViewItem>
#include <QTimerEvent>
#include <QWheelEvent>
#include <QWidget>
#include <qlistview.h>
#include "gen_qlistview.h"
class MiqtVirtualQListView : public virtual QListView {
public:
MiqtVirtualQListView(QWidget* parent): QListView(parent) {};
MiqtVirtualQListView(): QListView() {};

virtual ~MiqtVirtualQListView() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QListView::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QListView_MetaObject(const_cast<MiqtVirtualQListView*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QListView::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QListView::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QListView_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QListView::qt_metacast(param1);

	}
};
QListView* QListView_new(QWidget* parent) {
return new MiqtVirtualQListView(parent);
}
QListView* QListView_new2() {
return new MiqtVirtualQListView();
}
void QListView_virtbase(QListView* src
, QAbstractItemView** outptr_QAbstractItemView
) {
*outptr_QAbstractItemView = static_cast<QAbstractItemView*>(src);
}
QMetaObject* QListView_MetaObject(const QListView* self) {
	return (QMetaObject*) self->metaObject();
}
void* QListView_Metacast(QListView* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QListView_Tr(const char* s) {
	QString _ret = QListView::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QListView_SetMovement(QListView* self, Movement movement) {
	self->setMovement(movement);
}
Movement QListView_Movement(const QListView* self) {
	return self->movement();
}
void QListView_SetFlow(QListView* self, Flow flow) {
	self->setFlow(flow);
}
Flow QListView_Flow(const QListView* self) {
	return self->flow();
}
void QListView_SetWrapping(QListView* self, bool enable) {
	self->setWrapping(enable);
}
bool QListView_IsWrapping(const QListView* self) {
	return self->isWrapping();
}
void QListView_SetResizeMode(QListView* self, ResizeMode mode) {
	self->setResizeMode(mode);
}
ResizeMode QListView_ResizeMode(const QListView* self) {
	return self->resizeMode();
}
void QListView_SetLayoutMode(QListView* self, LayoutMode mode) {
	self->setLayoutMode(mode);
}
LayoutMode QListView_LayoutMode(const QListView* self) {
	return self->layoutMode();
}
void QListView_SetSpacing(QListView* self, int space) {
	self->setSpacing(static_cast<int>(space));
}
int QListView_Spacing(const QListView* self) {
	return self->spacing();
}
void QListView_SetBatchSize(QListView* self, int batchSize) {
	self->setBatchSize(static_cast<int>(batchSize));
}
int QListView_BatchSize(const QListView* self) {
	return self->batchSize();
}
void QListView_SetGridSize(QListView* self, QSize* size) {
	self->setGridSize(*size);
}
QSize* QListView_GridSize(const QListView* self) {
	return new QSize(self->gridSize());
}
void QListView_SetViewMode(QListView* self, ViewMode mode) {
	self->setViewMode(mode);
}
ViewMode QListView_ViewMode(const QListView* self) {
	return self->viewMode();
}
void QListView_ClearPropertyFlags(QListView* self) {
	self->clearPropertyFlags();
}
bool QListView_IsRowHidden(const QListView* self, int row) {
	return self->isRowHidden(static_cast<int>(row));
}
void QListView_SetRowHidden(QListView* self, int row, bool hide) {
	self->setRowHidden(static_cast<int>(row), hide);
}
void QListView_SetModelColumn(QListView* self, int column) {
	self->setModelColumn(static_cast<int>(column));
}
int QListView_ModelColumn(const QListView* self) {
	return self->modelColumn();
}
void QListView_SetUniformItemSizes(QListView* self, bool enable) {
	self->setUniformItemSizes(enable);
}
bool QListView_UniformItemSizes(const QListView* self) {
	return self->uniformItemSizes();
}
void QListView_SetWordWrap(QListView* self, bool on) {
	self->setWordWrap(on);
}
bool QListView_WordWrap(const QListView* self) {
	return self->wordWrap();
}
void QListView_SetSelectionRectVisible(QListView* self, bool show) {
	self->setSelectionRectVisible(show);
}
bool QListView_IsSelectionRectVisible(const QListView* self) {
	return self->isSelectionRectVisible();
}
void QListView_SetItemAlignment(QListView* self, int alignment) {
	self->setItemAlignment(static_cast<Qt::Alignment>(alignment));
}
int QListView_ItemAlignment(const QListView* self) {
	Qt::Alignment _ret = self->itemAlignment();
	return static_cast<int>(_ret);
}
QRect* QListView_VisualRect(const QListView* self, QModelIndex* index) {
	return new QRect(self->visualRect(*index));
}
void QListView_ScrollTo(QListView* self, QModelIndex* index, ScrollHint hint) {
	self->scrollTo(*index, hint);
}
QModelIndex* QListView_IndexAt(const QListView* self, QPoint* p) {
	return new QModelIndex(self->indexAt(*p));
}
void QListView_DoItemsLayout(QListView* self) {
	self->doItemsLayout();
}
void QListView_Reset(QListView* self) {
	self->reset();
}
void QListView_SetRootIndex(QListView* self, QModelIndex* index) {
	self->setRootIndex(*index);
}
void QListView_IndexesMoved(QListView* self, struct miqt_array /* of QModelIndex* */  indexes) {
	QModelIndexList indexes_QList;
	indexes_QList.reserve(indexes.len);
	QModelIndex** indexes_arr = static_cast<QModelIndex**>(indexes.data);
	for(size_t i = 0; i < indexes.len; ++i) {
		indexes_QList.push_back(*(indexes_arr[i]));
	}
	self->indexesMoved(indexes_QList);
}
void QListView_connect_IndexesMoved(QListView* self, intptr_t slot) {
	MiqtVirtualQListView::connect(self, static_cast<void (QListView::*)(const QModelIndexList&)>(&QListView::indexesMoved), self, [=](const QModelIndexList& indexes) {
		const QModelIndexList& indexes_ret = indexes;
		// Convert QList<> from C++ memory to manually-managed C memory
		QModelIndex** indexes_arr = static_cast<QModelIndex**>(malloc(sizeof(QModelIndex*) * indexes_ret.length()));
		for (size_t i = 0, e = indexes_ret.length(); i < e; ++i) {
			indexes_arr[i] = new QModelIndex(indexes_ret[i]);
		}
		struct miqt_array indexes_out;
		indexes_out.len = indexes_ret.length();
		indexes_out.data = static_cast<void*>(indexes_arr);
		struct miqt_array /* of QModelIndex* */  sigval1 = indexes_out;
		miqt_exec_callback_QListView_IndexesMoved(slot, sigval1);
	});
}
struct miqt_string QListView_Tr2(const char* s, const char* c) {
	QString _ret = QListView::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QListView_Tr3(const char* s, const char* c, int n) {
	QString _ret = QListView::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QListView_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQListView*>( (QListView*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QListView_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQListView*)(self) )->virtualbase_MetaObject();
}
void QListView_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQListView*>( (QListView*)(self) )->handle__Metacast = slot;
}
void* QListView_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQListView*)(self) )->virtualbase_Metacast(param1);
}
void QListView_Delete(QListView* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQListView*>( self );
	} else {
		delete self;
	}
}
