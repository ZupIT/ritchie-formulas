Function runFormula () {
  $input_text = $env:INPUT_TEXT
  $input_boolean = $env:INPUT_BOOLEAN
  $input_list = $env:INPUT_LIST
  $input_password = $env:INPUT_PASSWORD

  Write-Host "Hello World!"
  Write-Host "My name is $input_text."
  if ([boolean]$input_boolean) {
    Write-Host "I've already created formulas using Ritchie."
  } else {
    Write-Host "I'm excited in creating new formulas using Ritchie."
  }
  Write-Host "Today, I want to automate $input_list."
  Write-Host "My secret is $input_password."
}
