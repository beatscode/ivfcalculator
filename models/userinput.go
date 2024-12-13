package models

import (
	"fmt"
	"log"
	"math"
)

type UserInput struct {
	Age                      int  `form:"age" json:"age"`
	WeightLbs                int  `form:"weight_lbs" json:"weight_lbs"`
	HeightFt                 int  `form:"height_ft" json:"height_ft"`
	HeightIn                 int  `form:"height_in" json:"height_in"`
	TubalFactor              bool `form:"tubal_factor" json:"tubal_factor"`
	MaleFactorInfertility    bool `form:"male_factor_infertility" json:"male_factor_infertility"`
	Endometriosis            bool `form:"endometriosis" json:"endometriosis"`
	OvulatoryDisorder        bool `form:"ovulatory_disorder" json:"ovulatory_disorder"`
	DiminishedOvarianReserve bool `form:"diminished_ovarian_reserve" json:"diminished_ovarian_reserve"`
	UterineFactor            bool `form:"uterine_factor" json:"uterine_factor"`
	OtherReason              bool `form:"other_reason" json:"other_reason"`
	UnexplainedInfertility   bool `form:"unexplained_infertility" json:"unexplained_infertility"`
	PriorPregnancies         int  `form:"prior_pregnancies" json:"prior_pregnancies"`
	LiveBirths               int  `form:"live_births" json:"live_births"`
}

func (input *UserInput) Validate() error {
	var err error
	if input.Age < 20 || input.Age > 50 {
		return fmt.Errorf("age is outside of range(20 - 50) needed for estimation")
	}

	if input.WeightLbs < 80 || input.WeightLbs > 300 {
		return fmt.Errorf("bmi is outside of range(80 - 300) needed for estimation")
	}
	return err
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

	bmiLinear := fi.FormulaBMILinearCoefficient * bmi
	bmiPower := fi.FormulaBMIPowerCoefficient * (math.Pow(bmi, fi.FormulaBMIPowerFactor))

	score += bmiLinear + bmiPower
	log.Printf("bmi: %f", bmi)

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

func (input *UserInput) SuccessRate(fi FormulaParameters) (score float64, successRate float64) {
	var sRate float64
	tmpScore := input.Score(fi)
	exp := math.Exp(tmpScore)
	sRate = exp / (1 + exp)
	percentage := sRate * 100
	return tmpScore, percentage
}
