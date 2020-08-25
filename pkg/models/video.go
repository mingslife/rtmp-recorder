package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"rtmp-recorder/pkg/utils"
	"strings"
)

type Video struct {
	Id       string `json:"id"`
	Status   string `json:"status"`
	Duration string `json:"duration"`
	Name     string `json:"name"`
	Url      string `json:"url"`
}

// var image string = "nginx:alpine"
var image string = "m1n9.vip/rtmp-recorder-worker:latest"
var containerNamePrefix string = "worker-"

func GetVideos() (s []*Video) {
	s = make([]*Video, 0)
	execCmd := fmt.Sprintf(`docker ps -a --format="{{json .}}" --filter "name=%s*" --filter "label=masterId=%s"`, containerNamePrefix, MasterId)
	fmt.Println("cmd:", execCmd)
	var out bytes.Buffer
	cmd := exec.Command("sh", "-c", execCmd)
	cmd.Stdout = &out
	cmd.Run()
	if result := strings.TrimSpace(out.String()); result != "" {
		for _, line := range strings.Split(result, "\n") {
			var m map[string]string
			json.Unmarshal([]byte(line), &m)
			v := &Video{
				Id:       m["Names"][len(containerNamePrefix):],
				Status:   strings.Split(m["Status"], " ")[0],
				Duration: m["RunningFor"],
			}
			s = append(s, v)
		}
	}
	return
}

func GetVideo(id string) *Video {
	execCmd := fmt.Sprintf(`docker inspect --format="{{json .Config.Env}}" %s%s`, containerNamePrefix, id)
	fmt.Println("cmd:", execCmd)
	var out bytes.Buffer
	cmd := exec.Command("sh", "-c", execCmd)
	cmd.Stdout = &out
	cmd.Run()
	if result := strings.TrimSpace(out.String()); result != "" {
		var s []string
		json.Unmarshal([]byte(result), &s)
		env := make(map[string]string)
		for _, envPair := range s {
			parts := strings.Split(envPair, "=")
			env[parts[0]] = parts[1]
		}
		v := &Video{
			Id:   id,
			Name: env["VIDEO_NAME"],
			Url:  env["VIDEO_URL"],
			// Status:   strings.Split(m["Status"], " ")[0],
			// Duration: m["RunningFor"],
		}
		return v
	}
	return nil
}

func (m *Video) Save() error {
	m.Id = utils.NewId()
	pwd, _ := os.Getwd()
	dataPath := path.Join(pwd, "data")
	execCmd := fmt.Sprintf(`docker run --name %s%s --rm -v %s:/data -e VIDEO_NAME=%s -e VIDEO_URL=%s --label masterId=%s -d %s`, containerNamePrefix, m.Id, dataPath, m.Name, m.Url, MasterId, image)
	fmt.Println("cmd:", execCmd)
	cmd := exec.Command("sh", "-c", execCmd)
	return cmd.Run()
}

// func (m *Video) Update() error {
// 	return Db.Model(m).Update(map[string]interface{}{
// 		"role_name": m.RoleName,
// 		"active":    m.Active,
// 	}).Error
// }

func (m *Video) Delete() error {
	execCmd := fmt.Sprintf(`docker stop --time 300 %s%s`, containerNamePrefix, m.Id)
	fmt.Println("cmd:", execCmd)
	cmd := exec.Command("sh", "-c", execCmd)
	return cmd.Run()
}
