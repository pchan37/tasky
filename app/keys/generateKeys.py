import os


def generateAuthKey(size):
    with open('authKey', 'w') as file_:
        file_.write(os.urandom(size).encode('hex'))
    return


def generateEncryptKey(size):
    with open('encryptKey', 'w') as file_:
        file_.write(os.urandom(size).encode('hex'))
    return

def main():
    generateAuthKey(16)
    generateAuthKey(16)

if __name__ == '__main__':
    main()
