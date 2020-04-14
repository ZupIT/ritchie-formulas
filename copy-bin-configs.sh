#!/bin/sh

FORMULAS="$1"

create_formulas_dir() {
  mkdir -p formulas/"$formula"
}

find_config_files() {
  files=$(find "$formula" -type f -name "*config.json")
}

copy_config_files() {
  for file in $files; do
    cp "$file" formulas/"$formula"
  done
}

copy_formula_bin() {
  cp -rf "$formula"/dist formulas/"$formula"
}

rm_formula_bin() {
  rm -rf "$formula"/dist
}

create_formula_checksum() {
  find "${formula}"/dist -type f -exec md5sum {} \; | sort -k 2 | md5sum | cut -f1 -d ' ' > formulas/"${formula}.md5"
}


compact_formula_bin_and_remove_them() {
  for bin_dir in `find formulas/"$formula" -type d -name "dist"`; do
    for binary in `ls -1 $bin_dir`; do
      cd  ${bin_dir}/${binary}
      zip -r "${binary}.zip" "bin"
      mv "${binary}".zip ../../
      cd -
      rm -rf "${bin_dir}"
    done;
  done
}


init() {
  for formula in $FORMULAS; do
    create_formulas_dir
    find_config_files
    copy_config_files
    create_formula_checksum
    copy_formula_bin
    rm_formula_bin
    compact_formula_bin_and_remove_them
  done
}

init
