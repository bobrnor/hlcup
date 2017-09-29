
import sys
import re

if len(sys.argv) != 2:
    print('format: ./prepare.py <file.answ>')
    exit(1)

f0 = open(sys.argv[1], 'r')
f0_lines = []
for line in f0:
    f0_lines.append(re.sub(r'\d+ \d+ \d+ \d+ \d+\n(GET|POST) (\S+) HTTP/1\.1\nHost: travels\.com\nUser-Agent: Technolab/1\.0 \(Docker; CentOS\) Highload/1\.0\nAccept: \*/\*\nConnection: keep-alive\n\n\nHTTP/1\.1 (\d+).+\nServer: fasthttp\nDate: .+\n(Content-Type: text/plain; charset=utf-8\nContent-Length: \d+\n|Content-Length: \d+\n)\n(.*)\n', r'\1\t\2\t\3\t\5\n', line))

print(f0_lines)

