package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"

	"com.sunfish.ivfsuccesscalculator/models"
	"github.com/gin-gonic/gin"
)

type IVFController struct {
	formulas models.Formulas
}

func NewIVFController(formulas models.Formulas) *IVFController {
	return &IVFController{
		formulas: formulas,
	}
}
func (c *IVFController) Home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}

func (c *IVFController) CalculateSuccess(ctx *gin.Context) {
	var err error
	var formData models.RequestInput
	if err := ctx.ShouldBind(&formData); err != nil {
		log.Println(err)
		ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": "Invalid form data",
		})
		return
	}
	var ivfAttemptedPrev string
	switch formData.IvfUsed {
	case "0":
		ivfAttemptedPrev = "FALSE"
	case "1", "2", "3+":
		ivfAttemptedPrev = "TRUE"
	default:
		ivfAttemptedPrev = "N/A"
	}

	fi := models.FormulaInput{
		ParamUsingOwnEggs:                formData.EggSource == "Own",
		ParamAttemptedIVFPreviously:      ivfAttemptedPrev,
		ParamIsReasonForInfertilityKnown: !formData.NoIvfReason(),
	}
	log.Println(fi)

	bt, _ := json.MarshalIndent(formData, "", " ")
	log.Println(string(bt))
	sFormula, err := c.formulas.ChooseFormula(fi.ParamUsingOwnEggs, fi.ParamAttemptedIVFPreviously, fi.ParamIsReasonForInfertilityKnown)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	ui := formData.ConvertToUserInput()
	err = ui.Validate()
	if err != nil {
		log.Println(err)
		ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": fmt.Sprintf("validation: %s", err.Error()),
		})
		return
	}
	score, successRate := ui.SuccessRate(sFormula)
	log.Println("Score:", score, "Success Rate:", successRate)

	ctx.HTML(http.StatusOK, "result.html", gin.H{
		"successRate": roundFloat(successRate, 2),
		"score":       score,
	})
}
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
