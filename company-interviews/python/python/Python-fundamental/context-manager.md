A context manager is a Python construct that manages setup and teardown of resources automatically, 
typically used with the with statement. It ensures proper cleanup, making code safer and cleaner.


class MyCustomeContextManager:
    def __enter__(self):
        print("Enter Context Manager")
        return "Return Resource"
    def __exit__(self, exc_type, exc, tb):
        print("Resource Exits")
    
with MyCustomeContextManager() as r:
    print(r)


    from contextlib import contextmanager

@contextmanager
def myContextManager():
    try:
        a = 1//1
        print(a)
        yield "Resource working"
    except Exception as e:
        print("Exception happened.. ",e)
    finally:
        print("Resource clean up")
        
with myContextManager() as r:
    print(r)




import sqlite3
from contextlib import contextmanager

@contextmanager
def DB_CONNECTION(db_name):
    db = sqlite3.connect(db_name)  # Set Up Connection
    try:
        yield db                   # Give Connection 
        db.commit()                
    except Exception as e:
        db.rollback()
        print(e)
    finally:
        print("Connection Closed")
        db.close()
        
with DB_CONNECTION("test.db") as cx:
    cx = cx.cursor()
    for row in cx.execute("select * from users"):
        print(row)