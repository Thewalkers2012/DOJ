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

func GetAnswerMsg(code int) string {
	if code == 0 {
		return "Accept"
	} else if code == -1 {
		return "Wrong Answer"
	} else if code == 1 {
		return "Time Limit Exceeded"
	} else if code == 4 {
		return "RunTime Error"
	} else if code == 3 {
		return "Memory Limit Exceeded"
	} else {
		return "Compile Error"
	}
}

type JudgeResult struct {
	Time       int
	Memory     int
	ResultCode int
	Score      int
}
