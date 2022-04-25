package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli"
)

//创建文章
func CreatePost(filename string, title string, postdes string) {
	fmt.Println("create a new post")

	var (
		file *os.File
		err  error
	)

	now := time.Now()
	fmt.Println(now.Local())

	//使用追加模式打开文件
	file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("Open file err =", err)
		return
	}
	defer file.Close()
	//写入文件

	filetmp, err := os.Open("tml.md")
	if err != nil {
		panic(err)
	}
	i := 0
	rd := bufio.NewReader(filetmp)
	for {
		line, err := rd.ReadString('\n')
		// i++
		if err != nil || io.EOF == err {
			break
		} else {
			//fmt.Println(line)
			if i == 1 {
				line = "title: \"" + title + "\"\n"
			}

			if i == 2 {
				line = "date: " + time.Now().Local().Format(time.RFC3339) + "\n"
			}

			if i == 3 {
				line = "description: \"" + postdes + "\"\n"
			}

			_, err = file.Write([]byte(line))
			if err != nil {
				fmt.Println("Write file err =", err)
				return
			}
		}
		i++
	}
}

func main() {
	app := cli.NewApp()

	//参数定义及说明
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "post, p",
			Value: "Post Title Name Tile Of Post",
			Usage: "Title Of Post",
		},

		cli.StringFlag{
			Name:  "description, d",
			Value: "Description Of Post",
			Usage: "Description Of Post",
		},
	}

	app.Name = "Create New Hugo Post"
	app.Usage = "NewPost -t Title Of Post -d Description Of Post -a Author"
	app.Action = func(c *cli.Context) error {
		poststr := c.String("p")

		poststrarr := strings.Split(poststr, "/")
		postname := poststrarr[len(poststrarr)-1]

		title := strings.Split(postname, ".")[0]

		postdirpath := poststr[:(len(poststr) - len(postname))]

		os.MkdirAll(postdirpath, os.ModePerm)

		//postdes := c.String("d")
		CreatePost(poststr, title, title)
		return nil
	}
	app.Run(os.Args)
}
