
# Compact AGI with Versatile Language Models (MultiLM-AGI)

MultiLM-AGI is a project that aims to create a simple but efficient Artificial General Intelligence (AGI) agent using Golang and a variety of Large Language Models (LLMs) compatible with the LLMClient interface. The agent is engineered to handle a variety of tasks by refining its objectives, executing tasks, evaluating results, and prioritizing subsequent tasks. With its adaptable architecture, MultiLM-AGI can be applied across various domains, including gaming, problem-solving, and knowledge acquisition.

**Please note that this project is still under development, and its individual-level efficiency has not yet been thoroughly established. However, given the escalating interest in the development of general artificial intelligence, we believe it's crucial to share the potential worth of this project and collaboratively evolve it with the community. Hence, we request your continued support and understanding.**

## Key Features

- **Objective Refinement**: An innovative approach to refine a larger objective into smaller, more manageable objectives that allow the AGI agent to tackle complex issues.
- Task Creation: Generates tasks and milestones based on the refined objectives.
- Execution Agent: Executes tasks using the OpenAI GPT-4 model (or GPT-3.5 model).
- Evaluation Agent: Evaluates the task results and their effectiveness.
- Prioritization Agent: Prioritizes tasks based on their relevance and importance.
- Task Context Agent: Stores the task context for future reference.
- **In-Memory Vector Store**: Efficient in-memory vector store to save embedding vectors for similarity queries, enabling faster access and improved performance.

## Note: Currently Supported LLM Provider

Though MultiLM-AGI is structured to be versatile and work with various Large Language Models (LLMs) that gratify the LLMClient interface, as of now, only OpenAI's GPT-4 or GPT-3.5 are supported as the LLM providers. Future releases may include support for additional LLM providers. Stay connected for further developments in the MultiLM-AGI project.

## Install
1. Install Go and configure your Go workspace.
2. Clone the repo:

```bash
git clone https://github.com/ReEcho-AI/agi-golang-flexlm.git
```

3. Navigate into the repo:

```bash
cd agi-golang-flexlm
```

4. Download the necessary Go packages:

```bash
go mod download
```

5. Create a .env file with your OpenAI API key:

```makefile
OPENAI_API_KEY=your_openai_api_key_here
```

## Usage

Start the main program:

```bash
go run ./cmd/main.go
```

The AGI agent will begin learning how to play chess by executing, reviewing, and refining tasks.

## Sequence Diagram

![Sequence Diagram](./img/agi-golang-flexlm-sequence-diagram.svg)

## Flowchart

![Flowchart](./img/agi-golang-flexlm-flowchart.svg)

## Example Output

```
======= Objective ======
I want to learn chess strategies.

======= Refined Objective ======
...

======= Milestone Objective ======
Learn the layout of the chessboard and the pieces' movements.

======= Refined Milestone Objective ======
Understand the chessboard layout, the starting positions of each piece, and the governing rules of movements and captures.

======= Task Result ======
...

======= Task Evaluation ======
This is a successful result because it precisely classifies each chess piece, their starting positions on the chessboard, and the rules governing their movements and captures.
```

## Contributing

Pull requests are welcomed. For substantial changes, kindly open an issue first for discussions.

## License

MIT