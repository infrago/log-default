package log_default

import (
	"io"
	"os"

	"github.com/infrago/log"
)

type (
	defaultDriver struct {
	}
	defaultConnect struct {
		instance       *log.Instance
		stdout, stderr io.Writer
	}
)

func (driver *defaultDriver) Connect(inst *log.Instance) (log.Connect, error) {
	return &defaultConnect{
		instance: inst, stdout: os.Stdout, stderr: os.Stderr,
	}, nil
}

// 打开连接
func (this *defaultConnect) Open() error {
	return nil
}

// 关闭连接
func (this *defaultConnect) Close() error {
	return nil
}

func (this *defaultConnect) Write(msgs ...log.Log) error {
	for _, msg := range msgs {
		if err := this.wrinting(msg); err != nil {
			return err
		}
	}
	return nil
}

func (this *defaultConnect) wrinting(msg log.Log) error {
	body := this.instance.Format(msg)
	if msg.Level <= log.LevelWarning {
		this.stderr.Write([]byte(body + "\n"))
	} else {
		this.stdout.Write([]byte(body + "\n"))
	}
	return nil
}
