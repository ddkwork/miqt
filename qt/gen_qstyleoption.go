package qt

import (
	"unsafe"
)

type QStyleOption__OptionType int

const (
	QStyleOption__SO_Default           QStyleOption__OptionType = 0
	QStyleOption__SO_FocusRect         QStyleOption__OptionType = 1
	QStyleOption__SO_Button            QStyleOption__OptionType = 2
	QStyleOption__SO_Tab               QStyleOption__OptionType = 3
	QStyleOption__SO_MenuItem          QStyleOption__OptionType = 4
	QStyleOption__SO_Frame             QStyleOption__OptionType = 5
	QStyleOption__SO_ProgressBar       QStyleOption__OptionType = 6
	QStyleOption__SO_ToolBox           QStyleOption__OptionType = 7
	QStyleOption__SO_Header            QStyleOption__OptionType = 8
	QStyleOption__SO_DockWidget        QStyleOption__OptionType = 9
	QStyleOption__SO_ViewItem          QStyleOption__OptionType = 10
	QStyleOption__SO_TabWidgetFrame    QStyleOption__OptionType = 11
	QStyleOption__SO_TabBarBase        QStyleOption__OptionType = 12
	QStyleOption__SO_RubberBand        QStyleOption__OptionType = 13
	QStyleOption__SO_ToolBar           QStyleOption__OptionType = 14
	QStyleOption__SO_GraphicsItem      QStyleOption__OptionType = 15
	QStyleOption__SO_Complex           QStyleOption__OptionType = 983040
	QStyleOption__SO_Slider            QStyleOption__OptionType = 983041
	QStyleOption__SO_SpinBox           QStyleOption__OptionType = 983042
	QStyleOption__SO_ToolButton        QStyleOption__OptionType = 983043
	QStyleOption__SO_ComboBox          QStyleOption__OptionType = 983044
	QStyleOption__SO_TitleBar          QStyleOption__OptionType = 983045
	QStyleOption__SO_GroupBox          QStyleOption__OptionType = 983046
	QStyleOption__SO_SizeGrip          QStyleOption__OptionType = 983047
	QStyleOption__SO_CustomBase        QStyleOption__OptionType = 3840
	QStyleOption__SO_ComplexCustomBase QStyleOption__OptionType = 251658240
)

type QStyleOption__StyleOptionType int

const (
	QStyleOption__Type QStyleOption__StyleOptionType = 0
)

type QStyleOption__StyleOptionVersion int

const (
	QStyleOption__Version QStyleOption__StyleOptionVersion = 1
)

type QStyleOptionFocusRect__StyleOptionType int

const (
	QStyleOptionFocusRect__Type QStyleOptionFocusRect__StyleOptionType = 1
)

type QStyleOptionFocusRect__StyleOptionVersion int

const (
	QStyleOptionFocusRect__Version QStyleOptionFocusRect__StyleOptionVersion = 1
)

type QStyleOptionFrame__StyleOptionType int

const (
	QStyleOptionFrame__Type QStyleOptionFrame__StyleOptionType = 5
)

type QStyleOptionFrame__StyleOptionVersion int

const (
	QStyleOptionFrame__Version QStyleOptionFrame__StyleOptionVersion = 1
)

type QStyleOptionFrame__FrameFeature int

const (
	QStyleOptionFrame__None    QStyleOptionFrame__FrameFeature = 0
	QStyleOptionFrame__Flat    QStyleOptionFrame__FrameFeature = 1
	QStyleOptionFrame__Rounded QStyleOptionFrame__FrameFeature = 2
)

type QStyleOptionTabWidgetFrame__StyleOptionType int

const (
	QStyleOptionTabWidgetFrame__Type QStyleOptionTabWidgetFrame__StyleOptionType = 11
)

type QStyleOptionTabWidgetFrame__StyleOptionVersion int

const (
	QStyleOptionTabWidgetFrame__Version QStyleOptionTabWidgetFrame__StyleOptionVersion = 1
)

type QStyleOptionTabBarBase__StyleOptionType int

const (
	QStyleOptionTabBarBase__Type QStyleOptionTabBarBase__StyleOptionType = 12
)

type QStyleOptionTabBarBase__StyleOptionVersion int

const (
	QStyleOptionTabBarBase__Version QStyleOptionTabBarBase__StyleOptionVersion = 1
)

type QStyleOptionHeader__StyleOptionType int

const (
	QStyleOptionHeader__Type QStyleOptionHeader__StyleOptionType = 8
)

type QStyleOptionHeader__StyleOptionVersion int

const (
	QStyleOptionHeader__Version QStyleOptionHeader__StyleOptionVersion = 1
)

type QStyleOptionHeader__SectionPosition int

const (
	QStyleOptionHeader__Beginning      QStyleOptionHeader__SectionPosition = 0
	QStyleOptionHeader__Middle         QStyleOptionHeader__SectionPosition = 1
	QStyleOptionHeader__End            QStyleOptionHeader__SectionPosition = 2
	QStyleOptionHeader__OnlyOneSection QStyleOptionHeader__SectionPosition = 3
)

type QStyleOptionHeader__SelectedPosition int

const (
	QStyleOptionHeader__NotAdjacent                QStyleOptionHeader__SelectedPosition = 0
	QStyleOptionHeader__NextIsSelected             QStyleOptionHeader__SelectedPosition = 1
	QStyleOptionHeader__PreviousIsSelected         QStyleOptionHeader__SelectedPosition = 2
	QStyleOptionHeader__NextAndPreviousAreSelected QStyleOptionHeader__SelectedPosition = 3
)

type QStyleOptionHeader__SortIndicator int

const (
	QStyleOptionHeader__None     QStyleOptionHeader__SortIndicator = 0
	QStyleOptionHeader__SortUp   QStyleOptionHeader__SortIndicator = 1
	QStyleOptionHeader__SortDown QStyleOptionHeader__SortIndicator = 2
)

type QStyleOptionHeaderV2__StyleOptionType int

const (
	QStyleOptionHeaderV2__Type QStyleOptionHeaderV2__StyleOptionType = 8
)

type QStyleOptionHeaderV2__StyleOptionVersion int

const (
	QStyleOptionHeaderV2__Version QStyleOptionHeaderV2__StyleOptionVersion = 2
)

type QStyleOptionButton__StyleOptionType int

const (
	QStyleOptionButton__Type QStyleOptionButton__StyleOptionType = 2
)

type QStyleOptionButton__StyleOptionVersion int

const (
	QStyleOptionButton__Version QStyleOptionButton__StyleOptionVersion = 1
)

type QStyleOptionButton__ButtonFeature int

const (
	QStyleOptionButton__None              QStyleOptionButton__ButtonFeature = 0
	QStyleOptionButton__Flat              QStyleOptionButton__ButtonFeature = 1
	QStyleOptionButton__HasMenu           QStyleOptionButton__ButtonFeature = 2
	QStyleOptionButton__DefaultButton     QStyleOptionButton__ButtonFeature = 4
	QStyleOptionButton__AutoDefaultButton QStyleOptionButton__ButtonFeature = 8
	QStyleOptionButton__CommandLinkButton QStyleOptionButton__ButtonFeature = 16
)

type QStyleOptionTab__StyleOptionType int

const (
	QStyleOptionTab__Type QStyleOptionTab__StyleOptionType = 3
)

type QStyleOptionTab__StyleOptionVersion int

const (
	QStyleOptionTab__Version QStyleOptionTab__StyleOptionVersion = 1
)

type QStyleOptionTab__TabPosition int

const (
	QStyleOptionTab__Beginning  QStyleOptionTab__TabPosition = 0
	QStyleOptionTab__Middle     QStyleOptionTab__TabPosition = 1
	QStyleOptionTab__End        QStyleOptionTab__TabPosition = 2
	QStyleOptionTab__OnlyOneTab QStyleOptionTab__TabPosition = 3
	QStyleOptionTab__Moving     QStyleOptionTab__TabPosition = 4
)

type QStyleOptionTab__SelectedPosition int

const (
	QStyleOptionTab__NotAdjacent        QStyleOptionTab__SelectedPosition = 0
	QStyleOptionTab__NextIsSelected     QStyleOptionTab__SelectedPosition = 1
	QStyleOptionTab__PreviousIsSelected QStyleOptionTab__SelectedPosition = 2
)

type QStyleOptionTab__CornerWidget int

const (
	QStyleOptionTab__NoCornerWidgets   QStyleOptionTab__CornerWidget = 0
	QStyleOptionTab__LeftCornerWidget  QStyleOptionTab__CornerWidget = 1
	QStyleOptionTab__RightCornerWidget QStyleOptionTab__CornerWidget = 2
)

type QStyleOptionTab__TabFeature int

const (
	QStyleOptionTab__None            QStyleOptionTab__TabFeature = 0
	QStyleOptionTab__HasFrame        QStyleOptionTab__TabFeature = 1
	QStyleOptionTab__MinimumSizeHint QStyleOptionTab__TabFeature = 2
)

type QStyleOptionToolBar__StyleOptionType int

const (
	QStyleOptionToolBar__Type QStyleOptionToolBar__StyleOptionType = 14
)

type QStyleOptionToolBar__StyleOptionVersion int

const (
	QStyleOptionToolBar__Version QStyleOptionToolBar__StyleOptionVersion = 1
)

type QStyleOptionToolBar__ToolBarPosition int

const (
	QStyleOptionToolBar__Beginning QStyleOptionToolBar__ToolBarPosition = 0
	QStyleOptionToolBar__Middle    QStyleOptionToolBar__ToolBarPosition = 1
	QStyleOptionToolBar__End       QStyleOptionToolBar__ToolBarPosition = 2
	QStyleOptionToolBar__OnlyOne   QStyleOptionToolBar__ToolBarPosition = 3
)

type QStyleOptionToolBar__ToolBarFeature int

const (
	QStyleOptionToolBar__None    QStyleOptionToolBar__ToolBarFeature = 0
	QStyleOptionToolBar__Movable QStyleOptionToolBar__ToolBarFeature = 1
)

type QStyleOptionProgressBar__StyleOptionType int

const (
	QStyleOptionProgressBar__Type QStyleOptionProgressBar__StyleOptionType = 6
)

type QStyleOptionProgressBar__StyleOptionVersion int

const (
	QStyleOptionProgressBar__Version QStyleOptionProgressBar__StyleOptionVersion = 1
)

type QStyleOptionMenuItem__StyleOptionType int

const (
	QStyleOptionMenuItem__Type QStyleOptionMenuItem__StyleOptionType = 4
)

type QStyleOptionMenuItem__StyleOptionVersion int

const (
	QStyleOptionMenuItem__Version QStyleOptionMenuItem__StyleOptionVersion = 1
)

type QStyleOptionMenuItem__MenuItemType int

const (
	QStyleOptionMenuItem__Normal      QStyleOptionMenuItem__MenuItemType = 0
	QStyleOptionMenuItem__DefaultItem QStyleOptionMenuItem__MenuItemType = 1
	QStyleOptionMenuItem__Separator   QStyleOptionMenuItem__MenuItemType = 2
	QStyleOptionMenuItem__SubMenu     QStyleOptionMenuItem__MenuItemType = 3
	QStyleOptionMenuItem__Scroller    QStyleOptionMenuItem__MenuItemType = 4
	QStyleOptionMenuItem__TearOff     QStyleOptionMenuItem__MenuItemType = 5
	QStyleOptionMenuItem__Margin      QStyleOptionMenuItem__MenuItemType = 6
	QStyleOptionMenuItem__EmptyArea   QStyleOptionMenuItem__MenuItemType = 7
)

type QStyleOptionMenuItem__CheckType int

const (
	QStyleOptionMenuItem__NotCheckable QStyleOptionMenuItem__CheckType = 0
	QStyleOptionMenuItem__Exclusive    QStyleOptionMenuItem__CheckType = 1
	QStyleOptionMenuItem__NonExclusive QStyleOptionMenuItem__CheckType = 2
)

type QStyleOptionDockWidget__StyleOptionType int

const (
	QStyleOptionDockWidget__Type QStyleOptionDockWidget__StyleOptionType = 9
)

type QStyleOptionDockWidget__StyleOptionVersion int

const (
	QStyleOptionDockWidget__Version QStyleOptionDockWidget__StyleOptionVersion = 1
)

type QStyleOptionViewItem__StyleOptionType int

const (
	QStyleOptionViewItem__Type QStyleOptionViewItem__StyleOptionType = 10
)

type QStyleOptionViewItem__StyleOptionVersion int

const (
	QStyleOptionViewItem__Version QStyleOptionViewItem__StyleOptionVersion = 1
)

type QStyleOptionViewItem__Position int

const (
	QStyleOptionViewItem__Left   QStyleOptionViewItem__Position = 0
	QStyleOptionViewItem__Right  QStyleOptionViewItem__Position = 1
	QStyleOptionViewItem__Top    QStyleOptionViewItem__Position = 2
	QStyleOptionViewItem__Bottom QStyleOptionViewItem__Position = 3
)

type QStyleOptionViewItem__ViewItemFeature int

const (
	QStyleOptionViewItem__None              QStyleOptionViewItem__ViewItemFeature = 0
	QStyleOptionViewItem__WrapText          QStyleOptionViewItem__ViewItemFeature = 1
	QStyleOptionViewItem__Alternate         QStyleOptionViewItem__ViewItemFeature = 2
	QStyleOptionViewItem__HasCheckIndicator QStyleOptionViewItem__ViewItemFeature = 4
	QStyleOptionViewItem__HasDisplay        QStyleOptionViewItem__ViewItemFeature = 8
	QStyleOptionViewItem__HasDecoration     QStyleOptionViewItem__ViewItemFeature = 16
)

type QStyleOptionViewItem__ViewItemPosition int

const (
	QStyleOptionViewItem__Invalid   QStyleOptionViewItem__ViewItemPosition = 0
	QStyleOptionViewItem__Beginning QStyleOptionViewItem__ViewItemPosition = 1
	QStyleOptionViewItem__Middle    QStyleOptionViewItem__ViewItemPosition = 2
	QStyleOptionViewItem__End       QStyleOptionViewItem__ViewItemPosition = 3
	QStyleOptionViewItem__OnlyOne   QStyleOptionViewItem__ViewItemPosition = 4
)

type QStyleOptionToolBox__StyleOptionType int

const (
	QStyleOptionToolBox__Type QStyleOptionToolBox__StyleOptionType = 7
)

type QStyleOptionToolBox__StyleOptionVersion int

const (
	QStyleOptionToolBox__Version QStyleOptionToolBox__StyleOptionVersion = 1
)

type QStyleOptionToolBox__TabPosition int

const (
	QStyleOptionToolBox__Beginning  QStyleOptionToolBox__TabPosition = 0
	QStyleOptionToolBox__Middle     QStyleOptionToolBox__TabPosition = 1
	QStyleOptionToolBox__End        QStyleOptionToolBox__TabPosition = 2
	QStyleOptionToolBox__OnlyOneTab QStyleOptionToolBox__TabPosition = 3
)

type QStyleOptionToolBox__SelectedPosition int

const (
	QStyleOptionToolBox__NotAdjacent        QStyleOptionToolBox__SelectedPosition = 0
	QStyleOptionToolBox__NextIsSelected     QStyleOptionToolBox__SelectedPosition = 1
	QStyleOptionToolBox__PreviousIsSelected QStyleOptionToolBox__SelectedPosition = 2
)

type QStyleOptionRubberBand__StyleOptionType int

const (
	QStyleOptionRubberBand__Type QStyleOptionRubberBand__StyleOptionType = 13
)

type QStyleOptionRubberBand__StyleOptionVersion int

const (
	QStyleOptionRubberBand__Version QStyleOptionRubberBand__StyleOptionVersion = 1
)

type QStyleOptionComplex__StyleOptionType int

const (
	QStyleOptionComplex__Type QStyleOptionComplex__StyleOptionType = 983040
)

type QStyleOptionComplex__StyleOptionVersion int

const (
	QStyleOptionComplex__Version QStyleOptionComplex__StyleOptionVersion = 1
)

type QStyleOptionSlider__StyleOptionType int

const (
	QStyleOptionSlider__Type QStyleOptionSlider__StyleOptionType = 983041
)

type QStyleOptionSlider__StyleOptionVersion int

const (
	QStyleOptionSlider__Version QStyleOptionSlider__StyleOptionVersion = 1
)

type QStyleOptionSpinBox__StyleOptionType int

const (
	QStyleOptionSpinBox__Type QStyleOptionSpinBox__StyleOptionType = 983042
)

type QStyleOptionSpinBox__StyleOptionVersion int

const (
	QStyleOptionSpinBox__Version QStyleOptionSpinBox__StyleOptionVersion = 1
)

type QStyleOptionToolButton__StyleOptionType int

const (
	QStyleOptionToolButton__Type QStyleOptionToolButton__StyleOptionType = 983043
)

type QStyleOptionToolButton__StyleOptionVersion int

const (
	QStyleOptionToolButton__Version QStyleOptionToolButton__StyleOptionVersion = 1
)

type QStyleOptionToolButton__ToolButtonFeature int

const (
	QStyleOptionToolButton__None            QStyleOptionToolButton__ToolButtonFeature = 0
	QStyleOptionToolButton__Arrow           QStyleOptionToolButton__ToolButtonFeature = 1
	QStyleOptionToolButton__Menu            QStyleOptionToolButton__ToolButtonFeature = 4
	QStyleOptionToolButton__MenuButtonPopup QStyleOptionToolButton__ToolButtonFeature = 4
	QStyleOptionToolButton__PopupDelay      QStyleOptionToolButton__ToolButtonFeature = 8
	QStyleOptionToolButton__HasMenu         QStyleOptionToolButton__ToolButtonFeature = 16
)

type QStyleOptionComboBox__StyleOptionType int

const (
	QStyleOptionComboBox__Type QStyleOptionComboBox__StyleOptionType = 983044
)

type QStyleOptionComboBox__StyleOptionVersion int

const (
	QStyleOptionComboBox__Version QStyleOptionComboBox__StyleOptionVersion = 1
)

type QStyleOptionTitleBar__StyleOptionType int

const (
	QStyleOptionTitleBar__Type QStyleOptionTitleBar__StyleOptionType = 983045
)

type QStyleOptionTitleBar__StyleOptionVersion int

const (
	QStyleOptionTitleBar__Version QStyleOptionTitleBar__StyleOptionVersion = 1
)

type QStyleOptionGroupBox__StyleOptionType int

const (
	QStyleOptionGroupBox__Type QStyleOptionGroupBox__StyleOptionType = 983046
)

type QStyleOptionGroupBox__StyleOptionVersion int

const (
	QStyleOptionGroupBox__Version QStyleOptionGroupBox__StyleOptionVersion = 1
)

type QStyleOptionSizeGrip__StyleOptionType int

const (
	QStyleOptionSizeGrip__Type QStyleOptionSizeGrip__StyleOptionType = 983047
)

type QStyleOptionSizeGrip__StyleOptionVersion int

const (
	QStyleOptionSizeGrip__Version QStyleOptionSizeGrip__StyleOptionVersion = 1
)

type QStyleOptionGraphicsItem__StyleOptionType int

const (
	QStyleOptionGraphicsItem__Type QStyleOptionGraphicsItem__StyleOptionType = 15
)

type QStyleOptionGraphicsItem__StyleOptionVersion int

const (
	QStyleOptionGraphicsItem__Version QStyleOptionGraphicsItem__StyleOptionVersion = 1
)

type QStyleHintReturn__HintReturnType int

const (
	QStyleHintReturn__SH_Default QStyleHintReturn__HintReturnType = 61440
	QStyleHintReturn__SH_Mask    QStyleHintReturn__HintReturnType = 61441
	QStyleHintReturn__SH_Variant QStyleHintReturn__HintReturnType = 61442
)

type QStyleHintReturn__StyleOptionType int

const (
	QStyleHintReturn__Type QStyleHintReturn__StyleOptionType = 61440
)

type QStyleHintReturn__StyleOptionVersion int

const (
	QStyleHintReturn__Version QStyleHintReturn__StyleOptionVersion = 1
)

type QStyleHintReturnMask__StyleOptionType int

const (
	QStyleHintReturnMask__Type QStyleHintReturnMask__StyleOptionType = 61441
)

type QStyleHintReturnMask__StyleOptionVersion int

const (
	QStyleHintReturnMask__Version QStyleHintReturnMask__StyleOptionVersion = 1
)

type QStyleHintReturnVariant__StyleOptionType int

const (
	QStyleHintReturnVariant__Type QStyleHintReturnVariant__StyleOptionType = 61442
)

type QStyleHintReturnVariant__StyleOptionVersion int

const (
	QStyleHintReturnVariant__Version QStyleHintReturnVariant__StyleOptionVersion = 1
)

type QStyleOption struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOption constructs a new QStyleOption object.
func NewQStyleOption() *QStyleOption {

	ret := newQStyleOption(QStyleOption_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOption2 constructs a new QStyleOption object.
func NewQStyleOption2(other *QStyleOption) *QStyleOption {

	ret := newQStyleOption(QStyleOption_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQStyleOption3 constructs a new QStyleOption object.
func NewQStyleOption3(version int) *QStyleOption {

	ret := newQStyleOption(QStyleOption_new3((int)(version)))
	ret.isSubclass = true
	return ret
}

// NewQStyleOption4 constructs a new QStyleOption object.
func NewQStyleOption4(version int, typeVal int) *QStyleOption {

	ret := newQStyleOption(QStyleOption_new4((int)(version), (int)(typeVal)))
	ret.isSubclass = true
	return ret
}

func (this *QStyleOption) InitFrom(w *QWidget) {
	QStyleOption_InitFrom(this.h, w.cPointer())
}

func (this *QStyleOption) OperatorAssign(other *QStyleOption) {
	QStyleOption_OperatorAssign(this.h, other.cPointer())
}

type QStyleOptionFocusRect struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionFocusRect constructs a new QStyleOptionFocusRect object.
func NewQStyleOptionFocusRect() *QStyleOptionFocusRect {

	ret := newQStyleOptionFocusRect(QStyleOptionFocusRect_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionFocusRect2 constructs a new QStyleOptionFocusRect object.
func NewQStyleOptionFocusRect2(other *QStyleOptionFocusRect) *QStyleOptionFocusRect {

	ret := newQStyleOptionFocusRect(QStyleOptionFocusRect_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

type QStyleOptionFrame struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionFrame constructs a new QStyleOptionFrame object.
func NewQStyleOptionFrame() *QStyleOptionFrame {

	ret := newQStyleOptionFrame(QStyleOptionFrame_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionFrame2 constructs a new QStyleOptionFrame object.
func NewQStyleOptionFrame2(other *QStyleOptionFrame) *QStyleOptionFrame {

	ret := newQStyleOptionFrame(QStyleOptionFrame_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

type QStyleOptionTabWidgetFrame struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionTabWidgetFrame constructs a new QStyleOptionTabWidgetFrame object.
func NewQStyleOptionTabWidgetFrame() *QStyleOptionTabWidgetFrame {

	ret := newQStyleOptionTabWidgetFrame(QStyleOptionTabWidgetFrame_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionTabWidgetFrame2 constructs a new QStyleOptionTabWidgetFrame object.
func NewQStyleOptionTabWidgetFrame2(other *QStyleOptionTabWidgetFrame) *QStyleOptionTabWidgetFrame {

	ret := newQStyleOptionTabWidgetFrame(QStyleOptionTabWidgetFrame_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

type QStyleOptionTabBarBase struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionTabBarBase constructs a new QStyleOptionTabBarBase object.
func NewQStyleOptionTabBarBase() *QStyleOptionTabBarBase {

	ret := newQStyleOptionTabBarBase(QStyleOptionTabBarBase_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionTabBarBase2 constructs a new QStyleOptionTabBarBase object.
func NewQStyleOptionTabBarBase2(other *QStyleOptionTabBarBase) *QStyleOptionTabBarBase {

	ret := newQStyleOptionTabBarBase(QStyleOptionTabBarBase_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

type QStyleOptionHeader struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionHeader constructs a new QStyleOptionHeader object.
func NewQStyleOptionHeader() *QStyleOptionHeader {

	ret := newQStyleOptionHeader(QStyleOptionHeader_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionHeader2 constructs a new QStyleOptionHeader object.
func NewQStyleOptionHeader2(other *QStyleOptionHeader) *QStyleOptionHeader {

	ret := newQStyleOptionHeader(QStyleOptionHeader_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

type QStyleOptionHeaderV2 struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionHeaderV2 constructs a new QStyleOptionHeaderV2 object.
func NewQStyleOptionHeaderV2() *QStyleOptionHeaderV2 {

	ret := newQStyleOptionHeaderV2(QStyleOptionHeaderV2_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionHeaderV22 constructs a new QStyleOptionHeaderV2 object.
func NewQStyleOptionHeaderV22(other *QStyleOptionHeaderV2) *QStyleOptionHeaderV2 {

	ret := newQStyleOptionHeaderV2(QStyleOptionHeaderV2_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

type QStyleOptionButton struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionButton constructs a new QStyleOptionButton object.
func NewQStyleOptionButton() *QStyleOptionButton {

	ret := newQStyleOptionButton(QStyleOptionButton_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionButton2 constructs a new QStyleOptionButton object.
func NewQStyleOptionButton2(other *QStyleOptionButton) *QStyleOptionButton {

	ret := newQStyleOptionButton(QStyleOptionButton_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

type QStyleOptionTab struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionTab constructs a new QStyleOptionTab object.
func NewQStyleOptionTab() *QStyleOptionTab {

	ret := newQStyleOptionTab(QStyleOptionTab_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionTab2 constructs a new QStyleOptionTab object.
func NewQStyleOptionTab2(other *QStyleOptionTab) *QStyleOptionTab {

	ret := newQStyleOptionTab(QStyleOptionTab_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

type QStyleOptionToolBar struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionToolBar constructs a new QStyleOptionToolBar object.
func NewQStyleOptionToolBar() *QStyleOptionToolBar {

	ret := newQStyleOptionToolBar(QStyleOptionToolBar_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionToolBar2 constructs a new QStyleOptionToolBar object.
func NewQStyleOptionToolBar2(other *QStyleOptionToolBar) *QStyleOptionToolBar {

	ret := newQStyleOptionToolBar(QStyleOptionToolBar_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

type QStyleOptionProgressBar struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionProgressBar constructs a new QStyleOptionProgressBar object.
func NewQStyleOptionProgressBar() *QStyleOptionProgressBar {

	ret := newQStyleOptionProgressBar(QStyleOptionProgressBar_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionProgressBar2 constructs a new QStyleOptionProgressBar object.
func NewQStyleOptionProgressBar2(other *QStyleOptionProgressBar) *QStyleOptionProgressBar {

	ret := newQStyleOptionProgressBar(QStyleOptionProgressBar_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

type QStyleOptionMenuItem struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionMenuItem constructs a new QStyleOptionMenuItem object.
func NewQStyleOptionMenuItem() *QStyleOptionMenuItem {

	ret := newQStyleOptionMenuItem(QStyleOptionMenuItem_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionMenuItem2 constructs a new QStyleOptionMenuItem object.
func NewQStyleOptionMenuItem2(other *QStyleOptionMenuItem) *QStyleOptionMenuItem {

	ret := newQStyleOptionMenuItem(QStyleOptionMenuItem_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

type QStyleOptionDockWidget struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionDockWidget constructs a new QStyleOptionDockWidget object.
func NewQStyleOptionDockWidget() *QStyleOptionDockWidget {

	ret := newQStyleOptionDockWidget(QStyleOptionDockWidget_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionDockWidget2 constructs a new QStyleOptionDockWidget object.
func NewQStyleOptionDockWidget2(other *QStyleOptionDockWidget) *QStyleOptionDockWidget {

	ret := newQStyleOptionDockWidget(QStyleOptionDockWidget_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

type QStyleOptionViewItem struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionViewItem constructs a new QStyleOptionViewItem object.
func NewQStyleOptionViewItem() *QStyleOptionViewItem {

	ret := newQStyleOptionViewItem(QStyleOptionViewItem_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionViewItem2 constructs a new QStyleOptionViewItem object.
func NewQStyleOptionViewItem2(other *QStyleOptionViewItem) *QStyleOptionViewItem {

	ret := newQStyleOptionViewItem(QStyleOptionViewItem_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

type QStyleOptionToolBox struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionToolBox constructs a new QStyleOptionToolBox object.
func NewQStyleOptionToolBox() *QStyleOptionToolBox {

	ret := newQStyleOptionToolBox(QStyleOptionToolBox_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionToolBox2 constructs a new QStyleOptionToolBox object.
func NewQStyleOptionToolBox2(other *QStyleOptionToolBox) *QStyleOptionToolBox {

	ret := newQStyleOptionToolBox(QStyleOptionToolBox_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

type QStyleOptionRubberBand struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionRubberBand constructs a new QStyleOptionRubberBand object.
func NewQStyleOptionRubberBand() *QStyleOptionRubberBand {

	ret := newQStyleOptionRubberBand(QStyleOptionRubberBand_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionRubberBand2 constructs a new QStyleOptionRubberBand object.
func NewQStyleOptionRubberBand2(other *QStyleOptionRubberBand) *QStyleOptionRubberBand {

	ret := newQStyleOptionRubberBand(QStyleOptionRubberBand_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

type QStyleOptionComplex struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionComplex constructs a new QStyleOptionComplex object.
func NewQStyleOptionComplex() *QStyleOptionComplex {

	ret := newQStyleOptionComplex(QStyleOptionComplex_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionComplex2 constructs a new QStyleOptionComplex object.
func NewQStyleOptionComplex2(other *QStyleOptionComplex) *QStyleOptionComplex {

	ret := newQStyleOptionComplex(QStyleOptionComplex_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionComplex3 constructs a new QStyleOptionComplex object.
func NewQStyleOptionComplex3(version int) *QStyleOptionComplex {

	ret := newQStyleOptionComplex(QStyleOptionComplex_new3((int)(version)))
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionComplex4 constructs a new QStyleOptionComplex object.
func NewQStyleOptionComplex4(version int, typeVal int) *QStyleOptionComplex {

	ret := newQStyleOptionComplex(QStyleOptionComplex_new4((int)(version), (int)(typeVal)))
	ret.isSubclass = true
	return ret
}

type QStyleOptionSlider struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionSlider constructs a new QStyleOptionSlider object.
func NewQStyleOptionSlider() *QStyleOptionSlider {

	ret := newQStyleOptionSlider(QStyleOptionSlider_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionSlider2 constructs a new QStyleOptionSlider object.
func NewQStyleOptionSlider2(other *QStyleOptionSlider) *QStyleOptionSlider {

	ret := newQStyleOptionSlider(QStyleOptionSlider_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

type QStyleOptionSpinBox struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionSpinBox constructs a new QStyleOptionSpinBox object.
func NewQStyleOptionSpinBox() *QStyleOptionSpinBox {

	ret := newQStyleOptionSpinBox(QStyleOptionSpinBox_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionSpinBox2 constructs a new QStyleOptionSpinBox object.
func NewQStyleOptionSpinBox2(other *QStyleOptionSpinBox) *QStyleOptionSpinBox {

	ret := newQStyleOptionSpinBox(QStyleOptionSpinBox_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

type QStyleOptionToolButton struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionToolButton constructs a new QStyleOptionToolButton object.
func NewQStyleOptionToolButton() *QStyleOptionToolButton {

	ret := newQStyleOptionToolButton(QStyleOptionToolButton_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionToolButton2 constructs a new QStyleOptionToolButton object.
func NewQStyleOptionToolButton2(other *QStyleOptionToolButton) *QStyleOptionToolButton {

	ret := newQStyleOptionToolButton(QStyleOptionToolButton_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

type QStyleOptionComboBox struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionComboBox constructs a new QStyleOptionComboBox object.
func NewQStyleOptionComboBox() *QStyleOptionComboBox {

	ret := newQStyleOptionComboBox(QStyleOptionComboBox_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionComboBox2 constructs a new QStyleOptionComboBox object.
func NewQStyleOptionComboBox2(other *QStyleOptionComboBox) *QStyleOptionComboBox {

	ret := newQStyleOptionComboBox(QStyleOptionComboBox_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

type QStyleOptionTitleBar struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionTitleBar constructs a new QStyleOptionTitleBar object.
func NewQStyleOptionTitleBar() *QStyleOptionTitleBar {

	ret := newQStyleOptionTitleBar(QStyleOptionTitleBar_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionTitleBar2 constructs a new QStyleOptionTitleBar object.
func NewQStyleOptionTitleBar2(other *QStyleOptionTitleBar) *QStyleOptionTitleBar {

	ret := newQStyleOptionTitleBar(QStyleOptionTitleBar_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

type QStyleOptionGroupBox struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionGroupBox constructs a new QStyleOptionGroupBox object.
func NewQStyleOptionGroupBox() *QStyleOptionGroupBox {

	ret := newQStyleOptionGroupBox(QStyleOptionGroupBox_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionGroupBox2 constructs a new QStyleOptionGroupBox object.
func NewQStyleOptionGroupBox2(other *QStyleOptionGroupBox) *QStyleOptionGroupBox {

	ret := newQStyleOptionGroupBox(QStyleOptionGroupBox_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

type QStyleOptionSizeGrip struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionSizeGrip constructs a new QStyleOptionSizeGrip object.
func NewQStyleOptionSizeGrip() *QStyleOptionSizeGrip {

	ret := newQStyleOptionSizeGrip(QStyleOptionSizeGrip_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionSizeGrip2 constructs a new QStyleOptionSizeGrip object.
func NewQStyleOptionSizeGrip2(other *QStyleOptionSizeGrip) *QStyleOptionSizeGrip {

	ret := newQStyleOptionSizeGrip(QStyleOptionSizeGrip_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

type QStyleOptionGraphicsItem struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleOptionGraphicsItem constructs a new QStyleOptionGraphicsItem object.
func NewQStyleOptionGraphicsItem() *QStyleOptionGraphicsItem {

	ret := newQStyleOptionGraphicsItem(QStyleOptionGraphicsItem_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleOptionGraphicsItem2 constructs a new QStyleOptionGraphicsItem object.
func NewQStyleOptionGraphicsItem2(other *QStyleOptionGraphicsItem) *QStyleOptionGraphicsItem {

	ret := newQStyleOptionGraphicsItem(QStyleOptionGraphicsItem_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func QStyleOptionGraphicsItem_LevelOfDetailFromTransform(worldTransform *QTransform) float64 {
	return (float64)(QStyleOptionGraphicsItem_LevelOfDetailFromTransform(worldTransform.cPointer()))
}

type QStyleHintReturn struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleHintReturn constructs a new QStyleHintReturn object.
func NewQStyleHintReturn() *QStyleHintReturn {

	ret := newQStyleHintReturn(QStyleHintReturn_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleHintReturn2 constructs a new QStyleHintReturn object.
func NewQStyleHintReturn2(param1 *QStyleHintReturn) *QStyleHintReturn {

	ret := newQStyleHintReturn(QStyleHintReturn_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQStyleHintReturn3 constructs a new QStyleHintReturn object.
func NewQStyleHintReturn3(version int) *QStyleHintReturn {

	ret := newQStyleHintReturn(QStyleHintReturn_new3((int)(version)))
	ret.isSubclass = true
	return ret
}

// NewQStyleHintReturn4 constructs a new QStyleHintReturn object.
func NewQStyleHintReturn4(version int, typeVal int) *QStyleHintReturn {

	ret := newQStyleHintReturn(QStyleHintReturn_new4((int)(version), (int)(typeVal)))
	ret.isSubclass = true
	return ret
}

func (this *QStyleHintReturn) OperatorAssign(param1 *QStyleHintReturn) {
	QStyleHintReturn_OperatorAssign(this.h, param1.cPointer())
}

type QStyleHintReturnMask struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleHintReturnMask constructs a new QStyleHintReturnMask object.
func NewQStyleHintReturnMask() *QStyleHintReturnMask {

	ret := newQStyleHintReturnMask(QStyleHintReturnMask_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleHintReturnMask2 constructs a new QStyleHintReturnMask object.
func NewQStyleHintReturnMask2(param1 *QStyleHintReturnMask) *QStyleHintReturnMask {

	ret := newQStyleHintReturnMask(QStyleHintReturnMask_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QStyleHintReturnMask) OperatorAssign(param1 *QStyleHintReturnMask) {
	QStyleHintReturnMask_OperatorAssign(this.h, param1.cPointer())
}

type QStyleHintReturnVariant struct {
	h          uintptr
	isSubclass bool
}

// NewQStyleHintReturnVariant constructs a new QStyleHintReturnVariant object.
func NewQStyleHintReturnVariant() *QStyleHintReturnVariant {

	ret := newQStyleHintReturnVariant(QStyleHintReturnVariant_new())
	ret.isSubclass = true
	return ret
}

// NewQStyleHintReturnVariant2 constructs a new QStyleHintReturnVariant object.
func NewQStyleHintReturnVariant2(param1 *QStyleHintReturnVariant) *QStyleHintReturnVariant {

	ret := newQStyleHintReturnVariant(QStyleHintReturnVariant_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QStyleHintReturnVariant) OperatorAssign(param1 *QStyleHintReturnVariant) {
	QStyleHintReturnVariant_OperatorAssign(this.h, param1.cPointer())
}
