// +build ignore

#include <QAbstractItemModel>
#include <QItemSelection>
#include <QItemSelectionModel>
#include <QItemSelectionRange>
#include <QList>
#include <QMetaObject>
#include <QModelIndex>
#include <QObject>
#include <QPersistentModelIndex>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qitemselectionmodel.h>
#include "gen_qitemselectionmodel.h"
QItemSelectionRange* QItemSelectionRange_new() {
return new QItemSelectionRange();
}
QItemSelectionRange* QItemSelectionRange_new2(QModelIndex* topL, QModelIndex* bottomR) {
return new QItemSelectionRange(*topL, *bottomR);
}
QItemSelectionRange* QItemSelectionRange_new3(QModelIndex* index) {
return new QItemSelectionRange(*index);
}
void QItemSelectionRange_Swap(QItemSelectionRange* self, QItemSelectionRange* other) {
	self->swap(*other);
}
int QItemSelectionRange_Top(const QItemSelectionRange* self) {
	return self->top();
}
int QItemSelectionRange_Left(const QItemSelectionRange* self) {
	return self->left();
}
int QItemSelectionRange_Bottom(const QItemSelectionRange* self) {
	return self->bottom();
}
int QItemSelectionRange_Right(const QItemSelectionRange* self) {
	return self->right();
}
int QItemSelectionRange_Width(const QItemSelectionRange* self) {
	return self->width();
}
int QItemSelectionRange_Height(const QItemSelectionRange* self) {
	return self->height();
}
QPersistentModelIndex* QItemSelectionRange_TopLeft(const QItemSelectionRange* self) {
	const QPersistentModelIndex& _ret = self->topLeft();
	// Cast returned reference into pointer
	return const_cast<QPersistentModelIndex*>(&_ret);
}
QPersistentModelIndex* QItemSelectionRange_BottomRight(const QItemSelectionRange* self) {
	const QPersistentModelIndex& _ret = self->bottomRight();
	// Cast returned reference into pointer
	return const_cast<QPersistentModelIndex*>(&_ret);
}
QModelIndex* QItemSelectionRange_Parent(const QItemSelectionRange* self) {
	return new QModelIndex(self->parent());
}
QAbstractItemModel* QItemSelectionRange_Model(const QItemSelectionRange* self) {
	return (QAbstractItemModel*) self->model();
}
bool QItemSelectionRange_Contains(const QItemSelectionRange* self, QModelIndex* index) {
	return self->contains(*index);
}
bool QItemSelectionRange_Contains2(const QItemSelectionRange* self, int row, int column, QModelIndex* parentIndex) {
	return self->contains(static_cast<int>(row), static_cast<int>(column), *parentIndex);
}
bool QItemSelectionRange_Intersects(const QItemSelectionRange* self, QItemSelectionRange* other) {
	return self->intersects(*other);
}
QItemSelectionRange* QItemSelectionRange_Intersected(const QItemSelectionRange* self, QItemSelectionRange* other) {
	return new QItemSelectionRange(self->intersected(*other));
}
bool QItemSelectionRange_IsValid(const QItemSelectionRange* self) {
	return self->isValid();
}
bool QItemSelectionRange_IsEmpty(const QItemSelectionRange* self) {
	return self->isEmpty();
}
struct miqt_array /* of QModelIndex* */  QItemSelectionRange_Indexes(const QItemSelectionRange* self) {
	QModelIndexList _ret = self->indexes();
	// Convert QList<> from C++ memory to manually-managed C memory
	QModelIndex** _arr = static_cast<QModelIndex**>(malloc(sizeof(QModelIndex*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QModelIndex(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
void QItemSelectionRange_Delete(QItemSelectionRange* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QItemSelectionRange*>( self );
	} else {
		delete self;
	}
}
class MiqtVirtualQItemSelectionModel : public virtual QItemSelectionModel {
public:
MiqtVirtualQItemSelectionModel(): QItemSelectionModel() {};
MiqtVirtualQItemSelectionModel(QAbstractItemModel* model, QObject* parent): QItemSelectionModel(model, parent) {};
MiqtVirtualQItemSelectionModel(QAbstractItemModel* model): QItemSelectionModel(model) {};

virtual ~MiqtVirtualQItemSelectionModel() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QItemSelectionModel::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QItemSelectionModel_MetaObject(const_cast<MiqtVirtualQItemSelectionModel*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QItemSelectionModel::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QItemSelectionModel::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QItemSelectionModel_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QItemSelectionModel::qt_metacast(param1);

	}
};
QItemSelectionModel* QItemSelectionModel_new() {
return new MiqtVirtualQItemSelectionModel();
}
QItemSelectionModel* QItemSelectionModel_new2(QAbstractItemModel* model, QObject* parent) {
return new MiqtVirtualQItemSelectionModel(model, parent);
}
QItemSelectionModel* QItemSelectionModel_new3(QAbstractItemModel* model) {
return new MiqtVirtualQItemSelectionModel(model);
}
void QItemSelectionModel_virtbase(QItemSelectionModel* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QItemSelectionModel_MetaObject(const QItemSelectionModel* self) {
	return (QMetaObject*) self->metaObject();
}
void* QItemSelectionModel_Metacast(QItemSelectionModel* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QItemSelectionModel_Tr(const char* s) {
	QString _ret = QItemSelectionModel::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QModelIndex* QItemSelectionModel_CurrentIndex(const QItemSelectionModel* self) {
	return new QModelIndex(self->currentIndex());
}
bool QItemSelectionModel_IsSelected(const QItemSelectionModel* self, QModelIndex* index) {
	return self->isSelected(*index);
}
bool QItemSelectionModel_IsRowSelected(const QItemSelectionModel* self, int row) {
	return self->isRowSelected(static_cast<int>(row));
}
bool QItemSelectionModel_IsColumnSelected(const QItemSelectionModel* self, int column) {
	return self->isColumnSelected(static_cast<int>(column));
}
bool QItemSelectionModel_RowIntersectsSelection(const QItemSelectionModel* self, int row) {
	return self->rowIntersectsSelection(static_cast<int>(row));
}
bool QItemSelectionModel_ColumnIntersectsSelection(const QItemSelectionModel* self, int column) {
	return self->columnIntersectsSelection(static_cast<int>(column));
}
bool QItemSelectionModel_HasSelection(const QItemSelectionModel* self) {
	return self->hasSelection();
}
struct miqt_array /* of QModelIndex* */  QItemSelectionModel_SelectedIndexes(const QItemSelectionModel* self) {
	QModelIndexList _ret = self->selectedIndexes();
	// Convert QList<> from C++ memory to manually-managed C memory
	QModelIndex** _arr = static_cast<QModelIndex**>(malloc(sizeof(QModelIndex*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QModelIndex(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
struct miqt_array /* of QModelIndex* */  QItemSelectionModel_SelectedRows(const QItemSelectionModel* self) {
	QModelIndexList _ret = self->selectedRows();
	// Convert QList<> from C++ memory to manually-managed C memory
	QModelIndex** _arr = static_cast<QModelIndex**>(malloc(sizeof(QModelIndex*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QModelIndex(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
struct miqt_array /* of QModelIndex* */  QItemSelectionModel_SelectedColumns(const QItemSelectionModel* self) {
	QModelIndexList _ret = self->selectedColumns();
	// Convert QList<> from C++ memory to manually-managed C memory
	QModelIndex** _arr = static_cast<QModelIndex**>(malloc(sizeof(QModelIndex*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QModelIndex(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
QItemSelection* QItemSelectionModel_Selection(const QItemSelectionModel* self) {
	return new QItemSelection(self->selection());
}
QAbstractItemModel* QItemSelectionModel_Model(const QItemSelectionModel* self) {
	return (QAbstractItemModel*) self->model();
}
QAbstractItemModel* QItemSelectionModel_Model2(QItemSelectionModel* self) {
	return self->model();
}
void QItemSelectionModel_SetModel(QItemSelectionModel* self, QAbstractItemModel* model) {
	self->setModel(model);
}
void QItemSelectionModel_SetCurrentIndex(QItemSelectionModel* self, QModelIndex* index, int command) {
	self->setCurrentIndex(*index, static_cast<QItemSelectionModel::SelectionFlags>(command));
}
void QItemSelectionModel_Select(QItemSelectionModel* self, QModelIndex* index, int command) {
	self->select(*index, static_cast<QItemSelectionModel::SelectionFlags>(command));
}
void QItemSelectionModel_Select2(QItemSelectionModel* self, QItemSelection* selection, int command) {
	self->select(*selection, static_cast<QItemSelectionModel::SelectionFlags>(command));
}
void QItemSelectionModel_Clear(QItemSelectionModel* self) {
	self->clear();
}
void QItemSelectionModel_Reset(QItemSelectionModel* self) {
	self->reset();
}
void QItemSelectionModel_ClearSelection(QItemSelectionModel* self) {
	self->clearSelection();
}
void QItemSelectionModel_ClearCurrentIndex(QItemSelectionModel* self) {
	self->clearCurrentIndex();
}
void QItemSelectionModel_SelectionChanged(QItemSelectionModel* self, QItemSelection* selected, QItemSelection* deselected) {
	self->selectionChanged(*selected, *deselected);
}
void QItemSelectionModel_connect_SelectionChanged(QItemSelectionModel* self, intptr_t slot) {
	MiqtVirtualQItemSelectionModel::connect(self, static_cast<void (QItemSelectionModel::*)(const QItemSelection&, const QItemSelection&)>(&QItemSelectionModel::selectionChanged), self, [=](const QItemSelection& selected, const QItemSelection& deselected) {
		const QItemSelection& selected_ret = selected;
		// Cast returned reference into pointer
		QItemSelection* sigval1 = const_cast<QItemSelection*>(&selected_ret);
		const QItemSelection& deselected_ret = deselected;
		// Cast returned reference into pointer
		QItemSelection* sigval2 = const_cast<QItemSelection*>(&deselected_ret);
		miqt_exec_callback_QItemSelectionModel_SelectionChanged(slot, sigval1, sigval2);
	});
}
void QItemSelectionModel_CurrentChanged(QItemSelectionModel* self, QModelIndex* current, QModelIndex* previous) {
	self->currentChanged(*current, *previous);
}
void QItemSelectionModel_connect_CurrentChanged(QItemSelectionModel* self, intptr_t slot) {
	MiqtVirtualQItemSelectionModel::connect(self, static_cast<void (QItemSelectionModel::*)(const QModelIndex&, const QModelIndex&)>(&QItemSelectionModel::currentChanged), self, [=](const QModelIndex& current, const QModelIndex& previous) {
		const QModelIndex& current_ret = current;
		// Cast returned reference into pointer
		QModelIndex* sigval1 = const_cast<QModelIndex*>(&current_ret);
		const QModelIndex& previous_ret = previous;
		// Cast returned reference into pointer
		QModelIndex* sigval2 = const_cast<QModelIndex*>(&previous_ret);
		miqt_exec_callback_QItemSelectionModel_CurrentChanged(slot, sigval1, sigval2);
	});
}
void QItemSelectionModel_CurrentRowChanged(QItemSelectionModel* self, QModelIndex* current, QModelIndex* previous) {
	self->currentRowChanged(*current, *previous);
}
void QItemSelectionModel_connect_CurrentRowChanged(QItemSelectionModel* self, intptr_t slot) {
	MiqtVirtualQItemSelectionModel::connect(self, static_cast<void (QItemSelectionModel::*)(const QModelIndex&, const QModelIndex&)>(&QItemSelectionModel::currentRowChanged), self, [=](const QModelIndex& current, const QModelIndex& previous) {
		const QModelIndex& current_ret = current;
		// Cast returned reference into pointer
		QModelIndex* sigval1 = const_cast<QModelIndex*>(&current_ret);
		const QModelIndex& previous_ret = previous;
		// Cast returned reference into pointer
		QModelIndex* sigval2 = const_cast<QModelIndex*>(&previous_ret);
		miqt_exec_callback_QItemSelectionModel_CurrentRowChanged(slot, sigval1, sigval2);
	});
}
void QItemSelectionModel_CurrentColumnChanged(QItemSelectionModel* self, QModelIndex* current, QModelIndex* previous) {
	self->currentColumnChanged(*current, *previous);
}
void QItemSelectionModel_connect_CurrentColumnChanged(QItemSelectionModel* self, intptr_t slot) {
	MiqtVirtualQItemSelectionModel::connect(self, static_cast<void (QItemSelectionModel::*)(const QModelIndex&, const QModelIndex&)>(&QItemSelectionModel::currentColumnChanged), self, [=](const QModelIndex& current, const QModelIndex& previous) {
		const QModelIndex& current_ret = current;
		// Cast returned reference into pointer
		QModelIndex* sigval1 = const_cast<QModelIndex*>(&current_ret);
		const QModelIndex& previous_ret = previous;
		// Cast returned reference into pointer
		QModelIndex* sigval2 = const_cast<QModelIndex*>(&previous_ret);
		miqt_exec_callback_QItemSelectionModel_CurrentColumnChanged(slot, sigval1, sigval2);
	});
}
void QItemSelectionModel_ModelChanged(QItemSelectionModel* self, QAbstractItemModel* model) {
	self->modelChanged(model);
}
void QItemSelectionModel_connect_ModelChanged(QItemSelectionModel* self, intptr_t slot) {
	MiqtVirtualQItemSelectionModel::connect(self, static_cast<void (QItemSelectionModel::*)(QAbstractItemModel*)>(&QItemSelectionModel::modelChanged), self, [=](QAbstractItemModel* model) {
		QAbstractItemModel* sigval1 = model;
		miqt_exec_callback_QItemSelectionModel_ModelChanged(slot, sigval1);
	});
}
struct miqt_string QItemSelectionModel_Tr2(const char* s, const char* c) {
	QString _ret = QItemSelectionModel::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QItemSelectionModel_Tr3(const char* s, const char* c, int n) {
	QString _ret = QItemSelectionModel::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
bool QItemSelectionModel_IsRowSelected2(const QItemSelectionModel* self, int row, QModelIndex* parent) {
	return self->isRowSelected(static_cast<int>(row), *parent);
}
bool QItemSelectionModel_IsColumnSelected2(const QItemSelectionModel* self, int column, QModelIndex* parent) {
	return self->isColumnSelected(static_cast<int>(column), *parent);
}
bool QItemSelectionModel_RowIntersectsSelection2(const QItemSelectionModel* self, int row, QModelIndex* parent) {
	return self->rowIntersectsSelection(static_cast<int>(row), *parent);
}
bool QItemSelectionModel_ColumnIntersectsSelection2(const QItemSelectionModel* self, int column, QModelIndex* parent) {
	return self->columnIntersectsSelection(static_cast<int>(column), *parent);
}
struct miqt_array /* of QModelIndex* */  QItemSelectionModel_SelectedRows1(const QItemSelectionModel* self, int column) {
	QModelIndexList _ret = self->selectedRows(static_cast<int>(column));
	// Convert QList<> from C++ memory to manually-managed C memory
	QModelIndex** _arr = static_cast<QModelIndex**>(malloc(sizeof(QModelIndex*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QModelIndex(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
struct miqt_array /* of QModelIndex* */  QItemSelectionModel_SelectedColumns1(const QItemSelectionModel* self, int row) {
	QModelIndexList _ret = self->selectedColumns(static_cast<int>(row));
	// Convert QList<> from C++ memory to manually-managed C memory
	QModelIndex** _arr = static_cast<QModelIndex**>(malloc(sizeof(QModelIndex*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QModelIndex(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
void QItemSelectionModel_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQItemSelectionModel*>( (QItemSelectionModel*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QItemSelectionModel_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQItemSelectionModel*)(self) )->virtualbase_MetaObject();
}
void QItemSelectionModel_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQItemSelectionModel*>( (QItemSelectionModel*)(self) )->handle__Metacast = slot;
}
void* QItemSelectionModel_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQItemSelectionModel*)(self) )->virtualbase_Metacast(param1);
}
void QItemSelectionModel_Delete(QItemSelectionModel* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQItemSelectionModel*>( self );
	} else {
		delete self;
	}
}
QItemSelection* QItemSelection_new(QModelIndex* topLeft, QModelIndex* bottomRight) {
return new QItemSelection(*topLeft, *bottomRight);
}
QItemSelection* QItemSelection_new2() {
return new QItemSelection();
}
QItemSelection* QItemSelection_new3(QItemSelection* param1) {
return new QItemSelection(*param1);
}
void QItemSelection_Select(QItemSelection* self, QModelIndex* topLeft, QModelIndex* bottomRight) {
	self->select(*topLeft, *bottomRight);
}
bool QItemSelection_Contains(const QItemSelection* self, QModelIndex* index) {
	return self->contains(*index);
}
struct miqt_array /* of QModelIndex* */  QItemSelection_Indexes(const QItemSelection* self) {
	QModelIndexList _ret = self->indexes();
	// Convert QList<> from C++ memory to manually-managed C memory
	QModelIndex** _arr = static_cast<QModelIndex**>(malloc(sizeof(QModelIndex*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QModelIndex(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
void QItemSelection_Merge(QItemSelection* self, QItemSelection* other, int command) {
	self->merge(*other, static_cast<QItemSelectionModel::SelectionFlags>(command));
}
void QItemSelection_Split(QItemSelectionRange* rangeVal, QItemSelectionRange* other, QItemSelection* result) {
	QItemSelection::split(*rangeVal, *other, result);
}
void QItemSelection_Delete(QItemSelection* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QItemSelection*>( self );
	} else {
		delete self;
	}
}
