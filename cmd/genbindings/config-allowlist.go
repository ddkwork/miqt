package main

import (
	"path/filepath"
	"strings"

	"github.com/ddkwork/golibrary/mylog"
)

func InsertTypedefs() {
	// Seed well-known typedefs
	const pkgName = "qt"

	// QString is deleted from this binding
	KnownTypedefs["QStringList"] = lookupResultTypedef{pkgName, CppTypedef{"QStringList", parseSingleTypeString("QList<QString>")}}

	// FIXME this isn't picked up automatically because QFile inherits QFileDevice and the name refers to its parent class
	KnownTypedefs["QFile::FileTime"] = lookupResultTypedef{pkgName, CppTypedef{"QFile::FileTime", parseSingleTypeString("QFileDevice::FileTime")}}

	KnownTypedefs["QLineF::IntersectionType"] = lookupResultTypedef{pkgName, CppTypedef{"QLineF::IntersectionType", parseSingleTypeString("QLineF::IntersectType")}}

	// Not sure the reason for this one
	KnownTypedefs["QSocketDescriptor::DescriptorType"] = lookupResultTypedef{pkgName, CppTypedef{"QSocketDescriptor::DescriptorType", parseSingleTypeString("QSocketNotifier::Type")}}

	// QFile doesn't see QFileDevice parent class enum
	KnownTypedefs["QFile::Permissions"] = lookupResultTypedef{pkgName, CppTypedef{"QFile::Permissions", parseSingleTypeString("QFileDevice::Permissions")}}
	KnownTypedefs["QFileDevice::Permissions"] = lookupResultTypedef{pkgName, CppTypedef{"QFile::Permissions", parseSingleTypeString("QFlags<QFileDevice::Permission>")}}
	KnownTypedefs["QIODevice::OpenMode"] = lookupResultTypedef{pkgName, CppTypedef{"QIODevice::OpenMode", parseSingleTypeString("QIODeviceBase::OpenMode")}}

	// Qt 5 WebKit - use of an empty enum (should be possible to support?)
	KnownEnums["QWebPluginFactory::Extension"] = lookupResultEnum{"qt/webkit", CppEnum{
		EnumName: "QWebPluginFactory::Extension",
		UnderlyingType: CppParameter{
			ParameterType: "int",
		},
	}}
	// Qt 6 QVariant helper types - needs investigation
	KnownTypedefs["QVariantHash"] = lookupResultTypedef{pkgName, CppTypedef{"QVariantHash", parseSingleTypeString("QHash<QString,QVariant>")}}
	KnownTypedefs["QVariantList"] = lookupResultTypedef{pkgName, CppTypedef{"QVariantList", parseSingleTypeString("QList<QVariant>")}}
	KnownTypedefs["QVariantMap"] = lookupResultTypedef{pkgName, CppTypedef{"QVariantMap", parseSingleTypeString("QMap<QString,QVariant>")}}

	// Qt 6 renamed the enum to LibraryPath, but left some uses of LibraryLocation with a typedef
	// We don't find the typedef - needs investigation
	// ONLY add this on Qt 6 builds, breaks Qt 5
	KnownTypedefs["QLibraryInfo::LibraryLocation"] = lookupResultTypedef{pkgName, CppTypedef{"QLibraryInfo::LibraryLocation", parseSingleTypeString("QLibraryInfo::LibraryPath")}}

	// Enums
	// QSysInfo.h is being truncated and not finding any content
	KnownEnums["QSysInfo::Endian"] = lookupResultEnum{pkgName, CppEnum{
		EnumName: "QSysInfo::Endian",
		UnderlyingType: CppParameter{
			ParameterType: "int",
		},
	}}
}

func Widgets_AllowHeader(fullpath string) bool {
	fname := filepath.Base(fullpath)
	if strings.HasSuffix(fname, `_impl.h`) {
		return false // Not meant to be imported
	}
	fname_lc := strings.ToLower(fname)
	if strings.Contains(fname_lc, `opengl`) || strings.Contains(fname_lc, `vulkan`) {
		return false // Too hard
	}
	switch fname {
	case "qatomic_bootstrap.h",
		"qatomic_cxx11.h",
		"qatomic_msvc.h",
		"qgenericatomic.h",             // Clang error
		"qt_windows.h",                 // Clang error
		"qmaccocoaviewcontainer_mac.h", // Needs NSView* headers. TODO allow with darwin build tag
		"qmacnativewidget_mac.h",       // Needs NSView* headers. TODO allow with darwin build tag
		"qstring.h",                    // QString does not exist in this binding
		"qbytearray.h",                 // QByteArray does not exist in this binding
		"qlist.h",                      // QList does not exist in this binding
		"qvector.h",                    // QVector does not exist in this binding
		"qhash.h",                      // QHash does not exist in this binding
		"qmap.h",                       // QMap does not exist in this binding
		"qtcoreexports.h",              // Nothing bindable here and has Q_CORE_EXPORT definition issues
		"q20algorithm.h",               // Qt 6 unstable header
		"q20functional.h",              // Qt 6 unstable header
		"q20iterator.h",                // Qt 6 unstable header
		"q23functional.h",              // Qt 6 unstable header
		"qguiapplication_platform.h",   // Qt 6 - can be built for X11 but then platform-specific code fails to build on Windows
		"____last____":
		return false
	}
	return true
}

func ImportHeaderForClass(className string) bool {
	switch {
	case className[0] != 'Q':
		return false
	case strings.HasPrefix(className, "QPlatform"):
		// e.g. QPlatformPixmap, QPlatformWindow, QPlatformScreen
		// These classes don't have a <> version to include
		return false
	case strings.HasPrefix(className, "Qsci"):
		// QScintilla - does not produce imports
		return false
	}
	// TODO this could be implict by checking if files exist in known header paths
	switch className {
	case "QGraphicsEffectSource", // e.g. qgraphicseffect.h
		"QAbstractConcatenable", // qstringbuilder.h
		"QTextEngine",           // qtextlayout.h
		"QText",                 // e.g. qtextcursor.h
		"QVLABaseBase",          // e.g. Qt 6 qvarlengtharray.h
		"QAdoptSharedDataTag",   // Qt 6 qshareddata.h
		"____last____":
		return false
	}
	return true
}

func AllowClass(className string) bool {
	switch {
	case strings.HasSuffix(className, "Private") || strings.HasSuffix(className, "PrivateShared") ||
		strings.Contains(className, "Private::") || strings.HasSuffix(className, "PrivateShared::"):
		return false
	case strings.Contains(className, "QPrivateSignal"):
		return false
	case strings.HasPrefix(className, `std::`):
		return false // Scintilla bindings find some of these
	}
	switch className {
	case
		"QTextStreamManipulator", // Only seems to contain garbage methods
		"QException",             // Extends std::exception, too hard
		"QUnhandledException",    // As above (child class)
		// "QItemSelection",             // Extends a QList<>, too hard
		"QXmlStreamAttributes",       // Extends a QList<>, too hard
		"QPolygon",                   // Extends a QVector<QPoint> template class, too hard
		"QPolygonF",                  // Extends a QVector<QPoint> template class, too hard
		"QAssociativeIterator",       // Qt 6. Extends a QIterator<>, too hard todo use go1.23 iter pkg ?
		"QAssociativeConstIterator",  // Qt 6. Extends a QIterator<>, too hard
		"QAssociativeIterable",       // Qt 6. Extends a QIterator<>, too hard
		"QSequentialIterator",        // Qt 6. Extends a QIterator<>, too hard
		"QSequentialConstIterator",   // Qt 6. Extends a QIterator<>, too hard
		"QSequentialIterable",        // Qt 6. Extends a QIterator<>, too hard
		"QBrushDataPointerDeleter",   // Qt 6 qbrush.h. Appears in header but cannot be linked
		"QPropertyBindingPrivatePtr", // Qt 6 qpropertyprivate.h. Appears in header but cannot be linked
		"QDeferredDeleteEvent",       // Qt 6. Hidden/undocumented class in Qt 6.4, moved to private header in Qt 6.7. Intended for test use only

		"QUntypedPropertyData::InheritsQUntypedPropertyData", // qpropertyprivate.h . Hidden/undocumented class in Qt 6.4, removed in 6.7
		"____last____":
		return false
	}

	return true
}

func AllowSignal(mm CppMethod) bool {
	if !mm.ReturnType.Void() {
		// This affects how we cast the signal function pointer for connect
		// It would be fixable, but, real signals always have void return types anyway
		return false
	}
	switch mm.MethodName {
	case `metaObject`, `qt_metacast`,
		`clone`: // Qt 6 - qcoreevent.h
		return false
	default:
		return true
	}
}

func AllowVirtual(mm CppMethod) bool {
	return mm.MethodName == "metaObject" || mm.MethodName == "qt_metacast" // AllowSignal(mm)
}

func AllowVirtualForClass(className string) bool {
	// Allowing the subclassing of QAccessibleWidget compiles fine,
	// but, always gives a linker error:
	//
	//   /usr/lib/go-1.19/pkg/tool/linux_amd64/link: running g++ failed: exit status 1
	//   /usr/bin/ld: /tmp/go-link-1745036494/000362.o: in function `MiqtVirtualQAccessibleWidget::MiqtVirtualQAccessibleWidget(QWidget*)':
	//   undefined reference to `vtable for MiqtVirtualQAccessibleWidget'
	//
	// An undefined vtable usually indicates that the virtual class is missing
	// definitions for some virtual methods, but AFAICT we have complete coverage.
	switch className {
	case "QAccessibleWidget", "QAccessibleObject", "QWebHapticFeedbackPlayer":
		return false
	case "QFutureWatcherBase": // Pure virtual method futureInterface() returns an unprojectable template type
		return false
	case "QObjectData": // Pure virtual dtor (should be possible to support)
		return false
	case "QAbstractEventDispatcher":
		// Pure virtual method registerEventNotifier takes a QWinEventNotifier* on Windows
		// which is platform-specific
		return false
	case "QWebNotificationPresenter": // Qt 5 QWebkit: undefined reference to typeinfo
		return false
	}
	return true
}

func AllowMethod(className string, mm CppMethod) error {
	for _, p := range mm.Parameters {
		if strings.HasSuffix(p.ParameterType, "Private") {
			return ErrTooComplex // Skip private type
		}
	}
	if strings.HasSuffix(mm.ReturnType.ParameterType, "Private") {
		return ErrTooComplex // Skip private type
	}
	if strings.Contains(mm.MethodName, `QGADGET`) {
		return ErrTooComplex // Skipping method with weird QGADGET behaviour
	}
	if mm.IsReceiverMethod() {
		// Non-projectable receiver pattern parameters
		return ErrTooComplex
	}
	if className == "QBitArray" && mm.MethodName == "operator~" {
		return ErrTooComplex // Present in Qt 5.15 and 6.4, missing in Qt 6.7
	}
	if className == "QTimeZone" && (mm.MethodName == "operator==" || mm.MethodName == "operator!=") {
		return ErrTooComplex // Present in Qt 5.15 and 6.4, missing in Qt 6.7
	}
	if className == "QWaveDecoder" && mm.MethodName == "setIODevice" {
		return ErrTooComplex // Qt 6: Present in header, but no-op method was not included in compiled library
	}
	if className == "QDeadlineTimer" && mm.MethodName == "_q_data" {
		// Qt 6.4: Present in header with "not a public method" comment, not present in Qt 6.6
		// @ref https://github.com/qt/qtbase/blob/v6.4.0/src/corelib/kernel/qdeadlinetimer.h#L156C29-L156C36
		return ErrTooComplex
	}
	if className == "QXmlStreamEntityResolver" && mm.MethodName == "operator=" {
		// Present in Qt 6.7, but marked as =delete by Q_DISABLE_COPY_MOVE in Qt 6.8
		return ErrTooComplex
	}
	return nil // OK, allow
}

// AllowType controls whether to permit binding of a method, if a method uses
// this type in its parameter list or return type.
// Any type not permitted by AllowClass is also not permitted by this method.
func AllowType(p CppParameter, isReturnType bool) error {
	if t, ok := p.QSetOf(); ok {
		mylog.Check(AllowType(t, isReturnType))
	}
	if t, ok := p.QListOf(); ok {
		if e := AllowType(t, isReturnType); e != nil { // e.g. QGradientStops is a QVector<> (OK) of QGradientStop (not OK)
			return e
		}
		// qsciscintilla.h QsciScintilla_Annotate4: no copy ctor for private type QsciStyledText
		// Works fine normally, but not in a list
		if t.ParameterType == "QsciStyledText" {
			return ErrTooComplex
		}
	}
	if kType, vType, ok := p.QMapOf(); ok {
		mylog.Check(AllowType(kType, isReturnType))
		mylog.Check(AllowType(vType, isReturnType))
		// Additionally, Go maps do not support []byte keys
		// This affects qnetwork qsslconfiguration BackendConfiguration
		if kType.ParameterType == "QByteArray" {
			return ErrTooComplex
		}
	}
	if kType, vType, ok := p.QPairOf(); ok {
		mylog.Check(AllowType(kType, isReturnType))
		mylog.Check(AllowType(vType, isReturnType))
	}
	switch {
	case p.QMultiMapOf():
		return ErrTooComplex // e.g. Qt5 QNetwork qsslcertificate.h has a QMultiMap<QSsl::AlternativeNameEntryType, QString>
	case !AllowClass(p.ParameterType):
		return ErrTooComplex // This whole class type has been blocked, not only as a parameter/return type
	case strings.Contains(p.ParameterType, "(*)"): // Function pointer.
		return ErrTooComplex // e.g. QAccessible_InstallFactory
	case strings.HasPrefix(p.ParameterType, "StringResult<"):
		return ErrTooComplex // e.g. qcborstreamreader.h
	case strings.HasPrefix(p.ParameterType, "QScopedPointer<"):
		return ErrTooComplex // e.g. qbrush.h
	case strings.HasPrefix(p.ParameterType, "QExplicitlySharedDataPointer<"):
		return ErrTooComplex // e.g. qpicture.h
	case strings.HasPrefix(p.ParameterType, "QSharedDataPointer<"):
		return ErrTooComplex // e.g. qurlquery.h
	case strings.HasPrefix(p.ParameterType, "QTypedArrayData<"):
		return ErrTooComplex // e.g. qbitarray.h
	case strings.HasPrefix(p.ParameterType, "QGenericMatrix<"):
		return ErrTooComplex // e.g. qmatrix4x4.h
	case strings.HasPrefix(p.ParameterType, "QUrlTwoFlags<"):
		return ErrTooComplex // e.g. qurl.h
	case strings.HasPrefix(p.ParameterType, "FillResult<"):
		return ErrTooComplex // Scintilla
	case strings.HasPrefix(p.ParameterType, "QBindable<"):
		return ErrTooComplex // e.g. Qt 6 qabstractanimation.h
	case strings.HasPrefix(p.ParameterType, "QRgbaFloat<"):
		return ErrTooComplex // e.g. Qt 6 qcolortransform.h
	case strings.HasPrefix(p.ParameterType, "QPointer<"):
		return ErrTooComplex // e.g. Qt 6 qevent.h . It should be possible to support this
	case strings.HasPrefix(p.ParameterType, "EncodedData<"):
		return ErrTooComplex // e.g. Qt 6 qstringconverter.h
	case strings.HasPrefix(p.ParameterType, "QQmlListProperty<"):
		return ErrTooComplex // e.g. Qt 5 QWebChannel qmlwebchannel.h . Supporting this will be required for QML in future
	case strings.HasPrefix(p.ParameterType, "QWebEngineCallback<"):
		return ErrTooComplex // Function pointer types in QtWebEngine
	case strings.HasPrefix(p.ParameterType, "std::"):
		// std::initializer           e.g. qcborarray.h
		// std::string                QByteArray->toStdString(). There are QString overloads already
		// std::nullptr_t             Qcborstreamwriter
		// std::chrono::nanoseconds   QDeadlineTimer_RemainingTimeAsDuration
		// std::seed_seq              QRandom
		// std::exception             Scintilla
		return ErrTooComplex
	case strings.Contains(p.ParameterType, `Iterator::value_type`):
		return ErrTooComplex // e.g. qcbormap
	case strings.Contains(p.ParameterType, `>::iterator`) ||
		strings.Contains(p.ParameterType, `>::const_iterator`):
		// qresultstore.h tries to create a
		// NewQtPrivate__ResultIteratorBase2(_mapIterator QMap<int, ResultItem>__const_iterator)
		return ErrTooComplex
	case strings.Contains(p.ParameterType, `::QPrivate`):
		return ErrTooComplex // e.g. QAbstractItemModel::QPrivateSignal
	case strings.Contains(p.GetQtCppType().ParameterType, `::DataPtr`):
		return ErrTooComplex // e.g. QImage::data_ptr()
	case strings.Contains(p.ParameterType, `::DataPointer`):
		return ErrTooComplex // Qt 6 qbytearray.h. This could probably be made to work
	case strings.HasPrefix(p.ParameterType, `QArrayDataPointer<`):
		return ErrTooComplex // Qt 6 qbytearray.h. This could probably be made to work
	case p.ParameterType[0] == 'Q' && strings.HasSuffix(p.ParameterType, "Private"):
		// Some QFoo constructors take a QFooPrivate
		// QIcon also returns a QIconPrivate
		return ErrTooComplex
	case strings.HasPrefix(p.ParameterType, "QtPrivate::"):
		return ErrTooComplex // e.g. Qt 6 qbindingstorage.h
	case p.ParameterType == "QString" && p.Pointer && !isReturnType: // Out-parameters
		// If any parameters are QString*, skip the method
		// QDebug constructor
		// QXmlStreamWriter constructor
		// QFile::moveToTrash
		// QLockFile::getLockInfo
		// QTextDecoder::toUnicode
		// QTextStream::readLineInto
		// QFileDialog::getOpenFileName selectedFilter* param
		return ErrTooComplex
	case p.ParameterType == "QByteArray" && p.Pointer && !isReturnType:
		// QBuffer can accept a raw pointer to an internal QByteArray, but that
		// doesn't work when QByteArray is deleted
		// QDataStream
		return ErrTooComplex
	case p.ParameterType == "QFormLayout::ItemRole" && p.Pointer && !isReturnType:
		return ErrTooComplex // Out-parameters in QFormLayout
	case p.Pointer && p.PointerCount >= 2: // Out-parameters
		if p.ParameterType != "char" {
			return ErrTooComplex // e.g. QGraphicsItem_IsBlockedByModalPanel1
		}
		if p.ParameterType == "char" && p.ParameterName == "xpm" {
			// Array parameters:
			// - QPixmap and QImage ctors from an xpm char*[]
			// TODO support these
			return ErrTooComplex
		}
	}

	switch p.ParameterType {
	case
		"QList<QVariant>",       // e.g. QVariant constructor - this has a deleted copy-constructor so we can't get it over the CABI boundary by value
		"QPolygon", "QPolygonF", // QPolygon extends a template type
		"QGenericMatrix", "QMatrix3x3", // extends a template type
		"QLatin1String", "QStringView", // e.g. QColor constructors and QColor::SetNamedColor() overloads. These are usually optional alternatives to QString
		"QLatin1StringView",               // Qt 6 - used in qanystringview
		"QUtf8StringView",                 // Qt 6 - used in qdebug
		"QStringRef",                      // e.g. QLocale::toLongLong and similar overloads. As above
		"qfloat16",                        // e.g. QDataStream - there is no such half-float type in C or Go
		"char16_t",                        // e.g. QChar() constructor overload, just unnecessary
		"char32_t",                        // e.g. QDebug().operator<< overload, unnecessary
		"wchar_t",                         // e.g. qstringview.h overloads, unnecessary
		"FILE",                            // e.g. qfile.h constructors
		"sockaddr",                        // Qt network Qhostaddress. Should be possible to make this work but may be platform-specific
		"qInternalCallback",               // e.g. qnamespace.h
		"QGraphicsEffectSource",           // e.g. used by qgraphicseffect.h, but the definition is in ????
		"QXmlStreamEntityDeclarations",    // e.g. qxmlstream.h. The class definition was blacklisted for ???? reason so don't allow it as a parameter either
		"QXmlStreamNamespaceDeclarations", // e.g. qxmlstream.h. As above
		"QXmlStreamNotationDeclarations",  // e.g. qxmlstream.h. As above
		"QXmlStreamAttributes",            // e.g. qxmlstream.h
		"LineLayout::ValidLevel",          // ..
		"QtMsgType",                       // e.g. qdebug.h TODO Defined in qlogging.h, but omitted because it's predefined in qglobal.h, and our clangexec is too agressive
		"QTextStreamFunction",             // e.g. qdebug.h
		"QFactoryInterface",               // qfactoryinterface.h
		"QTextEngine",                     // used by qtextlayout.h, also blocked in ImportHeaderForClass above
		"QVulkanInstance",                 // e.g. qwindow.h. Not tackling vulkan yet
		"QPlatformNativeInterface",        // e.g. QGuiApplication::platformNativeInterface(). Private type, could probably expose as uintptr. n.b. Changes in Qt6
		"QPlatformBackingStore",           // e.g. qbackingstore.h, as below
		"QPlatformMenuBar",                // e.g. qfutureinterface.h, as below
		"QPlatformOffscreenSurface",       // e.g. qoffscreensurface.h, as below
		"QPlatformPixmap",                 // e.g. qpixmap.h. as below
		"QPlatformScreen",                 // e.g. qscreen.h. as below
		"QPlatformWindow",                 // e.g. qwindow.h, as below
		"QPlatformSurface",                // e.g. qsurface.h. as below
		"QPlatformMenu",                   // e.g. QMenu_PlatformMenu. Defined in the QPA, could probably expose as uintptr
		"QPlatformMediaCaptureSession",    // Qt 6 Multimedia qmediacapturesession.h
		"QPlatformMediaRecorder",          // Qt 6 Multimedia qmediarecorder.h
		"QPlatformVideoSink",              // Qt 6 Multimedia qvideosink.h
		"QTextDocument::ResourceProvider", // Qt 6 typedef for unsupported std::function<QVariant(const QUrl&)>
		"QTransform::Affine",              // Qt 6 qtransform.h - public method returning private type
		"QAbstractAudioBuffer",            // Qt 5 Multimedia, this is a private/internal type only
		"QAbstractVideoBuffer",            // Works in Qt 5, but in Qt 6 Multimedia this type is used in qvideoframe.h but is not defined anywhere (it was later added in Qt 6.8)
		"QRhi",                            // Qt 6 unstable types, used in Multimedia
		"QPostEventList",                  // Qt QCoreApplication: private headers required
		"QMetaCallEvent",                  // ..
		"QPostEvent",                      // ..
		"QWebFrameAdapter",                // Qt 5 Webkit: Used by e.g. qwebframe.h but never defined anywhere
		"QWebPageAdapter",                 // ...
		"QQmlWebChannelAttached",          // Qt 5 qqmlwebchannel.h. Need to add QML support for this to work
		"____last____":
		return ErrTooComplex
	}
	// Should be OK
	return nil
}

// LinuxWindowsCompatCheck checks if the parameter is incompatible between the
// generated headers (generated on Linux) with other OSes such as Windows.
// These methods will be blocked on non-Linux OSes.
func LinuxWindowsCompatCheck(p CppParameter) bool {
	if p.GetQtCppType().ParameterType == "Q_PID" {
		return true // int64 on Linux, _PROCESS_INFORMATION* on Windows
	}
	if p.GetQtCppType().ParameterType == "QSocketDescriptor::DescriptorType" {
		return true // uintptr_t-compatible on Linux, void* on Windows
	}
	return false
}

func ApplyQuirks(className string, mm *CppMethod) {
	if className == "QArrayData" && mm.MethodName == "needsDetach" && mm.IsConst {
		mm.BecomesNonConstInVersion = addr("6.7")
	}
	if className == "QFileDialog" && mm.MethodName == "saveFileContent" && mm.IsStatic {
		// The prototype was changed from
		// [Qt 5 - 6.6] void QFileDialog::saveFileContent(const QByteArray &fileContent, const QString &fileNameHint = QString())
		// [Qt 6.7]     void QFileDialog::saveFileContent(const QByteArray &fileContent, const QString &fileNameHint, QWidget *parent = nullptr)
		// The 2nd parameter is no longer optional
		// As a compromise, make it non-optional everywhere
		mm.Parameters[1].Optional = false
	}
}
