package models

import (
	"fmt"
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
		return fmt.Errorf("weight is outside of range(80 - 300) needed for estimation")
	}
	return err
}
func (input *UserInput) BMI() float64 {
	h := input.HeightFt*12 + input.HeightIn
	return float64(input.WeightLbs) * 703 / math.Pow(float64(h), 2)
}

// Score calulates with formula
func (input *UserInput) Score(fi FormulaParameters) float64 {
	var score float64
	bmi := input.BMI()
	age := float64(input.Age)
	score = fi.FormulaIntercept

	ageLinear := fi.FormulaAgeLinearCoefficient * age
	agePower := fi.FormulaAgePowerCoefficient * math.Pow(age, fi.FormulaAgePowerFactor)
	score += ageLinear + agePower

	bmiLinear := fi.FormulaBMILinearCoefficient * bmi
	bmiPower := fi.FormulaBMIPowerCoefficient * (math.Pow(bmi, fi.FormulaBMIPowerFactor))

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

func (input *UserInput) SuccessRate(fi FormulaParameters) (score float64, successRate float64) {
	var sRate float64
	tmpScore := input.Score(fi)
	exp := math.Exp(tmpScore)
	sRate = exp / (1 + exp)
	percentage := sRate * 100
	return tmpScore, percentage
}
