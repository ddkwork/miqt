// +build ignore

#include <QLatin1StringMatcher>
#include <qlatin1stringmatcher.h>
#include "gen_qlatin1stringmatcher.h"

#ifndef _Bool
#define _Bool bool
#endif

QLatin1StringMatcher* QLatin1StringMatcher_new() {
	return new QLatin1StringMatcher();
}

void QLatin1StringMatcher_SetCaseSensitivity(QLatin1StringMatcher* self, int cs) {
	self->setCaseSensitivity(static_cast<Qt::CaseSensitivity>(cs));
}

int QLatin1StringMatcher_CaseSensitivity(const QLatin1StringMatcher* self) {
	Qt::CaseSensitivity _ret = self->caseSensitivity();
	return static_cast<int>(_ret);
}

void QLatin1StringMatcher_Delete(QLatin1StringMatcher* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QLatin1StringMatcher*>( self );
	} else {
		delete self;
	}
}

