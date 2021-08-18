// https://dev.w3.org/SVG/profiles/1.1F2/publish/
package set

// svg elemtent has baseProfile in attributes twice
// svg attribuites not in full list
// onunload onabort onerror onresize onscroll onzoom
// style title attribute not there

import "errors"

var (
	// Elements
	animationElements = []svgType{
		elemAnimate, elemAnimateColor, elemAnimateMotion,
		elemAnimateTransform, elemSet,
	}
	descriptiveElements = []svgType{
		elemDesc, elemMetadata, elemTitle,
	}
	shapeElements = []svgType{
		elemCircle, elemEllipse, elemLine, elemPath,
		elemPolygon, elemPolyline, elemRect,
	}
	structuralElements = []svgType{
		elemDefs, elemG, elemSvg, elemSymbol, elemUse,
	}
	gradientElements = []svgType{
		elemLinearGradient, elemRadialGradient,
	}
	// Attributes
	coreAttributes = []svgType{
		attrId, attrXmlBase, attrXmlLang, attrXmlSpace,
	}
	presentationAttributes = []svgType{
		attrAlignmentBaseline, attrBaselineShift, attrClip,
		attrClipPath, attrClipRule, attrColor,
		attrColorInterpolation, attrColorInterpolationFilters,
		attrColorProfile, attrColorRendering, attrCursor,
		attrDirection, attrDisplay, attrDominantBaseline,
		attrEnableBackground, attrFill, attrFillOpacity,
		attrFillRule, attrFilter, attrFloodColor,
		attrFloodOpacity, attrFontFamily, attrFontSize,
		attrFontSizeAdjust, attrFontStretch, attrFontStyle,
		attrFontVariant, attrFontWeight,
		attrGlyphOrientationHorizontal,
		attrGlyphOrientationVertical, attrImageRendering,
		attrKerning, attrLetterSpacing, attrLightingColor,
		attrMarkerEnd, attrMarkerMid, attrMarkerStart, attrMask,
		attrOpacity, attrOverflow, attrPointerEvents,
		attrShapeRendering, attrStopColor, attrStopOpacity,
		attrStroke, attrStrokeDasharray, attrStrokeDashoffset,
		attrStrokeLinecap, attrStrokeLinejoin,
		attrStrokeMiterlimit, attrStrokeOpacity,
		attrStrokeWidth, attrTextAnchor, attrTextDecoration,
		attrTextRendering, attrUnicodeBidi, attrVisibility,
		attrWordSpacing, attrWritingMode,
	}
	conditionalProcessingAttributes = []svgType{
		attrRequiredExtensions, attrRequiredFeatures,
		attrSystemLanguage,
	}
	graphicalEventAttributes = []svgType{
		attrOnfocusin, attrOnfocusout, attrOnactivate,
		attrOnclick, attrOnmousedown, attrOnmouseup,
		attrOnmouseover, attrOnmousemove, attrOnmouseout,
		attrOnload,
	}
	xlinkAttributes = []svgType{
		attrXlinkHref, attrXlinkShow, attrXlinkActuate,
		attrXlinkType, attrXlinkRole, attrXlinkArcrole,
		attrXlinkTitle,
	}
	animationEventAttributes = []svgType{
		attrOnbegin, attrOnend, attrOnrepeat, attrOnload,
	}
	animationAttributeTargetAttributes = []svgType{
		attributeType, attributeName,
	}
	animationTimingAttributes = []svgType{
		attrBegin, attrDur, attrEnd, attrMin, attrMax,
		attrRestart, attrRepeatCount, attrRepeatDur, attrFill,
	}
	animationValueAttributes = []svgType{
		attrBegin, attrDur, attrEnd, attrMin, attrMax,
		attrRestart, attrRepeatCount, attrRepeatDur, attrFill,
	}
	animationAdditionAttributes = []svgType{
		attrAdditive, attrAccumulate,
	}
	filterPrimitiveAttributes = []svgType{
		attrX, attrY, attrWidth, attrHeight, attrResult,
	}
	lightSourceElements = []svgType{
		elemFeDiffuseLighting, elemFeSpecularLighting,
		elemFeDistantLight, elemFePointLight, elemFeSpotLight,
	}
	transferFunctionElementAttributes = []svgType{
		attrType, attrTableValues, attrSlope, attrIntercept,
		attrAmplitude, attrExponent, attrOffset,
	}
	filterPrimiveElements = []svgType{
		elemFeBlend, elemFeColorMatrix, elemFeComponentTransfer,
		elemFeComposite, elemFeConvolveMatrix,
		elemFeDiffuseLighting, elemFeDisplacementMap,
		elemFeFlood, elemFeGaussianBlur, elemFeImage,
		elemFeMerge, elemFeMorphology, elemFeOffset,
		elemFeSpecularLighting, elemFeTile, elemFeTurbulence,
	}
	documentEventAttributes = []svgType{
		attrOnunload, attrOnabort, attrOnerror, attrOnresize,
		attrOnscroll, attrOnzoom,
	}
)

func init() {
	v := group{bitVector: new(bitVector)}

	// <a>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/linking.html#AElement
	// Catagory
	v.Add(cataContainerElement)
	// Content model
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	v.AddAll(shapeElements...)
	v.AddAll(structuralElements...)
	v.AddAll(gradientElements...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(
		elemA, elemAltGlyphDef, elemClipPath, elemColorProfile,
		elemCursor, elemFilter, elemFont, elemFontFace,
		elemForeignObject, elemImage, elemMarker, elemMask,
		elemPattern, elemScript, elemStyle, elemSwitch,
		elemText, elemView,
	)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrTransform, attrTarget,
	)
	// DOM interface
	v.Add(domSVGAElement)
	elements[elemA] = v
	v.Clear()

	// <altGlyph>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/text.html#AltGlyphElement
	// Catagory
	v.AddAll(cataTextContentElement, cataTextContentChildElement)
	// Content Model
	// Any elements or character data.
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrX, attrY, attrDx, attrDy, attrGlyphRef, attrFormat,
		attrRotate, attrXlinkHref,
	)
	// DOM interface
	v.Add(domSVGAltGlyphElement)
	elements[elemAltGlyph] = v
	v.Clear()

	// <altGlyphDef>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/text.html#AltGlyphDefElement
	// Catagory
	// none
	// Content Model
	// Either:
	//	one or more ‘glyphRef’ elements, or
	//	one or more ‘altGlyphItem’ elements.
	v.AddAll(elemGlyphRef, elemAltGlyphItem)
	// Attributes
	v.AddAll(coreAttributes...)
	// DOM interface
	v.Add(domSVGAltGlyphDefElement)
	v.contract = append(v.contract, func(g group) error {
		if g.Has(elemGlyphRef) && !g.Has(elemAltGlyphItem) ||
			!g.Has(elemGlyphRef) && g.Has(elemAltGlyphItem) {
			return nil
		}
		return errors.New("<altGlyphDef> requires either: one or more " +
			"‘glyphRef’ elements, or one or more ‘altGlyphItem’ elements.")
	})
	elements[elemAltGlyphDef] = v
	v.Clear()

	// <altGlyphItem>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/text.html#AltGlyphItemElement
	// Catagory
	// none
	// Content Model
	// One or more ‘glyphRef’ elements.
	v.Add(elemGlyphRef)
	// Attributes
	v.AddAll(coreAttributes...)
	// DOM interface
	v.Add(domSVGAltGlyphItemElement)
	v.contract = append(v.contract, func(g group) error {
		if g.Has(elemGlyphRef) {
			return nil
		}
		return errors.New("<altGlyphItem> requires one or more ‘glyphRef’ child elements.")
	})
	elements[elemAltGlyphItem] = v
	v.Clear()

	// <animate>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/animate.html#AnimateElement
	// Catagory
	v.Add(cataAnimationElement)
	// Content Model
	v.AddAll(descriptiveElements...)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(animationEventAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(animationAttributeTargetAttributes...)
	v.AddAll(animationTimingAttributes...)
	v.AddAll(animationValueAttributes...)
	v.AddAll(animationAdditionAttributes...)
	v.AddAll(presentationAttributes...)
	v.Add(attrExternalResourcesRequired)
	// DOM interface
	v.Add(domSVGAnimateElement)
	elements[elemAnimate] = v
	v.Clear()

	// <animateColor>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/animate.html#AnimateColorElement
	// Catagory
	v.Add(cataAnimationElement)
	// Content model
	v.AddAll(descriptiveElements...)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(animationEventAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(animationAttributeTargetAttributes...)
	v.AddAll(animationTimingAttributes...)
	v.AddAll(animationValueAttributes...)
	v.AddAll(animationAdditionAttributes...)
	v.AddAll(presentationAttributes...)
	v.Add(attrExternalResourcesRequired)
	// DOM interface
	v.Add(domSVGAnimateColorElement)
	elements[elemAnimateColor] = v
	v.Clear()

	// <animateMotion>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/animate.html#AnimateMotionElement
	// Catagory
	v.Add(cataAnimationElement)
	// Content Model
	v.AddAll(descriptiveElements...)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(animationEventAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(animationTimingAttributes...)
	v.AddAll(animationValueAttributes...)
	v.AddAll(animationAdditionAttributes...)
	v.AddAll(
		attrExternalResourcesRequired, attrPath, attrKeyPoints,
		attrRotate, attrOrigin,
	)
	// DOM interface
	v.Add(domSVGAnimateMotionElement)
	v.contract = append(v.contract, func(g group) error {
		if g.Has(elemMpath) {
			return nil
		}
		return errors.New("<animateMotion> can have most one ‘mpath’ element.")
	})
	elements[elemAnimateMotion] = v
	v.Clear()

	// <animateTransform>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/animate.html#AnimateTransformElement
	// Catagory
	v.Add(cataAnimationElement)
	// Content Model
	v.AddAll(descriptiveElements...)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(animationEventAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(animationAttributeTargetAttributes...)
	v.AddAll(animationTimingAttributes...)
	v.AddAll(animationValueAttributes...)
	v.AddAll(animationAdditionAttributes...)
	v.AddAll(
		attrExternalResourcesRequired,
		attrType,
	)
	// DOM interface
	v.Add(domSVGAnimateTransformElement)
	elements[elemAnimateTransform] = v
	v.Clear()

	// <elemCircle>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/shapes.html#CircleElement
	// Catagory
	v.AddAll(
		cataBasicShapeElement, cataGraphicsElement,
		cataShapeElement,
	)
	// Content Model
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrTransform, attrCx, attrCy, attrR,
	)
	// DOM interface
	v.Add(domSVGCircleElement)
	elements[elemCircle] = v
	v.Clear()

	// <clipPath>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/masking.html#ClipPathElement

	// Catagory
	// none
	// Content Model
	v.AddAll(descriptiveElements...)
	v.AddAll(animationElements...)
	v.AddAll(shapeElements...)
	v.AddAll(elemText, elemUse)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrTransform, attrClipPathUnits,
	)
	// DOM interface
	v.Add(domSVGClipPathElement)
	elements[elemClipPath] = v
	v.Clear()

	// <color-profile>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/color.html#ColorProfileElement
	// Catagory
	// none
	// Content Model
	v.AddAll(descriptiveElements...)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(attrLocal, attrName, attrRenderingIntent, attrXlinkHref)
	// DOM interface
	v.Add(domSVGColorProfileElement)
	elements[elemColorProfile] = v
	v.Clear()

	// <cursor>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/interact.html#CursorElement
	// Catagory
	// none
	// Content Model
	v.AddAll(descriptiveElements...)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(
		attrExternalResourcesRequired, attrX, attrY,
		attrXlinkHref,
	)
	// DOM interface
	v.Add(domSVGCursorElement)
	elements[elemCursor] = v
	v.Clear()

	// <defs>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/struct.html#DefsElement
	// Catagory
	v.AddAll(cataContainerElement, cataStructuralElement)
	// Content Model
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	v.AddAll(shapeElements...)
	v.AddAll(structuralElements...)
	v.AddAll(gradientElements...)
	v.AddAll(
		elemA, elemAltGlyphDef, elemClipPath, elemColorProfile,
		elemCursor, elemFilter, elemFont, elemFontFace,
		elemForeignObject, elemImage, elemMarker, elemMask,
		elemPattern, elemScript, elemStyle, elemSwitch,
		elemText, elemView,
	)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrTransform,
	)
	// DOM interface
	v.Add(domSVGDefsElement)
	elements[elemDefs] = v
	v.Clear()

	// <desc>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/struct.html#DescElement
	// Catagory
	v.Add(cataDescriptiveElement)
	// Content Model
	// Any elements or character data.
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(attrClass, attrStyle)
	// DOM interface
	v.Add(domSVGDescElement)
	elements[elemDesc] = v
	v.Clear()

	// <ellipse>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/shapes.html#EllipseElement
	// Catagory
	v.AddAll(
		cataBasicShapeElement, cataGraphicsElement,
		cataShapeElement,
	)
	// Content Model
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrTransform, attrCx, attrCy, attrRx, attrRy,
	)
	// DOM interface
	v.Add(domSVGEllipseElement)
	elements[elemEllipse] = v
	v.Clear()

	// <feBlend>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feBlendElement
	// Catagory
	v.Add(cataFilterPrimitiveElement)
	// Content Model
	v.AddAll(elemAnimate, elemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(attrClass, attrStyle, attrIn, attrIn2, attrMode)
	// DOM interface
	v.Add(domSVGFEBlendElement)
	elements[elemFeBlend] = v
	v.Clear()

	// <feColorMatrix>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feColorMatrixElement
	// Catagory
	v.Add(cataFilterPrimitiveElement)
	// Content Model
	v.AddAll(elemAnimate, elemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(attrClass, attrStyle, attrIn, attrType, attrValues)
	// DOM interface
	v.Add(domSVGFEColorMatrixElement)
	elements[elemFeColorMatrix] = v
	v.Clear()

	// <feComponentTransfer>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feComponentTransferElement
	// Catagory
	v.Add(cataFilterPrimitiveElement)
	// Content Model
	v.AddAll(elemFeFuncA, elemFeFuncB, elemFeFuncG, elemFeFuncR)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(attrClass, attrStyle, attrIn)
	// DOM interface
	v.Add(domSVGFEComponentTransferElement)
	elements[elemFeComponentTransfer] = v
	v.Clear()

	// <feComposite>
	// Catagory
	v.Add(cataFilterPrimitiveElement)
	// Content Model
	v.AddAll(elemAnimate, elemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrIn, attrIn2, attrOperator,
		attrK1, attrK2, attrK3, attrK4,
	)
	// DOM interface
	v.Add(domSVGFECompositeElement)
	elements[elemFeComposite] = v
	v.Clear()

	// <feConvolveMatrix>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feConvolveMatrixElement
	// Catagory
	v.Add(cataFilterPrimitiveElement)
	// Content Model
	v.AddAll(elemAnimate, elemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrIn, attrOrder,
		attrKernelMatrix, attrDivisor, attrBias, attrTargetX,
		attrTargetY, attrEdgeMode, attrKernelUnitLength,
		attrPreserveAlpha,
	)
	// DOM interface
	v.Add(domSVGFEConvolveMatrixElement)
	elements[elemFeConvolveMatrix] = v
	v.Clear()

	// <feDiffuseLighting>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feDiffuseLightingElement
	// Catagory
	v.Add(cataFilterPrimitiveElement)
	// Content Model
	// Any number of descriptive elements and exactly one light
	// source element, in any order.
	v.AddAll(descriptiveElements...)
	v.AddAll(lightSourceElements...)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrIn, attrSurfaceScale,
		attrDiffuseConstant, attrKernelUnitLength,
	)
	// DOM interface
	v.Add(domSVGFEDiffuseLightingElement)
	v.contract = append(v.contract, func(g group) error {
		var count int
		err := errors.New("<feDiffuseLighting> only one light " +
			"source element permited")
		for _, v := range lightSourceElements {
			if g.Has(v) {
				count++
				if count > 1 {
					return err
				}
			}
		}
		return nil
	})
	elements[elemFeDiffuseLighting] = v
	v.Clear()

	// <feDisplacementMap>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feDisplacementMapElement
	// Catagory
	v.Add(cataFilterPrimitiveElement)
	// Content Model
	v.AddAll(elemAnimate, elemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrIn, attrIn2, attrScale,
		attrXChannelSelector, attrYChannelSelector,
	)
	// DOM interface
	v.Add(domSVGFEDisplacementMapElement)
	elements[elemFeDisplacementMap] = v
	v.Clear()

	// <feDistantLight>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feDistantLightElement
	// Catagory
	v.Add(cataLightSourceElement)
	// Content Model
	// Any number of the following elements, in any order:
	v.AddAll(elemAnimate, elemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(attrAzimuth, attrElevation)
	// DOM interface
	v.Add(domSVGFEDistantLightElement)
	elements[elemFeDistantLight] = v
	v.Clear()

	// <feFlood>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feFloodElement
	// Catagory
	v.Add(cataFilterPrimitiveElement)
	// Content Model
	// Any number of the following elements, in any order:
	v.AddAll(elemAnimate, elemAnimateColor, elemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(attrClass, attrStyle)
	// DOM interface
	v.Add(domSVGFEFloodElement)
	elements[elemFeFlood] = v
	v.Clear()

	// <feFuncA]>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feFuncAElement
	// Catagory
	// none
	// Content Model
	// Any number of the following elements, in any order:
	v.AddAll(elemAnimate, elemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(transferFunctionElementAttributes...)
	// DOM interface
	v.Add(domSVGFEFuncAElement)
	elements[elemFeFuncA] = v
	v.Clear()

	// <feFuncB>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feFuncBElement
	// Catagory
	// none
	// Content Model
	// Any number of the following elements, in any order:
	v.AddAll(elemAnimate, elemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(transferFunctionElementAttributes...)
	// DOM interface
	v.Add(domSVGFEFuncBElement)
	elements[elemFeFuncB] = v
	v.Clear()

	// <feFuncG>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feFuncGElement
	// Catagory
	// none
	// Content Model
	// Any number of the following elements, in any order:
	v.AddAll(elemAnimate, elemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(transferFunctionElementAttributes...)
	// DOM interface
	v.Add(domSVGFEFuncGElement)
	elements[elemFeFuncG] = v
	v.Clear()

	// <feFuncR>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feFuncRElement
	// Catagory
	// none
	// Content Model
	// Any number of the following elements, in any order:
	v.AddAll(elemAnimate, elemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(transferFunctionElementAttributes...)
	// DOM interface
	v.Add(domSVGFEFuncRElement)
	elements[elemFeFuncR] = v
	v.Clear()

	// <feGaussianBlur>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feGaussianBlurElement
	// catagory
	v.Add(cataFilterPrimitiveElement)
	// Content Model
	// Any number of the following elements, in any order:
	v.AddAll(elemAnimate, elemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(attrClass, attrStyle, attrIn, attrStdDeviation)
	// dom interface
	v.Add(domSVGFEGaussianBlurElement)
	elements[elemFeGaussianBlur] = v
	v.Clear()

	// <feImage>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feImageElement
	// Catagory
	v.Add(cataFilterPrimitiveElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(elemAnimate, elemAnimateTransform, elemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrPreserveAspectRatio, attrXlinkHref,
	)
	// DOM interface
	v.Add(domSVGFEImageElement)
	elements[elemFeImage] = v
	v.Clear()

	// <feMerge>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feMergeElement
	// Catagory
	v.Add(cataFilterPrimitiveElement)
	// Content model
	// Any number of the following elements, in any order:
	v.Add(elemFeMergeNode)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(attrClass, attrStyle)
	// DOM interface
	v.Add(domSVGFEMergeElement)
	elements[elemFeMerge] = v
	v.Clear()

	// <feMergeNode>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feMergeNodeElement
	// Catagory
	// none
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(elemAnimate, elemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(
		attrIn,
	)
	// DOM interface
	v.Add(domSVGFEMergeNodeElement)
	elements[elemFeMergeNode] = v
	v.Clear()

	// <feMorphology>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feMorphologyElement
	// Catagory
	v.Add(cataFilterPrimitiveElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(elemAnimate, elemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(attrClass, attrStyle, attrIn, attrOperator, attrRadius)
	// DOM interface
	v.Add(domSVGFEMorphologyElement)
	elements[elemFeMorphology] = v
	v.Clear()

	// <feOffset>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feOffsetElement
	// Catagory
	v.Add(cataFilterPrimitiveElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(elemAnimate, elemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(attrClass, attrStyle, attrIn, attrDx, attrDy)
	// DOM interface
	v.Add(domSVGFEOffsetElement)
	elements[elemFeOffset] = v
	v.Clear()

	// <fePointLight>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#fePointLightElement
	// Catagory
	v.Add(cataLightSourceElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(elemAnimate, elemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(attrX, attrY, attrZ)
	// DOM interface
	v.Add(domSVGFEPointLightElement)
	elements[elemFePointLight] = v
	v.Clear()

	// <feSpecularLighting>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feSpecularLightingElement
	// Catagory
	v.Add(cataFilterPrimitiveElement)
	// Content model
	// Any number of descriptive elements and exactly one light
	// source element, in any order.
	v.AddAll(descriptiveElements...)
	v.AddAll(lightSourceElements...)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrIn, attrSurfaceScale,
		attrSpecularConstant, attrSpecularExponent,
		attrKernelUnitLength,
	)
	// DOM interface
	v.Add(domSVGFESpecularLightingElement)
	v.contract = append(v.contract, func(g group) error {
		var count int
		err := errors.New("<feSpecularLighting> only one light " +
			"source element permited")
		for _, v := range lightSourceElements {
			if g.Has(v) {
				count++
				if count > 1 {
					return err
				}
			}
		}
		return nil
	})
	elements[elemFeSpecularLighting] = v
	v.Clear()

	// <feSpotLight>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feSpotLightElement
	// Catagory
	v.Add(cataLightSourceElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(elemAnimate, elemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(
		attrX, attrY, attrZ, attrPointsAtX, attrPointsAtY,
		attrPointsAtZ, attrSpecularExponent,
		attrLimitingConeAngle,
	)
	// DOM interface
	v.Add(domSVGFESpotLightElement)
	elements[elemFeSpotLight] = v
	v.Clear()

	// <feTile>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feTileElement
	// Catagory
	v.Add(cataFilterPrimitiveElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(elemAnimate, elemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(attrClass, attrStyle, attrIn)
	// DOM interface
	v.Add(domSVGFETileElement)
	elements[elemFeTile] = v
	v.Clear()

	// <feTurbulence>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feTurbulenceElement
	// Catagory
	v.Add(cataFilterPrimitiveElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(elemAnimate, elemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrBaseFrequency, attrNumOctaves,
		attrSeed, attrStitchTiles, attrType,
	)
	// DOM interface
	v.Add(domSVGFETurbulenceElement)
	elements[elemFeTurbulence] = v
	v.Clear()

	// <filter>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#FilterElement
	// Catagory
	// none
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(descriptiveElements...)
	v.AddAll(filterPrimiveElements...)
	v.AddAll(elemAnimate, elemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrX, attrY, attrWidth, attrHeight, attrFilterRes,
		attrFilterUnits, attrPrimitiveUnits, attrXlinkHref,
	)
	// DOM interface
	v.Add(domSVGFilterElement)
	elements[elemFilter] = v
	v.Clear()

	// <font>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/fonts.html#FontElement
	// Catagory
	// none
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(descriptiveElements...)
	v.AddAll(
		elemFontFace, elemGlyph, elemHkern, elemMissingGlyph,
		elemVkern,
	)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrHorizOriginX, attrHorizOriginY, attrHorizAdvX,
		attrVertOriginX, attrVertOriginY, attrVertAdvY,
	)
	// DOM interface
	v.Add(domSVGFontElement)
	elements[elemFont] = v
	v.Clear()

	// <font-face>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/fonts.html#FontFaceElement
	// Catagory
	// none
	// Content model
	// Any number of descriptive elements and at most one
	// ‘font-face-src’ element, in any order.
	v.AddAll(descriptiveElements...)
	v.Add(elemFontFaceSrc)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(
		attrFontFamily, attrFontStyle, attrFontVariant,
		attrFontWeight, attrFontStretch, attrFontSize,
		attrUnicodeRange, attrUnitsPerEm, attrPanose1,
		attrStemv, attrStemh, attrSlope, attrCapHeight,
		attrXHeight, attrAccentHeight, attrAscent, attrDescent,
		attrWidths, attrBbox, attrIdeographic, attrAlphabetic,
		attrMathematical, attrHanging, attrVIdeographic,
		attrVAlphabetic, attrVMathematical, attrVHanging,
		attrUnderlinePosition, attrUnderlineThickness,
		attrStrikethroughPosition, attrStrikethroughThickness,
		attrOverlinePosition, attrOverlineThickness,
	)
	// DOM interface
	v.Add(domSVGFontFaceElement)
	v.contract = append(v.contract, func(g group) error {
		// Any number of descriptive elements and at most one
		// ‘font-face-src’ element, in any order.
		return nil
	})
	elements[elemFontFace] = v
	v.Clear()

	// <font-face-format>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/fonts.html#FontFaceFormatElement
	// Catagory
	// none
	// Content model
	// empty
	// Attributes
	v.AddAll(coreAttributes...)
	v.Add(attrString)
	// DOM interface
	v.Add(domSVGFontFaceFormatElement)
	v.contract = append(v.contract, func(g group) error {
		// should be empty
		return nil
	})
	elements[elemFontFaceFormat] = v
	v.Clear()

	// <font-face-name>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/fonts.html#FontFaceNameElement
	// Catagory
	// none
	// Content model
	// empty
	// Attributes
	v.AddAll(coreAttributes...)
	v.Add(attrName)
	// DOM interface
	v.Add(domSVGFontFaceNameElement)
	v.contract = append(v.contract, func(g group) error {
		// empty
		return nil
	})
	elements[elemFontFaceName] = v
	v.Clear()

	// <font-face-src>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/fonts.html#FontFaceSrcElement
	// Catagory
	// none
	// Content model
	// One or more of the following elements, in any order:
	v.AddAll(elemFontFaceName, elemFontFaceUri)
	// Attributes
	v.AddAll(coreAttributes...)
	// DOM interface
	v.Add(domSVGFontFaceSrcElement)
	v.contract = append(v.contract, func(g group) error {
		// One or more of the following elements, in any order:
		//	elemFontFaceName, elemFontFaceUri
		return nil
	})
	elements[elemFontFaceSrc] = v
	v.Clear()

	// <font-face-uri>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/fonts.html#FontFaceURIElement
	// Catagory
	// none
	// Content model
	// Any number of the following elements, in any order:
	v.Add(elemFontFaceFormat)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(xlinkAttributes...)
	v.Add(attrXlinkHref)
	// DOM interface
	v.AddAll(domSVGFontFaceUriElement)
	elements[elemFontFaceUri] = v
	v.Clear()

	// <foreignObject>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/extend.html#ForeignObjectElement
	// Catagory
	// none
	// Content model
	// Any elements or character data.
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrTransform, attrX, attrY, attrWidth, attrHeight,
	)
	// DOM interface
	v.Add(domSVGForeignObjectElement)
	elements[elemForeignObject] = v
	v.Clear()

	// <foreignObject>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/extend.html#ForeignObjectElement
	// Catagory
	// none
	// Content model
	// Any elements or character data.
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrTransform, attrX, attrY, attrWidth, attrHeight,
	)
	// DOM interface
	v.Add(domSVGForeignObjectElement)
	elements[elemForeignObject] = v
	v.Clear()

	// <g>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/struct.html#GElement
	// Catagory
	v.AddAll(cataContainerElement, cataStructuralElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	v.AddAll(shapeElements...)
	v.AddAll(structuralElements...)
	v.AddAll(gradientElements...)
	v.AddAll(
		elemA, elemAltGlyphDef, elemClipPath, elemColorProfile,
		elemCursor, elemFilter, elemFont, elemFontFace,
		elemForeignObject, elemImage, elemMarker, elemMask,
		elemPattern, elemScript, elemStyle, elemSwitch,
		elemText, elemView,
	)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrTransform,
	)
	// DOM interface
	v.Add(domSVGGElement)
	elements[elemG] = v
	v.Clear()

	// <glyph>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/fonts.html#GlyphElement
	// Catagory
	v.Add(cataContainerElement)
	//
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	v.AddAll(shapeElements...)
	v.AddAll(structuralElements...)
	v.AddAll(gradientElements...)
	v.AddAll(
		elemA, elemAltGlyphDef, elemClipPath, elemColorProfile,
		elemCursor, elemFilter, elemFont, elemFontFace,
		elemForeignObject, elemImage, elemMarker, elemMask,
		elemPattern, elemScript, elemStyle, elemSwitch,
		elemText, elemView,
	)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrD, attrHorizAdvX,
		attrVertOriginX, attrVertOriginY, attrVertAdvY,
		attrUnicode, attrGlyphName, attrOrientation,
		attrArabicForm, attrLang,
	)
	// DOM interface
	v.Add(domSVGGlyphElement)
	elements[elemGlyph] = v
	v.Clear()

	// <glyphRef>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/text.html#GlyphRefElement
	// Catagory
	// none
	// Content model
	// empty
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrX, attrY, attrDx, attrDy,
		attrGlyphRef, attrFormat, attrXlinkHref,
	)
	// DOM interface
	v.Add(domSVGGlyphRefElement)
	elements[elemGlyphRef] = v
	v.Clear()

	// <hkern>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/fonts.html#HKernElement
	// Catagory
	// none
	// Content model
	// empty
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(attrU1, attrG1, attrU2, attrG2, attrK)
	// DOM interface
	v.Add(domSVGHKernElement)
	elements[elemHkern] = v
	v.Clear()

	// <image>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/struct.html#ImageElement
	// Catagory
	v.AddAll(cataGraphicsElement, cataGraphicsReferencingElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrPreserveAspectRatio, attrTransform, attrX, attrY,
		attrWidth, attrHeight, attrXlinkHref,
	)
	// DOM interface
	v.Add(domSVGImageElement)
	elements[elemImage] = v
	v.Clear()

	// <line>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/shapes.html#LineElement
	// Catagory
	v.AddAll(
		cataBasicShapeElement, cataGraphicsElement,
		cataShapeElement,
	)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrTransform, attrX1, attrY1, attrX2, attrY2,
	)
	// DOM interface
	v.Add(domSVGLineElement)
	v.contract = append(v.contract, func(g group) error {
		// Any number of the following elements, in any order:
		// descriptive elements
		// animation elements
		return nil
	})
	elements[elemLine] = v
	v.Clear()

	// <linearGradient>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/pservers.html#LinearGradientElement
	// Catagory
	v.Add(cataGradientElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(descriptiveElements...)
	v.AddAll(elemAnimate, elemAnimateTransform, elemSet, elemStop)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrX1, attrY1, attrX2, attrY2, attrGradientUnits,
		attrGradientTransform, attrSpreadMethod, attrXlinkHref,
	)
	// DOM interface
	v.Add(domSVGLinearGradientElement)
	v.contract = append(v.contract, func(g group) error {
		// Any number of the following elements, in any order:
		// descriptive elements
		// ‘animate’ ‘animateTransform’ ‘set’ ‘stop’
		return nil
	})
	elements[elemLinearGradient] = v
	v.Clear()

	// <marker>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/painting.html#MarkerElement
	// Catagory
	v.Add(cataContainerElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	v.AddAll(shapeElements...)
	v.AddAll(structuralElements...)
	v.AddAll(gradientElements...)
	v.AddAll(
		elemA, elemAltGlyphDef, elemClipPath, elemColorProfile,
		elemCursor, elemFilter, elemFont, elemFontFace,
		elemForeignObject, elemImage, elemMarker, elemMask,
		elemPattern, elemScript, elemStyle, elemSwitch,
		elemText, elemView,
	)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrViewBox, attrPreserveAspectRatio, attrRefX,
		attrRefY, attrMarkerUnits, attrMarkerWidth,
		attrMarkerHeight, attrOrient,
	)
	// DOM interface
	v.Add(domSVGMarkerElement)
	v.contract = append(v.contract, func(g group) error {
		// Any number of the following elements, in any order:
		// animationElements
		// descriptiveElements
		// shapeElements
		// structuralElements
		// gradientElements
		//
		// 	elemA, elemAltGlyphDef, elemClipPath,
		// 	elemColorProfile, elemCursor, elemFilter,
		// 	elemFont, elemFontFace, elemForeignObject,
		// 	elemImage, elemMarker, elemMask, elemPattern,
		// 	elemScript, elemStyle, elemSwitch, elemText,
		// 	elemView,
		return nil
	})
	elements[elemMarker] = v
	v.Clear()

	// <mask>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/masking.html#MaskElement
	// Catagory
	v.Add(cataContainerElement)
	//
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	v.AddAll(shapeElements...)
	v.AddAll(structuralElements...)
	v.AddAll(gradientElements...)
	v.AddAll(
		elemA, elemAltGlyphDef, elemClipPath, elemColorProfile,
		elemCursor, elemFilter, elemFont, elemFontFace,
		elemForeignObject, elemImage, elemMarker, elemMask,
		elemPattern, elemScript, elemStyle, elemSwitch,
		elemText, elemView,
	)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrX, attrY, attrWidth, attrHeight, attrMaskUnits,
		attrMaskContentUnits,
	)
	// DOM interface
	v.Add(domSVGMaskElement)
	elements[elemMask] = v
	v.Clear()

	// <metadata>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/metadata.html#MetadataElement
	// Catagory
	v.Add(cataDescriptiveElement)
	//
	// Content model
	// Any elements or character data.
	// Attributes
	v.AddAll(coreAttributes...)
	// DOM interface
	v.Add(domSVGMetadataElement)
	elements[elemMetadata] = v
	v.Clear()

	// <missing-glyph>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/fonts.html#MissingGlyphElement
	// Catagory
	v.Add(cataContainerElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	v.AddAll(shapeElements...)
	v.AddAll(structuralElements...)
	v.AddAll(gradientElements...)
	v.AddAll(
		elemA, elemAltGlyphDef, elemClipPath, elemColorProfile,
		elemCursor, elemFilter, elemFont, elemFontFace,
		elemForeignObject, elemImage, elemMarker, elemMask,
		elemPattern, elemScript, elemStyle, elemSwitch,
		elemText, elemView,
	)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrD, attrHorizAdvX,
		attrVertOriginX, attrVertOriginY, attrVertAdvY,
	)
	// DOM interface
	v.Add(domSVGMissingGlyphElement)
	elements[elemMissingGlyph] = v
	v.Clear()

	// <mpath>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/animate.html#MPathElement
	// Catagory
	//
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(descriptiveElements...)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(attrExternalResourcesRequired, attrXlinkHref)
	// DOM interface
	v.Add(domSVGMPathElement)
	elements[elemMpath] = v
	v.Clear()

	// <path>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/paths.html#PathElement
	// Catagory
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	//
	// Content model
	//
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrTransform, attrD,
	)
	// DOM interface
	v.Add(domSVGPathElement)
	elements[elemPath] = v
	v.Clear()

	// <pattern>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/pservers.html#PatternElement
	// Catagory
	v.Add(cataContainerElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	v.AddAll(shapeElements...)
	v.AddAll(structuralElements...)
	v.AddAll(gradientElements...)
	v.AddAll(
		elemA, elemAltGlyphDef, elemClipPath, elemColorProfile,
		elemCursor, elemFilter, elemFont, elemFontFace,
		elemForeignObject, elemImage, elemMarker, elemMask,
		elemPattern, elemScript, elemStyle, elemSwitch,
		elemText, elemView,
	)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrViewBox, attrPreserveAspectRatio, attrX, attrY,
		attrWidth, attrHeight, attrPatternUnits,
		attrPatternContentUnits, attrPatternTransform,
		attrXlinkHref,
	)
	// DOM interface
	v.Add(domSVGPatternElement)
	elements[elemPattern] = v
	v.Clear()

	// <polygon>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/shapes.html#PolygonElement
	// Catagory
	v.AddAll(
		cataBasicShapeElement, cataGraphicsElement,
		cataShapeElement,
	)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrTransform, attrPoints,
	)
	// DOM interface
	v.Add(domSVGPolygonElement)
	elements[elemPolygon] = v
	v.Clear()

	// <polyline>
	// Catagory
	v.AddAll(
		cataBasicShapeElement, cataGraphicsElement,
		cataShapeElement,
	)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrTransform, attrPoints,
	)
	// DOM interface
	v.Add(domSVGPolylineElement)
	elements[elemPolyline] = v
	v.Clear()

	// <radialGradient>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/pservers.html#RadialGradientElement
	// Catagory
	v.Add(cataGradientElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(descriptiveElements...)
	v.AddAll(elemAnimate, elemAnimateTransform, elemSet, elemStop)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrCx, attrCy, attrR, attrFx, attrFy,
		attrGradientUnits, attrGradientTransform,
		attrSpreadMethod, attrXlinkHref,
	)
	// DOM interface
	v.Add(domSVGRadialGradientElement)
	elements[elemRadialGradient] = v
	v.Clear()

	// <rect>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/shapes.html#RectElement
	// Catagory
	v.AddAll(
		cataBasicShapeElement, cataGraphicsElement,
		cataShapeElement,
	)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrTransform, attrX, attrY, attrWidth, attrHeight,
		attrRx, attrRy,
	)
	// DOM interface
	v.Add(domSVGRectElement)
	elements[elemRect] = v
	v.Clear()

	// <script>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/script.html#ScriptElement
	// Catagory
	// none
	// Content model
	// Any elements or character data.
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(attrExternalResourcesRequired, attrType, attrXlinkHref)
	// DOM interface
	v.Add(domSVGScriptElement)
	elements[elemScript] = v
	v.Clear()

	// <set>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/animate.html#SetElement
	// Catagory
	v.Add(cataAnimationElement)
	//
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(descriptiveElements...)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(animationEventAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(animationAttributeTargetAttributes...)
	v.AddAll(animationTimingAttributes...)
	v.AddAll(
		attrExternalResourcesRequired,
		attrTo,
	)
	// DOM interface
	v.Add(domSVGSetElement)
	elements[elemSet] = v
	v.Clear()

	// <stop>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/pservers.html#StopElement
	// Catagory
	// none
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(elemAnimate, elemAnimateColor, elemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(attrClass, attrStyle, attrOffset)
	// DOM interface
	v.Add(domSVGStopElement)
	elements[elemStop] = v
	v.Clear()

	// <style>
	// Catagory
	// none
	// Content model
	// Any elements or character data.
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(attrType, attrMedia, attrTitle)
	// DOM interface
	v.Add(domSVGStyleElement)
	elements[elemStyle] = v
	v.Clear()

	// <svg>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/styling.html#StyleElement
	// Catagory
	v.AddAll(cataContainerElement, cataStructuralElement)
	//
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	v.AddAll(shapeElements...)
	v.AddAll(structuralElements...)
	v.AddAll(gradientElements...)
	v.AddAll(
		elemA, elemAltGlyphDef, elemClipPath, elemColorProfile,
		elemCursor, elemFilter, elemFont, elemFontFace,
		elemForeignObject, elemImage, elemMarker, elemMask,
		elemPattern, elemScript, elemStyle, elemSwitch,
		elemText, elemView,
	)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(documentEventAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrX, attrY, attrWidth, attrHeight, attrViewBox,
		attrPreserveAspectRatio, attrZoomAndPan, attrVersion,
		attrContentScriptType, attrContentStyleType, attrX,
		attrY, attrWidth, attrHeight, attrVersion,
		attrBaseProfile,
	)
	// DOM interface
	v.Add(domSVGSVGElement)
	elements[elemSvg] = v
	v.Clear()

	// <switch>
	// Catagory
	v.Add(cataContainerElement)
	//
	// Content model
	// Any number of the following elements, in any order:
	// animation elements
	// descriptive elements
	// shape elements
	v.AddAll(
		elemA, elemForeignObject, elemG, elemImage, elemSvg,
		elemSwitch, elemText, elemUse,
	)
	//
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrTransform,
	)
	// DOM interface
	v.Add(domSVGSwitchElement)
	elements[elemSwitch] = v
	v.Clear()

	// <symbol>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/struct.html#SymbolElement
	// Catagory
	v.AddAll(cataContainerElement, cataStructuralElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	v.AddAll(shapeElements...)
	v.AddAll(structuralElements...)
	v.AddAll(gradientElements...)
	v.AddAll(
		elemA, elemAltGlyphDef, elemClipPath, elemColorProfile,
		elemCursor, elemFilter, elemFont, elemFontFace,
		elemForeignObject, elemImage, elemMarker, elemMask,
		elemPattern, elemScript, elemStyle, elemSwitch,
		elemText, elemView,
	)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrPreserveAspectRatio, attrViewBox,
	)
	// DOM interface
	v.Add(domSVGSymbolElement)
	elements[elemSymbol] = v
	v.Clear()

	// <text>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/text.html#TextElement
	// Catagory
	v.AddAll(cataGraphicsElement, cataTextContentElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	v.AddAll(cataTextContentChildElement)
	v.Add(elemA)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrTransform, attrLengthAdjust, attrX, attrY, attrDx,
		attrDy, attrRotate, attrTextLength,
	)
	// DOM interface
	v.Add(domSVGTextElement)
	elements[elemText] = v
	v.Clear()

	// <textPath>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/text.html#TextPathElement
	// Catagory
	v.AddAll(cataTextContentElement, cataTextContentChildElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(descriptiveElements...)
	v.AddAll(
		elemA, elemAltGlyph, elemAnimate, elemAnimateColor,
		elemSet, elemTref, elemTspan,
	)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrXlinkHref, attrStartOffset, attrMethod, attrSpacing,
	)
	// DOM interface
	v.Add(domSVGTextPathElement)
	elements[elemTextPath] = v
	v.Clear()

	// <title>
	// Catagory
	v.Add(cataDescriptiveElement)
	// Content model
	// Any elements or character data.
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(attrClass, attrStyle)
	// DOM interface
	v.Add(domSVGTitleElement)
	elements[elemTitle] = v
	v.Clear()

	// <tref>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/text.html#TRefElement
	// Catagory
	v.AddAll(cataTextContentElement, cataTextContentChildElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(descriptiveElements...)
	v.AddAll(elemAnimate, elemAnimateColor, elemSet)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrXlinkHref,
	)
	// DOM interface
	v.Add(domSVGTRefElement)
	elements[elemTref] = v
	v.Clear()

	// <tspan>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/text.html#TSpanElement
	// Catagory
	v.AddAll(cataTextContentElement, cataTextContentChildElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(descriptiveElements...)
	v.AddAll(
		elemA, elemAltGlyph, elemAnimate, elemAnimateColor,
		elemSet, elemTref, elemTspan,
	)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrX, attrY, attrDx, attrDy, attrRotate,
		attrTextLength, attrLengthAdjust,
	)
	// DOM interface
	v.Add(domSVGTSpanElement)
	elements[elemTspan] = v
	v.Clear()

	// <use>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/struct.html#UseElement
	// Catagory
	v.AddAll(
		cataGraphicsElement, cataGraphicsReferencingElement,
		cataStructuralElement,
	)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(
		attrClass, attrStyle, attrExternalResourcesRequired,
		attrTransform, attrX, attrY, attrWidth, attrHeight,
		attrXlinkHref,
	)
	// DOM interface
	v.Add(domSVGUseElement)
	elements[elemUse] = v
	v.Clear()

	// <view>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/linking.html#ViewElement
	// Catagory
	// none
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(descriptiveElements...)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(
		attrExternalResourcesRequired, attrViewBox,
		attrPreserveAspectRatio, attrZoomAndPan, attrViewTarget,
	)
	// DOM interface
	v.Add(domSVGViewElement)
	elements[elemView] = v
	v.Clear()

	// <vkern>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/fonts.html#VKernElement
	// Catagory
	// none
	// Content model
	// empty
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(attrU1, attrG1, attrU2, attrG2, attrK)
	// DOM interface
	v.Add(domSVGHKernElement)
	elements[elemVkern] = v
	v.Clear()
}
