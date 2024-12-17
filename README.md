# FTP File Upload & Clean Utility

This Go program allows you to interact with an FTP server to upload files or clean remote directories. You can either upload files from a local directory to a remote FTP server or clean a specified remote directory by deleting its files and subdirectories.

## Prerequisites

- Go (1.16 or higher) installed.
- FTP server details (address, username, password).
- The `github.com/jlaffaye/ftp` package is required. You can install it by running:

    ```bash
    go get github.com/jlaffaye/ftp
    ```

## Configuration

You can configure the program using either environment variables or command-line parameters.

### Option 1: Using Command-Line Parameters

Run the program with the necessary flags to specify FTP server details and local/remote file paths. Here is the full list of available flags:

```bash
go run main.go --server <ftp_server> --port <ftp_port> --username <ftp_user> --password <ftp_password> --file <local_file_or_directory> --remote <remote_directory> [--clean] [--push]

Parameters:
--server: FTP server address (e.g., ftpupload.net).
--port: FTP server port (default is 21 if not specified).
--username: FTP username.
--password: FTP password.
--file: Path to the local file or directory to upload. For clean operation, this is not required.
--remote: Path to the remote directory where files will be uploaded or cleaned.
--clean: Clean the remote path (deletes all files and folders).
--push: Upload files to the remote path (default is no action).
Example usage:

Uploading Files:

bash
Copy code
go run main.go --server ftpupload.net --port 21 --username "username" --password "password" --file ./test --remote /domain.com/htdocs/
Cleaning Remote Directory:

bash
Copy code
go run main.go --server ftpupload.net --port 21 --username "username" --password "password" --remote /domain.com/htdocs/ --clean
Option 2: Using Environment Variables
Alternatively, you can set the following environment variables and run the program without parameters.

bash
Copy code
export FTP_SERVER="ftpupload.net"
export FTP_PORT="21"
export FTP_USER="username"
export FTP_PASSWORD="password"
export LOCALPATH="/test"
export REMOTEPATH="/domain.com/htdocs/"
Then, run the program with the go run command:

Uploading Files:

bash
Copy code
go run main.go --push
Cleaning Remote Directory:

bash
Copy code
go run main.go --clean
Functionality
Uploading Files (--push flag): The program will walk through the specified local directory and upload all files and subdirectories to the remote FTP server. If the local path is a file, it will upload that file to the remote directory.

Cleaning Remote Directory (--clean flag): The program will delete all files and subdirectories in the specified remote directory.

Example:
bash
Copy code
go run main.go --server ftpupload.net --port 21 --username "username" --password "password" --file ./test --remote /domain.com/htdocs/ --push
This will upload the contents of the ./test folder to the /domain.com/htdocs/ folder on the FTP server.

bash
Copy code
go run main.go --server ftpupload.net --port 21 --username "username" --password "password" --remote /domain.com/htdocs/ --clean
This will clean (delete) all files and subdirectories inside the /domain.com/htdocs/ folder on the FTP server.
