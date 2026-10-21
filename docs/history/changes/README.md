# Cambios documentados

Cada caso que cierra el workflow `adw` deja aquí su documento de cierre: qué se construyó, cómo se usa, cómo se probó y qué conviene saber para trabajar sobre ello. Lo escribe la skill `document` en la rama del caso, y viaja en el PR con el código.

Este índice es condicional: cada entrada dice cuándo conviene cargar su documento. Se llega a él desde [`docs/README.md`](../../README.md): quien planifica, como las skills `feature`, `bug` y `chore`, lee solo los documentos cuyas condiciones encajan con el encargo.

Un documento de cierre cuenta qué pasó en su caso, y no se reescribe cuando otro caso cambia el mismo código. Lo vigente está en `docs/knowledge/` y `docs/procedures/`.

## Formato de las entradas

```markdown
- [<tipo>-<N>-<slug>.md](<tipo>-<N>-<slug>.md): <qué cambió, en una línea>
  - Cuando se trabaje con <rutas, acciones o plantillas concretas>
  - Cuando se modifique <modelo, campo o regla relacionados>
  - Cuando se investigue <un problema concreto de esa área>
```

Entre tres y cinco condiciones por documento, y las entradas en orden por número de issue. Si un documento ya tiene entrada, se actualizan sus condiciones en lugar de añadir otra.

## Índice

- [feature-1-fecha-limite-tareas.md](feature-1-fecha-limite-tareas.md): fecha límite opcional en las tareas, al crear y editar, visible en la lista
  - Cuando se trabaje con `models/task.go`, `actions/tasks.go` o los formularios de `templates/tasks/`
  - Cuando se modifique la columna `tasks.due_date` o se añada otro campo a `Task` que pase por `taskForm`
  - Cuando se añada vencimiento, orden, filtro o recordatorio basado en la fecha de una tarea
  - Cuando se investigue una fecha que sale desplazada un día o un error de validación que marca el campo equivocado
