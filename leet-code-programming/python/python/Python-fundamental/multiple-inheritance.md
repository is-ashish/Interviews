class A:
    def say(self):
        print("says- A")
        
class B(A):
    def say(self):
        super().say()
        print("says- B")

class C(A):
    def say(self):
        super().say()
        print("says- C")

class D(C,B):
    pass

d = D()

d.say()