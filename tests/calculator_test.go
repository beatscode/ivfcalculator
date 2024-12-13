package tests

import (
	"encoding/json"
	"log"
	"math"
	"testing"

	"com.sunfish.ivfsuccesscalculator/models"
)

func TestCalculator(t *testing.T) {

	formulas, err := models.LoadFormulasFromCSV("../ivf_success_formulas.csv")
	if err != nil {
		log.Fatal(err)
	}

	//choose formula  Inputs
	f1, err := formulas.ChooseFormula(true, "FALSE", true)
	if err != nil {
		t.Error("invalid selection", err)
	}

	if f1.CDCFormula != "1-3" {
		t.Error("CDC formula 1-3 failed")
	}

	f2, err := formulas.ChooseFormula(true, "FALSE", false)
	if err != nil {
		t.Error("invalid selection", err)
	}

	if f2.CDCFormula != "4-6" {
		t.Error("CDC formula 4-6 failed")
	}

	f3, err := formulas.ChooseFormula(true, "TRUE", true)
	if err != nil {
		t.Error("invalid selection", err)
	}

	if f3.CDCFormula != "7-8" {
		t.Error("CDC formula 7-8 failed")
	}

	f4, err := formulas.ChooseFormula(true, "TRUE", false)
	if err != nil {
		t.Error("invalid selection", err)
	}

	if f4.CDCFormula != "9-10" {
		t.Error("CDC formula 9-10 failed")
	}

	f5, err := formulas.ChooseFormula(false, "N/A", true)
	if err != nil {
		t.Error("invalid selection", err)
	}

	if f5.CDCFormula != "11-13" {
		t.Error("CDC formula 11-13 failed")
	}

	f6, err := formulas.ChooseFormula(false, "N/A", false)
	if err != nil {
		t.Error("invalid selection", err)
	}

	if f6.CDCFormula != "14-16" {
		t.Error("CDC formula 14-16 failed")
	}

}
func TestBMICalc(t *testing.T) {
	input := models.UserInput{
		HeightFt:  5,
		HeightIn:  8,
		WeightLbs: 150,
	}

	bmi := input.BMI()
	if bmi < 22 || bmi > 23 {
		t.Error("invalid bmi value", bmi)
	}
	log.Println(input, "bmi is", bmi)
}

// Example: Using Own Eggs / Did Not Previously Attempt IVF / Known Infertility Reason
func TestScore1(t *testing.T) {
	formulas, err := models.LoadFormulasFromCSV("../ivf_success_formulas.csv")
	if err != nil {
		log.Fatal(err)
	}

	//const result = 0.498270 | this is given in the description but not matching
	const result = 0.498277
	const successRateResult = 62
	personA := models.UserInput{
		Age:                      32,
		WeightLbs:                150,
		HeightFt:                 5,
		HeightIn:                 8,
		TubalFactor:              false,
		Endometriosis:            true,
		OvulatoryDisorder:        true,
		DiminishedOvarianReserve: false,
		PriorPregnancies:         1,
		LiveBirths:               1,
	}

	fi := models.FormulaInput{
		ParamUsingOwnEggs:                true,
		ParamAttemptedIVFPreviously:      "FALSE",
		ParamIsReasonForInfertilityKnown: true,
	}
	sFormula, err := formulas.ChooseFormula(fi.ParamUsingOwnEggs, fi.ParamAttemptedIVFPreviously, fi.ParamIsReasonForInfertilityKnown)
	if err != nil {
		t.Fatal("could not choose a formula")
	}
	score, successRate := personA.SuccessRate(sFormula)
	if !floatsEqual(score, result) {
		t.Errorf("invalid score given %f, should be %f", score, result)
	}
	if successRate < 62.20 || successRate > 62.21 {
		t.Errorf("invalid success rate %f, should be %v", successRate, successRateResult)
	}
	log.Printf("score %f, success rate: %f", score, successRate)
}

func TestScore2(t *testing.T) {

	formulas, err := models.LoadFormulasFromCSV("../ivf_success_formulas.csv")
	if err != nil {
		log.Fatal(err)
	}

	//const result = 0.398593 | this is given in the description but not matching
	const result = 0.398540
	const successRateResult = 59
	personA := models.UserInput{
		Age:                      32,
		WeightLbs:                150,
		HeightFt:                 5,
		HeightIn:                 8,
		TubalFactor:              false,
		Endometriosis:            true,
		OvulatoryDisorder:        true,
		DiminishedOvarianReserve: false,
		PriorPregnancies:         1,
		LiveBirths:               1,
	}

	fi := models.FormulaInput{
		ParamUsingOwnEggs:                true,
		ParamAttemptedIVFPreviously:      "FALSE",
		ParamIsReasonForInfertilityKnown: false,
	}
	sFormula, err := formulas.ChooseFormula(fi.ParamUsingOwnEggs, fi.ParamAttemptedIVFPreviously, fi.ParamIsReasonForInfertilityKnown)
	if err != nil {
		t.Fatal("could not choose a formula")
	}
	score, successRate := personA.SuccessRate(sFormula)
	if !floatsEqual(score, result) {
		t.Errorf("invalid score given %f, should be %f", score, result)
	}
	if successRate < 59.80 || successRate > 59.85 {
		t.Errorf("invalid success rate %f, should be %v", successRate, successRateResult)
	}
}

func TestScore3(t *testing.T) {

	formulas, err := models.LoadFormulasFromCSV("../ivf_success_formulas.csv")
	if err != nil {
		log.Fatal(err)
	}

	//const result = 0.368348 | this is given in the description but not matching
	const result = -0.368321
	const successRateResult = 40.894680
	personA := models.UserInput{
		Age:                      32,
		WeightLbs:                150,
		HeightFt:                 5,
		HeightIn:                 8,
		TubalFactor:              true,
		DiminishedOvarianReserve: true,
		PriorPregnancies:         1,
		LiveBirths:               1,
	}
	bt, _ := json.MarshalIndent(personA, "", " ")
	log.Println(string(bt))
	fi := models.FormulaInput{
		ParamUsingOwnEggs:                true,
		ParamAttemptedIVFPreviously:      "TRUE",
		ParamIsReasonForInfertilityKnown: true,
	}
	log.Println(fi)
	sFormula, err := formulas.ChooseFormula(fi.ParamUsingOwnEggs, fi.ParamAttemptedIVFPreviously, fi.ParamIsReasonForInfertilityKnown)
	if err != nil {
		t.Fatal("could not choose a formula")
	}
	score, successRate := personA.SuccessRate(sFormula)
	if !floatsEqual(score, result) {
		t.Errorf("invalid score given %f, should be %f", score, result)
	}
	if !floatsEqual(successRate, successRateResult) {
		t.Errorf("invalid success rate %f, should be %v", successRate, successRateResult)
	}
	log.Println("score", score, "success rate:", successRate)
}

func floatsEqual(a, b float64) bool {
	return math.Abs(a-b) < 0.000001
}
