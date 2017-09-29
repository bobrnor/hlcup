
import sys
import re

r = re.compile(r'\d+ \d+ \d+ \d+ \d+\n(GET|POST) (\S+) HTTP/1\.1\nHost: travels\.com\nUser-Agent: Technolab/1\.0 \(Docker; CentOS\) Highload/1\.0\nAccept: \*/\*\nConnection: keep-alive\n\n\nHTTP/1\.1 (\d+).+\nServer: fasthttp\nDate: .+\n(Content-Type: text/plain; charset=utf-8\nContent-Length: \d+\n|Content-Length: \d+\n)\n(.*)\n', re.MULTILINE)

if len(sys.argv) != 3:
    print('format: ./prepare.py <in.answ> <out.answ>')
    exit(1)

in_file = open(sys.argv[1], 'r')
content = in_file.read()
in_file.close()

for match in r.finditer(content):
    print(match)

# content = r.sub(r'\1\t\2\t\3\t\5\n', content)
# print(content)

out_file = open(sys.argv[2], 'w')
out_file.write(content)
out_file.close()
