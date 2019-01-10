// package main
// import (   
//     . "github.com/lxn/walk/declarative"
// )
// func main() {
//     mainWindow.Run()
// }
// var LableHello=Label{
//   Text: "Hello world!",
// }
// var widget=[]Widget{
//    LableHello,
// }
// var mainWindow=MainWindow{
//   Title:"MainWindow",
//   MinSize:Size{400, 200},
//   Layout:VBox{},
//   Children:widget,
// }

// func main() {
//     var label1, label2 *walk.Label
//     mv := MainWindow{
//         Title:  "go window",
//         Size:   Size{Width:800, Height:600},
//         Layout: VBox{MarginsZero: true},
//         Children: []Widget{
//             Label{Text: "Hello World", AssignTo: &label1,
//                 Background: SolidColorBrush{walk.RGB(255, 0, 0)},
//                 OnSizeChanged: func() {
//                         label1.SetBounds(walk.Rectangle{100, 100, 100, 100})
//             }},
//             Label{Text: "Hello World", AssignTo: &label2,
//                 Background: SolidColorBrush{walk.RGB(0, 255, 0)},
//                 OnSizeChanged: func() {
//                     label2.SetBounds(walk.Rectangle{100, 200, 100, 100})
//             }},
//         },
//     }
//     mv.Run()
// }


package main

import (
    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
    "strings"
)

func main() {
    var inTE, outTE *walk.TextEdit

    MainWindow{
        Title:   "SCREAMO",
        MinSize: Size{600, 400},
        Layout:  VBox{},
        Children: []Widget{
            HSplitter{
                Children: []Widget{
                    TextEdit{AssignTo: &inTE},
                    TextEdit{AssignTo: &outTE, ReadOnly: true},
                },
            },
            PushButton{
                Text: "SCREAM",
                OnClicked: func() {
                    outTE.SetText(strings.ToUpper(inTE.Text()))
                },
            },
        },
    }.Run()
}
