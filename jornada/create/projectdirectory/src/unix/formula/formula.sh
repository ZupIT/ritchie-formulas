#!/bin/sh

runFormula() {
 
  clear
  # valida se este usuário informado existe
  if ! [ -e "/Users/$RIT_INPUT_USER" ];then
        validation
        echoColor "red" "Usuário Informado não encontrado!"
        echoColor "blue" "Favor validar se o usuário informado está correto."
        endValidation
        exit
  fi

  echoColor "green" "Olá, $RIT_INPUT_USER. Seja Bem vindo!"
  echoColor "blue" "Vamos te auxiliar na configuração da sua estação de trabalho."
  
  echoColor "blue" "Percebemos que está no projeto $RIT_INPUT_PROJECT"
  echoColor "blue" "E que está utilizando um SO $RIT_INPUT_SO"


  if [ $RIT_INPUT_SO = "Mac" ];then
   
    if [ $RIT_INPUT_PROJECT = "Java" ];then
      createDirectory "/Users/$RIT_INPUT_USER/Java"
    fi

    if [ $RIT_INPUT_PROJECT = "Python" ];then
      createDirectory "/Users/$RIT_INPUT_USER/Python"
    fi

    if [ $RIT_INPUT_PROJECT = "Java e Phyton" ];then
      createDirectory "/Users/$RIT_INPUT_USER/Java"
      createDirectory "/Users/$RIT_INPUT_USER/Python"
    fi

  elif [ $RIT_INPUT_SO = "Windows" ];then
    
    if [ $RIT_INPUT_PROJECT = "Java" ];then
      createDirectory "C:\\Users\\$RIT_INPUT_USER\\Java"
    elif [ $RIT_INPUT_PROJECT = "Python" ];then
      createDirectory "C:\\Users\\$RIT_INPUT_USER\\Python"
    fi

  fi

  echoColor "cyan" "Seu ambiente foi configurado!"

}

validation() {
echo "==========  [ V A L I D A T I O N ]  =============="
}

endValidation() {
echo "==================================================="
}

createDirectory() {

 DIRETORIO=$1

#verifica se existe o diretorio
if ! [ -e "$DIRETORIO" ];then
  mkdir "$DIRETORIO"
else
  validation
  echo "O diretório informado já existe!"
  echo "Sua Máquina já estava configurada."
  endValidation
  exit
fi

}

echoColor() {
  case $1 in
    red)
      echo "$(printf '\033[31m')$2$(printf '\033[0m')"
      ;;
    green)
      echo "$(printf '\033[32m')$2$(printf '\033[0m')"
      ;;
    yellow)
      echo "$(printf '\033[33m')$2$(printf '\033[0m')"
      ;;
    blue)
      echo "$(printf '\033[34m')$2$(printf '\033[0m')"
      ;;
    cyan)
      echo "$(printf '\033[36m')$2$(printf '\033[0m')"
      ;;
  esac
}

