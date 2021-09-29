package judge

import "testing"

func TestClient_JudgeWithRequest(t *testing.T) {
	javaSrc := `
	import java.util.Scanner;
		public class Main{
			public static void main(String[] args){
			Scanner in=new Scanner(System.in);
			int a=in.nextInt();
			int b=in.nextInt();
			System.out.println(a + b);
		}
	}
	`
	cSrc := `
	#include <stdio.h>
	int main(){
		int a, b;
		scanf("%d%d", &a, &b);
		printf("%d\n", a+b);
		return 0;
	}
	`
	var tests = []struct {
		JudgeRequest *JudgeRequest
	}{
		{
			JudgeRequest: &JudgeRequest{
				Src:            javaSrc,
				LanguageConfig: JavaLangConfig,
				MaxCpuTime:     1000,
				MaxMemory:      256 * 1024 * 1024,
				TestCaseId:     "normal",
			},
		},
		{
			JudgeRequest: &JudgeRequest{
				Src:            cSrc,
				LanguageConfig: CPPLangConfig,
				MaxCpuTime:     1000,
				MaxMemory:      256 * 1024 * 1024,
				TestCaseId:     "normal",
			},
		},
	}

	for _, test := range tests {
		resp, err := client.JudgeWithRequest(test.JudgeRequest)
		if err != nil {
			t.Errorf("unexpected1 error. error: %+v", err)
			return
		}
		if resp.Err() != nil {
			t.Errorf("unexpected2 error. error: %+v", resp.Err())
			return
		}
		if resp.RespData == nil {
			t.Error("resp.RespData unexpected nil")
			return
		}
	}
}
