package main

import (
	"github.com/go-vgo/robotgo"
	"fmt"
)

func main() {
	robotgo.MouseClick("left", true)

	for i := 1; i < 1000; i ++ {
		x, y := robotgo.GetMousePos()
		fmt.Println("pos:", x, y)
		color := robotgo.GetPixelColor(x, y)
		fmt.Println("color----", color)
	}


	keve := robotgo.AddEvent("k")
	if keve == 0 {
		fmt.Println("you press...", "k")
	}
	//
	//mleft := robotgo.AddEvent("mleft")
	//if mleft == 0 {
	//	fmt.Println("you press...", "mouse left button")
	//}
	//
	//fpid, err := robotgo.FindIds("Google")
	//if err == nil {
	//	fmt.Println("pids...", fpid)
	//
	//	if len(fpid) > 0 {
	//		robotgo.ActivePID(fpid[0])
	//
	//		robotgo.Kill(fpid[0])
	//	}
	//}
	//
	//robotgo.ActiveName("chrome")
	//
	//isExist, err := robotgo.PidExists(100)
	//if err == nil && isExist {
	//	fmt.Println("pid exists is", isExist)
	//
	//	robotgo.Kill(100)
	//}
	//
	//abool := robotgo.ShowAlert("test", "robotgo")
	//if abool == 0 {
	//	fmt.Println("ok@@@", "ok")
	//}
	//
	//title := robotgo.GetTitle()
	//fmt.Println("title@@@", title)

	//
	//robotgo.ScrollMouse(1000, "up")
	////robotgo.MoveMouseSmooth(1100, 200, 1.0, 100.0)
	//
	//robotgo.TypeString("Hello World")
	//robotgo.TypeString("测试")
	//robotgo.TypeStr("测试")
	//ustr := uint32(robotgo.CharCodeAt("测试", 0))
	//robotgo.UnicodeType(ustr)
	////
	//robotgo.KeyTap("enter")
	//robotgo.TypeString("en")
	//robotgo.KeyTap("i", "alt", "command")
	//arr := []string{"alt", "command"}
	//robotgo.KeyTap("i", arr)
	//
	//robotgo.WriteAll("测试111")
	//text, err := robotgo.ReadAll()
	//if err == nil {
	//	fmt.Println(text)
	//}
	//

}
