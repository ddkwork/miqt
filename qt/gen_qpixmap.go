package qt

import (
	"unsafe"
)

type QPixmap struct {
	h          uintptr
	isSubclass bool
}

// NewQPixmap constructs a new QPixmap object.
func NewQPixmap() *QPixmap {
	g := newQPixmap(QPixmap_new())
	g.isSubclass = true
	return g
}

// NewQPixmap2 constructs a new QPixmap object.
func NewQPixmap2(w int, h int) *QPixmap {
	g := newQPixmap(QPixmap_new2((int)(w), (int)(h)))
	g.isSubclass = true
	return g
}

// NewQPixmap3 constructs a new QPixmap object.
func NewQPixmap3(param1 *QSize) *QPixmap {
	g := newQPixmap(QPixmap_new3(param1.cPointer()))
	g.isSubclass = true
	return g
}

// NewQPixmap4 constructs a new QPixmap object.
func NewQPixmap4(fileName string) *QPixmap {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	g := newQPixmap(QPixmap_new4(fileName_ms))
	g.isSubclass = true
	return g
}

// NewQPixmap5 constructs a new QPixmap object.
func NewQPixmap5(param1 *QPixmap) *QPixmap {
	g := newQPixmap(QPixmap_new5(param1.cPointer()))
	g.isSubclass = true
	return g
}

// NewQPixmap6 constructs a new QPixmap object.
func NewQPixmap6(fileName string, format string) *QPixmap {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	format_Cstring := CString(format)
	defer free(unsafe.Pointer(format_Cstring))
	g := newQPixmap(QPixmap_new6(fileName_ms, format_Cstring))
	g.isSubclass = true
	return g
}

// NewQPixmap7 constructs a new QPixmap object.
func NewQPixmap7(fileName string, format string, flags ImageConversionFlag) *QPixmap {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	format_Cstring := CString(format)
	defer free(unsafe.Pointer(format_Cstring))
	g := newQPixmap(QPixmap_new7(fileName_ms, format_Cstring, (int)(flags)))
	g.isSubclass = true
	return g
}

func (this *QPixmap) OperatorAssign(param1 *QPixmap) {
	QPixmap_OperatorAssign(this.h, param1.cPointer())
}

func (this *QPixmap) Swap(other *QPixmap) {
	QPixmap_Swap(this.h, other.cPointer())
}

func (this *QPixmap) IsNull() bool {
	return (bool)(QPixmap_IsNull(this.h))
}

func (this *QPixmap) DevType() int {
	return (int)(QPixmap_DevType(this.h))
}

func (this *QPixmap) Width() int {
	return (int)(QPixmap_Width(this.h))
}

func (this *QPixmap) Height() int {
	return (int)(QPixmap_Height(this.h))
}

func (this *QPixmap) Size() *QSize {
	_goptr := newQSize(QPixmap_Size(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) Rect() *QRect {
	_goptr := newQRect(QPixmap_Rect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) Depth() int {
	return (int)(QPixmap_Depth(this.h))
}

func QPixmap_DefaultDepth() int {
	return (int)(QPixmap_DefaultDepth())
}

func (this *QPixmap) Fill() {
	QPixmap_Fill(this.h)
}

func (this *QPixmap) Mask() *QBitmap {
	_goptr := newQBitmap(QPixmap_Mask(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) SetMask(mask *QBitmap) {
	QPixmap_SetMask(this.h, mask.cPointer())
}

func (this *QPixmap) DevicePixelRatio() float64 {
	return (float64)(QPixmap_DevicePixelRatio(this.h))
}

func (this *QPixmap) SetDevicePixelRatio(scaleFactor float64) {
	QPixmap_SetDevicePixelRatio(this.h, (double)(scaleFactor))
}

func (this *QPixmap) DeviceIndependentSize() *QSizeF {
	_goptr := newQSizeF(QPixmap_DeviceIndependentSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) HasAlpha() bool {
	return (bool)(QPixmap_HasAlpha(this.h))
}

func (this *QPixmap) HasAlphaChannel() bool {
	return (bool)(QPixmap_HasAlphaChannel(this.h))
}

func (this *QPixmap) CreateHeuristicMask() *QBitmap {
	_goptr := newQBitmap(QPixmap_CreateHeuristicMask(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) CreateMaskFromColor(maskColor *QColor) *QBitmap {
	_goptr := newQBitmap(QPixmap_CreateMaskFromColor(this.h, maskColor.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) Scaled(w int, h int) *QPixmap {
	_goptr := newQPixmap(QPixmap_Scaled(this.h, (int)(w), (int)(h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) ScaledWithQSize(s *QSize) *QPixmap {
	_goptr := newQPixmap(QPixmap_ScaledWithQSize(this.h, s.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) ScaledToWidth(w int) *QPixmap {
	_goptr := newQPixmap(QPixmap_ScaledToWidth(this.h, (int)(w)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) ScaledToHeight(h int) *QPixmap {
	_goptr := newQPixmap(QPixmap_ScaledToHeight(this.h, (int)(h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) Transformed(param1 *QTransform) *QPixmap {
	_goptr := newQPixmap(QPixmap_Transformed(this.h, param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QPixmap_TrueMatrix(m *QTransform, w int, h int) *QTransform {
	_goptr := newQTransform(QPixmap_TrueMatrix(m.cPointer(), (int)(w), (int)(h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) ToImage() *QImage {
	_goptr := newQImage(QPixmap_ToImage(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QPixmap_FromImage(image *QImage) *QPixmap {
	_goptr := newQPixmap(QPixmap_FromImage(image.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QPixmap_FromImageReader(imageReader *QImageReader) *QPixmap {
	_goptr := newQPixmap(QPixmap_FromImageReader(imageReader.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) Load(fileName string) bool {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	return (bool)(QPixmap_Load(this.h, fileName_ms))
}

func (this *QPixmap) LoadFromData(buf *byte, lenVal uint) bool {
	return (bool)(QPixmap_LoadFromData(this.h, (*uchar)(unsafe.Pointer(buf)), (uint)(lenVal)))
}

func (this *QPixmap) LoadFromDataWithData(data []byte) bool {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	return (bool)(QPixmap_LoadFromDataWithData(this.h, data_alias))
}

func (this *QPixmap) Save(fileName string) bool {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	return (bool)(QPixmap_Save(this.h, fileName_ms))
}

func (this *QPixmap) SaveWithDevice(device *QIODevice) bool {
	return (bool)(QPixmap_SaveWithDevice(this.h, device.cPointer()))
}

func (this *QPixmap) ConvertFromImage(img *QImage) bool {
	return (bool)(QPixmap_ConvertFromImage(this.h, img.cPointer()))
}

func (this *QPixmap) Copy(x int, y int, width int, height int) *QPixmap {
	_goptr := newQPixmap(QPixmap_Copy(this.h, (int)(x), (int)(y), (int)(width), (int)(height)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) Copy2() *QPixmap {
	_goptr := newQPixmap(QPixmap_Copy2(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) Scroll(dx int, dy int, x int, y int, width int, height int) {
	QPixmap_Scroll(this.h, (int)(dx), (int)(dy), (int)(x), (int)(y), (int)(width), (int)(height))
}

func (this *QPixmap) Scroll2(dx int, dy int, rect *QRect) {
	QPixmap_Scroll2(this.h, (int)(dx), (int)(dy), rect.cPointer())
}

func (this *QPixmap) CacheKey() int64 {
	return (int64)(QPixmap_CacheKey(this.h))
}

func (this *QPixmap) IsDetached() bool {
	return (bool)(QPixmap_IsDetached(this.h))
}

func (this *QPixmap) Detach() {
	QPixmap_Detach(this.h)
}

func (this *QPixmap) IsQBitmap() bool {
	return (bool)(QPixmap_IsQBitmap(this.h))
}

func (this *QPixmap) PaintEngine() *QPaintEngine {
	return newQPaintEngine(QPixmap_PaintEngine(this.h))
}

func (this *QPixmap) OperatorNot() bool {
	return (bool)(QPixmap_OperatorNot(this.h))
}

func (this *QPixmap) DataPtr() *DataPtr {
	xxxxxxxxx
}

func (this *QPixmap) Fill1(fillColor *QColor) {
	QPixmap_Fill1(this.h, fillColor.cPointer())
}

func (this *QPixmap) CreateHeuristicMask1(clipTight bool) *QBitmap {
	_goptr := newQBitmap(QPixmap_CreateHeuristicMask1(this.h, (bool)(clipTight)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) CreateMaskFromColor2(maskColor *QColor, mode MaskMode) *QBitmap {
	_goptr := newQBitmap(QPixmap_CreateMaskFromColor2(this.h, maskColor.cPointer(), (int)(mode)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) Scaled3(w int, h int, aspectMode AspectRatioMode) *QPixmap {
	_goptr := newQPixmap(QPixmap_Scaled3(this.h, (int)(w), (int)(h), (int)(aspectMode)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) Scaled4(w int, h int, aspectMode AspectRatioMode, mode TransformationMode) *QPixmap {
	_goptr := newQPixmap(QPixmap_Scaled4(this.h, (int)(w), (int)(h), (int)(aspectMode), (int)(mode)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) Scaled2(s *QSize, aspectMode AspectRatioMode) *QPixmap {
	_goptr := newQPixmap(QPixmap_Scaled2(this.h, s.cPointer(), (int)(aspectMode)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) Scaled32(s *QSize, aspectMode AspectRatioMode, mode TransformationMode) *QPixmap {
	_goptr := newQPixmap(QPixmap_Scaled32(this.h, s.cPointer(), (int)(aspectMode), (int)(mode)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) ScaledToWidth2(w int, mode TransformationMode) *QPixmap {
	_goptr := newQPixmap(QPixmap_ScaledToWidth2(this.h, (int)(w), (int)(mode)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) ScaledToHeight2(h int, mode TransformationMode) *QPixmap {
	_goptr := newQPixmap(QPixmap_ScaledToHeight2(this.h, (int)(h), (int)(mode)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) Transformed2(param1 *QTransform, mode TransformationMode) *QPixmap {
	_goptr := newQPixmap(QPixmap_Transformed2(this.h, param1.cPointer(), (int)(mode)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QPixmap_FromImage2(image *QImage, flags ImageConversionFlag) *QPixmap {
	_goptr := newQPixmap(QPixmap_FromImage2(image.cPointer(), (int)(flags)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QPixmap_FromImageReader2(imageReader *QImageReader, flags ImageConversionFlag) *QPixmap {
	_goptr := newQPixmap(QPixmap_FromImageReader2(imageReader.cPointer(), (int)(flags)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) Load2(fileName string, format string) bool {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	format_Cstring := CString(format)
	defer free(unsafe.Pointer(format_Cstring))
	return (bool)(QPixmap_Load2(this.h, fileName_ms, format_Cstring))
}

func (this *QPixmap) Load3(fileName string, format string, flags ImageConversionFlag) bool {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	format_Cstring := CString(format)
	defer free(unsafe.Pointer(format_Cstring))
	return (bool)(QPixmap_Load3(this.h, fileName_ms, format_Cstring, (int)(flags)))
}

func (this *QPixmap) LoadFromData3(buf *byte, lenVal uint, format string) bool {
	format_Cstring := CString(format)
	defer free(unsafe.Pointer(format_Cstring))
	return (bool)(QPixmap_LoadFromData3(this.h, (*uchar)(unsafe.Pointer(buf)), (uint)(lenVal), format_Cstring))
}

func (this *QPixmap) LoadFromData4(buf *byte, lenVal uint, format string, flags ImageConversionFlag) bool {
	format_Cstring := CString(format)
	defer free(unsafe.Pointer(format_Cstring))
	return (bool)(QPixmap_LoadFromData4(this.h, (*uchar)(unsafe.Pointer(buf)), (uint)(lenVal), format_Cstring, (int)(flags)))
}

func (this *QPixmap) LoadFromData2(data []byte, format string) bool {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	format_Cstring := CString(format)
	defer free(unsafe.Pointer(format_Cstring))
	return (bool)(QPixmap_LoadFromData2(this.h, data_alias, format_Cstring))
}

func (this *QPixmap) LoadFromData32(data []byte, format string, flags ImageConversionFlag) bool {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	format_Cstring := CString(format)
	defer free(unsafe.Pointer(format_Cstring))
	return (bool)(QPixmap_LoadFromData32(this.h, data_alias, format_Cstring, (int)(flags)))
}

func (this *QPixmap) Save2(fileName string, format string) bool {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	format_Cstring := CString(format)
	defer free(unsafe.Pointer(format_Cstring))
	return (bool)(QPixmap_Save2(this.h, fileName_ms, format_Cstring))
}

func (this *QPixmap) Save3(fileName string, format string, quality int) bool {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	format_Cstring := CString(format)
	defer free(unsafe.Pointer(format_Cstring))
	return (bool)(QPixmap_Save3(this.h, fileName_ms, format_Cstring, (int)(quality)))
}

func (this *QPixmap) Save22(device *QIODevice, format string) bool {
	format_Cstring := CString(format)
	defer free(unsafe.Pointer(format_Cstring))
	return (bool)(QPixmap_Save22(this.h, device.cPointer(), format_Cstring))
}

func (this *QPixmap) Save32(device *QIODevice, format string, quality int) bool {
	format_Cstring := CString(format)
	defer free(unsafe.Pointer(format_Cstring))
	return (bool)(QPixmap_Save32(this.h, device.cPointer(), format_Cstring, (int)(quality)))
}

func (this *QPixmap) ConvertFromImage2(img *QImage, flags ImageConversionFlag) bool {
	return (bool)(QPixmap_ConvertFromImage2(this.h, img.cPointer(), (int)(flags)))
}

func (this *QPixmap) Copy1(rect *QRect) *QPixmap {
	_goptr := newQPixmap(QPixmap_Copy1(this.h, rect.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPixmap) Scroll7(dx int, dy int, x int, y int, width int, height int, exposed *QRegion) {
	QPixmap_Scroll7(this.h, (int)(dx), (int)(dy), (int)(x), (int)(y), (int)(width), (int)(height), exposed.cPointer())
}

func (this *QPixmap) Scroll4(dx int, dy int, rect *QRect, exposed *QRegion) {
	QPixmap_Scroll4(this.h, (int)(dx), (int)(dy), rect.cPointer(), exposed.cPointer())
}
