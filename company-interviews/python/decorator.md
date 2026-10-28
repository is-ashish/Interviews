**Decorators in Python**

---

**What is a Decorator?**

A decorator is a function that **wraps another function** to add extra behavior — without modifying the original function's code.

```
Original Function
      +
Extra Behavior (logging, timing, auth etc.)
      =
Decorated Function
```

Think of it like a **gift wrap** — the gift (function) stays the same inside, but the wrap adds something extra outside.

---

**Before Decorators — The Problem**

```python
def greet():
    print("Hello!")

# Want to add logging to every function?
# Without decorators you'd do this manually:

print("Function starting...")
greet()
print("Function finished.")

# Imagine doing this for 50 functions ❌ — messy!
```

---

**How Decorators Work — Step by Step**

**Step 1 — Functions are First Class Objects**

In Python, functions can be:
- Passed as arguments
- Returned from functions
- Assigned to variables

```python
def greet():
    print("Hello!")

say_hi = greet       # assign to variable
say_hi()             # → Hello!

def run(func):
    func()           # pass as argument

run(greet)           # → Hello!
```

**Step 2 — A Function Inside a Function**

```python
def outer():
    def inner():     # inner function
        print("I am inner")
    inner()          # call inner

outer()              # → I am inner
```

**Step 3 — Return a Function**

```python
def outer():
    def inner():
        print("I am inner")
    return inner     # return function (no parentheses!)

func = outer()
func()               # → I am inner
```

---

**Building a Simple Decorator**

```python
def my_decorator(func):         # takes a function
    def wrapper():
        print("Before function runs")
        func()                  # call original function
        print("After function runs")
    return wrapper              # return wrapped function


def greet():
    print("Hello!")


# Apply decorator manually
greet = my_decorator(greet)
greet()
```

**Output:**
```
Before function runs
Hello!
After function runs
```

---

**Using @ Syntax (Syntactic Sugar)**

The `@` symbol is just a cleaner way to write the same thing:

```python
def my_decorator(func):
    def wrapper():
        print("Before function runs")
        func()
        print("After function runs")
    return wrapper


@my_decorator           # same as greet = my_decorator(greet)
def greet():
    print("Hello!")


greet()
```

**Output:**
```
Before function runs
Hello!
After function runs
```

---

**Decorator with Arguments**

Use `*args` and `**kwargs` to handle any function signature:

```python
def my_decorator(func):
    def wrapper(*args, **kwargs):    # accept any arguments
        print(f"Calling {func.__name__}")
        result = func(*args, **kwargs)
        print(f"Finished {func.__name__}")
        return result
    return wrapper


@my_decorator
def add(a, b):
    print(f"Sum = {a + b}")
    return a + b


add(3, 5)
```

**Output:**
```
Calling add
Sum = 8
Finished add
```

---

**Real World Examples**

**1. Logger Decorator**
```python
def logger(func):
    def wrapper(*args, **kwargs):
        print(f"LOG: {func.__name__} called with {args}")
        result = func(*args, **kwargs)
        print(f"LOG: {func.__name__} returned {result}")
        return result
    return wrapper


@logger
def multiply(a, b):
    return a * b


multiply(3, 4)
```
```
LOG: multiply called with (3, 4)
LOG: multiply returned 12
```

---

**2. Timer Decorator**
```python
import time

def timer(func):
    def wrapper(*args, **kwargs):
        start  = time.time()
        result = func(*args, **kwargs)
        end    = time.time()
        print(f"{func.__name__} took {end - start:.4f} seconds")
        return result
    return wrapper


@timer
def slow_function():
    time.sleep(1)
    print("Done!")


slow_function()
```
```
Done!
slow_function took 1.0012 seconds
```

---

**3. Authentication Decorator**
```python
def requires_auth(func):
    def wrapper(user, *args, **kwargs):
        if not user.get("is_authenticated"):
            print("Access Denied!")
            return
        return func(user, *args, **kwargs)
    return wrapper


@requires_auth
def dashboard(user):
    print(f"Welcome {user['name']}!")


dashboard({"name": "Ashish", "is_authenticated": True})
# → Welcome Ashish!

dashboard({"name": "Bob", "is_authenticated": False})
# → Access Denied!
```

---

**Stacking Multiple Decorators**

```python
@logger
@timer
def add(a, b):
    return a + b

# Execution order (bottom up):
# 1. timer wraps add first
# 2. logger wraps timer(add) second
# Result: logger → timer → add
```

---

**Summary**

```
┌─────────────────────────────────────────┐
│            my_decorator(func)           │
│  ┌───────────────────────────────────┐  │
│  │           wrapper()               │  │
│  │   print("Before")                 │  │
│  │   ┌─────────────────────────┐     │  │
│  │   │    original func()      │     │  │
│  │   └─────────────────────────┘     │  │
│  │   print("After")                  │  │
│  └───────────────────────────────────┘  │
└─────────────────────────────────────────┘
```

| Concept | Purpose |
|---|---|
| `@decorator` | Syntactic sugar to apply decorator |
| `wrapper()` | Adds extra behavior around original |
| `*args, **kwargs` | Makes decorator work with any function |
| Stacking | Apply multiple decorators on one function |

---

**Key Insight for Interview**

> A decorator is just a function that takes a function, wraps it inside another function to add behavior, and returns the wrapper. The `@` syntax is just clean shorthand for `func = decorator(func)`. They follow the **Open/Closed Principle** — open for extension, closed for modification.

---

Want to dive into **Generators**, **Context Managers**, or **`functools.wraps`** next? 🚀