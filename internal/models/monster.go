package models

type Monster struct {
	Name           string
	Description    string
	ImageIcon      string
	ImageWallpaper string
	Health         int16
	Size           float32
	Quests         []Quest
}

func Fatalis() *Monster {
	return &Monster{
		Name: "Fatalis",
		Description: "A legendary black dragon known only as Fatalis. Rumored to have destroyed a kingdom in a single night, and has taken its castle for a roost.\n" +
			"As Long as its horns are intact, overcoming its final form's breath attack may be impossible. Cannons and ballistae can topple it. Flinch shots when " +
			"its flying or standing will lower its head.",
		ImageIcon:      "/static/images/fatalis.jpg",
		ImageWallpaper: "/static/images/fatalis_icon.webp",
		Health:         5500,
		Size:           5088,
	}
}

func Lagiacrus() *Monster {

	return &Monster{
		Name: "Lagiacrus",
		Description: "Known as Sea Wyverns, Lagiacrus are at the top of the aquatic food chain. Feared by sailors as 'The Lords of the Seas'," +
			" they store enough electricity in their spinal organs to make the oceans surge. Occasionally seen resting on land.",
		ImageIcon:      "/static/images/lagiacrus.jpg",
		ImageWallpaper: "/static/images/lagiacrus_icon.webp",
		Health:         3845,
		Size:           2648,
	}

}
