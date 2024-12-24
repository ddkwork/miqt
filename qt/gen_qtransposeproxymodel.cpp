// +build ignore

#include <QAbstractItemModel>
#include <QAbstractProxyModel>
#include <QMap>
#include <QMetaObject>
#include <QModelIndex>
#include <QObject>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTransposeProxyModel>
#include <QVariant>
#include <qtransposeproxymodel.h>
#include "gen_qtransposeproxymodel.h"
class MiqtVirtualQTransposeProxyModel : public virtual QTransposeProxyModel {
public:
MiqtVirtualQTransposeProxyModel(): QTransposeProxyModel() {};
MiqtVirtualQTransposeProxyModel(QObject* parent): QTransposeProxyModel(parent) {};

virtual ~MiqtVirtualQTransposeProxyModel() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QTransposeProxyModel::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QTransposeProxyModel_MetaObject(const_cast<MiqtVirtualQTransposeProxyModel*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QTransposeProxyModel::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QTransposeProxyModel::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QTransposeProxyModel_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QTransposeProxyModel::qt_metacast(param1);

	}
};
QTransposeProxyModel* QTransposeProxyModel_new() {
return new MiqtVirtualQTransposeProxyModel();
}
QTransposeProxyModel* QTransposeProxyModel_new2(QObject* parent) {
return new MiqtVirtualQTransposeProxyModel(parent);
}
void QTransposeProxyModel_virtbase(QTransposeProxyModel* src
, QAbstractProxyModel** outptr_QAbstractProxyModel
) {
*outptr_QAbstractProxyModel = static_cast<QAbstractProxyModel*>(src);
}
QMetaObject* QTransposeProxyModel_MetaObject(const QTransposeProxyModel* self) {
	return (QMetaObject*) self->metaObject();
}
void* QTransposeProxyModel_Metacast(QTransposeProxyModel* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QTransposeProxyModel_Tr(const char* s) {
	QString _ret = QTransposeProxyModel::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QTransposeProxyModel_SetSourceModel(QTransposeProxyModel* self, QAbstractItemModel* newSourceModel) {
	self->setSourceModel(newSourceModel);
}
int QTransposeProxyModel_RowCount(const QTransposeProxyModel* self, QModelIndex* parent) {
	return self->rowCount(*parent);
}
int QTransposeProxyModel_ColumnCount(const QTransposeProxyModel* self, QModelIndex* parent) {
	return self->columnCount(*parent);
}
QVariant* QTransposeProxyModel_HeaderData(const QTransposeProxyModel* self, int section, int orientation, int role) {
	return new QVariant(self->headerData(static_cast<int>(section), static_cast<Qt::Orientation>(orientation), static_cast<int>(role)));
}
bool QTransposeProxyModel_SetHeaderData(QTransposeProxyModel* self, int section, int orientation, QVariant* value, int role) {
	return self->setHeaderData(static_cast<int>(section), static_cast<Qt::Orientation>(orientation), *value, static_cast<int>(role));
}
bool QTransposeProxyModel_SetItemData(QTransposeProxyModel* self, QModelIndex* index, struct miqt_map /* of int to QVariant* */  roles) {
	QMap<int, QVariant> roles_QMap;
	int* roles_karr = static_cast<int*>(roles.keys);
	QVariant** roles_varr = static_cast<QVariant**>(roles.values);
	for(size_t i = 0; i < roles.len; ++i) {
		roles_QMap[static_cast<int>(roles_karr[i])] = *(roles_varr[i]);
	}
	return self->setItemData(*index, roles_QMap);
}
QSize* QTransposeProxyModel_Span(const QTransposeProxyModel* self, QModelIndex* index) {
	return new QSize(self->span(*index));
}
struct miqt_map /* of int to QVariant* */  QTransposeProxyModel_ItemData(const QTransposeProxyModel* self, QModelIndex* index) {
	QMap<int, QVariant> _ret = self->itemData(*index);
	// Convert QMap<> from C++ memory to manually-managed C memory
	int* _karr = static_cast<int*>(malloc(sizeof(int) * _ret.size()));
	QVariant** _varr = static_cast<QVariant**>(malloc(sizeof(QVariant*) * _ret.size()));
	int _ctr = 0;
	for (auto _itr = _ret.keyValueBegin(); _itr != _ret.keyValueEnd(); ++_itr) {
		_karr[_ctr] = _itr->first;
		_varr[_ctr] = new QVariant(_itr->second);
		_ctr++;
	}
	struct miqt_map _out;
	_out.len = _ret.size();
	_out.keys = static_cast<void*>(_karr);
	_out.values = static_cast<void*>(_varr);
	return _out;
}
QModelIndex* QTransposeProxyModel_MapFromSource(const QTransposeProxyModel* self, QModelIndex* sourceIndex) {
	return new QModelIndex(self->mapFromSource(*sourceIndex));
}
QModelIndex* QTransposeProxyModel_MapToSource(const QTransposeProxyModel* self, QModelIndex* proxyIndex) {
	return new QModelIndex(self->mapToSource(*proxyIndex));
}
QModelIndex* QTransposeProxyModel_Parent(const QTransposeProxyModel* self, QModelIndex* index) {
	return new QModelIndex(self->parent(*index));
}
QModelIndex* QTransposeProxyModel_Index(const QTransposeProxyModel* self, int row, int column, QModelIndex* parent) {
	return new QModelIndex(self->index(static_cast<int>(row), static_cast<int>(column), *parent));
}
bool QTransposeProxyModel_InsertRows(QTransposeProxyModel* self, int row, int count, QModelIndex* parent) {
	return self->insertRows(static_cast<int>(row), static_cast<int>(count), *parent);
}
bool QTransposeProxyModel_RemoveRows(QTransposeProxyModel* self, int row, int count, QModelIndex* parent) {
	return self->removeRows(static_cast<int>(row), static_cast<int>(count), *parent);
}
bool QTransposeProxyModel_MoveRows(QTransposeProxyModel* self, QModelIndex* sourceParent, int sourceRow, int count, QModelIndex* destinationParent, int destinationChild) {
	return self->moveRows(*sourceParent, static_cast<int>(sourceRow), static_cast<int>(count), *destinationParent, static_cast<int>(destinationChild));
}
bool QTransposeProxyModel_InsertColumns(QTransposeProxyModel* self, int column, int count, QModelIndex* parent) {
	return self->insertColumns(static_cast<int>(column), static_cast<int>(count), *parent);
}
bool QTransposeProxyModel_RemoveColumns(QTransposeProxyModel* self, int column, int count, QModelIndex* parent) {
	return self->removeColumns(static_cast<int>(column), static_cast<int>(count), *parent);
}
bool QTransposeProxyModel_MoveColumns(QTransposeProxyModel* self, QModelIndex* sourceParent, int sourceColumn, int count, QModelIndex* destinationParent, int destinationChild) {
	return self->moveColumns(*sourceParent, static_cast<int>(sourceColumn), static_cast<int>(count), *destinationParent, static_cast<int>(destinationChild));
}
void QTransposeProxyModel_Sort(QTransposeProxyModel* self, int column, int order) {
	self->sort(static_cast<int>(column), static_cast<Qt::SortOrder>(order));
}
struct miqt_string QTransposeProxyModel_Tr2(const char* s, const char* c) {
	QString _ret = QTransposeProxyModel::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QTransposeProxyModel_Tr3(const char* s, const char* c, int n) {
	QString _ret = QTransposeProxyModel::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QTransposeProxyModel_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQTransposeProxyModel*>( (QTransposeProxyModel*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QTransposeProxyModel_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQTransposeProxyModel*)(self) )->virtualbase_MetaObject();
}
void QTransposeProxyModel_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQTransposeProxyModel*>( (QTransposeProxyModel*)(self) )->handle__Metacast = slot;
}
void* QTransposeProxyModel_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQTransposeProxyModel*)(self) )->virtualbase_Metacast(param1);
}
void QTransposeProxyModel_Delete(QTransposeProxyModel* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQTransposeProxyModel*>( self );
	} else {
		delete self;
	}
}
