package testcase

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/settings"
	"go.uber.org/zap"
)

func CreateTestcase(args *model.Info, pid int64) error {
	path := settings.Config.TestPath
	// 1. 到 testcase 所在的目录
	err := os.Chdir(path)
	if err != nil {
		zap.L().Error("os.Chdir(path) failed", zap.Error(err))
		return err
	}

	// 2. 创建和 problemid 相同的文件夹
	pidStr := strconv.Itoa(int(pid))
	err = os.Mkdir(pidStr, 0777)
	if err != nil {
		zap.L().Error("os.Mkdir() failed", zap.Error(err))
		return err
	}

	// 3. 进去到那个题目的文件夹之中
	err = os.Chdir(pidStr)
	if err != nil {
		zap.L().Error("os.Chdir() failed", zap.Error(err))
		return err
	}

	// 4. 将所有的输入输出全部写各个样例
	err = CreateTestFile(args.TestCaseNum, args.TestCases, path+"/"+pidStr)
	if err != nil {
		return nil
	}

	// 生成 info 文件
	err = GenerateInfoFile(args)

	return err
}

// 将内容写到文件里面
func writeToFile(content string, file *os.File) error {
	defer file.Close()
	_, err := io.WriteString(file, content)
	return err
}

// 获取文件的大小
func getFileSize(filename string) (int64, error) {
	file, err := os.Stat(filename)
	if err != nil {
		zap.L().Error("getFileSize failed", zap.Error(err))
		return 0, err
	}
	return file.Size(), nil
}

func CreateTestFile(num int, testcase []model.TestCase, base string) error {
	for i := 1; i <= num; i++ {
		// 创建 xx.in 文件
		testcase[i-1].InputName = fmt.Sprintf("%d.in", i)
		inputFile, err := os.Create(testcase[i-1].InputName)
		if err != nil {
			zap.L().Error("create input file failed", zap.Error(err))
			return err
		}
		// 将 testcase[i].Input 写入 xx.in 文件
		err = writeToFile(testcase[i-1].Input, inputFile)
		if err != nil {
			zap.L().Error("(input) writeToFile failed", zap.Error(err))
			return err
		}
		// 更新 inputSize 大小
		inputSize, err := getFileSize(base + "/" + testcase[i-1].InputName)
		if err != nil {
			zap.L().Error("getFileSize failed", zap.Error(err))
			return err
		}
		testcase[i-1].InputSize = inputSize
		// 创建 xx.out 文件
		testcase[i-1].OutputName = fmt.Sprintf("%d.out", i)
		outputFile, err := os.Create(testcase[i-1].OutputName)
		if err != nil {
			zap.L().Error("create output file failed", zap.Error(err))
			return err
		}
		// 将 testcase[i].Output 写入 xx.out 文件
		err = writeToFile(testcase[i-1].Output, outputFile)
		if err != nil {
			zap.L().Error("(output) writeToFile failed", zap.Error(err))
			return err
		}
		// 更新 outputSize 大小
		outputSize, err := getFileSize(base + "/" + testcase[i-1].OutputName)
		if err != nil {
			zap.L().Error("getFileSize failed", zap.Error(err))
			return err
		}
		testcase[i-1].OutputSize = outputSize
		// todo：更新 output Md5 的值
		testcase[i-1].StrippedOutputMd5 = generateMd5(testcase[i-1].Output)
		testcase[i-1].OutputMd5 = generateMd5(testcase[i-1].Output)
	}

	return nil
}

func generateMd5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GenerateInfoFile(args *model.Info) error {
	res := &strings.Builder{}
	res.WriteByte('{')
	res.WriteString(fmt.Sprintf("\"test_case_number\": %d,", args.TestCaseNum))
	res.WriteString("\"spj\": false,")
	res.WriteString(GenerateTestCase(args.TestCases))
	res.WriteByte('}')

	_, err := os.Create("info")
	if err != nil {
		zap.L().Error("create info file failed", zap.Error(err))
		return err
	}

	err = ioutil.WriteFile("info", []byte(res.String()), 0777)
	return err
}

func GenerateTestCase(params []model.TestCase) string {
	n := len(params)
	res := &strings.Builder{}
	res.WriteString("\"test_cases\":")
	res.WriteByte('{')
	for i := 1; i <= n; i++ {
		str := fmt.Sprintf("\"%d\": {\"stripped_output_md5\": \"%s\", \"output_size\": %d, \"output_md5\": \"%s\", \"input_name\": \"%s\", \"input_size\": %d, \"output_name\": \"%s\"}",
			i,
			params[i-1].StrippedOutputMd5,
			params[i-1].OutputSize,
			params[i-1].OutputMd5,
			params[i-1].InputName,
			params[i-1].InputSize,
			params[i-1].OutputName,
		)
		res.WriteString(str)
		if i != n {
			res.WriteByte(',')
		}
	}
	res.WriteByte('}')

	return res.String()
}
