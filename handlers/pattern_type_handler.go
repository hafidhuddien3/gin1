package handlers

import (
    "fmt"
	"net/http"
    // "gin-quickstart/db"
    // "gin-quickstart/models"
    "github.com/gin-gonic/gin"
)

// nextInSequence determines if the sequence is arithmetic or geometric
// and returns the next number.
func nextInSequenceOld(seq []int) (string, int) {
    if len(seq) < 2 {
        return "unknown", 0
    }

    // Check arithmetic
    diff := seq[1] - seq[0]
    isArithmetic := true
    for i := 1; i < len(seq); i++ {
        if seq[i]-seq[i-1] != diff {
            isArithmetic = false
            break
        }
    }

    if isArithmetic {
        return "arithmetic", seq[len(seq)-1] + diff
    }

    // Check geometric
    if seq[0] == 0 {
        return "unknown", 0
    }
    ratio := seq[1] / seq[0]
    isGeometric := true
    for i := 1; i < len(seq); i++ {
        if seq[i-1] == 0 || seq[i]/seq[i-1] != ratio {
            isGeometric = false
            break
        }
    }

    if isGeometric {
        return "geometric", seq[len(seq)-1] * ratio
    }

    return "unknown", 0
}

type RuleFunc func([]int) (string, int, bool)
func nextInSequence(seq []int, rules ...RuleFunc) (string, int) {
    if len(seq) < 2 {
        return "unknown", 0
    }

    // Arithmetic
    diff := seq[1] - seq[0]
    isArithmetic := true
    for i := 1; i < len(seq); i++ {
        if seq[i]-seq[i-1] != diff {
            isArithmetic = false
            break
        }
    }
    if isArithmetic {
        return "arithmetic", seq[len(seq)-1] + diff
    }

    // Geometric
    if seq[0] != 0 {
        ratio := float64(seq[1]) / float64(seq[0])
		fmt.Println("ratio", ratio)
        isGeometric := true
        for i := 1; i < len(seq); i++ {
			// fmt.Println("ratioFor", seq[i]/seq[i-1])
            if seq[i-1] == 0 || float64(seq[i])/float64(seq[i-1]) != ratio {
                isGeometric = false
                break
            }
        }
        if isGeometric {
			next := int(float64(seq[len(seq)-1]) * ratio)
            return "geometric", next
        }
    }

    // Fibonacci
    if len(seq) >= 3 {
        isFibonacci := true
        for i := 2; i < len(seq); i++ {
            if seq[i] != seq[i-1]+seq[i-2] {
                isFibonacci = false
                break
            }
        }
        if isFibonacci {
            return "fibonacci", seq[len(seq)-1] + seq[len(seq)-2]
        }
    }

	// custom rules 
	for _, rule := range rules { if name, next, ok := rule(seq); ok { return name, next } }

    return "unknown", 0
}



// data["seq"] → map access (dynamic, key lookup).
// data.Seq → struct field access (static, strongly typed).

	type InputPattern struct {
		Seq []int `json:"seq"`
	}

// @Tags Pattern 
// @Summary Post pattern 
// @Description Detect the type of sequence and return the next number 
// @Accept json 
// @Produce json 
// @Param input body InputPattern true "Sequence input" 
// @Success 200 {object} map[string]interface{} 
// @Failure 400 {object} map[string]string 
// @Router /pattern [post]
func SearchPattern(c *gin.Context) {
	// example input with map
	// var input map[string]interface{} 
	// if err := c.ShouldBindJSON(&input); err != nil { c.JSON(400, gin.H{"error": err.Error()}) return } 
	// input["seq"] will be []interface{}

	var inputPattern InputPattern
    if err := c.ShouldBindJSON(&inputPattern); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	// fmt.Println(inputPattern["seq"])

	
	// Custom rule: double then add 1
	customRuleDoubleAddOne := func(seq []int) (string, int, bool) {
		if len(seq) < 2 {
			return "", 0, false
		}
		// check if each term = prev*2 + 1
		for i := 1; i < len(seq); i++ {
			if seq[i] != seq[i-1]*2+1 {
				return "", 0, false
			}
		}
		return "double+1", seq[len(seq)-1]*2 + 1, true
	}


	// pattern, next := nextInSequence([]int{3, 7, 15}, customRule)
	// → "double+1", 31

	pattern, next := nextInSequence(inputPattern.Seq, customRuleDoubleAddOne)
    c.JSON(http.StatusOK, gin.H{  
		"pattern": pattern, 
		"next": next,
	 })
}

// func main() {
//     // Example 1: geometric
//     seq1 := []int{2, 4, 8}
//     pattern1, next1 := nextInSequence(seq1)
//     fmt.Printf("Sequence: %v → %s, next = %d\n", seq1, pattern1, next1)

//     // Example 2: arithmetic
//     seq2 := []int{2, 4, 6}
//     pattern2, next2 := nextInSequence(seq2)
//     fmt.Printf("Sequence: %v → %s, next = %d\n", seq2, pattern2, next2)
// }
