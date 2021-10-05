package model

type Info struct {
	TestCaseNum int  `json:"test_case_number"` // 测试样例的数目
	Spj         bool `json:"spj"`
	TestCases   []TestCase
}

type TestCase struct {
	InputSize         int64  `json:"input_size"`
	OutputSize        int64  `json:"output_size"`
	InputName         string `json:"input_name"`
	OutputName        string `json:"output_name"`
	Input             string
	Output            string
	OutputMd5         string `json:"output_md5"`
	StrippedOutputMd5 string `json:"stripped_output_md5"`
}
