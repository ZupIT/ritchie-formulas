const run = require("./collaborator/collaborator")

//const permission = process.env.permission_list
const collaborator = process.env.COLLABORATOR_USER
const repository = process.env.REPOSITORY_NAME
const username = process.env.OWNER_USERNAME
const token = process.env.OWNER_TOKEN


//console.log(username, repository, collaborator, token)

run.exec (username, repository, collaborator, token)