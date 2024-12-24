// +build ignore

#include <QAbstractItemModel>
#include <QAbstractProxyModel>
#include <QIdentityProxyModel>
#include <QItemSelection>
#include <QList>
#include <QMetaObject>
#include <QMimeData>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QVariant>
#include <qidentityproxymodel.h>
#include "gen_qidentityproxymodel.h"
class MiqtVirtualQIdentityProxyModel : public virtual QIdentityProxyModel {
public:
MiqtVirtualQIdentityProxyModel(): QIdentityProxyModel() {};
MiqtVirtualQIdentityProxyModel(QObject* parent): QIdentityProxyModel(parent) {};

virtual ~MiqtVirtualQIdentityProxyModel() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QIdentityProxyModel::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QIdentityProxyModel_MetaObject(const_cast<MiqtVirtualQIdentityProxyModel*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QIdentityProxyModel::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QIdentityProxyModel::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QIdentityProxyModel_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QIdentityProxyModel::qt_metacast(param1);

	}
};
QIdentityProxyModel* QIdentityProxyModel_new() {
return new MiqtVirtualQIdentityProxyModel();
}
QIdentityProxyModel* QIdentityProxyModel_new2(QObject* parent) {
return new MiqtVirtualQIdentityProxyModel(parent);
}
void QIdentityProxyModel_virtbase(QIdentityProxyModel* src
, QAbstractProxyModel** outptr_QAbstractProxyModel
) {
*outptr_QAbstractProxyModel = static_cast<QAbstractProxyModel*>(src);
}
QMetaObject* QIdentityProxyModel_MetaObject(const QIdentityProxyModel* self) {
	return (QMetaObject*) self->metaObject();
}
void* QIdentityProxyModel_Metacast(QIdentityProxyModel* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QIdentityProxyModel_Tr(const char* s) {
	QString _ret = QIdentityProxyModel::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
int QIdentityProxyModel_ColumnCount(const QIdentityProxyModel* self, QModelIndex* parent) {
	return self->columnCount(*parent);
}
QModelIndex* QIdentityProxyModel_Index(const QIdentityProxyModel* self, int row, int column, QModelIndex* parent) {
	return new QModelIndex(self->index(static_cast<int>(row), static_cast<int>(column), *parent));
}
QModelIndex* QIdentityProxyModel_MapFromSource(const QIdentityProxyModel* self, QModelIndex* sourceIndex) {
	return new QModelIndex(self->mapFromSource(*sourceIndex));
}
QModelIndex* QIdentityProxyModel_MapToSource(const QIdentityProxyModel* self, QModelIndex* proxyIndex) {
	return new QModelIndex(self->mapToSource(*proxyIndex));
}
QModelIndex* QIdentityProxyModel_Parent(const QIdentityProxyModel* self, QModelIndex* child) {
	return new QModelIndex(self->parent(*child));
}
int QIdentityProxyModel_RowCount(const QIdentityProxyModel* self, QModelIndex* parent) {
	return self->rowCount(*parent);
}
QVariant* QIdentityProxyModel_HeaderData(const QIdentityProxyModel* self, int section, int orientation, int role) {
	return new QVariant(self->headerData(static_cast<int>(section), static_cast<Qt::Orientation>(orientation), static_cast<int>(role)));
}
bool QIdentityProxyModel_DropMimeData(QIdentityProxyModel* self, QMimeData* data, int action, int row, int column, QModelIndex* parent) {
	return self->dropMimeData(data, static_cast<Qt::DropAction>(action), static_cast<int>(row), static_cast<int>(column), *parent);
}
QModelIndex* QIdentityProxyModel_Sibling(const QIdentityProxyModel* self, int row, int column, QModelIndex* idx) {
	return new QModelIndex(self->sibling(static_cast<int>(row), static_cast<int>(column), *idx));
}
QItemSelection* QIdentityProxyModel_MapSelectionFromSource(const QIdentityProxyModel* self, QItemSelection* selection) {
	return new QItemSelection(self->mapSelectionFromSource(*selection));
}
QItemSelection* QIdentityProxyModel_MapSelectionToSource(const QIdentityProxyModel* self, QItemSelection* selection) {
	return new QItemSelection(self->mapSelectionToSource(*selection));
}
struct miqt_array /* of QModelIndex* */  QIdentityProxyModel_Match(const QIdentityProxyModel* self, QModelIndex* start, int role, QVariant* value, int hits, int flags) {
	QModelIndexList _ret = self->match(*start, static_cast<int>(role), *value, static_cast<int>(hits), static_cast<Qt::MatchFlags>(flags));
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
void QIdentityProxyModel_SetSourceModel(QIdentityProxyModel* self, QAbstractItemModel* sourceModel) {
	self->setSourceModel(sourceModel);
}
bool QIdentityProxyModel_InsertColumns(QIdentityProxyModel* self, int column, int count, QModelIndex* parent) {
	return self->insertColumns(static_cast<int>(column), static_cast<int>(count), *parent);
}
bool QIdentityProxyModel_InsertRows(QIdentityProxyModel* self, int row, int count, QModelIndex* parent) {
	return self->insertRows(static_cast<int>(row), static_cast<int>(count), *parent);
}
bool QIdentityProxyModel_RemoveColumns(QIdentityProxyModel* self, int column, int count, QModelIndex* parent) {
	return self->removeColumns(static_cast<int>(column), static_cast<int>(count), *parent);
}
bool QIdentityProxyModel_RemoveRows(QIdentityProxyModel* self, int row, int count, QModelIndex* parent) {
	return self->removeRows(static_cast<int>(row), static_cast<int>(count), *parent);
}
bool QIdentityProxyModel_MoveRows(QIdentityProxyModel* self, QModelIndex* sourceParent, int sourceRow, int count, QModelIndex* destinationParent, int destinationChild) {
	return self->moveRows(*sourceParent, static_cast<int>(sourceRow), static_cast<int>(count), *destinationParent, static_cast<int>(destinationChild));
}
bool QIdentityProxyModel_MoveColumns(QIdentityProxyModel* self, QModelIndex* sourceParent, int sourceColumn, int count, QModelIndex* destinationParent, int destinationChild) {
	return self->moveColumns(*sourceParent, static_cast<int>(sourceColumn), static_cast<int>(count), *destinationParent, static_cast<int>(destinationChild));
}
bool QIdentityProxyModel_HandleSourceLayoutChanges(const QIdentityProxyModel* self) {
	return self->handleSourceLayoutChanges();
}
bool QIdentityProxyModel_HandleSourceDataChanges(const QIdentityProxyModel* self) {
	return self->handleSourceDataChanges();
}
struct miqt_string QIdentityProxyModel_Tr2(const char* s, const char* c) {
	QString _ret = QIdentityProxyModel::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QIdentityProxyModel_Tr3(const char* s, const char* c, int n) {
	QString _ret = QIdentityProxyModel::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QIdentityProxyModel_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQIdentityProxyModel*>( (QIdentityProxyModel*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QIdentityProxyModel_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQIdentityProxyModel*)(self) )->virtualbase_MetaObject();
}
void QIdentityProxyModel_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQIdentityProxyModel*>( (QIdentityProxyModel*)(self) )->handle__Metacast = slot;
}
void* QIdentityProxyModel_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQIdentityProxyModel*)(self) )->virtualbase_Metacast(param1);
}
void QIdentityProxyModel_Delete(QIdentityProxyModel* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQIdentityProxyModel*>( self );
	} else {
		delete self;
	}
}
