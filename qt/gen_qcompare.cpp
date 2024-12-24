// +build ignore

#include <QPartialOrdering>
#define WORKAROUND_INNER_CLASS_DEFINITION_partial_ordering
#define WORKAROUND_INNER_CLASS_DEFINITION_strong_ordering
#define WORKAROUND_INNER_CLASS_DEFINITION_weak_ordering
#include <qcompare.h>
#include "gen_qcompare.h"

partial_ordering* partial_ordering_new(const partial_ordering* param1) {
	return new Qt::partial_ordering(*param1);
}

void partial_ordering_Delete(partial_ordering* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<Qt::partial_ordering*>( self );
	} else {
		delete self;
	}
}

weak_ordering* weak_ordering_new(const weak_ordering* param1) {
	return new Qt::weak_ordering(*param1);
}

void weak_ordering_Delete(weak_ordering* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<Qt::weak_ordering*>( self );
	} else {
		delete self;
	}
}

strong_ordering* strong_ordering_new(const strong_ordering* param1) {
	return new Qt::strong_ordering(*param1);
}

void strong_ordering_Delete(strong_ordering* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<Qt::strong_ordering*>( self );
	} else {
		delete self;
	}
}

QPartialOrdering* QPartialOrdering_new(partial_ordering* order) {
	return new QPartialOrdering(*order);
}

QPartialOrdering* QPartialOrdering_new2(weak_ordering* stdorder) {
	return new QPartialOrdering(*stdorder);
}

QPartialOrdering* QPartialOrdering_new3(strong_ordering* stdorder) {
	return new QPartialOrdering(*stdorder);
}

QPartialOrdering* QPartialOrdering_new4(QPartialOrdering* param1) {
	return new QPartialOrdering(*param1);
}

void QPartialOrdering_Delete(QPartialOrdering* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QPartialOrdering*>( self );
	} else {
		delete self;
	}
}

