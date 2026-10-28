# **Vector Search & RAG Interview Questions**

## **🟢 BASIC LEVEL Questions**

### **Conceptual Understanding**

1. **What is a vector embedding? How is it different from plain text?**
   - Expected Answer: Vectors are numerical representations capturing semantic meaning; text is just characters

2. **What is RAG? Why do we need it?**
   - Expected Answer: Retrieval-Augmented Generation - retrieves relevant documents then generates answers with that context

3. **What is vector similarity? Name 2 similarity metrics.**
   - Expected Answer: Cosine similarity, Euclidean distance, Manhattan distance, Dot product

4. **What is a vector database? Give examples.**

5. **What's the difference between traditional search and semantic search?**
   - Expected Answer: Traditional = keyword matching; Semantic = meaning-based matching

6. **Explain the RAG pipeline in 3-4 steps.**
   ```
   Query → Embed → Search → Retrieve Docs → Generate Answer
   ```

---

## **🟡 INTERMEDIATE LEVEL Questions**

### **Algorithm & Performance**

7. **Why is Cosine Similarity preferred for text embeddings over Euclidean distance?**
   - Expected Answer: Cosine measures angle (direction), not magnitude; two long documents with same meaning get high scores

8. **What is Brute Force search? When would you use it?**
   - Expected Answer: Compares query with all documents; O(n) complexity; use for small datasets (<100K docs)

9. **What is ANN? Why is it important for large-scale RAG?**
   - Expected Answer: Approximate Nearest Neighbor; trades slight accuracy for massive speed improvement; essential for >1M documents

10. **Explain HNSW algorithm in simple terms. What's its advantage?**
    - Expected Answer: Multi-layer hierarchical graph; searches efficiently through layers; used by Pinecone, Weaviate

11. **What is IVF? How does clustering help in search?**
    - Expected Answer: Inverted File Index; divides vectors into clusters; narrows search space; faster than brute force

12. **Difference between HNSW and IVF algorithms.**
    ```
    HNSW: Hierarchical graphs (better accuracy ~95%)
    IVF: Cluster-based (faster, ~90% accuracy)
    ```

13. **What is the time complexity of different search methods?**
    - Brute Force: O(n)
    - HNSW: O(log n)
    - IVF: O(√n)

14. **How does Product Quantization work? What problem does it solve?**
    - Expected Answer: Compresses vectors; reduces memory usage by 10-100x; trades speed for accuracy

---

## **🔴 ADVANCED LEVEL Questions**

### **Design & Architecture**

15. **Design a RAG system for 1 billion documents. What challenges would you face?**
    - Memory constraints
    - Latency requirements
    - Embedding model choice
    - Vector DB selection
    - Caching strategies
    

### ========================== Done ======================================

16. **How do you handle out-of-date information in your vector database?**
    - Expected Answer: Regular re-embeddings, incremental updates, TTL (time-to-live), soft deletes

### ========================== Done ======================================
17. **What's the difference between vector store and vector index?**
    - Store = full document storage
    - Index = optimized search structure
### ========================== Done ======================================

18. **Explain the "curse of dimensionality" in vector search.**
    - Higher dimensions → distances become more uniform
    - All vectors look equidistant
    - Solution: Dimensionality reduction techniques

### ========================== Done ======================================

19. **How do you evaluate embedding model quality for your domain?**
    - Expected Answer: Benchmark on domain-specific datasets, measure retrieval accuracy, compare with baselines

### ========================== Done ======================================
20. **What's the tradeoff between embedding model size and accuracy?**
    - Larger models (384+ dims): Better accuracy, slower, more memory
    - Smaller models (768 dims): Faster, less accurate, less memory

---

## **💻 CODING/IMPLEMENTATION Questions**

### **Practical Problems**

21. **Write pseudocode for a basic vector search function using Cosine Similarity.**
    ```
    function cosineSearch(query, documents, top_k):
        queryVector = embed(query)
        scores = []
        for each doc in documents:
            docVector = embed(doc)
            similarity = cosine(queryVector, docVector)
            scores.append({doc, similarity})
        return top_k highest scores
    ```

22. **How would you implement approximate search for 10M documents in Python?**
    - Expected: Use Faiss, Qdrant, or Pinecone SDK

23. **Write code to embed documents and store in vector DB.**
    ```python
    from pinecone import Pinecone
    from sentence_transformers import SentenceTransformer
    
    model = SentenceTransformer('all-MiniLM-L6-v2')
    pc = Pinecone()
    
    documents = ["Python guide", "Java tutorial", ...]
    embeddings = model.encode(documents)
    
    pc.index("my-index").upsert(
        vectors=list(zip(range(len(embeddings)), embeddings))
    )
    ```

24. **How would you query the vector store?**
    ```python
    query = "How to learn Python?"
    query_vector = model.encode(query)
    results = pc.index("my-index").query(
        vector=query_vector, 
        top_k=5
    )
    ```

25. **Implement a simple RAG pipeline.**
    ```python
    # 1. Retrieve
    query = "What is AI?"
    results = vectorDB.search(query, top_k=3)
    context = "\n".join([r.text for r in results])
    
    # 2. Generate
    prompt = f"Based on: {context}\n\nAnswer: {query}"
    answer = llm.generate(prompt)
    ```

---

## **🤔 SCENARIO-BASED Questions**

### **Real-World Problems**

26. **Your RAG system returns irrelevant documents. What could be wrong?**
    - Poor embedding model
    - Mismatch between query and documents
    - Wrong similarity metric
    - Incorrect preprocessing
    - Low top_k value

27. **Search latency is 2 seconds. How would you optimize?**
    - Use HNSW instead of IVF
    - Add caching layer
    - Use smaller embedding model
    - Reduce vector dimensions
    - Implement query batching

28. **How would you handle multi-language queries in your RAG system?**
    - Use multilingual embedding model
    - Translate to common language
    - Separate indices per language
    - Cross-lingual embeddings

29. **What happens when embedding model is updated?**
    - Need to re-embed all documents
    - Re-index vector database
    - Plan migration carefully
    - Gradual rollout strategy

30. **How do you prevent hallucinations in RAG?**
    - Retrieve high-quality documents
    - Set confidence threshold
    - Use retrieval-augmented generation
    - Fact-check with source documents
    - Temperature tuning in LLM

---

## **📊 Data & Quality Questions**

31. **How do you measure retrieval quality? What metrics would you use?**
    - Precision@K
    - Recall@K
    - MRR (Mean Reciprocal Rank)
    - NDCG (Normalized Discounted Cumulative Gain)
    - F1 Score

32. **What's the difference between index time and query time?**
    - Index time: Building the search structure (can be slow)
    - Query time: Searching (must be fast)

33. **How do you handle document updates in vector database?**
    - Upsert (update + insert)
    - Delete old, insert new
    - Versioning strategy
    - Soft deletes with metadata

34. **What preprocessing is important before embedding?**
    - Chunking documents
    - Removing noise (HTML tags, special chars)
    - Normalization
    - Tokenization

---

## **⚡ Quick Fire Round**

35. **What embedding dimension should I use? (256, 768, 1536?)**
    - Start with 384-768 (good balance)
    - Larger = better accuracy (more memory)
    - Smaller = faster (less accurate)

36. **HNSW vs IVF - which is better?**
    - HNSW: Better accuracy, higher memory
    - IVF: Faster, lower memory

37. **Cosine vs Euclidean - when to use which?**
    - Cosine: Text embeddings (default)
    - Euclidean: Dense numerical data

38. **How many top_k results should I retrieve?**
    - Small RAG: top_k=5
    - Large RAG: top_k=10-20
    - Balance accuracy vs token usage

39. **What's the typical latency for vector search?**
    - Brute force: 100ms-1s
    - HNSW: 1-50ms
    - With caching: <10ms

40. **How often should you re-embed documents?**
    - When content changes significantly
    - When embedding model updates
    - Incremental approach (not all at once)

---

## **🎯 Follow-Up Questions (Interviewer Will Ask)**

After your answer, expect:
- **"Why?"** - Explain your reasoning
- **"What if...?"** - Handle edge cases
- **"How would you measure it?"** - Metrics & monitoring
- **"What are tradeoffs?"** - Performance vs accuracy
- **"Any production issues?"** - Real-world experience

---

## **💡 Best Interview Tips**

| ✅ Do | ❌ Don't |
|------|---------|
| Explain with examples | Use jargon without context |
| Ask clarifying questions | Assume requirements |
| Discuss tradeoffs | Say "always use X" |
| Show system design thinking | Jump to implementation |
| Mention monitoring/metrics | Ignore performance |
| Discuss edge cases | Oversimplify problems |

---

**Would you like me to:**
- ✅ Create detailed answer explanations for any question?
- ✅ Build a code repository with examples?
- ✅ Create a practice problem set?
- ✅ Design a complete interview prep guide?