# HttpSh 
**Remote execute scripts tool / 远程脚本执行工具**

## Usage of HttpSh

```shell
./httpSh_linux_amd64 -h

Usage of ./httpSh_mac_amd64:
  -c string
        --c {"Access to the address/访问地址 1":"Execute script absolute path/执行脚本绝对路径1","Access to the address/访问地址 2":"Execute script absolute path/执行脚本绝对路径 2"} (default "{}")
  -p string
        --p port/端口 (default "8088")
```

## E.G


```shell
// 1. creat shell scripts
echo -e '#!/bin/sh\necho "success"' > /data/test.sh
chmod +x /data/test.sh

// 2. run Httpsh
./httpSh_linux_amd64 -c '{"test":"/data/test.sh"}'

// 3. test
curl http://127.0.0.1:8088/test
success success
```