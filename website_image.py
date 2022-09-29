import pymysql
import base64,io
from PIL import Image


# link mysql db
def creatdb(): 
    db = pymysql.connect(
        host='localhost',
        port=33777,
        user='root',
        password='cat020605',
        database='lookcat'
    )
    return db

# image trans base64
def toBase64(): 
    img_path = "./test.jpg"

def main() :
    with open("./test.jpg","rb") as f: 
        img_d = f.read()
        base64_d = base64.b64decode(img_d)
        xx = base64.decode('utf-8')
        print(xx)
        print(type(base64_d))
    

if __name__ == '__main__':
    main()
