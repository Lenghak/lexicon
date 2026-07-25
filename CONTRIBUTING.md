# Contributing

## Local Setup

Install dependencies:

    bun install

Install git hooks:

    bunx lefthook install

Validate hooks:

    bunx lefthook validate

## Commit Messages

Follow Conventional Commits.

Examples:

    feat(iam): add user lockout
    fix(account): prevent duplicate registration
    docs(repo): update setup guide

## Branch Naming

    feature/*
    bugfix/*
    hotfix/*
    build/*

## Pull Requests

- Keep PRs focused
- One logical change per PR
- CI must pass
