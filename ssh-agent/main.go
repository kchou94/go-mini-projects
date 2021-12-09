package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func check(err error) {
	if err != nil {
		fmt.Printf("Error happened: %s \n", err)
		os.Exit(1)
	}
}

var (
	sshUserName        = "vagrant"
	sshPassword        = "vagrant"
	sshKeyPath         = "/home/kchou/.vagrant.d/insecure_private_key"
	sshHostName        = "192.168.56.22:22"
	commandToExec      = "ls -alh /home/vagrant"
	fileToUpload       = "./upload.txt"
	fileUploadLocation = "/home/vagrant/upload.txt"
	fileToDownload     = "/home/vagrant/download.txt"
)

func main() {
	fmt.Println("......Golang SSH Demo......")

	// conf := sshDemoWithPassword()
	conf := sshDemoWithPrivateKey()

	// open ssh connection
	sshClient, err := ssh.Dial("tcp", sshHostName, conf)
	check(err)
	session, err := sshClient.NewSession()
	check(err)
	defer session.Close()

	// execute command on remote server
	var b bytes.Buffer
	session.Stdout = &b
	err = session.Run(commandToExec)
	check(err)
	log.Printf("%s: %s", commandToExec, b.String())

	// open sftp connection
	sftpClient, err := sftp.NewClient(sshClient)
	check(err)
	defer sftpClient.Close()

	// create a file
	createFile, err := sftpClient.Create(fileToDownload)
	check(err)
	text := "This file created by Golang SSH.\nThis will be downloaded by Golang SSH\n"
	_, err = createFile.Write([]byte(text))
	check(err)
	fmt.Println("Created file", fileToDownload)

	// upload a file
	srcFile, err := os.Open(fileToUpload)
	check(err)
	defer srcFile.Close()

	dstFile, err := sftpClient.Create(fileUploadLocation)
	check(err)
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	check(err)
	fmt.Println("File uploaded successfully", fileUploadLocation)

	// download a file
	remoteFile, err := sftpClient.Open(fileToDownload)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open remote file: %v\n", err)
		return
	}
	defer remoteFile.Close()

	localFile, err := os.Create("./download.txt")
	check(err)
	defer localFile.Close()

	_, err = io.Copy(localFile, remoteFile)
	check(err)
	fmt.Println("File downloaded successfully")
}

func sshDemoWithPassword() *ssh.ClientConfig {
	conf := &ssh.ClientConfig{
		User: sshUserName,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return conf
}

func sshDemoWithPrivateKey() *ssh.ClientConfig {
	keyByte, err := ioutil.ReadFile(sshKeyPath)
	check(err)
	key, err := ssh.ParsePrivateKey(keyByte)
	check(err)

	conf := &ssh.ClientConfig{
		User: sshUserName,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return conf
}
