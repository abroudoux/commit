function commit() {
    local current_dir=$(pwd)

    local git_root
    git_root=$(git rev-parse --show-toplevel 2>/dev/null)

    if [[ $? -ne 0 ]]; then
        echo "Error: Not a git repository"
        return 1
    fi

    cd "$git_root" || return 1

    git add .
    git commit .

    local branch_name
    branch_name=$(git rev-parse --abbrev-ref HEAD)

    git rev-parse --abbrev-ref --symbolic-full-name @{u} &>/dev/null

    if [[ $? -ne 0 ]]; then
        echo "Branch '$branch_name' has no upstream branch. Creating it..."
        git push -u origin "$branch_name"
    else
        git push
    fi

    cd "$current_dir" || return 1
}