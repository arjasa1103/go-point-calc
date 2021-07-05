package main

type TrelloBoard struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	Desc              string            `json:"desc"`
	DescData          interface{}       `json:"descData"`
	Closed            bool              `json:"closed"`
	DateClosed        interface{}       `json:"dateClosed"`
	IDOrganization    string            `json:"idOrganization"`
	ShortLink         string            `json:"shortLink"`
	PowerUPS          []interface{}     `json:"powerUps"`
	DateLastActivity  string            `json:"dateLastActivity"`
	IDTags            []interface{}     `json:"idTags"`
	DatePluginDisable interface{}       `json:"datePluginDisable"`
	CreationMethod    interface{}       `json:"creationMethod"`
	IDBoardSource     interface{}       `json:"idBoardSource"`
	IDMemberCreator   string            `json:"idMemberCreator"`
	IDEnterprise      interface{}       `json:"idEnterprise"`
	Pinned            bool              `json:"pinned"`
	Starred           bool              `json:"starred"`
	URL               string            `json:"url"`
	ShortURL          string            `json:"shortUrl"`
	EnterpriseOwned   bool              `json:"enterpriseOwned"`
	IxUpdate          string            `json:"ixUpdate"`
	Limits            TrelloBoardLimits `json:"limits"`
	Subscribed        bool              `json:"subscribed"`
	TemplateGallery   interface{}       `json:"templateGallery"`
	PremiumFeatures   []interface{}     `json:"premiumFeatures"`
	DateLastView      string            `json:"dateLastView"`
	LabelNames        LabelNames        `json:"labelNames"`
	Prefs             Prefs             `json:"prefs"`
	Actions           []Action          `json:"actions"`
	Cards             []CardElement     `json:"cards"`
	Labels            []Label           `json:"labels"`
	Lists             []ListElement     `json:"lists"`
	Members           []MemberElement   `json:"members"`
	Checklists        []interface{}     `json:"checklists"`
	CustomFields      []CustomField     `json:"customFields"`
	Memberships       []Membership      `json:"memberships"`
	PluginData        []interface{}     `json:"pluginData"`
}

type Action struct {
	ID              string         `json:"id"`
	IDMemberCreator string         `json:"idMemberCreator"`
	Data            Data           `json:"data"`
	Type            string         `json:"type"`
	Date            string         `json:"date"`
	AppCreator      interface{}    `json:"appCreator"`
	Limits          NonPublicClass `json:"limits"`
	Member          *ActionMember  `json:"member,omitempty"`
	MemberCreator   MemberCreator  `json:"memberCreator"`
}

type Data struct {
	IDMember      *string         `json:"idMember,omitempty"`
	Board         Board           `json:"board"`
	Old           *Old            `json:"old,omitempty"`
	Card          *DataCard       `json:"card,omitempty"`
	ListBefore    *ListAfterClass `json:"listBefore,omitempty"`
	ListAfter     *ListAfterClass `json:"listAfter,omitempty"`
	List          *ListAfterClass `json:"list,omitempty"`
	MemberType    *string         `json:"memberType,omitempty"`
	IDMemberAdded *string         `json:"idMemberAdded,omitempty"`
	CardSource    *Board          `json:"cardSource,omitempty"`
}

type Board struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShortLink string `json:"shortLink"`
	IDShort   *int64 `json:"idShort,omitempty"`
}

type DataCard struct {
	IDList    *string `json:"idList,omitempty"`
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	IDShort   int64   `json:"idShort"`
	ShortLink string  `json:"shortLink"`
	Pos       *int64  `json:"pos,omitempty"`
}

type ListAfterClass struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Old struct {
	IDList *string `json:"idList,omitempty"`
	Pos    *int64  `json:"pos,omitempty"`
}

type NonPublicClass struct {
}

type ActionMember struct {
	ID                 string      `json:"id"`
	Username           string      `json:"username"`
	ActivityBlocked    bool        `json:"activityBlocked"`
	AvatarHash         interface{} `json:"avatarHash"`
	AvatarURL          interface{} `json:"avatarUrl"`
	FullName           string      `json:"fullName"`
	IDMemberReferrer   interface{} `json:"idMemberReferrer"`
	Initials           string      `json:"initials"`
	NonPublic          NonPublic   `json:"nonPublic"`
	NonPublicAvailable bool        `json:"nonPublicAvailable"`
}

type NonPublic struct {
	FullName *string `json:"fullName,omitempty"`
	Initials *string `json:"initials,omitempty"`
}

type MemberCreator struct {
	ID                 string         `json:"id"`
	Username           string         `json:"username"`
	ActivityBlocked    bool           `json:"activityBlocked"`
	AvatarHash         interface{}    `json:"avatarHash"`
	AvatarURL          interface{}    `json:"avatarUrl"`
	FullName           string         `json:"fullName"`
	IDMemberReferrer   interface{}    `json:"idMemberReferrer"`
	Initials           string         `json:"initials"`
	NonPublic          NonPublicClass `json:"nonPublic"`
	NonPublicAvailable bool           `json:"nonPublicAvailable"`
}

type CardElement struct {
	ID                    string            `json:"id"`
	Address               interface{}       `json:"address"`
	CheckItemStates       interface{}       `json:"checkItemStates"`
	Closed                bool              `json:"closed"`
	Coordinates           interface{}       `json:"coordinates"`
	CreationMethod        interface{}       `json:"creationMethod"`
	DateLastActivity      string            `json:"dateLastActivity"`
	Desc                  string            `json:"desc"`
	DescData              interface{}       `json:"descData"`
	DueReminder           interface{}       `json:"dueReminder"`
	IDBoard               string            `json:"idBoard"`
	IDLabels              []interface{}     `json:"idLabels"`
	IDList                string            `json:"idList"`
	IDMembersVoted        []interface{}     `json:"idMembersVoted"`
	IDShort               int64             `json:"idShort"`
	IDAttachmentCover     interface{}       `json:"idAttachmentCover"`
	LocationName          interface{}       `json:"locationName"`
	ManualCoverAttachment bool              `json:"manualCoverAttachment"`
	Name                  string            `json:"name"`
	Pos                   int64             `json:"pos"`
	ShortLink             string            `json:"shortLink"`
	IsTemplate            bool              `json:"isTemplate"`
	CardRole              interface{}       `json:"cardRole"`
	Badges                Badges            `json:"badges"`
	DueComplete           bool              `json:"dueComplete"`
	Due                   *string           `json:"due"`
	Email                 string            `json:"email"`
	IDChecklists          []interface{}     `json:"idChecklists"`
	IDMembers             []interface{}     `json:"idMembers"`
	Labels                []interface{}     `json:"labels"`
	Limits                CardLimits        `json:"limits"`
	ShortURL              string            `json:"shortUrl"`
	Start                 interface{}       `json:"start"`
	Subscribed            bool              `json:"subscribed"`
	URL                   string            `json:"url"`
	Cover                 Cover             `json:"cover"`
	Attachments           []Attachment      `json:"attachments"`
	PluginData            []interface{}     `json:"pluginData"`
	CustomFieldItems      []CustomFieldItem `json:"customFieldItems"`
}

type Attachment struct {
	Bytes     *int64        `json:"bytes"`
	Date      string        `json:"date"`
	EdgeColor interface{}   `json:"edgeColor"`
	IDMember  string        `json:"idMember"`
	IsUpload  bool          `json:"isUpload"`
	MIMEType  string        `json:"mimeType"`
	Name      string        `json:"name"`
	Previews  []interface{} `json:"previews"`
	URL       string        `json:"url"`
	Pos       int64         `json:"pos"`
	FileName  *string       `json:"fileName,omitempty"`
	ID        string        `json:"id"`
}

type Badges struct {
	AttachmentsByType     AttachmentsByType `json:"attachmentsByType"`
	Location              bool              `json:"location"`
	Votes                 int64             `json:"votes"`
	ViewingMemberVoted    bool              `json:"viewingMemberVoted"`
	Subscribed            bool              `json:"subscribed"`
	Fogbugz               string            `json:"fogbugz"`
	CheckItems            int64             `json:"checkItems"`
	CheckItemsChecked     int64             `json:"checkItemsChecked"`
	CheckItemsEarliestDue interface{}       `json:"checkItemsEarliestDue"`
	Comments              int64             `json:"comments"`
	Attachments           int64             `json:"attachments"`
	Description           bool              `json:"description"`
	Due                   *string           `json:"due"`
	DueComplete           bool              `json:"dueComplete"`
	Start                 interface{}       `json:"start"`
}

type AttachmentsByType struct {
	Trello Trello `json:"trello"`
}

type Trello struct {
	Board int64 `json:"board"`
	Card  int64 `json:"card"`
}

type Cover struct {
	IDAttachment         interface{} `json:"idAttachment"`
	Color                interface{} `json:"color"`
	IDUploadedBackground interface{} `json:"idUploadedBackground"`
	Size                 string      `json:"size"`
	Brightness           string      `json:"brightness"`
	IDPlugin             interface{} `json:"idPlugin"`
}

type CustomFieldItem struct {
	ID            string `json:"id"`
	Value         Value  `json:"value"`
	IDCustomField string `json:"idCustomField"`
	IDModel       string `json:"idModel"`
	ModelType     string `json:"modelType"`
}

type Value struct {
	Number string `json:"number"`
}

type CardLimits struct {
	Attachments Stickers `json:"attachments"`
	Checklists  Stickers `json:"checklists"`
	Stickers    Stickers `json:"stickers"`
}

type Stickers struct {
	PerCard PerBoard `json:"perCard"`
}

type PerBoard struct {
	Status    string `json:"status"`
	DisableAt int64  `json:"disableAt"`
	WarnAt    int64  `json:"warnAt"`
}

type CustomField struct {
	ID               string  `json:"id"`
	IDModel          string  `json:"idModel"`
	ModelType        string  `json:"modelType"`
	FieldGroup       string  `json:"fieldGroup"`
	Display          Display `json:"display"`
	Name             string  `json:"name"`
	Pos              int64   `json:"pos"`
	Type             string  `json:"type"`
	IsSuggestedField bool    `json:"isSuggestedField"`
}

type Display struct {
	CardFront bool `json:"cardFront"`
}

type LabelNames struct {
	Green  string `json:"green"`
	Yellow string `json:"yellow"`
	Orange string `json:"orange"`
	Red    string `json:"red"`
	Purple string `json:"purple"`
	Blue   string `json:"blue"`
	Sky    string `json:"sky"`
	Lime   string `json:"lime"`
	Pink   string `json:"pink"`
	Black  string `json:"black"`
}

type Label struct {
	ID      string `json:"id"`
	IDBoard string `json:"idBoard"`
	Name    string `json:"name"`
	Color   string `json:"color"`
}

type TrelloBoardLimits struct {
	Attachments        Attachments        `json:"attachments"`
	Boards             Boards             `json:"boards"`
	Cards              PurpleCards        `json:"cards"`
	Checklists         Attachments        `json:"checklists"`
	CheckItems         CheckItems         `json:"checkItems"`
	CustomFields       CustomFields       `json:"customFields"`
	CustomFieldOptions CustomFieldOptions `json:"customFieldOptions"`
	Labels             CustomFields       `json:"labels"`
	Lists              Lists              `json:"lists"`
	Stickers           Stickers           `json:"stickers"`
	Reactions          Reactions          `json:"reactions"`
}

type Attachments struct {
	PerBoard PerBoard `json:"perBoard"`
	PerCard  PerBoard `json:"perCard"`
}

type Boards struct {
	TotalMembersPerBoard PerBoard `json:"totalMembersPerBoard"`
}

type PurpleCards struct {
	OpenPerBoard  PerBoard `json:"openPerBoard"`
	OpenPerList   PerBoard `json:"openPerList"`
	TotalPerBoard PerBoard `json:"totalPerBoard"`
	TotalPerList  PerBoard `json:"totalPerList"`
}

type CheckItems struct {
	PerChecklist PerBoard `json:"perChecklist"`
}

type CustomFieldOptions struct {
	PerField PerBoard `json:"perField"`
}

type CustomFields struct {
	PerBoard PerBoard `json:"perBoard"`
}

type Lists struct {
	OpenPerBoard  PerBoard `json:"openPerBoard"`
	TotalPerBoard PerBoard `json:"totalPerBoard"`
}

type Reactions struct {
	PerAction       PerBoard `json:"perAction"`
	UniquePerAction PerBoard `json:"uniquePerAction"`
}

type ListElement struct {
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	Closed         bool        `json:"closed"`
	Pos            int64       `json:"pos"`
	SoftLimit      interface{} `json:"softLimit"`
	CreationMethod interface{} `json:"creationMethod"`
	IDBoard        string      `json:"idBoard"`
	Limits         ListLimits  `json:"limits"`
	Subscribed     bool        `json:"subscribed"`
}

type ListLimits struct {
	Cards FluffyCards `json:"cards"`
}

type FluffyCards struct {
	OpenPerList  PerBoard `json:"openPerList"`
	TotalPerList PerBoard `json:"totalPerList"`
}

type MemberElement struct {
	ID                       string        `json:"id"`
	Bio                      string        `json:"bio"`
	BioData                  interface{}   `json:"bioData"`
	Confirmed                bool          `json:"confirmed"`
	MemberType               string        `json:"memberType"`
	Username                 string        `json:"username"`
	ActivityBlocked          bool          `json:"activityBlocked"`
	AvatarHash               interface{}   `json:"avatarHash"`
	AvatarURL                interface{}   `json:"avatarUrl"`
	FullName                 string        `json:"fullName"`
	IDEnterprise             interface{}   `json:"idEnterprise"`
	IDEnterprisesDeactivated []interface{} `json:"idEnterprisesDeactivated"`
	IDMemberReferrer         interface{}   `json:"idMemberReferrer"`
	IDPremOrgsAdmin          []interface{} `json:"idPremOrgsAdmin"`
	Initials                 string        `json:"initials"`
	NonPublic                NonPublic     `json:"nonPublic"`
	NonPublicAvailable       bool          `json:"nonPublicAvailable"`
	Products                 []interface{} `json:"products"`
	URL                      string        `json:"url"`
	Status                   string        `json:"status"`
}

type Membership struct {
	ID          string `json:"id"`
	IDMember    string `json:"idMember"`
	MemberType  string `json:"memberType"`
	Unconfirmed bool   `json:"unconfirmed"`
	Deactivated bool   `json:"deactivated"`
}

type Prefs struct {
	PermissionLevel       string                  `json:"permissionLevel"`
	HideVotes             bool                    `json:"hideVotes"`
	Voting                string                  `json:"voting"`
	Comments              string                  `json:"comments"`
	Invitations           string                  `json:"invitations"`
	SelfJoin              bool                    `json:"selfJoin"`
	CardCovers            bool                    `json:"cardCovers"`
	IsTemplate            bool                    `json:"isTemplate"`
	CardAging             string                  `json:"cardAging"`
	CalendarFeedEnabled   bool                    `json:"calendarFeedEnabled"`
	Background            string                  `json:"background"`
	BackgroundImage       string                  `json:"backgroundImage"`
	BackgroundImageScaled []BackgroundImageScaled `json:"backgroundImageScaled"`
	BackgroundTile        bool                    `json:"backgroundTile"`
	BackgroundBrightness  string                  `json:"backgroundBrightness"`
	BackgroundBottomColor string                  `json:"backgroundBottomColor"`
	BackgroundTopColor    string                  `json:"backgroundTopColor"`
	CanBePublic           bool                    `json:"canBePublic"`
	CanBeEnterprise       bool                    `json:"canBeEnterprise"`
	CanBeOrg              bool                    `json:"canBeOrg"`
	CanBePrivate          bool                    `json:"canBePrivate"`
	CanInvite             bool                    `json:"canInvite"`
}

type BackgroundImageScaled struct {
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
	URL    string `json:"url"`
}

type CardNamePoints struct {
	Name   string `json: name`
	Points string `json: points`
}
