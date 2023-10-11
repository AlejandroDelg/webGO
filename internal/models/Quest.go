package models

type Quest struct {
	Name           string
	MonsterInQuest *Monster
	Description    string
	Location       string
	MaxTime        int8
}

func QuestFatalis() Quest {
	return Quest{
		Name:           "M★6 THE BLACK DRAGON",
		MonsterInQuest: Fatalis(),
		Description:    "It is time for us to confront Fatalis. We, along with every nation and organization known, shall rid this world of this fabled scourge. When you're ready, we'll set sail.",
		Location:       "Castle Schrade",
		MaxTime:        30,
	}
}
