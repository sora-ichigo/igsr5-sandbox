github_user_prompt() {
    local green='\033[0;32m'
    local red='\033[0;31m'
    local reset='\033[0m'

    if [ ! -z "${GITHUB_USER}" ]; then
        echo -en "🌈 ${green}@${GITHUB_USER} ${reset}${red}:${reset}"
    else
        echo -en "${green}@\u ${reset}${red}➜${reset}"
    fi
}

current_directory_prompt() {
    local blue='\033[1;34m'
    local reset='\033[0m'

    echo -en "${blue}$(pwd)${reset}"
}

git_branch_prompt() {
    local cyan='\033[0;36m'
    local red='\033[1;31m'
    local reset='\033[0m'
    local branch=$(git symbolic-ref --short HEAD 2>/dev/null || git rev-parse --short HEAD 2>/dev/null)

    if [ ! -z "$branch" ]; then
        echo -en "🌳 ${red}${branch}${reset}"
    else
        echo -en "🌳 ${red}no branch${reset}"
    fi
}

kube_fork_prompt() {
    local cyan='\033[0;36m'
    local red='\033[1;31m'
    local reset='\033[0m'

    if [ ! -z "${KUBE_FORK_TARGET_ENV}" ]; then
        echo -en " | 🍴 ${red}${KUBE_FORK_TARGET_ENV}${reset}"
    else
        echo -en " | 🍴 ${red}no fork${reset}"
    fi
}

export PS1='$(github_user_prompt) $(current_directory_prompt)   $(git_branch_prompt)$(kube_fork_prompt) \n  %: '
