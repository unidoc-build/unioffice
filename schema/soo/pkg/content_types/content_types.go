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

package content_types ;import (_ga "encoding/xml";_ad "fmt";_f "github.com/unidoc/unioffice";_a "regexp";);func NewCT_Default ()*CT_Default {_fc :=&CT_Default {};_fc .ExtensionAttr ="\u0078\u006d\u006c";_fc .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";return _fc ;};

// Validate validates the Default and its children
func (_dgae *Default )Validate ()error {return _dgae .ValidateWithPath ("\u0044e\u0066\u0061\u0075\u006c\u0074");};type Types struct{CT_Types };func NewOverride ()*Override {_gac :=&Override {};_gac .CT_Override =*NewCT_Override ();return _gac };func NewTypes ()*Types {_eg :=&Types {};_eg .CT_Types =*NewCT_Types ();return _eg };

// ValidateWithPath validates the CT_Default and its children, prefixing error messages with path
func (_cc *CT_Default )ValidateWithPath (path string )error {if !ST_ExtensionPatternRe .MatchString (_cc .ExtensionAttr ){return _ad .Errorf ("\u0025s\u002f\u006d.\u0045\u0078\u0074\u0065n\u0073\u0069\u006fn\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074 m\u0061\u0074\u0063h\u0020\u0027%\u0073\u0027\u0020\u0028\u0068\u0061v\u0065\u0020%\u0076\u0029",path ,ST_ExtensionPatternRe ,_cc .ExtensionAttr );};if !ST_ContentTypePatternRe .MatchString (_cc .ContentTypeAttr ){return _ad .Errorf ("\u0025\u0073/\u006d\u002e\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0027\u0025\u0073\u0027\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029",path ,ST_ContentTypePatternRe ,_cc .ContentTypeAttr );};return nil ;};var ST_ExtensionPatternRe =_a .MustCompile (ST_ExtensionPattern );func (_adg *Default )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {return _adg .CT_Default .MarshalXML (e ,start );};func (_ea *CT_Override )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {_ea .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";for _ ,_cb :=range start .Attr {if _cb .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_af ,_ae :=_cb .Value ,error (nil );if _ae !=nil {return _ae ;};_ea .ContentTypeAttr =_af ;continue ;};if _cb .Name .Local =="\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065"{_eef ,_gb :=_cb .Value ,error (nil );if _gb !=nil {return _gb ;};_ea .PartNameAttr =_eef ;continue ;};};for {_bd ,_de :=d .Token ();if _de !=nil {return _ad .Errorf ("\u0070\u0061\u0072si\u006e\u0067\u0020\u0043\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u003a\u0020\u0025\u0073",_de );};if _bg ,_ebg :=_bd .(_ga .EndElement );_ebg &&_bg .Name ==start .Name {break ;};};return nil ;};func (_aaa *CT_Types )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {_fa :for {_cd ,_dac :=d .Token ();if _dac !=nil {return _dac ;};switch _ab :=_cd .(type ){case _ga .StartElement :switch _ab .Name {case _ga .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u0044e\u0066\u0061\u0075\u006c\u0074"}:_ge :=NewDefault ();if _afa :=d .DecodeElement (_ge ,&_ab );_afa !=nil {return _afa ;};_aaa .Default =append (_aaa .Default ,_ge );case _ga .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}:_bb :=NewOverride ();if _ec :=d .DecodeElement (_bb ,&_ab );_ec !=nil {return _ec ;};_aaa .Override =append (_aaa .Override ,_bb );default:_f .Log ("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073\u0020\u0025\u0076",_ab .Name );if _deg :=d .Skip ();_deg !=nil {return _deg ;};};case _ga .EndElement :break _fa ;case _ga .CharData :};};return nil ;};

// Validate validates the CT_Override and its children
func (_aa *CT_Override )Validate ()error {return _aa .ValidateWithPath ("C\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065");};

// Validate validates the CT_Default and its children
func (_ag *CT_Default )Validate ()error {return _ag .ValidateWithPath ("\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074");};

// Validate validates the Types and its children
func (_abfb *Types )Validate ()error {return _abfb .ValidateWithPath ("\u0054\u0079\u0070e\u0073")};

// ValidateWithPath validates the Types and its children, prefixing error messages with path
func (_ebgc *Types )ValidateWithPath (path string )error {if _ba :=_ebgc .CT_Types .ValidateWithPath (path );_ba !=nil {return _ba ;};return nil ;};func (_deb *Default )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {_deb .CT_Default =*NewCT_Default ();for _ ,_fbe :=range start .Attr {if _fbe .Name .Local =="\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"{_ef ,_be :=_fbe .Value ,error (nil );if _be !=nil {return _be ;};_deb .ExtensionAttr =_ef ;continue ;};if _fbe .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_daf ,_eaf :=_fbe .Value ,error (nil );if _eaf !=nil {return _eaf ;};_deb .ContentTypeAttr =_daf ;continue ;};};for {_bed ,_eff :=d .Token ();if _eff !=nil {return _ad .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u003a\u0020\u0025\u0073",_eff );};if _fgb ,_beb :=_bed .(_ga .EndElement );_beb &&_fgb .Name ==start .Name {break ;};};return nil ;};func (_eaa *Override )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {return _eaa .CT_Override .MarshalXML (e ,start );};

// ValidateWithPath validates the CT_Types and its children, prefixing error messages with path
func (_dd *CT_Types )ValidateWithPath (path string )error {for _cf ,_bc :=range _dd .Default {if _gec :=_bc .ValidateWithPath (_ad .Sprintf ("\u0025\u0073\u002f\u0044\u0065\u0066\u0061\u0075\u006ct\u005b\u0025\u0064\u005d",path ,_cf ));_gec !=nil {return _gec ;};};for _fgf ,_dga :=range _dd .Override {if _fd :=_dga .ValidateWithPath (_ad .Sprintf ("\u0025s\u002fO\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u005b\u0025\u0064\u005d",path ,_fgf ));_fd !=nil {return _fd ;};};return nil ;};func (_dge *CT_Types )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {e .EncodeToken (start );if _dge .Default !=nil {_gce :=_ga .StartElement {Name :_ga .Name {Local :"\u0044e\u0066\u0061\u0075\u006c\u0074"}};for _ ,_cg :=range _dge .Default {e .EncodeElement (_cg ,_gce );};};if _dge .Override !=nil {_ccd :=_ga .StartElement {Name :_ga .Name {Local :"\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}};for _ ,_ac :=range _dge .Override {e .EncodeElement (_ac ,_ccd );};};e .EncodeToken (_ga .EndElement {Name :start .Name });return nil ;};func (_aae *Types )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {start .Attr =append (start .Attr ,_ga .Attr {Name :_ga .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s"});start .Attr =append (start .Attr ,_ga .Attr {Name :_ga .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u0054\u0079\u0070e\u0073";return _aae .CT_Types .MarshalXML (e ,start );};func (_fg *CT_Default )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {_fg .ExtensionAttr ="\u0078\u006d\u006c";_fg .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";for _ ,_ce :=range start .Attr {if _ce .Name .Local =="\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"{_e ,_d :=_ce .Value ,error (nil );if _d !=nil {return _d ;};_fg .ExtensionAttr =_e ;continue ;};if _ce .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_dg ,_gc :=_ce .Value ,error (nil );if _gc !=nil {return _gc ;};_fg .ContentTypeAttr =_dg ;continue ;};};for {_ee ,_da :=d .Token ();if _da !=nil {return _ad .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u003a\u0020%\u0073",_da );};if _eb ,_ceg :=_ee .(_ga .EndElement );_ceg &&_eb .Name ==start .Name {break ;};};return nil ;};type Override struct{CT_Override };

// ValidateWithPath validates the Default and its children, prefixing error messages with path
func (_cce *Default )ValidateWithPath (path string )error {if _ccc :=_cce .CT_Default .ValidateWithPath (path );_ccc !=nil {return _ccc ;};return nil ;};type CT_Types struct{Default []*Default ;Override []*Override ;};type Default struct{CT_Default };

// Validate validates the CT_Types and its children
func (_abf *CT_Types )Validate ()error {return _abf .ValidateWithPath ("\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073");};

// ValidateWithPath validates the CT_Override and its children, prefixing error messages with path
func (_gba *CT_Override )ValidateWithPath (path string )error {if !ST_ContentTypePatternRe .MatchString (_gba .ContentTypeAttr ){return _ad .Errorf ("\u0025\u0073/\u006d\u002e\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0027\u0025\u0073\u0027\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029",path ,ST_ContentTypePatternRe ,_gba .ContentTypeAttr );};return nil ;};var ST_ContentTypePatternRe =_a .MustCompile (ST_ContentTypePattern );func (_fcf *Types )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {_fcf .CT_Types =*NewCT_Types ();_debc :for {_fcb ,_afc :=d .Token ();if _afc !=nil {return _afc ;};switch _caa :=_fcb .(type ){case _ga .StartElement :switch _caa .Name {case _ga .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u0044e\u0066\u0061\u0075\u006c\u0074"}:_fga :=NewDefault ();if _def :=d .DecodeElement (_fga ,&_caa );_def !=nil {return _def ;};_fcf .Default =append (_fcf .Default ,_fga );case _ga .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}:_gf :=NewOverride ();if _geca :=d .DecodeElement (_gf ,&_caa );_geca !=nil {return _geca ;};_fcf .Override =append (_fcf .Override ,_gf );default:_f .Log ("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006fn \u0054\u0079\u0070e\u0073 \u0025\u0076",_caa .Name );if _agd :=d .Skip ();_agd !=nil {return _agd ;};};case _ga .EndElement :break _debc ;case _ga .CharData :};};return nil ;};type CT_Override struct{ContentTypeAttr string ;PartNameAttr string ;};const ST_ExtensionPattern ="\u0028\u005b\u0021\u0024\u0026\u0027\\\u0028\u005c\u0029\u005c\u002a\\\u002b\u002c\u003a\u003d\u005d\u007c(\u0025\u005b\u0030\u002d\u0039a\u002d\u0066\u0041\u002d\u0046\u005d\u005b\u0030\u002d\u0039\u0061\u002df\u0041\u002d\u0046\u005d\u0029\u007c\u005b\u003a\u0040\u005d\u007c\u005b\u0061\u002d\u007aA\u002d\u005a\u0030\u002d\u0039\u005c\u002d\u005f~\u005d\u0029\u002b";const ST_ContentTypePattern ="\u005e\\\u0070{\u004c\u0061\u0074\u0069\u006e\u007d\u002b\u002f\u002e\u002a\u0024";func NewCT_Types ()*CT_Types {_ff :=&CT_Types {};return _ff };

// ValidateWithPath validates the Override and its children, prefixing error messages with path
func (_caf *Override )ValidateWithPath (path string )error {if _bgf :=_caf .CT_Override .ValidateWithPath (path );_bgf !=nil {return _bgf ;};return nil ;};func (_ed *CT_Override )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {start .Attr =append (start .Attr ,_ga .Attr {Name :_ga .Name {Local :"C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"},Value :_ad .Sprintf ("\u0025\u0076",_ed .ContentTypeAttr )});start .Attr =append (start .Attr ,_ga .Attr {Name :_ga .Name {Local :"\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065"},Value :_ad .Sprintf ("\u0025\u0076",_ed .PartNameAttr )});e .EncodeToken (start );e .EncodeToken (_ga .EndElement {Name :start .Name });return nil ;};func (_c *CT_Default )MarshalXML (e *_ga .Encoder ,start _ga .StartElement )error {start .Attr =append (start .Attr ,_ga .Attr {Name :_ga .Name {Local :"\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"},Value :_ad .Sprintf ("\u0025\u0076",_c .ExtensionAttr )});start .Attr =append (start .Attr ,_ga .Attr {Name :_ga .Name {Local :"C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"},Value :_ad .Sprintf ("\u0025\u0076",_c .ContentTypeAttr )});e .EncodeToken (start );e .EncodeToken (_ga .EndElement {Name :start .Name });return nil ;};func NewDefault ()*Default {_fb :=&Default {};_fb .CT_Default =*NewCT_Default ();return _fb };type CT_Default struct{ExtensionAttr string ;ContentTypeAttr string ;};func NewCT_Override ()*CT_Override {_fgd :=&CT_Override {};_fgd .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";return _fgd ;};func (_ca *Override )UnmarshalXML (d *_ga .Decoder ,start _ga .StartElement )error {_ca .CT_Override =*NewCT_Override ();for _ ,_gad :=range start .Attr {if _gad .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_bga ,_eeg :=_gad .Value ,error (nil );if _eeg !=nil {return _eeg ;};_ca .ContentTypeAttr =_bga ;continue ;};if _gad .Name .Local =="\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065"{_gbg ,_ccca :=_gad .Value ,error (nil );if _ccca !=nil {return _ccca ;};_ca .PartNameAttr =_gbg ;continue ;};};for {_aba ,_dbe :=d .Token ();if _dbe !=nil {return _ad .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u004f\u0076\u0065r\u0072\u0069\u0064\u0065: \u0025\u0073",_dbe );};if _aaab ,_gg :=_aba .(_ga .EndElement );_gg &&_aaab .Name ==start .Name {break ;};};return nil ;};

// Validate validates the Override and its children
func (_fbeb *Override )Validate ()error {return _fbeb .ValidateWithPath ("\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065");};func init (){_f .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073",NewCT_Types );_f .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074",NewCT_Default );_f .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","C\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065",NewCT_Override );_f .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0054\u0079\u0070e\u0073",NewTypes );_f .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0044e\u0066\u0061\u0075\u006c\u0074",NewDefault );_f .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065",NewOverride );};