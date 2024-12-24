package qt

import (
	"unsafe"
)

type QGradient__Type int

const (
	QGradient__LinearGradient  QGradient__Type = 0
	QGradient__RadialGradient  QGradient__Type = 1
	QGradient__ConicalGradient QGradient__Type = 2
	QGradient__NoGradient      QGradient__Type = 3
)

type QGradient__Spread int

const (
	QGradient__PadSpread     QGradient__Spread = 0
	QGradient__ReflectSpread QGradient__Spread = 1
	QGradient__RepeatSpread  QGradient__Spread = 2
)

type QGradient__CoordinateMode int

const (
	QGradient__LogicalMode         QGradient__CoordinateMode = 0
	QGradient__StretchToDeviceMode QGradient__CoordinateMode = 1
	QGradient__ObjectBoundingMode  QGradient__CoordinateMode = 2
	QGradient__ObjectMode          QGradient__CoordinateMode = 3
)

type QGradient__InterpolationMode int

const (
	QGradient__ColorInterpolation     QGradient__InterpolationMode = 0
	QGradient__ComponentInterpolation QGradient__InterpolationMode = 1
)

type QGradient__Preset int

const (
	QGradient__WarmFlame        QGradient__Preset = 1
	QGradient__NightFade        QGradient__Preset = 2
	QGradient__SpringWarmth     QGradient__Preset = 3
	QGradient__JuicyPeach       QGradient__Preset = 4
	QGradient__YoungPassion     QGradient__Preset = 5
	QGradient__LadyLips         QGradient__Preset = 6
	QGradient__SunnyMorning     QGradient__Preset = 7
	QGradient__RainyAshville    QGradient__Preset = 8
	QGradient__FrozenDreams     QGradient__Preset = 9
	QGradient__WinterNeva       QGradient__Preset = 10
	QGradient__DustyGrass       QGradient__Preset = 11
	QGradient__TemptingAzure    QGradient__Preset = 12
	QGradient__HeavyRain        QGradient__Preset = 13
	QGradient__AmyCrisp         QGradient__Preset = 14
	QGradient__MeanFruit        QGradient__Preset = 15
	QGradient__DeepBlue         QGradient__Preset = 16
	QGradient__RipeMalinka      QGradient__Preset = 17
	QGradient__CloudyKnoxville  QGradient__Preset = 18
	QGradient__MalibuBeach      QGradient__Preset = 19
	QGradient__NewLife          QGradient__Preset = 20
	QGradient__TrueSunset       QGradient__Preset = 21
	QGradient__MorpheusDen      QGradient__Preset = 22
	QGradient__RareWind         QGradient__Preset = 23
	QGradient__NearMoon         QGradient__Preset = 24
	QGradient__WildApple        QGradient__Preset = 25
	QGradient__SaintPetersburg  QGradient__Preset = 26
	QGradient__PlumPlate        QGradient__Preset = 28
	QGradient__EverlastingSky   QGradient__Preset = 29
	QGradient__HappyFisher      QGradient__Preset = 30
	QGradient__Blessing         QGradient__Preset = 31
	QGradient__SharpeyeEagle    QGradient__Preset = 32
	QGradient__LadogaBottom     QGradient__Preset = 33
	QGradient__LemonGate        QGradient__Preset = 34
	QGradient__ItmeoBranding    QGradient__Preset = 35
	QGradient__ZeusMiracle      QGradient__Preset = 36
	QGradient__OldHat           QGradient__Preset = 37
	QGradient__StarWine         QGradient__Preset = 38
	QGradient__HappyAcid        QGradient__Preset = 41
	QGradient__AwesomePine      QGradient__Preset = 42
	QGradient__NewYork          QGradient__Preset = 43
	QGradient__ShyRainbow       QGradient__Preset = 44
	QGradient__MixedHopes       QGradient__Preset = 46
	QGradient__FlyHigh          QGradient__Preset = 47
	QGradient__StrongBliss      QGradient__Preset = 48
	QGradient__FreshMilk        QGradient__Preset = 49
	QGradient__SnowAgain        QGradient__Preset = 50
	QGradient__FebruaryInk      QGradient__Preset = 51
	QGradient__KindSteel        QGradient__Preset = 52
	QGradient__SoftGrass        QGradient__Preset = 53
	QGradient__GrownEarly       QGradient__Preset = 54
	QGradient__SharpBlues       QGradient__Preset = 55
	QGradient__ShadyWater       QGradient__Preset = 56
	QGradient__DirtyBeauty      QGradient__Preset = 57
	QGradient__GreatWhale       QGradient__Preset = 58
	QGradient__TeenNotebook     QGradient__Preset = 59
	QGradient__PoliteRumors     QGradient__Preset = 60
	QGradient__SweetPeriod      QGradient__Preset = 61
	QGradient__WideMatrix       QGradient__Preset = 62
	QGradient__SoftCherish      QGradient__Preset = 63
	QGradient__RedSalvation     QGradient__Preset = 64
	QGradient__BurningSpring    QGradient__Preset = 65
	QGradient__NightParty       QGradient__Preset = 66
	QGradient__SkyGlider        QGradient__Preset = 67
	QGradient__HeavenPeach      QGradient__Preset = 68
	QGradient__PurpleDivision   QGradient__Preset = 69
	QGradient__AquaSplash       QGradient__Preset = 70
	QGradient__SpikyNaga        QGradient__Preset = 72
	QGradient__LoveKiss         QGradient__Preset = 73
	QGradient__CleanMirror      QGradient__Preset = 75
	QGradient__PremiumDark      QGradient__Preset = 76
	QGradient__ColdEvening      QGradient__Preset = 77
	QGradient__CochitiLake      QGradient__Preset = 78
	QGradient__SummerGames      QGradient__Preset = 79
	QGradient__PassionateBed    QGradient__Preset = 80
	QGradient__MountainRock     QGradient__Preset = 81
	QGradient__DesertHump       QGradient__Preset = 82
	QGradient__JungleDay        QGradient__Preset = 83
	QGradient__PhoenixStart     QGradient__Preset = 84
	QGradient__OctoberSilence   QGradient__Preset = 85
	QGradient__FarawayRiver     QGradient__Preset = 86
	QGradient__AlchemistLab     QGradient__Preset = 87
	QGradient__OverSun          QGradient__Preset = 88
	QGradient__PremiumWhite     QGradient__Preset = 89
	QGradient__MarsParty        QGradient__Preset = 90
	QGradient__EternalConstance QGradient__Preset = 91
	QGradient__JapanBlush       QGradient__Preset = 92
	QGradient__SmilingRain      QGradient__Preset = 93
	QGradient__CloudyApple      QGradient__Preset = 94
	QGradient__BigMango         QGradient__Preset = 95
	QGradient__HealthyWater     QGradient__Preset = 96
	QGradient__AmourAmour       QGradient__Preset = 97
	QGradient__RiskyConcrete    QGradient__Preset = 98
	QGradient__StrongStick      QGradient__Preset = 99
	QGradient__ViciousStance    QGradient__Preset = 100
	QGradient__PaloAlto         QGradient__Preset = 101
	QGradient__HappyMemories    QGradient__Preset = 102
	QGradient__MidnightBloom    QGradient__Preset = 103
	QGradient__Crystalline      QGradient__Preset = 104
	QGradient__PartyBliss       QGradient__Preset = 106
	QGradient__ConfidentCloud   QGradient__Preset = 107
	QGradient__LeCocktail       QGradient__Preset = 108
	QGradient__RiverCity        QGradient__Preset = 109
	QGradient__FrozenBerry      QGradient__Preset = 110
	QGradient__ChildCare        QGradient__Preset = 112
	QGradient__FlyingLemon      QGradient__Preset = 113
	QGradient__NewRetrowave     QGradient__Preset = 114
	QGradient__HiddenJaguar     QGradient__Preset = 115
	QGradient__AboveTheSky      QGradient__Preset = 116
	QGradient__Nega             QGradient__Preset = 117
	QGradient__DenseWater       QGradient__Preset = 118
	QGradient__Seashore         QGradient__Preset = 120
	QGradient__MarbleWall       QGradient__Preset = 121
	QGradient__CheerfulCaramel  QGradient__Preset = 122
	QGradient__NightSky         QGradient__Preset = 123
	QGradient__MagicLake        QGradient__Preset = 124
	QGradient__YoungGrass       QGradient__Preset = 125
	QGradient__ColorfulPeach    QGradient__Preset = 126
	QGradient__GentleCare       QGradient__Preset = 127
	QGradient__PlumBath         QGradient__Preset = 128
	QGradient__HappyUnicorn     QGradient__Preset = 129
	QGradient__AfricanField     QGradient__Preset = 131
	QGradient__SolidStone       QGradient__Preset = 132
	QGradient__OrangeJuice      QGradient__Preset = 133
	QGradient__GlassWater       QGradient__Preset = 134
	QGradient__NorthMiracle     QGradient__Preset = 136
	QGradient__FruitBlend       QGradient__Preset = 137
	QGradient__MillenniumPine   QGradient__Preset = 138
	QGradient__HighFlight       QGradient__Preset = 139
	QGradient__MoleHall         QGradient__Preset = 140
	QGradient__SpaceShift       QGradient__Preset = 142
	QGradient__ForestInei       QGradient__Preset = 143
	QGradient__RoyalGarden      QGradient__Preset = 144
	QGradient__RichMetal        QGradient__Preset = 145
	QGradient__JuicyCake        QGradient__Preset = 146
	QGradient__SmartIndigo      QGradient__Preset = 147
	QGradient__SandStrike       QGradient__Preset = 148
	QGradient__NorseBeauty      QGradient__Preset = 149
	QGradient__AquaGuidance     QGradient__Preset = 150
	QGradient__SunVeggie        QGradient__Preset = 151
	QGradient__SeaLord          QGradient__Preset = 152
	QGradient__BlackSea         QGradient__Preset = 153
	QGradient__GrassShampoo     QGradient__Preset = 154
	QGradient__LandingAircraft  QGradient__Preset = 155
	QGradient__WitchDance       QGradient__Preset = 156
	QGradient__SleeplessNight   QGradient__Preset = 157
	QGradient__AngelCare        QGradient__Preset = 158
	QGradient__CrystalRiver     QGradient__Preset = 159
	QGradient__SoftLipstick     QGradient__Preset = 160
	QGradient__SaltMountain     QGradient__Preset = 161
	QGradient__PerfectWhite     QGradient__Preset = 162
	QGradient__FreshOasis       QGradient__Preset = 163
	QGradient__StrictNovember   QGradient__Preset = 164
	QGradient__MorningSalad     QGradient__Preset = 165
	QGradient__DeepRelief       QGradient__Preset = 166
	QGradient__SeaStrike        QGradient__Preset = 167
	QGradient__NightCall        QGradient__Preset = 168
	QGradient__SupremeSky       QGradient__Preset = 169
	QGradient__LightBlue        QGradient__Preset = 170
	QGradient__MindCrawl        QGradient__Preset = 171
	QGradient__LilyMeadow       QGradient__Preset = 172
	QGradient__SugarLollipop    QGradient__Preset = 173
	QGradient__SweetDessert     QGradient__Preset = 174
	QGradient__MagicRay         QGradient__Preset = 175
	QGradient__TeenParty        QGradient__Preset = 176
	QGradient__FrozenHeat       QGradient__Preset = 177
	QGradient__GagarinView      QGradient__Preset = 178
	QGradient__FabledSunset     QGradient__Preset = 179
	QGradient__PerfectBlue      QGradient__Preset = 180
	QGradient__NumPresets       QGradient__Preset = 181
)

type QBrush struct {
	h          uintptr
	isSubclass bool
}

// NewQBrush constructs a new QBrush object.
func NewQBrush() *QBrush {

	ret := newQBrush(QBrush_new())
	ret.isSubclass = true
	return ret
}

// NewQBrush2 constructs a new QBrush object.
func NewQBrush2(bs BrushStyle) *QBrush {

	ret := newQBrush(QBrush_new2((int)(bs)))
	ret.isSubclass = true
	return ret
}

// NewQBrush3 constructs a new QBrush object.
func NewQBrush3(color *QColor) *QBrush {

	ret := newQBrush(QBrush_new3(color.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQBrush4 constructs a new QBrush object.
func NewQBrush4(color GlobalColor) *QBrush {

	ret := newQBrush(QBrush_new4((int)(color)))
	ret.isSubclass = true
	return ret
}

// NewQBrush5 constructs a new QBrush object.
func NewQBrush5(color *QColor, pixmap *QPixmap) *QBrush {

	ret := newQBrush(QBrush_new5(color.cPointer(), pixmap.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQBrush6 constructs a new QBrush object.
func NewQBrush6(color GlobalColor, pixmap *QPixmap) *QBrush {

	ret := newQBrush(QBrush_new6((int)(color), pixmap.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQBrush7 constructs a new QBrush object.
func NewQBrush7(pixmap *QPixmap) *QBrush {

	ret := newQBrush(QBrush_new7(pixmap.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQBrush8 constructs a new QBrush object.
func NewQBrush8(image *QImage) *QBrush {

	ret := newQBrush(QBrush_new8(image.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQBrush9 constructs a new QBrush object.
func NewQBrush9(brush *QBrush) *QBrush {

	ret := newQBrush(QBrush_new9(brush.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQBrush10 constructs a new QBrush object.
func NewQBrush10(gradient *QGradient) *QBrush {

	ret := newQBrush(QBrush_new10(gradient.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQBrush11 constructs a new QBrush object.
func NewQBrush11(color *QColor, bs BrushStyle) *QBrush {

	ret := newQBrush(QBrush_new11(color.cPointer(), (int)(bs)))
	ret.isSubclass = true
	return ret
}

// NewQBrush12 constructs a new QBrush object.
func NewQBrush12(color GlobalColor, bs BrushStyle) *QBrush {

	ret := newQBrush(QBrush_new12((int)(color), (int)(bs)))
	ret.isSubclass = true
	return ret
}

func (this *QBrush) OperatorAssign(brush *QBrush) {
	QBrush_OperatorAssign(this.h, brush.cPointer())
}

func (this *QBrush) Swap(other *QBrush) {
	QBrush_Swap(this.h, other.cPointer())
}

func (this *QBrush) OperatorAssignWithStyle(style BrushStyle) {
	QBrush_OperatorAssignWithStyle(this.h, (int)(style))
}

func (this *QBrush) OperatorAssignWithColor(color QColor) {
	QBrush_OperatorAssignWithColor(this.h, color.cPointer())
}

func (this *QBrush) OperatorAssign2(color GlobalColor) {
	QBrush_OperatorAssign2(this.h, (int)(color))
}

func (this *QBrush) Style() BrushStyle {
	return (BrushStyle)(QBrush_Style(this.h))
}

func (this *QBrush) SetStyle(style BrushStyle) {
	QBrush_SetStyle(this.h, (int)(style))
}

func (this *QBrush) Transform() *QTransform {
	_goptr := newQTransform(QBrush_Transform(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QBrush) SetTransform(transform *QTransform) {
	QBrush_SetTransform(this.h, transform.cPointer())
}

func (this *QBrush) Texture() *QPixmap {
	_goptr := newQPixmap(QBrush_Texture(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QBrush) SetTexture(pixmap *QPixmap) {
	QBrush_SetTexture(this.h, pixmap.cPointer())
}

func (this *QBrush) TextureImage() *QImage {
	_goptr := newQImage(QBrush_TextureImage(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QBrush) SetTextureImage(image *QImage) {
	QBrush_SetTextureImage(this.h, image.cPointer())
}

func (this *QBrush) Color() *QColor {
	return newQColor(QBrush_Color(this.h))
}

func (this *QBrush) SetColor(color *QColor) {
	QBrush_SetColor(this.h, color.cPointer())
}

func (this *QBrush) SetColorWithColor(color GlobalColor) {
	QBrush_SetColorWithColor(this.h, (int)(color))
}

func (this *QBrush) Gradient() *QGradient {
	return newQGradient(QBrush_Gradient(this.h))
}

func (this *QBrush) IsOpaque() bool {
	return (bool)(QBrush_IsOpaque(this.h))
}

func (this *QBrush) OperatorEqual(b *QBrush) bool {
	return (bool)(QBrush_OperatorEqual(this.h, b.cPointer()))
}

func (this *QBrush) OperatorNotEqual(b *QBrush) bool {
	return (bool)(QBrush_OperatorNotEqual(this.h, b.cPointer()))
}

func (this *QBrush) IsDetached() bool {
	return (bool)(QBrush_IsDetached(this.h))
}

func (this *QBrush) DataPtr() *DataPtr {
	xxxxxxxxx
}

type QBrushData struct {
	h          uintptr
	isSubclass bool
}

// NewQBrushData constructs a new QBrushData object.
func NewQBrushData(param1 *QBrushData) *QBrushData {

	ret := newQBrushData(QBrushData_new(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QBrushData) OperatorAssign(param1 *QBrushData) {
	QBrushData_OperatorAssign(this.h, param1.cPointer())
}

type QGradient struct {
	h          uintptr
	isSubclass bool
}

// NewQGradient constructs a new QGradient object.
func NewQGradient() *QGradient {

	ret := newQGradient(QGradient_new())
	ret.isSubclass = true
	return ret
}

// NewQGradient2 constructs a new QGradient object.
func NewQGradient2(param1 Preset) *QGradient {

	ret := newQGradient(QGradient_new2(param1))
	ret.isSubclass = true
	return ret
}

// NewQGradient3 constructs a new QGradient object.
func NewQGradient3(param1 *QGradient) *QGradient {

	ret := newQGradient(QGradient_new3(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QGradient) Type() Type {
	xxxxxxxxx
}

func (this *QGradient) SetSpread(spread Spread) {
	QGradient_SetSpread(this.h, spread)
}

func (this *QGradient) Spread() Spread {
	xxxxxxxxx
}

func (this *QGradient) SetColorAt(pos float64, color *QColor) {
	QGradient_SetColorAt(this.h, (double)(pos), color.cPointer())
}

func (this *QGradient) SetStops(stops []struct {
	First  float64
	Second QColor
}) {
	stops_CArray := (*[0xffff]struct_miqt_map)(malloc(size_t(8 * len(stops))))
	defer free(unsafe.Pointer(stops_CArray))
	for i := range stops {
		stops_i_First_CArray := (*[0xffff]double)(malloc(size_t(8)))
		defer free(unsafe.Pointer(stops_i_First_CArray))
		stops_i_Second_CArray := (*[0xffff]*QColor)(malloc(size_t(8)))
		defer free(unsafe.Pointer(stops_i_Second_CArray))
		stops_i_First_CArray[0] = (double)(stops[i].First)
		stops_i_Second_CArray[0] = stops[i].Second.cPointer()
		stops_i_pair := struct_miqt_map{
			len:    1,
			keys:   unsafe.Pointer(stops_i_First_CArray),
			values: unsafe.Pointer(stops_i_Second_CArray),
		}
		stops_CArray[i] = stops_i_pair
	}
	stops_ma := struct_miqt_array{len: size_t(len(stops)), data: unsafe.Pointer(stops_CArray)}
	QGradient_SetStops(this.h, stops_ma)
}

func (this *QGradient) Stops() []struct {
	First  float64
	Second QColor
} {
	var _ma struct_miqt_array = QGradient_Stops(this.h)
	_ret := make([]struct {
		First  float64
		Second QColor
	}, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_map)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_mm struct_miqt_map = _outCast[i]
		_lv_First_CArray := (*[0xffff]double)(unsafe.Pointer(_lv_mm.keys))
		_lv_Second_CArray := (*[0xffff]*QColor)(unsafe.Pointer(_lv_mm.values))
		_lv_entry_First := (float64)(_lv_First_CArray[0])

		_lv_second_goptr := newQColor(_lv_Second_CArray[0])
		_lv_second_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_lv_entry_Second := *_lv_second_goptr

		_ret[i] = struct {
			First  float64
			Second QColor
		}{First: _lv_entry_First, Second: _lv_entry_Second}
	}
	return _ret
}

func (this *QGradient) CoordinateMode() CoordinateMode {
	xxxxxxxxx
}

func (this *QGradient) SetCoordinateMode(mode CoordinateMode) {
	QGradient_SetCoordinateMode(this.h, mode)
}

func (this *QGradient) InterpolationMode() InterpolationMode {
	xxxxxxxxx
}

func (this *QGradient) SetInterpolationMode(mode InterpolationMode) {
	QGradient_SetInterpolationMode(this.h, mode)
}

func (this *QGradient) OperatorEqual(gradient *QGradient) bool {
	return (bool)(QGradient_OperatorEqual(this.h, gradient.cPointer()))
}

func (this *QGradient) OperatorNotEqual(other *QGradient) bool {
	return (bool)(QGradient_OperatorNotEqual(this.h, other.cPointer()))
}

type QLinearGradient struct {
	h          uintptr
	isSubclass bool
}

// NewQLinearGradient constructs a new QLinearGradient object.
func NewQLinearGradient() *QLinearGradient {

	ret := newQLinearGradient(QLinearGradient_new())
	ret.isSubclass = true
	return ret
}

// NewQLinearGradient2 constructs a new QLinearGradient object.
func NewQLinearGradient2(start *QPointF, finalStop *QPointF) *QLinearGradient {

	ret := newQLinearGradient(QLinearGradient_new2(start.cPointer(), finalStop.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQLinearGradient3 constructs a new QLinearGradient object.
func NewQLinearGradient3(xStart float64, yStart float64, xFinalStop float64, yFinalStop float64) *QLinearGradient {

	ret := newQLinearGradient(QLinearGradient_new3((double)(xStart), (double)(yStart), (double)(xFinalStop), (double)(yFinalStop)))
	ret.isSubclass = true
	return ret
}

// NewQLinearGradient4 constructs a new QLinearGradient object.
func NewQLinearGradient4(param1 *QLinearGradient) *QLinearGradient {

	ret := newQLinearGradient(QLinearGradient_new4(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QLinearGradient) Start() *QPointF {
	_goptr := newQPointF(QLinearGradient_Start(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLinearGradient) SetStart(start *QPointF) {
	QLinearGradient_SetStart(this.h, start.cPointer())
}

func (this *QLinearGradient) SetStart2(x float64, y float64) {
	QLinearGradient_SetStart2(this.h, (double)(x), (double)(y))
}

func (this *QLinearGradient) FinalStop() *QPointF {
	_goptr := newQPointF(QLinearGradient_FinalStop(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLinearGradient) SetFinalStop(stop *QPointF) {
	QLinearGradient_SetFinalStop(this.h, stop.cPointer())
}

func (this *QLinearGradient) SetFinalStop2(x float64, y float64) {
	QLinearGradient_SetFinalStop2(this.h, (double)(x), (double)(y))
}

type QRadialGradient struct {
	h          uintptr
	isSubclass bool
}

// NewQRadialGradient constructs a new QRadialGradient object.
func NewQRadialGradient() *QRadialGradient {

	ret := newQRadialGradient(QRadialGradient_new())
	ret.isSubclass = true
	return ret
}

// NewQRadialGradient2 constructs a new QRadialGradient object.
func NewQRadialGradient2(center *QPointF, radius float64, focalPoint *QPointF) *QRadialGradient {

	ret := newQRadialGradient(QRadialGradient_new2(center.cPointer(), (double)(radius), focalPoint.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQRadialGradient3 constructs a new QRadialGradient object.
func NewQRadialGradient3(cx float64, cy float64, radius float64, fx float64, fy float64) *QRadialGradient {

	ret := newQRadialGradient(QRadialGradient_new3((double)(cx), (double)(cy), (double)(radius), (double)(fx), (double)(fy)))
	ret.isSubclass = true
	return ret
}

// NewQRadialGradient4 constructs a new QRadialGradient object.
func NewQRadialGradient4(center *QPointF, radius float64) *QRadialGradient {

	ret := newQRadialGradient(QRadialGradient_new4(center.cPointer(), (double)(radius)))
	ret.isSubclass = true
	return ret
}

// NewQRadialGradient5 constructs a new QRadialGradient object.
func NewQRadialGradient5(cx float64, cy float64, radius float64) *QRadialGradient {

	ret := newQRadialGradient(QRadialGradient_new5((double)(cx), (double)(cy), (double)(radius)))
	ret.isSubclass = true
	return ret
}

// NewQRadialGradient6 constructs a new QRadialGradient object.
func NewQRadialGradient6(center *QPointF, centerRadius float64, focalPoint *QPointF, focalRadius float64) *QRadialGradient {

	ret := newQRadialGradient(QRadialGradient_new6(center.cPointer(), (double)(centerRadius), focalPoint.cPointer(), (double)(focalRadius)))
	ret.isSubclass = true
	return ret
}

// NewQRadialGradient7 constructs a new QRadialGradient object.
func NewQRadialGradient7(cx float64, cy float64, centerRadius float64, fx float64, fy float64, focalRadius float64) *QRadialGradient {

	ret := newQRadialGradient(QRadialGradient_new7((double)(cx), (double)(cy), (double)(centerRadius), (double)(fx), (double)(fy), (double)(focalRadius)))
	ret.isSubclass = true
	return ret
}

// NewQRadialGradient8 constructs a new QRadialGradient object.
func NewQRadialGradient8(param1 *QRadialGradient) *QRadialGradient {

	ret := newQRadialGradient(QRadialGradient_new8(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QRadialGradient) Center() *QPointF {
	_goptr := newQPointF(QRadialGradient_Center(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRadialGradient) SetCenter(center *QPointF) {
	QRadialGradient_SetCenter(this.h, center.cPointer())
}

func (this *QRadialGradient) SetCenter2(x float64, y float64) {
	QRadialGradient_SetCenter2(this.h, (double)(x), (double)(y))
}

func (this *QRadialGradient) FocalPoint() *QPointF {
	_goptr := newQPointF(QRadialGradient_FocalPoint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRadialGradient) SetFocalPoint(focalPoint *QPointF) {
	QRadialGradient_SetFocalPoint(this.h, focalPoint.cPointer())
}

func (this *QRadialGradient) SetFocalPoint2(x float64, y float64) {
	QRadialGradient_SetFocalPoint2(this.h, (double)(x), (double)(y))
}

func (this *QRadialGradient) Radius() float64 {
	return (float64)(QRadialGradient_Radius(this.h))
}

func (this *QRadialGradient) SetRadius(radius float64) {
	QRadialGradient_SetRadius(this.h, (double)(radius))
}

func (this *QRadialGradient) CenterRadius() float64 {
	return (float64)(QRadialGradient_CenterRadius(this.h))
}

func (this *QRadialGradient) SetCenterRadius(radius float64) {
	QRadialGradient_SetCenterRadius(this.h, (double)(radius))
}

func (this *QRadialGradient) FocalRadius() float64 {
	return (float64)(QRadialGradient_FocalRadius(this.h))
}

func (this *QRadialGradient) SetFocalRadius(radius float64) {
	QRadialGradient_SetFocalRadius(this.h, (double)(radius))
}

type QConicalGradient struct {
	h          uintptr
	isSubclass bool
}

// NewQConicalGradient constructs a new QConicalGradient object.
func NewQConicalGradient() *QConicalGradient {

	ret := newQConicalGradient(QConicalGradient_new())
	ret.isSubclass = true
	return ret
}

// NewQConicalGradient2 constructs a new QConicalGradient object.
func NewQConicalGradient2(center *QPointF, startAngle float64) *QConicalGradient {

	ret := newQConicalGradient(QConicalGradient_new2(center.cPointer(), (double)(startAngle)))
	ret.isSubclass = true
	return ret
}

// NewQConicalGradient3 constructs a new QConicalGradient object.
func NewQConicalGradient3(cx float64, cy float64, startAngle float64) *QConicalGradient {

	ret := newQConicalGradient(QConicalGradient_new3((double)(cx), (double)(cy), (double)(startAngle)))
	ret.isSubclass = true
	return ret
}

// NewQConicalGradient4 constructs a new QConicalGradient object.
func NewQConicalGradient4(param1 *QConicalGradient) *QConicalGradient {

	ret := newQConicalGradient(QConicalGradient_new4(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QConicalGradient) Center() *QPointF {
	_goptr := newQPointF(QConicalGradient_Center(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QConicalGradient) SetCenter(center *QPointF) {
	QConicalGradient_SetCenter(this.h, center.cPointer())
}

func (this *QConicalGradient) SetCenter2(x float64, y float64) {
	QConicalGradient_SetCenter2(this.h, (double)(x), (double)(y))
}

func (this *QConicalGradient) Angle() float64 {
	return (float64)(QConicalGradient_Angle(this.h))
}

func (this *QConicalGradient) SetAngle(angle float64) {
	QConicalGradient_SetAngle(this.h, (double)(angle))
}

type QGradient__QGradientData struct {
	h          uintptr
	isSubclass bool
}

// NewQGradient__QGradientData constructs a new QGradient::QGradientData object.
func NewQGradient__QGradientData(param1 *QGradientData) *QGradient__QGradientData {

	ret := newQGradient__QGradientData(QGradient__QGradientData_new(param1))
	ret.isSubclass = true
	return ret
}
