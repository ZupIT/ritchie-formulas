#!/bin/bash

USER=$(pwd)

declare -rx CMDS=(
    'sleep 1'
)

RED=$(tput setaf 1)
GREEN=$(tput setaf 2)
YELLOW=$(tput setaf 3)
LIME_YELLOW=$(tput setaf 190)
CYAN=$(tput setaf 6)

runConfigLinux() {
    sayHello "$SYSTEM"
    createDirLog

    case "$CONFIGURATION" in
    "Complete configuration")
        completeConfiguration
        ;;
    "Select programs")
        selectConfiguration
        ;;
    *) echo -e "${RED}Invalid option, try again!" ;;
    esac
}

completeConfiguration() {
    installSPropertiesCommon
    installWget
    installCertificates
    installCurl
    installGnupgAgent
    installSnap
    installJdk
    installPython
    installNode
    installNpm
    installVsCode
    installGit
    installMaven
    installDocker
    installIntellij
    installPostman
    installMySql
    installMySqlWorkbench
    installChrome
    installSpotify
    installDiscord

    finishInstall
}

selectConfiguration() {
    init "Starting manual installation"
    clearLine 1

    echo -ne "${YELLOW}Select the tool you want to install:"

    selections=(
        "Properties Common"
        "Wget"
        "Curl"
        "Gnupg Agent"
        "Snap"
        "Jdk"
        "Python"
        "Node"
        "Npm"
        "VsCode"
        "Git"
        "Maven"
        "Docker"
        "Intellij"
        "Postman"
        "MySql"
        "MySql Workbench"
        "Chrome"
        "Spotify"
        "Discord"
        "Quit"
    )

    while opt=$(zenity --title="Installation" --text="Tools" --list \
        --column="Name" "${selections[@]}"); do

        case "$opt" in
        "${selections[0]}") installSPropertiesCommon ;;
        "${selections[1]}") installWget ;;
        "${selections[2]}") installCurl ;;
        "${selections[3]}") installGnupgAgent ;;
        "${selections[4]}") installSnap ;;
        "${selections[5]}") installJdk ;;
        "${selections[6]}") installPython ;;
        "${selections[7]}") installNode ;;
        "${selections[8]}") installNpm ;;
        "${selections[9]}") installVsCode ;;
        "${selections[10]}") installGit ;;
        "${selections[11]}") installMaven ;;
        "${selections[12]}") installDocker ;;
        "${selections[13]}") installIntellij ;;
        "${selections[14]}") installPostman ;;
        "${selections[15]}") installMySql ;;
        "${selections[16]}") installMySqlWorkbench ;;
        "${selections[17]}") installChrome ;;
        "${selections[18]}") installSpotify ;;
        "${selections[19]}") installDiscord ;;
        "${selections[20]}") break ;;
        *) zenity --error --text="Invalid option. Try another one." ;;
        esac

    done

    finishInstall
}

sayHello() {
    checkIfIsInstalled "lolcat"
    if mycmd; then
        sudo apt-get install lolcat -y 2>&1 | tee "$USER"/log/tmp.log
    fi

    checkIfIsInstalled "cowsay"
    if mycmd; then
        sudo apt-get install cowsay -y 2>&1 | tee "$USER"/log/tmp.log
    fi

    cowsay -f gnu Setting up environment "$1" with RITCHIE! 🦸🚀 | lolcat

    printf "${CYAN}OS version: %s"
    lsb_release -r
}

installSPropertiesCommon() {
    init "Starting installation of Properties Common" "Properties Common installation complete"

    sudo apt-get install software-properties-common -y 2>&1 | tee "$USER"/log/tmp.log
}

installWget() {
    init "Starting installation of GNU Wget" "GNU Wget installation completed"

    checkIfIsInstalled "wget"
    if [ "$?" -eq 1 ]; then
        information "GNU Wget is already installed!"
        return
    fi

    sudo apt-get install apt-transport-https wget -y 2>&1 | tee "$USER"/log/tmp.log
}

installCertificates() {
    init "Starting installation of Certificates" "Certificate Installation Complete"

    sudo apt-get install ca-certificates -y 2>&1 | tee "$USER"/log/tmp.log
}

installCurl() {
    init "Starting installation of Curl Wget" "Curl Installation Complete"

    checkIfIsInstalled "curl"
    if [ "$?" -eq 1 ]; then
        information "Curl is already installed!"
        return
    fi

    sudo apt-get install curl -y 2>&1 | tee "$USER"/log/tmp.log
}

installGnupgAgent() {
    init "Starting installation of the GNUPG Agent" "GNUPG installation completed"

    sudo apt-get install gnupg-agent -y 2>&1 | tee "$USER"/log/tmp.log
}

installSnap() {
    init "Starting Snap installation" "Snap Installation Complete"

    checkIfIsInstalled "snap"
    if [ "$?" -eq 1 ]; then
        information "Snap is already installed!"
        return
    fi

    sudo apt-get install snapd -y 2>&1 | tee "$USER"/log/tmp.log
}

installJdk() {
    init "Starting JDK installation"
    clearLine 1

    sudo add-apt-repository ppa:linuxuprising/java -y

    sudo apt-get update -y 2>&1 | tee "$USER"/log/tmp.log

    echo -ne "${YELLOW}Select the version of Openjdk to be installed:\\n"

    options=("8" "11" "12" "Quit")

    select opt in "${options[@]}"; do
        case $opt in
        "8")
            init "Starting JDK 8 installation" "JDK 8 installation completed"
            sudo apt-get install openjdk-8-jdk -y 2>&1 | tee "$USER"/log/tmp.log
            break
            ;;
        "11")
            init "Starting JDK 11 installation" "JDK 11 installation completed"
            sudo apt-get install openjdk-11-jdk -y 2>&1 | tee "$USER"/log/tmp.log
            break
            ;;
        "12")
            init "Starting JDK 12 installation" "JDK 12 installation completed"
            sudo apt-get install openjdk-12-jdk -y 2>&1 | tee "$USER"/log/tmp.log
            break
            ;;
        "Quit")
            clearLine 1
            echo -ne "${RED}[ ❌ ] JDK not installed!"
            break
            ;;
        *) echo -e "${RED}Invalid option, try again!" ;;
        esac
    done
}

installPython() {
    init "Starting Python3 installation" "Python3 installation completed"

    checkIfIsInstalled "python3"
    if [ "$?" -eq 1 ]; then
        information "Python3 is already installed!"
        return
    fi

    sudo add-apt-repository ppa:deadsnakes/ppa -y 2>&1 | tee "$USER"/log/tmp.log
    sudo apt-get install python3.8 -y 2>&1 | tee "$USER"/log/tmp.log
}

installNode() {
    init "Starting Node installation" "Node installation complete"

    checkIfIsInstalled "node"
    if [ "$?" -eq 1 ]; then
        information "Node is already installed!"
        return
    fi

    sudo apt-get install nodejs -y 2>&1 | tee "$USER"/log/tmp.log
}

installNpm() {
    init "Starting Npm installation" "Npm Installation Complete"

    checkIfIsInstalled "npm"
    if [ "$?" -eq 1 ]; then
        information "Npm is already installed!"
        return
    fi

    sudo apt-get install npm -y 2>&1 | tee "$USER"/log/tmp.log
}

installVsCode() {
    init "Starting Visual Studio Code installation" "Visual Studio Code installation complete"

    checkIfIsInstalled "code"
    if [ "$?" -eq 1 ]; then
        information "Visual Studio Code is already installed!"
        return
    fi

    wget -q https://packages.microsoft.com/keys/microsoft.asc -O- | sudo apt-key add - 2>&1 | tee "$USER"/log/tmp.log
    sudo add-apt-repository "deb [arch=amd64] https://packages.microsoft.com/repos/vscode stable main" 2>&1 | tee "$USER"/log/tmp.log
    sudo apt-get update -y 2>&1 | tee "$USER"/log/tmp.log
    sudo apt-get install code -y
}

installGit() {
    init "Starting Git installation" "Git installation complete"

    checkIfIsInstalled "git"
    if [ "$?" -eq 1 ]; then
        information "Git is already installed!"
        return
    fi

    sudo apt-get install git -y 2>&1 | tee "$USER"/log/tmp.log

    configGit
}

configGit() {
    echo -ne "${GREEN}? ${YELLOW}GIT username: "
    read -r user

    echo -ne "${GREEN}? ${YELLOW}GIT email: "
    read -r email

    git config --global user.name "$user"
    git config --global user.email "$email"

    echo -ne "${YELLOW}Would you like to configure git credentials using rit? (y|n): "

    while true; do
        read -r var
        if [[ "$var" == "y" || "$var" == "Y" ]]; then
            rit set credential
            break
        elif [[ "$var" == "n" || "$var" == "N" ]]; then
            break
        else
            echo >&2 "Please, enter with 'y' or 'n'"
        fi
    done
}

installMaven() {
    init "Starting Maven installation" "Installation of Maven Complete"

    checkIfIsInstalled "mvn"
    if [ "$?" -eq 1 ]; then
        information "Maven is already installed!"
        return
    fi

    sudo apt-get install maven -y 2>&1 | tee "$USER"/log/tmp.log
}

installDocker() {
    init "Starting Docker installation" "Docker installation complete"

    checkIfIsInstalled "docker"
    if [ "$?" -eq 1 ]; then
        information "Docker is already installed!"
        return
    fi

    sudo apt-get install docker.io -y 2>&1 | tee "$USER"/log/tmp.log
    sudo systemctl start docker 2>&1 | tee "$USER"/log/tmp.log
    sudo systemctl enable docker 2>&1 | tee "$USER"/log/tmp.log
}

installIntellij() {
    init "Starting installation of IntelliJ IDEA"

    checkIfIsInstalled "idea"
    if [ "$?" -eq 1 ]; then
        information "IntelliJ IDEA is already installed!"
        return
    fi

    sudo apt-get update -y 2>&1 | tee "$USER"/log/tmp.log

    echo -ne "${YELLOW}Select the version of IntelliJ to be installed:\\n"

    options=("Community" "Ultimate" "Quit")

    select opt in "${options[@]}"; do
        case $opt in
        "Community")
            clearLine 1
            init "Starting installation of IntelliJ IDEA Community" "IntelliJ IDEA Community installation completed"
            sudo snap install intellij-idea-community --classic 2>&1 | tee "$USER"/log/tmp.log
            break
            ;;
        "Ultimate")
            clearLine 1
            init "Starting installation of IntelliJ IDEA" "IntelliJ IDEA Ultimate installation completed"
            sudo snap install intellij-idea-ultimate --classic 2>&1 | tee "$USER"/log/tmp.log
            break
            ;;
        "Quit")
            clearLine 1
            echo -e "${RED}[ ❌ ] Intellij not installed!"
            break
            ;;
        *) echo -e "${RED}Invalid option, try again!" ;;
        esac
    done
}

installPostman() {
    init "Starting Postman installation" "Postman installation complete"

    checkIfIsInstalled "postman"
    if [ "$?" -eq 1 ]; then
        information "Postman is already installed!"
        return
    fi

    sudo snap install postman 2>&1 | tee "$USER"/log/tmp.log
}

installMySql() {
    init "Starting MySql installation" "MySql Server installation complete"

    checkIfIsInstalled "mysql"
    if [ "$?" -eq 1 ]; then
        information "MySql is already installed!"
        return
    fi

    sudo apt-get install mysql-server mysql-client -y 2>&1 | tee "$USER"/log/tmp.log
}

installMySqlWorkbench() {
    init "Starting installation of MySql Workbench" "MySql Workbench Installation Complete"

    checkIfIsInstalled "mysql-workbench"
    if [ "$?" -eq 1 ]; then
        information "MySql Workbench is already installed!"
        return
    fi

    sudo apt-get install mysql-workbench -y 2>&1 | tee "$USER"/log/tmp.log
}

installChrome() {
    init "Starting installation of Google Chrome" "Google Chrome installation complete"

    checkIfIsInstalled "google-chrome"
    if [ "$?" -eq 1 ]; then
        information "Google Chrome is already installed!"
        return
    fi

    wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb 2>&1 | tee "$USER"/log/tmp.log
    sudo dpkg -i google-chrome-stable_current_amd64.deb 2>&1 | tee "$USER"/log/tmp.log
}

installSpotify() {
    init "Starting Spotify installation" "Spotify Installation Complete"

    checkIfIsInstalled "spotify"
    if [ "$?" -eq 1 ]; then
        information "Spotify is already installed!"
        return
    fi

    snap install spotify 2>&1 | tee "$USER"/log/tmp.log
}

installDiscord() {
    init "Starting Discord installation" "Discord Installation Complete"

    checkIfIsInstalled "discord"
    if [ "$?" -eq 1 ]; then
        information "Discord is already installed!"
        return
    fi

    sudo snap install discord 2>&1 | tee "$USER"/log/tmp.log
}

checkIfIsInstalled() {
    if command -v "$1"; then
        return 1
    else
        return 0
    fi
}

information() {
    clearLine 2
    echo -ne "\\r${LIME_YELLOW}[ ⚠️  ] $1\\n"
}

init() {
    FRAME=("⠋" "⠙" "⠹" "⠸" "⠼" "⠴" "⠦" "⠧" "⠇" "⠏")
    FRAME_INTERVAL=0.1

    local step=0

    tput civis -- invisible

    while [ "$step" -lt "${#CMDS[@]}" ]; do
        ${CMDS[$step]} &
        pid=$!

        while ps -p $pid &>/dev/null; do
            echo -ne "\\r${CYAN}[   ] $1 ..."

            for k in "${!FRAME[@]}"; do
                echo -ne "\\r${CYAN}[ ${FRAME[k]} ]"
                sleep $FRAME_INTERVAL
            done
        done

        tput dl1
        echo -ne "\\r${GREEN}[ ✅ ] $2\\n"
        step=$((step + 1))
    done

    tput cnorm -- normal
}

finishInstall() {
    sudo apt-get update && sudo apt-get dist-upgrade -y && sudo apt-get clean && sudo apt-get autoremove -y

    if mycmd; then
        echo -ne "${RED}Erro ao atualizar sistema!"
    else
        echo -ne "${GREEN}Installed programs."
        echo "Updated Repository."
        echo "Updated System."
        echo "...................."
        echo Press Enter to Continue
        read -r #pausa
        exit
    fi
}

createDirLog() {
    if [ ! -d log ]; then
        mkdir "$USER"/log
    else
        rm "$USER"/log/tmp.log
    fi
}

clearLine() {
    for ((c = 1; c <= "$1"; c++)); do
        tput cuu1
        tput dl1
    done
}
