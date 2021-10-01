package judge

const (
	postUrl = "http://127.0.0.1:12358"
	token   = "YOUR_TOKEN_HERE"
)

func SubmitCode(code string, language string, maxCpuTime, maxMemory int64, testCaseID string) *Resp {
	client := NewClient(WithTimeout(0))
	client.SetOptions(WithEndpointURL(postUrl), WithToken(token))

	resp, _ := client.JudgeWithRequest(&JudgeRequest{
		Src:            code,
		LanguageConfig: getConfig(language),
		MaxCpuTime:     maxCpuTime,
		MaxMemory:      maxMemory,
		TestCaseId:     testCaseID,
	})

	return resp
}

func getConfig(language string) *LangConfig {
	if language == "cpp" {
		return CPPLangConfig
	} else if language == "java" {
		return JavaLangConfig
	}
	return PY2LangConfig
}
