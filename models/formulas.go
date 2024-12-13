package models

import (
	"errors"
	"os"

	"github.com/gocarina/gocsv"
)

type FormulaParameters struct {
	FormulaInput

	FormulaIntercept                          float64 `csv:"formula_intercept"`
	FormulaAgeLinearCoefficient               float64 `csv:"formula_age_linear_coefficient"`
	FormulaAgePowerCoefficient                float64 `csv:"formula_age_power_coefficient"`
	FormulaAgePowerFactor                     float64 `csv:"formula_age_power_factor"`
	FormulaBMILinearCoefficient               float64 `csv:"formula_bmi_linear_coefficient"`
	FormulaBMIPowerCoefficient                float64 `csv:"formula_bmi_power_coefficient"`
	FormulaBMIPowerFactor                     float64 `csv:"formula_bmi_power_factor"`
	FormulaTubalFactorTrueValue               float64 `csv:"formula_tubal_factor_true_value"`
	FormulaTubalFactorFalseValue              float64 `csv:"formula_tubal_factor_false_value"`
	FormulaMaleFactorInfertilityTrueValue     float64 `csv:"formula_male_factor_infertility_true_value"`
	FormulaMaleFactorInfertilityFalseValue    float64 `csv:"formula_male_factor_infertility_false_value"`
	FormulaEndometriosisTrueValue             float64 `csv:"formula_endometriosis_true_value"`
	FormulaEndometriosisFalseValue            float64 `csv:"formula_endometriosis_false_value"`
	FormulaOvulatoryDisorderTrueValue         float64 `csv:"formula_ovulatory_disorder_true_value"`
	FormulaOvulatoryDisorderFalseValue        float64 `csv:"formula_ovulatory_disorder_false_value"`
	FormulaDiminishedOvarianReserveTrueValue  float64 `csv:"formula_diminished_ovarian_reserve_true_value"`
	FormulaDiminishedOvarianReserveFalseValue float64 `csv:"formula_diminished_ovarian_reserve_false_value"`
	FormulaUterineFactorTrueValue             float64 `csv:"formula_uterine_factor_true_value"`
	FormulaUterineFactorFalseValue            float64 `csv:"formula_uterine_factor_false_value"`
	FormulaOtherReasonTrueValue               float64 `csv:"formula_other_reason_true_value"`
	FormulaOtherReasonFalseValue              float64 `csv:"formula_other_reason_false_value"`
	FormulaUnexplainedInfertilityTrueValue    float64 `csv:"formula_unexplained_infertility_true_value"`
	FormulaUnexplainedInfertilityFalseValue   float64 `csv:"formula_unexplained_infertility_false_value"`
	FormulaPriorPregnancies0Value             float64 `csv:"formula_prior_pregnancies_0_value"`
	FormulaPriorPregnancies1Value             float64 `csv:"formula_prior_pregnancies_1_value"`
	FormulaPriorPregnancies2PlusValue         float64 `csv:"formula_prior_pregnancies_2+_value"`
	FormulaPriorLiveBirths0Value              float64 `csv:"formula_prior_live_births_0_value"`
	FormulaPriorLiveBirths1Value              float64 `csv:"formula_prior_live_births_1_value"`
	FormulaPriorLiveBirths2PlusValue          float64 `csv:"formula_prior_live_births_2+_value"`
}

type FormulaInput struct {
	ParamUsingOwnEggs                bool   `csv:"param_using_own_eggs"`
	ParamAttemptedIVFPreviously      string `csv:"param_attempted_ivf_previously"`
	ParamIsReasonForInfertilityKnown bool   `csv:"param_is_reason_for_infertility_known"`
	CDCFormula                       string `csv:"cdc_formula"`
}

type Formulas []FormulaParameters

func LoadFormulasFromCSV(filename string) (Formulas, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var formulas []FormulaParameters
	if err := gocsv.UnmarshalFile(file, &formulas); err != nil {
		return nil, err
	}

	return formulas, nil
}

func (formulas Formulas) ChooseFormula(usingOwnEggs bool, attemptedIVFPreviously string, reasonForInfertilityKnown bool) (FormulaParameters, error) {
	for _, formula := range formulas {
		if formula.ParamUsingOwnEggs == usingOwnEggs && formula.ParamAttemptedIVFPreviously == attemptedIVFPreviously && formula.ParamIsReasonForInfertilityKnown == reasonForInfertilityKnown {
			return formula, nil
		}
	}
	return FormulaParameters{}, errors.New("could not find formula")
}

func (formula FormulaParameters) GetTubalFactorValue(param bool) float64 {
	factor := formula.FormulaTubalFactorTrueValue
	if !param {
		factor = formula.FormulaTubalFactorFalseValue
	}
	return factor
}

func (formula FormulaParameters) GetMaleInfertilityFactorValue(param bool) float64 {
	factor := formula.FormulaMaleFactorInfertilityTrueValue
	if !param {
		factor = formula.FormulaMaleFactorInfertilityFalseValue
	}
	return factor
}

func (formula FormulaParameters) GetEndometriosisFactorValue(param bool) float64 {
	factor := formula.FormulaEndometriosisTrueValue
	if !param {
		factor = formula.FormulaEndometriosisFalseValue
	}
	return factor
}

func (formula FormulaParameters) GetOvulatorDisorderValue(param bool) float64 {
	factor := formula.FormulaOvulatoryDisorderTrueValue
	if !param {
		factor = formula.FormulaOvulatoryDisorderFalseValue
	}
	return factor
}

func (formula FormulaParameters) GetDiminishedOvarianResereValue(param bool) float64 {
	factor := formula.FormulaDiminishedOvarianReserveTrueValue
	if !param {
		factor = formula.FormulaDiminishedOvarianReserveFalseValue
	}
	return factor
}
func (formula FormulaParameters) GetUterineFactorValue(param bool) float64 {
	factor := formula.FormulaUterineFactorTrueValue
	if !param {
		factor = formula.FormulaUterineFactorFalseValue
	}
	return factor
}
func (formula FormulaParameters) GetOtherReasonValue(param bool) float64 {
	factor := formula.FormulaOtherReasonTrueValue
	if !param {
		factor = formula.FormulaOtherReasonFalseValue
	}
	return factor
}
func (formula FormulaParameters) GetUnexplainedInfertilityValue(param bool) float64 {
	factor := formula.FormulaUnexplainedInfertilityTrueValue
	if !param {
		factor = formula.FormulaUnexplainedInfertilityFalseValue
	}
	return factor
}
func (formula FormulaParameters) GetPriorPregnanciesValue(param int) float64 {
	var factor float64
	if param == 0 {
		factor = formula.FormulaPriorPregnancies0Value
	} else if param == 1 {
		factor = formula.FormulaPriorPregnancies1Value

	} else if param >= 2 {
		factor = formula.FormulaPriorPregnancies2PlusValue
	}
	return factor
}
func (formula FormulaParameters) GetLiveBirthsValue(param int) float64 {
	var factor float64
	if param == 0 {
		factor = formula.FormulaPriorLiveBirths0Value
	} else if param == 1 {
		factor = formula.FormulaPriorLiveBirths1Value

	} else if param >= 2 {
		factor = formula.FormulaPriorLiveBirths2PlusValue
	}
	return factor
}
