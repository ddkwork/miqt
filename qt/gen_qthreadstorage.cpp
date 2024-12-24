// +build ignore

#include <QThreadStorageData>
#include <qthreadstorage.h>
#include "gen_qthreadstorage.h"

void QThreadStorageData_Delete(QThreadStorageData* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QThreadStorageData*>( self );
	} else {
		delete self;
	}
}

