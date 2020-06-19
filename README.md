# GoExpandIPRanges

Golang script to expand network ranges in following formats to individual IP ranges: 
* Network masks (`network/<mask>`)

Based on work done by `kotakanbe` and written here: https://gist.github.com/kotakanbe/d3059af990252ba89a82

IP addresses may not be in order as IP expansion occurs via goroutines

## Examples

To expand IP masks provided in file `/tmp/ranges.txt` to individual IPs, run the command:
```
$ cat /tmp/ranges.txt
1.1.1.0/24
2.2.2.0/24

$ cat /tmp/ranges.txt | goexpandipranges
1.1.1.1
1.1.1.2
...
2.2.2.254
```
