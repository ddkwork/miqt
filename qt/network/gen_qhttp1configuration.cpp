// +build ignore

#include <QHttp1Configuration>
#include <qhttp1configuration.h>
#include "gen_qhttp1configuration.h"

QHttp1Configuration* QHttp1Configuration_new() {
	return new QHttp1Configuration();
}

QHttp1Configuration* QHttp1Configuration_new2(QHttp1Configuration* other) {
	return new QHttp1Configuration(*other);
}

void QHttp1Configuration_OperatorAssign(QHttp1Configuration* self, QHttp1Configuration* other) {
	self->operator=(*other);
}

void QHttp1Configuration_SetNumberOfConnectionsPerHost(QHttp1Configuration* self, ptrdiff_t amount) {
	self->setNumberOfConnectionsPerHost((qsizetype)(amount));
}

ptrdiff_t QHttp1Configuration_NumberOfConnectionsPerHost(const QHttp1Configuration* self) {
	qsizetype _ret = self->numberOfConnectionsPerHost();
	return static_cast<ptrdiff_t>(_ret);
}

void QHttp1Configuration_Swap(QHttp1Configuration* self, QHttp1Configuration* other) {
	self->swap(*other);
}

void QHttp1Configuration_Delete(QHttp1Configuration* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QHttp1Configuration*>( self );
	} else {
		delete self;
	}
}

