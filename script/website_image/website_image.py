import pymysql
import os

def write_file(data, filename):
    # Convert binary data to proper format and write it on Hard Disk
    with open(filename, 'wb') as file:
        file.write(data)

class SqlOpt:
    addImage = "insert into website_images(name_image,data_image,type_image) values (%s,%s,%s)"
    downImage = "select * from website_images;"


    def __init__(self,values ):
        self.db = connectDB()
        self.cursor = self.db.cursor()
        self.values = values
    
    def addSql(self):
        self.cursor.execute(self.addImage,self.values)
        print("ee",self.cursor.rowcount)

        self.db.commit()
        # self.db.close()
    def readBLOB(self):
        self.cursor.execute(self.downImage)
        record = self.cursor.fetchall()
        for row in record: 
            img = row[2]
            nameimg = row[1]
            typeimg = row[-1]
            write_file(img,nameimg + typeimg)


# link mysql db
def connectDB(): 
    db = pymysql.connect(
        host='localhost',
        port=33777,
        user='root',
        password='cat020605',
        database='lookcat'
    )
    return db

# file = open('imgBase64.txt', 'wb')
# file.write(encodeImg)
# file.close()

def convertToBinaryData():
    with open("./test.jpg","rb") as f:
        binaryData = f.read()
        
    return binaryData

# get file to name and ext
def fileExt(filePath): 
    basename = os.path.basename(filePath)
    ( file, ext ) = os.path.splitext(basename)
    print(file,ext)

    
def main():
    # x = convertToBinaryData()
    # v = ('xx',x,'jpg')
    # SqlOpt(v).readBLOB()
    fileExt('./test.jpg')

if __name__ == '__main__':
    main()

