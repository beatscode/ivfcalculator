package models

type RequestInput struct {
	Age                      int    `form:"age"`
	WeightLbs                int    `form:"weight_lbs"`
	HeightFt                 int    `form:"height_ft"`
	HeightIn                 int    `form:"height_in"`
	TubalFactor              string `form:"tubal_factor"`               // Changed to string
	MaleFactorInfertility    string `form:"male_factor_infertility"`    // Changed to string
	Endometriosis            string `form:"endometriosis"`              // Changed to string
	OvulatoryDisorder        string `form:"ovulatory_disorder"`         // Changed to string
	DiminishedOvarianReserve string `form:"diminished_ovarian_reserve"` // Changed to string
	UterineFactor            string `form:"uterine_factor"`             // Changed to string
	OtherReason              string `form:"other_reason"`               // Changed to string
	UnexplainedInfertility   string `form:"unexplained_infertility"`    // Changed to string
	PriorPregnancies         int    `form:"prior_pregnancies"`
	LiveBirths               int    `form:"live_births"`
	EggSource                string `form:"eggSource"`
	IvfUsed                  string `form:"ivfUsed"`
	IvfReason                string `form:"donotknow"`
}

// Helper method to convert the string "on" to boolean
func (u RequestInput) IsTubalFactor() bool {
	return u.TubalFactor == "on"
}

func (u RequestInput) IsMaleFactorInfertility() bool {
	return u.MaleFactorInfertility == "on"
}

func (u RequestInput) IsEndometriosis() bool {
	return u.Endometriosis == "on"
}

func (u RequestInput) IsOvulatoryDisorder() bool {
	return u.OvulatoryDisorder == "on"
}

func (u RequestInput) IsDiminishedOvarianReserve() bool {
	return u.DiminishedOvarianReserve == "on"
}

func (u RequestInput) IsUterineFactor() bool {
	return u.UterineFactor == "on"
}

func (u RequestInput) IsOtherReason() bool {
	return u.OtherReason == "on"
}

func (u RequestInput) IsUnexplainedInfertility() bool {
	return u.UnexplainedInfertility == "on"
}

func (u RequestInput) NoIvfReason() bool {
	return u.IvfReason == "on"
}

func (u RequestInput) ConvertToUserInput() UserInput {
	ui := UserInput{}
	ui.Age = u.Age
	ui.WeightLbs = u.WeightLbs
	ui.HeightFt = u.HeightFt
	ui.HeightIn = u.HeightIn
	ui.TubalFactor = u.IsTubalFactor()
	ui.MaleFactorInfertility = u.IsMaleFactorInfertility()
	ui.Endometriosis = u.IsEndometriosis()
	ui.OvulatoryDisorder = u.IsOvulatoryDisorder()
	ui.DiminishedOvarianReserve = u.IsDiminishedOvarianReserve()
	ui.UterineFactor = u.IsUterineFactor()
	ui.OtherReason = u.IsOtherReason()
	ui.PriorPregnancies = u.PriorPregnancies
	ui.LiveBirths = u.LiveBirths
	return ui
}
