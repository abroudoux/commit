function commit() {
    local current_dir=$(pwd)

    if ! git rev-parse --is-inside-work-tree &>/dev/null; then
        echo "Error: This is not a Git repository."
        exit 1
    fi

    local local_branch_name=$(git rev-parse --abbrev-ref HEAD)
    local git_root=$(git rev-parse --show-toplevel 2>/dev/null)

    git add .

    if ! git commit; then
        echo "Commit aborted or failed."
        exit 1
    fi

    if git rev-parse --abbrev-ref --symbolic-full-name @{u} &>/dev/null; then
        local remote_branch=$(git rev-parse --abbrev-ref @{u})
        echo "The local branch '$local_branch_name' is tracking the remote branch '$remote_branch'."
        git push
    else
        echo "The local branch '$local_branch_name' does not have a remote branch configured."
        read "response?Do you want to create it? (yes/no): "

        if [[ "$response" =~ ^(yes|y|Y)$ ]]; then
            echo "Creating upstream branch and pushing..."
            git push -u origin "$local_branch_name"
        else
            echo "Skipping upstream branch creation."
        fi
    fi
}
