package slinguist

// CODE GENERATED AUTOMATICALLY WITH github.com/src-d/simple-linguist/cli/slinguist-generate
// THIS FILE SHOULD NOT BE EDITED BY HAND
// Extracted from github/linguist commit: dae33dc2b20cddc85d1300435c3be7118a7115a9

type Type int

const (
	TypeUnknown Type = iota
	TypeData
	TypeProgramming
	TypeMarkup
	TypeProse
)

func GetLanguageType(language string) (langType Type) {
	langType, _ = languagesType[language]
	return langType
}

var languagesType = map[string]Type{
	"1C Enterprise":    TypeProgramming,
	"ABAP":             TypeProgramming,
	"ABNF":             TypeData,
	"AGS Script":       TypeProgramming,
	"AMPL":             TypeProgramming,
	"ANTLR":            TypeProgramming,
	"API Blueprint":    TypeMarkup,
	"APL":              TypeProgramming,
	"ASN.1":            TypeData,
	"ASP":              TypeProgramming,
	"ATS":              TypeProgramming,
	"ActionScript":     TypeProgramming,
	"Ada":              TypeProgramming,
	"Agda":             TypeProgramming,
	"Alloy":            TypeProgramming,
	"Alpine Abuild":    TypeProgramming,
	"Ant Build System": TypeData,
	"ApacheConf":       TypeMarkup,
	"Apex":             TypeProgramming,
	"Apollo Guidance Computer": TypeProgramming,
	"AppleScript":              TypeProgramming,
	"Arc":                      TypeProgramming,
	"Arduino":                  TypeProgramming,
	"AsciiDoc":                 TypeProse,
	"AspectJ":                  TypeProgramming,
	"Assembly":                 TypeProgramming,
	"Augeas":                   TypeProgramming,
	"AutoHotkey":               TypeProgramming,
	"AutoIt":                   TypeProgramming,
	"Awk":                      TypeProgramming,
	"Batchfile":                TypeProgramming,
	"Befunge":                  TypeProgramming,
	"Bison":                    TypeProgramming,
	"BitBake":                  TypeProgramming,
	"Blade":                    TypeMarkup,
	"BlitzBasic":               TypeProgramming,
	"BlitzMax":                 TypeProgramming,
	"Bluespec":                 TypeProgramming,
	"Boo":                      TypeProgramming,
	"Brainfuck":                TypeProgramming,
	"Brightscript":             TypeProgramming,
	"Bro":                      TypeProgramming,
	"C":                        TypeProgramming,
	"C#":                       TypeProgramming,
	"C++":                      TypeProgramming,
	"C-ObjDump":                TypeData,
	"C2hs Haskell":             TypeProgramming,
	"CLIPS":                    TypeProgramming,
	"CMake":                    TypeProgramming,
	"COBOL":                    TypeProgramming,
	"COLLADA":                  TypeData,
	"CSON":                     TypeData,
	"CSS":                      TypeMarkup,
	"CSV":                      TypeData,
	"Cap'n Proto":              TypeProgramming,
	"CartoCSS":                 TypeProgramming,
	"Ceylon":                   TypeProgramming,
	"Chapel":                   TypeProgramming,
	"Charity":                  TypeProgramming,
	"ChucK":                    TypeProgramming,
	"Cirru":                    TypeProgramming,
	"Clarion":                  TypeProgramming,
	"Clean":                    TypeProgramming,
	"Click":                    TypeProgramming,
	"Clojure":                  TypeProgramming,
	"CoffeeScript":             TypeProgramming,
	"ColdFusion":               TypeProgramming,
	"ColdFusion CFC":           TypeProgramming,
	"Common Lisp":              TypeProgramming,
	"Component Pascal":         TypeProgramming,
	"Cool":                     TypeProgramming,
	"Coq":                      TypeProgramming,
	"Cpp-ObjDump":              TypeData,
	"Creole":                   TypeProse,
	"Crystal":                  TypeProgramming,
	"Csound":                   TypeProgramming,
	"Csound Document":          TypeProgramming,
	"Csound Score":             TypeProgramming,
	"Cuda":                     TypeProgramming,
	"Cycript":                  TypeProgramming,
	"Cython":                   TypeProgramming,
	"D":                        TypeProgramming,
	"D-ObjDump":                TypeData,
	"DIGITAL Command Language": TypeProgramming,
	"DM":             TypeProgramming,
	"DNS Zone":       TypeData,
	"DTrace":         TypeProgramming,
	"Darcs Patch":    TypeData,
	"Dart":           TypeProgramming,
	"Diff":           TypeData,
	"Dockerfile":     TypeData,
	"Dogescript":     TypeProgramming,
	"Dylan":          TypeProgramming,
	"E":              TypeProgramming,
	"EBNF":           TypeData,
	"ECL":            TypeProgramming,
	"ECLiPSe":        TypeProgramming,
	"EJS":            TypeMarkup,
	"EQ":             TypeProgramming,
	"Eagle":          TypeMarkup,
	"Ecere Projects": TypeData,
	"Eiffel":         TypeProgramming,
	"Elixir":         TypeProgramming,
	"Elm":            TypeProgramming,
	"Emacs Lisp":     TypeProgramming,
	"EmberScript":    TypeProgramming,
	"Erlang":         TypeProgramming,
	"F#":             TypeProgramming,
	"FLUX":           TypeProgramming,
	"Factor":         TypeProgramming,
	"Fancy":          TypeProgramming,
	"Fantom":         TypeProgramming,
	"Filebench WML":  TypeProgramming,
	"Filterscript":   TypeProgramming,
	"Formatted":      TypeData,
	"Forth":          TypeProgramming,
	"Fortran":        TypeProgramming,
	"FreeMarker":     TypeProgramming,
	"Frege":          TypeProgramming,
	"G-code":         TypeData,
	"GAMS":           TypeProgramming,
	"GAP":            TypeProgramming,
	"GCC Machine Description": TypeProgramming,
	"GDB":      TypeProgramming,
	"GDScript": TypeProgramming,
	"GLSL":     TypeProgramming,
	"GN":       TypeData,
	"Game Maker Language":     TypeProgramming,
	"Genie":                   TypeProgramming,
	"Genshi":                  TypeProgramming,
	"Gentoo Ebuild":           TypeProgramming,
	"Gentoo Eclass":           TypeProgramming,
	"Gettext Catalog":         TypeProse,
	"Gherkin":                 TypeProgramming,
	"Glyph":                   TypeProgramming,
	"Gnuplot":                 TypeProgramming,
	"Go":                      TypeProgramming,
	"Golo":                    TypeProgramming,
	"Gosu":                    TypeProgramming,
	"Grace":                   TypeProgramming,
	"Gradle":                  TypeData,
	"Grammatical Framework":   TypeProgramming,
	"Graph Modeling Language": TypeData,
	"GraphQL":                 TypeData,
	"Graphviz (DOT)":          TypeData,
	"Groovy":                  TypeProgramming,
	"Groovy Server Pages":     TypeProgramming,
	"HCL":                      TypeProgramming,
	"HLSL":                     TypeProgramming,
	"HTML":                     TypeMarkup,
	"HTML+Django":              TypeMarkup,
	"HTML+ECR":                 TypeMarkup,
	"HTML+EEX":                 TypeMarkup,
	"HTML+ERB":                 TypeMarkup,
	"HTML+PHP":                 TypeMarkup,
	"HTTP":                     TypeData,
	"Hack":                     TypeProgramming,
	"Haml":                     TypeMarkup,
	"Handlebars":               TypeMarkup,
	"Harbour":                  TypeProgramming,
	"Haskell":                  TypeProgramming,
	"Haxe":                     TypeProgramming,
	"Hy":                       TypeProgramming,
	"HyPhy":                    TypeProgramming,
	"IDL":                      TypeProgramming,
	"IGOR Pro":                 TypeProgramming,
	"INI":                      TypeData,
	"IRC log":                  TypeData,
	"Idris":                    TypeProgramming,
	"Inform 7":                 TypeProgramming,
	"Inno Setup":               TypeProgramming,
	"Io":                       TypeProgramming,
	"Ioke":                     TypeProgramming,
	"Isabelle":                 TypeProgramming,
	"Isabelle ROOT":            TypeProgramming,
	"J":                        TypeProgramming,
	"JFlex":                    TypeProgramming,
	"JSON":                     TypeData,
	"JSON5":                    TypeData,
	"JSONLD":                   TypeData,
	"JSONiq":                   TypeProgramming,
	"JSX":                      TypeProgramming,
	"Jasmin":                   TypeProgramming,
	"Java":                     TypeProgramming,
	"Java Server Pages":        TypeProgramming,
	"JavaScript":               TypeProgramming,
	"Jison":                    TypeProgramming,
	"Jison Lex":                TypeProgramming,
	"Julia":                    TypeProgramming,
	"Jupyter Notebook":         TypeMarkup,
	"KRL":                      TypeProgramming,
	"KiCad":                    TypeProgramming,
	"Kit":                      TypeMarkup,
	"Kotlin":                   TypeProgramming,
	"LFE":                      TypeProgramming,
	"LLVM":                     TypeProgramming,
	"LOLCODE":                  TypeProgramming,
	"LSL":                      TypeProgramming,
	"LabVIEW":                  TypeProgramming,
	"Lasso":                    TypeProgramming,
	"Latte":                    TypeMarkup,
	"Lean":                     TypeProgramming,
	"Less":                     TypeMarkup,
	"Lex":                      TypeProgramming,
	"LilyPond":                 TypeProgramming,
	"Limbo":                    TypeProgramming,
	"Linker Script":            TypeData,
	"Linux Kernel Module":      TypeData,
	"Liquid":                   TypeMarkup,
	"Literate Agda":            TypeProgramming,
	"Literate CoffeeScript":    TypeProgramming,
	"Literate Haskell":         TypeProgramming,
	"LiveScript":               TypeProgramming,
	"Logos":                    TypeProgramming,
	"Logtalk":                  TypeProgramming,
	"LookML":                   TypeProgramming,
	"LoomScript":               TypeProgramming,
	"Lua":                      TypeProgramming,
	"M":                        TypeProgramming,
	"M4":                       TypeProgramming,
	"M4Sugar":                  TypeProgramming,
	"MAXScript":                TypeProgramming,
	"MQL4":                     TypeProgramming,
	"MQL5":                     TypeProgramming,
	"MTML":                     TypeMarkup,
	"MUF":                      TypeProgramming,
	"Makefile":                 TypeProgramming,
	"Mako":                     TypeProgramming,
	"Markdown":                 TypeProse,
	"Marko":                    TypeMarkup,
	"Mask":                     TypeMarkup,
	"Mathematica":              TypeProgramming,
	"Matlab":                   TypeProgramming,
	"Maven POM":                TypeData,
	"Max":                      TypeProgramming,
	"MediaWiki":                TypeProse,
	"Mercury":                  TypeProgramming,
	"Meson":                    TypeProgramming,
	"Metal":                    TypeProgramming,
	"MiniD":                    TypeProgramming,
	"Mirah":                    TypeProgramming,
	"Modelica":                 TypeProgramming,
	"Modula-2":                 TypeProgramming,
	"Module Management System": TypeProgramming,
	"Monkey":                   TypeProgramming,
	"Moocode":                  TypeProgramming,
	"MoonScript":               TypeProgramming,
	"Myghty":                   TypeProgramming,
	"NCL":                      TypeProgramming,
	"NL":                       TypeData,
	"NSIS":                     TypeProgramming,
	"Nemerle":                  TypeProgramming,
	"NetLinx":                  TypeProgramming,
	"NetLinx+ERB":              TypeProgramming,
	"NetLogo":                  TypeProgramming,
	"NewLisp":                  TypeProgramming,
	"Nginx":                    TypeMarkup,
	"Nim":                      TypeProgramming,
	"Ninja":                    TypeData,
	"Nit":                      TypeProgramming,
	"Nix":                      TypeProgramming,
	"Nu":                       TypeProgramming,
	"NumPy":                    TypeProgramming,
	"OCaml":                    TypeProgramming,
	"ObjDump":                  TypeData,
	"Objective-C":              TypeProgramming,
	"Objective-C++":            TypeProgramming,
	"Objective-J":              TypeProgramming,
	"Omgrofl":                  TypeProgramming,
	"Opa":                      TypeProgramming,
	"Opal":                     TypeProgramming,
	"OpenCL":                   TypeProgramming,
	"OpenEdge ABL":             TypeProgramming,
	"OpenRC runscript":         TypeProgramming,
	"OpenSCAD":                 TypeProgramming,
	"OpenType Feature File":    TypeData,
	"Org":                            TypeProse,
	"Ox":                             TypeProgramming,
	"Oxygene":                        TypeProgramming,
	"Oz":                             TypeProgramming,
	"P4":                             TypeProgramming,
	"PAWN":                           TypeProgramming,
	"PHP":                            TypeProgramming,
	"PLSQL":                          TypeProgramming,
	"PLpgSQL":                        TypeProgramming,
	"POV-Ray SDL":                    TypeProgramming,
	"Pan":                            TypeProgramming,
	"Papyrus":                        TypeProgramming,
	"Parrot":                         TypeProgramming,
	"Parrot Assembly":                TypeProgramming,
	"Parrot Internal Representation": TypeProgramming,
	"Pascal":                       TypeProgramming,
	"Perl":                         TypeProgramming,
	"Perl6":                        TypeProgramming,
	"Pic":                          TypeMarkup,
	"Pickle":                       TypeData,
	"PicoLisp":                     TypeProgramming,
	"PigLatin":                     TypeProgramming,
	"Pike":                         TypeProgramming,
	"Pod":                          TypeProse,
	"PogoScript":                   TypeProgramming,
	"Pony":                         TypeProgramming,
	"PostScript":                   TypeMarkup,
	"PowerBuilder":                 TypeProgramming,
	"PowerShell":                   TypeProgramming,
	"Processing":                   TypeProgramming,
	"Prolog":                       TypeProgramming,
	"Propeller Spin":               TypeProgramming,
	"Protocol Buffer":              TypeMarkup,
	"Public Key":                   TypeData,
	"Pug":                          TypeMarkup,
	"Puppet":                       TypeProgramming,
	"Pure Data":                    TypeProgramming,
	"PureBasic":                    TypeProgramming,
	"PureScript":                   TypeProgramming,
	"Python":                       TypeProgramming,
	"Python console":               TypeProgramming,
	"Python traceback":             TypeData,
	"QML":                          TypeProgramming,
	"QMake":                        TypeProgramming,
	"R":                            TypeProgramming,
	"RAML":                         TypeMarkup,
	"RDoc":                         TypeProse,
	"REALbasic":                    TypeProgramming,
	"REXX":                         TypeProgramming,
	"RHTML":                        TypeMarkup,
	"RMarkdown":                    TypeProse,
	"RPM Spec":                     TypeData,
	"RUNOFF":                       TypeMarkup,
	"Racket":                       TypeProgramming,
	"Ragel":                        TypeProgramming,
	"Rascal":                       TypeProgramming,
	"Raw token data":               TypeData,
	"Reason":                       TypeProgramming,
	"Rebol":                        TypeProgramming,
	"Red":                          TypeProgramming,
	"Redcode":                      TypeProgramming,
	"Regular Expression":           TypeData,
	"Ren'Py":                       TypeProgramming,
	"RenderScript":                 TypeProgramming,
	"RobotFramework":               TypeProgramming,
	"Roff":                         TypeMarkup,
	"Rouge":                        TypeProgramming,
	"Ruby":                         TypeProgramming,
	"Rust":                         TypeProgramming,
	"SAS":                          TypeProgramming,
	"SCSS":                         TypeMarkup,
	"SMT":                          TypeProgramming,
	"SPARQL":                       TypeData,
	"SQF":                          TypeProgramming,
	"SQL":                          TypeData,
	"SQLPL":                        TypeProgramming,
	"SRecode Template":             TypeMarkup,
	"STON":                         TypeData,
	"SVG":                          TypeData,
	"Sage":                         TypeProgramming,
	"SaltStack":                    TypeProgramming,
	"Sass":                         TypeMarkup,
	"Scala":                        TypeProgramming,
	"Scaml":                        TypeMarkup,
	"Scheme":                       TypeProgramming,
	"Scilab":                       TypeProgramming,
	"Self":                         TypeProgramming,
	"Shell":                        TypeProgramming,
	"ShellSession":                 TypeProgramming,
	"Shen":                         TypeProgramming,
	"Slash":                        TypeProgramming,
	"Slim":                         TypeMarkup,
	"Smali":                        TypeProgramming,
	"Smalltalk":                    TypeProgramming,
	"Smarty":                       TypeProgramming,
	"SourcePawn":                   TypeProgramming,
	"Spline Font Database":         TypeData,
	"Squirrel":                     TypeProgramming,
	"Stan":                         TypeProgramming,
	"Standard ML":                  TypeProgramming,
	"Stata":                        TypeProgramming,
	"Stylus":                       TypeMarkup,
	"SubRip Text":                  TypeData,
	"Sublime Text Config":          TypeData,
	"SuperCollider":                TypeProgramming,
	"Swift":                        TypeProgramming,
	"SystemVerilog":                TypeProgramming,
	"TI Program":                   TypeProgramming,
	"TLA":                          TypeProgramming,
	"TOML":                         TypeData,
	"TXL":                          TypeProgramming,
	"Tcl":                          TypeProgramming,
	"Tcsh":                         TypeProgramming,
	"TeX":                          TypeMarkup,
	"Tea":                          TypeMarkup,
	"Terra":                        TypeProgramming,
	"Text":                         TypeProse,
	"Textile":                      TypeProse,
	"Thrift":                       TypeProgramming,
	"Turing":                       TypeProgramming,
	"Turtle":                       TypeData,
	"Twig":                         TypeMarkup,
	"TypeScript":                   TypeProgramming,
	"Unified Parallel C":           TypeProgramming,
	"Unity3D Asset":                TypeData,
	"Unix Assembly":                TypeProgramming,
	"Uno":                          TypeProgramming,
	"UnrealScript":                 TypeProgramming,
	"UrWeb":                        TypeProgramming,
	"VCL":                          TypeProgramming,
	"VHDL":                         TypeProgramming,
	"Vala":                         TypeProgramming,
	"Verilog":                      TypeProgramming,
	"Vim script":                   TypeProgramming,
	"Visual Basic":                 TypeProgramming,
	"Volt":                         TypeProgramming,
	"Vue":                          TypeMarkup,
	"Wavefront Material":           TypeData,
	"Wavefront Object":             TypeData,
	"Web Ontology Language":        TypeMarkup,
	"WebIDL":                       TypeProgramming,
	"World of Warcraft Addon Data": TypeData,
	"X10":              TypeProgramming,
	"XC":               TypeProgramming,
	"XCompose":         TypeData,
	"XML":              TypeData,
	"XPages":           TypeProgramming,
	"XProc":            TypeProgramming,
	"XQuery":           TypeProgramming,
	"XS":               TypeProgramming,
	"XSLT":             TypeProgramming,
	"Xojo":             TypeProgramming,
	"Xtend":            TypeProgramming,
	"YAML":             TypeData,
	"YANG":             TypeData,
	"Yacc":             TypeProgramming,
	"Zephir":           TypeProgramming,
	"Zimpl":            TypeProgramming,
	"desktop":          TypeData,
	"eC":               TypeProgramming,
	"edn":              TypeData,
	"fish":             TypeProgramming,
	"mupad":            TypeProgramming,
	"nesC":             TypeProgramming,
	"ooc":              TypeProgramming,
	"reStructuredText": TypeProse,
	"wisp":             TypeProgramming,
	"xBase":            TypeProgramming,
}