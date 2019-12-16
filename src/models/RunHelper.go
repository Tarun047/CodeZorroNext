package models

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func RunSpecificCode(code *Code){
	if code.Language == "C" {
		exec.Command("gcc", "sampleC.c", "-o", "sampleC")
		cmd := exec.Command("./sampleC")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%q\n", out.String())
	}
	if code.Language == "CPP" {
		exec.Command("g++", "sampleCPP.cpp", "-o", "sampleCPP")
		cmd := exec.Command("./sampleCPP")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%q\n", out.String())
	}
	if code.Language == "JAVA" {
		exec.Command("javac", "sampleJAVA.java")
		cmd := exec.Command("java","sampleJAVA")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%q\n", out.String())
	}
	if code.Language == "PYTH" {
		cmd := exec.Command("python3","samplePyth.py")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%q\n", out.String())
	}
}
