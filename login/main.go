package main

import (
    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
)

func main() {
    var usernameTE, passwordTE *walk.LineEdit
    MainWindow{
        Title:   "登录",
        MinSize: Size{270, 300},
        Layout:  VBox{},
        Children: []Widget{
            Composite{
                // MaxSize: Size{0, 200},
                Layout: VBox{},
                // Layout: Grid{Columns: 2},
                Children: []Widget{
                    Label{ Text: "用户名" },
                    LineEdit{ AssignTo: &usernameTE, MaxLength: 10},
                    Label{ Text: "密码" },
                    LineEdit{ AssignTo: &passwordTE, ReadOnly: true},
                    PushButton{
                        Text:     "登录",
                        MinSize: Size{30, 50},
                        OnClicked: func() {
                            if usernameTE.Text() == "" {
                                var tmp walk.Form
                                walk.MsgBox(tmp, "提示", "用户名不能为空", walk.MsgBoxIconInformation)
                                return
                            }
                        },
					},
                },
            },
            // Composite{
            //     Layout:  HBox{},
            //     Children: []Widget{
            //         Label{ Text: "密码" },
            //         LineEdit{ AssignTo: &passwordTE},
            //     },
            // },
            // Composite{
            //     PushButton{
            //         Text:    "登录",
            //         MinSize: Size{120, 50},
            //         OnClicked: func() {
            //             if usernameTE.Text() == "" {
            //                 var tmp walk.Form
            //                 walk.MsgBox(tmp, "用户名为空", "", walk.MsgBoxIconInformation)
            //                 return
            //             }
            //             if passwordTE.Text() == "" {
            //                 var tmp walk.Form
            //                 walk.MsgBox(tmp, "密码为空", "", walk.MsgBoxIconInformation)
            //                 return
            //             }
            //         },
            //     },
            // },
        },
    }.Run()
}