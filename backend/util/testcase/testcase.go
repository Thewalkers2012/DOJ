package testcase

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/settings"
	"go.uber.org/zap"
)

func CreateTestcase(args *model.Info, pid int64) error {
	path := settings.Config.TestPath
	err := os.Chdir(path)
	if err != nil {
		zap.L().Error("os.Chdir(path) failed", zap.Error(err))
		return err
	}

	pidStr := strconv.Itoa(int(pid))
	err = os.Mkdir(pidStr, 0777)
	if err != nil {
		zap.L().Error("os.Mkdir() failed", zap.Error(err))
		return err
	}

	err = os.Chdir(pidStr)
	if err != nil {
		zap.L().Error("os.Chdir() failed", zap.Error(err))
		return err
	}

	// 将所有的输入输出全部写各个样例
	err = CreateTestFile(args.TestCaseNum, args.TestCases, path+"/"+pidStr)

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
		fmt.Println(testcase[i-1].InputName)
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
	}
	return nil
}

// func GenerateInfoFile() {

// }
