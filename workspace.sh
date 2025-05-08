#!/bin/bash

session=$1

[ "$session" = "" ] && "The session name (first argument) is required" && exit 1

tmux new-session -d -s "$session"

tmux rename-window -t 1 'IDE'

tmux send-keys -t 'IDE' 'devbox run nvim' C-m

tmux new-window -t $session:2 -n 'BE Shell'

tmux send-keys -t 'BE Shell' 'devbox run npm run dev -w=backend' C-m
#tmux send-keys -t 'BE Shell' 'cd backend' C-m

tmux new-window -t $session:3 -n 'FE Shell'

tmux send-keys -t 'FE Shell' 'devbox run npm run dev -w=frontend' C-m
#tmux send-keys -t 'FE Shell' 'cd frontend' C-m
