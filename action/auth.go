package action

import (
	`encoding/json`
	`fmt`
	`io/ioutil`
	`os`
	`path/filepath`
	`strings`

	`github.com/tal-tech/go-zero/tools/goctl/plugin`
	`github.com/tal-tech/go-zero/tools/goctl/util`
	"github.com/urfave/cli/v2"
)

type Url struct {
	Base string `json:"base"`
	Doc string `json:"doc"`
	Url string `json:"url"`
	Handler string `json:"handler"`
}

func AuthApi(ctx *cli.Context) error {
	// 生成mkdir
	dir, err := mkDir("./auth")
	if err != nil {
		return err
	}
	// 生成文件
	filename := filepath.Join(dir, "auth.json")
	fmt.Printf("文件位置:%v", filename)

	if !util.FileExists(filename) {
		file, err := util.CreateIfNotExist(filename)
		if err != nil {
			return err
		}
		_ = file.Close()
	}
	// 读取文件
	all, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	root := make([]string, 0)
	if len(all) > 0 {
		err = json.Unmarshal(all, &root)
		if err != nil {
			return err
		}
	}
	// 添加文件
	newPlugin, err := plugin.NewPlugin()
	if err != nil {
		return err
	}
	base := newPlugin.Api.Service.Name
	groups := newPlugin.Api.Service.Groups
	for _, ser := range groups {
		middleware := ser.Annotation.Properties["middleware"]
		if strings.Contains(middleware, "AdminAuth") {
			for _, ro := range ser.Routes {
				summary := strings.Trim(ro.AtDoc.Properties["summary"], `"`)
				if len(summary) == 0 {
					summary = ro.Doc[0]
				}
				url := ro.Path
				handler := ro.AtServerAnnotation.Properties["handler"]
				if len(handler) == 0 {
					handler = ro.Handler
				}
				data, _ := json.Marshal(Url{
					Base:    base,
					Doc:     summary,
					Url:     url,
					Handler: handler,
				})
				text := string(data)
				isExist := false
				for _, v := range root {
					if text == v {
						isExist = true
					}
				}
				if !isExist {
					root = append(root, text)
				}
			}
		}
	}
	// 清空内容
	data, err := json.Marshal(root)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 077)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func mkDir(target string) (string, error) {
	abs, err := filepath.Abs(target)
	if err != nil {
		return "", err
	}

	err = util.MkdirIfNotExist(abs)
	if err != nil {
		return "", err
	}
	return abs, nil
}