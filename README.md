# gh-review-cli

`gh-review-cli` is a lightweight CLI tool written in Go that lists the GitHub Pull Requests assigned to you for review.

It uses the GitHub REST API and requires a Personal Access Token.

## 🚀 Features

<ul>
    <li>List all open PRs where you are requested as a reviewer</li>
    <li>Show PR title, author, and link</li>
    <li>Display the total number of PRs</li>
</ul>

## 📦 Installation

Clone the repository:
```
git clone https://github.com/yourusername/gh-review-cli.git
cd gh-review-cli
```

Build the binary:
```
go build -o gh-review
```

Move it to your PATH (Linux/macOS):
```
sudo mv gh-review /usr/local/bin/
```

Now you can run gh-review from anywhere.

## 🔑 Setup

You need a GitHub Personal Access Token with repo and pull_request scopes.

Create one at:
<a href="https://github.com/settings/tokens">GitHub → Settings → Developer settings → Personal Access Tokens
</a>

Export it as an environment variable:
```
export GITHUB_TOKEN=your_token_here
```

## 🖥 Usage
List PRs for review
```
gh-review list
```

## 📫 Contributing to `gh-review-cli`

To contribute to `gh-review-cli`, follow these steps:

1. Fork this repository.
2. Create a branch: git checkout -b <branch_name>.
3. Make your changes and commit them: `git commit -m 'commit_message'`.
4. Push to your branch.
5. Create a pull request.

Alternatively, see the GitHub documentation on <a href="https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request">creating a pull request</a>