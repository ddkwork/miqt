// +build ignore

#include <QAnyStringView>
#include <QByteArrayView>
#include <QFormDataBuilder>
#include <QFormDataPartBuilder>
#include <QHttpHeaders>
#include <QIODevice>
#include <qformdatabuilder.h>
#include "gen_qformdatabuilder.h"
QFormDataPartBuilder* QFormDataPartBuilder_new() {
return new QFormDataPartBuilder();
}
QFormDataPartBuilder* QFormDataPartBuilder_new2(QFormDataPartBuilder* param1) {
return new QFormDataPartBuilder(*param1);
}
void QFormDataPartBuilder_Swap(QFormDataPartBuilder* self, QFormDataPartBuilder* other) {
	self->swap(*other);
}
QFormDataPartBuilder* QFormDataPartBuilder_SetBody(QFormDataPartBuilder* self, QByteArrayView* data) {
	return new QFormDataPartBuilder(self->setBody(*data));
}
QFormDataPartBuilder* QFormDataPartBuilder_SetBodyDevice(QFormDataPartBuilder* self, QIODevice* body) {
	return new QFormDataPartBuilder(self->setBodyDevice(body));
}
QFormDataPartBuilder* QFormDataPartBuilder_SetHeaders(QFormDataPartBuilder* self, QHttpHeaders* headers) {
	return new QFormDataPartBuilder(self->setHeaders(*headers));
}
QFormDataPartBuilder* QFormDataPartBuilder_SetBody2(QFormDataPartBuilder* self, QByteArrayView* data, QAnyStringView* fileName) {
	return new QFormDataPartBuilder(self->setBody(*data, *fileName));
}
QFormDataPartBuilder* QFormDataPartBuilder_SetBody3(QFormDataPartBuilder* self, QByteArrayView* data, QAnyStringView* fileName, QAnyStringView* mimeType) {
	return new QFormDataPartBuilder(self->setBody(*data, *fileName, *mimeType));
}
QFormDataPartBuilder* QFormDataPartBuilder_SetBodyDevice2(QFormDataPartBuilder* self, QIODevice* body, QAnyStringView* fileName) {
	return new QFormDataPartBuilder(self->setBodyDevice(body, *fileName));
}
QFormDataPartBuilder* QFormDataPartBuilder_SetBodyDevice3(QFormDataPartBuilder* self, QIODevice* body, QAnyStringView* fileName, QAnyStringView* mimeType) {
	return new QFormDataPartBuilder(self->setBodyDevice(body, *fileName, *mimeType));
}
void QFormDataPartBuilder_Delete(QFormDataPartBuilder* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QFormDataPartBuilder*>( self );
	} else {
		delete self;
	}
}
QFormDataBuilder* QFormDataBuilder_new() {
return new QFormDataBuilder();
}
void QFormDataBuilder_Swap(QFormDataBuilder* self, QFormDataBuilder* other) {
	self->swap(*other);
}
QFormDataPartBuilder* QFormDataBuilder_Part(QFormDataBuilder* self, QAnyStringView* name) {
	return new QFormDataPartBuilder(self->part(*name));
}
void QFormDataBuilder_Delete(QFormDataBuilder* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QFormDataBuilder*>( self );
	} else {
		delete self;
	}
}
