#!/usr/bin/ruby

require 'bundler/setup'
require_relative 'formula/formula'

INPUT1 = ENV["INPUT_TEXT"]
INPUT2 = ENV["INPUT_LIST"]
INPUT3 = ENV["INPUT_BOOLEAN"]
INPUT4 = ENV["INPUT_PASSWORD"]

Run(INPUT1, INPUT2, INPUT3, INPUT4)
