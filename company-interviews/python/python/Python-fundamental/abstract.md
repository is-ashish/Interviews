<!-- An abstract method is a method that is declared but not implemented in a base class.

but their subclasses must implement, ensuring consistent behavior across different implementations.


• Useful for defining a common interface across different subclasses.


⚡ Why use abstract methods?
• They enforce consistency: every subclass must implement the required behavior.
• They define a blueprint for subclasses.
• Common in frameworks and large systems where you want to guarantee certain methods exist. -->



from abc import ABC, abstractmethod

class A:
    def __init__(self, name):
        self.name = name
        
    def get_Name(self):
        return self.name
    
    @abstractmethod
    def details(self):
        pass

class B(A):
    def details(self):
        return f"definition of class B -: {self.name}"

class C(A):
    def details(self):
        return f"definition of class C -: {self.name}"
    

b= B("ankit")
c= C("manish")

print(b.details())
print(c.details())
