const bcrypt = require('bcryptjs')

function randomString (length) {
  const chars = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'
  let result = ''
  for (let i = length; i > 0; --i) {
    result += chars[Math.floor(Math.random() * chars.length)]
  }
  return result
}

function runTest () {
  // 18 + 55 + 1 = 74, so above 72 characters' limit of BCrypt
  const userId = randomString(18)
  const username = randomString(55)
  const password = 'super-duper-secure-password'

  const combinedString = `${userId}:${username}:${password}`

  const combinedHash = bcrypt.hashSync(combinedString)

  // let's try to break it
  const wrongPassword = 'wrong-password'
  const wrongCombinedString = `${userId}:${username}:${wrongPassword}`

  if (bcrypt.compareSync(wrongCombinedString, combinedHash)) {
    console.log('Password is correct')
  } else {
    console.log('Password is wrong')
  }
}

runTest()
