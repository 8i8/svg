package set

type svgType uint

const (
	/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	 *  element
	 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */
	elemA svgType = iota
	elemAltGlyph
	elemAltGlyphDef
	elemAltGlyphItem
	elemAnimate
	elemAnimateColor
	elemAnimateMotion
	elemAnimateTransform
	elemCircle
	elemClipPath
	elemColorProfile
	elemCursor
	elemDefs
	elemDesc
	elemEllipse
	elemFeBlend
	elemFeColorMatrix
	elemFeComponentTransfer
	elemFeComposite
	elemFeConvolveMatrix
	elemFeDiffuseLighting
	elemFeDisplacementMap
	elemFeDistantLight
	elemFeFlood
	elemFeFuncA
	elemFeFuncB
	elemFeFuncG
	elemFeFuncR
	elemFeGaussianBlur
	elemFeImage
	elemFeMerge
	elemFeMergeNode
	elemFeMorphology
	elemFeOffset
	elemFePointLight
	elemFeSpecularLighting
	elemFeSpotLight
	elemFeTile
	elemFeTurbulence
	elemFilter
	elemFont
	elemFontFace
	elemFontFaceFormat
	elemFontFaceName
	elemFontFaceSrc
	elemFontFaceUri
	elemForeignObject
	elemG
	elemGlyph
	elemGlyphRef
	elemHkern
	elemImage
	elemLine
	elemLinearGradient
	elemMarker
	elemMask
	elemMetadata
	elemMissingGlyph
	elemMpath
	elemPath
	elemPattern
	elemPolygon
	elemPolyline
	elemRadialGradient
	elemRect
	elemScript
	elemSet
	elemStop
	elemStyle
	elemSvg
	elemSwitch
	elemSymbol
	elemText
	elemTextPath
	elemTitle
	elemTref
	elemTspan
	elemUse
	elemView
	elemVkern

	/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	 *  attributes
	 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */
	attrAccentHeight
	attrAccumulate
	attrAdditive
	attrAlignmentBaseline
	attrAlphabetic
	attrAmplitude
	attrArabicForm
	attrAscent
	attrAttributeName
	attrAttributeType
	attrAzimuth
	attrBaseFrequency
	attrBaselineShift
	attrBaseProfile
	attrBbox
	attrBegin
	attrBias
	attrBy
	attrCalcMode
	attrCapHeight
	attrClass
	attrClip
	attrClipPathUnits
	attrClipPath
	attrClipRule
	attrColor
	attrColorInterpolation
	attrColorInterpolationFilters
	attrColorProfile
	attrColorRendering
	attrContentScriptType
	attrContentStyleType
	attrCrossorigin
	attrCursor
	attrCx
	attrCy
	attrD
	attrDecelerate
	attrDescent
	attrDiffuseConstant
	attrDirection
	attrDisplay
	attrDivisor
	attrDominantBaseline
	attrDur
	attrDx
	attrDy
	attrEdgeMode
	attrElevation
	attrEnableBackground
	attrEnd
	attrExponent
	attrFill
	attrFillOpacity
	attrFillRule
	attrFilter
	attrFilterRes
	attrFilterUnits
	attrFloodColor
	attrFloodOpacity
	attrFontFamily
	attrFontSize
	attrFontSizeAdjust
	attrFontStretch
	attrFontStyle
	attrFontVariant
	attrFontWeight
	attrFormat
	attrFrom
	attrFr
	attrFx
	attrFy
	attrG1
	attrG2
	attrGlyphName
	attrGlyphOrientationHorizontal
	attrGlyphOrientationVertical
	attrGlyphRef
	attrGradientTransform
	attrGradientUnits
	attrHanging
	attrHeight
	attrHref
	attrHreflang
	attrHorizAdvX
	attrHorizOriginX
	attrHorizOriginY
	attrId
	attrIdeographic
	attrImageRendering
	attrIn
	attrIn2
	attrIntercept
	attrK
	attrK1
	attrK2
	attrK3
	attrK4
	attrKernelMatrix
	attrKernelUnitLength
	attrKerning
	attrKeyPoints
	attrKeySplines
	attrKeyTimes
	attrLang
	attrLengthAdjust
	attrLetterSpacing
	attrLightingColor
	attrLimitingConeAngle
	attrLocal
	attrMarkerEnd
	attrMarkerMid
	attrMarkerStart
	attrMarkerHeight
	attrMarkerUnits
	attrMarkerWidth
	attrMask
	attrMaskContentUnits
	attrMaskUnits
	attrMathematical
	attrMax
	attrMedia
	attrMethod
	attrMin
	attrMode
	attrName
	attrNumOctaves
	attrOffset
	attrOnunload
	attrOnabort
	attrOnerror
	attrOnresize
	attrOnscroll
	attrOnzoom
	attrOpacity
	attrOperator
	attrOrder
	attrOrient
	attrOrientation
	attrOrigin
	attrOverflow
	attrOverlinePosition
	attrOverlineThickness
	attrPanose1
	attrPaintOrder
	attrPath
	attrPathLength
	attrPatternContentUnits
	attrPatternTransform
	attrPatternUnits
	attrPing
	attrPointerEvents
	attrPoints
	attrPointsAtX
	attrPointsAtY
	attrPointsAtZ
	attrPreserveAlpha
	attrPreserveAspectRatio
	attrPrimitiveUnits
	attrR
	attrRadius
	attrReferrerPolicy
	attrRefX
	attrRefY
	attrRel
	attrRenderingIntent
	attrRepeatCount
	attrRepeatDur
	attrRequiredExtensions
	attrRequiredFeatures
	attrRestart
	attrResult
	attrRotate
	attrRx
	attrRy
	attrScale
	attrSeed
	attrShapeRendering
	attrSlope
	attrSpacing
	attrSpecularConstant
	attrSpecularExponent
	attrSpeed
	attrSpreadMethod
	attrStartOffset
	attrStdDeviation
	attrStemh
	attrStemv
	attrStitchTiles
	attrStopColor
	attrStopOpacity
	attrStrikethroughPosition
	attrStrikethroughThickness
	attrString
	attrStroke
	attrStrokeDasharray
	attrStrokeDashoffset
	attrStrokeLinecap
	attrStrokeLinejoin
	attrStrokeMiterlimit
	attrStrokeOpacity
	attrStrokeWidth
	attrStyle
	attrSurfaceScale
	attrSystemLanguage
	attrTabindex
	attrTableValues
	attrTarget
	attrTargetX
	attrTargetY
	attrTextAnchor
	attrTextDecoration
	attrTextRendering
	attrTextLength
	attrTitle
	attrTo
	attrTransform
	attrTransformOrigin
	attrType
	attrU1
	attrU2
	attrUnderlinePosition
	attrUnderlineThickness
	attrUnicode
	attrUnicodeBidi
	attrUnicodeRange
	attrUnitsPerEm
	attrVAlphabetic
	attrVHanging
	attrVIdeographic
	attrVMathematical
	attrValues
	attrVectorEffect
	attrVersion
	attrVertAdvY
	attrVertOriginX
	attrVertOriginY
	attrViewBox
	attrViewTarget
	attrVisibility
	attrWidth
	attrWidths
	attrWordSpacing
	attrWritingMode
	attrX
	attrXHeight
	attrX1
	attrX2
	attrXChannelSelector
	attrXlinkActuate
	attrXlinkArcrole
	attrXlinkHref
	attrXlinkRole
	attrXlinkShow
	attrXlinkTitle
	attrXlinkType
	attrXmlBase
	attrXmlLang
	attrXmlSpace
	attrY
	attrY1
	attrY2
	attrYChannelSelector
	attrZ
	attrZoomAndPan
	// Graphical attributes
	attrOnfocusin
	attrOnfocusout
	attrOnactivate
	attrOnclick
	attrOnmousedown
	attrOnmouseup
	attrOnmouseover
	attrOnmousemove
	attrOnmouseout
	// Animation
	attrOnload
	attrOnbegin
	attrOnend
	attrOnrepeat
	attributeType
	attributeName
	//
	attrExternalResourcesRequired

	/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	 *  catagories
	 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */
	cataContainerElement
	cataAnimationElement
	cataBasicShapeElement
	cataGraphicsElement
	cataShapeElement
	cataStructuralElement
	cataDescriptiveElement
	cataTextContentElement
	cataTextContentChildElement
	cataFilterPrimitiveElement
	cataLightSourceElement
	cataGraphicsReferencingElement
	cataGradientElement

	/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	 *  DOMinterface
	 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */
	domSVGAElement
	domSVGAnimateElement
	domSVGAnimateColorElement
	domSVGAnimateMotionElement
	domSVGCircleElement
	domSVGAnimateTransformElement
	domSVGClipPathElement
	domSVGDefsElement
	domSVGDescElement
	domSVGAltGlyphElement
	domSVGAltGlyphDefElement
	domSVGAltGlyphItemElement
	domSVGColorProfileElement
	domSVGCursorElement
	domSVGEllipseElement
	domSVGFEBlendElement
	domSVGFEColorMatrixElement
	domSVGFEComponentTransferElement
	domSVGFECompositeElement
	domSVGFEConvolveMatrixElement
	domSVGFEDiffuseLightingElement
	domSVGFEDisplacementMapElement
	domSVGFEDistantLightElement
	domSVGFEFloodElement
	domSVGFEFuncAElement
	domSVGFEFuncBElement
	domSVGFEFuncGElement
	domSVGFEFuncRElement
	domSVGFEGaussianBlurElement
	domSVGFEImageElement
	domSVGFEMergeElement
	domSVGFEMergeNodeElement
	domSVGFEMorphologyElement
	domSVGFEOffsetElement
	domSVGFEPointLightElement
	domSVGFESpecularLightingElement
	domSVGFESpotLightElement
	domSVGFETileElement
	domSVGFETurbulenceElement
	domSVGFilterElement
	domSVGFontElement
	domSVGFontFaceFormatElement
	domSVGFontFaceElement
	domSVGFontFaceNameElement
	domSVGFontFaceUriElement
	domSVGFontFaceSrcElement
	domSVGForeignObjectElement
	domSVGGElement
	domSVGGlyphElement
	domSVGGlyphRefElement
	domSVGHKernElement
	domSVGImageElement
	domSVGLineElement
	domSVGLinearGradientElement
	domSVGMarkerElement
	domSVGMaskElement
	domSVGMetadataElement
	domSVGMissingGlyphElement
	domSVGMPathElement
	domSVGPathElement
	domSVGPatternElement
	domSVGPolygonElement
	domSVGPolylineElement
	domSVGRadialGradientElement
	domSVGRectElement
	domSVGScriptElement
	domSVGSetElement
	domSVGStopElement
	domSVGStyleElement
	domSVGSVGElement
	domSVGSwitchElement
	domSVGSymbolElement
	domSVGTextElement
	domSVGTextPathElement
	domSVGTitleElement
	domSVGTRefElement
	domSVGTSpanElement
	domSVGUseElement
	domSVGViewElement
)

var elements = make(map[svgType]group)

type group struct {
	*bitVector
	contract []func(group) error
}

// cm 	centimeters
// mm 	millimeters
// in 	inches (1in = 96px = 2.54cm)
// px  	pixels (1px = 1/96th of 1in)
// pt 	points (1pt = 1/72 of 1in)
// pc 	picas (1pc = 12 pt)

// Relative length units specify a length relative to another length
// property. Relative length units scale better between different
// rendering medium.

// em 	Relative to the font-size of the element (2em means 2 times the size of the current font)
// ex 	Relative to the x-height of the current font (rarely used)
// ch 	Relative to the width of the "0" (zero)
// rem 	Relative to font-size of the root element
// vw 	Relative to 1% of the width of the viewport*
// vh 	Relative to 1% of the height of the viewport*
// vmin 	Relative to 1% of viewport's* smaller dimension
// vmax 	Relative to 1% of viewport's* larger dimension
// % 	Relative to the parent element

var UnitCheck = map[string]bool{
	"cm":   true,
	"mm":   true,
	"in":   true,
	"px":   true,
	"pt":   true,
	"pc":   true,
	"em":   true,
	"ex":   true,
	"ch":   true,
	"rem":  true,
	"vw":   true,
	"vh":   true,
	"vmin": true,
	"vmax": true,
	"%":    true,
}

var ColourCheck = map[string]string{
	"aliceblue":            "rgb(240, 248, 255)",
	"antiquewhite":         "rgb(250, 235, 215)",
	"aqua":                 "rgb( 0, 255, 255)",
	"aquamarine":           "rgb(127, 255, 212)",
	"azure":                "rgb(240, 255, 255)",
	"beige":                "rgb(245, 245, 220)",
	"bisque":               "rgb(255, 228, 196)",
	"black":                "rgb( 0, 0, 0)",
	"blanchedalmond":       "rgb(255, 235, 205)",
	"blue":                 "rgb( 0, 0, 255)",
	"blueviolet":           "rgb(138, 43, 226)",
	"brown":                "rgb(165, 42, 42)",
	"burlywood":            "rgb(222, 184, 135)",
	"cadetblue":            "rgb( 95, 158, 160)",
	"chartreuse":           "rgb(127, 255, 0)",
	"chocolate":            "rgb(210, 105, 30)",
	"coral":                "rgb(255, 127, 80)",
	"cornflowerblue":       "rgb(100, 149, 237)",
	"cornsilk":             "rgb(255, 248, 220)",
	"crimson":              "rgb(220, 20, 60)",
	"cyan":                 "rgb( 0, 255, 255)",
	"darkblue":             "rgb( 0, 0, 139)",
	"darkcyan":             "rgb( 0, 139, 139)",
	"darkgoldenrod":        "rgb(184, 134, 11)",
	"darkgray":             "rgb(169, 169, 169)",
	"darkgreen":            "rgb( 0, 100, 0)",
	"darkgrey":             "rgb(169, 169, 169)",
	"darkkhaki":            "rgb(189, 183, 107)",
	"darkmagenta":          "rgb(139, 0, 139)",
	"darkolivegreen":       "rgb( 85, 107, 47)",
	"darkorange":           "rgb(255, 140, 0)",
	"darkorchid":           "rgb(153, 50, 204)",
	"darkred":              "rgb(139, 0, 0)",
	"darksalmon":           "rgb(233, 150, 122)",
	"darkseagreen":         "rgb(143, 188, 143)",
	"darkslateblue":        "rgb( 72, 61, 139)",
	"darkslategray":        "rgb( 47, 79, 79)",
	"darkturquoise":        "rgb( 0, 206, 209)",
	"darkviolet":           "rgb(148, 0, 211)",
	"deeppink":             "rgb(255, 20, 147)",
	"deepskyblue":          "rgb( 0, 191, 255)",
	"dimgrey":              "rgb(105, 105, 105)",
	"dodgerblue":           "rgb( 30, 144, 255)",
	"firebrick":            "rgb(178, 34, 34)",
	"floralwhite":          "rgb(255, 250, 240)",
	"forestgreen":          "rgb( 34, 139, 34)",
	"fuchsia":              "rgb(255, 0, 255)",
	"gainsboro":            "rgb(220, 220, 220)",
	"ghostwhite":           "rgb(248, 248, 255)",
	"gold":                 "rgb(255, 215, 0)",
	"goldenrod":            "rgb(218, 165, 32)",
	"grey":                 "rgb(128, 128, 128)",
	"green":                "rgb( 0, 128, 0)",
	"greenyellow":          "rgb(173, 255, 47)",
	"honeydew":             "rgb(240, 255, 240)",
	"hotpink":              "rgb(255, 105, 180)",
	"indianred":            "rgb(205, 92, 92)",
	"indigo":               "rgb( 75, 0, 130)",
	"ivory":                "rgb(255, 255, 240)",
	"khaki":                "rgb(240, 230, 140)",
	"lavender":             "rgb(230, 230, 250)",
	"lavenderblush":        "rgb(255, 240, 245)",
	"lawngreen":            "rgb(124, 252, 0)",
	"lemonchiffon":         "rgb(255, 250, 205)",
	"lightblue":            "rgb(173, 216, 230)",
	"lightcoral":           "rgb(240, 128, 128)",
	"lightcyan":            "rgb(224, 255, 255)",
	"lightgoldenrodyellow": "rgb(250, 250, 210)",
	"lightgreen":           "rgb(144, 238, 144)",
	"lightgrey":            "rgb(211, 211, 211)",
	"lightpink":            "rgb(255, 182, 193)",
	"lightsalmon":          "rgb(255, 160, 122)",
	"lightseagreen":        "rgb( 32, 178, 170)",
	"lightskyblue":         "rgb(135, 206, 250)",
	"lightslategray":       "rgb(119, 136, 153)",
	"lightslategrey":       "rgb(119, 136, 153)",
	"lightsteelblue":       "rgb(176, 196, 222)",
	"lightyellow":          "rgb(255, 255, 224)",
	"lime":                 "rgb( 0, 255, 0)",
	"limegreen":            "rgb( 50, 205, 50)",
	"linen":                "rgb(250, 240, 230)",
	"magenta":              "rgb(255, 0, 255)",
	"maroon":               "rgb(128, 0, 0)",
	"mediumaquamarine":     "rgb(102, 205, 170)",
	"mediumblue":           "rgb( 0, 0, 205)",
	"mediumorchid":         "rgb(186, 85, 211)",
	"mediumpurple":         "rgb(147, 112, 219)",
	"mediumseagreen":       "rgb( 60, 179, 113)",
	"mediumslateblue":      "rgb(123, 104, 238)",
	"mediumspringgreen":    "rgb( 0, 250, 154)",
	"mediumturquoise":      "rgb( 72, 209, 204)",
	"mediumvioletred":      "rgb(199, 21, 133)",
	"midnightblue":         "rgb( 25, 25, 112)",
	"mintcream":            "rgb(245, 255, 250)",
	"mistyrose":            "rgb(255, 228, 225)",
	"moccasin":             "rgb(255, 228, 181)",
	"navajowhite":          "rgb(255, 222, 173)",
	"navy":                 "rgb( 0, 0, 128)",
	"oldlace":              "rgb(253, 245, 230)",
	"olive":                "rgb(128, 128, 0)",
	"olivedrab":            "rgb(107, 142, 35)",
	"orange":               "rgb(255, 165, 0)",
	"orangered":            "rgb(255, 69, 0)",
	"orchid":               "rgb(218, 112, 214)",
	"palegoldenrod":        "rgb(238, 232, 170)",
	"palegreen":            "rgb(152, 251, 152)",
	"paleturquoise":        "rgb(175, 238, 238)",
	"palevioletred":        "rgb(219, 112, 147)",
	"papayawhip":           "rgb(255, 239, 213)",
	"peachpuff":            "rgb(255, 218, 185)",
	"peru":                 "rgb(205, 133, 63)",
	"pink":                 "rgb(255, 192, 203)",
	"plum":                 "rgb(221, 160, 221)",
	"powderblue":           "rgb(176, 224, 230)",
	"purple":               "rgb(128, 0, 128)",
	"red":                  "rgb(255, 0, 0)",
	"rosybrown":            "rgb(188, 143, 143)",
	"royalblue":            "rgb( 65, 105, 225)",
	"saddlebrown":          "rgb(139, 69, 19)",
	"salmon":               "rgb(250, 128, 114)",
	"sandybrown":           "rgb(244, 164, 96)",
	"seagreen":             "rgb( 46, 139, 87)",
	"seashell":             "rgb(255, 245, 238)",
	"sienna":               "rgb(160, 82, 45)",
	"silver":               "rgb(192, 192, 192)",
	"skyblue":              "rgb(135, 206, 235)",
	"slateblue":            "rgb(106, 90, 205)",
	"slategrey":            "rgb(112, 128, 144)",
	"snow":                 "rgb(255, 250, 250)",
	"springgreen":          "rgb( 0, 255, 127)",
	"steelblue":            "rgb( 70, 130, 180)",
	"tan":                  "rgb(210, 180, 140)",
	"teal":                 "rgb( 0, 128, 128)",
	"thistle":              "rgb(216, 191, 216)",
	"tomato":               "rgb(255, 99, 71)",
	"turquoise":            "rgb( 64, 224, 208)",
	"violet":               "rgb(238, 130, 238)",
	"wheat":                "rgb(245, 222, 179)",
	"white":                "rgb(255, 255, 255)",
	"whitesmoke":           "rgb(245, 245, 245)",
	"yellow":               "rgb(255, 255, 0)",
	"yellowgreen":          "rgb(154, 205, 50)",
}
