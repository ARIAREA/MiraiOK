package main

/*
 * Copyright 2020 LXY1226, Mamoe Technologies and contributors.
 *
 * 采用与mirai相同的LICENSE
 * 此源代码的使用受 GNU AFFERO GENERAL PUBLIC LICENSE version 3 许可证的约束, 可以在以下链接找到该许可证.
 * Use of this source code is governed by the GNU AGPLv3 license that can be found through the following link.
 *
 * https://github.com/mamoe/mirai/blob/master/LICENSE
 */

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

const (
	mainClass = "net.mamoe.mirai.console.pure.MiraiConsolePureLoader"
	libDIR    = "libs/"
	RTStr     = runtime.GOOS + "-" + runtime.GOARCH
)

var javaPath = "./jre/bin/java"
var arg0 = os.Args[0]
var libs []lib
var globalWG = sync.WaitGroup{}

func main() {
	INFO("MiraiOK", BUILDTIME, RTStr)
	INFO("此程序以Affero GPL3.0协议发布，使用时请遵守协议")
	INFO("部分开源于: github.com/LXY1226/MiraiOK")
	doUpdate()
	globalWG.Wait()
	classpath := "CLASSPATH="
	for _, lib := range libs {
		classpath += lib.LibPath() + ";"
	}
	cmd := exec.Command(javaPath, mainClass)
	cmd.Env = append(cmd.Env, classpath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	INFO("启动Mirai...")
	err := os.Remove(arg0 + ".old")
	if err == nil {
		INFO("删除旧版MiraiOK")
	}
	err = cmd.Run()
	if err != nil {
		log.Print("java退出，", err)
		ERROR("运行失败，尝试更新mirai三件套", err.Error())
		updateMirai(true)
		WARN("请重新启动MiraiOK")
		var str string
		INFO("按任意键退出")
		fmt.Scan(&str)
	}
}
