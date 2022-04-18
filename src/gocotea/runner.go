package gocotea

import (
	"fmt"
	"os"

	gopython "github.com/ispras/gopython/src/gopython"
)

type Runner struct {
	playbookPath       string
	inventoryPath      string
	classNamePy        string
	argMaker           *ArgumentMaker
	runnerPythonObject *gopython.PythonObject
	emptyArgs          gopython.PythonMethodArguments
}

func (r *Runner) InitRunner(argmaker *ArgumentMaker, pbPath, debugMode, logFile string) error {
	moduleNamePy := "cotea.runner"
	r.playbookPath = pbPath
	r.argMaker = argmaker
	r.classNamePy = "runner"

	var runnerModule gopython.PythonModule
	runnerModule.SetModuleName(moduleNamePy)
	err := runnerModule.MakeImport()
	if err != nil {
		return &PythonImportError{ModuleName: moduleNamePy,
			ErrorMsg: err.Error()}
	}

	var runnerClass *gopython.PythonClass
	runnerClass, err = runnerModule.GetClass(r.classNamePy)
	if err != nil {
		return &PythonAttrError{SourceUnitName: moduleNamePy,
			AttrName: r.classNamePy,
			ErrorMsg: err.Error()}
	}

	var initArgs gopython.PythonMethodArguments
	initArgsCount := 2

	initArgs.SetArgCount(initArgsCount)
	initArgs.SetNextArgument(pbPath)
	initArgs.SetNextArgument(argmaker.argMakerPythonObject)
	//initArgs.SetNextArgument(debugMode)
	//initArgs.SetNextArgument(logFile)

	r.runnerPythonObject, err = runnerClass.CreateObject(&initArgs)
	if err != nil {
		return &PythonObjectCreationError{ClassName: r.classNamePy,
			ErrorMsg: err.Error()}
	}

	fmt.Println(r.runnerPythonObject)

	// python runner object has a lot of methods who takes 0 arguments
	// we're creating it on time here
	r.emptyArgs.SetArgCount(0)

	return nil
}

func (r *Runner) HasNextPlay() bool {
	methodNamePy := "has_next_play"
	resObjects, err := r.runnerPythonObject.CallMethod(methodNamePy, &r.emptyArgs)
	if err != nil {
		gotErr := PythonCallMethodError{MethodName: methodNamePy,
			ClassName: r.classNamePy, ErrorMsg: err.Error()}
		fmt.Println(gotErr.Error())
		os.Exit(1)
	}

	hasNextPlayPy := resObjects[0]
	res, typeError := hasNextPlayPy.ToStandartGoType()
	if typeError != nil {
		fmt.Println(typeError)
		os.Exit(1)
	}
	//fmt.Println("HasNextPlay result = ", res.(bool))
	return res.(bool)
}

// func (r *Runner) SetupPlayForRun() bool {
// 	methodNamePy := "setup_play_for_run"
// 	resObjects, err := r.runnerPythonObject.CallMethod(methodNamePy, &r.emptyArgs)
// 	if err != nil {
// 		gotErr := PythonCallMethodError{MethodName: methodNamePy,
// 			ClassName: r.classNamePy, ErrorMsg: err.Error()}
// 		fmt.Println(gotErr.Error())
// 		os.Exit(1)
// 	}

// 	setupOkPy := resObjects[0]
// 	res, typeError := setupOkPy.ToStandartGoType()
// 	if typeError != nil {
// 		fmt.Println(typeError)
// 		os.Exit(1)
// 	}

// 	return res.(bool)
// }

func (r *Runner) HasNextTask() bool {
	methodNamePy := "has_next_task"
	resObjects, err := r.runnerPythonObject.CallMethod(methodNamePy, &r.emptyArgs)
	if err != nil {
		gotErr := PythonCallMethodError{MethodName: methodNamePy,
			ClassName: r.classNamePy, ErrorMsg: err.Error()}
		fmt.Println(gotErr.Error())
		os.Exit(1)
	}

	hasNextTaskPy := resObjects[0]
	hasNextTaskGo, typeError := hasNextTaskPy.ToStandartGoType()
	if typeError != nil {
		fmt.Println(typeError)
		os.Exit(1)
	}

	var res bool
	switch v := hasNextTaskGo.(type) {
	case int:
		res = true
	case bool:
		res = v
	}
	// default and handle posible error ???

	return res
}

func (r *Runner) RunNextTask() bool {
	methodNamePy := "run_next_task"
	_, err := r.runnerPythonObject.CallMethod(methodNamePy, &r.emptyArgs)
	if err != nil {
		gotErr := PythonCallMethodError{MethodName: methodNamePy,
			ClassName: r.classNamePy, ErrorMsg: err.Error()}
		fmt.Println(gotErr.Error())
		os.Exit(1)
	}

	// runNextPy := resObjects[0]
	// res, typeError := runNextPy.ToStandartGoType()
	// if typeError != nil {
	// 	fmt.Println(typeError)
	// 	os.Exit(1)
	// }

	return true
}

func (r *Runner) FinishAnsibleWork() bool {
	methodNamePy := "finish_ansible"
	_, err := r.runnerPythonObject.CallMethod(methodNamePy, &r.emptyArgs)
	if err != nil {
		gotErr := PythonCallMethodError{MethodName: methodNamePy,
			ClassName: r.classNamePy, ErrorMsg: err.Error()}
		fmt.Println(gotErr.Error())
		os.Exit(1)
	}

	return true
}

func (r *Runner) GetIP() string {
	methodNamePy := "_getIP"
	resObjects, err := r.runnerPythonObject.CallMethod(methodNamePy, &r.emptyArgs)
	if err != nil {
		gotErr := PythonCallMethodError{MethodName: methodNamePy,
			ClassName: r.classNamePy, ErrorMsg: err.Error()}
		fmt.Println(gotErr.Error())
		os.Exit(1)
	}

	runNextPy := resObjects[0]
	res, typeError := runNextPy.ToStandartGoType()
	if typeError != nil {
		fmt.Println(typeError)
		os.Exit(1)
	}

	return res.(string)
}

func (r *Runner) WasError() bool {
	methodNamePy := "was_error"
	resObjects, err := r.runnerPythonObject.CallMethod(methodNamePy, &r.emptyArgs)
	if err != nil {
		gotErr := PythonCallMethodError{MethodName: methodNamePy,
			ClassName: r.classNamePy, ErrorMsg: err.Error()}
		fmt.Println(gotErr.Error())
		os.Exit(1)
	}

	runNextPy := resObjects[0]
	res, typeError := runNextPy.ToStandartGoType()
	if typeError != nil {
		fmt.Println(typeError)
		os.Exit(1)
	}

	return res.(bool)
}

func (r *Runner) GetErrorMsg() string {
	methodNamePy := "get_error_msg"
	resObjects, err := r.runnerPythonObject.CallMethod(methodNamePy, &r.emptyArgs)
	if err != nil {
		gotErr := PythonCallMethodError{MethodName: methodNamePy,
			ClassName: r.classNamePy, ErrorMsg: err.Error()}
		fmt.Println(gotErr.Error())
		os.Exit(1)
	}

	fmt.Println(resObjects[0])
	runNextPy := resObjects[0]
	res, typeError := runNextPy.ToStandartGoType()
	if typeError != nil {
		fmt.Println(typeError)
		os.Exit(1)
	}

	return res.(string)
}

func (r *Runner) GetAnsibleRunString() string {
	methodNamePy := "get_run_string"
	resObjects, err := r.runnerPythonObject.CallMethod(methodNamePy, &r.emptyArgs)
	if err != nil {
		gotErr := PythonCallMethodError{MethodName: methodNamePy,
			ClassName: r.classNamePy, ErrorMsg: err.Error()}
		fmt.Println(gotErr.Error())
		os.Exit(1)
	}

	runNextPy := resObjects[0]
	res, typeError := runNextPy.ToStandartGoType()
	if typeError != nil {
		fmt.Println(typeError)
		os.Exit(1)
	}

	return res.(string)
}

func (r *Runner) GetNextTaskName() string {
	methodNamePy := "get_next_task_name"
	resObjects, err := r.runnerPythonObject.CallMethod(methodNamePy, &r.emptyArgs)
	if err != nil {
		gotErr := PythonCallMethodError{MethodName: methodNamePy,
			ClassName: r.classNamePy, ErrorMsg: err.Error()}
		fmt.Println(gotErr.Error())
		os.Exit(1)
	}

	runNextPy := resObjects[0]
	res, typeError := runNextPy.ToStandartGoType()
	if typeError != nil {
		fmt.Println(typeError)
		os.Exit(1)
	}

	return res.(string)
}

func (r *Runner) GetCurrentPlayName() string {
	methodNamePy := "get_cur_play_name"
	resObjects, err := r.runnerPythonObject.CallMethod(methodNamePy, &r.emptyArgs)
	if err != nil {
		gotErr := PythonCallMethodError{MethodName: methodNamePy,
			ClassName: r.classNamePy, ErrorMsg: err.Error()}
		fmt.Println(gotErr.Error())
		os.Exit(1)
	}

	runNextPy := resObjects[0]
	res, typeError := runNextPy.ToStandartGoType()
	if typeError != nil {
		fmt.Println(typeError)
		os.Exit(1)
	}

	return res.(string)
}
