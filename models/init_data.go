package models

import (
	"io/ioutil"
	"log"
)

// InitializeData 初始化数据库数据
func InitializeData() error {
	// 读取西安攻略内容
	xianContent, err := ioutil.ReadFile("templates/xianthree.html")
	if err != nil {
		log.Printf("读取西安攻略文件失败: %v", err)
		return err
	}

	// 读取山东攻略内容
	shandongContent, err := ioutil.ReadFile("templates/shandonglaodong.html")
	if err != nil {
		log.Printf("读取山东攻略文件失败: %v", err)
		return err
	}

	// 创建西安攻略页面
	xianPage := &Page{
		Title:   "西安4日暴走攻略",
		Content: string(xianContent),
		Path:    "/xianthree",
	}
	if err := xianPage.Save(); err != nil {
		log.Printf("保存西安攻略失败: %v", err)
		return err
	}

	// 创建山东攻略页面
	shandongPage := &Page{
		Title:   "山东环海自驾游5天攻略",
		Content: string(shandongContent),
		Path:    "/shandong",
	}
	if err := shandongPage.Save(); err != nil {
		log.Printf("保存山东攻略失败: %v", err)
		return err
	}

	log.Println("数据初始化完成")
	return nil
} 