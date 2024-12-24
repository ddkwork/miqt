package qt

import (
	"unsafe"
)

type QLoggingCategory struct {
	h          uintptr
	isSubclass bool
}

// NewQLoggingCategory constructs a new QLoggingCategory object.
func NewQLoggingCategory(category string) *QLoggingCategory {
	category_Cstring := CString(category)
	defer free(unsafe.Pointer(category_Cstring))

	ret := newQLoggingCategory(QLoggingCategory_new(category_Cstring))
	ret.isSubclass = true
	return ret
}

func (this *QLoggingCategory) IsDebugEnabled() bool {
	return (bool)(QLoggingCategory_IsDebugEnabled(this.h))
}

func (this *QLoggingCategory) IsInfoEnabled() bool {
	return (bool)(QLoggingCategory_IsInfoEnabled(this.h))
}

func (this *QLoggingCategory) IsWarningEnabled() bool {
	return (bool)(QLoggingCategory_IsWarningEnabled(this.h))
}

func (this *QLoggingCategory) IsCriticalEnabled() bool {
	return (bool)(QLoggingCategory_IsCriticalEnabled(this.h))
}

func (this *QLoggingCategory) CategoryName() string {
	_ret := QLoggingCategory_CategoryName(this.h)
	return GoString(_ret)
}

func (this *QLoggingCategory) OperatorCall() *QLoggingCategory {
	return newQLoggingCategory(QLoggingCategory_OperatorCall(this.h))
}

func (this *QLoggingCategory) OperatorCall2() *QLoggingCategory {
	return newQLoggingCategory(QLoggingCategory_OperatorCall2(this.h))
}

func QLoggingCategory_DefaultCategory() *QLoggingCategory {
	return newQLoggingCategory(QLoggingCategory_DefaultCategory())
}

func QLoggingCategory_InstallFilter(param1 CategoryFilter) CategoryFilter {
	xxxxxxxxx
}

func QLoggingCategory_SetFilterRules(rules string) {
	rules_ms := struct_miqt_string{}
	rules_ms.data = CString(rules)
	rules_ms.len = size_t(len(rules))
	defer free(unsafe.Pointer(rules_ms.data))
	QLoggingCategory_SetFilterRules(rules_ms)
}
