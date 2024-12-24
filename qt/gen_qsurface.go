package qt

import (
	"unsafe"
)

type QSurface__SurfaceClass int

const (
	QSurface__Window    QSurface__SurfaceClass = 0
	QSurface__Offscreen QSurface__SurfaceClass = 1
)

type QSurface__SurfaceType int

const (
	QSurface__RasterSurface   QSurface__SurfaceType = 0
	QSurface__OpenGLSurface   QSurface__SurfaceType = 1
	QSurface__RasterGLSurface QSurface__SurfaceType = 2
	QSurface__OpenVGSurface   QSurface__SurfaceType = 3
	QSurface__VulkanSurface   QSurface__SurfaceType = 4
	QSurface__MetalSurface    QSurface__SurfaceType = 5
	QSurface__Direct3DSurface QSurface__SurfaceType = 6
)

type QSurface struct {
	h          uintptr
	isSubclass bool
}

func (this *QSurface) SurfaceClass() SurfaceClass {
	xxxxxxxxx
}

func (this *QSurface) Format() *QSurfaceFormat {
	_goptr := newQSurfaceFormat(QSurface_Format(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSurface) SurfaceType() SurfaceType {
	xxxxxxxxx
}

func (this *QSurface) SupportsOpenGL() bool {
	return (bool)(QSurface_SupportsOpenGL(this.h))
}

func (this *QSurface) Size() *QSize {
	_goptr := newQSize(QSurface_Size(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
