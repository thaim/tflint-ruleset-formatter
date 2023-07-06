plugin "formatter" {
  enabled = true
  # remove source and version to use local installed plugin
  source = "github.com/thaim/tflint-ruleset-formatter"
  version = "0.2.2"
}

rule "formatter_max_len" {
  enabled = true
  length = 80
}
