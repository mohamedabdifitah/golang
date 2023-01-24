package variables

import (
	"fmt"
	"strconv"
)

func Variables() {
	/* =========================== =====|| Variables initialize||================================ */
	var stringVar string = "string variable"
	var floatNum float32 = 46.6
	var doubleNum float64 = 56.7
	var intVar int = 444
	var int64Var int = 5
	var unsignedVar uint = 7
	var boolVar bool = true
	var complexVar complex64 = 65 + 2i
	var runeVar rune = 'h'
	/* =========================== =====|| Print to console all variables ||================================ */
	fmt.Println(stringVar, floatNum, doubleNum, intVar, int64Var, unsignedVar, boolVar, complexVar, runeVar)
	// /* ================================|| integer conver to string  ||================================ */
	int2str := strconv.Itoa(int64Var)
	fmt.Printf("variable type %v", int2str)
	/* ================================|| Float to int ||================================ */
	floa2int := int(floatNum)
	fmt.Printf("variable type %v", floa2int)

}
