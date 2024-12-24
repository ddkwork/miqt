// +build ignore

#define WORKAROUND_INNER_CLASS_DEFINITION_QtMocHelpers__NoType
#include <qtmochelpers.h>
#include "gen_qtmochelpers.h"

#ifndef _Bool
#define _Bool bool
#endif

void QtMocHelpers__NoType_Delete(QtMocHelpers__NoType* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QtMocHelpers::NoType*>( self );
	} else {
		delete self;
	}
}

