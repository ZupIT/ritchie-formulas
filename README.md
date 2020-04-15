<img class="special-img-class" src="/docs/img/ritchie-banner.png" />

[Contribute to the Ritchie community](https://github.com/ZupIT/ritchie-formulas/blob/master/CONTRIBUTING.md)

## Documentation

This repository contains the community formulas which can be executed by the [ritchie-cli](https://github.com/ZupIT/ritchie-cli).

- [Gitbook](https://docs.ritchiecli.io)

- [Zup Products](https://www.zup.com.br/en/about) 

## Build and test formulas locally

Execute one of the following commands through the `ritchie-formulas/Makefile` file :

- `make test-local`: Build all formulas and add them to the `~/.rit/formulas` folder.

- `make test-local form={FORMULA_NAME}`: Build the specific formula and add it to the `~/.rit/formulas` folder.

**The formula(s) can then be tested locally through the terminal using the associated ritchie command.**

## Contribute to the Ritchie community with your formulas

1. Fork the repository
2. Create a branch: `git checkout -b <branch_name>`
3. Check the step by step of [how to create formulas on Ritchie](https://codelabs-preview.appspot.com/?file_id=1B3sNi3_btVWh80uZRZpIymcEi1c1SLCcZAH6-3WEVCc#0)
4. Add your formulas to the repository and commit your implementation: `git commit -m '<commit_message>'`
5. Push your branch: `git push origin <project_name>/<location>`
6. Open a pull request on the repository for analysis.
