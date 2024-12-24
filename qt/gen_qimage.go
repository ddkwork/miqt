package qt
import (
	"unsafe"
)

		type QImage__InvertMode int
		const (
QImage__InvertRgb QImage__InvertMode = 0
QImage__InvertRgba QImage__InvertMode = 1

)


		type QImage__Format int
		const (
QImage__Format_Invalid QImage__Format = 0
QImage__Format_Mono QImage__Format = 1
QImage__Format_MonoLSB QImage__Format = 2
QImage__Format_Indexed8 QImage__Format = 3
QImage__Format_RGB32 QImage__Format = 4
QImage__Format_ARGB32 QImage__Format = 5
QImage__Format_ARGB32_Premultiplied QImage__Format = 6
QImage__Format_RGB16 QImage__Format = 7
QImage__Format_ARGB8565_Premultiplied QImage__Format = 8
QImage__Format_RGB666 QImage__Format = 9
QImage__Format_ARGB6666_Premultiplied QImage__Format = 10
QImage__Format_RGB555 QImage__Format = 11
QImage__Format_ARGB8555_Premultiplied QImage__Format = 12
QImage__Format_RGB888 QImage__Format = 13
QImage__Format_RGB444 QImage__Format = 14
QImage__Format_ARGB4444_Premultiplied QImage__Format = 15
QImage__Format_RGBX8888 QImage__Format = 16
QImage__Format_RGBA8888 QImage__Format = 17
QImage__Format_RGBA8888_Premultiplied QImage__Format = 18
QImage__Format_BGR30 QImage__Format = 19
QImage__Format_A2BGR30_Premultiplied QImage__Format = 20
QImage__Format_RGB30 QImage__Format = 21
QImage__Format_A2RGB30_Premultiplied QImage__Format = 22
QImage__Format_Alpha8 QImage__Format = 23
QImage__Format_Grayscale8 QImage__Format = 24
QImage__Format_RGBX64 QImage__Format = 25
QImage__Format_RGBA64 QImage__Format = 26
QImage__Format_RGBA64_Premultiplied QImage__Format = 27
QImage__Format_Grayscale16 QImage__Format = 28
QImage__Format_BGR888 QImage__Format = 29
QImage__Format_RGBX16FPx4 QImage__Format = 30
QImage__Format_RGBA16FPx4 QImage__Format = 31
QImage__Format_RGBA16FPx4_Premultiplied QImage__Format = 32
QImage__Format_RGBX32FPx4 QImage__Format = 33
QImage__Format_RGBA32FPx4 QImage__Format = 34
QImage__Format_RGBA32FPx4_Premultiplied QImage__Format = 35
QImage__Format_CMYK8888 QImage__Format = 36
QImage__NImageFormats QImage__Format = 37

)


		type QImage struct {
			h uintptr
			isSubclass bool
}
		
			// NewQImage constructs a new QImage object.
			func NewQImage() *QImage {
								
				ret := newQImage(QImage_new())
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQImage2 constructs a new QImage object.
			func NewQImage2(size *QSize, format Format) *QImage {
								
				ret := newQImage(QImage_new2(size.cPointer(), format))
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQImage3 constructs a new QImage object.
			func NewQImage3(width int, height int, format Format) *QImage {
								
				ret := newQImage(QImage_new3((int)(width), (int)(height), format))
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQImage4 constructs a new QImage object.
			func NewQImage4(data *byte, width int, height int, format Format) *QImage {
								
				ret := newQImage(QImage_new4((*uchar)(unsafe.Pointer(data)), (int)(width), (int)(height), format))
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQImage5 constructs a new QImage object.
			func NewQImage5(data *byte, width int, height int, format Format) *QImage {
								
				ret := newQImage(QImage_new5((*uchar)(unsafe.Pointer(data)), (int)(width), (int)(height), format))
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQImage6 constructs a new QImage object.
			func NewQImage6(data *byte, width int, height int, bytesPerLine int64, format Format) *QImage {
								
				ret := newQImage(QImage_new6((*uchar)(unsafe.Pointer(data)), (int)(width), (int)(height), (ptrdiff_t)(bytesPerLine), format))
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQImage7 constructs a new QImage object.
			func NewQImage7(data *byte, width int, height int, bytesPerLine int64, format Format) *QImage {
								
				ret := newQImage(QImage_new7((*uchar)(unsafe.Pointer(data)), (int)(width), (int)(height), (ptrdiff_t)(bytesPerLine), format))
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQImage8 constructs a new QImage object.
			func NewQImage8(fileName string) *QImage {
				fileName_ms := struct_miqt_string{}
fileName_ms.data = CString(fileName)
fileName_ms.len = size_t(len(fileName))
defer free(unsafe.Pointer(fileName_ms.data))
				
				ret := newQImage(QImage_new8(fileName_ms))
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQImage9 constructs a new QImage object.
			func NewQImage9(param1 *QImage) *QImage {
								
				ret := newQImage(QImage_new9(param1.cPointer()))
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQImage10 constructs a new QImage object.
			func NewQImage10(fileName string, format string) *QImage {
				fileName_ms := struct_miqt_string{}
fileName_ms.data = CString(fileName)
fileName_ms.len = size_t(len(fileName))
defer free(unsafe.Pointer(fileName_ms.data))
format_Cstring := CString(format)
defer free(unsafe.Pointer(format_Cstring))
				
				ret := newQImage(QImage_new10(fileName_ms, format_Cstring))
				ret.isSubclass = true
				return ret
			}
			
			
			func (this *QImage) OperatorAssign(param1 *QImage)  {
				 QImage_OperatorAssign(this.h, param1.cPointer())
}
			
			func (this *QImage) Swap(other *QImage)  {
				 QImage_Swap(this.h, other.cPointer())
}
			
			func (this *QImage) IsNull() bool {
				return (bool)(QImage_IsNull(this.h))
}
			
			func (this *QImage) DevType() int {
				return (int)(QImage_DevType(this.h))
}
			
			func (this *QImage) OperatorEqual(param1 *QImage) bool {
				return (bool)(QImage_OperatorEqual(this.h, param1.cPointer()))
}
			
			func (this *QImage) OperatorNotEqual(param1 *QImage) bool {
				return (bool)(QImage_OperatorNotEqual(this.h, param1.cPointer()))
}
			
			func (this *QImage) Detach()  {
				 QImage_Detach(this.h)
}
			
			func (this *QImage) IsDetached() bool {
				return (bool)(QImage_IsDetached(this.h))
}
			
			func (this *QImage) Copy() *QImage {
				_goptr := newQImage(QImage_Copy(this.h))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) Copy2(x int, y int, w int, h int) *QImage {
				_goptr := newQImage(QImage_Copy2(this.h, (int)(x), (int)(y), (int)(w), (int)(h)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) Format() Format {
				xxxxxxxxx}
			
			func (this *QImage) ConvertToFormat(f Format) *QImage {
				_goptr := newQImage(QImage_ConvertToFormat(this.h, f))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) ConvertToFormat2(f Format, colorTable []uint) *QImage {
				colorTable_CArray := (*[0xffff]uint)(malloc(size_t(8 * len(colorTable))))
defer free(unsafe.Pointer(colorTable_CArray))
for i := range colorTable{
colorTable_CArray[i] = (uint)(colorTable[i])
}
colorTable_ma := struct_miqt_array{len: size_t(len(colorTable)), data: unsafe.Pointer(colorTable_CArray)}
_goptr := newQImage(QImage_ConvertToFormat2(this.h, f, colorTable_ma))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) ReinterpretAsFormat(f Format) bool {
				return (bool)(QImage_ReinterpretAsFormat(this.h, f))
}
			
			func (this *QImage) ConvertedTo(f Format) *QImage {
				_goptr := newQImage(QImage_ConvertedTo(this.h, f))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) ConvertTo(f Format)  {
				 QImage_ConvertTo(this.h, f)
}
			
			func (this *QImage) Width() int {
				return (int)(QImage_Width(this.h))
}
			
			func (this *QImage) Height() int {
				return (int)(QImage_Height(this.h))
}
			
			func (this *QImage) Size() *QSize {
				_goptr := newQSize(QImage_Size(this.h))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) Rect() *QRect {
				_goptr := newQRect(QImage_Rect(this.h))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) Depth() int {
				return (int)(QImage_Depth(this.h))
}
			
			func (this *QImage) ColorCount() int {
				return (int)(QImage_ColorCount(this.h))
}
			
			func (this *QImage) BitPlaneCount() int {
				return (int)(QImage_BitPlaneCount(this.h))
}
			
			func (this *QImage) Color(i int) uint {
				return (uint)(QImage_Color(this.h, (int)(i)))
}
			
			func (this *QImage) SetColor(i int, c uint)  {
				 QImage_SetColor(this.h, (int)(i), (uint)(c))
}
			
			func (this *QImage) SetColorCount(colorCount int)  {
				 QImage_SetColorCount(this.h, (int)(colorCount))
}
			
			func (this *QImage) AllGray() bool {
				return (bool)(QImage_AllGray(this.h))
}
			
			func (this *QImage) IsGrayscale() bool {
				return (bool)(QImage_IsGrayscale(this.h))
}
			
			func (this *QImage) Bits() *byte {
				return (*byte)(unsafe.Pointer(QImage_Bits(this.h)))
}
			
			func (this *QImage) Bits2() *byte {
				return (*byte)(unsafe.Pointer(QImage_Bits2(this.h)))
}
			
			func (this *QImage) ConstBits() *byte {
				return (*byte)(unsafe.Pointer(QImage_ConstBits(this.h)))
}
			
			func (this *QImage) SizeInBytes() int64 {
				return (int64)(QImage_SizeInBytes(this.h))
}
			
			func (this *QImage) ScanLine(param1 int) *byte {
				return (*byte)(unsafe.Pointer(QImage_ScanLine(this.h, (int)(param1))))
}
			
			func (this *QImage) ScanLineWithInt(param1 int) *byte {
				return (*byte)(unsafe.Pointer(QImage_ScanLineWithInt(this.h, (int)(param1))))
}
			
			func (this *QImage) ConstScanLine(param1 int) *byte {
				return (*byte)(unsafe.Pointer(QImage_ConstScanLine(this.h, (int)(param1))))
}
			
			func (this *QImage) BytesPerLine() int64 {
				return (int64)(QImage_BytesPerLine(this.h))
}
			
			func (this *QImage) Valid(x int, y int) bool {
				return (bool)(QImage_Valid(this.h, (int)(x), (int)(y)))
}
			
			func (this *QImage) ValidWithPt(pt *QPoint) bool {
				return (bool)(QImage_ValidWithPt(this.h, pt.cPointer()))
}
			
			func (this *QImage) PixelIndex(x int, y int) int {
				return (int)(QImage_PixelIndex(this.h, (int)(x), (int)(y)))
}
			
			func (this *QImage) PixelIndexWithPt(pt *QPoint) int {
				return (int)(QImage_PixelIndexWithPt(this.h, pt.cPointer()))
}
			
			func (this *QImage) Pixel(x int, y int) uint {
				return (uint)(QImage_Pixel(this.h, (int)(x), (int)(y)))
}
			
			func (this *QImage) PixelWithPt(pt *QPoint) uint {
				return (uint)(QImage_PixelWithPt(this.h, pt.cPointer()))
}
			
			func (this *QImage) SetPixel(x int, y int, index_or_rgb uint)  {
				 QImage_SetPixel(this.h, (int)(x), (int)(y), (uint)(index_or_rgb))
}
			
			func (this *QImage) SetPixel2(pt *QPoint, index_or_rgb uint)  {
				 QImage_SetPixel2(this.h, pt.cPointer(), (uint)(index_or_rgb))
}
			
			func (this *QImage) PixelColor(x int, y int) *QColor {
				_goptr := newQColor(QImage_PixelColor(this.h, (int)(x), (int)(y)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) PixelColorWithPt(pt *QPoint) *QColor {
				_goptr := newQColor(QImage_PixelColorWithPt(this.h, pt.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) SetPixelColor(x int, y int, c *QColor)  {
				 QImage_SetPixelColor(this.h, (int)(x), (int)(y), c.cPointer())
}
			
			func (this *QImage) SetPixelColor2(pt *QPoint, c *QColor)  {
				 QImage_SetPixelColor2(this.h, pt.cPointer(), c.cPointer())
}
			
			func (this *QImage) ColorTable() []uint {
				var _ma struct_miqt_array =  QImage_ColorTable(this.h)
_ret := make([]uint, int(_ma.len))
_outCast := (*[0xffff]uint)(unsafe.Pointer(_ma.data)) // hey ya
for i := 0; i < int(_ma.len); i++ {
_ret[i] = (uint)(_outCast[i])
}
return  _ret
}
			
			func (this *QImage) SetColorTable(colors []uint)  {
				colors_CArray := (*[0xffff]uint)(malloc(size_t(8 * len(colors))))
defer free(unsafe.Pointer(colors_CArray))
for i := range colors{
colors_CArray[i] = (uint)(colors[i])
}
colors_ma := struct_miqt_array{len: size_t(len(colors)), data: unsafe.Pointer(colors_CArray)}
 QImage_SetColorTable(this.h, colors_ma)
}
			
			func (this *QImage) DevicePixelRatio() float64 {
				return (float64)(QImage_DevicePixelRatio(this.h))
}
			
			func (this *QImage) SetDevicePixelRatio(scaleFactor float64)  {
				 QImage_SetDevicePixelRatio(this.h, (double)(scaleFactor))
}
			
			func (this *QImage) DeviceIndependentSize() *QSizeF {
				_goptr := newQSizeF(QImage_DeviceIndependentSize(this.h))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) Fill(pixel uint)  {
				 QImage_Fill(this.h, (uint)(pixel))
}
			
			func (this *QImage) FillWithColor(color *QColor)  {
				 QImage_FillWithColor(this.h, color.cPointer())
}
			
			func (this *QImage) Fill2(color GlobalColor)  {
				 QImage_Fill2(this.h, (int)(color))
}
			
			func (this *QImage) HasAlphaChannel() bool {
				return (bool)(QImage_HasAlphaChannel(this.h))
}
			
			func (this *QImage) SetAlphaChannel(alphaChannel *QImage)  {
				 QImage_SetAlphaChannel(this.h, alphaChannel.cPointer())
}
			
			func (this *QImage) CreateAlphaMask() *QImage {
				_goptr := newQImage(QImage_CreateAlphaMask(this.h))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) CreateHeuristicMask() *QImage {
				_goptr := newQImage(QImage_CreateHeuristicMask(this.h))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) CreateMaskFromColor(color uint) *QImage {
				_goptr := newQImage(QImage_CreateMaskFromColor(this.h, (uint)(color)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) Scaled(w int, h int) *QImage {
				_goptr := newQImage(QImage_Scaled(this.h, (int)(w), (int)(h)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) ScaledWithQSize(s *QSize) *QImage {
				_goptr := newQImage(QImage_ScaledWithQSize(this.h, s.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) ScaledToWidth(w int) *QImage {
				_goptr := newQImage(QImage_ScaledToWidth(this.h, (int)(w)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) ScaledToHeight(h int) *QImage {
				_goptr := newQImage(QImage_ScaledToHeight(this.h, (int)(h)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) Transformed(matrix *QTransform) *QImage {
				_goptr := newQImage(QImage_Transformed(this.h, matrix.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func QImage_TrueMatrix(param1 *QTransform, w int, h int) *QTransform {
				_goptr := newQTransform(QImage_TrueMatrix(param1.cPointer(), (int)(w), (int)(h)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) Mirrored() *QImage {
				_goptr := newQImage(QImage_Mirrored(this.h))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) Mirror()  {
				 QImage_Mirror(this.h)
}
			
			func (this *QImage) RgbSwapped() *QImage {
				_goptr := newQImage(QImage_RgbSwapped(this.h))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) Flipped() *QImage {
				_goptr := newQImage(QImage_Flipped(this.h))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) Flip()  {
				 QImage_Flip(this.h)
}
			
			func (this *QImage) RgbSwap()  {
				 QImage_RgbSwap(this.h)
}
			
			func (this *QImage) InvertPixels()  {
				 QImage_InvertPixels(this.h)
}
			
			func (this *QImage) ColorSpace() *QColorSpace {
				_goptr := newQColorSpace(QImage_ColorSpace(this.h))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) ConvertedToColorSpace(colorSpace *QColorSpace) *QImage {
				_goptr := newQImage(QImage_ConvertedToColorSpace(this.h, colorSpace.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) ConvertedToColorSpace2(colorSpace *QColorSpace, format QImage__Format) *QImage {
				_goptr := newQImage(QImage_ConvertedToColorSpace2(this.h, colorSpace.cPointer(), (int)(format)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) ConvertToColorSpace(colorSpace *QColorSpace)  {
				 QImage_ConvertToColorSpace(this.h, colorSpace.cPointer())
}
			
			func (this *QImage) ConvertToColorSpace2(colorSpace *QColorSpace, format QImage__Format)  {
				 QImage_ConvertToColorSpace2(this.h, colorSpace.cPointer(), (int)(format))
}
			
			func (this *QImage) SetColorSpace(colorSpace *QColorSpace)  {
				 QImage_SetColorSpace(this.h, colorSpace.cPointer())
}
			
			func (this *QImage) ColorTransformed(transform *QColorTransform) *QImage {
				_goptr := newQImage(QImage_ColorTransformed(this.h, transform.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) ColorTransformed2(transform *QColorTransform, format QImage__Format) *QImage {
				_goptr := newQImage(QImage_ColorTransformed2(this.h, transform.cPointer(), (int)(format)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) ApplyColorTransform(transform *QColorTransform)  {
				 QImage_ApplyColorTransform(this.h, transform.cPointer())
}
			
			func (this *QImage) ApplyColorTransform2(transform *QColorTransform, format QImage__Format)  {
				 QImage_ApplyColorTransform2(this.h, transform.cPointer(), (int)(format))
}
			
			func (this *QImage) Load(device *QIODevice, format string) bool {
				format_Cstring := CString(format)
defer free(unsafe.Pointer(format_Cstring))
return (bool)(QImage_Load(this.h, device.cPointer(), format_Cstring))
}
			
			func (this *QImage) LoadWithFileName(fileName string) bool {
				fileName_ms := struct_miqt_string{}
fileName_ms.data = CString(fileName)
fileName_ms.len = size_t(len(fileName))
defer free(unsafe.Pointer(fileName_ms.data))
return (bool)(QImage_LoadWithFileName(this.h, fileName_ms))
}
			
			func (this *QImage) LoadFromData(data QByteArrayView) bool {
				return (bool)(QImage_LoadFromData(this.h, data.cPointer()))
}
			
			func (this *QImage) LoadFromData2(buf *byte, lenVal int) bool {
				return (bool)(QImage_LoadFromData2(this.h, (*uchar)(unsafe.Pointer(buf)), (int)(lenVal)))
}
			
			func (this *QImage) LoadFromDataWithData(data []byte) bool {
				data_alias := struct_miqt_string{}
data_alias.data = (char)(unsafe.Pointer(&data[0]))
data_alias.len = size_t(len(data))
return (bool)(QImage_LoadFromDataWithData(this.h, data_alias))
}
			
			func (this *QImage) Save(fileName string) bool {
				fileName_ms := struct_miqt_string{}
fileName_ms.data = CString(fileName)
fileName_ms.len = size_t(len(fileName))
defer free(unsafe.Pointer(fileName_ms.data))
return (bool)(QImage_Save(this.h, fileName_ms))
}
			
			func (this *QImage) SaveWithDevice(device *QIODevice) bool {
				return (bool)(QImage_SaveWithDevice(this.h, device.cPointer()))
}
			
			func QImage_FromData(data QByteArrayView) *QImage {
				_goptr := newQImage(QImage_FromData(data.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func QImage_FromData2(data *byte, size int) *QImage {
				_goptr := newQImage(QImage_FromData2((*uchar)(unsafe.Pointer(data)), (int)(size)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func QImage_FromDataWithData(data []byte) *QImage {
				data_alias := struct_miqt_string{}
data_alias.data = (char)(unsafe.Pointer(&data[0]))
data_alias.len = size_t(len(data))
_goptr := newQImage(QImage_FromDataWithData(data_alias))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) CacheKey() int64 {
				return (int64)(QImage_CacheKey(this.h))
}
			
			func (this *QImage) PaintEngine() *QPaintEngine {
				return newQPaintEngine(QImage_PaintEngine(this.h))
}
			
			func (this *QImage) DotsPerMeterX() int {
				return (int)(QImage_DotsPerMeterX(this.h))
}
			
			func (this *QImage) DotsPerMeterY() int {
				return (int)(QImage_DotsPerMeterY(this.h))
}
			
			func (this *QImage) SetDotsPerMeterX(dotsPerMeterX int)  {
				 QImage_SetDotsPerMeterX(this.h, (int)(dotsPerMeterX))
}
			
			func (this *QImage) SetDotsPerMeterY(dotsPerMeterY int)  {
				 QImage_SetDotsPerMeterY(this.h, (int)(dotsPerMeterY))
}
			
			func (this *QImage) Offset() *QPoint {
				_goptr := newQPoint(QImage_Offset(this.h))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) SetOffset(offset *QPoint)  {
				 QImage_SetOffset(this.h, offset.cPointer())
}
			
			func (this *QImage) TextKeys() []string {
				var _ma struct_miqt_array =  QImage_TextKeys(this.h)
_ret := make([]string, int(_ma.len))
_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
for i := 0; i < int(_ma.len); i++ {
var _lv_ms struct_miqt_string =  _outCast[i]
_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
free(unsafe.Pointer(_lv_ms.data))
_ret[i] = _lv_ret}
return  _ret
}
			
			func (this *QImage) Text() string {
				var _ms struct_miqt_string =  QImage_Text(this.h)
_ret := GoStringN(_ms.data, int(int64(_ms.len)))
free(unsafe.Pointer(_ms.data))
return _ret}
			
			func (this *QImage) SetText(key string, value string)  {
				key_ms := struct_miqt_string{}
key_ms.data = CString(key)
key_ms.len = size_t(len(key))
defer free(unsafe.Pointer(key_ms.data))
value_ms := struct_miqt_string{}
value_ms.data = CString(value)
value_ms.len = size_t(len(value))
defer free(unsafe.Pointer(value_ms.data))
 QImage_SetText(this.h, key_ms, value_ms)
}
			
			func (this *QImage) PixelFormat() *QPixelFormat {
				_goptr := newQPixelFormat(QImage_PixelFormat(this.h))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func QImage_ToPixelFormat(format QImage__Format) *QPixelFormat {
				_goptr := newQPixelFormat(QImage_ToPixelFormat((int)(format)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func QImage_ToImageFormat(format QPixelFormat) QImage__Format {
				return (QImage__Format)(QImage_ToImageFormat(format.cPointer()))
}
			
			func (this *QImage) ToHBITMAP() *struct HBITMAP__ {
				return (*struct HBITMAP__)(unsafe.Pointer(QImage_ToHBITMAP(this.h)))
}
			
			func (this *QImage) ToHICON() *struct HICON__ {
				return (*struct HICON__)(unsafe.Pointer(QImage_ToHICON(this.h)))
}
			
			func QImage_FromHBITMAP(hbitmap *struct HBITMAP__) *QImage {
				_goptr := newQImage(QImage_FromHBITMAP(hbitmap))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func QImage_FromHICON(icon *struct HICON__) *QImage {
				_goptr := newQImage(QImage_FromHICON(icon))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) DataPtr() *DataPtr {
				xxxxxxxxx}
			
			func (this *QImage) Copy1(rect *QRect) *QImage {
				_goptr := newQImage(QImage_Copy1(this.h, rect.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) ConvertToFormat22(f Format, flags ImageConversionFlag) *QImage {
				_goptr := newQImage(QImage_ConvertToFormat22(this.h, f, (int)(flags)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) ConvertToFormat3(f Format, colorTable []uint, flags ImageConversionFlag) *QImage {
				colorTable_CArray := (*[0xffff]uint)(malloc(size_t(8 * len(colorTable))))
defer free(unsafe.Pointer(colorTable_CArray))
for i := range colorTable{
colorTable_CArray[i] = (uint)(colorTable[i])
}
colorTable_ma := struct_miqt_array{len: size_t(len(colorTable)), data: unsafe.Pointer(colorTable_CArray)}
_goptr := newQImage(QImage_ConvertToFormat3(this.h, f, colorTable_ma, (int)(flags)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) ConvertedTo2(f Format, flags ImageConversionFlag) *QImage {
				_goptr := newQImage(QImage_ConvertedTo2(this.h, f, (int)(flags)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) ConvertTo2(f Format, flags ImageConversionFlag)  {
				 QImage_ConvertTo2(this.h, f, (int)(flags))
}
			
			func (this *QImage) CreateAlphaMask1(flags ImageConversionFlag) *QImage {
				_goptr := newQImage(QImage_CreateAlphaMask1(this.h, (int)(flags)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) CreateHeuristicMask1(clipTight bool) *QImage {
				_goptr := newQImage(QImage_CreateHeuristicMask1(this.h, (bool)(clipTight)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) CreateMaskFromColor2(color uint, mode MaskMode) *QImage {
				_goptr := newQImage(QImage_CreateMaskFromColor2(this.h, (uint)(color), (int)(mode)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) Scaled3(w int, h int, aspectMode AspectRatioMode) *QImage {
				_goptr := newQImage(QImage_Scaled3(this.h, (int)(w), (int)(h), (int)(aspectMode)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) Scaled4(w int, h int, aspectMode AspectRatioMode, mode TransformationMode) *QImage {
				_goptr := newQImage(QImage_Scaled4(this.h, (int)(w), (int)(h), (int)(aspectMode), (int)(mode)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) Scaled2(s *QSize, aspectMode AspectRatioMode) *QImage {
				_goptr := newQImage(QImage_Scaled2(this.h, s.cPointer(), (int)(aspectMode)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) Scaled32(s *QSize, aspectMode AspectRatioMode, mode TransformationMode) *QImage {
				_goptr := newQImage(QImage_Scaled32(this.h, s.cPointer(), (int)(aspectMode), (int)(mode)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) ScaledToWidth2(w int, mode TransformationMode) *QImage {
				_goptr := newQImage(QImage_ScaledToWidth2(this.h, (int)(w), (int)(mode)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) ScaledToHeight2(h int, mode TransformationMode) *QImage {
				_goptr := newQImage(QImage_ScaledToHeight2(this.h, (int)(h), (int)(mode)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) Transformed2(matrix *QTransform, mode TransformationMode) *QImage {
				_goptr := newQImage(QImage_Transformed2(this.h, matrix.cPointer(), (int)(mode)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) Mirrored1(horizontally bool) *QImage {
				_goptr := newQImage(QImage_Mirrored1(this.h, (bool)(horizontally)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) Mirrored2(horizontally bool, vertically bool) *QImage {
				_goptr := newQImage(QImage_Mirrored2(this.h, (bool)(horizontally), (bool)(vertically)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) Mirror1(horizontally bool)  {
				 QImage_Mirror1(this.h, (bool)(horizontally))
}
			
			func (this *QImage) Mirror2(horizontally bool, vertically bool)  {
				 QImage_Mirror2(this.h, (bool)(horizontally), (bool)(vertically))
}
			
			func (this *QImage) Flipped1(orient Orientation) *QImage {
				_goptr := newQImage(QImage_Flipped1(this.h, (int)(orient)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) Flip1(orient Orientation)  {
				 QImage_Flip1(this.h, (int)(orient))
}
			
			func (this *QImage) InvertPixels1(param1 InvertMode)  {
				 QImage_InvertPixels1(this.h, param1)
}
			
			func (this *QImage) ConvertedToColorSpace3(colorSpace *QColorSpace, format QImage__Format, flags ImageConversionFlag) *QImage {
				_goptr := newQImage(QImage_ConvertedToColorSpace3(this.h, colorSpace.cPointer(), (int)(format), (int)(flags)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) ConvertToColorSpace3(colorSpace *QColorSpace, format QImage__Format, flags ImageConversionFlag)  {
				 QImage_ConvertToColorSpace3(this.h, colorSpace.cPointer(), (int)(format), (int)(flags))
}
			
			func (this *QImage) ColorTransformed3(transform *QColorTransform, format QImage__Format, flags ImageConversionFlag) *QImage {
				_goptr := newQImage(QImage_ColorTransformed3(this.h, transform.cPointer(), (int)(format), (int)(flags)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) ApplyColorTransform3(transform *QColorTransform, format QImage__Format, flags ImageConversionFlag)  {
				 QImage_ApplyColorTransform3(this.h, transform.cPointer(), (int)(format), (int)(flags))
}
			
			func (this *QImage) Load2(fileName string, format string) bool {
				fileName_ms := struct_miqt_string{}
fileName_ms.data = CString(fileName)
fileName_ms.len = size_t(len(fileName))
defer free(unsafe.Pointer(fileName_ms.data))
format_Cstring := CString(format)
defer free(unsafe.Pointer(format_Cstring))
return (bool)(QImage_Load2(this.h, fileName_ms, format_Cstring))
}
			
			func (this *QImage) LoadFromData22(data QByteArrayView, format string) bool {
				format_Cstring := CString(format)
defer free(unsafe.Pointer(format_Cstring))
return (bool)(QImage_LoadFromData22(this.h, data.cPointer(), format_Cstring))
}
			
			func (this *QImage) LoadFromData3(buf *byte, lenVal int, format string) bool {
				format_Cstring := CString(format)
defer free(unsafe.Pointer(format_Cstring))
return (bool)(QImage_LoadFromData3(this.h, (*uchar)(unsafe.Pointer(buf)), (int)(lenVal), format_Cstring))
}
			
			func (this *QImage) LoadFromData23(data []byte, format string) bool {
				data_alias := struct_miqt_string{}
data_alias.data = (char)(unsafe.Pointer(&data[0]))
data_alias.len = size_t(len(data))
format_Cstring := CString(format)
defer free(unsafe.Pointer(format_Cstring))
return (bool)(QImage_LoadFromData23(this.h, data_alias, format_Cstring))
}
			
			func (this *QImage) Save2(fileName string, format string) bool {
				fileName_ms := struct_miqt_string{}
fileName_ms.data = CString(fileName)
fileName_ms.len = size_t(len(fileName))
defer free(unsafe.Pointer(fileName_ms.data))
format_Cstring := CString(format)
defer free(unsafe.Pointer(format_Cstring))
return (bool)(QImage_Save2(this.h, fileName_ms, format_Cstring))
}
			
			func (this *QImage) Save3(fileName string, format string, quality int) bool {
				fileName_ms := struct_miqt_string{}
fileName_ms.data = CString(fileName)
fileName_ms.len = size_t(len(fileName))
defer free(unsafe.Pointer(fileName_ms.data))
format_Cstring := CString(format)
defer free(unsafe.Pointer(format_Cstring))
return (bool)(QImage_Save3(this.h, fileName_ms, format_Cstring, (int)(quality)))
}
			
			func (this *QImage) Save22(device *QIODevice, format string) bool {
				format_Cstring := CString(format)
defer free(unsafe.Pointer(format_Cstring))
return (bool)(QImage_Save22(this.h, device.cPointer(), format_Cstring))
}
			
			func (this *QImage) Save32(device *QIODevice, format string, quality int) bool {
				format_Cstring := CString(format)
defer free(unsafe.Pointer(format_Cstring))
return (bool)(QImage_Save32(this.h, device.cPointer(), format_Cstring, (int)(quality)))
}
			
			func QImage_FromData22(data QByteArrayView, format string) *QImage {
				format_Cstring := CString(format)
defer free(unsafe.Pointer(format_Cstring))
_goptr := newQImage(QImage_FromData22(data.cPointer(), format_Cstring))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func QImage_FromData3(data *byte, size int, format string) *QImage {
				format_Cstring := CString(format)
defer free(unsafe.Pointer(format_Cstring))
_goptr := newQImage(QImage_FromData3((*uchar)(unsafe.Pointer(data)), (int)(size), format_Cstring))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func QImage_FromData23(data []byte, format string) *QImage {
				data_alias := struct_miqt_string{}
data_alias.data = (char)(unsafe.Pointer(&data[0]))
data_alias.len = size_t(len(data))
format_Cstring := CString(format)
defer free(unsafe.Pointer(format_Cstring))
_goptr := newQImage(QImage_FromData23(data_alias, format_Cstring))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QImage) Text1(key string) string {
				key_ms := struct_miqt_string{}
key_ms.data = CString(key)
key_ms.len = size_t(len(key))
defer free(unsafe.Pointer(key_ms.data))
var _ms struct_miqt_string =  QImage_Text1(this.h, key_ms)
_ret := GoStringN(_ms.data, int(int64(_ms.len)))
free(unsafe.Pointer(_ms.data))
return _ret}
			
			func (this *QImage) ToHICON1(mask *QImage) *struct HICON__ {
				return (*struct HICON__)(unsafe.Pointer(QImage_ToHICON1(this.h, mask.cPointer())))
}
			
				func (this *QImage) callVirtualBase_DevType() int {
					
					return (int)(QImage_virtualbase_DevType(unsafe.Pointer(this.h)))

				}
			func (this *QImage) OnDevType(slot func(super func() int) int) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QImage_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QImage_DevType
				func miqt_exec_callback_QImage_DevType(self QImage, cb intptr_t) int{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			
virtualReturn := gofunc((&QImage{h: self}).callVirtualBase_DevType )

return (int)(virtualReturn)

				}
			
				func (this *QImage) callVirtualBase_PaintEngine() *QPaintEngine {
					
					return newQPaintEngine(QImage_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

				}
			func (this *QImage) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QImage_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QImage_PaintEngine
				func miqt_exec_callback_QImage_PaintEngine(self QImage, cb intptr_t) *QPaintEngine{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			
virtualReturn := gofunc((&QImage{h: self}).callVirtualBase_PaintEngine )

return virtualReturn.cPointer()

				}
			
				func (this *QImage) callVirtualBase_Metric(metric PaintDeviceMetric) int {
					
					return (int)(QImage_virtualbase_Metric(unsafe.Pointer(this.h), metric))

				}
			func (this *QImage) OnMetric(slot func(super func(metric PaintDeviceMetric) int, metric PaintDeviceMetric) int) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QImage_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QImage_Metric
				func miqt_exec_callback_QImage_Metric(self QImage, cb intptr_t, metric PaintDeviceMetric) int{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func(metric PaintDeviceMetric) int, metric PaintDeviceMetric) int)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			// Convert all CABI parameters to Go parameters
xxxxxxxxx

virtualReturn := gofunc((&QImage{h: self}).callVirtualBase_Metric, slotval1 )

return (int)(virtualReturn)

				}
			
				func (this *QImage) callVirtualBase_InitPainter(painter *QPainter)  {
					
					 QImage_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

				}
			func (this *QImage) OnInitPainter(slot func(super func(painter *QPainter) , painter *QPainter) ) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QImage_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QImage_InitPainter
				func miqt_exec_callback_QImage_InitPainter(self QImage, cb intptr_t, painter *QPainter) {
					gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter) , painter *QPainter) )
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			// Convert all CABI parameters to Go parameters
slotval1 := newQPainter(painter)


gofunc((&QImage{h: self}).callVirtualBase_InitPainter, slotval1 )

				}
			
				func (this *QImage) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {
					
					return newQPaintDevice(QImage_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

				}
			func (this *QImage) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QImage_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QImage_Redirected
				func miqt_exec_callback_QImage_Redirected(self QImage, cb intptr_t, offset *QPoint) *QPaintDevice{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			// Convert all CABI parameters to Go parameters
slotval1 := newQPoint(offset)


virtualReturn := gofunc((&QImage{h: self}).callVirtualBase_Redirected, slotval1 )

return virtualReturn.cPointer()

				}
			
				func (this *QImage) callVirtualBase_SharedPainter() *QPainter {
					
					return newQPainter(QImage_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

				}
			func (this *QImage) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QImage_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QImage_SharedPainter
				func miqt_exec_callback_QImage_SharedPainter(self QImage, cb intptr_t) *QPainter{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			
virtualReturn := gofunc((&QImage{h: self}).callVirtualBase_SharedPainter )

return virtualReturn.cPointer()

				}
			