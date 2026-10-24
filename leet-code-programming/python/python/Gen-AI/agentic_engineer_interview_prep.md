# Agentic Engineer Interview Prep — Multi-Agent Systems & LangGraph

Based on your Agentic Ticket Resolution project (LangGraph + LangChain + FastAPI + Chroma + Groq).

---

## Section 1: Multi-Agent Fundamentals

**Q1. What is a multi-agent system, and why use one instead of a single agent with many tools?**
A multi-agent system splits a task across several specialized agents, each with a narrow scope and small toolset, coordinated by an orchestrator. Reasons to split:
- **Smaller tool count per agent** → less hallucinated tool calls, better tool-selection accuracy
- **Separation of concerns** → billing logic doesn't leak into shipping prompts
- **Independent iteration** → you can improve/re-prompt the billing agent without touching others
- **Context window efficiency** → each agent only sees relevant instructions/tools, not all 30 tools

**Q2. What are the common multi-agent design patterns?**
- **Supervisor / router pattern** — a central node classifies and routes to specialist subgraphs (what you used).
- **Sequential pipeline** — agents run one after another (e.g., researcher → writer → editor).
- **Hierarchical** — supervisors of supervisors, used for very large tool ecosystems.
- **Peer-to-peer / debate** — agents critique each other's output (used in evaluation or research agents).
- **Blackboard pattern** — agents read/write to shared state without direct handoff (closest to LangGraph's shared state model).

**Q3. Why did you choose the supervisor/router pattern over a single agent with all tools?**
A single agent with billing + shipping + account tools (10+ tools) increases the chance of the LLM picking the wrong tool or hallucinating parameters, especially with smaller/faster models like Groq's Llama. Routing first with a cheap classifier keeps each specialist's tool surface small (2-4 tools), which measurably improves tool-call accuracy.

**Q4. What's the difference between orchestration and choreography in multi-agent systems?**
- **Orchestration**: a central controller (router/supervisor) explicitly decides the next agent — deterministic, easy to debug, easy to add approval gates. (Your project uses this.)
- **Choreography**: agents communicate peer-to-peer / event-driven with no central controller — more flexible but much harder to debug, trace, and bound (risk of infinite agent-to-agent loops).

**Q5. How do you prevent agents from stepping on each other's context or state?**
Use a single shared state object (TypedDict in LangGraph) with well-defined keys per concern (e.g., `issue_type`, `sensitive_action`, `retry_count`). Each node reads only what it needs and writes only its own fields back via partial state updates — LangGraph merges these automatically, so agents don't need to know about each other's internals.

---

## Section 2: LangGraph-Specific

**Q6. What is LangGraph and how does it differ from a plain LangChain AgentExecutor?**
LangGraph models an agent workflow as an explicit state machine (nodes + edges) with a shared, typed state object, instead of an implicit ReAct loop. This gives you: explicit conditional branching, cycles/retries, checkpointing (pause/resume), and human-in-the-loop interrupts — none of which are first-class in a basic AgentExecutor loop.

**Q7. Explain `StateGraph`, nodes, and edges.**
- `StateGraph(TicketState)` defines the shape of shared state (a TypedDict).
- **Nodes** are functions `(state) -> partial_state_dict` — each node's return value is merged into state.
- **Edges** connect nodes. `add_edge` is unconditional; `add_conditional_edges` routes based on a function's return value, matched against a dict of possible next nodes.

**Q8. What's the difference between `add_edge` and `add_conditional_edges`?**
`add_edge("A", "B")` always goes A → B. `add_conditional_edges("A", router_fn, {"x": "B", "y": "C"})` calls `router_fn(state)` after node A runs, and routes to whichever node the returned key maps to. This is how your intent router (billing/shipping/account/escalate) and the retry loop are implemented.

**Q9. How does state merging work across nodes? What happens with list-type fields like `history` or `messages`?**
By default, each key in the state dict is *overwritten* by the latest node's return value, unless you use a reducer (`Annotated[list, add_messages]` or similar) — with `add_messages`, LangGraph appends instead of overwriting, which is why the `messages` list keeps growing correctly across nodes. For `history` in your project, it's handled manually — each node does `state["history"] + [new_log]` before returning, since it's a plain list field without a reducer.

**Q10. What is a checkpointer and why is it critical for production agents?**
A checkpointer persists the full state at every node transition, keyed by `thread_id`. It enables:
- **Resume after interrupt** (human approval flows)
- **Crash recovery** (resume from last checkpoint, not from scratch)
- **Audit/observability** (full history of state transitions)
- **"Track my ticket" features for free** — just query state by `thread_id`, no separate tracking DB needed
You used `SqliteSaver` for prototyping; production would move to `PostgresSaver` for concurrency and durability.

**Q11. How does `interrupt_before` work, and how do you resume?**
`graph.compile(checkpointer=..., interrupt_before=["human_approval"])` pauses execution right before that node runs, persisting state. To resume: `graph.update_state(config, {"approved": True}, as_node="human_approval")` then `graph.stream(None, config)` — passing `None` as input resumes from the last checkpoint instead of restarting.

**Q12. Why interrupt *before* the node rather than inside it?**
`interrupt_before` halts the graph entirely — no code in `human_approval_node` runs until resumed, so there's no risk of partial side effects. If you tried to pause *inside* a node using custom logic, you'd risk partial execution (e.g., a tool call fired before the pause registered).

**Q13. How would you scale this from SQLite to Postgres?**
Swap `SqliteSaver(conn)` for `PostgresSaver` (`langgraph-checkpoint-postgres`), using a connection pool instead of a single sqlite3 connection. This is necessary because SQLite doesn't handle concurrent writers well — a production system with many simultaneous tickets needs Postgres's proper locking/concurrency.

**Q14. What's the difference between `graph.invoke()` and `graph.stream()`?**
`invoke()` runs the graph to completion (or until an interrupt) and returns only the final state. `stream()` yields intermediate state after each node — useful for showing live progress in a UI, or for consuming a generator to force full execution as done in your `submit_ticket` endpoint (`for _ in graph.stream(...): pass`).

---

## Section 3: Agent Design & Tool Use

**Q15. What makes a good tool definition for an LLM agent?**
- Clear, unambiguous docstring describing *when* to use it, not just what it does
- Strongly typed arguments (Pydantic schema) so the LLM has a clear contract
- Narrow scope — one tool, one responsibility (e.g., `get_invoices` vs. a mega `billing_action` tool)
- Return structured, parseable output (JSON) rather than free text, so downstream nodes (like your sensitivity checker) can inspect `tool_call.args` reliably

**Q16. How do you decide which actions require human approval vs. autonomous execution?**
Risk-based thresholds: irreversible or costly actions (refund > $100, account deletion, production reboot) require approval; read-only or low-risk actions (lookup, status check, small refund) execute autonomously. This should be a configurable policy, not hardcoded — in your project it's a simple `if name == "process_refund" and amount > 100`, but production systems typically externalize this as a rules engine or policy table.

**Q17. How do you prevent an agent from looping forever?**
Bounded retry counts (`retry_count < 3` in your `route_verification_result`) with an escalation fallback once exhausted. Also worth adding: a hard step-count limit on the whole graph execution (LangGraph supports `recursion_limit`) as a safety net independent of your own retry logic.

**Q18. Why have a separate "verify" node instead of trusting the tool's success response?**
A tool returning HTTP 200 doesn't mean the *customer's problem* is solved — e.g., a refund API might succeed but for the wrong amount, or a tracking lookup might return "no data" without erroring. The verify node does a semantic check (rule-based or LLM-judged) against the original customer intent, not just the API status code.

**Q19. What's the risk of using an LLM to "verify" its own agent's work?**
Self-verification bias — the same model family may rubber-stamp its own output. Mitigations: use a different/cheaper model for verification than for the acting agent, use rule-based checks where possible (as your Mock fallback does — scanning for "error" strings), and log verification reasoning for human audit.

**Q20. How would you evaluate whether your multi-agent system is actually working well in production?**
- **Tool-call accuracy** — % of tool calls with valid, sensible arguments
- **Resolution rate** — % tickets closed without escalation
- **Escalation rate & reasons** — categorize why tickets fail (bad classification vs. missing tool vs. actual limitation)
- **Retry rate** — high retry rate signals verify/agent misalignment
- **Latency per node** — via LangSmith tracing, to find bottleneck nodes
- **Human override rate** on approvals — if humans reject often, your sensitivity threshold is miscalibrated

---

## Section 4: RAG vs Agentic — Comparison Questions

**Q21. When should a system stay RAG-only vs. go agentic?**
RAG-only suffices for pure information retrieval (policy lookup, FAQ). Go agentic when the request requires: taking an *action* (not just answering), account-specific data lookups (order status, balance), multi-step reasoning across systems, or explicit escalation/human handoff.

**Q22. How do RAG and agents work together in your system?**
Chroma DB (RAG) is used as a knowledge layer specialist agents can query for policy/procedure context (e.g., "what's our refund policy for damaged goods") while the *agentic* layer handles the actual state-changing actions (calling `process_refund`). RAG answers "what should happen"; the agent makes it happen.

**Q23. What's the cascade/fallback architecture from RAG to agent, and what triggers the handoff?**
Start with RAG (cheap, fast) for the initial answer. Escalate to the agentic resolver when: retrieval confidence is low, the query contains an action verb (reset, refund, cancel), the customer explicitly asks to escalate, or the query needs account-specific data RAG can't access. This is a cost-control pattern — you don't want every simple FAQ hitting a full multi-agent graph.

---

## Section 5: Error Handling & Reliability

**Q24. What happens in your system if a tool call throws an exception (e.g., API timeout)?**
The `ToolNode` (or your wrapper) catches it and returns a `ToolMessage` containing the error, which flows into `verify_result_node`. The verifier detects the error string and marks `resolved=False`, incrementing `retry_count`. This keeps failures inside the normal state flow rather than crashing the graph — the agent gets a chance to self-correct or eventually escalate.

**Q25. What's your fallback strategy if the LLM provider (Groq) is down or rate-limited?**
Your project has a `MockChatGroq` fallback used explicitly for testing, but in production you'd want: retry with exponential backoff, a secondary model provider fallback, and circuit-breaker logic so a full provider outage escalates tickets to humans gracefully instead of hanging.

**Q26. How do you debug a multi-agent system when something goes wrong three nodes deep?**
LangSmith tracing — every node execution, prompt, tool call, and state diff is logged per `thread_id`/run. Combined with your own `history` audit log in state, you get both a technical trace (LangSmith) and a human-readable timeline (your custom log) for the same execution.

**Q27. Why keep an explicit `history` audit log in state in addition to LangSmith tracing?**
LangSmith is for developer debugging; the `history` list in state is domain-facing — it's what gets shown to the customer/support agent in the UI timeline, and it persists with the ticket via the checkpointer independent of your LangSmith retention window. It's also queryable without hitting an external tracing service.

---

## Section 6: Architecture / Systems Design Questions

**Q28. How would you handle 10,000 concurrent tickets with this architecture?**
- Move checkpointer to Postgres with connection pooling
- Run the FastAPI service behind multiple workers/replicas (stateless app layer — state lives in Postgres, not memory)
- Queue-based ingestion (SQS/Celery) instead of synchronous `graph.stream()` in the request path, so ticket submission returns immediately and processing happens async
- Rate-limit/batch LLM calls to stay within provider limits

**Q29. How would you version-control changes to agent prompts/behavior safely?**
Treat prompts as versioned artifacts (not inline strings) — store in a prompt registry (or LangSmith's prompt hub), A/B test new agent prompts against a held-out ticket set before rollout, and keep the verify node's pass/fail rate as a regression signal.

**Q30. LangGraph vs CrewAI vs AutoGen — how do you choose?**
- **LangGraph**: best when you need explicit control flow, cycles, human-in-the-loop interrupts, and durable state — i.e., production workflows with approval gates (your use case).
- **CrewAI**: higher-level, role-based abstraction (crew of agents with roles/goals) — faster to prototype, less fine-grained control.
- **AutoGen**: strong for conversational multi-agent patterns (agents debating/collaborating) — less suited for deterministic business workflows with hard approval gates.
For a ticket resolution system needing auditability and interrupts, LangGraph's explicit graph model is the right fit — this is a good talking point to justify your tech choice.

---

## Section 7: Rapid-Fire / Likely Curveballs

**Q31. What's the difference between an "agent" and a simple LLM call with tools bound?**
An agent has a loop: it can observe a tool result and decide the *next* action, potentially calling more tools, until it reaches a stopping condition. A single LLM call with tools bound is one-shot — if it makes a tool call, something else (your code) has to explicitly decide what happens next. LangGraph nodes generally still need this decision logic (routing edges) — the "agency" is expressed by the graph structure, not automatically inside the LLM call.

**Q32. What are the risks of giving an agent too much autonomy?**
Unintended irreversible actions (hence your human-approval gate), cost blowup from excessive tool/LLM calls in retry loops (hence bounded retries), and cascading errors where one bad tool call misleads verification and downstream agents.

**Q33. How do you keep prompt/token costs under control in a multi-agent system?**
Use a small/cheap model for classification and verification (routing doesn't need your best model), keep per-agent context minimal (don't pass full conversation history to every node if not needed), and cap retries so failures don't multiply LLM calls indefinitely.

**Q34. If asked to add a new specialist agent (e.g., "technical support"), what would you change?**
Add a new node function (`tech_agent_node`) with its own tool subset, add it to `route_by_intent`'s classification options, register it as a target in the `check_sensitive_actions` and retry-routing conditional edges, and add its edge to `check_sensitive_actions`. The graph structure and checkpointing require no changes — this is the main selling point of the supervisor pattern's extensibility.

**Q35. How would you unit test a LangGraph node in isolation?**
Since nodes are plain functions `(state) -> dict`, you can test them directly without running the full graph — pass a hand-constructed `TicketState` dict and assert on the returned partial state. For conditional edge functions, same approach: call `route_by_intent({"issue_type": "billing"})` and assert it returns `"billing_agent"`.

---

## Quick Reference: Your Project's Numbers/Facts to Have Ready

- Retry cap: 2 attempts before escalation (`if retries < 2`)
- Sensitivity threshold: refund > $100 triggers human approval
- Checkpointer: SQLite (prototype) → Postgres (production path)
- Tracing: LangSmith
- LLM: Groq (Llama 3.x family) with mock fallback for offline testing
- Knowledge layer: Chroma DB for policy/RAG context
- API: FastAPI exposing submit / resume / status / list endpoints, all keyed by `thread_id`
