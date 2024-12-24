package qt

import (
	"unsafe"
)

type QTextOption__TabType int

const (
	QTextOption__LeftTab      QTextOption__TabType = 0
	QTextOption__RightTab     QTextOption__TabType = 1
	QTextOption__CenterTab    QTextOption__TabType = 2
	QTextOption__DelimiterTab QTextOption__TabType = 3
)

type QTextOption__WrapMode int

const (
	QTextOption__NoWrap                       QTextOption__WrapMode = 0
	QTextOption__WordWrap                     QTextOption__WrapMode = 1
	QTextOption__ManualWrap                   QTextOption__WrapMode = 2
	QTextOption__WrapAnywhere                 QTextOption__WrapMode = 3
	QTextOption__WrapAtWordBoundaryOrAnywhere QTextOption__WrapMode = 4
)

type QTextOption__Flag int

const (
	QTextOption__ShowTabsAndSpaces                     QTextOption__Flag = 1
	QTextOption__ShowLineAndParagraphSeparators        QTextOption__Flag = 2
	QTextOption__AddSpaceForLineAndParagraphSeparators QTextOption__Flag = 4
	QTextOption__SuppressColors                        QTextOption__Flag = 8
	QTextOption__ShowDocumentTerminator                QTextOption__Flag = 16
	QTextOption__ShowDefaultIgnorables                 QTextOption__Flag = 32
	QTextOption__DisableEmojiParsing                   QTextOption__Flag = 64
	QTextOption__IncludeTrailingSpaces                 QTextOption__Flag = 2147483648
)

type QTextOption struct {
	h          uintptr
	isSubclass bool
}

// NewQTextOption constructs a new QTextOption object.
func NewQTextOption() *QTextOption {

	ret := newQTextOption(QTextOption_new())
	ret.isSubclass = true
	return ret
}

// NewQTextOption2 constructs a new QTextOption object.
func NewQTextOption2(alignment AlignmentFlag) *QTextOption {

	ret := newQTextOption(QTextOption_new2((int)(alignment)))
	ret.isSubclass = true
	return ret
}

// NewQTextOption3 constructs a new QTextOption object.
func NewQTextOption3(o *QTextOption) *QTextOption {

	ret := newQTextOption(QTextOption_new3(o.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTextOption) OperatorAssign(o *QTextOption) {
	QTextOption_OperatorAssign(this.h, o.cPointer())
}

func (this *QTextOption) SetAlignment(alignment AlignmentFlag) {
	QTextOption_SetAlignment(this.h, (int)(alignment))
}

func (this *QTextOption) Alignment() AlignmentFlag {
	return (AlignmentFlag)(QTextOption_Alignment(this.h))
}

func (this *QTextOption) SetTextDirection(aDirection LayoutDirection) {
	QTextOption_SetTextDirection(this.h, (int)(aDirection))
}

func (this *QTextOption) TextDirection() LayoutDirection {
	return (LayoutDirection)(QTextOption_TextDirection(this.h))
}

func (this *QTextOption) SetWrapMode(wrap WrapMode) {
	QTextOption_SetWrapMode(this.h, wrap)
}

func (this *QTextOption) WrapMode() WrapMode {
	xxxxxxxxx
}

func (this *QTextOption) SetFlags(flags Flags) {
	QTextOption_SetFlags(this.h, flags)
}

func (this *QTextOption) Flags() Flags {
	xxxxxxxxx
}

func (this *QTextOption) SetTabStopDistance(tabStopDistance float64) {
	QTextOption_SetTabStopDistance(this.h, (double)(tabStopDistance))
}

func (this *QTextOption) TabStopDistance() float64 {
	return (float64)(QTextOption_TabStopDistance(this.h))
}

func (this *QTextOption) SetTabArray(tabStops []float64) {
	tabStops_CArray := (*[0xffff]double)(malloc(size_t(8 * len(tabStops))))
	defer free(unsafe.Pointer(tabStops_CArray))
	for i := range tabStops {
		tabStops_CArray[i] = (double)(tabStops[i])
	}
	tabStops_ma := struct_miqt_array{len: size_t(len(tabStops)), data: unsafe.Pointer(tabStops_CArray)}
	QTextOption_SetTabArray(this.h, tabStops_ma)
}

func (this *QTextOption) TabArray() []float64 {
	var _ma struct_miqt_array = QTextOption_TabArray(this.h)
	_ret := make([]float64, int(_ma.len))
	_outCast := (*[0xffff]double)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (float64)(_outCast[i])
	}
	return _ret
}

func (this *QTextOption) SetTabs(tabStops []Tab) {
	tabStops_CArray := (*[0xffff]Tab)(malloc(size_t(8 * len(tabStops))))
	defer free(unsafe.Pointer(tabStops_CArray))
	for i := range tabStops {
		tabStops_CArray[i] = tabStops[i]
	}
	tabStops_ma := struct_miqt_array{len: size_t(len(tabStops)), data: unsafe.Pointer(tabStops_CArray)}
	QTextOption_SetTabs(this.h, tabStops_ma)
}

func (this *QTextOption) Tabs() []Tab {
	var _ma struct_miqt_array = QTextOption_Tabs(this.h)
	_ret := make([]Tab, int(_ma.len))
	_outCast := (*[0xffff]Tab)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		xxxxxxxxx
	}
	return _ret
}

func (this *QTextOption) SetUseDesignMetrics(b bool) {
	QTextOption_SetUseDesignMetrics(this.h, (bool)(b))
}

func (this *QTextOption) UseDesignMetrics() bool {
	return (bool)(QTextOption_UseDesignMetrics(this.h))
}

type QTextOption__Tab struct {
	h          uintptr
	isSubclass bool
}

// NewQTextOption__Tab constructs a new QTextOption::Tab object.
func NewQTextOption__Tab() *QTextOption__Tab {

	ret := newQTextOption__Tab(QTextOption__Tab_new())
	ret.isSubclass = true
	return ret
}

// NewQTextOption__Tab2 constructs a new QTextOption::Tab object.
func NewQTextOption__Tab2(pos float64, tabType TabType) *QTextOption__Tab {

	ret := newQTextOption__Tab(QTextOption__Tab_new2((double)(pos), tabType))
	ret.isSubclass = true
	return ret
}

// NewQTextOption__Tab3 constructs a new QTextOption::Tab object.
func NewQTextOption__Tab3(pos float64, tabType TabType, delim QChar) *QTextOption__Tab {

	ret := newQTextOption__Tab(QTextOption__Tab_new3((double)(pos), tabType, delim.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTextOption__Tab) OperatorEqual(other *Tab) bool {
	return (bool)(QTextOption__Tab_OperatorEqual(this.h, other))
}

func (this *QTextOption__Tab) OperatorNotEqual(other *Tab) bool {
	return (bool)(QTextOption__Tab_OperatorNotEqual(this.h, other))
}
