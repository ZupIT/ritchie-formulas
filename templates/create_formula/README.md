# How to add some language on Rit create formula

- Create a folder on languages folder
- The folders on language's folder will be presented to user on `rit create formula`
- Inside the created folder you should have:
    - Makefile, this file should do:
        - create run.sh and run.bat, these files should explain how to run the formula, run.sh is called on linux and mac system and build.bat will be called on windows system
        - copy files that run.sh and run.bat needs to run
        - remember makefile can be called inside a docker using the dockerImageBuilder as the docker image name
    - config.json, with this file rit can ask the inputs and use dockerImageBuilder to build the formula
    - src folder, on this folder you can create a simple formula using the language that you will add.
    - Dockerfile, this file will be use when --docker is pass to the formula, the objective is to run the formula inside the docker, so in this file you need to create a dockerfile that can run any formula of this language
    