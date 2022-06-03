package utils

import (
	"context"
	"fmt"
	"github.com/pkg/sftp"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"time"
)

// 上传图片到七牛云，然后返回状态和图片的url
func UploadToQiNiu(file *multipart.FileHeader) (int, string, string) {

	//var AccessKey = "_OBtN_TsdTp7_CFMu423WOBCQe3Nf9rAk5GURp1E" // 秘钥对
	//var SerectKey = "VK49p-yXFAjWvycZnzwKi0VHW4DDqFCabn_RdNY4"
	//var Bucket = "douyin-video" // 空间名称
	////var ImgUrl = "http://rbppmzeve.hn-bkt.clouddn.com/" // 自定义域名或测试域名

	var AccessKey = "1-kzJOvB1oQ40HMO9xSTjWDBv3ARrZWDMBFWxjDS"
	var SerectKey = "7uAxrKuIp2XXC10kAkVBrPNTkZ4l_2Q9otT3rUFt"
	var Bucket = "videodouyin"
	var ImgUrl = "http://rcw3mo7gu.hn-bkt.clouddn.com/"

	src, err := file.Open()
	if err != nil {
		return 10011, err.Error(), ""
	}
	defer src.Close()
	imgKey := "video/" + strconv.FormatInt(time.Now().UnixNano(), 10)
	putPlicy := storage.PutPolicy{
		Scope:         Bucket,
		PersistentOps: "vframe/jpg/offset/0/w/480/h/240",
		ForceSaveKey:  true,
		SaveKey:       imgKey,
		//PersistentNotifyURL: "http://rbppmzeve.hn-bkt.clouddn.com/Q9GLAFFqfCrYF6YfQAcON4w4Ezs=",
	}
	mac := qbox.NewMac(AccessKey, SerectKey)

	// 获取上传凭证
	upToken := putPlicy.UploadToken(mac)

	// 配置参数
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan, // 华南区
		UseCdnDomains: false,
		UseHTTPS:      false, // 非https
	}
	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}        // 上传后返回的结果
	putExtra := storage.PutExtra{} // 额外参数

	// 上传 自定义key，可以指定上传目录及文件名和后缀，
	key := "image/" + file.Filename // 上传路径，如果当前目录中已存在相同文件，则返回上传失败错误
	err = formUploader.Put(context.Background(), &ret, upToken, key, src, file.Size, &putExtra)

	// 以默认key方式上传
	// err = formUploader.PutWithoutKey(context.Background(), &ret, upToken, src, fileSize, &putExtra)

	// 自定义key，上传指定路径的文件
	// localFilePath = "./aa.jpg"
	// err = formUploader.PutFile(context.Background(), &ret, upToken, key, localFilePath, &putExtra)

	// 默认key，上传指定路径的文件
	// localFilePath = "./aa.jpg"
	// err = formUploader.PutFile(context.Background(), &ret, upToken, key, localFilePath, &putExtra)

	if err != nil {
		code := 501
		return code, err.Error(), ""
	}

	url := ImgUrl + ret.Key // 返回上传后的文件访问路径
	imgUrl := ImgUrl + "XlMSXysTOOFC84dEL3HgtjKxtAY=" + ret.Hash
	return 0, url, imgUrl
}

func connect(user, password, host string, port int) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))
	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //ssh.FixedHostKey(hostKey),
	}
	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)
	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}
	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}
	return sftpClient, nil
}
func uploadFile(sftpClient *sftp.Client, localFilePath string, remotePath string) {
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		fmt.Println("os.Open error : ", localFilePath)
		log.Fatal(err)
	}
	defer srcFile.Close()
	var remoteFileName = path.Base(localFilePath)
	dstFile, err := sftpClient.Create(path.Join(remotePath, remoteFileName))
	if err != nil {
		fmt.Println("sftpClient.Create error : ", path.Join(remotePath, remoteFileName))
		log.Fatal(err)
	}
	defer dstFile.Close()
	ff, err := ioutil.ReadAll(srcFile)
	if err != nil {
		fmt.Println("ReadAll error : ", localFilePath)
		log.Fatal(err)
	}
	dstFile.Write(ff)
	fmt.Println(localFilePath + " copy file to remote server finished!")
}
func uploadDirectory(sftpClient *sftp.Client, localPath string, remotePath string) {
	localFiles, err := ioutil.ReadDir(localPath)
	if err != nil {
		log.Fatal("read dir list fail ", err)
	}
	for _, backupDir := range localFiles {
		localFilePath := path.Join(localPath, backupDir.Name())
		remoteFilePath := path.Join(remotePath, backupDir.Name())
		if backupDir.IsDir() {
			sftpClient.Mkdir(remoteFilePath)
			uploadDirectory(sftpClient, localFilePath, remoteFilePath)
		} else {
			uploadFile(sftpClient, path.Join(localPath, backupDir.Name()), remotePath)
		}
	}
	fmt.Println(localPath + " copy directory to remote server finished!")
}

func DoBackup(host string, port int, userName string, password string, localPath string, remotePath string) {
	var (
		err        error
		sftpClient *sftp.Client
	)
	start := time.Now()
	sftpClient, err = connect(userName, password, host, port)
	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()
	_, errStat := sftpClient.Stat(remotePath)
	if errStat != nil {
		log.Fatal(remotePath + " remote path not exists!")
	}
	_, err = ioutil.ReadDir(localPath)
	if err != nil {
		log.Fatal(localPath + " local path not exists!")
	}
	uploadDirectory(sftpClient, localPath, remotePath)
	elapsed := time.Since(start)
	fmt.Println("elapsed time : ", elapsed)
}
