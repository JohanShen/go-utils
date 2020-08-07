package logger

import (
	"fmt"
	"github.com/JohanShen/go-utils/utils"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type TimeWriter struct {
	dir, fileName string
	lastMakeTime  time.Time
	fileMode      os.FileMode
	CurrentPath   string
	PathTemplate  string
	files         map[string]*fileObject
	locker        *sync.Mutex
}

type fileObject struct {
	file       *os.File
	writeLen   int
	createTime time.Time
}

func NewTimeWriter(logTemplate string) (obj *TimeWriter) {
	obj = &TimeWriter{
		PathTemplate: logTemplate,
		locker:       &sync.Mutex{},
		fileMode:     0666,
		files:        make(map[string]*fileObject),
	}
	return obj
}

// 创建新的文件对象
func (l *TimeWriter) makeNewFile(logPath string) (file *os.File, err error) {
	//now := time.Now()
	//logPath := utils.XTime(now).Format(l.PathTemplate)
	l.dir, l.fileName = filepath.Split(logPath)
	if ok, _ := utils.IsDirExists(l.dir); !ok {
		if err := os.MkdirAll(l.dir, l.fileMode); err != nil {
			return nil, err
		}
	}

	file, err = os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, l.fileMode)
	return file, err
}

func (l *TimeWriter) Write(p []byte) (n int, err error) {
	l.locker.Lock()
	defer l.locker.Unlock()

	now := time.Now()
	logPath := utils.XTime(now).Format(l.PathTemplate)

	var file *fileObject
	ok := false
	for {
		if file, ok = l.files[logPath]; ok {
			break
		}
		if obj, err := l.makeNewFile(logPath); err == nil {
			l.files[logPath] = &fileObject{file: obj, createTime: now}
			l.CurrentPath = logPath
			l.lastMakeTime = now
		} else {
			fmt.Printf("write fail, msg(%s)\n", err)
			return 0, err
		}
		go l.closeFile()
	}

	n, err = file.file.Write(p)
	file.writeLen += n
	return n, err
}

func (l *TimeWriter) closeFile() {

	if len(l.files) > 1 {
		now := time.Now()
		closeFiles := make([]string, 0, len(l.files))
		// 需要将之前的文件句柄关闭
		for key, item := range l.files {
			// 将非当前文件，且超过15分钟的文件句柄关闭
			if key == l.CurrentPath || item.createTime.Sub(now) > time.Minute*-15 {
				continue
			}
			closeFiles = append(closeFiles, key)
			if err := item.file.Sync(); err != nil {
				//有问题但是不处理
				println("写入文件时 closeFile() ", err)
			}
			if err := item.file.Close(); err != nil {
				//有问题但是不处理
				println("关闭文件时 closeFile() ", err)
			}
		}
		if len(closeFiles) > 0 {
			l.locker.Lock()
			for _, v := range closeFiles {
				delete(l.files, v)
			}
			l.locker.Unlock()
		}
	}
}

func (l *TimeWriter) Close() error {
	l.locker.Lock()
	defer l.locker.Unlock()

	for _, item := range l.files {
		if err := item.file.Sync(); err != nil {
			//有问题但是不处理
			println("写入文件时 Close() ", err)
		}
		if err := item.file.Close(); err != nil {
			//有问题但是不处理
			println("关闭文件时 Close() ", err)
		}
	}
	return nil
}
