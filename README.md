# Token-management-system-in-go

# Features
1. Token Initialization:
   - Creates 1000 tokens with unique identifiers (`Token 1`, `Token 2`, ..., `Token 1000`).
   - All tokens start with a usage count of `0`.

2. Token Selection:
   - Selects the token with the least usage count.
   - If multiple tokens have the same usage count, one is selected randomly.

3. Token Usage Update:
   - Each time a token is selected, its usage count increases by `1`.

4. 24-Hour Reset:
   - A background process resets all token usage counts to `0` every 24 hours without affecting ongoing operations.

5. Simulation:
   - Simulates a specified number of token usage operations.
   - Updates token usage counts during each operation.

6. Reporting:
   - Displays usage counts for all tokens after the simulation.
   - Identifies and displays the least-used tokens.

# How It Works (Simple Approach)

Tokens are created: A total of 1000 tokens are initialized.

Token selection happens: When a token is needed, the system finds the one used the least and selects it.

Usage count increases: The selected token's usage count goes up by 1.

Simulation runs: The system performs token selection multiple times based on user input.

Reset process works in the background: Every 24 hours, all tokens are reset to zero, without slowing down the program.

Results are displayed: At the end, you see how many times each token was used and which ones were used the least.
