In this section, we present the concept of schedules to help identify those executions that are guaranteed to ensure the isolation property and thus database consistency.

konsep schedules membantu meng-identifikasi eksekusi yang dijamin memastikan properti isolation dan konsistensi DB.

Given example banking system:

which has
several accounts, and a set of transactions that access and update those accounts.
Let T 1 and T 2 be two transactions that transfer funds from one account to another.
Transaction T 1 transfers $50 from account A to account B. It is defined as: