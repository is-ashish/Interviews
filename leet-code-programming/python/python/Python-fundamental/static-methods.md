class A:
    def __init__(self, name):
        self.name = name
        
    def get_Name(self):
        return self.name
    
    @staticmethod
    # Belongs to the class itself, not to any instance
    # Do not take self (instance) or cls (class) as the first argument
    # Useful for utility/helper functions that logically relate to the class but don’t need access to instance 
    def is_valid_name(n) -> bool:
        """check if the name is valid - non empty and non integer"""
        return n != "" and n.replace(" ", "").isalpha()


a = A("ashish Jain")
print(a.is_valid_name("ankit"))


====================== Class mathods ==================

class Person:
    species = "Homo sapiens"
    
    def __init__(self, name):
        self.name = name
    
    @staticmethod
    def is_valid_name(name: str) -> bool:
        """Utility check: doesn't depend on class or instance."""
        return bool(name) and name.isalpha()
    
    @classmethod
    def create_anonymous(cls):
        """Factory method: uses class reference to construct object."""
        return cls("Anonymous")


# Usage
print(Person.is_valid_name("Ashish"))   # ✅ Static method → True
print(Person.is_valid_name("123"))      # ❌ Static method → False

anon = Person.create_anonymous()        # ✅ Class method → creates Person("Anonymous")
print(anon.name)                        # Output: Anonymous
print(anon.species)                     # Accesses class attribute



class User:
    def __init__(self, name, sum):
        self.sum = sum
        self.name = name
        
    def getName(self):
        return self.name
    
    @staticmethod
    def checkNumber(n:int):
        return n>=18

    @classmethod
    def addition(cls,name, a,b):
        sum = a+b
        return cls(name, sum)
        
    
a = User("ashish jain", 10)
print(a.getName())
print(User.checkNumber(19))
b = User.addition("ankit",8,9)
print(b.sum)

