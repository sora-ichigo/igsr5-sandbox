# ctrl+r search history
fh() {
  local selected_command
  selected_command=$(cat ~/.bash_history | fzf)
  echo $selected_command
}
bind -x '"\C-r": "READLINE_LINE=$(fh); READLINE_POINT=${#READLINE_LINE}"'
