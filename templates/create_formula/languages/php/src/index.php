<?php
require __DIR__ . '/vendor/autoload.php';
include 'formula/formula.php';

$input1 = getenv('INPUT_TEXT');
$input2 = getenv('INPUT_BOOLEAN');
$input3 = getenv('INPUT_LIST');
$input4 = getenv('INPUT_PASSWORD');
Run($input1, $input2, $input3, $input4);
?>
