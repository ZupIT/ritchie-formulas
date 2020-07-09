const run = require("./collaborator/collaborator")

const collaborator = process.env.COLLABORATOR_USER
const repository = process.env.REPOSITORY_NAME
const username = process.env.OWNER_USERNAME
const token = process.env.OWNER_TOKEN

run.exec (username, repository, collaborator, token)