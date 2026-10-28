[Ticket In] ──> [Classifier / Router Node]
                       │ (Conditional intent routing)
                       ▼
         ┌─────────────┼─────────────┐
         ▼             ▼             ▼
    [Billing]     [Shipping]     [Account]
      Agent         Agent          Agent
         │             │             │ (Decides tool calls)
         └─────────────┼─────────────┘
                       ▼
                 [Tool Node] ──> (Is sensitive? e.g. refund > $100)
                       │
             ┌─────────┴─────────┐
             ▼ (No)              ▼ (Yes)
       [Execute API]       [Human Approval] (Interrupt Halt)
             │                   │
             ▼                   ▼ (Approved?)
             │           ┌───────┴───────┐
             │           ▼ (Yes)         ▼ (No)
             │     [Execute API]   [Inject Error to Agent]
             │           │               │
             └───────────┼───────────────┘
                         ▼
                  [Verify Result] <──────┐ (Retry Loop, max 3)
                   /          \          │
               (Yes)          (No) ──────┘ (Increment retry count)
                 /              \
                ▼                ▼ (Exhausted retries)
             [Close]        [Escalate Node] (LLM summary -> Human Queue)