package umacache

var (
	// Errors
	activeErrors map[string]string = make(map[string]string)

	// API URLs
	// Tracen Academy - Japanese data, though offers more stats
	tracBase             = "https://www.tracenacademy.com/api/"
	tracCharacters       = tracBase + "BasicCharaDataInfo"
	tracCharacterData    = tracBase + "CharaData"
	tracSupportCards     = tracBase + "BasicSupportCardDataInfo"
	tracSupportCardsData = tracBase + "SupportCardData"

	// Umapyoi - English, less data but offers Images
	umapBase            = "https://umapyoi.net/api/v1/"
	umapCharacterInfo   = umapBase + "character/" // Character ID on end
	umapSupportCardInfo = umapBase + "support/"   // Support Card ID on end
)

func Init() {
	getCharacters()
	getCharacterData()
	getSupportCards()
	getSupportCardData()
}
