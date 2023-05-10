# cernner
A platform to grok yourself. A place where is possivel organize, to see and to show what you know and who are you.


## Purpose
Provide a plaftorm to handle what you know. Storage, visualize, make it growth, grok what youy know


## References
- obsidian + logseq + roam + remnote + dendron + heptabase
- todoist + remember-the-milk + simplenote + trello
- quora + dev.to + medium + reddit + wordpress
- excalidraw + draw.io + miro + mural + whimsical + mindmeister
- toggl + desktime + activitywatch + timewattior
- taiga-io + asana + basecamp
- notion + evernote + joplin
- figma + sketch + zeplin
- taskwarrior + calcurse
- habitica + lifeup
- exercism + codewars + hackerhank + freecodecamp


## Adjectives
- portable
- customizable
- extensible
- complex
- intuitive
- effective
- cooperative
- funny
- rhizomatic


## Use cases
- Career
- Sport
- Personal Life
- University Graduation


## Global Schema
grok > skill > journey > project,training,study,award,tracker,note


## Modules
- groks
    - description
        - A bunch of pieces, of different types and sources, used to compile a solid knowledge about something. It's a place where you go when you are looking for some accurate information, empirical data, previously compiled, reliable, selected knowledge.
    - domain
    - schema
    - features
        - whiteboard
            - drag & drop media
            - knowledge spider graph + back-links + outgoing-links
        - rhizomatic graph
    - design
- skills
    - description
    - domain
        ```markdown
        Tempo de exposição + (Quantidade de projetos concluídos * Complexidade dos projetos concluídos)) /  Tempo de conclusão dos projetos) + Conhecimento adquirido + Conhecimentos relacionados
        
        journey:kubernetes | label:k8s ⇒ get_mastery_level { journey[time] + ( count(projects[status]:completed) + projects[difficulty] + project[time]) + learnings-wikis[size] + get_mastery_level(all[labels]:k8s) }
        
        skill_depth = qtde * complexity
        skill_level = time / depth
        ```
    - schema
        ```yaml
        skill:
        	subject:
        		topic:
        			properties: [level, depth] # water glass approach: more water, smaller depth
        ```
    - features
    - design
- journeys
    - description
    - domain
    - schema
        ```yaml
        journey:
        	path:
        		quest:
        		side-quest:
        		glory-quest: 
        		adventue:
        		challange: 
        ```
    - features
    - design
- projects
    - description
    - domain
    - schema
        ```yaml
        project:
        	initiative:
        		epic:
        			story:
        				kanban:
        					step: 
        						 columns: [backlog, validating, to-do, doing, testing, done, owner, collaborators, tags]
        					task: 
        						properties: [name, description, created-at, updated-at, deadline, priority, urgency, DoD, DoR]
        ```
    - features
        - timebox
            > Possibilidade de dividir período de tempo para focar em tasks específicas
            > 
        - task inception
            - Existem tasks com tarefas aninhadas onde, para concluir uma task, é necessário concluir sub-taskscom suas respectivas complexidades.
            - Para tal, será possível criar kanbans dentro de tasks: kanban > task > kanban > sub-task
        - pomodoro
            - merge restarts
            - action time, paused time, session-count
    - design
- trainings
    - description
    - domain
    - schema
        ```yaml
        ```
    - features
    - design
- studies
    - description
    - domain
    - schema
        ```yaml
        study:
        	multiverse: ## onion + helicoid approach
        		layer:
        			archive:
        				wiki:
        					articles:
        					videos: 
        					infographics:
        				  sketches:
        				  notes:
        				  books:
        				  podcasts:
        				  bookmarks: 
        ```
    - features
        - spaced repetition
        - active study
        - flashcards
        - memory palace
        - mnemonics
    - design
- awards
    - description
    - domain
    - scheme
        ```yaml
        scope:
            self:
                - medal
                - trophy
            work:
                - promotion
                - bonus
                - reward
            social:
                - certificate
                - certifications
                - degree
        ```
    - features
    - design
- trackers
    - description
    - domain
    - scheme
    - features
        - books
            - books
            - hqs
            - mangás
        - movies
            - movies
            - animated movies
        - animations
            - animations
            - animated series
            - animes
        - series
            - series
            - brazilian series
            - documentaries
        - games
            - pc-games
            - console-games
    - design
- notes
    - description
    - domain
    - scheme
    - features
    - design

## Technologies
- platforms
    - backend
        - REST API: golang - fiber/chi/gingonic/gorillamux
            - doc: swagger
            - logs: plain/json
            - data: json/protobuf/faltbuf
            - traces: opentelemetry
    - frontend
        -  PWA: nodejs - react/vuejs
            - drag-and-drop
            - whiteboard
    - cmd
        -  cli: golang - cobra/viper
            - config: yaml


## Architecture
- domain
    - entities
    - value-objects
- application
    - services
    - use-cases
- adapters
    - controllers
    - repositories
- infra
    - cmd
        - cli: viper + cobra
        - grpc: protobuf/flatbuf
    - web
        - api: fiber/chi/gin/mux + swagger
        - bff: graphql
    - core
        - database: mongodb
        - cache: redis
        - streaming: rabbitmq

## brainstorm
- features
    - todo-list
        - views: day, week, month, quarter, year
        - contexts: skill, journey, project, study, training
    - time-tracker
- weekend-projects
    - todo-list
        - poc
            - models
                - task
            - entities
                - task
                    - methods
                        - create
                        - describe
                        - list
                        - update
                        - remove
        - mvc
            - models
                - task
                    ```python
                    class Task(BaseModel):
                    	id: UUID
                    	title: str
                    	description: str
                    	owner: str
                    	status: str
                    	tags: list[str]
                    	creation_time: datetime
                    	update_time: datetime
                    	start_time: datetime
                    	end_time: datetime
                    	due_date: datetime	
                    	duration_time: timedelta
                    ```
                - to_do
                    ```python
                     class ToDo(BaseModel):
                    	id: UUID
                    	title: str
                    	description: str
                    	owner: str
                    	status: str
                    	tags: list[str]
                    	creation_time: datetime
                    	update_time: datetime
                    	start_time: datetime
                    	end_time: datetime
                    	due_date: datetime	
                    	duration_time: timedelta
                    
                    	tasks: list[Task] = None
                    
                    	skill: UUID = None
                    	journey: UUID = None
                    	project: UUID = None
                    	learning: UUID = None
                    	training: UUID = None
                    ```
            - entities
                - task
                    - methods
                        - create
                        - describe
                        - list
                        - update
                        - remove
                - to_do
                    - methods
                        - create
                        - describe
                        - list
                        - update
                        - remove
                        - add_task
                        - del_task
            - interfaces
                - database
                - cli
        - project
    - stopwatch

- movies / series / animations
    - imdb
    - rottentomatoes
    - metacritic
- books
    - goodreads
    - skoob
- games
    - steam




---
``` yaml
skill:
    - skill:
        - entity:
            ``` yaml
            id: uuid
            kind: string
            name: string
            description: string
            scope: string
            category: string
            purpose: string
            level: int
            priority: int
            tags: string
            ```
        - example:
            ``` yaml
            id: 1234-1234-1234-1234-1234
            kind: skill
            name: golang
            description: Be able to work with golang development language
            scope: development
            category: language
            purpose: general
            level: 1
            priority: 1
            tags: golang,go
            ```
```