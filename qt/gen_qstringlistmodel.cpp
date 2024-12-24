// +build ignore

#include <QAbstractItemModel>
#include <QAbstractListModel>
#include <QList>
#include <QMap>
#include <QMetaObject>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStringListModel>
#include <QVariant>
#include <qstringlistmodel.h>
#include "gen_qstringlistmodel.h"
class MiqtVirtualQStringListModel : public virtual QStringListModel {
public:
MiqtVirtualQStringListModel(): QStringListModel() {};
MiqtVirtualQStringListModel(const QStringList& strings): QStringListModel(strings) {};
MiqtVirtualQStringListModel(QObject* parent): QStringListModel(parent) {};
MiqtVirtualQStringListModel(const QStringList& strings, QObject* parent): QStringListModel(strings, parent) {};

virtual ~MiqtVirtualQStringListModel() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QStringListModel::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QStringListModel_MetaObject(const_cast<MiqtVirtualQStringListModel*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QStringListModel::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QStringListModel::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QStringListModel_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QStringListModel::qt_metacast(param1);

	}
};
QStringListModel* QStringListModel_new() {
return new MiqtVirtualQStringListModel();
}
QStringListModel* QStringListModel_new2(struct miqt_array /* of struct miqt_string */  strings) {
QStringList strings_QList;
	strings_QList.reserve(strings.len);
	struct miqt_string* strings_arr = static_cast<struct miqt_string*>(strings.data);
	for(size_t i = 0; i < strings.len; ++i) {
		QString strings_arr_i_QString = QString::fromUtf8(strings_arr[i].data, strings_arr[i].len);
		strings_QList.push_back(strings_arr_i_QString);
	}
	return new MiqtVirtualQStringListModel(strings_QList);
}
QStringListModel* QStringListModel_new3(QObject* parent) {
return new MiqtVirtualQStringListModel(parent);
}
QStringListModel* QStringListModel_new4(struct miqt_array /* of struct miqt_string */  strings, QObject* parent) {
QStringList strings_QList;
	strings_QList.reserve(strings.len);
	struct miqt_string* strings_arr = static_cast<struct miqt_string*>(strings.data);
	for(size_t i = 0; i < strings.len; ++i) {
		QString strings_arr_i_QString = QString::fromUtf8(strings_arr[i].data, strings_arr[i].len);
		strings_QList.push_back(strings_arr_i_QString);
	}
	return new MiqtVirtualQStringListModel(strings_QList, parent);
}
void QStringListModel_virtbase(QStringListModel* src
, QAbstractListModel** outptr_QAbstractListModel
) {
*outptr_QAbstractListModel = static_cast<QAbstractListModel*>(src);
}
QMetaObject* QStringListModel_MetaObject(const QStringListModel* self) {
	return (QMetaObject*) self->metaObject();
}
void* QStringListModel_Metacast(QStringListModel* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QStringListModel_Tr(const char* s) {
	QString _ret = QStringListModel::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
int QStringListModel_RowCount(const QStringListModel* self, QModelIndex* parent) {
	return self->rowCount(*parent);
}
QModelIndex* QStringListModel_Sibling(const QStringListModel* self, int row, int column, QModelIndex* idx) {
	return new QModelIndex(self->sibling(static_cast<int>(row), static_cast<int>(column), *idx));
}
QVariant* QStringListModel_Data(const QStringListModel* self, QModelIndex* index, int role) {
	return new QVariant(self->data(*index, static_cast<int>(role)));
}
bool QStringListModel_SetData(QStringListModel* self, QModelIndex* index, QVariant* value, int role) {
	return self->setData(*index, *value, static_cast<int>(role));
}
bool QStringListModel_ClearItemData(QStringListModel* self, QModelIndex* index) {
	return self->clearItemData(*index);
}
int QStringListModel_Flags(const QStringListModel* self, QModelIndex* index) {
	Qt::ItemFlags _ret = self->flags(*index);
	return static_cast<int>(_ret);
}
bool QStringListModel_InsertRows(QStringListModel* self, int row, int count, QModelIndex* parent) {
	return self->insertRows(static_cast<int>(row), static_cast<int>(count), *parent);
}
bool QStringListModel_RemoveRows(QStringListModel* self, int row, int count, QModelIndex* parent) {
	return self->removeRows(static_cast<int>(row), static_cast<int>(count), *parent);
}
bool QStringListModel_MoveRows(QStringListModel* self, QModelIndex* sourceParent, int sourceRow, int count, QModelIndex* destinationParent, int destinationChild) {
	return self->moveRows(*sourceParent, static_cast<int>(sourceRow), static_cast<int>(count), *destinationParent, static_cast<int>(destinationChild));
}
struct miqt_map /* of int to QVariant* */  QStringListModel_ItemData(const QStringListModel* self, QModelIndex* index) {
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
bool QStringListModel_SetItemData(QStringListModel* self, QModelIndex* index, struct miqt_map /* of int to QVariant* */  roles) {
	QMap<int, QVariant> roles_QMap;
	int* roles_karr = static_cast<int*>(roles.keys);
	QVariant** roles_varr = static_cast<QVariant**>(roles.values);
	for(size_t i = 0; i < roles.len; ++i) {
		roles_QMap[static_cast<int>(roles_karr[i])] = *(roles_varr[i]);
	}
	return self->setItemData(*index, roles_QMap);
}
void QStringListModel_Sort(QStringListModel* self, int column, int order) {
	self->sort(static_cast<int>(column), static_cast<Qt::SortOrder>(order));
}
struct miqt_array /* of struct miqt_string */  QStringListModel_StringList(const QStringListModel* self) {
	QStringList _ret = self->stringList();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
void QStringListModel_SetStringList(QStringListModel* self, struct miqt_array /* of struct miqt_string */  strings) {
	QStringList strings_QList;
	strings_QList.reserve(strings.len);
	struct miqt_string* strings_arr = static_cast<struct miqt_string*>(strings.data);
	for(size_t i = 0; i < strings.len; ++i) {
		QString strings_arr_i_QString = QString::fromUtf8(strings_arr[i].data, strings_arr[i].len);
		strings_QList.push_back(strings_arr_i_QString);
	}
	self->setStringList(strings_QList);
}
int QStringListModel_SupportedDropActions(const QStringListModel* self) {
	Qt::DropActions _ret = self->supportedDropActions();
	return static_cast<int>(_ret);
}
struct miqt_string QStringListModel_Tr2(const char* s, const char* c) {
	QString _ret = QStringListModel::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QStringListModel_Tr3(const char* s, const char* c, int n) {
	QString _ret = QStringListModel::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QStringListModel_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQStringListModel*>( (QStringListModel*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QStringListModel_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQStringListModel*)(self) )->virtualbase_MetaObject();
}
void QStringListModel_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQStringListModel*>( (QStringListModel*)(self) )->handle__Metacast = slot;
}
void* QStringListModel_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQStringListModel*)(self) )->virtualbase_Metacast(param1);
}
void QStringListModel_Delete(QStringListModel* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQStringListModel*>( self );
	} else {
		delete self;
	}
}
