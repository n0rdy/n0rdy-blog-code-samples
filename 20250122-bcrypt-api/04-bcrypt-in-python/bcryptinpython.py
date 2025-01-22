import random
import string

import bcrypt

def random_string(length):
    return ''.join(random.choice(string.ascii_letters) for i in range(length))

if __name__ == '__main__':
    # 18 + 55 + 1 = 74, so above 72 characters' limit of BCrypt
    user_id = random_string(18)
    username = random_string(55)
    password = "super-duper-secure-password"

    combined_string = "{0}:{1}:{2}".format(user_id, username, password)

    combined_hash = bcrypt.hashpw(combined_string.encode('utf-8'), bcrypt.gensalt())

    # let's try to break it
    wrong_password = "wrong-password"
    wrong_combined_string = "{0}:{1}:{2}".format(user_id, username, wrong_password)

    if bcrypt.checkpw(wrong_combined_string.encode('utf-8'), combined_hash):
        print("Password is correct")
    else:
        print("Password is incorrect")