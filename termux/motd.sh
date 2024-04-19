#!/data/data/com.termux/files/usr/bin/bash

if ! command -v gum >/dev/null 2>&1; then
  echo "You need gum installed on the system"
  exit 1
fi

main() {
  local -r quote=$(shuf -n 1 ~/.termux/quotes.txt)
  local -r size=$(stty size | cut -d ' ' -f 2)
  local width=$((size - 2))

  gum style \
    --foreground 212 --border-foreground 212 --border normal \
    --align center --width $width --padding "1 1" \
    "$quote"
}

main "$@"
