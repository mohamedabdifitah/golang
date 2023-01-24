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
	/* =========================== =====|| Print to console all variables ||================================ */
	fmt.Println(stringVar, floatNum, doubleNum, intVar, int64Var)
	// /* ================================|| integer conver to string  ||================================ */
	int2str := strconv.Itoa(int64Var)
	fmt.Printf("variable type %v", int2str)
	/* ================================|| Float to int ||================================ */
	floa2int := int(floatNum)
	fmt.Printf("variable type %v", floa2int)

}
