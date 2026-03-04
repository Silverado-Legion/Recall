# Recall
Fast command-line tool for recalling notes and tasks right from the Terminal. Built with Go.<br>
`⚠️ ALPHA STAGE - Core Features are Not Complete.`

This README will be filled out more as project progresses. README will be completed alongside the commit adding the search feature, which will also mark completion of alpha.

## Quick Demo:
```
$ recall add Recall To-Do No.9: Implement a cool TUI. # - This will have an auto-generated title.
Note saved as: Recall To-do No.9: Implement

$ recall add Remember you have an appointment tomorrow! --title "Appointment Tomorrow" # - This has an explicit title.
Note saved as: Appointment Tomorrow

$ recall list # - lists the IDs and titles of all notes, as well as when they were created.
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃  5. Appointment Tomorrow                               [2026-02-27T03:17:51Z]┃
┃  4. Recall To-Do No.9: Implement                       [2026-02-27T03:17:42Z]┃
┃  3. Note content here but in                           [2026-02-27T03:17:23Z]┃
┃  2. Custom Title                                       [2026-02-27T03:17:16Z]┃
┃  1. Note content here.                                 [2026-02-27T03:17:09Z]┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛

$ recall view Appointment Tomorrow # - gets a note by it's exact title. You can also do an ID.
Title: Appointment Tomorrow [5]
...
Remember you have an appointment tomorrow!
...
Created at: [2026-02-27T03:17:51Z]
```
## Usage:
### Add Note:
```
$ recall add Note content here.
$ recall add Note content here again. --title "Custom Title"
$ recall add "Note content here but in quotes so you can put symbols like: #."
```

### List Notes:
```
$ recall list
```

### View Note:
```
$ recall view Title Here (case sensitive, must be exact)
$ recall view [note id]
```

### Complete Note:
```
$ recall done [note id]
```

## To Do:
✅ - Add notes, with automatic titling.<br>
✅ - Write to database.<br>
✅ - List all notes.<br>
✅ - Recall notes from database.<br>
✅ - Mark notes complete.<br>
🔁 - Search for notes.<br>
⏭️ - Make the outputs more "snappy".<br>
⏭️ - Implement a smarter title autogeneration system.<br>
⏭️ - Terminal User Interface (in addition to CLI controls).<br>
⏭️ - Tasks, which are notes with reminders (or time limits).<br>

✅ = Completed.<br>
🔁 = In Progress.<br>
➡️ = Next Up.<br>
⏭️ = Future.<br>

## Tech Stack:
Golang, Sqlite, Cobra. More information on tech stack will be added with time.

## License
Licensed under MIT. View the LICENSE file for more information.
Attributions will be added soon in ATTRIBUTIONS.md.
