package models

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"
)

func RunSpecificCode(code *Code,fname string) (string,bool){
	var resp string
	retCode:=true
	if code.Language == "C" || code.Language=="CPP" {
		var cFamily string
		if code.Language=="C"{
			cFamily="gcc"
		} else {
			cFamily = "g++"
		}
		cmd1:=exec.Command(cFamily, fname, "-o", "sampleC")
		_,err1:=cmd1.Output()
		if err1 != nil{
			resp = err1.Error()
			retCode=false
		} else {
			cmd := exec.Command("./sampleC")
			var out bytes.Buffer
			cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
				resp=err.Error()
				retCode=false
			} else {
				resp = out.String()
			}
		}
	}
	if code.Language == "JAVA" {
		cmd1:=exec.Command("javac", fname)
		_,err1:=cmd1.Output()
		if err1 != nil{
			resp = err1.Error()
			retCode=false
		}else {

			cmd := exec.Command("java", strings.TrimSuffix(fname, filepath.Ext(fname)))
			var out bytes.Buffer
			cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
				retCode = false
				resp=err.Error()
			}
			resp = out.String()
		}
	}
	if code.Language == "PYTH" {
		cmd := exec.Command("python3",fname)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			resp=err.Error()
			retCode = false
		}
		resp=out.String()
	}
	return resp,retCode
}
