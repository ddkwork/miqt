package qt
import (
	"unsafe"
)

		type QUuid__Variant int
		const (
QUuid__VarUnknown QUuid__Variant = -1
QUuid__NCS QUuid__Variant = 0
QUuid__DCE QUuid__Variant = 2
QUuid__Microsoft QUuid__Variant = 6
QUuid__Reserved QUuid__Variant = 7

)


		type QUuid__Version int
		const (
QUuid__VerUnknown QUuid__Version = -1
QUuid__Time QUuid__Version = 1
QUuid__EmbeddedPOSIX QUuid__Version = 2
QUuid__Md5 QUuid__Version = 3
QUuid__Name QUuid__Version = 3
QUuid__Random QUuid__Version = 4
QUuid__Sha1 QUuid__Version = 5
QUuid__UnixEpoch QUuid__Version = 7

)


		type QUuid__StringFormat int
		const (
QUuid__WithBraces QUuid__StringFormat = 0
QUuid__WithoutBraces QUuid__StringFormat = 1
QUuid__Id128 QUuid__StringFormat = 3

)


		type QUuid struct {
			h uintptr
			isSubclass bool
}
		
			// NewQUuid constructs a new QUuid object.
			func NewQUuid() *QUuid {
								
				ret := newQUuid(QUuid_new())
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQUuid2 constructs a new QUuid object.
			func NewQUuid2(l uint, w1 uint16, w2 uint16, b1 byte, b2 byte, b3 byte, b4 byte, b5 byte, b6 byte, b7 byte, b8 byte) *QUuid {
								
				ret := newQUuid(QUuid_new2((uint)(l), (uint16_t)(w1), (uint16_t)(w2), (uchar)(b1), (uchar)(b2), (uchar)(b3), (uchar)(b4), (uchar)(b5), (uchar)(b6), (uchar)(b7), (uchar)(b8)))
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQUuid3 constructs a new QUuid object.
			func NewQUuid3(id128 Id128Bytes) *QUuid {
								
				ret := newQUuid(QUuid_new3(id128))
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQUuid4 constructs a new QUuid object.
			func NewQUuid4(stringVal QAnyStringView) *QUuid {
								
				ret := newQUuid(QUuid_new4(stringVal.cPointer()))
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQUuid5 constructs a new QUuid object.
			func NewQUuid5(guid *struct _GUID) *QUuid {
								
				ret := newQUuid(QUuid_new5(guid))
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQUuid6 constructs a new QUuid object.
			func NewQUuid6(param1 *QUuid) *QUuid {
								
				ret := newQUuid(QUuid_new6(param1.cPointer()))
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQUuid7 constructs a new QUuid object.
			func NewQUuid7(id128 Id128Bytes, order QSysInfo__Endian) *QUuid {
								
				ret := newQUuid(QUuid_new7(id128, (int)(order)))
				ret.isSubclass = true
				return ret
			}
			
			
			func QUuid_FromString(stringVal QAnyStringView) *QUuid {
				_goptr := newQUuid(QUuid_FromString(stringVal.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QUuid) ToString() string {
				var _ms struct_miqt_string =  QUuid_ToString(this.h)
_ret := GoStringN(_ms.data, int(int64(_ms.len)))
free(unsafe.Pointer(_ms.data))
return _ret}
			
			func (this *QUuid) ToByteArray() []byte {
				var _bytearray struct_miqt_string =  QUuid_ToByteArray(this.h)
_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
free(unsafe.Pointer(_bytearray.data))
return _ret}
			
			func (this *QUuid) ToBytes() Id128Bytes {
				xxxxxxxxx}
			
			func (this *QUuid) ToRfc4122() []byte {
				var _bytearray struct_miqt_string =  QUuid_ToRfc4122(this.h)
_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
free(unsafe.Pointer(_bytearray.data))
return _ret}
			
			func QUuid_FromBytes(bytes unsafe.Pointer) *QUuid {
				_goptr := newQUuid(QUuid_FromBytes(bytes))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func QUuid_FromRfc4122(param1 QByteArrayView) *QUuid {
				_goptr := newQUuid(QUuid_FromRfc4122(param1.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QUuid) IsNull() bool {
				return (bool)(QUuid_IsNull(this.h))
}
			
			func (this *QUuid) OperatorAssign(guid *struct _GUID)  {
				 QUuid_OperatorAssign(this.h, guid)
}
			
			func QUuid_CreateUuid() *QUuid {
				_goptr := newQUuid(QUuid_CreateUuid())
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func QUuid_CreateUuidV5(ns QUuid, baseData QByteArrayView) *QUuid {
				_goptr := newQUuid(QUuid_CreateUuidV5(ns.cPointer(), baseData.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func QUuid_CreateUuidV3(ns QUuid, baseData QByteArrayView) *QUuid {
				_goptr := newQUuid(QUuid_CreateUuidV3(ns.cPointer(), baseData.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func QUuid_CreateUuidV7() *QUuid {
				_goptr := newQUuid(QUuid_CreateUuidV7())
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QUuid) Variant() Variant {
				xxxxxxxxx}
			
			func (this *QUuid) Version() Version {
				xxxxxxxxx}
			
			func (this *QUuid) OperatorAssignWithQUuid(param1 *QUuid)  {
				 QUuid_OperatorAssignWithQUuid(this.h, param1.cPointer())
}
			
			func (this *QUuid) ToString1(mode StringFormat) string {
				var _ms struct_miqt_string =  QUuid_ToString1(this.h, mode)
_ret := GoStringN(_ms.data, int(int64(_ms.len)))
free(unsafe.Pointer(_ms.data))
return _ret}
			
			func (this *QUuid) ToByteArray1(mode StringFormat) []byte {
				var _bytearray struct_miqt_string =  QUuid_ToByteArray1(this.h, mode)
_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
free(unsafe.Pointer(_bytearray.data))
return _ret}
			
			func (this *QUuid) ToBytes1(order QSysInfo__Endian) Id128Bytes {
				xxxxxxxxx}
			
			func QUuid_FromBytes2(bytes unsafe.Pointer, order QSysInfo__Endian) *QUuid {
				_goptr := newQUuid(QUuid_FromBytes2(bytes, (int)(order)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QUuid) IsNull1(param1 Disambiguated_t) bool {
				return (bool)(QUuid_IsNull1(this.h, param1.cPointer()))
}
			
			func (this *QUuid) Variant1(param1 Disambiguated_t) Variant {
				xxxxxxxxx}
			
			func (this *QUuid) Version1(param1 Disambiguated_t) Version {
				xxxxxxxxx}
			
		type QUuid__Id128Bytes struct {
			h uintptr
			isSubclass bool
}
		
			// NewQUuid__Id128Bytes constructs a new QUuid::Id128Bytes object.
			func NewQUuid__Id128Bytes() *QUuid__Id128Bytes {
								
				ret := newQUuid__Id128Bytes(QUuid__Id128Bytes_new())
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQUuid__Id128Bytes2 constructs a new QUuid::Id128Bytes object.
			func NewQUuid__Id128Bytes2(param1 *Id128Bytes) *QUuid__Id128Bytes {
								
				ret := newQUuid__Id128Bytes(QUuid__Id128Bytes_new2(param1))
				ret.isSubclass = true
				return ret
			}
			
			