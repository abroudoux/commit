function commit() {
    local current_dir=$(pwd)

    local git_root
    git_root=$(git rev-parse --show-toplevel 2>/dev/null)

    if [[ $? -ne 0 ]]; then
        echo "Error: Not a git repository"
        return 1
    fi

    cd "$git_root" || return 1

    local branch_name
    branch_name=$(git rev-parse --abbrev-ref HEAD)

    git rev-parse --abbrev-ref --symbolic-full-name @{u} &>/dev/null

    if [[ $? -ne 0 ]]; then
        read -p "Branch '$branch_name' has no upstream branch. Do you want to create it? (yes/no): " response

        if [[ "$response" == "yes" ]] || [[ "$response" == "Y" ]] || [[ "$response" == "y" ]]; then
            echo "Creating upstream branch and pushing..."
            git add .
            git commit .
            git push -u origin "$branch_name"
        else
            echo "Skipping upstream branch creation."
            git add .
            git commit .
        fi
    else
        git add .
        git commit .
        git push
    fi

    cd "$current_dir" || return 1
}
