# Tasks (portfolio demo)

Minimal task list for a paired Flutter + Go sample: users see tasks immediately and can add or mark them complete. Not a production product.

## Language

**Task**:
A single to-do item with a title and whether it is done.
_Avoid_: Todo, item (in user-facing copy unless UI stays generic)

**Complete** (verb):
The act of marking a Task as done; idempotent if already done.
_Avoid_: Finish, close, archive

**Open Task**:
A Task that is not yet complete.
_Avoid_: Pending, active

**Completed Task**:
A Task that has been marked complete.
_Avoid_: Done task (as a formal term)

**Pair** (repos):
The Flutter client repo and the Go API repo that run together locally for this portfolio demo.
_Avoid_: Monorepo, full stack (as repo name)
