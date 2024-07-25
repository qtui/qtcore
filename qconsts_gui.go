package qtcore

type QAccessible__Event = int

const QAccessible__SoundPlayed QAccessible__Event = 1
const QAccessible__Alert QAccessible__Event = 2
const QAccessible__ForegroundChanged QAccessible__Event = 3
const QAccessible__MenuStart QAccessible__Event = 4
const QAccessible__MenuEnd QAccessible__Event = 5
const QAccessible__PopupMenuStart QAccessible__Event = 6
const QAccessible__PopupMenuEnd QAccessible__Event = 7
const QAccessible__ContextHelpStart QAccessible__Event = 12
const QAccessible__ContextHelpEnd QAccessible__Event = 13
const QAccessible__DragDropStart QAccessible__Event = 14
const QAccessible__DragDropEnd QAccessible__Event = 15
const QAccessible__DialogStart QAccessible__Event = 16
const QAccessible__DialogEnd QAccessible__Event = 17
const QAccessible__ScrollingStart QAccessible__Event = 18
const QAccessible__ScrollingEnd QAccessible__Event = 19
const QAccessible__MenuCommand QAccessible__Event = 24
const QAccessible__ActionChanged QAccessible__Event = 257
const QAccessible__ActiveDescendantChanged QAccessible__Event = 258
const QAccessible__AttributeChanged QAccessible__Event = 259
const QAccessible__DocumentContentChanged QAccessible__Event = 260
const QAccessible__DocumentLoadComplete QAccessible__Event = 261
const QAccessible__DocumentLoadStopped QAccessible__Event = 262
const QAccessible__DocumentReload QAccessible__Event = 263
const QAccessible__HyperlinkEndIndexChanged QAccessible__Event = 264
const QAccessible__HyperlinkNumberOfAnchorsChanged QAccessible__Event = 265
const QAccessible__HyperlinkSelectedLinkChanged QAccessible__Event = 266
const QAccessible__HypertextLinkActivated QAccessible__Event = 267
const QAccessible__HypertextLinkSelected QAccessible__Event = 268
const QAccessible__HyperlinkStartIndexChanged QAccessible__Event = 269
const QAccessible__HypertextChanged QAccessible__Event = 270
const QAccessible__HypertextNLinksChanged QAccessible__Event = 271
const QAccessible__ObjectAttributeChanged QAccessible__Event = 272
const QAccessible__PageChanged QAccessible__Event = 273
const QAccessible__SectionChanged QAccessible__Event = 274
const QAccessible__TableCaptionChanged QAccessible__Event = 275
const QAccessible__TableColumnDescriptionChanged QAccessible__Event = 276
const QAccessible__TableColumnHeaderChanged QAccessible__Event = 277
const QAccessible__TableModelChanged QAccessible__Event = 278
const QAccessible__TableRowDescriptionChanged QAccessible__Event = 279
const QAccessible__TableRowHeaderChanged QAccessible__Event = 280
const QAccessible__TableSummaryChanged QAccessible__Event = 281
const QAccessible__TextAttributeChanged QAccessible__Event = 282
const QAccessible__TextCaretMoved QAccessible__Event = 283
const QAccessible__TextColumnChanged QAccessible__Event = 285
const QAccessible__TextInserted QAccessible__Event = 286
const QAccessible__TextRemoved QAccessible__Event = 287
const QAccessible__TextUpdated QAccessible__Event = 288
const QAccessible__TextSelectionChanged QAccessible__Event = 289
const QAccessible__VisibleDataChanged QAccessible__Event = 290
const QAccessible__ObjectCreated QAccessible__Event = 32768
const QAccessible__ObjectDestroyed QAccessible__Event = 32769
const QAccessible__ObjectShow QAccessible__Event = 32770
const QAccessible__ObjectHide QAccessible__Event = 32771
const QAccessible__ObjectReorder QAccessible__Event = 32772
const QAccessible__Focus QAccessible__Event = 32773
const QAccessible__Selection QAccessible__Event = 32774
const QAccessible__SelectionAdd QAccessible__Event = 32775
const QAccessible__SelectionRemove QAccessible__Event = 32776
const QAccessible__SelectionWithin QAccessible__Event = 32777
const QAccessible__StateChanged QAccessible__Event = 32778
const QAccessible__LocationChanged QAccessible__Event = 32779
const QAccessible__NameChanged QAccessible__Event = 32780
const QAccessible__DescriptionChanged QAccessible__Event = 32781
const QAccessible__ValueChanged QAccessible__Event = 32782
const QAccessible__ParentChanged QAccessible__Event = 32783
const QAccessible__HelpChanged QAccessible__Event = 32928
const QAccessible__DefaultActionChanged QAccessible__Event = 32944
const QAccessible__AcceleratorChanged QAccessible__Event = 32960
const QAccessible__InvalidEvent QAccessible__Event = 32961

type QAccessible__Role = int

const QAccessible__NoRole QAccessible__Role = 0
const QAccessible__TitleBar QAccessible__Role = 1
const QAccessible__MenuBar QAccessible__Role = 2
const QAccessible__ScrollBar QAccessible__Role = 3
const QAccessible__Grip QAccessible__Role = 4
const QAccessible__Sound QAccessible__Role = 5
const QAccessible__Cursor QAccessible__Role = 6
const QAccessible__Caret QAccessible__Role = 7
const QAccessible__AlertMessage QAccessible__Role = 8
const QAccessible__Window QAccessible__Role = 9
const QAccessible__Client QAccessible__Role = 10
const QAccessible__PopupMenu QAccessible__Role = 11
const QAccessible__MenuItem QAccessible__Role = 12
const QAccessible__ToolTip QAccessible__Role = 13
const QAccessible__Application QAccessible__Role = 14
const QAccessible__Document QAccessible__Role = 15
const QAccessible__Pane QAccessible__Role = 16
const QAccessible__Chart QAccessible__Role = 17
const QAccessible__Dialog QAccessible__Role = 18
const QAccessible__Border QAccessible__Role = 19
const QAccessible__Grouping QAccessible__Role = 20
const QAccessible__Separator QAccessible__Role = 21
const QAccessible__ToolBar QAccessible__Role = 22
const QAccessible__StatusBar QAccessible__Role = 23
const QAccessible__Table QAccessible__Role = 24
const QAccessible__ColumnHeader QAccessible__Role = 25
const QAccessible__RowHeader QAccessible__Role = 26
const QAccessible__Column QAccessible__Role = 27
const QAccessible__Row QAccessible__Role = 28
const QAccessible__Cell QAccessible__Role = 29
const QAccessible__Link QAccessible__Role = 30
const QAccessible__HelpBalloon QAccessible__Role = 31
const QAccessible__Assistant QAccessible__Role = 32
const QAccessible__List QAccessible__Role = 33
const QAccessible__ListItem QAccessible__Role = 34
const QAccessible__Tree QAccessible__Role = 35
const QAccessible__TreeItem QAccessible__Role = 36
const QAccessible__PageTab QAccessible__Role = 37
const QAccessible__PropertyPage QAccessible__Role = 38
const QAccessible__Indicator QAccessible__Role = 39
const QAccessible__Graphic QAccessible__Role = 40
const QAccessible__StaticText QAccessible__Role = 41
const QAccessible__EditableText QAccessible__Role = 42
const QAccessible__Button QAccessible__Role = 43
const QAccessible__PushButton QAccessible__Role = 43
const QAccessible__CheckBox QAccessible__Role = 44
const QAccessible__RadioButton QAccessible__Role = 45
const QAccessible__ComboBox QAccessible__Role = 46
const QAccessible__ProgressBar QAccessible__Role = 48
const QAccessible__Dial QAccessible__Role = 49
const QAccessible__HotkeyField QAccessible__Role = 50
const QAccessible__Slider QAccessible__Role = 51
const QAccessible__SpinBox QAccessible__Role = 52
const QAccessible__Canvas QAccessible__Role = 53
const QAccessible__Animation QAccessible__Role = 54
const QAccessible__Equation QAccessible__Role = 55
const QAccessible__ButtonDropDown QAccessible__Role = 56
const QAccessible__ButtonMenu QAccessible__Role = 57
const QAccessible__ButtonDropGrid QAccessible__Role = 58
const QAccessible__Whitespace QAccessible__Role = 59
const QAccessible__PageTabList QAccessible__Role = 60
const QAccessible__Clock QAccessible__Role = 61
const QAccessible__Splitter QAccessible__Role = 62
const QAccessible__LayeredPane QAccessible__Role = 128
const QAccessible__Terminal QAccessible__Role = 129
const QAccessible__Desktop QAccessible__Role = 130
const QAccessible__Paragraph QAccessible__Role = 131
const QAccessible__WebDocument QAccessible__Role = 132
const QAccessible__Section QAccessible__Role = 133
const QAccessible__ColorChooser QAccessible__Role = 1028
const QAccessible__Footer QAccessible__Role = 1038
const QAccessible__Form QAccessible__Role = 1040
const QAccessible__Heading QAccessible__Role = 1044
const QAccessible__Note QAccessible__Role = 1051
const QAccessible__ComplementaryContent QAccessible__Role = 1068
const QAccessible__UserRole QAccessible__Role = 65535

type QAccessible__Text = int

const QAccessible__Name QAccessible__Text = 0
const QAccessible__Description QAccessible__Text = 1
const QAccessible__Value QAccessible__Text = 2
const QAccessible__Help QAccessible__Text = 3
const QAccessible__Accelerator QAccessible__Text = 4
const QAccessible__DebugDescription QAccessible__Text = 5
const QAccessible__UserText QAccessible__Text = 65535

type QAccessible__RelationFlag = int

const QAccessible__Label QAccessible__RelationFlag = 1
const QAccessible__Labelled QAccessible__RelationFlag = 2
const QAccessible__Controller QAccessible__RelationFlag = 4
const QAccessible__Controlled QAccessible__RelationFlag = 8
const QAccessible__AllRelations QAccessible__RelationFlag = -1

type QAccessible__InterfaceType = int

const QAccessible__TextInterface QAccessible__InterfaceType = 0
const QAccessible__EditableTextInterface QAccessible__InterfaceType = 1
const QAccessible__ValueInterface QAccessible__InterfaceType = 2
const QAccessible__ActionInterface QAccessible__InterfaceType = 3
const QAccessible__ImageInterface QAccessible__InterfaceType = 4
const QAccessible__TableInterface QAccessible__InterfaceType = 5
const QAccessible__TableCellInterface QAccessible__InterfaceType = 6

type QAccessible__TextBoundaryType = int

const QAccessible__CharBoundary QAccessible__TextBoundaryType = 0
const QAccessible__WordBoundary QAccessible__TextBoundaryType = 1
const QAccessible__SentenceBoundary QAccessible__TextBoundaryType = 2
const QAccessible__ParagraphBoundary QAccessible__TextBoundaryType = 3
const QAccessible__LineBoundary QAccessible__TextBoundaryType = 4
const QAccessible__NoBoundary QAccessible__TextBoundaryType = 5

type QAccessibleTableModelChangeEvent__ModelChangeType = int

const QAccessibleTableModelChangeEvent__ModelReset QAccessibleTableModelChangeEvent__ModelChangeType = 0
const QAccessibleTableModelChangeEvent__DataChanged QAccessibleTableModelChangeEvent__ModelChangeType = 1
const QAccessibleTableModelChangeEvent__RowsInserted QAccessibleTableModelChangeEvent__ModelChangeType = 2
const QAccessibleTableModelChangeEvent__ColumnsInserted QAccessibleTableModelChangeEvent__ModelChangeType = 3
const QAccessibleTableModelChangeEvent__RowsRemoved QAccessibleTableModelChangeEvent__ModelChangeType = 4
const QAccessibleTableModelChangeEvent__ColumnsRemoved QAccessibleTableModelChangeEvent__ModelChangeType = 5

type QClipboard__Mode = int

const QClipboard__Clipboard QClipboard__Mode = 0
const QClipboard__Selection QClipboard__Mode = 1
const QClipboard__FindBuffer QClipboard__Mode = 2
const QClipboard__LastMode QClipboard__Mode = 2

type QColor__Spec = int

const QColor__Invalid QColor__Spec = 0
const QColor__Rgb QColor__Spec = 1
const QColor__Hsv QColor__Spec = 2
const QColor__Cmyk QColor__Spec = 3
const QColor__Hsl QColor__Spec = 4

type QColor__NameFormat = int

const QColor__HexRgb QColor__NameFormat = 0
const QColor__HexArgb QColor__NameFormat = 1

type QContextMenuEvent__Reason = int

const QContextMenuEvent__Mouse QContextMenuEvent__Reason = 0
const QContextMenuEvent__Keyboard QContextMenuEvent__Reason = 1
const QContextMenuEvent__Other QContextMenuEvent__Reason = 2

type QDoubleValidator__Notation = int

const QDoubleValidator__StandardNotation QDoubleValidator__Notation = 0
const QDoubleValidator__ScientificNotation QDoubleValidator__Notation = 1

type QFont__StyleHint = int

const QFont__Helvetica QFont__StyleHint = 0
const QFont__SansSerif QFont__StyleHint = 0
const QFont__Times QFont__StyleHint = 1
const QFont__Serif QFont__StyleHint = 1
const QFont__Courier QFont__StyleHint = 2
const QFont__TypeWriter QFont__StyleHint = 2
const QFont__OldEnglish QFont__StyleHint = 3
const QFont__Decorative QFont__StyleHint = 3
const QFont__System QFont__StyleHint = 4
const QFont__AnyStyle QFont__StyleHint = 5
const QFont__Cursive QFont__StyleHint = 6
const QFont__Monospace QFont__StyleHint = 7
const QFont__Fantasy QFont__StyleHint = 8

type QFont__StyleStrategy = int

const QFont__PreferDefault QFont__StyleStrategy = 1
const QFont__PreferBitmap QFont__StyleStrategy = 2
const QFont__PreferDevice QFont__StyleStrategy = 4
const QFont__PreferOutline QFont__StyleStrategy = 8
const QFont__ForceOutline QFont__StyleStrategy = 16
const QFont__PreferMatch QFont__StyleStrategy = 32
const QFont__PreferQuality QFont__StyleStrategy = 64
const QFont__PreferAntialias QFont__StyleStrategy = 128
const QFont__NoAntialias QFont__StyleStrategy = 256
const QFont__OpenGLCompatible QFont__StyleStrategy = 512
const QFont__ForceIntegerMetrics QFont__StyleStrategy = 1024
const QFont__NoSubpixelAntialias QFont__StyleStrategy = 2048
const QFont__PreferNoShaping QFont__StyleStrategy = 4096
const QFont__NoFontMerging QFont__StyleStrategy = 32768

type QFont__HintingPreference = int

const QFont__PreferDefaultHinting QFont__HintingPreference = 0
const QFont__PreferNoHinting QFont__HintingPreference = 1
const QFont__PreferVerticalHinting QFont__HintingPreference = 2
const QFont__PreferFullHinting QFont__HintingPreference = 3

type QFont__Weight = int

const QFont__Thin QFont__Weight = 0
const QFont__ExtraLight QFont__Weight = 12
const QFont__Light QFont__Weight = 25
const QFont__Normal QFont__Weight = 50
const QFont__Medium QFont__Weight = 57
const QFont__DemiBold QFont__Weight = 63
const QFont__Bold QFont__Weight = 75
const QFont__ExtraBold QFont__Weight = 81
const QFont__Black QFont__Weight = 87

type QFont__Style = int

const QFont__StyleNormal QFont__Style = 0
const QFont__StyleItalic QFont__Style = 1
const QFont__StyleOblique QFont__Style = 2

type QFont__Stretch = int

const QFont__AnyStretch QFont__Stretch = 0
const QFont__UltraCondensed QFont__Stretch = 50
const QFont__ExtraCondensed QFont__Stretch = 62
const QFont__Condensed QFont__Stretch = 75
const QFont__SemiCondensed QFont__Stretch = 87
const QFont__Unstretched QFont__Stretch = 100
const QFont__SemiExpanded QFont__Stretch = 112
const QFont__Expanded QFont__Stretch = 125
const QFont__ExtraExpanded QFont__Stretch = 150
const QFont__UltraExpanded QFont__Stretch = 200

type QFont__Capitalization = int

const QFont__MixedCase QFont__Capitalization = 0
const QFont__AllUppercase QFont__Capitalization = 1
const QFont__AllLowercase QFont__Capitalization = 2
const QFont__SmallCaps QFont__Capitalization = 3
const QFont__Capitalize QFont__Capitalization = 4

type QFont__SpacingType = int

const QFont__PercentageSpacing QFont__SpacingType = 0
const QFont__AbsoluteSpacing QFont__SpacingType = 1

type QFont__ResolveProperties = int

const QFont__FamilyResolved QFont__ResolveProperties = 1
const QFont__SizeResolved QFont__ResolveProperties = 2
const QFont__StyleHintResolved QFont__ResolveProperties = 4
const QFont__StyleStrategyResolved QFont__ResolveProperties = 8
const QFont__WeightResolved QFont__ResolveProperties = 16
const QFont__StyleResolved QFont__ResolveProperties = 32
const QFont__UnderlineResolved QFont__ResolveProperties = 64
const QFont__OverlineResolved QFont__ResolveProperties = 128
const QFont__StrikeOutResolved QFont__ResolveProperties = 256
const QFont__FixedPitchResolved QFont__ResolveProperties = 512
const QFont__StretchResolved QFont__ResolveProperties = 1024
const QFont__KerningResolved QFont__ResolveProperties = 2048
const QFont__CapitalizationResolved QFont__ResolveProperties = 4096
const QFont__LetterSpacingResolved QFont__ResolveProperties = 8192
const QFont__WordSpacingResolved QFont__ResolveProperties = 16384
const QFont__HintingPreferenceResolved QFont__ResolveProperties = 32768
const QFont__StyleNameResolved QFont__ResolveProperties = 65536
const QFont__AllPropertiesResolved QFont__ResolveProperties = 131071

type QFontDatabase__WritingSystem = int

const QFontDatabase__Any QFontDatabase__WritingSystem = 0
const QFontDatabase__Latin QFontDatabase__WritingSystem = 1
const QFontDatabase__Greek QFontDatabase__WritingSystem = 2
const QFontDatabase__Cyrillic QFontDatabase__WritingSystem = 3
const QFontDatabase__Armenian QFontDatabase__WritingSystem = 4
const QFontDatabase__Hebrew QFontDatabase__WritingSystem = 5
const QFontDatabase__Arabic QFontDatabase__WritingSystem = 6
const QFontDatabase__Syriac QFontDatabase__WritingSystem = 7
const QFontDatabase__Thaana QFontDatabase__WritingSystem = 8
const QFontDatabase__Devanagari QFontDatabase__WritingSystem = 9
const QFontDatabase__Bengali QFontDatabase__WritingSystem = 10
const QFontDatabase__Gurmukhi QFontDatabase__WritingSystem = 11
const QFontDatabase__Gujarati QFontDatabase__WritingSystem = 12
const QFontDatabase__Oriya QFontDatabase__WritingSystem = 13
const QFontDatabase__Tamil QFontDatabase__WritingSystem = 14
const QFontDatabase__Telugu QFontDatabase__WritingSystem = 15
const QFontDatabase__Kannada QFontDatabase__WritingSystem = 16
const QFontDatabase__Malayalam QFontDatabase__WritingSystem = 17
const QFontDatabase__Sinhala QFontDatabase__WritingSystem = 18
const QFontDatabase__Thai QFontDatabase__WritingSystem = 19
const QFontDatabase__Lao QFontDatabase__WritingSystem = 20
const QFontDatabase__Tibetan QFontDatabase__WritingSystem = 21
const QFontDatabase__Myanmar QFontDatabase__WritingSystem = 22
const QFontDatabase__Georgian QFontDatabase__WritingSystem = 23
const QFontDatabase__Khmer QFontDatabase__WritingSystem = 24
const QFontDatabase__SimplifiedChinese QFontDatabase__WritingSystem = 25
const QFontDatabase__TraditionalChinese QFontDatabase__WritingSystem = 26
const QFontDatabase__Japanese QFontDatabase__WritingSystem = 27
const QFontDatabase__Korean QFontDatabase__WritingSystem = 28
const QFontDatabase__Vietnamese QFontDatabase__WritingSystem = 29
const QFontDatabase__Symbol QFontDatabase__WritingSystem = 30
const QFontDatabase__Other QFontDatabase__WritingSystem = 30
const QFontDatabase__Ogham QFontDatabase__WritingSystem = 31
const QFontDatabase__Runic QFontDatabase__WritingSystem = 32
const QFontDatabase__Nko QFontDatabase__WritingSystem = 33
const QFontDatabase__WritingSystemsCount QFontDatabase__WritingSystem = 34

type QFontDatabase__SystemFont = int

const QFontDatabase__GeneralFont QFontDatabase__SystemFont = 0
const QFontDatabase__FixedFont QFontDatabase__SystemFont = 1
const QFontDatabase__TitleFont QFontDatabase__SystemFont = 2
const QFontDatabase__SmallestReadableFont QFontDatabase__SystemFont = 3

type QGlyphRun__GlyphRunFlag = int

const QGlyphRun__Overline QGlyphRun__GlyphRunFlag = 1
const QGlyphRun__Underline QGlyphRun__GlyphRunFlag = 2
const QGlyphRun__StrikeOut QGlyphRun__GlyphRunFlag = 4
const QGlyphRun__RightToLeft QGlyphRun__GlyphRunFlag = 8
const QGlyphRun__SplitLigature QGlyphRun__GlyphRunFlag = 16

type QGradient__Type = int

const QGradient__LinearGradient QGradient__Type = 0
const QGradient__RadialGradient QGradient__Type = 1
const QGradient__ConicalGradient QGradient__Type = 2
const QGradient__NoGradient QGradient__Type = 3

type QGradient__Spread = int

const QGradient__PadSpread QGradient__Spread = 0
const QGradient__ReflectSpread QGradient__Spread = 1
const QGradient__RepeatSpread QGradient__Spread = 2

type QGradient__CoordinateMode = int

const QGradient__LogicalMode QGradient__CoordinateMode = 0
const QGradient__StretchToDeviceMode QGradient__CoordinateMode = 1
const QGradient__ObjectBoundingMode QGradient__CoordinateMode = 2
const QGradient__ObjectMode QGradient__CoordinateMode = 3

type QGradient__InterpolationMode = int

const QGradient__ColorInterpolation QGradient__InterpolationMode = 0
const QGradient__ComponentInterpolation QGradient__InterpolationMode = 1

type QGradient__Preset = int

const QGradient__WarmFlame QGradient__Preset = 1
const QGradient__NightFade QGradient__Preset = 2
const QGradient__SpringWarmth QGradient__Preset = 3
const QGradient__JuicyPeach QGradient__Preset = 4
const QGradient__YoungPassion QGradient__Preset = 5
const QGradient__LadyLips QGradient__Preset = 6
const QGradient__SunnyMorning QGradient__Preset = 7
const QGradient__RainyAshville QGradient__Preset = 8
const QGradient__FrozenDreams QGradient__Preset = 9
const QGradient__WinterNeva QGradient__Preset = 10
const QGradient__DustyGrass QGradient__Preset = 11
const QGradient__TemptingAzure QGradient__Preset = 12
const QGradient__HeavyRain QGradient__Preset = 13
const QGradient__AmyCrisp QGradient__Preset = 14
const QGradient__MeanFruit QGradient__Preset = 15
const QGradient__DeepBlue QGradient__Preset = 16
const QGradient__RipeMalinka QGradient__Preset = 17
const QGradient__CloudyKnoxville QGradient__Preset = 18
const QGradient__MalibuBeach QGradient__Preset = 19
const QGradient__NewLife QGradient__Preset = 20
const QGradient__TrueSunset QGradient__Preset = 21
const QGradient__MorpheusDen QGradient__Preset = 22
const QGradient__RareWind QGradient__Preset = 23
const QGradient__NearMoon QGradient__Preset = 24
const QGradient__WildApple QGradient__Preset = 25
const QGradient__SaintPetersburg QGradient__Preset = 26
const QGradient__PlumPlate QGradient__Preset = 28
const QGradient__EverlastingSky QGradient__Preset = 29
const QGradient__HappyFisher QGradient__Preset = 30
const QGradient__Blessing QGradient__Preset = 31
const QGradient__SharpeyeEagle QGradient__Preset = 32
const QGradient__LadogaBottom QGradient__Preset = 33
const QGradient__LemonGate QGradient__Preset = 34
const QGradient__ItmeoBranding QGradient__Preset = 35
const QGradient__ZeusMiracle QGradient__Preset = 36
const QGradient__OldHat QGradient__Preset = 37
const QGradient__StarWine QGradient__Preset = 38
const QGradient__HappyAcid QGradient__Preset = 41
const QGradient__AwesomePine QGradient__Preset = 42
const QGradient__NewYork QGradient__Preset = 43
const QGradient__ShyRainbow QGradient__Preset = 44
const QGradient__MixedHopes QGradient__Preset = 46
const QGradient__FlyHigh QGradient__Preset = 47
const QGradient__StrongBliss QGradient__Preset = 48
const QGradient__FreshMilk QGradient__Preset = 49
const QGradient__SnowAgain QGradient__Preset = 50
const QGradient__FebruaryInk QGradient__Preset = 51
const QGradient__KindSteel QGradient__Preset = 52
const QGradient__SoftGrass QGradient__Preset = 53
const QGradient__GrownEarly QGradient__Preset = 54
const QGradient__SharpBlues QGradient__Preset = 55
const QGradient__ShadyWater QGradient__Preset = 56
const QGradient__DirtyBeauty QGradient__Preset = 57
const QGradient__GreatWhale QGradient__Preset = 58
const QGradient__TeenNotebook QGradient__Preset = 59
const QGradient__PoliteRumors QGradient__Preset = 60
const QGradient__SweetPeriod QGradient__Preset = 61
const QGradient__WideMatrix QGradient__Preset = 62
const QGradient__SoftCherish QGradient__Preset = 63
const QGradient__RedSalvation QGradient__Preset = 64
const QGradient__BurningSpring QGradient__Preset = 65
const QGradient__NightParty QGradient__Preset = 66
const QGradient__SkyGlider QGradient__Preset = 67
const QGradient__HeavenPeach QGradient__Preset = 68
const QGradient__PurpleDivision QGradient__Preset = 69
const QGradient__AquaSplash QGradient__Preset = 70
const QGradient__SpikyNaga QGradient__Preset = 72
const QGradient__LoveKiss QGradient__Preset = 73
const QGradient__CleanMirror QGradient__Preset = 75
const QGradient__PremiumDark QGradient__Preset = 76
const QGradient__ColdEvening QGradient__Preset = 77
const QGradient__CochitiLake QGradient__Preset = 78
const QGradient__SummerGames QGradient__Preset = 79
const QGradient__PassionateBed QGradient__Preset = 80
const QGradient__MountainRock QGradient__Preset = 81
const QGradient__DesertHump QGradient__Preset = 82
const QGradient__JungleDay QGradient__Preset = 83
const QGradient__PhoenixStart QGradient__Preset = 84
const QGradient__OctoberSilence QGradient__Preset = 85
const QGradient__FarawayRiver QGradient__Preset = 86
const QGradient__AlchemistLab QGradient__Preset = 87
const QGradient__OverSun QGradient__Preset = 88
const QGradient__PremiumWhite QGradient__Preset = 89
const QGradient__MarsParty QGradient__Preset = 90
const QGradient__EternalConstance QGradient__Preset = 91
const QGradient__JapanBlush QGradient__Preset = 92
const QGradient__SmilingRain QGradient__Preset = 93
const QGradient__CloudyApple QGradient__Preset = 94
const QGradient__BigMango QGradient__Preset = 95
const QGradient__HealthyWater QGradient__Preset = 96
const QGradient__AmourAmour QGradient__Preset = 97
const QGradient__RiskyConcrete QGradient__Preset = 98
const QGradient__StrongStick QGradient__Preset = 99
const QGradient__ViciousStance QGradient__Preset = 100
const QGradient__PaloAlto QGradient__Preset = 101
const QGradient__HappyMemories QGradient__Preset = 102
const QGradient__MidnightBloom QGradient__Preset = 103
const QGradient__Crystalline QGradient__Preset = 104
const QGradient__PartyBliss QGradient__Preset = 106
const QGradient__ConfidentCloud QGradient__Preset = 107
const QGradient__LeCocktail QGradient__Preset = 108
const QGradient__RiverCity QGradient__Preset = 109
const QGradient__FrozenBerry QGradient__Preset = 110
const QGradient__ChildCare QGradient__Preset = 112
const QGradient__FlyingLemon QGradient__Preset = 113
const QGradient__NewRetrowave QGradient__Preset = 114
const QGradient__HiddenJaguar QGradient__Preset = 115
const QGradient__AboveTheSky QGradient__Preset = 116
const QGradient__Nega QGradient__Preset = 117
const QGradient__DenseWater QGradient__Preset = 118
const QGradient__Seashore QGradient__Preset = 120
const QGradient__MarbleWall QGradient__Preset = 121
const QGradient__CheerfulCaramel QGradient__Preset = 122
const QGradient__NightSky QGradient__Preset = 123
const QGradient__MagicLake QGradient__Preset = 124
const QGradient__YoungGrass QGradient__Preset = 125
const QGradient__ColorfulPeach QGradient__Preset = 126
const QGradient__GentleCare QGradient__Preset = 127
const QGradient__PlumBath QGradient__Preset = 128
const QGradient__HappyUnicorn QGradient__Preset = 129
const QGradient__AfricanField QGradient__Preset = 131
const QGradient__SolidStone QGradient__Preset = 132
const QGradient__OrangeJuice QGradient__Preset = 133
const QGradient__GlassWater QGradient__Preset = 134
const QGradient__NorthMiracle QGradient__Preset = 136
const QGradient__FruitBlend QGradient__Preset = 137
const QGradient__MillenniumPine QGradient__Preset = 138
const QGradient__HighFlight QGradient__Preset = 139
const QGradient__MoleHall QGradient__Preset = 140
const QGradient__SpaceShift QGradient__Preset = 142
const QGradient__ForestInei QGradient__Preset = 143
const QGradient__RoyalGarden QGradient__Preset = 144
const QGradient__RichMetal QGradient__Preset = 145
const QGradient__JuicyCake QGradient__Preset = 146
const QGradient__SmartIndigo QGradient__Preset = 147
const QGradient__SandStrike QGradient__Preset = 148
const QGradient__NorseBeauty QGradient__Preset = 149
const QGradient__AquaGuidance QGradient__Preset = 150
const QGradient__SunVeggie QGradient__Preset = 151
const QGradient__SeaLord QGradient__Preset = 152
const QGradient__BlackSea QGradient__Preset = 153
const QGradient__GrassShampoo QGradient__Preset = 154
const QGradient__LandingAircraft QGradient__Preset = 155
const QGradient__WitchDance QGradient__Preset = 156
const QGradient__SleeplessNight QGradient__Preset = 157
const QGradient__AngelCare QGradient__Preset = 158
const QGradient__CrystalRiver QGradient__Preset = 159
const QGradient__SoftLipstick QGradient__Preset = 160
const QGradient__SaltMountain QGradient__Preset = 161
const QGradient__PerfectWhite QGradient__Preset = 162
const QGradient__FreshOasis QGradient__Preset = 163
const QGradient__StrictNovember QGradient__Preset = 164
const QGradient__MorningSalad QGradient__Preset = 165
const QGradient__DeepRelief QGradient__Preset = 166
const QGradient__SeaStrike QGradient__Preset = 167
const QGradient__NightCall QGradient__Preset = 168
const QGradient__SupremeSky QGradient__Preset = 169
const QGradient__LightBlue QGradient__Preset = 170
const QGradient__MindCrawl QGradient__Preset = 171
const QGradient__LilyMeadow QGradient__Preset = 172
const QGradient__SugarLollipop QGradient__Preset = 173
const QGradient__SweetDessert QGradient__Preset = 174
const QGradient__MagicRay QGradient__Preset = 175
const QGradient__TeenParty QGradient__Preset = 176
const QGradient__FrozenHeat QGradient__Preset = 177
const QGradient__GagarinView QGradient__Preset = 178
const QGradient__FabledSunset QGradient__Preset = 179
const QGradient__PerfectBlue QGradient__Preset = 180

type QIcon__Mode = int

const QIcon__Normal QIcon__Mode = 0
const QIcon__Disabled QIcon__Mode = 1
const QIcon__Active QIcon__Mode = 2
const QIcon__Selected QIcon__Mode = 3

type QIcon__State = int

const QIcon__On QIcon__State = 0
const QIcon__Off QIcon__State = 1

type QIconEngine__IconEngineHook = int

const QIconEngine__AvailableSizesHook QIconEngine__IconEngineHook = 1
const QIconEngine__IconNameHook QIconEngine__IconEngineHook = 2
const QIconEngine__IsNullHook QIconEngine__IconEngineHook = 3
const QIconEngine__ScaledPixmapHook QIconEngine__IconEngineHook = 4

type QImage__InvertMode = int

const QImage__InvertRgb QImage__InvertMode = 0
const QImage__InvertRgba QImage__InvertMode = 1

type QImage__Format = int

const QImage__Format_Invalid QImage__Format = 0
const QImage__Format_Mono QImage__Format = 1
const QImage__Format_MonoLSB QImage__Format = 2
const QImage__Format_Indexed8 QImage__Format = 3
const QImage__Format_RGB32 QImage__Format = 4
const QImage__Format_ARGB32 QImage__Format = 5
const QImage__Format_ARGB32_Premultiplied QImage__Format = 6
const QImage__Format_RGB16 QImage__Format = 7
const QImage__Format_ARGB8565_Premultiplied QImage__Format = 8
const QImage__Format_RGB666 QImage__Format = 9
const QImage__Format_ARGB6666_Premultiplied QImage__Format = 10
const QImage__Format_RGB555 QImage__Format = 11
const QImage__Format_ARGB8555_Premultiplied QImage__Format = 12
const QImage__Format_RGB888 QImage__Format = 13
const QImage__Format_RGB444 QImage__Format = 14
const QImage__Format_ARGB4444_Premultiplied QImage__Format = 15
const QImage__Format_RGBX8888 QImage__Format = 16
const QImage__Format_RGBA8888 QImage__Format = 17
const QImage__Format_RGBA8888_Premultiplied QImage__Format = 18
const QImage__Format_BGR30 QImage__Format = 19
const QImage__Format_A2BGR30_Premultiplied QImage__Format = 20
const QImage__Format_RGB30 QImage__Format = 21
const QImage__Format_A2RGB30_Premultiplied QImage__Format = 22
const QImage__Format_Alpha8 QImage__Format = 23
const QImage__Format_Grayscale8 QImage__Format = 24
const QImage__Format_RGBX64 QImage__Format = 25
const QImage__Format_RGBA64 QImage__Format = 26
const QImage__Format_RGBA64_Premultiplied QImage__Format = 27
const QImage__NImageFormats QImage__Format = 28

type QImageIOHandler__ImageOption = int

const QImageIOHandler__Size QImageIOHandler__ImageOption = 0
const QImageIOHandler__ClipRect QImageIOHandler__ImageOption = 1
const QImageIOHandler__Description QImageIOHandler__ImageOption = 2
const QImageIOHandler__ScaledClipRect QImageIOHandler__ImageOption = 3
const QImageIOHandler__ScaledSize QImageIOHandler__ImageOption = 4
const QImageIOHandler__CompressionRatio QImageIOHandler__ImageOption = 5
const QImageIOHandler__Gamma QImageIOHandler__ImageOption = 6
const QImageIOHandler__Quality QImageIOHandler__ImageOption = 7
const QImageIOHandler__Name QImageIOHandler__ImageOption = 8
const QImageIOHandler__SubType QImageIOHandler__ImageOption = 9
const QImageIOHandler__IncrementalReading QImageIOHandler__ImageOption = 10
const QImageIOHandler__Endianness QImageIOHandler__ImageOption = 11
const QImageIOHandler__Animation QImageIOHandler__ImageOption = 12
const QImageIOHandler__BackgroundColor QImageIOHandler__ImageOption = 13
const QImageIOHandler__ImageFormat QImageIOHandler__ImageOption = 14
const QImageIOHandler__SupportedSubTypes QImageIOHandler__ImageOption = 15
const QImageIOHandler__OptimizedWrite QImageIOHandler__ImageOption = 16
const QImageIOHandler__ProgressiveScanWrite QImageIOHandler__ImageOption = 17
const QImageIOHandler__ImageTransformation QImageIOHandler__ImageOption = 18
const QImageIOHandler__TransformedByDefault QImageIOHandler__ImageOption = 19

type QImageIOHandler__Transformation = int

const QImageIOHandler__TransformationNone QImageIOHandler__Transformation = 0
const QImageIOHandler__TransformationMirror QImageIOHandler__Transformation = 1
const QImageIOHandler__TransformationFlip QImageIOHandler__Transformation = 2
const QImageIOHandler__TransformationRotate180 QImageIOHandler__Transformation = 3
const QImageIOHandler__TransformationRotate90 QImageIOHandler__Transformation = 4
const QImageIOHandler__TransformationMirrorAndRotate90 QImageIOHandler__Transformation = 5
const QImageIOHandler__TransformationFlipAndRotate90 QImageIOHandler__Transformation = 6
const QImageIOHandler__TransformationRotate270 QImageIOHandler__Transformation = 7

type QImageIOPlugin__Capability = int

const QImageIOPlugin__CanRead QImageIOPlugin__Capability = 1
const QImageIOPlugin__CanWrite QImageIOPlugin__Capability = 2
const QImageIOPlugin__CanReadIncremental QImageIOPlugin__Capability = 4

type QImageReader__ImageReaderError = int

const QImageReader__UnknownError QImageReader__ImageReaderError = 0
const QImageReader__FileNotFoundError QImageReader__ImageReaderError = 1
const QImageReader__DeviceError QImageReader__ImageReaderError = 2
const QImageReader__UnsupportedFormatError QImageReader__ImageReaderError = 3
const QImageReader__InvalidDataError QImageReader__ImageReaderError = 4

type QImageWriter__ImageWriterError = int

const QImageWriter__UnknownError QImageWriter__ImageWriterError = 0
const QImageWriter__DeviceError QImageWriter__ImageWriterError = 1
const QImageWriter__UnsupportedFormatError QImageWriter__ImageWriterError = 2
const QImageWriter__InvalidImageError QImageWriter__ImageWriterError = 3

type QInputMethod__Action = int

const QInputMethod__Click QInputMethod__Action = 0
const QInputMethod__ContextMenu QInputMethod__Action = 1

type QInputMethodEvent__AttributeType = int

const QInputMethodEvent__TextFormat QInputMethodEvent__AttributeType = 0
const QInputMethodEvent__Cursor QInputMethodEvent__AttributeType = 1
const QInputMethodEvent__Language QInputMethodEvent__AttributeType = 2
const QInputMethodEvent__Ruby QInputMethodEvent__AttributeType = 3
const QInputMethodEvent__Selection QInputMethodEvent__AttributeType = 4

type QKeySequence__StandardKey = int

const QKeySequence__UnknownKey QKeySequence__StandardKey = 0
const QKeySequence__HelpContents QKeySequence__StandardKey = 1
const QKeySequence__WhatsThis QKeySequence__StandardKey = 2
const QKeySequence__Open QKeySequence__StandardKey = 3
const QKeySequence__Close QKeySequence__StandardKey = 4
const QKeySequence__Save QKeySequence__StandardKey = 5
const QKeySequence__New QKeySequence__StandardKey = 6
const QKeySequence__Delete QKeySequence__StandardKey = 7
const QKeySequence__Cut QKeySequence__StandardKey = 8
const QKeySequence__Copy QKeySequence__StandardKey = 9
const QKeySequence__Paste QKeySequence__StandardKey = 10
const QKeySequence__Undo QKeySequence__StandardKey = 11
const QKeySequence__Redo QKeySequence__StandardKey = 12
const QKeySequence__Back QKeySequence__StandardKey = 13
const QKeySequence__Forward QKeySequence__StandardKey = 14
const QKeySequence__Refresh QKeySequence__StandardKey = 15
const QKeySequence__ZoomIn QKeySequence__StandardKey = 16
const QKeySequence__ZoomOut QKeySequence__StandardKey = 17
const QKeySequence__Print QKeySequence__StandardKey = 18
const QKeySequence__AddTab QKeySequence__StandardKey = 19
const QKeySequence__NextChild QKeySequence__StandardKey = 20
const QKeySequence__PreviousChild QKeySequence__StandardKey = 21
const QKeySequence__Find QKeySequence__StandardKey = 22
const QKeySequence__FindNext QKeySequence__StandardKey = 23
const QKeySequence__FindPrevious QKeySequence__StandardKey = 24
const QKeySequence__Replace QKeySequence__StandardKey = 25
const QKeySequence__SelectAll QKeySequence__StandardKey = 26
const QKeySequence__Bold QKeySequence__StandardKey = 27
const QKeySequence__Italic QKeySequence__StandardKey = 28
const QKeySequence__Underline QKeySequence__StandardKey = 29
const QKeySequence__MoveToNextChar QKeySequence__StandardKey = 30
const QKeySequence__MoveToPreviousChar QKeySequence__StandardKey = 31
const QKeySequence__MoveToNextWord QKeySequence__StandardKey = 32
const QKeySequence__MoveToPreviousWord QKeySequence__StandardKey = 33
const QKeySequence__MoveToNextLine QKeySequence__StandardKey = 34
const QKeySequence__MoveToPreviousLine QKeySequence__StandardKey = 35
const QKeySequence__MoveToNextPage QKeySequence__StandardKey = 36
const QKeySequence__MoveToPreviousPage QKeySequence__StandardKey = 37
const QKeySequence__MoveToStartOfLine QKeySequence__StandardKey = 38
const QKeySequence__MoveToEndOfLine QKeySequence__StandardKey = 39
const QKeySequence__MoveToStartOfBlock QKeySequence__StandardKey = 40
const QKeySequence__MoveToEndOfBlock QKeySequence__StandardKey = 41
const QKeySequence__MoveToStartOfDocument QKeySequence__StandardKey = 42
const QKeySequence__MoveToEndOfDocument QKeySequence__StandardKey = 43
const QKeySequence__SelectNextChar QKeySequence__StandardKey = 44
const QKeySequence__SelectPreviousChar QKeySequence__StandardKey = 45
const QKeySequence__SelectNextWord QKeySequence__StandardKey = 46
const QKeySequence__SelectPreviousWord QKeySequence__StandardKey = 47
const QKeySequence__SelectNextLine QKeySequence__StandardKey = 48
const QKeySequence__SelectPreviousLine QKeySequence__StandardKey = 49
const QKeySequence__SelectNextPage QKeySequence__StandardKey = 50
const QKeySequence__SelectPreviousPage QKeySequence__StandardKey = 51
const QKeySequence__SelectStartOfLine QKeySequence__StandardKey = 52
const QKeySequence__SelectEndOfLine QKeySequence__StandardKey = 53
const QKeySequence__SelectStartOfBlock QKeySequence__StandardKey = 54
const QKeySequence__SelectEndOfBlock QKeySequence__StandardKey = 55
const QKeySequence__SelectStartOfDocument QKeySequence__StandardKey = 56
const QKeySequence__SelectEndOfDocument QKeySequence__StandardKey = 57
const QKeySequence__DeleteStartOfWord QKeySequence__StandardKey = 58
const QKeySequence__DeleteEndOfWord QKeySequence__StandardKey = 59
const QKeySequence__DeleteEndOfLine QKeySequence__StandardKey = 60
const QKeySequence__InsertParagraphSeparator QKeySequence__StandardKey = 61
const QKeySequence__InsertLineSeparator QKeySequence__StandardKey = 62
const QKeySequence__SaveAs QKeySequence__StandardKey = 63
const QKeySequence__Preferences QKeySequence__StandardKey = 64
const QKeySequence__Quit QKeySequence__StandardKey = 65
const QKeySequence__FullScreen QKeySequence__StandardKey = 66
const QKeySequence__Deselect QKeySequence__StandardKey = 67
const QKeySequence__DeleteCompleteLine QKeySequence__StandardKey = 68
const QKeySequence__Backspace QKeySequence__StandardKey = 69
const QKeySequence__Cancel QKeySequence__StandardKey = 70

type QKeySequence__SequenceFormat = int

const QKeySequence__NativeText QKeySequence__SequenceFormat = 0
const QKeySequence__PortableText QKeySequence__SequenceFormat = 1

type QKeySequence__SequenceMatch = int

const QKeySequence__NoMatch QKeySequence__SequenceMatch = 0
const QKeySequence__PartialMatch QKeySequence__SequenceMatch = 1
const QKeySequence__ExactMatch QKeySequence__SequenceMatch = 2

type QMatrix4x4__ = int

const QMatrix4x4__Identity QMatrix4x4__ = 0
const QMatrix4x4__Translation QMatrix4x4__ = 1
const QMatrix4x4__Scale QMatrix4x4__ = 2
const QMatrix4x4__Rotation2D QMatrix4x4__ = 4
const QMatrix4x4__Rotation QMatrix4x4__ = 8
const QMatrix4x4__Perspective QMatrix4x4__ = 16
const QMatrix4x4__General QMatrix4x4__ = 31

type QMovie__MovieState = int

const QMovie__NotRunning QMovie__MovieState = 0
const QMovie__Paused QMovie__MovieState = 1
const QMovie__Running QMovie__MovieState = 2

type QMovie__CacheMode = int

const QMovie__CacheNone QMovie__CacheMode = 0
const QMovie__CacheAll QMovie__CacheMode = 1

type QPagedPaintDevice__PageSize = int

const QPagedPaintDevice__A4 QPagedPaintDevice__PageSize = 0
const QPagedPaintDevice__B5 QPagedPaintDevice__PageSize = 1
const QPagedPaintDevice__Letter QPagedPaintDevice__PageSize = 2
const QPagedPaintDevice__Legal QPagedPaintDevice__PageSize = 3
const QPagedPaintDevice__Executive QPagedPaintDevice__PageSize = 4
const QPagedPaintDevice__A0 QPagedPaintDevice__PageSize = 5
const QPagedPaintDevice__A1 QPagedPaintDevice__PageSize = 6
const QPagedPaintDevice__A2 QPagedPaintDevice__PageSize = 7
const QPagedPaintDevice__A3 QPagedPaintDevice__PageSize = 8
const QPagedPaintDevice__A5 QPagedPaintDevice__PageSize = 9
const QPagedPaintDevice__A6 QPagedPaintDevice__PageSize = 10
const QPagedPaintDevice__A7 QPagedPaintDevice__PageSize = 11
const QPagedPaintDevice__A8 QPagedPaintDevice__PageSize = 12
const QPagedPaintDevice__A9 QPagedPaintDevice__PageSize = 13
const QPagedPaintDevice__B0 QPagedPaintDevice__PageSize = 14
const QPagedPaintDevice__B1 QPagedPaintDevice__PageSize = 15
const QPagedPaintDevice__B10 QPagedPaintDevice__PageSize = 16
const QPagedPaintDevice__B2 QPagedPaintDevice__PageSize = 17
const QPagedPaintDevice__B3 QPagedPaintDevice__PageSize = 18
const QPagedPaintDevice__B4 QPagedPaintDevice__PageSize = 19
const QPagedPaintDevice__B6 QPagedPaintDevice__PageSize = 20
const QPagedPaintDevice__B7 QPagedPaintDevice__PageSize = 21
const QPagedPaintDevice__B8 QPagedPaintDevice__PageSize = 22
const QPagedPaintDevice__B9 QPagedPaintDevice__PageSize = 23
const QPagedPaintDevice__C5E QPagedPaintDevice__PageSize = 24
const QPagedPaintDevice__Comm10E QPagedPaintDevice__PageSize = 25
const QPagedPaintDevice__DLE QPagedPaintDevice__PageSize = 26
const QPagedPaintDevice__Folio QPagedPaintDevice__PageSize = 27
const QPagedPaintDevice__Ledger QPagedPaintDevice__PageSize = 28
const QPagedPaintDevice__Tabloid QPagedPaintDevice__PageSize = 29
const QPagedPaintDevice__Custom QPagedPaintDevice__PageSize = 30
const QPagedPaintDevice__A10 QPagedPaintDevice__PageSize = 31
const QPagedPaintDevice__A3Extra QPagedPaintDevice__PageSize = 32
const QPagedPaintDevice__A4Extra QPagedPaintDevice__PageSize = 33
const QPagedPaintDevice__A4Plus QPagedPaintDevice__PageSize = 34
const QPagedPaintDevice__A4Small QPagedPaintDevice__PageSize = 35
const QPagedPaintDevice__A5Extra QPagedPaintDevice__PageSize = 36
const QPagedPaintDevice__B5Extra QPagedPaintDevice__PageSize = 37
const QPagedPaintDevice__JisB0 QPagedPaintDevice__PageSize = 38
const QPagedPaintDevice__JisB1 QPagedPaintDevice__PageSize = 39
const QPagedPaintDevice__JisB2 QPagedPaintDevice__PageSize = 40
const QPagedPaintDevice__JisB3 QPagedPaintDevice__PageSize = 41
const QPagedPaintDevice__JisB4 QPagedPaintDevice__PageSize = 42
const QPagedPaintDevice__JisB5 QPagedPaintDevice__PageSize = 43
const QPagedPaintDevice__JisB6 QPagedPaintDevice__PageSize = 44
const QPagedPaintDevice__JisB7 QPagedPaintDevice__PageSize = 45
const QPagedPaintDevice__JisB8 QPagedPaintDevice__PageSize = 46
const QPagedPaintDevice__JisB9 QPagedPaintDevice__PageSize = 47
const QPagedPaintDevice__JisB10 QPagedPaintDevice__PageSize = 48
const QPagedPaintDevice__AnsiC QPagedPaintDevice__PageSize = 49
const QPagedPaintDevice__AnsiD QPagedPaintDevice__PageSize = 50
const QPagedPaintDevice__AnsiE QPagedPaintDevice__PageSize = 51
const QPagedPaintDevice__LegalExtra QPagedPaintDevice__PageSize = 52
const QPagedPaintDevice__LetterExtra QPagedPaintDevice__PageSize = 53
const QPagedPaintDevice__LetterPlus QPagedPaintDevice__PageSize = 54
const QPagedPaintDevice__LetterSmall QPagedPaintDevice__PageSize = 55
const QPagedPaintDevice__TabloidExtra QPagedPaintDevice__PageSize = 56
const QPagedPaintDevice__ArchA QPagedPaintDevice__PageSize = 57
const QPagedPaintDevice__ArchB QPagedPaintDevice__PageSize = 58
const QPagedPaintDevice__ArchC QPagedPaintDevice__PageSize = 59
const QPagedPaintDevice__ArchD QPagedPaintDevice__PageSize = 60
const QPagedPaintDevice__ArchE QPagedPaintDevice__PageSize = 61
const QPagedPaintDevice__Imperial7x9 QPagedPaintDevice__PageSize = 62
const QPagedPaintDevice__Imperial8x10 QPagedPaintDevice__PageSize = 63
const QPagedPaintDevice__Imperial9x11 QPagedPaintDevice__PageSize = 64
const QPagedPaintDevice__Imperial9x12 QPagedPaintDevice__PageSize = 65
const QPagedPaintDevice__Imperial10x11 QPagedPaintDevice__PageSize = 66
const QPagedPaintDevice__Imperial10x13 QPagedPaintDevice__PageSize = 67
const QPagedPaintDevice__Imperial10x14 QPagedPaintDevice__PageSize = 68
const QPagedPaintDevice__Imperial12x11 QPagedPaintDevice__PageSize = 69
const QPagedPaintDevice__Imperial15x11 QPagedPaintDevice__PageSize = 70
const QPagedPaintDevice__ExecutiveStandard QPagedPaintDevice__PageSize = 71
const QPagedPaintDevice__Note QPagedPaintDevice__PageSize = 72
const QPagedPaintDevice__Quarto QPagedPaintDevice__PageSize = 73
const QPagedPaintDevice__Statement QPagedPaintDevice__PageSize = 74
const QPagedPaintDevice__SuperA QPagedPaintDevice__PageSize = 75
const QPagedPaintDevice__SuperB QPagedPaintDevice__PageSize = 76
const QPagedPaintDevice__Postcard QPagedPaintDevice__PageSize = 77
const QPagedPaintDevice__DoublePostcard QPagedPaintDevice__PageSize = 78
const QPagedPaintDevice__Prc16K QPagedPaintDevice__PageSize = 79
const QPagedPaintDevice__Prc32K QPagedPaintDevice__PageSize = 80
const QPagedPaintDevice__Prc32KBig QPagedPaintDevice__PageSize = 81
const QPagedPaintDevice__FanFoldUS QPagedPaintDevice__PageSize = 82
const QPagedPaintDevice__FanFoldGerman QPagedPaintDevice__PageSize = 83
const QPagedPaintDevice__FanFoldGermanLegal QPagedPaintDevice__PageSize = 84
const QPagedPaintDevice__EnvelopeB4 QPagedPaintDevice__PageSize = 85
const QPagedPaintDevice__EnvelopeB5 QPagedPaintDevice__PageSize = 86
const QPagedPaintDevice__EnvelopeB6 QPagedPaintDevice__PageSize = 87
const QPagedPaintDevice__EnvelopeC0 QPagedPaintDevice__PageSize = 88
const QPagedPaintDevice__EnvelopeC1 QPagedPaintDevice__PageSize = 89
const QPagedPaintDevice__EnvelopeC2 QPagedPaintDevice__PageSize = 90
const QPagedPaintDevice__EnvelopeC3 QPagedPaintDevice__PageSize = 91
const QPagedPaintDevice__EnvelopeC4 QPagedPaintDevice__PageSize = 92
const QPagedPaintDevice__EnvelopeC6 QPagedPaintDevice__PageSize = 93
const QPagedPaintDevice__EnvelopeC65 QPagedPaintDevice__PageSize = 94
const QPagedPaintDevice__EnvelopeC7 QPagedPaintDevice__PageSize = 95
const QPagedPaintDevice__Envelope9 QPagedPaintDevice__PageSize = 96
const QPagedPaintDevice__Envelope11 QPagedPaintDevice__PageSize = 97
const QPagedPaintDevice__Envelope12 QPagedPaintDevice__PageSize = 98
const QPagedPaintDevice__Envelope14 QPagedPaintDevice__PageSize = 99
const QPagedPaintDevice__EnvelopeMonarch QPagedPaintDevice__PageSize = 100
const QPagedPaintDevice__EnvelopePersonal QPagedPaintDevice__PageSize = 101
const QPagedPaintDevice__EnvelopeChou3 QPagedPaintDevice__PageSize = 102
const QPagedPaintDevice__EnvelopeChou4 QPagedPaintDevice__PageSize = 103
const QPagedPaintDevice__EnvelopeInvite QPagedPaintDevice__PageSize = 104
const QPagedPaintDevice__EnvelopeItalian QPagedPaintDevice__PageSize = 105
const QPagedPaintDevice__EnvelopeKaku2 QPagedPaintDevice__PageSize = 106
const QPagedPaintDevice__EnvelopeKaku3 QPagedPaintDevice__PageSize = 107
const QPagedPaintDevice__EnvelopePrc1 QPagedPaintDevice__PageSize = 108
const QPagedPaintDevice__EnvelopePrc2 QPagedPaintDevice__PageSize = 109
const QPagedPaintDevice__EnvelopePrc3 QPagedPaintDevice__PageSize = 110
const QPagedPaintDevice__EnvelopePrc4 QPagedPaintDevice__PageSize = 111
const QPagedPaintDevice__EnvelopePrc5 QPagedPaintDevice__PageSize = 112
const QPagedPaintDevice__EnvelopePrc6 QPagedPaintDevice__PageSize = 113
const QPagedPaintDevice__EnvelopePrc7 QPagedPaintDevice__PageSize = 114
const QPagedPaintDevice__EnvelopePrc8 QPagedPaintDevice__PageSize = 115
const QPagedPaintDevice__EnvelopePrc9 QPagedPaintDevice__PageSize = 116
const QPagedPaintDevice__EnvelopePrc10 QPagedPaintDevice__PageSize = 117
const QPagedPaintDevice__EnvelopeYou4 QPagedPaintDevice__PageSize = 118
const QPagedPaintDevice__LastPageSize QPagedPaintDevice__PageSize = 118
const QPagedPaintDevice__NPageSize QPagedPaintDevice__PageSize = 118
const QPagedPaintDevice__NPaperSize QPagedPaintDevice__PageSize = 118
const QPagedPaintDevice__AnsiA QPagedPaintDevice__PageSize = 2
const QPagedPaintDevice__AnsiB QPagedPaintDevice__PageSize = 28
const QPagedPaintDevice__EnvelopeC5 QPagedPaintDevice__PageSize = 24
const QPagedPaintDevice__EnvelopeDL QPagedPaintDevice__PageSize = 26
const QPagedPaintDevice__Envelope10 QPagedPaintDevice__PageSize = 25

type QPagedPaintDevice__PdfVersion = int

const QPagedPaintDevice__PdfVersion_1_4 QPagedPaintDevice__PdfVersion = 0
const QPagedPaintDevice__PdfVersion_A1b QPagedPaintDevice__PdfVersion = 1
const QPagedPaintDevice__PdfVersion_1_6 QPagedPaintDevice__PdfVersion = 2

type QPageLayout__Unit = int

const QPageLayout__Millimeter QPageLayout__Unit = 0
const QPageLayout__Point QPageLayout__Unit = 1
const QPageLayout__Inch QPageLayout__Unit = 2
const QPageLayout__Pica QPageLayout__Unit = 3
const QPageLayout__Didot QPageLayout__Unit = 4
const QPageLayout__Cicero QPageLayout__Unit = 5

type QPageLayout__Orientation = int

const QPageLayout__Portrait QPageLayout__Orientation = 0
const QPageLayout__Landscape QPageLayout__Orientation = 1

type QPageLayout__Mode = int

const QPageLayout__StandardMode QPageLayout__Mode = 0
const QPageLayout__FullPageMode QPageLayout__Mode = 1

type QPageSize__PageSizeId = int

const QPageSize__A4 QPageSize__PageSizeId = 0
const QPageSize__B5 QPageSize__PageSizeId = 1
const QPageSize__Letter QPageSize__PageSizeId = 2
const QPageSize__Legal QPageSize__PageSizeId = 3
const QPageSize__Executive QPageSize__PageSizeId = 4
const QPageSize__A0 QPageSize__PageSizeId = 5
const QPageSize__A1 QPageSize__PageSizeId = 6
const QPageSize__A2 QPageSize__PageSizeId = 7
const QPageSize__A3 QPageSize__PageSizeId = 8
const QPageSize__A5 QPageSize__PageSizeId = 9
const QPageSize__A6 QPageSize__PageSizeId = 10
const QPageSize__A7 QPageSize__PageSizeId = 11
const QPageSize__A8 QPageSize__PageSizeId = 12
const QPageSize__A9 QPageSize__PageSizeId = 13
const QPageSize__B0 QPageSize__PageSizeId = 14
const QPageSize__B1 QPageSize__PageSizeId = 15
const QPageSize__B10 QPageSize__PageSizeId = 16
const QPageSize__B2 QPageSize__PageSizeId = 17
const QPageSize__B3 QPageSize__PageSizeId = 18
const QPageSize__B4 QPageSize__PageSizeId = 19
const QPageSize__B6 QPageSize__PageSizeId = 20
const QPageSize__B7 QPageSize__PageSizeId = 21
const QPageSize__B8 QPageSize__PageSizeId = 22
const QPageSize__B9 QPageSize__PageSizeId = 23
const QPageSize__C5E QPageSize__PageSizeId = 24
const QPageSize__Comm10E QPageSize__PageSizeId = 25
const QPageSize__DLE QPageSize__PageSizeId = 26
const QPageSize__Folio QPageSize__PageSizeId = 27
const QPageSize__Ledger QPageSize__PageSizeId = 28
const QPageSize__Tabloid QPageSize__PageSizeId = 29
const QPageSize__Custom QPageSize__PageSizeId = 30
const QPageSize__A10 QPageSize__PageSizeId = 31
const QPageSize__A3Extra QPageSize__PageSizeId = 32
const QPageSize__A4Extra QPageSize__PageSizeId = 33
const QPageSize__A4Plus QPageSize__PageSizeId = 34
const QPageSize__A4Small QPageSize__PageSizeId = 35
const QPageSize__A5Extra QPageSize__PageSizeId = 36
const QPageSize__B5Extra QPageSize__PageSizeId = 37
const QPageSize__JisB0 QPageSize__PageSizeId = 38
const QPageSize__JisB1 QPageSize__PageSizeId = 39
const QPageSize__JisB2 QPageSize__PageSizeId = 40
const QPageSize__JisB3 QPageSize__PageSizeId = 41
const QPageSize__JisB4 QPageSize__PageSizeId = 42
const QPageSize__JisB5 QPageSize__PageSizeId = 43
const QPageSize__JisB6 QPageSize__PageSizeId = 44
const QPageSize__JisB7 QPageSize__PageSizeId = 45
const QPageSize__JisB8 QPageSize__PageSizeId = 46
const QPageSize__JisB9 QPageSize__PageSizeId = 47
const QPageSize__JisB10 QPageSize__PageSizeId = 48
const QPageSize__AnsiC QPageSize__PageSizeId = 49
const QPageSize__AnsiD QPageSize__PageSizeId = 50
const QPageSize__AnsiE QPageSize__PageSizeId = 51
const QPageSize__LegalExtra QPageSize__PageSizeId = 52
const QPageSize__LetterExtra QPageSize__PageSizeId = 53
const QPageSize__LetterPlus QPageSize__PageSizeId = 54
const QPageSize__LetterSmall QPageSize__PageSizeId = 55
const QPageSize__TabloidExtra QPageSize__PageSizeId = 56
const QPageSize__ArchA QPageSize__PageSizeId = 57
const QPageSize__ArchB QPageSize__PageSizeId = 58
const QPageSize__ArchC QPageSize__PageSizeId = 59
const QPageSize__ArchD QPageSize__PageSizeId = 60
const QPageSize__ArchE QPageSize__PageSizeId = 61
const QPageSize__Imperial7x9 QPageSize__PageSizeId = 62
const QPageSize__Imperial8x10 QPageSize__PageSizeId = 63
const QPageSize__Imperial9x11 QPageSize__PageSizeId = 64
const QPageSize__Imperial9x12 QPageSize__PageSizeId = 65
const QPageSize__Imperial10x11 QPageSize__PageSizeId = 66
const QPageSize__Imperial10x13 QPageSize__PageSizeId = 67
const QPageSize__Imperial10x14 QPageSize__PageSizeId = 68
const QPageSize__Imperial12x11 QPageSize__PageSizeId = 69
const QPageSize__Imperial15x11 QPageSize__PageSizeId = 70
const QPageSize__ExecutiveStandard QPageSize__PageSizeId = 71
const QPageSize__Note QPageSize__PageSizeId = 72
const QPageSize__Quarto QPageSize__PageSizeId = 73
const QPageSize__Statement QPageSize__PageSizeId = 74
const QPageSize__SuperA QPageSize__PageSizeId = 75
const QPageSize__SuperB QPageSize__PageSizeId = 76
const QPageSize__Postcard QPageSize__PageSizeId = 77
const QPageSize__DoublePostcard QPageSize__PageSizeId = 78
const QPageSize__Prc16K QPageSize__PageSizeId = 79
const QPageSize__Prc32K QPageSize__PageSizeId = 80
const QPageSize__Prc32KBig QPageSize__PageSizeId = 81
const QPageSize__FanFoldUS QPageSize__PageSizeId = 82
const QPageSize__FanFoldGerman QPageSize__PageSizeId = 83
const QPageSize__FanFoldGermanLegal QPageSize__PageSizeId = 84
const QPageSize__EnvelopeB4 QPageSize__PageSizeId = 85
const QPageSize__EnvelopeB5 QPageSize__PageSizeId = 86
const QPageSize__EnvelopeB6 QPageSize__PageSizeId = 87
const QPageSize__EnvelopeC0 QPageSize__PageSizeId = 88
const QPageSize__EnvelopeC1 QPageSize__PageSizeId = 89
const QPageSize__EnvelopeC2 QPageSize__PageSizeId = 90
const QPageSize__EnvelopeC3 QPageSize__PageSizeId = 91
const QPageSize__EnvelopeC4 QPageSize__PageSizeId = 92
const QPageSize__EnvelopeC6 QPageSize__PageSizeId = 93
const QPageSize__EnvelopeC65 QPageSize__PageSizeId = 94
const QPageSize__EnvelopeC7 QPageSize__PageSizeId = 95
const QPageSize__Envelope9 QPageSize__PageSizeId = 96
const QPageSize__Envelope11 QPageSize__PageSizeId = 97
const QPageSize__Envelope12 QPageSize__PageSizeId = 98
const QPageSize__Envelope14 QPageSize__PageSizeId = 99
const QPageSize__EnvelopeMonarch QPageSize__PageSizeId = 100
const QPageSize__EnvelopePersonal QPageSize__PageSizeId = 101
const QPageSize__EnvelopeChou3 QPageSize__PageSizeId = 102
const QPageSize__EnvelopeChou4 QPageSize__PageSizeId = 103
const QPageSize__EnvelopeInvite QPageSize__PageSizeId = 104
const QPageSize__EnvelopeItalian QPageSize__PageSizeId = 105
const QPageSize__EnvelopeKaku2 QPageSize__PageSizeId = 106
const QPageSize__EnvelopeKaku3 QPageSize__PageSizeId = 107
const QPageSize__EnvelopePrc1 QPageSize__PageSizeId = 108
const QPageSize__EnvelopePrc2 QPageSize__PageSizeId = 109
const QPageSize__EnvelopePrc3 QPageSize__PageSizeId = 110
const QPageSize__EnvelopePrc4 QPageSize__PageSizeId = 111
const QPageSize__EnvelopePrc5 QPageSize__PageSizeId = 112
const QPageSize__EnvelopePrc6 QPageSize__PageSizeId = 113
const QPageSize__EnvelopePrc7 QPageSize__PageSizeId = 114
const QPageSize__EnvelopePrc8 QPageSize__PageSizeId = 115
const QPageSize__EnvelopePrc9 QPageSize__PageSizeId = 116
const QPageSize__EnvelopePrc10 QPageSize__PageSizeId = 117
const QPageSize__EnvelopeYou4 QPageSize__PageSizeId = 118
const QPageSize__LastPageSize QPageSize__PageSizeId = 118
const QPageSize__NPageSize QPageSize__PageSizeId = 118
const QPageSize__NPaperSize QPageSize__PageSizeId = 118
const QPageSize__AnsiA QPageSize__PageSizeId = 2
const QPageSize__AnsiB QPageSize__PageSizeId = 28
const QPageSize__EnvelopeC5 QPageSize__PageSizeId = 24
const QPageSize__EnvelopeDL QPageSize__PageSizeId = 26
const QPageSize__Envelope10 QPageSize__PageSizeId = 25

type QPageSize__Unit = int

const QPageSize__Millimeter QPageSize__Unit = 0
const QPageSize__Point QPageSize__Unit = 1
const QPageSize__Inch QPageSize__Unit = 2
const QPageSize__Pica QPageSize__Unit = 3
const QPageSize__Didot QPageSize__Unit = 4
const QPageSize__Cicero QPageSize__Unit = 5

type QPageSize__SizeMatchPolicy = int

const QPageSize__FuzzyMatch QPageSize__SizeMatchPolicy = 0
const QPageSize__FuzzyOrientationMatch QPageSize__SizeMatchPolicy = 1
const QPageSize__ExactMatch QPageSize__SizeMatchPolicy = 2

type QPaintDevice__PaintDeviceMetric = int

const QPaintDevice__PdmWidth QPaintDevice__PaintDeviceMetric = 1
const QPaintDevice__PdmHeight QPaintDevice__PaintDeviceMetric = 2
const QPaintDevice__PdmWidthMM QPaintDevice__PaintDeviceMetric = 3
const QPaintDevice__PdmHeightMM QPaintDevice__PaintDeviceMetric = 4
const QPaintDevice__PdmNumColors QPaintDevice__PaintDeviceMetric = 5
const QPaintDevice__PdmDepth QPaintDevice__PaintDeviceMetric = 6
const QPaintDevice__PdmDpiX QPaintDevice__PaintDeviceMetric = 7
const QPaintDevice__PdmDpiY QPaintDevice__PaintDeviceMetric = 8
const QPaintDevice__PdmPhysicalDpiX QPaintDevice__PaintDeviceMetric = 9
const QPaintDevice__PdmPhysicalDpiY QPaintDevice__PaintDeviceMetric = 10
const QPaintDevice__PdmDevicePixelRatio QPaintDevice__PaintDeviceMetric = 11
const QPaintDevice__PdmDevicePixelRatioScaled QPaintDevice__PaintDeviceMetric = 12

type QPaintEngine__PaintEngineFeature = int

const QPaintEngine__PrimitiveTransform QPaintEngine__PaintEngineFeature = 1
const QPaintEngine__PatternTransform QPaintEngine__PaintEngineFeature = 2
const QPaintEngine__PixmapTransform QPaintEngine__PaintEngineFeature = 4
const QPaintEngine__PatternBrush QPaintEngine__PaintEngineFeature = 8
const QPaintEngine__LinearGradientFill QPaintEngine__PaintEngineFeature = 16
const QPaintEngine__RadialGradientFill QPaintEngine__PaintEngineFeature = 32
const QPaintEngine__ConicalGradientFill QPaintEngine__PaintEngineFeature = 64
const QPaintEngine__AlphaBlend QPaintEngine__PaintEngineFeature = 128
const QPaintEngine__PorterDuff QPaintEngine__PaintEngineFeature = 256
const QPaintEngine__PainterPaths QPaintEngine__PaintEngineFeature = 512
const QPaintEngine__Antialiasing QPaintEngine__PaintEngineFeature = 1024
const QPaintEngine__BrushStroke QPaintEngine__PaintEngineFeature = 2048
const QPaintEngine__ConstantOpacity QPaintEngine__PaintEngineFeature = 4096
const QPaintEngine__MaskedBrush QPaintEngine__PaintEngineFeature = 8192
const QPaintEngine__PerspectiveTransform QPaintEngine__PaintEngineFeature = 16384
const QPaintEngine__BlendModes QPaintEngine__PaintEngineFeature = 32768
const QPaintEngine__ObjectBoundingModeGradients QPaintEngine__PaintEngineFeature = 65536
const QPaintEngine__RasterOpModes QPaintEngine__PaintEngineFeature = 131072
const QPaintEngine__PaintOutsidePaintEvent QPaintEngine__PaintEngineFeature = 536870912
const QPaintEngine__AllFeatures QPaintEngine__PaintEngineFeature = -1

type QPaintEngine__DirtyFlag = int

const QPaintEngine__DirtyPen QPaintEngine__DirtyFlag = 1
const QPaintEngine__DirtyBrush QPaintEngine__DirtyFlag = 2
const QPaintEngine__DirtyBrushOrigin QPaintEngine__DirtyFlag = 4
const QPaintEngine__DirtyFont QPaintEngine__DirtyFlag = 8
const QPaintEngine__DirtyBackground QPaintEngine__DirtyFlag = 16
const QPaintEngine__DirtyBackgroundMode QPaintEngine__DirtyFlag = 32
const QPaintEngine__DirtyTransform QPaintEngine__DirtyFlag = 64
const QPaintEngine__DirtyClipRegion QPaintEngine__DirtyFlag = 128
const QPaintEngine__DirtyClipPath QPaintEngine__DirtyFlag = 256
const QPaintEngine__DirtyHints QPaintEngine__DirtyFlag = 512
const QPaintEngine__DirtyCompositionMode QPaintEngine__DirtyFlag = 1024
const QPaintEngine__DirtyClipEnabled QPaintEngine__DirtyFlag = 2048
const QPaintEngine__DirtyOpacity QPaintEngine__DirtyFlag = 4096
const QPaintEngine__AllDirty QPaintEngine__DirtyFlag = 65535

type QPaintEngine__PolygonDrawMode = int

const QPaintEngine__OddEvenMode QPaintEngine__PolygonDrawMode = 0
const QPaintEngine__WindingMode QPaintEngine__PolygonDrawMode = 1
const QPaintEngine__ConvexMode QPaintEngine__PolygonDrawMode = 2
const QPaintEngine__PolylineMode QPaintEngine__PolygonDrawMode = 3

type QPaintEngine__Type = int

const QPaintEngine__X11 QPaintEngine__Type = 0
const QPaintEngine__Windows QPaintEngine__Type = 1
const QPaintEngine__QuickDraw QPaintEngine__Type = 2
const QPaintEngine__CoreGraphics QPaintEngine__Type = 3
const QPaintEngine__MacPrinter QPaintEngine__Type = 4
const QPaintEngine__QWindowSystem QPaintEngine__Type = 5
const QPaintEngine__PostScript QPaintEngine__Type = 6
const QPaintEngine__OpenGL QPaintEngine__Type = 7
const QPaintEngine__Picture QPaintEngine__Type = 8
const QPaintEngine__SVG QPaintEngine__Type = 9
const QPaintEngine__Raster QPaintEngine__Type = 10
const QPaintEngine__Direct3D QPaintEngine__Type = 11
const QPaintEngine__Pdf QPaintEngine__Type = 12
const QPaintEngine__OpenVG QPaintEngine__Type = 13
const QPaintEngine__OpenGL2 QPaintEngine__Type = 14
const QPaintEngine__PaintBuffer QPaintEngine__Type = 15
const QPaintEngine__Blitter QPaintEngine__Type = 16
const QPaintEngine__Direct2D QPaintEngine__Type = 17
const QPaintEngine__User QPaintEngine__Type = 50
const QPaintEngine__MaxUser QPaintEngine__Type = 100

type QPainter__RenderHint = int

const QPainter__Antialiasing QPainter__RenderHint = 1
const QPainter__TextAntialiasing QPainter__RenderHint = 2
const QPainter__SmoothPixmapTransform QPainter__RenderHint = 4
const QPainter__HighQualityAntialiasing QPainter__RenderHint = 8
const QPainter__NonCosmeticDefaultPen QPainter__RenderHint = 16
const QPainter__Qt4CompatiblePainting QPainter__RenderHint = 32

type QPainter__PixmapFragmentHint = int

const QPainter__OpaqueHint QPainter__PixmapFragmentHint = 1

type QPainter__CompositionMode = int

const QPainter__CompositionMode_SourceOver QPainter__CompositionMode = 0
const QPainter__CompositionMode_DestinationOver QPainter__CompositionMode = 1
const QPainter__CompositionMode_Clear QPainter__CompositionMode = 2
const QPainter__CompositionMode_Source QPainter__CompositionMode = 3
const QPainter__CompositionMode_Destination QPainter__CompositionMode = 4
const QPainter__CompositionMode_SourceIn QPainter__CompositionMode = 5
const QPainter__CompositionMode_DestinationIn QPainter__CompositionMode = 6
const QPainter__CompositionMode_SourceOut QPainter__CompositionMode = 7
const QPainter__CompositionMode_DestinationOut QPainter__CompositionMode = 8
const QPainter__CompositionMode_SourceAtop QPainter__CompositionMode = 9
const QPainter__CompositionMode_DestinationAtop QPainter__CompositionMode = 10
const QPainter__CompositionMode_Xor QPainter__CompositionMode = 11
const QPainter__CompositionMode_Plus QPainter__CompositionMode = 12
const QPainter__CompositionMode_Multiply QPainter__CompositionMode = 13
const QPainter__CompositionMode_Screen QPainter__CompositionMode = 14
const QPainter__CompositionMode_Overlay QPainter__CompositionMode = 15
const QPainter__CompositionMode_Darken QPainter__CompositionMode = 16
const QPainter__CompositionMode_Lighten QPainter__CompositionMode = 17
const QPainter__CompositionMode_ColorDodge QPainter__CompositionMode = 18
const QPainter__CompositionMode_ColorBurn QPainter__CompositionMode = 19
const QPainter__CompositionMode_HardLight QPainter__CompositionMode = 20
const QPainter__CompositionMode_SoftLight QPainter__CompositionMode = 21
const QPainter__CompositionMode_Difference QPainter__CompositionMode = 22
const QPainter__CompositionMode_Exclusion QPainter__CompositionMode = 23
const QPainter__RasterOp_SourceOrDestination QPainter__CompositionMode = 24
const QPainter__RasterOp_SourceAndDestination QPainter__CompositionMode = 25
const QPainter__RasterOp_SourceXorDestination QPainter__CompositionMode = 26
const QPainter__RasterOp_NotSourceAndNotDestination QPainter__CompositionMode = 27
const QPainter__RasterOp_NotSourceOrNotDestination QPainter__CompositionMode = 28
const QPainter__RasterOp_NotSourceXorDestination QPainter__CompositionMode = 29
const QPainter__RasterOp_NotSource QPainter__CompositionMode = 30
const QPainter__RasterOp_NotSourceAndDestination QPainter__CompositionMode = 31
const QPainter__RasterOp_SourceAndNotDestination QPainter__CompositionMode = 32
const QPainter__RasterOp_NotSourceOrDestination QPainter__CompositionMode = 33
const QPainter__RasterOp_SourceOrNotDestination QPainter__CompositionMode = 34
const QPainter__RasterOp_ClearDestination QPainter__CompositionMode = 35
const QPainter__RasterOp_SetDestination QPainter__CompositionMode = 36
const QPainter__RasterOp_NotDestination QPainter__CompositionMode = 37

type QPainterPath__ElementType = int

const QPainterPath__MoveToElement QPainterPath__ElementType = 0
const QPainterPath__LineToElement QPainterPath__ElementType = 1
const QPainterPath__CurveToElement QPainterPath__ElementType = 2
const QPainterPath__CurveToDataElement QPainterPath__ElementType = 3

type QPalette__ColorGroup = int

const QPalette__Active QPalette__ColorGroup = 0
const QPalette__Disabled QPalette__ColorGroup = 1
const QPalette__Inactive QPalette__ColorGroup = 2
const QPalette__NColorGroups QPalette__ColorGroup = 3
const QPalette__Current QPalette__ColorGroup = 4
const QPalette__All QPalette__ColorGroup = 5
const QPalette__Normal QPalette__ColorGroup = 0

type QPalette__ColorRole = int

const QPalette__WindowText QPalette__ColorRole = 0
const QPalette__Button QPalette__ColorRole = 1
const QPalette__Light QPalette__ColorRole = 2
const QPalette__Midlight QPalette__ColorRole = 3
const QPalette__Dark QPalette__ColorRole = 4
const QPalette__Mid QPalette__ColorRole = 5
const QPalette__Text QPalette__ColorRole = 6
const QPalette__BrightText QPalette__ColorRole = 7
const QPalette__ButtonText QPalette__ColorRole = 8
const QPalette__Base QPalette__ColorRole = 9
const QPalette__Window QPalette__ColorRole = 10
const QPalette__Shadow QPalette__ColorRole = 11
const QPalette__Highlight QPalette__ColorRole = 12
const QPalette__HighlightedText QPalette__ColorRole = 13
const QPalette__Link QPalette__ColorRole = 14
const QPalette__LinkVisited QPalette__ColorRole = 15
const QPalette__AlternateBase QPalette__ColorRole = 16
const QPalette__NoRole QPalette__ColorRole = 17
const QPalette__ToolTipBase QPalette__ColorRole = 18
const QPalette__ToolTipText QPalette__ColorRole = 19
const QPalette__PlaceholderText QPalette__ColorRole = 20
const QPalette__NColorRoles QPalette__ColorRole = 21
const QPalette__Foreground QPalette__ColorRole = 0
const QPalette__Background QPalette__ColorRole = 10

type QPixelFormat__FieldWidth = int

const QPixelFormat__ModelFieldWidth QPixelFormat__FieldWidth = 4
const QPixelFormat__FirstFieldWidth QPixelFormat__FieldWidth = 6
const QPixelFormat__SecondFieldWidth QPixelFormat__FieldWidth = 6
const QPixelFormat__ThirdFieldWidth QPixelFormat__FieldWidth = 6
const QPixelFormat__FourthFieldWidth QPixelFormat__FieldWidth = 6
const QPixelFormat__FifthFieldWidth QPixelFormat__FieldWidth = 6
const QPixelFormat__AlphaFieldWidth QPixelFormat__FieldWidth = 6
const QPixelFormat__AlphaUsageFieldWidth QPixelFormat__FieldWidth = 1
const QPixelFormat__AlphaPositionFieldWidth QPixelFormat__FieldWidth = 1
const QPixelFormat__PremulFieldWidth QPixelFormat__FieldWidth = 1
const QPixelFormat__TypeInterpretationFieldWidth QPixelFormat__FieldWidth = 4
const QPixelFormat__ByteOrderFieldWidth QPixelFormat__FieldWidth = 2
const QPixelFormat__SubEnumFieldWidth QPixelFormat__FieldWidth = 6
const QPixelFormat__UnusedFieldWidth QPixelFormat__FieldWidth = 9
const QPixelFormat__TotalFieldWidthByWidths QPixelFormat__FieldWidth = 64

type QPixelFormat__Field = int

const QPixelFormat__ModelField QPixelFormat__Field = 0
const QPixelFormat__FirstField QPixelFormat__Field = 4
const QPixelFormat__SecondField QPixelFormat__Field = 10
const QPixelFormat__ThirdField QPixelFormat__Field = 16
const QPixelFormat__FourthField QPixelFormat__Field = 22
const QPixelFormat__FifthField QPixelFormat__Field = 28
const QPixelFormat__AlphaField QPixelFormat__Field = 34
const QPixelFormat__AlphaUsageField QPixelFormat__Field = 40
const QPixelFormat__AlphaPositionField QPixelFormat__Field = 41
const QPixelFormat__PremulField QPixelFormat__Field = 42
const QPixelFormat__TypeInterpretationField QPixelFormat__Field = 43
const QPixelFormat__ByteOrderField QPixelFormat__Field = 47
const QPixelFormat__SubEnumField QPixelFormat__Field = 49
const QPixelFormat__UnusedField QPixelFormat__Field = 55
const QPixelFormat__TotalFieldWidthByOffsets QPixelFormat__Field = 64

type QPixelFormat__ColorModel = int

const QPixelFormat__RGB QPixelFormat__ColorModel = 0
const QPixelFormat__BGR QPixelFormat__ColorModel = 1
const QPixelFormat__Indexed QPixelFormat__ColorModel = 2
const QPixelFormat__Grayscale QPixelFormat__ColorModel = 3
const QPixelFormat__CMYK QPixelFormat__ColorModel = 4
const QPixelFormat__HSL QPixelFormat__ColorModel = 5
const QPixelFormat__HSV QPixelFormat__ColorModel = 6
const QPixelFormat__YUV QPixelFormat__ColorModel = 7
const QPixelFormat__Alpha QPixelFormat__ColorModel = 8

type QPixelFormat__AlphaUsage = int

const QPixelFormat__UsesAlpha QPixelFormat__AlphaUsage = 0
const QPixelFormat__IgnoresAlpha QPixelFormat__AlphaUsage = 1

type QPixelFormat__AlphaPosition = int

const QPixelFormat__AtBeginning QPixelFormat__AlphaPosition = 0
const QPixelFormat__AtEnd QPixelFormat__AlphaPosition = 1

type QPixelFormat__AlphaPremultiplied = int

const QPixelFormat__NotPremultiplied QPixelFormat__AlphaPremultiplied = 0
const QPixelFormat__Premultiplied QPixelFormat__AlphaPremultiplied = 1

type QPixelFormat__TypeInterpretation = int

const QPixelFormat__UnsignedInteger QPixelFormat__TypeInterpretation = 0
const QPixelFormat__UnsignedShort QPixelFormat__TypeInterpretation = 1
const QPixelFormat__UnsignedByte QPixelFormat__TypeInterpretation = 2
const QPixelFormat__FloatingPoint QPixelFormat__TypeInterpretation = 3

type QPixelFormat__YUVLayout = int

const QPixelFormat__YUV444 QPixelFormat__YUVLayout = 0
const QPixelFormat__YUV422 QPixelFormat__YUVLayout = 1
const QPixelFormat__YUV411 QPixelFormat__YUVLayout = 2
const QPixelFormat__YUV420P QPixelFormat__YUVLayout = 3
const QPixelFormat__YUV420SP QPixelFormat__YUVLayout = 4
const QPixelFormat__YV12 QPixelFormat__YUVLayout = 5
const QPixelFormat__UYVY QPixelFormat__YUVLayout = 6
const QPixelFormat__YUYV QPixelFormat__YUVLayout = 7
const QPixelFormat__NV12 QPixelFormat__YUVLayout = 8
const QPixelFormat__NV21 QPixelFormat__YUVLayout = 9
const QPixelFormat__IMC1 QPixelFormat__YUVLayout = 10
const QPixelFormat__IMC2 QPixelFormat__YUVLayout = 11
const QPixelFormat__IMC3 QPixelFormat__YUVLayout = 12
const QPixelFormat__IMC4 QPixelFormat__YUVLayout = 13
const QPixelFormat__Y8 QPixelFormat__YUVLayout = 14
const QPixelFormat__Y16 QPixelFormat__YUVLayout = 15

type QPixelFormat__ByteOrder = int

const QPixelFormat__LittleEndian QPixelFormat__ByteOrder = 0
const QPixelFormat__BigEndian QPixelFormat__ByteOrder = 1
const QPixelFormat__CurrentSystemEndian QPixelFormat__ByteOrder = 2

type QPlatformSurfaceEvent__SurfaceEventType = int

const QPlatformSurfaceEvent__SurfaceCreated QPlatformSurfaceEvent__SurfaceEventType = 0
const QPlatformSurfaceEvent__SurfaceAboutToBeDestroyed QPlatformSurfaceEvent__SurfaceEventType = 1

type QRawFont__AntialiasingType = int

const QRawFont__PixelAntialiasing QRawFont__AntialiasingType = 0
const QRawFont__SubPixelAntialiasing QRawFont__AntialiasingType = 1

type QRawFont__LayoutFlag = int

const QRawFont__SeparateAdvances QRawFont__LayoutFlag = 0
const QRawFont__KernedAdvances QRawFont__LayoutFlag = 1
const QRawFont__UseDesignMetrics QRawFont__LayoutFlag = 2

type QRegion__RegionType = int

const QRegion__Rectangle QRegion__RegionType = 0
const QRegion__Ellipse QRegion__RegionType = 1

type QRgba64__Shifts = int

const QRgba64__RedShift QRgba64__Shifts = 0
const QRgba64__GreenShift QRgba64__Shifts = 16
const QRgba64__BlueShift QRgba64__Shifts = 32
const QRgba64__AlphaShift QRgba64__Shifts = 48

type QScrollEvent__ScrollState = int

const QScrollEvent__ScrollStarted QScrollEvent__ScrollState = 0
const QScrollEvent__ScrollUpdated QScrollEvent__ScrollState = 1
const QScrollEvent__ScrollFinished QScrollEvent__ScrollState = 2

type QSessionManager__RestartHint = int

const QSessionManager__RestartIfRunning QSessionManager__RestartHint = 0
const QSessionManager__RestartAnyway QSessionManager__RestartHint = 1
const QSessionManager__RestartImmediately QSessionManager__RestartHint = 2
const QSessionManager__RestartNever QSessionManager__RestartHint = 3

type QStandardItem__ItemType = int

const QStandardItem__Type QStandardItem__ItemType = 0
const QStandardItem__UserType QStandardItem__ItemType = 1000

type QStaticText__PerformanceHint = int

const QStaticText__ModerateCaching QStaticText__PerformanceHint = 0
const QStaticText__AggressiveCaching QStaticText__PerformanceHint = 1

type QSurface__SurfaceClass = int

const QSurface__Window QSurface__SurfaceClass = 0
const QSurface__Offscreen QSurface__SurfaceClass = 1

type QSurface__SurfaceType = int

const QSurface__RasterSurface QSurface__SurfaceType = 0
const QSurface__OpenGLSurface QSurface__SurfaceType = 1
const QSurface__RasterGLSurface QSurface__SurfaceType = 2
const QSurface__OpenVGSurface QSurface__SurfaceType = 3
const QSurface__VulkanSurface QSurface__SurfaceType = 4
const QSurface__MetalSurface QSurface__SurfaceType = 5

type QSurfaceFormat__FormatOption = int

const QSurfaceFormat__StereoBuffers QSurfaceFormat__FormatOption = 1
const QSurfaceFormat__DebugContext QSurfaceFormat__FormatOption = 2
const QSurfaceFormat__DeprecatedFunctions QSurfaceFormat__FormatOption = 4
const QSurfaceFormat__ResetNotification QSurfaceFormat__FormatOption = 8

type QSurfaceFormat__SwapBehavior = int

const QSurfaceFormat__DefaultSwapBehavior QSurfaceFormat__SwapBehavior = 0
const QSurfaceFormat__SingleBuffer QSurfaceFormat__SwapBehavior = 1
const QSurfaceFormat__DoubleBuffer QSurfaceFormat__SwapBehavior = 2
const QSurfaceFormat__TripleBuffer QSurfaceFormat__SwapBehavior = 3

type QSurfaceFormat__RenderableType = int

const QSurfaceFormat__DefaultRenderableType QSurfaceFormat__RenderableType = 0
const QSurfaceFormat__OpenGL QSurfaceFormat__RenderableType = 1
const QSurfaceFormat__OpenGLES QSurfaceFormat__RenderableType = 2
const QSurfaceFormat__OpenVG QSurfaceFormat__RenderableType = 4

type QSurfaceFormat__OpenGLContextProfile = int

const QSurfaceFormat__NoProfile QSurfaceFormat__OpenGLContextProfile = 0
const QSurfaceFormat__CoreProfile QSurfaceFormat__OpenGLContextProfile = 1
const QSurfaceFormat__CompatibilityProfile QSurfaceFormat__OpenGLContextProfile = 2

type QSurfaceFormat__ColorSpace = int

const QSurfaceFormat__DefaultColorSpace QSurfaceFormat__ColorSpace = 0
const QSurfaceFormat__sRGBColorSpace QSurfaceFormat__ColorSpace = 1

type QTabletEvent__TabletDevice = int

const QTabletEvent__NoDevice QTabletEvent__TabletDevice = 0
const QTabletEvent__Puck QTabletEvent__TabletDevice = 1
const QTabletEvent__Stylus QTabletEvent__TabletDevice = 2
const QTabletEvent__Airbrush QTabletEvent__TabletDevice = 3
const QTabletEvent__FourDMouse QTabletEvent__TabletDevice = 4
const QTabletEvent__XFreeEraser QTabletEvent__TabletDevice = 5
const QTabletEvent__RotationStylus QTabletEvent__TabletDevice = 6

type QTabletEvent__PointerType = int

const QTabletEvent__UnknownPointer QTabletEvent__PointerType = 0
const QTabletEvent__Pen QTabletEvent__PointerType = 1
const QTabletEvent__Cursor QTabletEvent__PointerType = 2
const QTabletEvent__Eraser QTabletEvent__PointerType = 3

type QTextBlockFormat__LineHeightTypes = int

const QTextBlockFormat__SingleHeight QTextBlockFormat__LineHeightTypes = 0
const QTextBlockFormat__ProportionalHeight QTextBlockFormat__LineHeightTypes = 1
const QTextBlockFormat__FixedHeight QTextBlockFormat__LineHeightTypes = 2
const QTextBlockFormat__MinimumHeight QTextBlockFormat__LineHeightTypes = 3
const QTextBlockFormat__LineDistanceHeight QTextBlockFormat__LineHeightTypes = 4

type QTextCharFormat__VerticalAlignment = int

const QTextCharFormat__AlignNormal QTextCharFormat__VerticalAlignment = 0
const QTextCharFormat__AlignSuperScript QTextCharFormat__VerticalAlignment = 1
const QTextCharFormat__AlignSubScript QTextCharFormat__VerticalAlignment = 2
const QTextCharFormat__AlignMiddle QTextCharFormat__VerticalAlignment = 3
const QTextCharFormat__AlignTop QTextCharFormat__VerticalAlignment = 4
const QTextCharFormat__AlignBottom QTextCharFormat__VerticalAlignment = 5
const QTextCharFormat__AlignBaseline QTextCharFormat__VerticalAlignment = 6

type QTextCharFormat__UnderlineStyle = int

const QTextCharFormat__NoUnderline QTextCharFormat__UnderlineStyle = 0
const QTextCharFormat__SingleUnderline QTextCharFormat__UnderlineStyle = 1
const QTextCharFormat__DashUnderline QTextCharFormat__UnderlineStyle = 2
const QTextCharFormat__DotLine QTextCharFormat__UnderlineStyle = 3
const QTextCharFormat__DashDotLine QTextCharFormat__UnderlineStyle = 4
const QTextCharFormat__DashDotDotLine QTextCharFormat__UnderlineStyle = 5
const QTextCharFormat__WaveUnderline QTextCharFormat__UnderlineStyle = 6
const QTextCharFormat__SpellCheckUnderline QTextCharFormat__UnderlineStyle = 7

type QTextCharFormat__FontPropertiesInheritanceBehavior = int

const QTextCharFormat__FontPropertiesSpecifiedOnly QTextCharFormat__FontPropertiesInheritanceBehavior = 0
const QTextCharFormat__FontPropertiesAll QTextCharFormat__FontPropertiesInheritanceBehavior = 1

type QTextCursor__MoveMode = int

const QTextCursor__MoveAnchor QTextCursor__MoveMode = 0
const QTextCursor__KeepAnchor QTextCursor__MoveMode = 1

type QTextCursor__MoveOperation = int

const QTextCursor__NoMove QTextCursor__MoveOperation = 0
const QTextCursor__Start QTextCursor__MoveOperation = 1
const QTextCursor__Up QTextCursor__MoveOperation = 2
const QTextCursor__StartOfLine QTextCursor__MoveOperation = 3
const QTextCursor__StartOfBlock QTextCursor__MoveOperation = 4
const QTextCursor__StartOfWord QTextCursor__MoveOperation = 5
const QTextCursor__PreviousBlock QTextCursor__MoveOperation = 6
const QTextCursor__PreviousCharacter QTextCursor__MoveOperation = 7
const QTextCursor__PreviousWord QTextCursor__MoveOperation = 8
const QTextCursor__Left QTextCursor__MoveOperation = 9
const QTextCursor__WordLeft QTextCursor__MoveOperation = 10
const QTextCursor__End QTextCursor__MoveOperation = 11
const QTextCursor__Down QTextCursor__MoveOperation = 12
const QTextCursor__EndOfLine QTextCursor__MoveOperation = 13
const QTextCursor__EndOfWord QTextCursor__MoveOperation = 14
const QTextCursor__EndOfBlock QTextCursor__MoveOperation = 15
const QTextCursor__NextBlock QTextCursor__MoveOperation = 16
const QTextCursor__NextCharacter QTextCursor__MoveOperation = 17
const QTextCursor__NextWord QTextCursor__MoveOperation = 18
const QTextCursor__Right QTextCursor__MoveOperation = 19
const QTextCursor__WordRight QTextCursor__MoveOperation = 20
const QTextCursor__NextCell QTextCursor__MoveOperation = 21
const QTextCursor__PreviousCell QTextCursor__MoveOperation = 22
const QTextCursor__NextRow QTextCursor__MoveOperation = 23
const QTextCursor__PreviousRow QTextCursor__MoveOperation = 24

type QTextCursor__SelectionType = int

const QTextCursor__WordUnderCursor QTextCursor__SelectionType = 0
const QTextCursor__LineUnderCursor QTextCursor__SelectionType = 1
const QTextCursor__BlockUnderCursor QTextCursor__SelectionType = 2
const QTextCursor__Document QTextCursor__SelectionType = 3

type QTextDocument__MetaInformation = int

const QTextDocument__DocumentTitle QTextDocument__MetaInformation = 0
const QTextDocument__DocumentUrl QTextDocument__MetaInformation = 1

type QTextDocument__FindFlag = int

const QTextDocument__FindBackward QTextDocument__FindFlag = 1
const QTextDocument__FindCaseSensitively QTextDocument__FindFlag = 2
const QTextDocument__FindWholeWords QTextDocument__FindFlag = 4

type QTextDocument__ResourceType = int

const QTextDocument__HtmlResource QTextDocument__ResourceType = 1
const QTextDocument__ImageResource QTextDocument__ResourceType = 2
const QTextDocument__StyleSheetResource QTextDocument__ResourceType = 3
const QTextDocument__UserResource QTextDocument__ResourceType = 100

type QTextDocument__Stacks = int

const QTextDocument__UndoStack QTextDocument__Stacks = 1
const QTextDocument__RedoStack QTextDocument__Stacks = 2
const QTextDocument__UndoAndRedoStacks QTextDocument__Stacks = 3

type QTextFormat__FormatType = int

const QTextFormat__InvalidFormat QTextFormat__FormatType = -1
const QTextFormat__BlockFormat QTextFormat__FormatType = 1
const QTextFormat__CharFormat QTextFormat__FormatType = 2
const QTextFormat__ListFormat QTextFormat__FormatType = 3
const QTextFormat__TableFormat QTextFormat__FormatType = 4
const QTextFormat__FrameFormat QTextFormat__FormatType = 5
const QTextFormat__UserFormat QTextFormat__FormatType = 100

type QTextFormat__Property = int

const QTextFormat__ObjectIndex QTextFormat__Property = 0
const QTextFormat__CssFloat QTextFormat__Property = 2048
const QTextFormat__LayoutDirection QTextFormat__Property = 2049
const QTextFormat__OutlinePen QTextFormat__Property = 2064
const QTextFormat__BackgroundBrush QTextFormat__Property = 2080
const QTextFormat__ForegroundBrush QTextFormat__Property = 2081
const QTextFormat__BackgroundImageUrl QTextFormat__Property = 2083
const QTextFormat__BlockAlignment QTextFormat__Property = 4112
const QTextFormat__BlockTopMargin QTextFormat__Property = 4144
const QTextFormat__BlockBottomMargin QTextFormat__Property = 4145
const QTextFormat__BlockLeftMargin QTextFormat__Property = 4146
const QTextFormat__BlockRightMargin QTextFormat__Property = 4147
const QTextFormat__TextIndent QTextFormat__Property = 4148
const QTextFormat__TabPositions QTextFormat__Property = 4149
const QTextFormat__BlockIndent QTextFormat__Property = 4160
const QTextFormat__LineHeight QTextFormat__Property = 4168
const QTextFormat__LineHeightType QTextFormat__Property = 4169
const QTextFormat__BlockNonBreakableLines QTextFormat__Property = 4176
const QTextFormat__BlockTrailingHorizontalRulerWidth QTextFormat__Property = 4192
const QTextFormat__HeadingLevel QTextFormat__Property = 4208
const QTextFormat__FirstFontProperty QTextFormat__Property = 8160
const QTextFormat__FontCapitalization QTextFormat__Property = 8160
const QTextFormat__FontLetterSpacingType QTextFormat__Property = 8243
const QTextFormat__FontLetterSpacing QTextFormat__Property = 8161
const QTextFormat__FontWordSpacing QTextFormat__Property = 8162
const QTextFormat__FontStretch QTextFormat__Property = 8244
const QTextFormat__FontStyleHint QTextFormat__Property = 8163
const QTextFormat__FontStyleStrategy QTextFormat__Property = 8164
const QTextFormat__FontKerning QTextFormat__Property = 8165
const QTextFormat__FontHintingPreference QTextFormat__Property = 8166
const QTextFormat__FontFamily QTextFormat__Property = 8192
const QTextFormat__FontPointSize QTextFormat__Property = 8193
const QTextFormat__FontSizeAdjustment QTextFormat__Property = 8194
const QTextFormat__FontSizeIncrement QTextFormat__Property = 8194
const QTextFormat__FontWeight QTextFormat__Property = 8195
const QTextFormat__FontItalic QTextFormat__Property = 8196
const QTextFormat__FontUnderline QTextFormat__Property = 8197
const QTextFormat__FontOverline QTextFormat__Property = 8198
const QTextFormat__FontStrikeOut QTextFormat__Property = 8199
const QTextFormat__FontFixedPitch QTextFormat__Property = 8200
const QTextFormat__FontPixelSize QTextFormat__Property = 8201
const QTextFormat__LastFontProperty QTextFormat__Property = 8201
const QTextFormat__TextUnderlineColor QTextFormat__Property = 8208
const QTextFormat__TextVerticalAlignment QTextFormat__Property = 8225
const QTextFormat__TextOutline QTextFormat__Property = 8226
const QTextFormat__TextUnderlineStyle QTextFormat__Property = 8227
const QTextFormat__TextToolTip QTextFormat__Property = 8228
const QTextFormat__IsAnchor QTextFormat__Property = 8240
const QTextFormat__AnchorHref QTextFormat__Property = 8241
const QTextFormat__AnchorName QTextFormat__Property = 8242
const QTextFormat__ObjectType QTextFormat__Property = 12032
const QTextFormat__ListStyle QTextFormat__Property = 12288
const QTextFormat__ListIndent QTextFormat__Property = 12289
const QTextFormat__ListNumberPrefix QTextFormat__Property = 12290
const QTextFormat__ListNumberSuffix QTextFormat__Property = 12291
const QTextFormat__FrameBorder QTextFormat__Property = 16384
const QTextFormat__FrameMargin QTextFormat__Property = 16385
const QTextFormat__FramePadding QTextFormat__Property = 16386
const QTextFormat__FrameWidth QTextFormat__Property = 16387
const QTextFormat__FrameHeight QTextFormat__Property = 16388
const QTextFormat__FrameTopMargin QTextFormat__Property = 16389
const QTextFormat__FrameBottomMargin QTextFormat__Property = 16390
const QTextFormat__FrameLeftMargin QTextFormat__Property = 16391
const QTextFormat__FrameRightMargin QTextFormat__Property = 16392
const QTextFormat__FrameBorderBrush QTextFormat__Property = 16393
const QTextFormat__FrameBorderStyle QTextFormat__Property = 16400
const QTextFormat__TableColumns QTextFormat__Property = 16640
const QTextFormat__TableColumnWidthConstraints QTextFormat__Property = 16641
const QTextFormat__TableCellSpacing QTextFormat__Property = 16642
const QTextFormat__TableCellPadding QTextFormat__Property = 16643
const QTextFormat__TableHeaderRowCount QTextFormat__Property = 16644
const QTextFormat__TableCellRowSpan QTextFormat__Property = 18448
const QTextFormat__TableCellColumnSpan QTextFormat__Property = 18449
const QTextFormat__TableCellTopPadding QTextFormat__Property = 18450
const QTextFormat__TableCellBottomPadding QTextFormat__Property = 18451
const QTextFormat__TableCellLeftPadding QTextFormat__Property = 18452
const QTextFormat__TableCellRightPadding QTextFormat__Property = 18453
const QTextFormat__ImageName QTextFormat__Property = 20480
const QTextFormat__ImageWidth QTextFormat__Property = 20496
const QTextFormat__ImageHeight QTextFormat__Property = 20497
const QTextFormat__ImageQuality QTextFormat__Property = 20500
const QTextFormat__FullWidthSelection QTextFormat__Property = 24576
const QTextFormat__PageBreakPolicy QTextFormat__Property = 28672
const QTextFormat__UserProperty QTextFormat__Property = 1048576

type QTextFormat__ObjectTypes = int

const QTextFormat__NoObject QTextFormat__ObjectTypes = 0
const QTextFormat__ImageObject QTextFormat__ObjectTypes = 1
const QTextFormat__TableObject QTextFormat__ObjectTypes = 2
const QTextFormat__TableCellObject QTextFormat__ObjectTypes = 3
const QTextFormat__UserObject QTextFormat__ObjectTypes = 4096

type QTextFormat__PageBreakFlag = int

const QTextFormat__PageBreak_Auto QTextFormat__PageBreakFlag = 0
const QTextFormat__PageBreak_AlwaysBefore QTextFormat__PageBreakFlag = 1
const QTextFormat__PageBreak_AlwaysAfter QTextFormat__PageBreakFlag = 16

type QTextFrameFormat__Position = int

const QTextFrameFormat__InFlow QTextFrameFormat__Position = 0
const QTextFrameFormat__FloatLeft QTextFrameFormat__Position = 1
const QTextFrameFormat__FloatRight QTextFrameFormat__Position = 2

type QTextFrameFormat__BorderStyle = int

const QTextFrameFormat__BorderStyle_None QTextFrameFormat__BorderStyle = 0
const QTextFrameFormat__BorderStyle_Dotted QTextFrameFormat__BorderStyle = 1
const QTextFrameFormat__BorderStyle_Dashed QTextFrameFormat__BorderStyle = 2
const QTextFrameFormat__BorderStyle_Solid QTextFrameFormat__BorderStyle = 3
const QTextFrameFormat__BorderStyle_Double QTextFrameFormat__BorderStyle = 4
const QTextFrameFormat__BorderStyle_DotDash QTextFrameFormat__BorderStyle = 5
const QTextFrameFormat__BorderStyle_DotDotDash QTextFrameFormat__BorderStyle = 6
const QTextFrameFormat__BorderStyle_Groove QTextFrameFormat__BorderStyle = 7
const QTextFrameFormat__BorderStyle_Ridge QTextFrameFormat__BorderStyle = 8
const QTextFrameFormat__BorderStyle_Inset QTextFrameFormat__BorderStyle = 9
const QTextFrameFormat__BorderStyle_Outset QTextFrameFormat__BorderStyle = 10

type QTextItem__RenderFlag = int

const QTextItem__RightToLeft QTextItem__RenderFlag = 1
const QTextItem__Overline QTextItem__RenderFlag = 16
const QTextItem__Underline QTextItem__RenderFlag = 32
const QTextItem__StrikeOut QTextItem__RenderFlag = 64
const QTextItem__Dummy QTextItem__RenderFlag = -1

type QTextLayout__CursorMode = int

const QTextLayout__SkipCharacters QTextLayout__CursorMode = 0
const QTextLayout__SkipWords QTextLayout__CursorMode = 1

type QTextLength__Type = int

const QTextLength__VariableLength QTextLength__Type = 0
const QTextLength__FixedLength QTextLength__Type = 1
const QTextLength__PercentageLength QTextLength__Type = 2

type QTextLine__Edge = int

const QTextLine__Leading QTextLine__Edge = 0
const QTextLine__Trailing QTextLine__Edge = 1

type QTextLine__CursorPosition = int

const QTextLine__CursorBetweenCharacters QTextLine__CursorPosition = 0
const QTextLine__CursorOnCharacter QTextLine__CursorPosition = 1

type QTextListFormat__Style = int

const QTextListFormat__ListDisc QTextListFormat__Style = -1
const QTextListFormat__ListCircle QTextListFormat__Style = -2
const QTextListFormat__ListSquare QTextListFormat__Style = -3
const QTextListFormat__ListDecimal QTextListFormat__Style = -4
const QTextListFormat__ListLowerAlpha QTextListFormat__Style = -5
const QTextListFormat__ListUpperAlpha QTextListFormat__Style = -6
const QTextListFormat__ListLowerRoman QTextListFormat__Style = -7
const QTextListFormat__ListUpperRoman QTextListFormat__Style = -8
const QTextListFormat__ListStyleUndefined QTextListFormat__Style = 0

type QTextOption__TabType = int

const QTextOption__LeftTab QTextOption__TabType = 0
const QTextOption__RightTab QTextOption__TabType = 1
const QTextOption__CenterTab QTextOption__TabType = 2
const QTextOption__DelimiterTab QTextOption__TabType = 3

type QTextOption__WrapMode = int

const QTextOption__NoWrap QTextOption__WrapMode = 0
const QTextOption__WordWrap QTextOption__WrapMode = 1
const QTextOption__ManualWrap QTextOption__WrapMode = 2
const QTextOption__WrapAnywhere QTextOption__WrapMode = 3
const QTextOption__WrapAtWordBoundaryOrAnywhere QTextOption__WrapMode = 4

type QTextOption__Flag = int

const QTextOption__ShowTabsAndSpaces QTextOption__Flag = 1
const QTextOption__ShowLineAndParagraphSeparators QTextOption__Flag = 2
const QTextOption__AddSpaceForLineAndParagraphSeparators QTextOption__Flag = 4
const QTextOption__SuppressColors QTextOption__Flag = 8
const QTextOption__ShowDocumentTerminator QTextOption__Flag = 16
const QTextOption__IncludeTrailingSpaces QTextOption__Flag = -2147483648

type QTouchDevice__DeviceType = int

const QTouchDevice__TouchScreen QTouchDevice__DeviceType = 0
const QTouchDevice__TouchPad QTouchDevice__DeviceType = 1

type QTouchDevice__CapabilityFlag = int

const QTouchDevice__Position QTouchDevice__CapabilityFlag = 1
const QTouchDevice__Area QTouchDevice__CapabilityFlag = 2
const QTouchDevice__Pressure QTouchDevice__CapabilityFlag = 4
const QTouchDevice__Velocity QTouchDevice__CapabilityFlag = 8
const QTouchDevice__RawPositions QTouchDevice__CapabilityFlag = 16
const QTouchDevice__NormalizedPosition QTouchDevice__CapabilityFlag = 32
const QTouchDevice__MouseEmulation QTouchDevice__CapabilityFlag = 64

type QTransform__TransformationType = int

const QTransform__TxNone QTransform__TransformationType = 0
const QTransform__TxTranslate QTransform__TransformationType = 1
const QTransform__TxScale QTransform__TransformationType = 2
const QTransform__TxRotate QTransform__TransformationType = 4
const QTransform__TxShear QTransform__TransformationType = 8
const QTransform__TxProject QTransform__TransformationType = 16

type QValidator__State = int

const QValidator__Invalid QValidator__State = 0
const QValidator__Intermediate QValidator__State = 1
const QValidator__Acceptable QValidator__State = 2

type QWheelEvent__ = int

const QWheelEvent__DefaultDeltasPerStep QWheelEvent__ = 120

type QWindow__Visibility = int

const QWindow__Hidden QWindow__Visibility = 0
const QWindow__AutomaticVisibility QWindow__Visibility = 1
const QWindow__Windowed QWindow__Visibility = 2
const QWindow__Minimized QWindow__Visibility = 3
const QWindow__Maximized QWindow__Visibility = 4
const QWindow__FullScreen QWindow__Visibility = 5

type QWindow__AncestorMode = int

const QWindow__ExcludeTransients QWindow__AncestorMode = 0
const QWindow__IncludeTransients QWindow__AncestorMode = 1
