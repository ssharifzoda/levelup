package domain

type Public struct {
	Gender        int    `json:"gender_id"`
	FamilyStatus  int    `json:"family_status_id"`
	Age           int    `json:"age"`
	GoalToLife    string `json:"goal_to_life"`
	BigFear       string `json:"big_fear"`
	TemperamentId int    `json:"temperament_id"`
}
type TemperamentTest struct {
	One int
}
