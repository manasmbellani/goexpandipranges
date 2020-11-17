# GoExpandIPRanges

Golang script to expand network ranges in following formats to individual IP ranges: 
* Network masks (`network/<mask>`).

The script can also optionally exclude the network address (e.g. `192.168.1.0` in `192.168.1.0/24`) and broadcast address (e.g. `192.168.1.255` in `192.168.1.0/24`) through flags when executing the script.

Based on work done by `kotakanbe` and written here: https://gist.github.com/kotakanbe/d3059af990252ba89a82

IP addresses may not be in order as IP expansion occurs via goroutines. 

Use `sort` command to sort the output results.

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

To increase the number of threads from default (20), run the command with `-t` flag:
```
$ cat /tmp/ranges.txt
1.1.1.0/24
2.2.2.0/24

$ cat /tmp/ranges.txt | goexpandipranges -t 50
1.1.1.1
1.1.1.2
...
2.2.2.254
``` 

To remove the network addresses and broadcast addresses from network, use flag `-en` and `-eb` respectively:
```
$ cat /tmp/ranges.txt
1.1.1.0/32
2.2.2.0/32

$ cat /tmp/ranges.txt | goexpandipranges -t 50 -en -eb
1.1.1.1
1.1.1.2
2.2.2.1
2.2.2.2
```