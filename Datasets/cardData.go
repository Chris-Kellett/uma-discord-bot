package datasets

type CardData struct {
	ID                int `json:"id"`
	CharacterID       int `json:"charaId"`
	DefaultRarity     int `json:"defaultRarity"`
	LimitedCharacter  int `json:"limitedChara"`
	AvailableSkillSet int `json:"availableSkillSetId"`
	TalentSpeed       int `json:"talentSpeed"`
	TalentStamina     int `json:"talentStamina"`
	TalentPow         int `json:"talentPow"`
	TalentGuts        int `json:"talentGuts"`
	TalentWiz         int `json:"talentWiz"`
	TalentGroupID     int `json:"talentGroupId"`
	BgID              int `json:"bgId"`
	GetPieceID        int `json:"getPieceId"`
	RunningStyle      int `json:"runningStyle"`
}
