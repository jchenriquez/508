package twosum

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"testing"
)

type Test struct {
	Nums   []int `json:"nums"`
	Target int   `json:"target"`
	Output []int `json:"output"`
}

func TestTwoSum(testUtil *testing.T) {
	f, err := os.Open("./tests.json")

	if err != nil {
		testUtil.Error(err)
		return
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	decoder := json.NewDecoder(reader)

	var tests map[string]Test

	for {
		err = decoder.Decode(&tests)

		if err == nil {
			for name, test := range tests {
				testUtil.Run(name, func(stUtils *testing.T) {
					result := TwoSum(test.Nums, test.Target)
					sort.Ints(result)
					sort.Ints(test.Output)

					if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", test.Output) {
						stUtils.Errorf("result that errored %v\n", result)
					}
				})
			}
		} else if err == io.EOF {
			break
		} else {
			testUtil.Error(err)
			return
		}
	}
}
