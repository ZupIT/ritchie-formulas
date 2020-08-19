const fs = require('fs');
const util = require('util');
const exec = util.promisify(require('child_process').exec);
const readline = require('readline');
const xml2js = require('xml2js');

const parser = new xml2js.Parser();
const builder = new xml2js.Builder();

async function checkXcodeOpened() {
  try {
    await exec('ps aux | grep -v grep | grep -c Xcode');
    return true;
  } catch (err) {
    return false;
  }
}

function readInput(text) {
  const reader = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
  });

  return new Promise((resolve, _) => {
    reader.question(text, async (answer) => {
      reader.close();
      resolve(answer);
    });
  });
}

async function closeXcode() {
  try {
    console.log('To execute this formula you need to close Xcode.');
    const answer = await readInput('Do you want to close Xcode? [y/n]: ');

    if (answer.toLowerCase() === 'y') {
      console.log('Closing Xcode...');
      await exec('pkill Xcode');
      return true;
    } else if (answer.toLowerCase() === 'n') {
      console.log('Please close XCode to execute the formula');
      return false;
    }

    throw new Error("Please enter with a valid string: 'y' or 'n'");
  } catch (err) {
    throw err;
  }
}

async function getProjectName() {
  const readdir = util.promisify(fs.readdir);

  try {
    const files = await readdir(process.env.CURRENT_PWD);

    for (let i = 0; i < files.length; i++) {
      if (files[i].includes('.xcodeproj')) {
        return files[i].replace('.xcodeproj', '');
      }
    }
  } catch (err) {
    throw err;
  }

  throw new Error('.xcodeproj file not found, please cd to your project');
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
            "# Type a script or drag a script file from your workspace to insert its path.\nfind ${BUILT_PRODUCTS_DIR} -name '" +
            projectName +
            ".framework' -exec rm -rf {} \\;",
        },
      },
    },
  };
}

async function createXCSharedData(projectName, dirPath, writeFile, readFile) {
  const mkdir = util.promisify(fs.mkdir);

  await mkdir(dirPath, { recursive: true });

  let template = await readFile('./formula/template.xcscheme', 'utf8');

  template = template.replace(/\${projectName}/g, projectName);

  await writeFile(`${dirPath}/${projectName}.xcscheme`, template);
}

async function Run() {
  const readFile = util.promisify(fs.readFile);
  const fileExists = util.promisify(fs.exists);
  const writeFile = util.promisify(fs.writeFile);

  try {
    if (await checkXcodeOpened()) {
      const success = await closeXcode();
      if (!success) {
        return;
      }
    }

    const projectName = await getProjectName();
    const path = `${process.env.CURRENT_PWD}/${projectName}.xcodeproj/xcshareddata/xcschemes/${projectName}.xcscheme`;

    if (await fileExists(path)) {
      const file = await readFile(path);

      const parsedXML = await parser.parseStringPromise(file);

      if ('PreActions' in parsedXML.Scheme.BuildAction[0]) {
        parsedXML.Scheme.BuildAction[0].PreActions[0].ExecutionAction.push(
          getPreAction(projectName).ExecutionAction
        );
      } else {
        parsedXML.Scheme.BuildAction[0].PreActions = getPreAction(projectName);
      }

      const xml = builder.buildObject(parsedXML);

      await writeFile(path, xml);
    } else {
      const dirPath = `${process.env.CURRENT_PWD}/${projectName}.xcodeproj/xcshareddata/xcschemes`;
      await createXCSharedData(projectName, dirPath, writeFile, readFile);
    }

    console.log('Script successfully added!!');
  } catch (err) {
    console.log(err.message);
    process.exit(1);
  }
}

const formula = Run;
module.exports = formula;
