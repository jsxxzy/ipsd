# Author: d1y<chenhonzhou@gmail.com>

import sys, hashlib

pid = "2"

calg = "12345678"

def createMD5(str):
  hl = hashlib.md5()
  hl.update(str.encode(encoding='utf-8'))
  hex = hl.hexdigest()
  return hex

def createPassword(str):
  r = pid + str + calg
  output = createMD5(r) + calg + pid
  return output

if __name__ == "__main__":
  args = sys.argv
  if len(args) == 1:
    print("请传入密码")
  else:
    p = args[1]
    output = createPassword(p)
    print(output)