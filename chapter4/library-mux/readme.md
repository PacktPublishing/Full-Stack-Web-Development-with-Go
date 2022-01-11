# library-mux

This is a pure default Go library implementation of a web server rendering "hello world"

## Getting Started

Running is as simple as:

```bash
SERVER_PORT=9004 go run main.go 
2022/01/01 14:16:40.233798 main.go:13: GET
2022/01/01 14:16:40.234017 main.go:14: /
2022/01/01 14:16:40.234098 main.go:15: map[Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9] Accept-Encoding:[gzip, deflate, br] Accept-Language:[en-GB,en-US;q=0.9,en;q=0.8] Cache-Control:[max-age=0] Connection:[keep-alive] Dnt:[1] Sec-Ch-Ua:[" Not A;Brand";v="99", "Chromium";v="96", "Google Chrome";v="96"] Sec-Ch-Ua-Mobile:[?0] Sec-Ch-Ua-Platform:["macOS"] Sec-Fetch-Dest:[document] Sec-Fetch-Mode:[navigate] Sec-Fetch-Site:[none] Sec-Fetch-User:[?1] Upgrade-Insecure-Requests:[1] User-Agent:[Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36]]
...
```

You can visit `http://0.0.0.0:9004/` in your browser to perform a GET request.
