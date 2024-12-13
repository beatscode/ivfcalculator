package models

import (
	"fmt"
	"math"
)

type UserInput struct {
	Age                      int  `json:"age"`
	WeightLbs                int  `json:"weight_lbs"`
	HeightFt                 int  `json:"height_ft"`
	HeightIn                 int  `json:"height_in"`
	TubalFactor              bool `json:"tubal_factor"`
	MaleFactorInfertility    bool `json:"male_factor_infertility"`
	Endometriosis            bool `json:"endometriosis"`
	OvulatoryDisorder        bool `json:"ovulatory_disorder"`
	DiminishedOvarianReserve bool `json:"diminished_ovarian_reserve"`
	UterineFactor            bool `json:"uterine_factor"`
	OtherReason              bool `json:"other_reason"`
	UnexplainedInfertility   bool `json:"unexplained_infertility"`
	PriorPregnancies         int  `json:"prior_pregnancies"`
	LiveBirths               int  `json:"live_births"`
}

func (input *UserInput) BMI() float64 {
	x := input.HeightFt*12 + input.HeightIn
	return float64(input.WeightLbs) * 703 / math.Pow(float64(x), 2)
}

//Score calulates with formula
/*
score =
formula_intercept +
formula_age_linear_component ✕ user_age + formula_age_power_coefficient ✕ (user_age ^ formula_age_power_factor) +
formula_bmi_linear_coefficient ✕ user_bmi + formula_bmi_power_coefficient ✕ (user_bmi ^ formula_bmi_power_factor) +
formula_tubal_factor_value +
formula_male_factor_infertility_value +
formula_endometriosis_value +
formula_ovulator_disorder_value +
formula_diminished_ovarian_reserve_value +
formula_uterine_factor_value +
formula_other_reason_value +
formula_unexplained_infertility_value +
formula_prior_pregnancies_value +
formula_live_births_value
*/
func (input *UserInput) Score(fi FormulaParameters) float64 {
	var score float64
	//pemdas
	bmi := input.BMI()
	age := float64(input.Age)
	score = fi.FormulaIntercept

	ageLinear := fi.FormulaAgeLinearCoefficient * age
	agePower := fi.FormulaAgePowerCoefficient * math.Pow(age, fi.FormulaAgePowerFactor)
	score += ageLinear + agePower
	fmt.Printf("After age terms (linear=%f, power=%f): %f\n", ageLinear, agePower, score)
	bmiLinear := fi.FormulaBMILinearCoefficient * bmi
	bmiPower := fi.FormulaBMIPowerCoefficient * (math.Pow(bmi, fi.FormulaBMIPowerFactor))
	fmt.Printf("After BMI terms (linear=%f, power=%f): %f\n", bmiLinear, bmiPower, score)

	score += bmiLinear + bmiPower

	score += fi.GetTubalFactorValue(input.TubalFactor)
	score += fi.GetMaleInfertilityFactorValue(input.MaleFactorInfertility)
	score += fi.GetEndometriosisFactorValue(input.Endometriosis)
	score += fi.GetOvulatorDisorderValue(input.OvulatoryDisorder)
	score += fi.GetDiminishedOvarianResereValue(input.DiminishedOvarianReserve)
	score += fi.GetUterineFactorValue(input.UterineFactor)
	score += fi.GetUnexplainedInfertilityValue(input.UnexplainedInfertility)
	score += fi.GetPriorPregnanciesValue(input.PriorPregnancies)
	score += fi.GetLiveBirthsValue(input.LiveBirths)

	return score
}
func (input *UserInput) SuccessRate() float64 {
	var rate float64

	return rate
}
