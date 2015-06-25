// copyright : tencent
// author : solomonooo
// github : github.com/tencentyun/go-sdk

// this is a demo for qcloud go sdk
package main

import (
	"fmt"
	"github.com/tencentyun/go-sdk"
)

func main() {
	pic_test()
	video_test()	
}

func pic_test(){
	var appid uint = 200943
	sid := "AKIDOXkiS878nYFvc4sggDRxTU56UsmN3LMy"
	skey := "gMoR2lGvMWzxFGrxJCRoZMhU48f0tsdm"

	
	cloud := qcloud.PicCloud{appid, sid, skey}
	fmt.Println("=========================================")
	var analyze qcloud.PicAnalyze
	info, err := cloud.Upload("123456", "./test.jpg", analyze)
	if err != nil {
		fmt.Printf("pic upload failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic upload success")
		info.Print()
	}

	fmt.Println("******************")
	picInfo, err := cloud.Stat("123456", info.Fileid)
	if err != nil {
		fmt.Printf("pic stat failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic stat success")
		picInfo.Print()
	}

	fmt.Println("=========================================")
	info2, err := cloud.Copy("123456", info.Fileid)
	if err != nil {
		fmt.Printf("pic copy failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic copy success")
		info2.Print()
	}

	fmt.Println("******************")
	picInfo, err = cloud.Stat("123456", info2.Fileid)
	if err != nil {
		fmt.Printf("copy pic stat failed, err = %s\n", err.Error())
	} else {
		fmt.Println("copy pic stat success")
		picInfo.Print()
	}

	fmt.Println("=========================================")
	err = cloud.Download("123456", info2.Fileid, "./test2.jpg")
	if err != nil {
		fmt.Printf("pic download failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic download success")
	}

	fmt.Println("=========================================")
	err = cloud.Delete("123456", info.Fileid)
	if err != nil {
		fmt.Printf("pic delete failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic delete success")
	}

	err = cloud.Delete("123456", info2.Fileid)
	if err != nil {
		fmt.Printf("copy pic delete failed, err = %s\n", err.Error())
	} else {
		fmt.Println("copy pic delete success")
	}
	
	fmt.Println("=========================================")
	analyze.Fuzzy = 1;
	analyze.Food = 1;
	//is fuzzy? is food?
	info, err = cloud.Upload("123456", "./fuzzy.jpg", analyze)
	if err != nil {
		fmt.Printf("pic analyze failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic analyze success, pic=./fuzzy.jpg")
		fmt.Printf("is fuzzy : %d\r\n", info.Analyze.Fuzzy)
		fmt.Printf("is food : %d\r\n", info.Analyze.Food)
	}
	///////
	info, err = cloud.Upload("123456", "./food.jpg", analyze)
	if err != nil {
		fmt.Printf("pic analyze failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic analyze success, pic=./food.jpg")
		fmt.Printf("is fuzzy : %d\r\n", info.Analyze.Fuzzy)
		fmt.Printf("is food : %d\r\n", info.Analyze.Food)
	}
}

func video_test(){
	var appid uint = 200943
	sid := "AKIDOXkiS878nYFvc4sggDRxTU56UsmN3LMy"
	skey := "gMoR2lGvMWzxFGrxJCRoZMhU48f0tsdm"

	cloud := qcloud.VideoCloud{appid, sid, skey}
	fmt.Println("=========================================")
	info, err := cloud.Upload("123456", "./test.mp4","test_title","test_desc","test_magic_context")
	if err != nil {
		fmt.Printf("video upload failed, err = %s\n", err.Error())
	} else {
		fmt.Println("video upload success")
		info.Print()
	}

	fmt.Println("******************")
	picInfo, err := cloud.Stat("123456", info.Fileid)
	if err != nil {
		fmt.Printf("video stat failed, err = %s\n", err.Error())
	} else {
		fmt.Println("video stat success")
		picInfo.Print()
	}
	
	fmt.Println("=========================================")
	err = cloud.Download("123456", info.Fileid, "./test2.mp4")
	if err != nil {
		fmt.Printf("video download failed, err = %s\n", err.Error())
	} else {
		fmt.Println("video download success")
	}

	fmt.Println("******************")
	err = cloud.Delete("123456", info.Fileid)
	if err != nil {
		fmt.Printf("video delete failed, err = %s\n", err.Error())
	} else {
		fmt.Println("video delete success")
	}
}
