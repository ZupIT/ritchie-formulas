const fs = require("fs");
const yaml = require("yaml");

function Run(project_name, database, registry, staging_deploy, production_deploy) {
    const workflow_name = 'build-deploy';
    const config = fs.readFileSync(__dirname + '/config.yml', 'utf8');
    const templates = fs.readFileSync(__dirname + '/templates.yml', 'utf8');
    const config_parsed = yaml.parse(config);
    const templates_parsed = yaml.parse(templates);

    if (database === "mongo") {
        config_parsed.jobs.test.docker.push({image: 'mongo:4.0'})
    } else if (database === "postgres") {
        config_parsed.jobs.test.docker.push({
            image: 'postgres:12.4-alpine',
            environment: {
                POSTGRES_PASSWORD: 'admin',
                POSTGRES_USER: 'postgres',
                POSTGRES_DB: 'postgres'
            }
        })
    }

    let last_job = 'test';
    if (registry !== "No Registry") {
        const registry_job = registry.toLowerCase() + "-push";
        config_parsed.jobs[registry_job] = templates_parsed.registries[registry];
        config_parsed.workflows[workflow_name].jobs.push({
            [registry_job]: {
                requires: ['test'],
                filters: { branches: {only: 'master'}}
            }
        });
        last_job = registry_job;
    }

    if (staging_deploy === 'approval') {
        config_parsed.workflows[workflow_name].jobs.push({
            'approve-staging': {
                ...templates_parsed.deploys[staging_deploy],
                requires: last_job
            }
        });
        last_job = 'approve-staging';
    }
    config_parsed.workflows[workflow_name].jobs.push({
        'deploy-staging': {
            ...templates_parsed.deploys[staging_deploy],
            requires: last_job
        }
    });

    last_job = 'deploy-staging';
    if (production_deploy === 'approval') {
        config_parsed.workflows[workflow_name].jobs.push({
            'approve-production': {
                ...templates_parsed.deploys['approval'],
                requires: last_job
            }
        });
        last_job = 'approve-production';
    }
    config_parsed.workflows[workflow_name].jobs.push({
        'deploy-production': {
            ...templates_parsed.deploys[production_deploy],
            requires: last_job
        }
    });

    const yaml_result = yaml.stringify(config_parsed);
    fs.writeFileSync(__dirname + "/new-config.yml", yaml_result)
}

const formula = Run;
module.exports = formula;
