# Use Mode

## build
``` shell
cd src && go build -o cern
```

## run
``` shell
## pass all paramaters
./cern skill \
    --name golang \
    --description "Be able to work with golang development language" \
    --scope development \
    --category language \
    --purpose general \
    --level 1 \
    --priority 1 \
    --tags golang,go

## pass some paramaters
./cern skill \
    --name golang \
    --level 1

## pass a yaml manifest file
cat << EOF > myskills.yaml
kind: skill
name: golang
description: Be able to work with golang development language
scope: "development"
category: "language"
purpose: "general"
level: 1
priority: 1
tags: ["golang", "go"]
EOF

./cern skill -f myskills.yaml
```