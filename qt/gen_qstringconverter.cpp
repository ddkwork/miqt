// +build ignore

#include <QAnyStringView>
#include <QByteArrayView>
#include <QChar>
#include <QStringConverter>
#include <QStringConverterBase>
#include <QStringDecoder>
#include <QStringEncoder>
#include <qstringconverter.h>
#include "gen_qstringconverter.h"

QStringEncoder* QStringEncoder_new() {
	return new QStringEncoder();
}

QStringEncoder* QStringEncoder_new2(Encoding encoding) {
	return new QStringEncoder(encoding);
}

QStringEncoder* QStringEncoder_new3(QAnyStringView* name) {
	return new QStringEncoder(*name);
}

QStringEncoder* QStringEncoder_new4(Encoding encoding, Flags flags) {
	return new QStringEncoder(encoding, flags);
}

QStringEncoder* QStringEncoder_new5(QAnyStringView* name, Flags flags) {
	return new QStringEncoder(*name, flags);
}

void QStringEncoder_virtbase(QStringEncoder* src, QStringConverter** outptr_QStringConverter) {
	*outptr_QStringConverter = static_cast<QStringConverter*>(src);
}

ptrdiff_t QStringEncoder_RequiredSpace(const QStringEncoder* self, ptrdiff_t inputLength) {
	qsizetype _ret = self->requiredSpace((qsizetype)(inputLength));
	return static_cast<ptrdiff_t>(_ret);
}

void QStringEncoder_Delete(QStringEncoder* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QStringEncoder*>( self );
	} else {
		delete self;
	}
}

QStringDecoder* QStringDecoder_new(Encoding encoding) {
	return new QStringDecoder(encoding);
}

QStringDecoder* QStringDecoder_new2() {
	return new QStringDecoder();
}

QStringDecoder* QStringDecoder_new3(QAnyStringView* name) {
	return new QStringDecoder(*name);
}

QStringDecoder* QStringDecoder_new4(Encoding encoding, Flags flags) {
	return new QStringDecoder(encoding, flags);
}

QStringDecoder* QStringDecoder_new5(QAnyStringView* name, Flags f) {
	return new QStringDecoder(*name, f);
}

void QStringDecoder_virtbase(QStringDecoder* src, QStringConverter** outptr_QStringConverter) {
	*outptr_QStringConverter = static_cast<QStringConverter*>(src);
}

ptrdiff_t QStringDecoder_RequiredSpace(const QStringDecoder* self, ptrdiff_t inputLength) {
	qsizetype _ret = self->requiredSpace((qsizetype)(inputLength));
	return static_cast<ptrdiff_t>(_ret);
}

QChar* QStringDecoder_AppendToBuffer(QStringDecoder* self, QChar* out, QByteArrayView* ba) {
	return self->appendToBuffer(out, *ba);
}

QStringDecoder* QStringDecoder_DecoderForHtml(QByteArrayView* data) {
	return new QStringDecoder(QStringDecoder::decoderForHtml(*data));
}

void QStringDecoder_Delete(QStringDecoder* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QStringDecoder*>( self );
	} else {
		delete self;
	}
}

