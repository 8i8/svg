// https://dev.w3.org/SVG/profiles/1.1F2/publish/
package spec

// svg elemtent has baseProfile in attributes twice
// svg attribuites not in full list
// onunload onabort onerror onresize onscroll onzoom
// style title attribute not there

import "errors"

var (
	// Elements
	animationElements = []svgType{
		ElemAnimate, ElemAnimateColor, ElemAnimateMotion,
		ElemAnimateTransform, ElemSet,
	}
	descriptiveElements = []svgType{
		ElemDesc, ElemMetadata, ElemTitle,
	}
	shapeElements = []svgType{
		ElemCircle, ElemEllipse, ElemLine, ElemPath,
		ElemPolygon, ElemPolyline, ElemRect,
	}
	structuralElements = []svgType{
		ElemDefs, ElemG, ElemSvg, ElemSymbol, ElemUse,
	}
	gradientElements = []svgType{
		ElemLinearGradient, ElemRadialGradient,
	}
	// Attributes
	coreAttributes = []svgType{
		AttrId, AttrXmlBase, AttrXmlLang, AttrXmlSpace,
	}
	presentationAttributes = []svgType{
		AttrAlignmentBaseline, AttrBaselineShift, AttrClip,
		AttrClipPath, AttrClipRule, AttrColor,
		AttrColorInterpolation, AttrColorInterpolationFilters,
		AttrColorProfile, AttrColorRendering, AttrCursor,
		AttrDirection, AttrDisplay, AttrDominantBaseline,
		AttrEnableBackground, AttrFill, AttrFillOpacity,
		AttrFillRule, AttrFilter, AttrFloodColor,
		AttrFloodOpacity, AttrFontFamily, AttrFontSize,
		AttrFontSizeAdjust, AttrFontStretch, AttrFontStyle,
		AttrFontVariant, AttrFontWeight,
		AttrGlyphOrientationHorizontal,
		AttrGlyphOrientationVertical, AttrImageRendering,
		AttrKerning, AttrLetterSpacing, AttrLightingColor,
		AttrMarkerEnd, AttrMarkerMid, AttrMarkerStart, AttrMask,
		AttrOpacity, AttrOverflow, AttrPointerEvents,
		AttrShapeRendering, AttrStopColor, AttrStopOpacity,
		AttrStroke, AttrStrokeDasharray, AttrStrokeDashoffset,
		AttrStrokeLinecap, AttrStrokeLinejoin,
		AttrStrokeMiterlimit, AttrStrokeOpacity,
		AttrStrokeWidth, AttrTextAnchor, AttrTextDecoration,
		AttrTextRendering, AttrUnicodeBidi, AttrVisibility,
		AttrWordSpacing, AttrWritingMode,
	}
	conditionalProcessingAttributes = []svgType{
		AttrRequiredExtensions, AttrRequiredFeatures,
		AttrSystemLanguage,
	}
	graphicalEventAttributes = []svgType{
		AttrOnfocusin, AttrOnfocusout, AttrOnactivate,
		AttrOnclick, AttrOnmousedown, AttrOnmouseup,
		AttrOnmouseover, AttrOnmousemove, AttrOnmouseout,
		AttrOnload,
	}
	xlinkAttributes = []svgType{
		AttrXlinkHref, AttrXlinkShow, AttrXlinkActuate,
		AttrXlinkType, AttrXlinkRole, AttrXlinkArcrole,
		AttrXlinkTitle,
	}
	animationEventAttributes = []svgType{
		AttrOnbegin, AttrOnend, AttrOnrepeat, AttrOnload,
	}
	animationAttributeTargetAttributes = []svgType{
		AttrType, AttrName,
	}
	animationTimingAttributes = []svgType{
		AttrBegin, AttrDur, AttrEnd, AttrMin, AttrMax,
		AttrRestart, AttrRepeatCount, AttrRepeatDur, AttrFill,
	}
	animationValueAttributes = []svgType{
		AttrBegin, AttrDur, AttrEnd, AttrMin, AttrMax,
		AttrRestart, AttrRepeatCount, AttrRepeatDur, AttrFill,
	}
	animationAdditionAttributes = []svgType{
		AttrAdditive, AttrAccumulate,
	}
	filterPrimitiveAttributes = []svgType{
		AttrX, AttrY, AttrWidth, AttrHeight, AttrResult,
	}
	lightSourceElements = []svgType{
		ElemFeDiffuseLighting, ElemFeSpecularLighting,
		ElemFeDistantLight, ElemFePointLight, ElemFeSpotLight,
	}
	transferFunctionElementAttributes = []svgType{
		AttrType, AttrTableValues, AttrSlope, AttrIntercept,
		AttrAmplitude, AttrExponent, AttrOffset,
	}
	filterPrimiveElements = []svgType{
		ElemFeBlend, ElemFeColorMatrix, ElemFeComponentTransfer,
		ElemFeComposite, ElemFeConvolveMatrix,
		ElemFeDiffuseLighting, ElemFeDisplacementMap,
		ElemFeFlood, ElemFeGaussianBlur, ElemFeImage,
		ElemFeMerge, ElemFeMorphology, ElemFeOffset,
		ElemFeSpecularLighting, ElemFeTile, ElemFeTurbulence,
	}
	documentEventAttributes = []svgType{
		AttrOnunload, AttrOnabort, AttrOnerror, AttrOnresize,
		AttrOnscroll, AttrOnzoom,
	}
)

func init() {
	v := group{Set: new(Set)}

	// <a>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/linking.html#AElement
	// Catagory
	v.Add(CataContainerElement)
	// Content model
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	v.AddAll(shapeElements...)
	v.AddAll(structuralElements...)
	v.AddAll(gradientElements...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(
		ElemA, ElemAltGlyphDef, ElemClipPath, ElemColorProfile,
		ElemCursor, ElemFilter, ElemFont, ElemFontFace,
		ElemForeignObject, ElemImage, ElemMarker, ElemMask,
		ElemPattern, ElemScript, ElemStyle, ElemSwitch,
		ElemText, ElemView,
	)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrTransform, AttrTarget,
	)
	// DOM interface
	v.Add(DomSVGAElement)
	Elements[ElemA] = v
	v.Clear()

	// <altGlyph>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/text.html#AltGlyphElement
	// Catagory
	v.AddAll(CataTextContentElement, CataTextContentChildElement)
	// Content Model
	// Any elements or character data.
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrX, AttrY, AttrDx, AttrDy, AttrGlyphRef, AttrFormat,
		AttrRotate, AttrXlinkHref,
	)
	// DOM interface
	v.Add(DomSVGAltGlyphElement)
	Elements[ElemAltGlyph] = v
	v.Clear()

	// <altGlyphDef>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/text.html#AltGlyphDefElement
	// Catagory
	// none
	// Content Model
	// Either:
	//	one or more ‘glyphRef’ elements, or
	//	one or more ‘altGlyphItem’ elements.
	v.AddAll(ElemGlyphRef, ElemAltGlyphItem)
	// Attributes
	v.AddAll(coreAttributes...)
	// DOM interface
	v.Add(DomSVGAltGlyphDefElement)
	v.contract = append(v.contract, func(g group) error {
		if g.Has(ElemGlyphRef) && !g.Has(ElemAltGlyphItem) ||
			!g.Has(ElemGlyphRef) && g.Has(ElemAltGlyphItem) {
			return nil
		}
		return errors.New("<altGlyphDef> requires either: one or more " +
			"‘glyphRef’ elements, or one or more ‘altGlyphItem’ elements.")
	})
	Elements[ElemAltGlyphDef] = v
	v.Clear()

	// <altGlyphItem>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/text.html#AltGlyphItemElement
	// Catagory
	// none
	// Content Model
	// One or more ‘glyphRef’ elements.
	v.Add(ElemGlyphRef)
	// Attributes
	v.AddAll(coreAttributes...)
	// DOM interface
	v.Add(DomSVGAltGlyphItemElement)
	v.contract = append(v.contract, func(g group) error {
		if g.Has(ElemGlyphRef) {
			return nil
		}
		return errors.New("<altGlyphItem> requires one or more ‘glyphRef’ child elements.")
	})
	Elements[ElemAltGlyphItem] = v
	v.Clear()

	// <animate>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/animate.html#AnimateElement
	// Catagory
	v.Add(CataAnimationElement)
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
	v.Add(AttrExternalResourcesRequired)
	// DOM interface
	v.Add(DomSVGAnimateElement)
	Elements[ElemAnimate] = v
	v.Clear()

	// <animateColor>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/animate.html#AnimateColorElement
	// Catagory
	v.Add(CataAnimationElement)
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
	v.Add(AttrExternalResourcesRequired)
	// DOM interface
	v.Add(DomSVGAnimateColorElement)
	Elements[ElemAnimateColor] = v
	v.Clear()

	// <animateMotion>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/animate.html#AnimateMotionElement
	// Catagory
	v.Add(CataAnimationElement)
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
		AttrExternalResourcesRequired, AttrPath, AttrKeyPoints,
		AttrRotate, AttrOrigin,
	)
	// DOM interface
	v.Add(DomSVGAnimateMotionElement)
	v.contract = append(v.contract, func(g group) error {
		if g.Has(ElemMpath) {
			return nil
		}
		return errors.New("<animateMotion> can have most one ‘mpath’ element.")
	})
	Elements[ElemAnimateMotion] = v
	v.Clear()

	// <animateTransform>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/animate.html#AnimateTransformElement
	// Catagory
	v.Add(CataAnimationElement)
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
		AttrExternalResourcesRequired,
		AttrType,
	)
	// DOM interface
	v.Add(DomSVGAnimateTransformElement)
	Elements[ElemAnimateTransform] = v
	v.Clear()

	// <ElemCircle>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/shapes.html#CircleElement
	// Catagory
	v.AddAll(
		CataBasicShapeElement, CataGraphicsElement,
		CataShapeElement,
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
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrTransform, AttrCx, AttrCy, AttrR,
	)
	// DOM interface
	v.Add(DomSVGCircleElement)
	Elements[ElemCircle] = v
	v.Clear()

	// <clipPath>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/masking.html#ClipPathElement

	// Catagory
	// none
	// Content Model
	v.AddAll(descriptiveElements...)
	v.AddAll(animationElements...)
	v.AddAll(shapeElements...)
	v.AddAll(ElemText, ElemUse)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrTransform, AttrClipPathUnits,
	)
	// DOM interface
	v.Add(DomSVGClipPathElement)
	Elements[ElemClipPath] = v
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
	v.AddAll(AttrLocal, AttrName, AttrRenderingIntent, AttrXlinkHref)
	// DOM interface
	v.Add(DomSVGColorProfileElement)
	Elements[ElemColorProfile] = v
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
		AttrExternalResourcesRequired, AttrX, AttrY,
		AttrXlinkHref,
	)
	// DOM interface
	v.Add(DomSVGCursorElement)
	Elements[ElemCursor] = v
	v.Clear()

	// <defs>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/struct.html#DefsElement
	// Catagory
	v.AddAll(CataContainerElement, CataStructuralElement)
	// Content Model
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	v.AddAll(shapeElements...)
	v.AddAll(structuralElements...)
	v.AddAll(gradientElements...)
	v.AddAll(
		ElemA, ElemAltGlyphDef, ElemClipPath, ElemColorProfile,
		ElemCursor, ElemFilter, ElemFont, ElemFontFace,
		ElemForeignObject, ElemImage, ElemMarker, ElemMask,
		ElemPattern, ElemScript, ElemStyle, ElemSwitch,
		ElemText, ElemView,
	)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrTransform,
	)
	// DOM interface
	v.Add(DomSVGDefsElement)
	Elements[ElemDefs] = v
	v.Clear()

	// <desc>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/struct.html#DescElement
	// Catagory
	v.Add(CataDescriptiveElement)
	// Content Model
	// Any elements or character data.
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(AttrClass, AttrStyle)
	// DOM interface
	v.Add(DomSVGDescElement)
	Elements[ElemDesc] = v
	v.Clear()

	// <ellipse>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/shapes.html#EllipseElement
	// Catagory
	v.AddAll(
		CataBasicShapeElement, CataGraphicsElement,
		CataShapeElement,
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
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrTransform, AttrCx, AttrCy, AttrRx, AttrRy,
	)
	// DOM interface
	v.Add(DomSVGEllipseElement)
	Elements[ElemEllipse] = v
	v.Clear()

	// <feBlend>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feBlendElement
	// Catagory
	v.Add(CataFilterPrimitiveElement)
	// Content Model
	v.AddAll(ElemAnimate, ElemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(AttrClass, AttrStyle, AttrIn, AttrIn2, AttrMode)
	// DOM interface
	v.Add(DomSVGFEBlendElement)
	Elements[ElemFeBlend] = v
	v.Clear()

	// <feColorMatrix>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feColorMatrixElement
	// Catagory
	v.Add(CataFilterPrimitiveElement)
	// Content Model
	v.AddAll(ElemAnimate, ElemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(AttrClass, AttrStyle, AttrIn, AttrType, AttrValues)
	// DOM interface
	v.Add(DomSVGFEColorMatrixElement)
	Elements[ElemFeColorMatrix] = v
	v.Clear()

	// <feComponentTransfer>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feComponentTransferElement
	// Catagory
	v.Add(CataFilterPrimitiveElement)
	// Content Model
	v.AddAll(ElemFeFuncA, ElemFeFuncB, ElemFeFuncG, ElemFeFuncR)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(AttrClass, AttrStyle, AttrIn)
	// DOM interface
	v.Add(DomSVGFEComponentTransferElement)
	Elements[ElemFeComponentTransfer] = v
	v.Clear()

	// <feComposite>
	// Catagory
	v.Add(CataFilterPrimitiveElement)
	// Content Model
	v.AddAll(ElemAnimate, ElemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrIn, AttrIn2, AttrOperator,
		AttrK1, AttrK2, AttrK3, AttrK4,
	)
	// DOM interface
	v.Add(DomSVGFECompositeElement)
	Elements[ElemFeComposite] = v
	v.Clear()

	// <feConvolveMatrix>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feConvolveMatrixElement
	// Catagory
	v.Add(CataFilterPrimitiveElement)
	// Content Model
	v.AddAll(ElemAnimate, ElemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrIn, AttrOrder,
		AttrKernelMatrix, AttrDivisor, AttrBias, AttrTargetX,
		AttrTargetY, AttrEdgeMode, AttrKernelUnitLength,
		AttrPreserveAlpha,
	)
	// DOM interface
	v.Add(DomSVGFEConvolveMatrixElement)
	Elements[ElemFeConvolveMatrix] = v
	v.Clear()

	// <feDiffuseLighting>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feDiffuseLightingElement
	// Catagory
	v.Add(CataFilterPrimitiveElement)
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
		AttrClass, AttrStyle, AttrIn, AttrSurfaceScale,
		AttrDiffuseConstant, AttrKernelUnitLength,
	)
	// DOM interface
	v.Add(DomSVGFEDiffuseLightingElement)
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
	Elements[ElemFeDiffuseLighting] = v
	v.Clear()

	// <feDisplacementMap>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feDisplacementMapElement
	// Catagory
	v.Add(CataFilterPrimitiveElement)
	// Content Model
	v.AddAll(ElemAnimate, ElemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrIn, AttrIn2, AttrScale,
		AttrXChannelSelector, AttrYChannelSelector,
	)
	// DOM interface
	v.Add(DomSVGFEDisplacementMapElement)
	Elements[ElemFeDisplacementMap] = v
	v.Clear()

	// <feDistantLight>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feDistantLightElement
	// Catagory
	v.Add(CataLightSourceElement)
	// Content Model
	// Any number of the following elements, in any order:
	v.AddAll(ElemAnimate, ElemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(AttrAzimuth, AttrElevation)
	// DOM interface
	v.Add(DomSVGFEDistantLightElement)
	Elements[ElemFeDistantLight] = v
	v.Clear()

	// <feFlood>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feFloodElement
	// Catagory
	v.Add(CataFilterPrimitiveElement)
	// Content Model
	// Any number of the following elements, in any order:
	v.AddAll(ElemAnimate, ElemAnimateColor, ElemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(AttrClass, AttrStyle)
	// DOM interface
	v.Add(DomSVGFEFloodElement)
	Elements[ElemFeFlood] = v
	v.Clear()

	// <feFuncA]>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feFuncAElement
	// Catagory
	// none
	// Content Model
	// Any number of the following elements, in any order:
	v.AddAll(ElemAnimate, ElemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(transferFunctionElementAttributes...)
	// DOM interface
	v.Add(DomSVGFEFuncAElement)
	Elements[ElemFeFuncA] = v
	v.Clear()

	// <feFuncB>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feFuncBElement
	// Catagory
	// none
	// Content Model
	// Any number of the following elements, in any order:
	v.AddAll(ElemAnimate, ElemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(transferFunctionElementAttributes...)
	// DOM interface
	v.Add(DomSVGFEFuncBElement)
	Elements[ElemFeFuncB] = v
	v.Clear()

	// <feFuncG>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feFuncGElement
	// Catagory
	// none
	// Content Model
	// Any number of the following elements, in any order:
	v.AddAll(ElemAnimate, ElemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(transferFunctionElementAttributes...)
	// DOM interface
	v.Add(DomSVGFEFuncGElement)
	Elements[ElemFeFuncG] = v
	v.Clear()

	// <feFuncR>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feFuncRElement
	// Catagory
	// none
	// Content Model
	// Any number of the following elements, in any order:
	v.AddAll(ElemAnimate, ElemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(transferFunctionElementAttributes...)
	// DOM interface
	v.Add(DomSVGFEFuncRElement)
	Elements[ElemFeFuncR] = v
	v.Clear()

	// <feGaussianBlur>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feGaussianBlurElement
	// catagory
	v.Add(CataFilterPrimitiveElement)
	// Content Model
	// Any number of the following elements, in any order:
	v.AddAll(ElemAnimate, ElemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(AttrClass, AttrStyle, AttrIn, AttrStdDeviation)
	// dom interface
	v.Add(DomSVGFEGaussianBlurElement)
	Elements[ElemFeGaussianBlur] = v
	v.Clear()

	// <feImage>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feImageElement
	// Catagory
	v.Add(CataFilterPrimitiveElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(ElemAnimate, ElemAnimateTransform, ElemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrPreserveAspectRatio, AttrXlinkHref,
	)
	// DOM interface
	v.Add(DomSVGFEImageElement)
	Elements[ElemFeImage] = v
	v.Clear()

	// <feMerge>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feMergeElement
	// Catagory
	v.Add(CataFilterPrimitiveElement)
	// Content model
	// Any number of the following elements, in any order:
	v.Add(ElemFeMergeNode)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(AttrClass, AttrStyle)
	// DOM interface
	v.Add(DomSVGFEMergeElement)
	Elements[ElemFeMerge] = v
	v.Clear()

	// <feMergeNode>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feMergeNodeElement
	// Catagory
	// none
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(ElemAnimate, ElemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(
		AttrIn,
	)
	// DOM interface
	v.Add(DomSVGFEMergeNodeElement)
	Elements[ElemFeMergeNode] = v
	v.Clear()

	// <feMorphology>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feMorphologyElement
	// Catagory
	v.Add(CataFilterPrimitiveElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(ElemAnimate, ElemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(AttrClass, AttrStyle, AttrIn, AttrOperator, AttrRadius)
	// DOM interface
	v.Add(DomSVGFEMorphologyElement)
	Elements[ElemFeMorphology] = v
	v.Clear()

	// <feOffset>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feOffsetElement
	// Catagory
	v.Add(CataFilterPrimitiveElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(ElemAnimate, ElemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(AttrClass, AttrStyle, AttrIn, AttrDx, AttrDy)
	// DOM interface
	v.Add(DomSVGFEOffsetElement)
	Elements[ElemFeOffset] = v
	v.Clear()

	// <fePointLight>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#fePointLightElement
	// Catagory
	v.Add(CataLightSourceElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(ElemAnimate, ElemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(AttrX, AttrY, AttrZ)
	// DOM interface
	v.Add(DomSVGFEPointLightElement)
	Elements[ElemFePointLight] = v
	v.Clear()

	// <feSpecularLighting>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feSpecularLightingElement
	// Catagory
	v.Add(CataFilterPrimitiveElement)
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
		AttrClass, AttrStyle, AttrIn, AttrSurfaceScale,
		AttrSpecularConstant, AttrSpecularExponent,
		AttrKernelUnitLength,
	)
	// DOM interface
	v.Add(DomSVGFESpecularLightingElement)
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
	Elements[ElemFeSpecularLighting] = v
	v.Clear()

	// <feSpotLight>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feSpotLightElement
	// Catagory
	v.Add(CataLightSourceElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(ElemAnimate, ElemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(
		AttrX, AttrY, AttrZ, AttrPointsAtX, AttrPointsAtY,
		AttrPointsAtZ, AttrSpecularExponent,
		AttrLimitingConeAngle,
	)
	// DOM interface
	v.Add(DomSVGFESpotLightElement)
	Elements[ElemFeSpotLight] = v
	v.Clear()

	// <feTile>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feTileElement
	// Catagory
	v.Add(CataFilterPrimitiveElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(ElemAnimate, ElemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(AttrClass, AttrStyle, AttrIn)
	// DOM interface
	v.Add(DomSVGFETileElement)
	Elements[ElemFeTile] = v
	v.Clear()

	// <feTurbulence>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#feTurbulenceElement
	// Catagory
	v.Add(CataFilterPrimitiveElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(ElemAnimate, ElemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(filterPrimitiveAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrBaseFrequency, AttrNumOctaves,
		AttrSeed, AttrStitchTiles, AttrType,
	)
	// DOM interface
	v.Add(DomSVGFETurbulenceElement)
	Elements[ElemFeTurbulence] = v
	v.Clear()

	// <filter>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/filters.html#FilterElement
	// Catagory
	// none
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(descriptiveElements...)
	v.AddAll(filterPrimiveElements...)
	v.AddAll(ElemAnimate, ElemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrX, AttrY, AttrWidth, AttrHeight, AttrFilterRes,
		AttrFilterUnits, AttrPrimitiveUnits, AttrXlinkHref,
	)
	// DOM interface
	v.Add(DomSVGFilterElement)
	Elements[ElemFilter] = v
	v.Clear()

	// <font>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/fonts.html#FontElement
	// Catagory
	// none
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(descriptiveElements...)
	v.AddAll(
		ElemFontFace, ElemGlyph, ElemHkern, ElemMissingGlyph,
		ElemVkern,
	)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrHorizOriginX, AttrHorizOriginY, AttrHorizAdvX,
		AttrVertOriginX, AttrVertOriginY, AttrVertAdvY,
	)
	// DOM interface
	v.Add(DomSVGFontElement)
	Elements[ElemFont] = v
	v.Clear()

	// <font-face>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/fonts.html#FontFaceElement
	// Catagory
	// none
	// Content model
	// Any number of descriptive elements and at most one
	// ‘font-face-src’ element, in any order.
	v.AddAll(descriptiveElements...)
	v.Add(ElemFontFaceSrc)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(
		AttrFontFamily, AttrFontStyle, AttrFontVariant,
		AttrFontWeight, AttrFontStretch, AttrFontSize,
		AttrUnicodeRange, AttrUnitsPerEm, AttrPanose1,
		AttrStemv, AttrStemh, AttrSlope, AttrCapHeight,
		AttrXHeight, AttrAccentHeight, AttrAscent, AttrDescent,
		AttrWidths, AttrBbox, AttrIdeographic, AttrAlphabetic,
		AttrMathematical, AttrHanging, AttrVIdeographic,
		AttrVAlphabetic, AttrVMathematical, AttrVHanging,
		AttrUnderlinePosition, AttrUnderlineThickness,
		AttrStrikethroughPosition, AttrStrikethroughThickness,
		AttrOverlinePosition, AttrOverlineThickness,
	)
	// DOM interface
	v.Add(DomSVGFontFaceElement)
	v.contract = append(v.contract, func(g group) error {
		// Any number of descriptive elements and at most one
		// ‘font-face-src’ element, in any order.
		return nil
	})
	Elements[ElemFontFace] = v
	v.Clear()

	// <font-face-format>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/fonts.html#FontFaceFormatElement
	// Catagory
	// none
	// Content model
	// empty
	// Attributes
	v.AddAll(coreAttributes...)
	v.Add(AttrString)
	// DOM interface
	v.Add(DomSVGFontFaceFormatElement)
	v.contract = append(v.contract, func(g group) error {
		// should be empty
		return nil
	})
	Elements[ElemFontFaceFormat] = v
	v.Clear()

	// <font-face-name>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/fonts.html#FontFaceNameElement
	// Catagory
	// none
	// Content model
	// empty
	// Attributes
	v.AddAll(coreAttributes...)
	v.Add(AttrName)
	// DOM interface
	v.Add(DomSVGFontFaceNameElement)
	v.contract = append(v.contract, func(g group) error {
		// empty
		return nil
	})
	Elements[ElemFontFaceName] = v
	v.Clear()

	// <font-face-src>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/fonts.html#FontFaceSrcElement
	// Catagory
	// none
	// Content model
	// One or more of the following elements, in any order:
	v.AddAll(ElemFontFaceName, ElemFontFaceUri)
	// Attributes
	v.AddAll(coreAttributes...)
	// DOM interface
	v.Add(DomSVGFontFaceSrcElement)
	v.contract = append(v.contract, func(g group) error {
		// One or more of the following elements, in any order:
		//	ElemFontFaceName, ElemFontFaceUri
		return nil
	})
	Elements[ElemFontFaceSrc] = v
	v.Clear()

	// <font-face-uri>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/fonts.html#FontFaceURIElement
	// Catagory
	// none
	// Content model
	// Any number of the following elements, in any order:
	v.Add(ElemFontFaceFormat)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(xlinkAttributes...)
	v.Add(AttrXlinkHref)
	// DOM interface
	v.AddAll(DomSVGFontFaceUriElement)
	Elements[ElemFontFaceUri] = v
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
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrTransform, AttrX, AttrY, AttrWidth, AttrHeight,
	)
	// DOM interface
	v.Add(DomSVGForeignObjectElement)
	Elements[ElemForeignObject] = v
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
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrTransform, AttrX, AttrY, AttrWidth, AttrHeight,
	)
	// DOM interface
	v.Add(DomSVGForeignObjectElement)
	Elements[ElemForeignObject] = v
	v.Clear()

	// <g>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/struct.html#GElement
	// Catagory
	v.AddAll(CataContainerElement, CataStructuralElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	v.AddAll(shapeElements...)
	v.AddAll(structuralElements...)
	v.AddAll(gradientElements...)
	v.AddAll(
		ElemA, ElemAltGlyphDef, ElemClipPath, ElemColorProfile,
		ElemCursor, ElemFilter, ElemFont, ElemFontFace,
		ElemForeignObject, ElemImage, ElemMarker, ElemMask,
		ElemPattern, ElemScript, ElemStyle, ElemSwitch,
		ElemText, ElemView,
	)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrTransform,
	)
	// DOM interface
	v.Add(DomSVGGElement)
	Elements[ElemG] = v
	v.Clear()

	// <glyph>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/fonts.html#GlyphElement
	// Catagory
	v.Add(CataContainerElement)
	//
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	v.AddAll(shapeElements...)
	v.AddAll(structuralElements...)
	v.AddAll(gradientElements...)
	v.AddAll(
		ElemA, ElemAltGlyphDef, ElemClipPath, ElemColorProfile,
		ElemCursor, ElemFilter, ElemFont, ElemFontFace,
		ElemForeignObject, ElemImage, ElemMarker, ElemMask,
		ElemPattern, ElemScript, ElemStyle, ElemSwitch,
		ElemText, ElemView,
	)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrD, AttrHorizAdvX,
		AttrVertOriginX, AttrVertOriginY, AttrVertAdvY,
		AttrUnicode, AttrGlyphName, AttrOrientation,
		AttrArabicForm, AttrLang,
	)
	// DOM interface
	v.Add(DomSVGGlyphElement)
	Elements[ElemGlyph] = v
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
		AttrClass, AttrStyle, AttrX, AttrY, AttrDx, AttrDy,
		AttrGlyphRef, AttrFormat, AttrXlinkHref,
	)
	// DOM interface
	v.Add(DomSVGGlyphRefElement)
	Elements[ElemGlyphRef] = v
	v.Clear()

	// <hkern>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/fonts.html#HKernElement
	// Catagory
	// none
	// Content model
	// empty
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(AttrU1, AttrG1, AttrU2, AttrG2, AttrK)
	// DOM interface
	v.Add(DomSVGHKernElement)
	Elements[ElemHkern] = v
	v.Clear()

	// <image>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/struct.html#ImageElement
	// Catagory
	v.AddAll(CataGraphicsElement, CataGraphicsReferencingElement)
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
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrPreserveAspectRatio, AttrTransform, AttrX, AttrY,
		AttrWidth, AttrHeight, AttrXlinkHref,
	)
	// DOM interface
	v.Add(DomSVGImageElement)
	Elements[ElemImage] = v
	v.Clear()

	// <line>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/shapes.html#LineElement
	// Catagory
	v.AddAll(
		CataBasicShapeElement, CataGraphicsElement,
		CataShapeElement,
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
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrTransform, AttrX1, AttrY1, AttrX2, AttrY2,
	)
	// DOM interface
	v.Add(DomSVGLineElement)
	v.contract = append(v.contract, func(g group) error {
		// Any number of the following elements, in any order:
		// descriptive elements
		// animation elements
		return nil
	})
	Elements[ElemLine] = v
	v.Clear()

	// <linearGradient>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/pservers.html#LinearGradientElement
	// Catagory
	v.Add(CataGradientElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(descriptiveElements...)
	v.AddAll(ElemAnimate, ElemAnimateTransform, ElemSet, ElemStop)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrX1, AttrY1, AttrX2, AttrY2, AttrGradientUnits,
		AttrGradientTransform, AttrSpreadMethod, AttrXlinkHref,
	)
	// DOM interface
	v.Add(DomSVGLinearGradientElement)
	v.contract = append(v.contract, func(g group) error {
		// Any number of the following elements, in any order:
		// descriptive elements
		// ‘animate’ ‘animateTransform’ ‘set’ ‘stop’
		return nil
	})
	Elements[ElemLinearGradient] = v
	v.Clear()

	// <marker>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/painting.html#MarkerElement
	// Catagory
	v.Add(CataContainerElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	v.AddAll(shapeElements...)
	v.AddAll(structuralElements...)
	v.AddAll(gradientElements...)
	v.AddAll(
		ElemA, ElemAltGlyphDef, ElemClipPath, ElemColorProfile,
		ElemCursor, ElemFilter, ElemFont, ElemFontFace,
		ElemForeignObject, ElemImage, ElemMarker, ElemMask,
		ElemPattern, ElemScript, ElemStyle, ElemSwitch,
		ElemText, ElemView,
	)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrViewBox, AttrPreserveAspectRatio, AttrRefX,
		AttrRefY, AttrMarkerUnits, AttrMarkerWidth,
		AttrMarkerHeight, AttrOrient,
	)
	// DOM interface
	v.Add(DomSVGMarkerElement)
	v.contract = append(v.contract, func(g group) error {
		// Any number of the following elements, in any order:
		// animationElements
		// descriptiveElements
		// shapeElements
		// structuralElements
		// gradientElements
		//
		// 	ElemA, ElemAltGlyphDef, ElemClipPath,
		// 	ElemColorProfile, ElemCursor, ElemFilter,
		// 	ElemFont, ElemFontFace, ElemForeignObject,
		// 	ElemImage, ElemMarker, ElemMask, ElemPattern,
		// 	ElemScript, ElemStyle, ElemSwitch, ElemText,
		// 	ElemView,
		return nil
	})
	Elements[ElemMarker] = v
	v.Clear()

	// <mask>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/masking.html#MaskElement
	// Catagory
	v.Add(CataContainerElement)
	//
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	v.AddAll(shapeElements...)
	v.AddAll(structuralElements...)
	v.AddAll(gradientElements...)
	v.AddAll(
		ElemA, ElemAltGlyphDef, ElemClipPath, ElemColorProfile,
		ElemCursor, ElemFilter, ElemFont, ElemFontFace,
		ElemForeignObject, ElemImage, ElemMarker, ElemMask,
		ElemPattern, ElemScript, ElemStyle, ElemSwitch,
		ElemText, ElemView,
	)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrX, AttrY, AttrWidth, AttrHeight, AttrMaskUnits,
		AttrMaskContentUnits,
	)
	// DOM interface
	v.Add(DomSVGMaskElement)
	Elements[ElemMask] = v
	v.Clear()

	// <metadata>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/metadata.html#MetadataElement
	// Catagory
	v.Add(CataDescriptiveElement)
	//
	// Content model
	// Any elements or character data.
	// Attributes
	v.AddAll(coreAttributes...)
	// DOM interface
	v.Add(DomSVGMetadataElement)
	Elements[ElemMetadata] = v
	v.Clear()

	// <missing-glyph>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/fonts.html#MissingGlyphElement
	// Catagory
	v.Add(CataContainerElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	v.AddAll(shapeElements...)
	v.AddAll(structuralElements...)
	v.AddAll(gradientElements...)
	v.AddAll(
		ElemA, ElemAltGlyphDef, ElemClipPath, ElemColorProfile,
		ElemCursor, ElemFilter, ElemFont, ElemFontFace,
		ElemForeignObject, ElemImage, ElemMarker, ElemMask,
		ElemPattern, ElemScript, ElemStyle, ElemSwitch,
		ElemText, ElemView,
	)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrD, AttrHorizAdvX,
		AttrVertOriginX, AttrVertOriginY, AttrVertAdvY,
	)
	// DOM interface
	v.Add(DomSVGMissingGlyphElement)
	Elements[ElemMissingGlyph] = v
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
	v.AddAll(AttrExternalResourcesRequired, AttrXlinkHref)
	// DOM interface
	v.Add(DomSVGMPathElement)
	Elements[ElemMpath] = v
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
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrTransform, AttrD,
	)
	// DOM interface
	v.Add(DomSVGPathElement)
	Elements[ElemPath] = v
	v.Clear()

	// <pattern>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/pservers.html#PatternElement
	// Catagory
	v.Add(CataContainerElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	v.AddAll(shapeElements...)
	v.AddAll(structuralElements...)
	v.AddAll(gradientElements...)
	v.AddAll(
		ElemA, ElemAltGlyphDef, ElemClipPath, ElemColorProfile,
		ElemCursor, ElemFilter, ElemFont, ElemFontFace,
		ElemForeignObject, ElemImage, ElemMarker, ElemMask,
		ElemPattern, ElemScript, ElemStyle, ElemSwitch,
		ElemText, ElemView,
	)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrViewBox, AttrPreserveAspectRatio, AttrX, AttrY,
		AttrWidth, AttrHeight, AttrPatternUnits,
		AttrPatternContentUnits, AttrPatternTransform,
		AttrXlinkHref,
	)
	// DOM interface
	v.Add(DomSVGPatternElement)
	Elements[ElemPattern] = v
	v.Clear()

	// <polygon>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/shapes.html#PolygonElement
	// Catagory
	v.AddAll(
		CataBasicShapeElement, CataGraphicsElement,
		CataShapeElement,
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
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrTransform, AttrPoints,
	)
	// DOM interface
	v.Add(DomSVGPolygonElement)
	Elements[ElemPolygon] = v
	v.Clear()

	// <polyline>
	// Catagory
	v.AddAll(
		CataBasicShapeElement, CataGraphicsElement,
		CataShapeElement,
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
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrTransform, AttrPoints,
	)
	// DOM interface
	v.Add(DomSVGPolylineElement)
	Elements[ElemPolyline] = v
	v.Clear()

	// <radialGradient>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/pservers.html#RadialGradientElement
	// Catagory
	v.Add(CataGradientElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(descriptiveElements...)
	v.AddAll(ElemAnimate, ElemAnimateTransform, ElemSet, ElemStop)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrCx, AttrCy, AttrR, AttrFx, AttrFy,
		AttrGradientUnits, AttrGradientTransform,
		AttrSpreadMethod, AttrXlinkHref,
	)
	// DOM interface
	v.Add(DomSVGRadialGradientElement)
	Elements[ElemRadialGradient] = v
	v.Clear()

	// <rect>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/shapes.html#RectElement
	// Catagory
	v.AddAll(
		CataBasicShapeElement, CataGraphicsElement,
		CataShapeElement,
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
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrTransform, AttrX, AttrY, AttrWidth, AttrHeight,
		AttrRx, AttrRy,
	)
	// DOM interface
	v.Add(DomSVGRectElement)
	Elements[ElemRect] = v
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
	v.AddAll(AttrExternalResourcesRequired, AttrType, AttrXlinkHref)
	// DOM interface
	v.Add(DomSVGScriptElement)
	Elements[ElemScript] = v
	v.Clear()

	// <set>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/animate.html#SetElement
	// Catagory
	v.Add(CataAnimationElement)
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
		AttrExternalResourcesRequired,
		AttrTo,
	)
	// DOM interface
	v.Add(DomSVGSetElement)
	Elements[ElemSet] = v
	v.Clear()

	// <stop>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/pservers.html#StopElement
	// Catagory
	// none
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(ElemAnimate, ElemAnimateColor, ElemSet)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(AttrClass, AttrStyle, AttrOffset)
	// DOM interface
	v.Add(DomSVGStopElement)
	Elements[ElemStop] = v
	v.Clear()

	// <style>
	// Catagory
	// none
	// Content model
	// Any elements or character data.
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(AttrType, AttrMedia, AttrTitle)
	// DOM interface
	v.Add(DomSVGStyleElement)
	Elements[ElemStyle] = v
	v.Clear()

	// <svg>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/styling.html#StyleElement
	// Catagory
	v.AddAll(CataContainerElement, CataStructuralElement)
	//
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	v.AddAll(shapeElements...)
	v.AddAll(structuralElements...)
	v.AddAll(gradientElements...)
	v.AddAll(
		ElemA, ElemAltGlyphDef, ElemClipPath, ElemColorProfile,
		ElemCursor, ElemFilter, ElemFont, ElemFontFace,
		ElemForeignObject, ElemImage, ElemMarker, ElemMask,
		ElemPattern, ElemScript, ElemStyle, ElemSwitch,
		ElemText, ElemView,
	)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(documentEventAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrX, AttrY, AttrWidth, AttrHeight, AttrViewBox,
		AttrPreserveAspectRatio, AttrZoomAndPan, AttrVersion,
		AttrContentScriptType, AttrContentStyleType, AttrX,
		AttrY, AttrWidth, AttrHeight, AttrVersion,
		AttrBaseProfile,
	)
	// DOM interface
	v.Add(DomSVGSVGElement)
	Elements[ElemSvg] = v
	v.Clear()

	// <switch>
	// Catagory
	v.Add(CataContainerElement)
	//
	// Content model
	// Any number of the following elements, in any order:
	// animation elements
	// descriptive elements
	// shape elements
	v.AddAll(
		ElemA, ElemForeignObject, ElemG, ElemImage, ElemSvg,
		ElemSwitch, ElemText, ElemUse,
	)
	//
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrTransform,
	)
	// DOM interface
	v.Add(DomSVGSwitchElement)
	Elements[ElemSwitch] = v
	v.Clear()

	// <symbol>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/struct.html#SymbolElement
	// Catagory
	v.AddAll(CataContainerElement, CataStructuralElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	v.AddAll(shapeElements...)
	v.AddAll(structuralElements...)
	v.AddAll(gradientElements...)
	v.AddAll(
		ElemA, ElemAltGlyphDef, ElemClipPath, ElemColorProfile,
		ElemCursor, ElemFilter, ElemFont, ElemFontFace,
		ElemForeignObject, ElemImage, ElemMarker, ElemMask,
		ElemPattern, ElemScript, ElemStyle, ElemSwitch,
		ElemText, ElemView,
	)
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrPreserveAspectRatio, AttrViewBox,
	)
	// DOM interface
	v.Add(DomSVGSymbolElement)
	Elements[ElemSymbol] = v
	v.Clear()

	// <text>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/text.html#TextElement
	// Catagory
	v.AddAll(CataGraphicsElement, CataTextContentElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(animationElements...)
	v.AddAll(descriptiveElements...)
	v.AddAll(CataTextContentChildElement)
	v.Add(ElemA)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrTransform, AttrLengthAdjust, AttrX, AttrY, AttrDx,
		AttrDy, AttrRotate, AttrTextLength,
	)
	// DOM interface
	v.Add(DomSVGTextElement)
	Elements[ElemText] = v
	v.Clear()

	// <textPath>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/text.html#TextPathElement
	// Catagory
	v.AddAll(CataTextContentElement, CataTextContentChildElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(descriptiveElements...)
	v.AddAll(
		ElemA, ElemAltGlyph, ElemAnimate, ElemAnimateColor,
		ElemSet, ElemTref, ElemTspan,
	)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrXlinkHref, AttrStartOffset, AttrMethod, AttrSpacing,
	)
	// DOM interface
	v.Add(DomSVGTextPathElement)
	Elements[ElemTextPath] = v
	v.Clear()

	// <title>
	// Catagory
	v.Add(CataDescriptiveElement)
	// Content model
	// Any elements or character data.
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(AttrClass, AttrStyle)
	// DOM interface
	v.Add(DomSVGTitleElement)
	Elements[ElemTitle] = v
	v.Clear()

	// <tref>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/text.html#TRefElement
	// Catagory
	v.AddAll(CataTextContentElement, CataTextContentChildElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(descriptiveElements...)
	v.AddAll(ElemAnimate, ElemAnimateColor, ElemSet)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(xlinkAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrXlinkHref,
	)
	// DOM interface
	v.Add(DomSVGTRefElement)
	Elements[ElemTref] = v
	v.Clear()

	// <tspan>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/text.html#TSpanElement
	// Catagory
	v.AddAll(CataTextContentElement, CataTextContentChildElement)
	// Content model
	// Any number of the following elements, in any order:
	v.AddAll(descriptiveElements...)
	v.AddAll(
		ElemA, ElemAltGlyph, ElemAnimate, ElemAnimateColor,
		ElemSet, ElemTref, ElemTspan,
	)
	// Attributes
	v.AddAll(conditionalProcessingAttributes...)
	v.AddAll(coreAttributes...)
	v.AddAll(graphicalEventAttributes...)
	v.AddAll(presentationAttributes...)
	v.AddAll(
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrX, AttrY, AttrDx, AttrDy, AttrRotate,
		AttrTextLength, AttrLengthAdjust,
	)
	// DOM interface
	v.Add(DomSVGTSpanElement)
	Elements[ElemTspan] = v
	v.Clear()

	// <use>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/struct.html#UseElement
	// Catagory
	v.AddAll(
		CataGraphicsElement, CataGraphicsReferencingElement,
		CataStructuralElement,
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
		AttrClass, AttrStyle, AttrExternalResourcesRequired,
		AttrTransform, AttrX, AttrY, AttrWidth, AttrHeight,
		AttrXlinkHref,
	)
	// DOM interface
	v.Add(DomSVGUseElement)
	Elements[ElemUse] = v
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
		AttrExternalResourcesRequired, AttrViewBox,
		AttrPreserveAspectRatio, AttrZoomAndPan, AttrViewTarget,
	)
	// DOM interface
	v.Add(DomSVGViewElement)
	Elements[ElemView] = v
	v.Clear()

	// <vkern>
	// https://dev.w3.org/SVG/profiles/1.1F2/publish/fonts.html#VKernElement
	// Catagory
	// none
	// Content model
	// empty
	// Attributes
	v.AddAll(coreAttributes...)
	v.AddAll(AttrU1, AttrG1, AttrU2, AttrG2, AttrK)
	// DOM interface
	v.Add(DomSVGHKernElement)
	Elements[ElemVkern] = v
	v.Clear()
}
