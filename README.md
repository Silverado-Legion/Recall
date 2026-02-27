# Recall
Fast command-line tool for recalling notes and tasks right from the Terminal. Built with Go.<br>
`⚠️ ALPHA STAGE - Core Features are Not Complete.`

This README will be filled out more as project progresses. README should be completely filled out before the core features are complete.

## Quick Demo:
```
$ recall add Recall To-Do No.9: Implement a cool TUI. # - This will have an auto-generated title.
Note saved as: Recall To-do No.9: Implement

$ recall add Remember you have an appointment tomorrow! --title "Appointment Tomorrow" # - This has an explicit title.
Note saved as: Appointment Tomorrow
```
## Usage:
### Add Note:
```
$ recall add Note content here.
$ recall add Note content here again. --title "Custom Title"
$ recall add "Note content here but in quotes so you can put symbols like: #."
```

## To Do:
✅ - Add notes, with automatic titling.<br>
✅ - Write to database.<br>
🔁 - List all notes.<br>
➡️ - Recall notes from database.<br>
➡️ - Search for notes.<br>
➡️ - Mark notes complete.<br>
⏭️ - Make the outputs more "snappy".
⏭️ - Implement a smarter title autogeneration system.
⏭️ - Terminal User Interface (in addition to CLI controls).<br>
⏭️ - Tasks, which are notes with reminders (or time limits).<br>

🔁 = In Progress.<br>
➡️ = Next Up.<br>
⏭️ = Future.<br>

## Tech Stack:
Golang, Sqlite, Cobra. More information on tech stack will be added with time.

## License
Licensed under MIT. View the LICENSE file for more information.
Attributions will be added soon in ATTRIBUTIONS.md.
