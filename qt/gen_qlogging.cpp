// +build ignore

#include <QDebug>
#include <QLoggingCategory>
#include <QMessageLogContext>
#include <QMessageLogger>
#include <QNoDebug>
#include <qlogging.h>
#include "gen_qlogging.h"

QMessageLogContext* QMessageLogContext_new() {
	return new QMessageLogContext();
}

QMessageLogContext* QMessageLogContext_new2(const char* fileName, int lineNumber, const char* functionName, const char* categoryName) {
	return new QMessageLogContext(fileName, static_cast<int>(lineNumber), functionName, categoryName);
}

void QMessageLogContext_Delete(QMessageLogContext* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QMessageLogContext*>( self );
	} else {
		delete self;
	}
}

QMessageLogger* QMessageLogger_new() {
	return new QMessageLogger();
}

QMessageLogger* QMessageLogger_new2(const char* file, int line, const char* function) {
	return new QMessageLogger(file, static_cast<int>(line), function);
}

QMessageLogger* QMessageLogger_new3(const char* file, int line, const char* function, const char* category) {
	return new QMessageLogger(file, static_cast<int>(line), function, category);
}

void QMessageLogger_Debug(const QMessageLogger* self, const char* msg, ... ) {
	self->debug(msg, );
}

void QMessageLogger_NoDebug(const QMessageLogger* self, const char* param1, ... ) {
	self->noDebug(param1, );
}

void QMessageLogger_Info(const QMessageLogger* self, const char* msg, ... ) {
	self->info(msg, );
}

void QMessageLogger_Warning(const QMessageLogger* self, const char* msg, ... ) {
	self->warning(msg, );
}

void QMessageLogger_Critical(const QMessageLogger* self, const char* msg, ... ) {
	self->critical(msg, );
}

void QMessageLogger_Fatal(const QMessageLogger* self, const char* msg, ... ) {
	self->fatal(msg, );
}

void QMessageLogger_Debug2(const QMessageLogger* self, QLoggingCategory* cat, const char* msg, ... ) {
	self->debug(*cat, msg, );
}

void QMessageLogger_Debug3(const QMessageLogger* self, CategoryFunction catFunc, const char* msg, ... ) {
	self->debug(catFunc, msg, );
}

void QMessageLogger_Info2(const QMessageLogger* self, QLoggingCategory* cat, const char* msg, ... ) {
	self->info(*cat, msg, );
}

void QMessageLogger_Info3(const QMessageLogger* self, CategoryFunction catFunc, const char* msg, ... ) {
	self->info(catFunc, msg, );
}

void QMessageLogger_Warning2(const QMessageLogger* self, QLoggingCategory* cat, const char* msg, ... ) {
	self->warning(*cat, msg, );
}

void QMessageLogger_Warning3(const QMessageLogger* self, CategoryFunction catFunc, const char* msg, ... ) {
	self->warning(catFunc, msg, );
}

void QMessageLogger_Critical2(const QMessageLogger* self, QLoggingCategory* cat, const char* msg, ... ) {
	self->critical(*cat, msg, );
}

void QMessageLogger_Critical3(const QMessageLogger* self, CategoryFunction catFunc, const char* msg, ... ) {
	self->critical(catFunc, msg, );
}

void QMessageLogger_Fatal2(const QMessageLogger* self, QLoggingCategory* cat, const char* msg, ... ) {
	self->fatal(*cat, msg, );
}

void QMessageLogger_Fatal3(const QMessageLogger* self, CategoryFunction catFunc, const char* msg, ... ) {
	self->fatal(catFunc, msg, );
}

QDebug* QMessageLogger_Debug4(const QMessageLogger* self) {
	return new QDebug(self->debug());
}

QDebug* QMessageLogger_DebugWithCat(const QMessageLogger* self, QLoggingCategory* cat) {
	return new QDebug(self->debug(*cat));
}

QDebug* QMessageLogger_DebugWithCatFunc(const QMessageLogger* self, CategoryFunction catFunc) {
	return new QDebug(self->debug(catFunc));
}

QDebug* QMessageLogger_Info4(const QMessageLogger* self) {
	return new QDebug(self->info());
}

QDebug* QMessageLogger_InfoWithCat(const QMessageLogger* self, QLoggingCategory* cat) {
	return new QDebug(self->info(*cat));
}

QDebug* QMessageLogger_InfoWithCatFunc(const QMessageLogger* self, CategoryFunction catFunc) {
	return new QDebug(self->info(catFunc));
}

QDebug* QMessageLogger_Warning4(const QMessageLogger* self) {
	return new QDebug(self->warning());
}

QDebug* QMessageLogger_WarningWithCat(const QMessageLogger* self, QLoggingCategory* cat) {
	return new QDebug(self->warning(*cat));
}

QDebug* QMessageLogger_WarningWithCatFunc(const QMessageLogger* self, CategoryFunction catFunc) {
	return new QDebug(self->warning(catFunc));
}

QDebug* QMessageLogger_Critical4(const QMessageLogger* self) {
	return new QDebug(self->critical());
}

QDebug* QMessageLogger_CriticalWithCat(const QMessageLogger* self, QLoggingCategory* cat) {
	return new QDebug(self->critical(*cat));
}

QDebug* QMessageLogger_CriticalWithCatFunc(const QMessageLogger* self, CategoryFunction catFunc) {
	return new QDebug(self->critical(catFunc));
}

QDebug* QMessageLogger_Fatal4(const QMessageLogger* self) {
	return new QDebug(self->fatal());
}

QDebug* QMessageLogger_FatalWithCat(const QMessageLogger* self, QLoggingCategory* cat) {
	return new QDebug(self->fatal(*cat));
}

QDebug* QMessageLogger_FatalWithCatFunc(const QMessageLogger* self, CategoryFunction catFunc) {
	return new QDebug(self->fatal(catFunc));
}

QNoDebug* QMessageLogger_NoDebug2(const QMessageLogger* self) {
	return new QNoDebug(self->noDebug());
}

void QMessageLogger_Delete(QMessageLogger* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QMessageLogger*>( self );
	} else {
		delete self;
	}
}

