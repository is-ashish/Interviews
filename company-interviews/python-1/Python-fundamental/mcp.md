=========================== Protocol ========================


“MCP is an open protocol that standardizes how LLMs connect to external tools and data. 
It runs on JSON-RPC 2.0, 
supports stateful connections, 
and requires capability negotiation. 
Servers expose resources, prompts, and tools; clients expose sampling, roots, and elicitation. 

Security is enforced through user consent, 
zero-trust design, and audit logging. 
Transports include stdio, HTTP, and WebSocket, 
performance features like async processing and caching.”

========================= Authentication ================

“In MCP, authentication is handled with a zero-trust approach. 
Servers expose tools and resources only after explicit user consent. 

Secure transports like HTTPS/WebSocket with TLS are used, 

and tokens such as API keys or OAuth2/JWT maintain session authentication. 
Capability negotiation ensures both sides agree on supported methods. 
Finally, audit logging provides transparency and accountability.”