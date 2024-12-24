package qt

import (
	"unsafe"
)

type QPaintDevice__PaintDeviceMetric int

const (
	QPaintDevice__PdmWidth                      QPaintDevice__PaintDeviceMetric = 1
	QPaintDevice__PdmHeight                     QPaintDevice__PaintDeviceMetric = 2
	QPaintDevice__PdmWidthMM                    QPaintDevice__PaintDeviceMetric = 3
	QPaintDevice__PdmHeightMM                   QPaintDevice__PaintDeviceMetric = 4
	QPaintDevice__PdmNumColors                  QPaintDevice__PaintDeviceMetric = 5
	QPaintDevice__PdmDepth                      QPaintDevice__PaintDeviceMetric = 6
	QPaintDevice__PdmDpiX                       QPaintDevice__PaintDeviceMetric = 7
	QPaintDevice__PdmDpiY                       QPaintDevice__PaintDeviceMetric = 8
	QPaintDevice__PdmPhysicalDpiX               QPaintDevice__PaintDeviceMetric = 9
	QPaintDevice__PdmPhysicalDpiY               QPaintDevice__PaintDeviceMetric = 10
	QPaintDevice__PdmDevicePixelRatio           QPaintDevice__PaintDeviceMetric = 11
	QPaintDevice__PdmDevicePixelRatioScaled     QPaintDevice__PaintDeviceMetric = 12
	QPaintDevice__PdmDevicePixelRatioF_EncodedA QPaintDevice__PaintDeviceMetric = 13
	QPaintDevice__PdmDevicePixelRatioF_EncodedB QPaintDevice__PaintDeviceMetric = 14
)

type QPaintDevice struct {
	h          uintptr
	isSubclass bool
}

func (this *QPaintDevice) DevType() int {
	return (int)(QPaintDevice_DevType(this.h))
}

func (this *QPaintDevice) PaintingActive() bool {
	return (bool)(QPaintDevice_PaintingActive(this.h))
}

func (this *QPaintDevice) PaintEngine() *QPaintEngine {
	return newQPaintEngine(QPaintDevice_PaintEngine(this.h))
}

func (this *QPaintDevice) Width() int {
	return (int)(QPaintDevice_Width(this.h))
}

func (this *QPaintDevice) Height() int {
	return (int)(QPaintDevice_Height(this.h))
}

func (this *QPaintDevice) WidthMM() int {
	return (int)(QPaintDevice_WidthMM(this.h))
}

func (this *QPaintDevice) HeightMM() int {
	return (int)(QPaintDevice_HeightMM(this.h))
}

func (this *QPaintDevice) LogicalDpiX() int {
	return (int)(QPaintDevice_LogicalDpiX(this.h))
}

func (this *QPaintDevice) LogicalDpiY() int {
	return (int)(QPaintDevice_LogicalDpiY(this.h))
}

func (this *QPaintDevice) PhysicalDpiX() int {
	return (int)(QPaintDevice_PhysicalDpiX(this.h))
}

func (this *QPaintDevice) PhysicalDpiY() int {
	return (int)(QPaintDevice_PhysicalDpiY(this.h))
}

func (this *QPaintDevice) DevicePixelRatio() float64 {
	return (float64)(QPaintDevice_DevicePixelRatio(this.h))
}

func (this *QPaintDevice) DevicePixelRatioF() float64 {
	return (float64)(QPaintDevice_DevicePixelRatioF(this.h))
}

func (this *QPaintDevice) ColorCount() int {
	return (int)(QPaintDevice_ColorCount(this.h))
}

func (this *QPaintDevice) Depth() int {
	return (int)(QPaintDevice_Depth(this.h))
}

func QPaintDevice_DevicePixelRatioFScale() float64 {
	return (float64)(QPaintDevice_DevicePixelRatioFScale())
}

func QPaintDevice_EncodeMetricF(metric PaintDeviceMetric, value float64) int {
	return (int)(QPaintDevice_EncodeMetricF(metric, (double)(value)))
}
