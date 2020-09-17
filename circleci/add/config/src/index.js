const run = require("./formula/formula");

run(
    process.env.PROJECT_NAME,
    process.env.DATABASE,
    process.env.REGISTRY,
    process.env.STAGING_DEPLOY,
    process.env.PRODUCTION_DEPLOY
);
