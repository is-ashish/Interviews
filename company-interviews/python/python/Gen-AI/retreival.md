I prepared direct questions for AI interviews around topics like -

Vector DB

Embeddings

Agentic AI

But interviewer had different plans.

They turned their vehicle to Scenario based questions.

And asked me these 2 questions on vector embeddings.

1️⃣ Scenario: Wrong Results Despite Correct Documents

You have built a RAG chatbot for a company's internal knowledge base.

A user asks:

"How do I apply for parental leave?"

The correct document exists in the vector database.

However, the chatbot retrieves documents about:

- Annual leave
- Sick leave
- Leave approval workflows

instead.

Question:

How would you investigate whether the issue is caused by:

- Poor embedding model selection
- Chunking strategy
- Embedding dimension
- Similarity search configuration
- Query formulation

What metrics would you use to validate your findings?


2️⃣ Scenario: Embedding Model Migration

Your company has generated embeddings for 50 million documents using Model A.

A new embedding model becomes available and demonstrates significantly better retrieval quality.

The business wants to migrate.

Question:

How would you design the migration strategy?

Would you:

- Re-embed everything at once?
- Run both embedding models in parallel?
- Gradually migrate data?

How would you avoid downtime, retrieval inconsistencies, and increased infrastructure costs during the transition?


🎯 Bonus Follow-Up

Suppose two sentences have completely different words:

- "The customer purchased a vehicle."
- "The client bought a car."

Why would a good embedding model place these vectors close together in vector space?

What does this tell you about semantic understanding?