# Branch protection and auto-merge

`main` requires **one approving review** and **passing CI** before a pull request merges. After approval, the [Automerge workflow](workflows/automerge.yml) enables GitHub **auto-merge** so the PR merges when checks finish (no manual merge click).

Repository maintainers must apply the settings below once (GitHub UI or API). The Cloud Agent token cannot change these settings.

## 1. Allow auto-merge

1. Open **Settings → General → Pull Requests**.
2. Enable **Allow auto-merge**.

## 2. Ruleset for `main`

1. Open **Settings → Rules → Rulesets**.
2. **New ruleset → New branch ruleset**.
3. **Ruleset name:** `main`
4. **Enforcement:** Active
5. **Bypass list:** leave empty (or only break-glass admins).
6. **Target branches:** `main` (default branch).
7. Add rules:

| Rule | Setting |
|------|---------|
| Restrict deletions | on |
| Require a pull request before merging | on |
| Required approvals | **1** |
| Dismiss stale pull request approvals when new commits are pushed | on |
| Require status checks to pass | on |
| Require branches to be up to date before merging | on |
| Status checks (required) | `Go`, `Solidity`, `Automerge` |

8. Save the ruleset.

Status check names match CI job names in [.github/workflows/ci.yml](workflows/ci.yml) and the `Automerge` job in [.github/workflows/automerge.yml](workflows/automerge.yml). If GitHub shows different labels, pick the checks that correspond to those workflows.

### Optional: import ruleset JSON

A reference ruleset lives at [.github/rulesets/main.json](rulesets/main.json). GitHub does not load this file automatically; import it via **Rules → Rulesets → Import a ruleset** (or the REST API) if you prefer.

## 3. Optional: `AUTOMERGE_PAT` secret

If the Automerge workflow cannot enable auto-merge with `GITHUB_TOKEN`, create a fine-grained PAT with **Contents** and **Pull requests** write access on this repository, then add it as repository secret **`AUTOMERGE_PAT`**.

## Contributor flow

1. Open a pull request against `main`.
2. Wait for CI (`Go`, `Solidity`).
3. Get **one approval** from a reviewer with write access.
4. The Automerge workflow queues auto-merge; GitHub merges when all required checks pass.
5. Do not push new commits after approval without expecting a re-review (stale approvals are dismissed).

See also [CONTRIBUTING.md](../CONTRIBUTING.md).

## 中文（维护者）

1. **Settings → General → Pull Requests**：开启 **Allow auto-merge**。
2. **Settings → Rules → Rulesets**：为 `main` 新建规则集，要求 1 个 approve、CI 通过（`Go`、`Solidity`、`Automerge`），并开启「新提交后作废旧 approve」。
3. 若工作流无法自动合并，可配置仓库 Secret `AUTOMERGE_PAT`（PAT 需有 Contents / Pull requests 写权限）。

贡献者：提 PR → CI 通过 → 获得 1 个 approve → 工作流启用 auto-merge → 检查全绿后自动合并。
