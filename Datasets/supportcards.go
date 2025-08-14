package datasets

type SupportCard struct {
	ID               int    `json:"supportCardId"`
	CharaID          int    `json:"charaId"`
	Rarity           int    `json:"rarity"`
	SupportCardTitle string `json:"supportCardTitle"`
	SupportCardName  string `json:"supportCardName"`
	Name             string
	Data             SupportCardData
	Info             SupportCardInfo
}

type SupportCardData struct {
	ID              int   `json:"id"`
	CharaID         int   `json:"charaId"`
	Rarity          int   `json:"rarity"`
	ExchangeItemID  int   `json:"exchangeItemId"`
	EffectTableID   int   `json:"effectTableId"`
	UniqueEffectID  int   `json:"uniqueEffectId"`
	CommandType     int   `json:"commandType"`
	CommandID       int   `json:"commandId"`
	SupportCardType int   `json:"supportCardType"`
	SkillSetID      int   `json:"skillSetId"`
	DetailPosX      int   `json:"detailPosX"`
	DetailPosY      int   `json:"detailPosY"`
	DetailScale     int   `json:"detailScale"`
	DetailRotZ      int   `json:"detailRotZ"`
	StartDate       int64 `json:"startDate"`
	OutingMax       int   `json:"outingMax"`
	EffectID        int   `json:"effectId"`
}

type SupportCardInfo struct {
	CharaID      int    `json:"chara_id"`
	Gametora     string `json:"gametora"`
	ID           int    `json:"id"`
	Rarity       int    `json:"rarity"`
	RarityString string `json:"rarity_string"`
	StartDate    int64  `json:"start_date"`
	Title        string `json:"title"`
	TitleEN      string `json:"title_en"`
	Type         string `json:"type"`
	TypeIconURL  string `json:"type_icon_url"`
}
