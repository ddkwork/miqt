// +build ignore

#include <QTileRules>
#include <qdrawutil.h>
#include "gen_qdrawutil.h"

#ifndef _Bool
#define _Bool bool
#endif

void _GUID_Delete(_GUID* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<_GUID*>( self );
	} else {
		delete self;
	}
}

void type_info_Delete(type_info* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<type_info*>( self );
	} else {
		delete self;
	}
}

QTileRules* QTileRules_new(int horizontalRule, int verticalRule) {
	return new QTileRules(static_cast<Qt::TileRule>(horizontalRule), static_cast<Qt::TileRule>(verticalRule));
}

QTileRules* QTileRules_new2() {
	return new QTileRules();
}

QTileRules* QTileRules_new3(QTileRules* param1) {
	return new QTileRules(*param1);
}

QTileRules* QTileRules_new4(int rule) {
	return new QTileRules(static_cast<Qt::TileRule>(rule));
}

void QTileRules_Delete(QTileRules* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QTileRules*>( self );
	} else {
		delete self;
	}
}

