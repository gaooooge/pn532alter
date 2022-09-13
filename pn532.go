package main

import (
	"fmt"
	"log"

	"github.com/asjdf/pn532"
	"github.com/asjdf/pn532/command"
)

// 常用秘钥
var SIGN = []string{
	"bba5c1aab4b4",
	"ffffffffffff",
	"4038ac9e7ff5",
	"40389c9a77c5",
}

// 驱动所在目录
var DeviceDir = "/dev/tty.usbserial-110"

func main() {
	pn532.Mode = pn532.Release
	device, err := pn532.QuickInit(DeviceDir)
	if err != nil {
		fmt.Printf("加载设备失败: %v", err)
	}
	fmt.Println("初始化成功！")

	defer device.Close()

	_, err = device.FirmwareVersion()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("准备读取单张卡")
	uid, err := device.ReadPassiveTarget(pn532.ISO14443A)
	if err != nil {
		log.Fatalf("读取单张卡失败: %v", err)
	}
	log.Printf("读取单张卡成功 卡号: % X", uid)

	//按照扇区读卡
	// for i := 1; i < 64; i++ {
	// 	b := byte(i)

	// 	for _, k := range SIGN {
	// 		key := []byte{}
	// 		for j := 0; j < len(k); j += 2 {
	// 			n, _ := strconv.ParseInt(k[j:j+2], 16, 32)
	// 			key = append(key, byte(n))
	// 		}

	// 		//bba5c1aab4b4
	// 		// key := []byte{0xBB, 0xA5, 0xC1, 0xAA, 0xB4, 0xB4}
	// 		// key := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
	// 		fmt.Println(key)
	// 		success, err := device.MifareClassicAuthenticateBlock(uid, b, command.MifareCmdAuthA, key)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			continue
	// 		}

	// 		if success {
	// 			log.Print("authenticate success", "扇区", i, "key:", key)
	// 		} else {
	// 			continue
	// 			log.Print("authenticate failed", "扇区", i, "key:", key)
	// 		}

	// 		block, err := device.MifareClassicReadBlock(b)
	// 		if err != nil {
	// 			log.Fatal(err, b)
	// 		}
	// 		fmt.Printf("%d: %X \n", i, block)
	// 	}

	// }

	// b := byte(0)

	// //bba5c1aab4b4
	b := byte(12)
	key := []byte{0xBB, 0xA5, 0xC1, 0xAA, 0xB4, 0xB4}
	// key := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
	success, err := device.MifareClassicAuthenticateBlock(uid, b, command.MifareCmdAuthA, key)
	// if err != nil {
	// 	return
	// }
	if success {
		log.Print("authenticate success")
	} else {
		log.Print("authenticate failed")
	}

	block, err := device.MifareClassicReadBlock(b)
	if err != nil {
		log.Fatal(err, b)
	}
	fmt.Printf("0: % X \n", block)

	//00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
	//03 22 06 18 22 09 10 7F FF FF FF FF 03 06 03 03
	// testBlock := []byte{0x9A, 0x9A, 0x77, 0xC1, 0xB6, 0x08, 0x04, 0x00, 0x02, 0xAE, 0xAD, 0xC0, 0xE3, 0x57, 0xC6, 0x1D}
	// testBlock := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	// testBlock := []byte{0x03, 0x22, 0x06, 0x18, 0x24, 0x09, 0x10, 0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0x03, 0x06, 0x03, 0x03}
	testBlock := []byte{0x03, 0x20, 0x09, 0x08, 0x26, 0x09, 0x10, 0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0x03, 0x06, 0x03, 0x03}

	success, err = device.MifareClassicWriteBlock(b, testBlock)
	if err != nil {
		log.Fatal(err)
	} else if !success {
		log.Fatal("write block failed")
	}
	// fmt.Println(success)
	// for i := 0; i < 64; i++ {
	// 	// fmt.Println(reflect.TypeOf(byte(i)))
	// 	// fmt.Println(byte(i))
	// 	block, err := device.MifareClassicReadBlock(byte(i))
	// 	if err != nil {
	// 		log.Fatal(err, byte(i))
	// 	}
	// 	fmt.Printf("block: % X", block)
	// }
}
