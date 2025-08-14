package datasets

type Character struct {
	ID             int    `json:"charaId"`
	CardID         int    `json:"cardId"`
	SupportCardID  int    `json:"supportCardId"`
	CharaStartDate string `json:"charaStartDate"`
	CharaName      string `json:"charaName"`
	CharaCV        string `json:"charaCv"`
	Name           string
	Stats          CharacterStats
	Info           CharacterInfo
}

type CharacterStats struct {
	ID                     int    `json:"id"`
	BirthYear              int    `json:"birthYear"`
	BirthMonth             int    `json:"birthMonth"`
	BirthDay               int    `json:"birthDay"`
	Sex                    int    `json:"sex"`
	ImageColorMain         string `json:"imageColorMain"`
	ImageColorSub          string `json:"imageColorSub"`
	UIColorMain            string `json:"uiColorMain"`
	UIColorSub             string `json:"uiColorSub"`
	UITrainingColor1       string `json:"uiTrainingColor1"`
	UITrainingColor2       string `json:"uiTrainingColor2"`
	UIBorderColor          string `json:"uiBorderColor"`
	UINumColor1            string `json:"uiNumColor1"`
	UINumColor2            string `json:"uiNumColor2"`
	UITurnColor            string `json:"uiTurnColor"`
	UIWipeColor1           string `json:"uiWipeColor1"`
	UIWipeColor2           string `json:"uiWipeColor2"`
	UIWipeColor3           string `json:"uiWipeColor3"`
	UISpeechColor1         string `json:"uiSpeechColor1"`
	UISpeechColor2         string `json:"uiSpeechColor2"`
	UINameplateColor1      string `json:"uiNameplateColor1"`
	UINameplateColor2      string `json:"uiNameplateColor2"`
	Height                 int    `json:"height"`
	Bust                   int    `json:"bust"`
	Scale                  int    `json:"scale"`
	Skin                   int    `json:"skin"`
	Shape                  int    `json:"shape"`
	Socks                  int    `json:"socks"`
	PersonalDress          int    `json:"personalDress"`
	TailModelID            int    `json:"tailModelId"`
	RaceRunningType        int    `json:"raceRunningType"`
	EarRandomTimeMin       int    `json:"earRandomTimeMin"`
	EarRandomTimeMax       int    `json:"earRandomTimeMax"`
	TailRandomTimeMin      int    `json:"tailRandomTimeMin"`
	TailRandomTimeMax      int    `json:"tailRandomTimeMax"`
	StoryEarRandomTimeMin  int    `json:"storyEarRandomTimeMin"`
	StoryEarRandomTimeMax  int    `json:"storyEarRandomTimeMax"`
	StoryTailRandomTimeMin int    `json:"storyTailRandomTimeMin"`
	StoryTailRandomTimeMax int    `json:"storyTailRandomTimeMax"`
	AttachmentModelID      int    `json:"attachmentModelId"`
	MiniMayuShaderType     int    `json:"miniMayuShaderType"`
	StartDate              int64  `json:"startDate"`
	CharaCategory          int    `json:"charaCategory"`
	LoveRankLimit          int    `json:"loveRankLimit"`
	LastYear               int    `json:"lastYear"`
}

type CharacterInfo struct {
	BirthDay        int    `json:"birth_day"`
	BirthMonth      int    `json:"birth_month"`
	CategoryLabel   string `json:"category_label"`
	CategoryLabelEn string `json:"category_label_en"`
	CategoryValue   string `json:"category_value"`
	ColorMain       string `json:"color_main"`
	ColorSub        string `json:"color_sub"`
	DateGMT         string `json:"date_gmt"`
	DetailImgPC     string `json:"detail_img_pc"`
	DetailImgSP     string `json:"detail_img_sp"`
	EarsFact        string `json:"ears_fact"`
	FamilyFact      string `json:"family_fact"`
	GameID          int    `json:"game_id"`
	Grade           string `json:"grade"`
	Height          int    `json:"height"`
	ID              int    `json:"id"`
	Link            string `json:"link"`
	ModifiedGMT     string `json:"modified_gmt"`
	NameEn          string `json:"name_en"`
	NameEnInternal  string `json:"name_en_internal"`
	NameJP          string `json:"name_jp"`
	PreferredURL    string `json:"preferred_url"`
	Profile         string `json:"profile"`
	Residence       string `json:"residence"`
	RowNumber       int    `json:"row_number"`
	ShoeSize        string `json:"shoe_size"`
	SizeB           int    `json:"size_b"`
	SizeH           int    `json:"size_h"`
	SizeW           int    `json:"size_w"`
	Slogan          string `json:"slogan"`
	SnsHeader       string `json:"sns_header"`
	SnsIcon         string `json:"sns_icon"`
	Strengths       string `json:"strengths"`
	TailFact        string `json:"tail_fact"`
	ThumbImg        string `json:"thumb_img"`
	Voice           string `json:"voice"`
	Weaknesses      string `json:"weaknesses"`
	Weight          string `json:"weight"`
}
