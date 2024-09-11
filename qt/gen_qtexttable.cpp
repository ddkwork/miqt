#include <QMetaObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTextCharFormat>
#include <QTextCursor>
#include <QTextDocument>
#define WORKAROUND_INNER_CLASS_DEFINITION_QTextFrame__iterator
#include <QTextTable>
#include <QTextTableCell>
#include <QTextTableFormat>
#include "qtexttable.h"

#include "gen_qtexttable.h"

extern "C" {
    extern void miqt_exec_callback(void* cb, int argc, void* argv);
}

QTextTableCell* QTextTableCell_new() {
	return new QTextTableCell();
}

QTextTableCell* QTextTableCell_new2(QTextTableCell* o) {
	return new QTextTableCell(*o);
}

void QTextTableCell_OperatorAssign(QTextTableCell* self, QTextTableCell* o) {
	self->operator=(*o);
}

void QTextTableCell_SetFormat(QTextTableCell* self, QTextCharFormat* format) {
	self->setFormat(*format);
}

QTextCharFormat* QTextTableCell_Format(const QTextTableCell* self) {
	QTextCharFormat ret = self->format();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QTextCharFormat*>(new QTextCharFormat(ret));
}

int QTextTableCell_Row(const QTextTableCell* self) {
	return self->row();
}

int QTextTableCell_Column(const QTextTableCell* self) {
	return self->column();
}

int QTextTableCell_RowSpan(const QTextTableCell* self) {
	return self->rowSpan();
}

int QTextTableCell_ColumnSpan(const QTextTableCell* self) {
	return self->columnSpan();
}

bool QTextTableCell_IsValid(const QTextTableCell* self) {
	return self->isValid();
}

QTextCursor* QTextTableCell_FirstCursorPosition(const QTextTableCell* self) {
	QTextCursor ret = self->firstCursorPosition();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QTextCursor*>(new QTextCursor(ret));
}

QTextCursor* QTextTableCell_LastCursorPosition(const QTextTableCell* self) {
	QTextCursor ret = self->lastCursorPosition();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QTextCursor*>(new QTextCursor(ret));
}

int QTextTableCell_FirstPosition(const QTextTableCell* self) {
	return self->firstPosition();
}

int QTextTableCell_LastPosition(const QTextTableCell* self) {
	return self->lastPosition();
}

bool QTextTableCell_OperatorEqual(const QTextTableCell* self, QTextTableCell* other) {
	return self->operator==(*other);
}

bool QTextTableCell_OperatorNotEqual(const QTextTableCell* self, QTextTableCell* other) {
	return self->operator!=(*other);
}

QTextFrame__iterator* QTextTableCell_Begin(const QTextTableCell* self) {
	QTextFrame::iterator ret = self->begin();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QTextFrame::iterator*>(new QTextFrame::iterator(ret));
}

QTextFrame__iterator* QTextTableCell_End(const QTextTableCell* self) {
	QTextFrame::iterator ret = self->end();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QTextFrame::iterator*>(new QTextFrame::iterator(ret));
}

int QTextTableCell_TableCellFormatIndex(const QTextTableCell* self) {
	return self->tableCellFormatIndex();
}

void QTextTableCell_Delete(QTextTableCell* self) {
	delete self;
}

QTextTable* QTextTable_new(QTextDocument* doc) {
	return new QTextTable(doc);
}

QMetaObject* QTextTable_MetaObject(const QTextTable* self) {
	return (QMetaObject*) self->metaObject();
}

void QTextTable_Tr(const char* s, char** _out, int* _out_Strlen) {
	QString ret = QTextTable::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QTextTable_TrUtf8(const char* s, char** _out, int* _out_Strlen) {
	QString ret = QTextTable::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QTextTable_Resize(QTextTable* self, int rows, int cols) {
	self->resize(static_cast<int>(rows), static_cast<int>(cols));
}

void QTextTable_InsertRows(QTextTable* self, int pos, int num) {
	self->insertRows(static_cast<int>(pos), static_cast<int>(num));
}

void QTextTable_InsertColumns(QTextTable* self, int pos, int num) {
	self->insertColumns(static_cast<int>(pos), static_cast<int>(num));
}

void QTextTable_AppendRows(QTextTable* self, int count) {
	self->appendRows(static_cast<int>(count));
}

void QTextTable_AppendColumns(QTextTable* self, int count) {
	self->appendColumns(static_cast<int>(count));
}

void QTextTable_RemoveRows(QTextTable* self, int pos, int num) {
	self->removeRows(static_cast<int>(pos), static_cast<int>(num));
}

void QTextTable_RemoveColumns(QTextTable* self, int pos, int num) {
	self->removeColumns(static_cast<int>(pos), static_cast<int>(num));
}

void QTextTable_MergeCells(QTextTable* self, int row, int col, int numRows, int numCols) {
	self->mergeCells(static_cast<int>(row), static_cast<int>(col), static_cast<int>(numRows), static_cast<int>(numCols));
}

void QTextTable_MergeCellsWithCursor(QTextTable* self, QTextCursor* cursor) {
	self->mergeCells(*cursor);
}

void QTextTable_SplitCell(QTextTable* self, int row, int col, int numRows, int numCols) {
	self->splitCell(static_cast<int>(row), static_cast<int>(col), static_cast<int>(numRows), static_cast<int>(numCols));
}

int QTextTable_Rows(const QTextTable* self) {
	return self->rows();
}

int QTextTable_Columns(const QTextTable* self) {
	return self->columns();
}

QTextTableCell* QTextTable_CellAt(const QTextTable* self, int row, int col) {
	QTextTableCell ret = self->cellAt(static_cast<int>(row), static_cast<int>(col));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QTextTableCell*>(new QTextTableCell(ret));
}

QTextTableCell* QTextTable_CellAtWithPosition(const QTextTable* self, int position) {
	QTextTableCell ret = self->cellAt(static_cast<int>(position));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QTextTableCell*>(new QTextTableCell(ret));
}

QTextTableCell* QTextTable_CellAtWithQTextCursor(const QTextTable* self, QTextCursor* c) {
	QTextTableCell ret = self->cellAt(*c);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QTextTableCell*>(new QTextTableCell(ret));
}

QTextCursor* QTextTable_RowStart(const QTextTable* self, QTextCursor* c) {
	QTextCursor ret = self->rowStart(*c);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QTextCursor*>(new QTextCursor(ret));
}

QTextCursor* QTextTable_RowEnd(const QTextTable* self, QTextCursor* c) {
	QTextCursor ret = self->rowEnd(*c);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QTextCursor*>(new QTextCursor(ret));
}

void QTextTable_SetFormat(QTextTable* self, QTextTableFormat* format) {
	self->setFormat(*format);
}

QTextTableFormat* QTextTable_Format(const QTextTable* self) {
	QTextTableFormat ret = self->format();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QTextTableFormat*>(new QTextTableFormat(ret));
}

void QTextTable_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen) {
	QString ret = QTextTable::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QTextTable_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen) {
	QString ret = QTextTable::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QTextTable_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen) {
	QString ret = QTextTable::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QTextTable_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen) {
	QString ret = QTextTable::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QTextTable_Delete(QTextTable* self) {
	delete self;
}

