const axios = require('axios')

const start = (user, repo, collaborator, token) => {
    var config = {
        method: 'get',
        url: `https://api.github.com/repos/${user}/${repo}/collaborators`,
        headers: {
            'Authorization': `Bearer ${token}`
        }
    }
    axios(config)
        .then(() => {
            addCollab(user, repo, collaborator, token)
        })
        .catch((error) => {
            if (error.response.status == 401) console.log('Bad Credentials.')
            if (error.response.status == 404) console.log('Invalid User or Repository.')
        });
}
const addCollab = (user, repo, collaborator, token) => {
    var config = {
        method: 'put',
        url: `https://api.github.com/repos/${user}/${repo}/collaborators/${collaborator}`,
        headers: {
            'Authorization': `Bearer ${token}`,
        }
    }
    axios(config)
        .then((response) => {
            if (response.status === 204) throw console.log('User is already a collaborator.')
            console.log(`Collaborator: ${collaborator} was invited to repository: ${repo}`)
        })
        .catch((error) => {
            if (error != undefined) {
                if (error.response.status == 404) console.log('Collaborator username not found.')
            }
        })

}
module.exports.exec = start