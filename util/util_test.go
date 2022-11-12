package util

import (
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestMaxA(t *testing.T) {
	if Max(2, 5) != 5 {
		t.FailNow()
	}
}

func TestMaxB(t *testing.T) {
	if Max(6, 5) != 6 {
		t.FailNow()
	}
}

func TestMaxC(t *testing.T) {
	if Max(-1, -1) != -1 {
		t.FailNow()
	}
}

func TestMinA(t *testing.T) {
	if Min(2, 5) != 2 {
		t.FailNow()
	}
}

func TestMinB(t *testing.T) {
	if Min(6, 5) != 5 {
		t.FailNow()
	}
}

func TestMinC(t *testing.T) {
	if Max(-1, -1) != -1 {
		t.FailNow()
	}
}

func TestContainsEmpty(t *testing.T) {
	if Contains(&[]int{}, 0) {
		t.FailNow()
	}
}

func TestContainsTrue(t *testing.T) {
	if !Contains(&[]int{1, 2, 4, 8}, 2) {
		t.FailNow()
	}
}

func TestContainsFalse(t *testing.T) {
	if Contains(&[]int{-1, 2, -4, 8}, 5) {
		t.FailNow()
	}
}

func TestReverse(t *testing.T) {
	s := []int{4, 3, 2, 1, 0}
	Reverse(&s)
	for e, i := range s {
		if e != i {
			t.Fatalf("%d found but %d expected in position #%d", e, i, i)
		}
	}
}

func TestIsAlphabeticEmpty(t *testing.T) {
	s := ""
	if !IsAlphabetic(&s) {
		t.FailNow()
	}
}

func TestIsAlphabeticTrue(t *testing.T) {
	s := "Hello"
	if !IsAlphabetic(&s) {
		t.FailNow()
	}
}

func TestIsAlphabeticFalse(t *testing.T) {
	s := "Hell o"
	if IsAlphabetic(&s) {
		t.FailNow()
	}
}

func TestTimeFormat(t *testing.T) {
	tm := time.Date(2022, time.September, 10, 18, 59, 00, 00, time.UTC)
	x := LastInstantAtGivenTime(tm, 20)
	if x != "2022-09-09T20:00:00Z" {
		t.Fatal("Error in formatting time #1: " + x)
	}
	tm = time.Date(2022, time.December, 01, 10, 19, 01, 00, time.UTC)
	x = LastInstantAtGivenTime(tm, 20)
	if x != "2022-11-30T20:00:00Z" {
		t.Fatal("Error in formatting time #2: " + x)
	}
}

func TestIdToObjectParsingError(t *testing.T) {
	ctx := GetTestingGinContext()
	ctx.AddParam("id", "zebra")
	_, err := IdToObject(ctx, map[int]int{0: 1, 1: 0})
	ctx.Params = gin.Params{}
	if err == nil {
		t.Fatal(err)
	}
}

func TestIdToObjectNotFound(t *testing.T) {
	ctx := GetTestingGinContext()
	ctx.AddParam("id", "2")
	_, err := IdToObject(ctx, map[int]int{0: 1, 1: 0})
	ctx.Params = gin.Params{}
	if err == nil {
		t.Fatal(err)
	}
}

func TestIdToObject(t *testing.T) {
	ctx := GetTestingGinContext()
	ctx.AddParam("id", "0")
	_, err := IdToObject(ctx, map[int]int{0: 1, 1: 0})
	if err != nil {
		t.Fatal(err)
	}
}

func TestInit(t *testing.T) {
	if gin.Mode() != gin.TestMode {
		t.Fatal("Gin is not in test mode")
	}
	if GetTestingResponseRecorder() != testingResponseRecorder {
		t.Fatal("Wrong GetTestingResponseRecorder")
	}
	if testingGinContext == nil {
		t.Fatal("testingGinContext is nil")
	}
	if GetTestingGinContext() != testingGinContext {
		t.Fatalf("Wrong GetTestingGinContext")
	}
	if testingGinEngine == nil {
		t.Fatal("testingGinEngine is nil")
	}
	if GetTestingGinEngine() != testingGinEngine {
		t.Fatalf("Wrong GetTestingGinEngine")
	}
}
