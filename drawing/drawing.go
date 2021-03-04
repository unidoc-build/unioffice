//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package drawing ;import (_a "github.com/unidoc/unioffice";_e "github.com/unidoc/unioffice/color";_aa "github.com/unidoc/unioffice/measurement";_g "github.com/unidoc/unioffice/schema/soo/dml";);

// SetJoin sets the line join style.
func (_ef LineProperties )SetJoin (e LineJoin ){_ef ._c .Round =nil ;_ef ._c .Miter =nil ;_ef ._c .Bevel =nil ;switch e {case LineJoinRound :_ef ._c .Round =_g .NewCT_LineJoinRound ();case LineJoinBevel :_ef ._c .Bevel =_g .NewCT_LineJoinBevel ();case LineJoinMiter :_ef ._c .Miter =_g .NewCT_LineJoinMiterProperties ();};};

// SetHeight sets the height of the shape.
func (_fcb ShapeProperties )SetHeight (h _aa .Distance ){_fcb .ensureXfrm ();if _fcb ._bgb .Xfrm .Ext ==nil {_fcb ._bgb .Xfrm .Ext =_g .NewCT_PositiveSize2D ();};_fcb ._bgb .Xfrm .Ext .CyAttr =int64 (h /_aa .EMU );};

// X returns the inner wrapped XML type.
func (_cde Run )X ()*_g .EG_TextRun {return _cde ._fe };

// SetSize sets the width and height of the shape.
func (_aeg ShapeProperties )SetSize (w ,h _aa .Distance ){_aeg .SetWidth (w );_aeg .SetHeight (h )};func (_cf LineProperties )SetNoFill (){_cf .clearFill ();_cf ._c .NoFill =_g .NewCT_NoFillProperties ()};

// SetFlipHorizontal controls if the shape is flipped horizontally.
func (_da ShapeProperties )SetFlipHorizontal (b bool ){_da .ensureXfrm ();if !b {_da ._bgb .Xfrm .FlipHAttr =nil ;}else {_da ._bgb .Xfrm .FlipHAttr =_a .Bool (true );};};

// AddBreak adds a new line break to a paragraph.
func (_cd Paragraph )AddBreak (){_cbg :=_g .NewEG_TextRun ();_cbg .Br =_g .NewCT_TextLineBreak ();_cd ._bc .EG_TextRun =append (_cd ._bc .EG_TextRun ,_cbg );};

// RunProperties controls the run properties.
type RunProperties struct{_fc *_g .CT_TextCharacterProperties ;};func (_d LineProperties )clearFill (){_d ._c .NoFill =nil ;_d ._c .GradFill =nil ;_d ._c .SolidFill =nil ;_d ._c .PattFill =nil ;};

// AddRun adds a new run to a paragraph.
func (_agd Paragraph )AddRun ()Run {_bd :=MakeRun (_g .NewEG_TextRun ());_agd ._bc .EG_TextRun =append (_agd ._bc .EG_TextRun ,_bd .X ());return _bd ;};func (_aaf ShapeProperties )ensureXfrm (){if _aaf ._bgb .Xfrm ==nil {_aaf ._bgb .Xfrm =_g .NewCT_Transform2D ();};};

// MakeParagraphProperties constructs a new ParagraphProperties wrapper.
func MakeParagraphProperties (x *_g .CT_TextParagraphProperties )ParagraphProperties {return ParagraphProperties {x };};

// X returns the inner wrapped XML type.
func (_ag Paragraph )X ()*_g .CT_TextParagraph {return _ag ._bc };

// SetNumbered controls if bullets are numbered or not.
func (_ada ParagraphProperties )SetNumbered (scheme _g .ST_TextAutonumberScheme ){if scheme ==_g .ST_TextAutonumberSchemeUnset {_ada ._bg .BuAutoNum =nil ;}else {_ada ._bg .BuAutoNum =_g .NewCT_TextAutonumberBullet ();_ada ._bg .BuAutoNum .TypeAttr =scheme ;};};

// GetPosition gets the position of the shape in EMU.
func (_dg ShapeProperties )GetPosition ()(int64 ,int64 ){_dg .ensureXfrm ();if _dg ._bgb .Xfrm .Off ==nil {_dg ._bgb .Xfrm .Off =_g .NewCT_Point2D ();};return *_dg ._bgb .Xfrm .Off .XAttr .ST_CoordinateUnqualified ,*_dg ._bgb .Xfrm .Off .YAttr .ST_CoordinateUnqualified ;};type LineProperties struct{_c *_g .CT_LineProperties };

// X returns the inner wrapped XML type.
func (_ec ParagraphProperties )X ()*_g .CT_TextParagraphProperties {return _ec ._bg };

// X returns the inner wrapped XML type.
func (_aae LineProperties )X ()*_g .CT_LineProperties {return _aae ._c };

// SetBold controls the bolding of a run.
func (_ff RunProperties )SetBold (b bool ){_ff ._fc .BAttr =_a .Bool (b )};type ShapeProperties struct{_bgb *_g .CT_ShapeProperties };

// Properties returns the paragraph properties.
func (_cbgg Paragraph )Properties ()ParagraphProperties {if _cbgg ._bc .PPr ==nil {_cbgg ._bc .PPr =_g .NewCT_TextParagraphProperties ();};return MakeParagraphProperties (_cbgg ._bc .PPr );};

// X returns the inner wrapped XML type.
func (_bf ShapeProperties )X ()*_g .CT_ShapeProperties {return _bf ._bgb };

// SetText sets the run's text contents.
func (_af Run )SetText (s string ){_af ._fe .Br =nil ;_af ._fe .Fld =nil ;if _af ._fe .R ==nil {_af ._fe .R =_g .NewCT_RegularTextRun ();};_af ._fe .R .T =s ;};func (_f LineProperties )SetSolidFill (c _e .Color ){_f .clearFill ();_f ._c .SolidFill =_g .NewCT_SolidColorFillProperties ();_f ._c .SolidFill .SrgbClr =_g .NewCT_SRgbColor ();_f ._c .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// MakeRunProperties constructs a new RunProperties wrapper.
func MakeRunProperties (x *_g .CT_TextCharacterProperties )RunProperties {return RunProperties {x }};

// SetLevel sets the level of indentation of a paragraph.
func (_efe ParagraphProperties )SetLevel (idx int32 ){_efe ._bg .LvlAttr =_a .Int32 (idx )};

// SetBulletChar sets the bullet character for the paragraph.
func (_cg ParagraphProperties )SetBulletChar (c string ){if c ==""{_cg ._bg .BuChar =nil ;}else {_cg ._bg .BuChar =_g .NewCT_TextCharBullet ();_cg ._bg .BuChar .CharAttr =c ;};};

// SetWidth sets the width of the shape.
func (_fb ShapeProperties )SetWidth (w _aa .Distance ){_fb .ensureXfrm ();if _fb ._bgb .Xfrm .Ext ==nil {_fb ._bgb .Xfrm .Ext =_g .NewCT_PositiveSize2D ();};_fb ._bgb .Xfrm .Ext .CxAttr =int64 (w /_aa .EMU );};

// Paragraph is a paragraph within a document.
type Paragraph struct{_bc *_g .CT_TextParagraph };

// LineJoin is the type of line join
type LineJoin byte ;const (LineJoinRound LineJoin =iota ;LineJoinBevel ;LineJoinMiter ;);

// SetBulletFont controls the font for the bullet character.
func (_ad ParagraphProperties )SetBulletFont (f string ){if f ==""{_ad ._bg .BuFont =nil ;}else {_ad ._bg .BuFont =_g .NewCT_TextFont ();_ad ._bg .BuFont .TypefaceAttr =f ;};};func (_feb ShapeProperties )SetSolidFill (c _e .Color ){_feb .clearFill ();_feb ._bgb .SolidFill =_g .NewCT_SolidColorFillProperties ();_feb ._bgb .SolidFill .SrgbClr =_g .NewCT_SRgbColor ();_feb ._bgb .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// SetPosition sets the position of the shape.
func (_ga ShapeProperties )SetPosition (x ,y _aa .Distance ){_ga .ensureXfrm ();if _ga ._bgb .Xfrm .Off ==nil {_ga ._bgb .Xfrm .Off =_g .NewCT_Point2D ();};_ga ._bgb .Xfrm .Off .XAttr .ST_CoordinateUnqualified =_a .Int64 (int64 (x /_aa .EMU ));_ga ._bgb .Xfrm .Off .YAttr .ST_CoordinateUnqualified =_a .Int64 (int64 (y /_aa .EMU ));};

// MakeRun constructs a new Run wrapper.
func MakeRun (x *_g .EG_TextRun )Run {return Run {x }};

// Run is a run within a paragraph.
type Run struct{_fe *_g .EG_TextRun };

// SetGeometry sets the shape type of the shape
func (_fcg ShapeProperties )SetGeometry (g _g .ST_ShapeType ){if _fcg ._bgb .PrstGeom ==nil {_fcg ._bgb .PrstGeom =_g .NewCT_PresetGeometry2D ();};_fcg ._bgb .PrstGeom .PrstAttr =g ;};func MakeShapeProperties (x *_g .CT_ShapeProperties )ShapeProperties {return ShapeProperties {x }};

// SetWidth sets the line width, MS products treat zero as the minimum width
// that can be displayed.
func (_cb LineProperties )SetWidth (w _aa .Distance ){_cb ._c .WAttr =_a .Int32 (int32 (w /_aa .EMU ))};

// ParagraphProperties allows controlling paragraph properties.
type ParagraphProperties struct{_bg *_g .CT_TextParagraphProperties ;};

// SetSolidFill controls the text color of a run.
func (_fg RunProperties )SetSolidFill (c _e .Color ){_fg ._fc .NoFill =nil ;_fg ._fc .BlipFill =nil ;_fg ._fc .GradFill =nil ;_fg ._fc .GrpFill =nil ;_fg ._fc .PattFill =nil ;_fg ._fc .SolidFill =_g .NewCT_SolidColorFillProperties ();_fg ._fc .SolidFill .SrgbClr =_g .NewCT_SRgbColor ();_fg ._fc .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// SetAlign controls the paragraph alignment
func (_de ParagraphProperties )SetAlign (a _g .ST_TextAlignType ){_de ._bg .AlgnAttr =a };

// SetFont controls the font of a run.
func (_ea RunProperties )SetFont (s string ){_ea ._fc .Latin =_g .NewCT_TextFont ();_ea ._fc .Latin .TypefaceAttr =s ;};

// Properties returns the run's properties.
func (_bge Run )Properties ()RunProperties {if _bge ._fe .R ==nil {_bge ._fe .R =_g .NewCT_RegularTextRun ();};if _bge ._fe .R .RPr ==nil {_bge ._fe .R .RPr =_g .NewCT_TextCharacterProperties ();};return RunProperties {_bge ._fe .R .RPr };};

// SetFlipVertical controls if the shape is flipped vertically.
func (_fcf ShapeProperties )SetFlipVertical (b bool ){_fcf .ensureXfrm ();if !b {_fcf ._bgb .Xfrm .FlipVAttr =nil ;}else {_fcf ._bgb .Xfrm .FlipVAttr =_a .Bool (true );};};func (_ae ShapeProperties )clearFill (){_ae ._bgb .NoFill =nil ;_ae ._bgb .BlipFill =nil ;_ae ._bgb .GradFill =nil ;_ae ._bgb .GrpFill =nil ;_ae ._bgb .SolidFill =nil ;_ae ._bgb .PattFill =nil ;};

// MakeParagraph constructs a new paragraph wrapper.
func MakeParagraph (x *_g .CT_TextParagraph )Paragraph {return Paragraph {x }};func (_add ShapeProperties )LineProperties ()LineProperties {if _add ._bgb .Ln ==nil {_add ._bgb .Ln =_g .NewCT_LineProperties ();};return LineProperties {_add ._bgb .Ln };};func (_db ShapeProperties )SetNoFill (){_db .clearFill ();_db ._bgb .NoFill =_g .NewCT_NoFillProperties ();};

// SetSize sets the font size of the run text
func (_ce RunProperties )SetSize (sz _aa .Distance ){_ce ._fc .SzAttr =_a .Int32 (int32 (sz /_aa .HundredthPoint ));};