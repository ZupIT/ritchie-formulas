const fs = require('fs');
const util = require('util');
const xml2js = require('xml2js');

const parser = new xml2js.Parser();
const builder = new xml2js.Builder();

async function parse(data) {
  try {
    const res = await parser.parseStringPromise(data);
    return res;
  } catch (err) {
    console.log(err);
    process.exit(1);
  }
}

function getPreAction(projectName) {
  return {
    ExecutionAction: {
      $: {
        ActionType:
          'Xcode.IDEStandardExecutionActionsCore.ExecutionActionType.ShellScriptAction',
      },
      ActionContent: {
        $: {
          title: 'Run Script',
          scriptText:
            "# Type a script or drag a script file from your workspace to insert its path.;find ${BUILT_PRODUCTS_DIR} -name '" +
            projectName +
            ".framework' -exec rm -rf {} \\;",
        },
      },
    },
  };
}

async function Run(projectName) {
  const readFile = util.promisify(fs.readFile);
  const writeFile = util.promisify(fs.writeFile);
  const path = `${process.env.CURRENT_PWD}/${projectName}.xcodeproj/xcshareddata/xcschemes/${projectName}.xcscheme`;

  try {
    const file = await readFile(path);

    const parsedXML = await parse(file);

    if ('PreActions' in parsedXML.Scheme.BuildAction[0]) {
      parsedXML.Scheme.BuildAction[0].PreActions.push(
        getPreAction(projectName)
      );
    } else {
      parsedXML.Scheme.BuildAction[0].PreActions = getPreAction(projectName);
    }

    const xml = builder.buildObject(parsedXML);

    await writeFile(path, xml);

    console.log('Script successfully added!!');
  } catch (err) {
    console.log(err);
    process.exit(1);
  }
}

const formula = Run;
module.exports = formula;
