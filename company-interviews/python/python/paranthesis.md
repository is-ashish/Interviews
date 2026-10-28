Great choice — parenthesis-based problems are a staple in coding interviews because they test stack usage, string manipulation, and validation logic. Here are the most common ones you’ll encounter:
---
🔟 Must-Practice Parenthesis Problems
Valid Parentheses (LeetCode 20)
Check if a string of ()[]{} is balanced.
Core idea: use a stack to match opening and closing brackets.
Longest Valid Parentheses (LeetCode 32)
Find the length of the longest valid substring of parentheses.
Tests dynamic programming or stack-based reasoning.
Generate Parentheses (LeetCode 22)
Given n, generate all valid combinations of n pairs of parentheses.
Classic recursion + backtracking problem.
Remove Invalid Parentheses (LeetCode 301)
Remove the minimum number of parentheses to make the string valid.
Requires BFS or DFS with pruning.
Minimum Add to Make Parentheses Valid (LeetCode 921)
Count how many parentheses you need to add to balance the string.
Simple counter-based solution.
Check Redundant Braces (InterviewBit)
Detect if there are unnecessary parentheses in an expression.
Stack-based parsing problem.
Score of Parentheses (LeetCode 856)
Compute score based on rules: () = 1, (A) = 2 * score(A), AB = score(A) + score(B).
Tests recursion or stack evaluation.
Valid Parenthesis String (LeetCode 678)
String contains (, ), and * (wildcard). Check if valid.
Greedy + range tracking solution.
Add Parentheses (LeetCode 241)
Add parentheses in different ways to compute all possible results of an expression.
Tests recursion + divide-and-conquer.
Different Ways to Add Parentheses (LeetCode 241 variant)
Similar to above, but focuses on arithmetic expression evaluation.
---
🎯 Why These Matter
They cover stack, recursion, dynamic programming, and greedy algorithms.
Interviewers love them because they reveal how you handle nested structures and edge cases.
Many are building blocks for parsing problems in compilers or interpreters.
---
👉 Since you’re actively practicing interview prep, I’d suggest starting with Valid Parentheses (easy, stack-based) and then moving to Longest Valid Parentheses (medium-hard, DP/stack).
Would you like me to rank these problems by difficulty progression (easy → medium → hard) so you can plan your practice path efficiently?




Design a Python LLM Client class that uses the Requests library to interact with external Large Language Model (LLM) endpoints.
LLM_URL = 
Client_URL =
Client_secret =
 
1) Create the get_access_token method to get the access token
2) create the get_llm_response method to call the LLM using the access token
3) use request module from python
# 4) handle RPM = 20 
 
 


 import requests

class LLMClient:

    def __init__(self, Client_secret, Client_URL):
        self.Client_URL = Client_URL
        self.Client_secret = Client_secret
        
    def get_access_token():
        return self.access_token
        
    def get_llm_response(self,llm_url):
        # call LLM URL Client
        return request.client(llm_url, self.Client_URL, self.Client_secret)
    
    
LLM_URL = "llm_URL" # from env
access_token  = #fromm env
client = LLMClient(access_token_secret, access_token_URL)
token = client.get_access_token()
result = client.get_llm_response(LLM_URL, token.URL, token.secret)
print(result)

    
        
