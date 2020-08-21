#!/bin/sh

USER=$(pwd)

declare -rx CMDS=(
    'sleep 1'
)

BLACK=$(tput setaf 0)
RED=$(tput setaf 1)
GREEN=$(tput setaf 2)
YELLOW=$(tput setaf 3)
LIME_YELLOW=$(tput setaf 190)
POWDER_BLUE=$(tput setaf 153)
BLUE=$(tput setaf 4)
MAGENTA=$(tput setaf 5)
CYAN=$(tput setaf 6)
WHITE=$(tput setaf 7)
BRIGHT=$(tput bold)
NORMAL=$(tput sgr0)
BLINK=$(tput blink)
REVERSE=$(tput smso)
UNDERLINE=$(tput smul)

runConfigLinux() {
    sayHello $SYSTEM
    createDirLog

    case $CONFIGURATION in
    "Complete configuration")
        completeConfiguration $GIT_NAME $GIT_EMAIL
        ;;
    "Select programs")
        selectConfiguration $GIT_NAME $GIT_EMAIL
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
    installGit $GIT_NAME $GIT_EMAIL
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
    start "Starting manual installation"
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

    while opt=$(zenity --title="$title" --text="$prompt" --list \
        --column="Tools" "${selections[@]}"); do

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
        "${selections[10]}") installGit $GIT_NAME $GIT_EMAIL ;;
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
    if [ $? -eq 0 ]; then
        sudo apt-get install lolcat -y 2>&1 | tee $USER/log/tmp.log
    fi

    checkIfIsInstalled "cowsay"
    if [ $? -eq 0 ]; then
        sudo apt-get install cowsay -y 2>&1 | tee $USER/log/tmp.log
    fi

    cowsay -f gnu Setting up environment $1 with RITCHIE! 🦸🚀 | lolcat

    printf "${CYAN}OS version: %s"
    lsb_release -r
}

installSPropertiesCommon() {
    start "Starting installation of Properties Common" "Properties Common installation complete"

    sudo apt-get install software-properties-common -y 2>&1 | tee $USER/log/tmp.log
}

installWget() {
    start "Starting installation of GNU Wget" "GNU Wget installation completed"

    checkIfIsInstalled "wget"
    if [ $? -eq 1 ]; then
        information "GNU Wget is already installed!"
        return
    fi

    sudo apt-get install apt-transport-https wget -y 2>&1 | tee $USER/log/tmp.log
}

installCertificates() {
    start "Starting installation of Certificates" "Certificate Installation Complete"

    sudo apt-get install ca-certificates -y 2>&1 | tee $USER/log/tmp.log
}

installCurl() {
    start "Starting installation of Curl Wget" "Curl Installation Complete"

    checkIfIsInstalled "curl"
    if [ $? -eq 1 ]; then
        information "Curl is already installed!"
        return
    fi

    sudo apt-get install curl -y 2>&1 | tee $USER/log/tmp.log
}

installGnupgAgent() {
    start "Starting installation of the GNUPG Agent" "GNUPG installation completed"

    sudo apt-get install gnupg-agent -y 2>&1 | tee $USER/log/tmp.log
}

installSnap() {
    start "Starting Snap installation" "Snap Installation Complete"

    checkIfIsInstalled "snap"
    if [ $? -eq 1 ]; then
        information "Snap is already installed!"
        return
    fi

    sudo apt-get install snapd -y 2>&1 | tee $USER/log/tmp.log
}

installJdk() {
    start "Starting JDK installation"
    clearLine 1

    sudo add-apt-repository ppa:linuxuprising/java -y

    sudo apt-get update -y 2>&1 | tee $USER/log/tmp.log

    echo -ne "${YELLOW}Select the version of Openjdk to be installed:"

    options=("8" "11" "12" "Quit")

    select opt in "${options[@]}"; do
        case $opt in
        "8")
            start "Starting JDK 8 installation" "JDK 8 installation completed"
            sudo apt-get install openjdk-8-jdk -y 2>&1 | tee $USER/log/tmp.log
            break
            ;;
        "11")
            start "Starting JDK 11 installation" "JDK 11 installation completed"
            sudo apt-get install openjdk-11-jdk -y 2>&1 | tee $USER/log/tmp.log
            break
            ;;
        "12")
            start "Starting JDK 12 installation" "JDK 12 installation completed"
            sudo apt-get install openjdk-12-jdk -y 2>&1 | tee $USER/log/tmp.log
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
    start "Starting Python3 installation" "Python3 installation completed"

    checkIfIsInstalled "python3"
    if [ $? -eq 1 ]; then
        information "Python3 is already installed!"
        return
    fi

    sudo add-apt-repository ppa:deadsnakes/ppa -y 2>&1 | tee $USER/log/tmp.log
    sudo apt-get install python3.8 -y 2>&1 | tee $USER/log/tmp.log
}

installNode() {
    start "Starting Node installation" "Node installation complete"

    checkIfIsInstalled "node"
    if [ $? -eq 1 ]; then
        information "Node is already installed!"
        return
    fi

    sudo apt-get install nodejs -y 2>&1 | tee $USER/log/tmp.log
}

installNpm() {
    start "Starting Npm installation" "Npm Installation Complete"

    checkIfIsInstalled "npm"
    if [ $? -eq 1 ]; then
        information "Npm is already installed!"
        return
    fi

    sudo apt-get install npm -y 2>&1 | tee $USER/log/tmp.log
}

installVsCode() {
    start "Starting Visual Studio Code installation" "Visual Studio Code installation complete"

    checkIfIsInstalled "code"
    if [ $? -eq 1 ]; then
        information "Visual Studio Code is already installed!"
        return
    fi

    wget -q https://packages.microsoft.com/keys/microsoft.asc -O- | sudo apt-key add - 2>&1 | tee $USER/log/tmp.log
    sudo add-apt-repository "deb [arch=amd64] https://packages.microsoft.com/repos/vscode stable main" 2>&1 | tee $USER/log/tmp.log
    sudo apt-get update -y 2>&1 | tee $USER/log/tmp.log
    sudo apt-get install code -y
}

installGit() {
    start "Starting Git installation" "Git installation complete"

    checkIfIsInstalled "git"
    if [ $? -eq 1 ]; then
        information "Git is already installed!"
        return
    fi

    sudo apt-get install git -y 2>&1 | tee $USER/log/tmp.log

    git config --global user.name "$1"
    git config --global user.email "$2"
}

installMaven() {
    start "Starting Maven installation" "Installation of Maven Complete"

    checkIfIsInstalled "mvn"
    if [ $? -eq 1 ]; then
        information "Maven is already installed!"
        return
    fi

    sudo apt-get install maven -y 2>&1 | tee $USER/log/tmp.log
}

installDocker() {
    start "Starting Docker installation" "Docker installation complete"

    checkIfIsInstalled "docker"
    if [ $? -eq 1 ]; then
        information "Docker is already installed!"
        return
    fi

    sudo apt-get install docker.io -y 2>&1 | tee $USER/log/tmp.log
    sudo systemctl start docker 2>&1 | tee $USER/log/tmp.log
    sudo systemctl enable docker 2>&1 | tee $USER/log/tmp.log
}

installIntellij() {
    start "Starting installation of IntelliJ IDEA"

    checkIfIsInstalled "idea"
    if [ $? -eq 1 ]; then
        information "IntelliJ IDEA is already installed!"
        return
    fi

    Sudo apt-add-repository ppa:mmk2410/intellij-idea -y 2>&1 | tee $USER/log/tmp.log
    sudo apt-get update -y 2>&1 | tee $USER/log/tmp.log

    echo -ne "${YELLOW}Select the version of IntelliJ to be installed:"

    options=("Community" "Ultimate" "Quit")

    select opt in "${options[@]}"; do
        case $opt in
        "Community")
            clearLine 1
            start "Starting installation of IntelliJ IDEA Community" "IntelliJ IDEA Community installation completed"
            sudo apt-get install intellij-idea-community -y 2>&1 | tee $USER/log/tmp.log
            break
            ;;
        "Ultimate")
            clearLine 1
            start "Starting installation of IntelliJ IDEA" "IntelliJ IDEA Ultimate installation completed"
            sudo apt-get install intellij-idea-ultimate -y 2>&1 | tee $USER/log/tmp.log
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
    start "Starting Postman installation" "Postman installation complete"

    checkIfIsInstalled "postman"
    if [ $? -eq 1 ]; then
        information "Postman is already installed!"
        return
    fi

    sudo snap install postman 2>&1 | tee $USER/log/tmp.log
}

installMySql() {
    start "Starting MySql installation" "MySql Server installation complete"

    checkIfIsInstalled "mysql"
    if [ $? -eq 1 ]; then
        information "MySql is already installed!"
        return
    fi

    sudo apt-get install mysql-server mysql-client -y 2>&1 | tee $USER/log/tmp.log
}

installMySqlWorkbench() {
    start "Starting installation of MySql Workbench" "MySql Workbench Installation Complete"

    checkIfIsInstalled "mysql-workbench"
    if [ $? -eq 1 ]; then
        information "MySql Workbench is already installed!"
        return
    fi

    sudo apt-get install mysql-workbench -y 2>&1 | tee $USER/log/tmp.log
}

installChrome() {
    start "Starting installation of Google Chrome" "Google Chrome installation complete"

    checkIfIsInstalled "google-chrome"
    if [ $? -eq 1 ]; then
        information "Google Chrome is already installed!"
        return
    fi

    wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb 2>&1 | tee $USER/log/tmp.log
    sudo dpkg -i google-chrome-stable_current_amd64.deb 2>&1 | tee $USER/log/tmp.log
}

installSpotify() {
    start "Starting Spotify installation" "Spotify Installation Complete"

    checkIfIsInstalled "spotify"
    if [ $? -eq 1 ]; then
        information "Spotify is already installed!"
        return
    fi

    snap install spotify 2>&1 | tee $USER/log/tmp.log
}

installDiscord() {
    start "Starting Discord installation" "Discord Installation Complete"

    checkIfIsInstalled "discord"
    if [ $? -eq 1 ]; then
        information "Discord is already installed!"
        return
    fi

    sudo snap install discord 2>&1 | tee $USER/log/tmp.log
}

checkIfIsInstalled() {
    if which -a $1; then
        return 1
    else
        return 0
    fi
}

information() {
    clearLine 3
    echo -ne "\\r${LIME_YELLOW}[ ⚠️  ] $1\n"
}

start() {
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

    if [ $? -ne 0 ]; then
        echo -ne "${RED}Erro ao atualizar sistema!"
    else
        echo -ne "${GREEN}Installed programs."
        echo "Updated Repository."
        echo "Updated System."
        echo "...................."
        echo Press Enter to Continue
        read #pausa
        exit
    fi
}

createDirLog() {
    if [ ! -d log ]; then
        mkdir $USER/log
    else
        rm $USER/log/tmp.log
    fi
}

clearLine() {
    for counter in $(seq 1 $1); do
        tput cuu1
        tput dl1
    done
}
