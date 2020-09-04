Function runFormula () {
  $sample_text = $env:SAMPLE_TEXT
  $sample_list = $env:SAMPLE_LIST
  $sample_bool = $env:SAMPLE_BOOL

  Write-Host "Hello, World!"
  Write-Host "You receive $sample_text in text. "
  Write-Host "You receive $sample_list in list. "
  Write-Host "You receive $sample_bool in boolean. "
}