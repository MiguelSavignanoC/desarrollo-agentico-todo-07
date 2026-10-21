# Fecha límite en las tareas

**Issue:** #1 · **Tipo:** feature

## Resumen
Cada tarea puede tener una fecha límite opcional, sin hora, para saber cuándo hay que terminarla. Se fija al crear o editar la tarea, se quita vaciando el campo al editar y se muestra en la lista.

## Qué se construyó
- Campo opcional «Fecha límite» en el formulario de añadir tarea y en el de editar.
- Edición con la fecha precargada; vaciarla la quita; cambiar solo el título la conserva, igual que el estado de completada.
- La lista muestra «Fecha límite: dd/mm/aaaa» bajo el título de las tareas que la tienen.
- Una fecha inválida responde 422 con «La fecha límite no es válida.» y no guarda nada.

## Implementación técnica
| Fichero | Papel |
|---|---|
| `migrations/20261021000000_add_due_date_to_tasks.{up,down}.fizz` | Añade/quita `tasks.due_date` (`date`, nullable). |
| `models/task.go` | `Task.DueDate *time.Time`, `SetDueDate`, `DueDateValue` (aaaa-mm-dd), `DueDateLabel` (dd/mm/aaaa) y la validación. |
| `actions/tasks.go` | `taskForm` común (título y `due_date`) para `TasksCreate` y `TasksUpdate`. |
| `templates/tasks/index.plush.html`, `edit.plush.html` | Campo `<input type="date">` y visualización con `<time datetime>`. |
| `public/assets/app.css` | Estilos nuevos; la hoja sube a `?v=4` en `application.plush.html`. |
| `e2e/fixtures.ts` | Las tareas de prueba admiten `dueDate`. |

Cambios que importan:
- Columna de tipo `date`: no hay hora ni huso, así que el día no se desplaza según el navegador.
- `SetDueDate` trata vacío como «sin fecha». Un texto que no es fecha deja la fecha en nil y marca un campo privado (`dueDateInvalid`) que `Validate` convierte en error de `due_date`. Rechaza también `2026-02-30`.
- Los formularios ya no enlazan directamente a `models.Task` en la creación: se usa `taskForm` y la fecha pasa por `SetDueDate`.
- Las plantillas pasan de `errors.HasAny()` a `errors.Get("title")` / `errors.Get("due_date")`, para que un error de un campo no marque el otro.
- Los dos campos de fecha usan el id `task-due-date`, en pantallas distintas (lista y edición).

## Cómo se usa
1. En la lista, escribe el título, elige opcionalmente una fecha y pulsa «Añadir».
2. La tarea aparece con «Fecha límite: dd/mm/aaaa» bajo el título; sin fecha no muestra nada.
3. Para cambiarla o quitarla, abre la edición de la tarea, modifica o vacía «Fecha límite» y pulsa «Guardar cambios».

## Configuración
Requiere aplicar la migración `20261021000000_add_due_date_to_tasks`; las tareas existentes quedan sin fecha. La migración se deshace con su `down`.

## Cómo se probó
Acta de verificación: [validado, intento 1 de 3](https://github.com/MiguelSavignanoC/desarrollo-agentico-todo-07/issues/1#issuecomment-6320589905).

| Afirmación | Resultado |
|---|---|
| Crear con fecha la muestra igual al recargar y con otro huso horario; la base guarda el día. Sin fecha, queda vacía. | Verificada |
| Editar cambia la fecha, vaciarla la quita; solo el título conserva fecha y completada. | Verificada |
| La lista muestra «Fecha límite: dd/mm/aaaa» sin resaltar ni reordenar. | Verificada |
| Sin regresiones: tareas existentes sin fecha tras migrar, migración reversible. | Verificada |
| Suite completa: 103 tests de Go y 7 de navegador. | Pasa, sin saltar ninguno |

Tests nuevos: `models/task_test.go`, `actions/tasks_test.go` y `e2e/tareas/anadir-una-tarea-con-fecha-limite.spec.ts`, `e2e/tareas/cambiar-y-quitar-la-fecha-limite.spec.ts`. Para repetirlos: `GO_ENV=test go test -p 1 ./...` (con la base de test migrada) y `npx playwright test`, o la skill `check` para todos los gates de la CI.

## Notas
- Fuera de alcance, según la petición: resaltar tareas vencidas, ordenar por fecha, filtros y recordatorios.
- Formato y sitio de la fecha (debajo del título): elección del cambio, pendiente de que el humano la confirme.
- Pendientes señalados en el acta, no resueltos:
  - La ayuda de la edición, «Los demás datos de la tarea no cambian.», queda dudosa porque la misma pantalla ahora edita la fecha.
  - Ante una fecha inválida el campo vuelve vacío, no conserva lo escrito. Solo afecta a clientes sin navegador.
  - El formato `02/01/2006` de `DueDateLabel` está escrito en línea, no en una constante como `DueDateLayout`.
