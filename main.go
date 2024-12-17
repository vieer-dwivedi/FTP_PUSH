package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jlaffaye/ftp"
)

func createRemoteDir(conn *ftp.ServerConn, remoteDir string) error {
	err := conn.ChangeDir(remoteDir)
	if err != nil {
		err = conn.MakeDir(remoteDir)
		if err != nil {
			return fmt.Errorf("error creating remote directory %s: %v", remoteDir, err)
		}
		err = conn.ChangeDir(remoteDir)
		if err != nil {
			return fmt.Errorf("error changing to remote directory %s: %v", remoteDir, err)
		}
	}
	return nil
}

func uploadFile(conn *ftp.ServerConn, localFile, remotePath string) error {
	file, err := os.Open(localFile)
	if err != nil {
		return fmt.Errorf("error opening local file %s: %v", localFile, err)
	}
	defer file.Close()

	dirPath := filepath.Dir(remotePath)
	err = createRemoteDir(conn, dirPath)
	if err != nil {
		return fmt.Errorf("error creating remote directory: %v", err)
	}

	err = conn.Stor(remotePath, file)
	if err != nil {
		return fmt.Errorf("error uploading file %s: %v", localFile, err)
	}

	return nil
}

func main() {
	ftpServer := flag.String("server", os.Getenv("FTP_SERVER"), "FTP server address")
	ftpPort := flag.String("port", os.Getenv("FTP_PORT"), "FTP server port")
	ftpUser := flag.String("username", os.Getenv("FTP_USER"), "FTP username")
	ftpPassword := flag.String("password", os.Getenv("FTP_PASSWORD"), "FTP password")
	localPath := flag.String("file", os.Getenv("LOCALPATH"), "Path to the local file or directory")
	remotePath := flag.String("remote", os.Getenv("REMOTEPATH"), "Remote path to upload file or directory")
	clean := flag.Bool("clean", false, "Clean the remote path")
	push := flag.Bool("push", false, "Push the files to the remote path")

	flag.Parse()

	if *ftpServer == "" || *ftpUser == "" || *ftpPassword == "" || (*localPath == "" && !*clean) || *remotePath == "" {
		fmt.Println("Usage: go run main.go --server <ftp_server> --port <ftp_port> --username <ftp_user> --password <ftp_password> --file <file_path> --remote <remote_path> [--clean] [--push]")
		return
	}

	conn, err := ftp.DialTimeout(fmt.Sprintf("%s:%s", *ftpServer, *ftpPort), 10*time.Second)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Quit()

	err = conn.Login(*ftpUser, *ftpPassword)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if *push {
		info, err := os.Stat(*localPath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if info.IsDir() {
			err = filepath.Walk(*localPath, func(localFilePath string, fileInfo os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				relPath, _ := filepath.Rel(*localPath, localFilePath)
				fullRemotePath := filepath.Join(*remotePath, relPath)

				if fileInfo.IsDir() {
					err := createRemoteDir(conn, fullRemotePath)
					if err != nil {
						fmt.Println("Error creating remote directory:", err)
						return err
					}
				} else {
					err := uploadFile(conn, localFilePath, fullRemotePath)
					if err != nil {
						fmt.Println("Error uploading file:", err)
						return err
					}
					fmt.Println("Uploaded:", localFilePath, "to", fullRemotePath)
				}
				return nil
			})
			if err != nil {
				fmt.Println("Error walking through directory:", err)
			}
		} else {
			err = uploadFile(conn, *localPath, *remotePath)
			if err != nil {
				fmt.Println("Error uploading file:", err)
				return
			}
			fmt.Println("Uploaded:", *localPath, "to", *remotePath)
		}
	}
}
