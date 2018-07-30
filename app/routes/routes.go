// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).URL
}

func (_ tApp) DbSearch(
		word string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "word", word)
	return revel.MainRouter.Reverse("App.DbSearch", args).URL
}

func (_ tApp) IsbnRegister(
		isbn string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "isbn", isbn)
	return revel.MainRouter.Reverse("App.IsbnRegister", args).URL
}

func (_ tApp) TitleRegister(
		title string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "title", title)
	return revel.MainRouter.Reverse("App.TitleRegister", args).URL
}

func (_ tApp) RegisterPost(
		title string,
		author string,
		date string,
		isbn string,
		image string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "title", title)
	revel.Unbind(args, "author", author)
	revel.Unbind(args, "date", date)
	revel.Unbind(args, "isbn", isbn)
	revel.Unbind(args, "image", image)
	return revel.MainRouter.Reverse("App.RegisterPost", args).URL
}

func (_ tApp) DeletePost(
		isbn string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "isbn", isbn)
	return revel.MainRouter.Reverse("App.DeletePost", args).URL
}

func (_ tApp) Display(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Display", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeDir(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeDir", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}

func (_ tStatic) ServeModuleDir(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModuleDir", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


