package cards

type CardType string
const (
	// Main Monsters
	EffectMonster       CardType = "Effect Monster"
	FlipEffectMonster   CardType = "Flip Effect Monster"
	GeminiMonster       CardType = "Gemini Monster"
	NormalMonster       CardType = "Normal Monster"
	RitualEffectMonster CardType = "Ritual Effect Monster"
	FusionMonster       CardType = "Fusion Monster"
	RitualMonster       CardType = "Ritual Monster"
	UnionEffectMonster  CardType = "Union Effect Monster"
	SpiritMonster       CardType = "Spirit Monster"
	ToonMonster         CardType = "Toon Monster"
	// Tuner
	NormalTunerMonster     CardType = "Normal Tuner Monster"
	TunerMonster           CardType = "Tuner Monster"
	FlipTunerEffectMonster CardType = "Flip Tuner Effect Monster"
	// Synchro
	SynchroMonster      CardType = "Synchro Monster"
	SynchroTunerMonster CardType = "Synchro Tuner Monster"
	// XYZ
	XYZMonster CardType = "XYZ Monster"

    // Link
	LinkMonster CardType = "Link Monster"

	// Pendulum
	PendulumNormalMonster       CardType = "Pendulum Normal Monster"
	PendulumEffectMonster       CardType = "Pendulum Effect Monster"
	PendulumTunerEffectMonster  CardType = "Pendulum Tuner Effect Monster"
	PendulumEffectRitualMonster CardType = "Pendulum Effect Ritual Monster"
	PendulumFlipEffectMonster   CardType = "Pendulum Flip Effect Monster"
	PendulumEffectFusionMonster CardType = "Pendulum Effect Fusion Monster"
	SynchroPendulumEffectMonster CardType = "Synchro Pendulum Effect Monster"
	XYZPendulumEffectMonster     CardType = "XYZ Pendulum Effect Monster"

	Spell CardType = "Spell Card"

	Skill CardType = "Skill Card"

	Trap CardType = "Trap Card"

	Token CardType = "Token"
)

type Card interface {
	ID() string
	Name() string
	CardType() CardType
	FrameType() string
	Description() string
}

type MonsterCard struct {
	ID           string
	Name         string
	CardType     CardType
	TypeLine     []string
	Attack       int
	Defence      int
	Level        int
	Attribute    string
	FrameType    string
	Description  string
	Race         string
	Archetype    string
	CardImageURL string
}

type SpellCard struct {
	ID           string
	Name         string
	CardType     CardType
	FrameType    string
	Description  string
	Race         string
	Archetype    string
	CardImageURL string
}

type TrapCard struct {
	ID           string
	Name         string
	CardType     CardType
	FrameType    string
	Description  string
	Race         string
	Archetype    string
	CardImageURL string
}

type SkillCard struct {
	ID           string
	Name         string
	CardType     CardType
	FrameType    string
	Description  string
	Race         string
	CardImageURL string
}

type TokenCard struct {
	ID           string
	Name         string
	CardType     CardType
	FrameType    string
	Description  string
	Race         string
	Archetype    string
	CardImageURL string	
}

type PendulumCard struct {
	ID           string
	Name         string
	CardType     CardType
	TypeLine     []string
	Attack       int
	Defence      int
	Level        int
	Attribute    string
	FrameType    string
	Description  string
	Race         string
	Archetype    string
	CardImageURL string
	Scale int
}

type LinkMonsterCard struct {
	ID           string
	Name         string
	CardType     CardType
	TypeLine     []string
	Attack       int
	Defence      int
	Level        int
	Attribute    string
	FrameType    string
	Description  string
	Race         string
	Archetype    string
	CardImageURL string
	LinkVal int
	LinkMarkers []string
}