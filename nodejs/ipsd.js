// Author: d1y<chenhonzhou@gmail.com>

const pid = "2"

const calg = "12345678"

const createMD5 = data=> {
  var crypto = require('crypto')
  return crypto.createHash('md5').update(data).digest("hex")
}

const createPassword = p=> {
  const r = pid + p + calg
  const output = createMD5(r) + calg + pid
  return output
}

;(()=> {
  const args = process.argv
  if (args.length == 2) {
    console.log("请传入密码")
    return
  }
  const p = args[2]
  const output = createPassword(p)
  console.log(output)
})()