package toolbox

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/sftp"
	"go.uber.org/zap"
	"strings"
)

type SSHSftpClient struct {
	SSHClient
	UploadFile chan *UploadFile
	confirmMap map[string]chan *ConfirmInfo
}

func (this_ *SSHSftpClient) listenUpload() {
	if this_.UploadFile == nil {
		this_.UploadFile = make(chan *UploadFile, 10)

		go func() {
			for {
				select {
				case uploadFile := <-this_.UploadFile:
					this_.work(&SFTPRequest{
						Work:     "upload",
						WorkId:   uploadFile.WorkId,
						Dir:      uploadFile.Dir,
						Place:    uploadFile.Place,
						File:     uploadFile.File,
						FullPath: uploadFile.FullPath,
					})
				}
			}

		}()
	}
	return
}
func (this_ *SSHSftpClient) newSftp() (sftpClient *sftp.Client, err error) {
	err = this_.initClient()
	if err != nil {
		return
	}

	sftpClient, err = sftp.NewClient(this_.sshClient)
	if err != nil {
		this_.WSWriteError("SSH FTP创建失败:" + err.Error())
		return
	}

	return
}

func (this_ *SSHSftpClient) start() {
	SSHSftpCache[this_.Token] = this_
	go this_.ListenWS(this_.onEvent, this_.onMessage, this_.CloseClient)
	this_.listenUpload()
	this_.WSWriteEvent("ftp ready")
}

func (this_ *SSHSftpClient) closeSftClient(sftpClient *sftp.Client) {
	if sftpClient == nil {
		return
	}
	err := sftpClient.Close()
	if err != nil {
		fmt.Println("sftp client close error", err)
		return
	}
}
func (this_ *SSHSftpClient) onEvent(event string) {
	var err error
	this_.Logger.Info("SSH FTP On Event:", zap.Any("event", event))
	switch strings.ToLower(event) {
	case "ftp start":
		var sftpClient *sftp.Client
		sftpClient, err = this_.newSftp()
		if err != nil {
			return
		}
		defer this_.closeSftClient(sftpClient)
		this_.WSWriteEvent("ftp created")
	}
}

func (this_ *SSHSftpClient) onMessage(bs []byte) {

	go func() {
		var request *SFTPRequest
		err := json.Unmarshal(bs, &request)
		if err != nil {
			fmt.Println("sftp ws message to struct err:", err)
			return
		}
		this_.work(request)
	}()
}
