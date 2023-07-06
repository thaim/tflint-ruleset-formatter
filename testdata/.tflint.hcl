plugin "formatter" {
  enabled = true
  # remove source and version to use local installed plugin
  source = "github.com/thaim/tflint-ruleset-formatter"
  version = "0.2.3"

  signing_key = <<-KEY
-----BEGIN PGP PUBLIC KEY BLOCK-----

mQINBGSmsbQBEADUMY+SN9wYEWjcMdMreKJswLy9nvCeqSP41+T7MQxNS2Wq8K7N
mOoszNwAo33KLeLBIWUBWkNXwAGJcl3AcbnEhi9PvAgg1AUNiTyep4QhPlYEEAte
gtvF+9JuLIceKiQhPHlIlRV71YrIP7+3TXUiwpoeuLQHown9JYo9N4y69ILS/rnN
vnGGr135rPx8UVbY7XpR1Vy/FjQkGrOxL61qpgZIyn6y6mCKDRR4tjGFvQ5id3Vs
R5qNsRTSrR+1Ag5mpu5aOyTsI1EZuEYs97gZpMu+tflt03kwi8u8EGvilkHf5/Zs
L1uqhdIpnaWeYjUJBYnwdSJnF2PsOg5c4Btljw/Q2S9K9qvL3bmNXRMqaNBUnkEW
HR8koHk4ndy+06NAwqVv/ZhlSP+hOT0V+hQtNWPbJ2Vdfy5UO5WyWEDip2Eb+cZe
/YwmPQej5tVknCfalLNkvC8gqXuRft0XFFQ7BWM70pVVtlia64It3xJPUQ4aP3If
QkAUl5/E4faPFfP5CyFaHAQA+X0/HDvI4Wh9xayZJyOEqa3obWqExLFd5UCzwDOh
iaMwpmxMwuSRaA4fg7OBc3/5zwt5gVLcy3Os0JQ+8iXoOw91x8CPVBIQAqDtbxyu
k8QliB3NEJSPrVQ8TEdwuaKEHPblyfVC/ngX80TvbHaR3bRy20HvD2rTJQARAQAB
tBl0aGFpbSA8dGhhaW0yNEBnbWFpbC5jb20+iQJUBBMBCgA+FiEEAyVq0WHEQXpk
/tLGKelk/4gwtyYFAmSmsbQCGwMFCQPCZwAFCwkIBwIGFQoJCAsCBBYCAwECHgEC
F4AACgkQKelk/4gwtyY3Xw/8Chu97d57GUHLAr55wmTqH+lV3iCCw3xh7Hgj/Jj5
7dIsjyKaSso0TJIgHLoCNmWn8O5yrAUBJ7vETDuag6VsJTm2onWknu8u1ds4FKLq
ERyMERTW2cR8kH/cXPi22/ht9/JPTGJQHBlURYQKOkTx1SJCLihpQtke6if3xk5h
QM0oL6TCm2piVunvX0p0SZIUt13a3gvnJColcO+F6G4/F5E+Io1KBEOKL8i6tdhz
ORU2LrMwtxY/PbyQDqUq5Wu9AKqgPYWJ50aDRzNQEkvWopBNh9tAX6V0yhyiA2tx
gEOfOotBntvnwaW6lrfk0TuVdQanGdBfzOGGWyscYq7ewGAY/mre4dsUYmDuRo1L
PmbnEGrFrdB+yGznr6bmoY5PQPeSxh4o1D4S22XhJIRqVf1wXFTTDoQLmnsXaM0x
rZUjnKxA+JV3zMTElQor2l0TJyYtQQUrJctnU++HENppgvhJ7+wHIlbH6SgJAc6T
cu7jJtrOcT9NwSLCWkGMFs2Ia58NEf2u1xRbG/6/wyhbCYUMPpiRm7QJ5D1CPbUs
fPv2A6N2surFlPWgVwoh39ccxHsjl9mp14Shb9eJlL9wkLhPrkNfkenHKq+m44ZT
jBPJ/5UvjFcGHd4SV19sUU9PgL/cbKxHhKwEGRUbqkKtNFieAz/NZdgJomWsFD8u
LvA=
=+9jm
-----END PGP PUBLIC KEY BLOCK-----
KEY
}

rule "formatter_max_len" {
  enabled = true
  length = 80
}
