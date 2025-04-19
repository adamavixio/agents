# Agent Protocol

A Go-based protocol for orchestrating multiple AI agents to work together on tasks.

## Overview

This protocol allows you to define a set of AI agents, their capabilities, communication patterns, and workflow steps in a declarative YAML configuration. The system then orchestrates these agents to complete complex tasks by working together.

## Installation

```bash
go get github.com/adamjohnston/agentprotocol
```

## Configuration

Create a YAML file that defines your agents and workflow:

```yaml
version: "1.0"

task:
  name: "Research and summarize Go concurrency patterns"
  description: "Analyze different Go concurrency patterns and produce a summarized report"

agents:
  - id: "researcher"
    name: "Research Agent"
    role: "Searches and collects information about Go concurrency patterns"
    can_communicate_with:
      - "writer"
    capabilities:
      - "web_search"
      - "read_files"
  
  - id: "writer"
    name: "Writer Agent"
    role: "Compiles and formats information into a cohesive report"
    can_communicate_with:
      - "researcher"
    capabilities:
      - "write_files"

workflow:
  - step: "initial_research"
    agent: "researcher"
    action: "collect_information"
    inputs:
      - "go concurrency patterns"
    outputs:
      - type: "message"
        to: "writer"
  
  - step: "compile_report"
    agent: "writer"
    action: "write_report"
    depends_on:
      - "initial_research"
    outputs:
      - type: "file"
        path: "report.md"
```

## Usage

```bash
export CLAUDE_API_KEY="your_api_key_here"
./agentprotocol -c config.yaml
```

## Features

- Define multiple agents with specific roles and capabilities
- Control which agents can communicate with each other
- Create multi-step workflows with dependencies
- Support for different output types (messages between agents, files)
- Isolated workspace for each run

## Extending

You can extend the agent capabilities by modifying the `agent.go` file to add new functionality such as web search, code execution, etc.