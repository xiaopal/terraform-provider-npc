/*
# https://www.terraform.io/guides/writing-custom-terraform-providers.html?_ga=2.125223497.250066774.1494901881-1900920754.1467905211
# providers {
#  npc = "/path/to/privatecloud"
# }
*/

variable "namespace" {
  default = "default"
}

provider "npc" {
}

resource "npc_service" "sample-web" {
	count = 3
	spec = "C1M1S20"
	namespace = "pre"
}
