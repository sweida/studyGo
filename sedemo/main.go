package main
import (
    "log"  
)
import (
    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
)
var le *walk.LineEdit
var sport,maths,english *walk.RadioButton
func main() {
    if _, err := MainWindow1.Run(); err != nil {
        log.Fatal(err)
    }
    log.Fatal(Bind("enabledCB.Checked"))
}
var MainWindow1=MainWindow{
  Title:   "MainWindow",
  MinSize: Size{300, 200},
  Layout:  VBox{},  
  Children: widget,
}
 
var widget=[]Widget{
  LineEdit1,
  RadioButtonSport,
  RadioButtonMaths,
  RadioButtonEnglish,
  PushButtonOK,
}
 
var LineEdit1=LineEdit{ 
  AssignTo: &le,
}
 
var RadioButtonSport=RadioButton{
  AssignTo: &sport,
  Text:    "体育系", 
}
var RadioButtonMaths=RadioButton{
  AssignTo: &maths,
  Text:    "数学系",
}
var RadioButtonEnglish=RadioButton{
  AssignTo: &english,
  Text:    "英语系",
}
var PushButtonOK=PushButton{
  Text: "OK",
  OnClicked:OK_Clicked,
}
func OK_Clicked(){ 
  switch {
    case english.Checked():
      le.SetText("英语系")
    case sport.Checked():
      le.SetText("体育系")
    case maths.Checked():
      le.SetText("数学系")
  }
}