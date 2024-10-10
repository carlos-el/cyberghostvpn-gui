# cyberghostvpn-gui
User interface written in Golang for [CyberGhostVPN](https://www.cyberghostvpn.com/) on Linux using the [Fyne](https://github.com/fyne-io/fyne) package.  
It has only been tested on Ubuntu 22.04.

## Usage requirements
- You need to have installed the [CyberGhostVPN CLI](https://support.cyberghostvpn.com/hc/en-us/articles/360020436274-Set-Up-CyberGhost-VPN-CLI-App-on-Linux) and login into your account.
- The GUI needs to be executed with sudo privileges as CyberGhostVPN CLI requires sudo as well.

## Installing
### Downloading the release from GitHub
Requirements:
- Makefile

Steps: 
1. Download the desired release from GitHub.  
2. Extract the contents of the compressed file.  
3. Give a look to the Makefile for the different install and uninstall options.  
4. After installing you will have the program available in you applications with a desktop icon.  
5. You can also run it with `sudo cyberghostvpn-gui`.

### From source
Requirements:
- Go
- Makefile
- Fyne Go package  

Steps: 
1. Execute `make build-prod` to get the executable in the `bin` directory.  
2. Alternatively execute `make build-linux` to get the linux bundle in the `bin` directory (this may take a while). 
3. Install as in the if you were downloading the release from GitHub (previous section).

## Running without invoking sudo
(Do it at your own risk) It is possible to run the GUI without having to invoke sudo every time using the terminal.
The steps to achieve are the following:
1. Add your user to the sudoers file for executing the GUI without password prompt. 
    - Use the `/etc/sudoers.d` file if possible and add the following line:  
    `your_user ALL=(root) NOPASSWD: /home/your_user/.local/bin/cyberghostvpn-gui`  
    Replace `your_user` with your username and make sure that the path to the executable is correct. It should be if you installed the GUI using the `make user-install` but it may be different if you did a system wide installation.
2. Edit the desktop file used to run the GUI to execute as with sudo.
    - Locate the `.desktop` file. If you did a user installation it should be under `/home/you_user/.local/share/applications/cyberghostvpn-gui.desktop`. If not it may be under `usr/local/share/applications/cyberghostvpn-gui.desktop`.
    - Edit the file and add `sudo` to the command used in the `Exec` directive. It should look like this if you did a user installation:  
    `Exec=sudo /home/carlos/.local/bin/cyberghostvpn-gui`
3. Now you should be able to run the GUI from the application launcher without any issues. Sometimes it takes some seconds for the application launcher to recognize the changes in the `.desktop` file, take it into account.

## Advice
The CyberGhostVPN CLI requires sudo privileges to run. I find this somewhat concerning. If you just want to always connect to the same servers from time to time, you are probably better off [configuring your desired servers directly in your system using OpenVPN](https://support.cyberghostvpn.com/hc/en-us/articles/360007929314-Set-Up-OpenVPN-on-Linux-Ubuntu-via-Network-Manager).

## Disclaimer
This project was developed independently for personal use. CyberGhost has no affiliation, nor has control over the content or availability of this project.
